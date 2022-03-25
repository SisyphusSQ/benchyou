/*
 * GPL License
 *
 * mybenchx
 * revised by alex.zhao @2022 Spring
 *
 * github.com/SisyphusSQ/mybenchx
 *
 */

package sysbench

import (
	"mybenchx/src/xcommon"
	"mybenchx/src/xworker"
	"fmt"
	"log"
	"math"
	"sync"
	"sync/atomic"
	"time"

	"github.com/xelabs/go-mysqlstack/sqlparser/depends/common"
)

// PreInsert tuple.
type PreInsert struct {
	stop     bool
	requests uint64
	conf     *xcommon.Conf
	workers  []xworker.Worker
	lock     sync.WaitGroup
}

// NewPreInsert creates the new insert handler.
func NewPreInsert(conf *xcommon.Conf, workers []xworker.Worker) xworker.Handler {
	return &PreInsert{
		conf:    conf,
		workers: workers,
	}
}

// Run used to start the worker.
func (preInsert *PreInsert) Run() {
	threads := len(preInsert.workers)
	for i := 0; i < threads; i++ {
		preInsert.lock.Add(1)
		go preInsert.Insert(&preInsert.workers[i], threads, i)
	}
}

// Stop used to stop the worker.
func (preInsert *PreInsert) Stop() {
	preInsert.stop = true
	preInsert.lock.Wait()
}

// Rows returns the row numbers.
func (preInsert *PreInsert) Rows() uint64 {
	return atomic.LoadUint64(&preInsert.requests)
}

// Insert used to execute the insert query.
func (preInsert *PreInsert) Insert(worker *xworker.Worker, num, id int) {
	var buf *common.Buffer
	session := worker.S
	bs := int32(math.MaxInt32) / int32(num)
	lo := bs * int32(id)
	hi := bs * int32(id+1)
	columns1 := "k,c,pad,created_at,unix_stamp"
	valfmt1 := "(%v,'%s', '%s', '%s', %d),"

	t := time.Now()

	for i := 0; i <= preInsert.conf.OltpTableSize; i++ {
		if i == 0 {
			buf = common.NewBuffer(256)
			continue
		}

		var sql, value string

		sql = fmt.Sprintf("insert into mybenchx%d(%s) values", id, columns1)

		pad := xcommon.RandString(xcommon.Padtemplate)
		c := xcommon.RandString(xcommon.Ctemplate)

		Loc, _ := time.LoadLocation("Asia/Shanghai")
		unixStamp := time.Now().Unix() + int64(i)
		createdAt := fmt.Sprintf(time.Unix(unixStamp, 0).In(Loc).Format("2006-01-02 15:04:05"))

		value = fmt.Sprintf(valfmt1,
			xcommon.RandInt32(lo, hi),
			c,
			pad,
			createdAt,
			int32(unixStamp),
		)
		buf.WriteString(value)

		if i%3000 == 0 || preInsert.conf.OltpTableSize-i == 0 {
			// -1 to trim right ','
			vals, err := buf.ReadString(buf.Length() - 1)
			if err != nil {
				log.Panicf("preInsert.error[%v]", err)
			}
			sql += vals

			if err = session.Exec(sql).Error; err != nil {
				log.Panicf("preInsert.error[%v]", err)
			}

			buf = common.NewBuffer(256)

			elapsedLocal := time.Since(t)
			log.Printf("[PROCESS] Table: mybenchx%d insert num: %d cost time: %s sec\n", id, i, fmt.Sprintf("%+v", elapsedLocal.Seconds()))
		}
	}

	elapsed := time.Since(t)
	log.Printf("[END] Table: mybenchx%d cost time: %s sec\n", id, fmt.Sprintf("%+v", elapsed.Seconds()))
	preInsert.lock.Done()
}
