/*
 * benchyou
 * xelabs.org
 *
 * Copyright (c) XeLabs
 * GPL License
 *
 */

package sysbench

import (
	"mybenchx/src/xcommon"
	"mybenchx/src/xworker"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSysbenchQuery(t *testing.T) {
	mysql, cleanup := xcommon.MockMySQL()
	defer cleanup()

	conf := xcommon.MockConf(mysql.Addr())

	workers := xworker.CreateWorkers(conf, 2)
	assert.NotNil(t, workers)

	job := NewQuery(conf, workers)
	job.Run()
	time.Sleep(time.Millisecond * 100)
	job.Stop()
	assert.True(t, job.Rows() > 0)
}

func TestSysbenchQueryRandom(t *testing.T) {
	mysql, cleanup := xcommon.MockMySQL()
	defer cleanup()

	conf := xcommon.MockConf(mysql.Addr())
	conf.Random = true

	workers := xworker.CreateWorkers(conf, 2)
	assert.NotNil(t, workers)

	job := NewQuery(conf, workers)
	job.Run()
	time.Sleep(time.Millisecond * 100)
	job.Stop()
	assert.True(t, job.Rows() > 0)
}

func TestSysbenchNewQuery(t *testing.T) {
	conf := &xcommon.Conf{
		MysqlHost:        "127.0.0.1",
		MysqlUser:        "root",
		MysqlPassword:    "root",
		MysqlPort:        3306,
		MysqlDb:          "sbtest",
		MysqlTableEngine: "innodb",
		OltpTablesCount:  10,
		ReadThreads:      10,
		Random:           false,
		RowsPerInsert:    20,
		BatchPerCommit:   10,
		QueryType:        "time_stamp",
	}

	workers := xworker.CreateWorkers(conf, 2)
	assert.NotNil(t, workers)

	job := NewQuery(conf, workers)
	job.Run()
	time.Sleep(time.Millisecond * 100)
	job.Stop()
	assert.True(t, job.Rows() > 0)
}
