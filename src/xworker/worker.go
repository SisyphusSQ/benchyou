/*
 * benchyou
 * xelabs.org
 *
 * Copyright (c) XeLabs
 * GPL License
 *
 */

package xworker

import (
	"mybenchx/src/xcommon"
	"fmt"
	"gorm.io/gorm/logger"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Metric tuple.
type Metric struct {
	WNums  uint64
	WCosts uint64
	WMax   uint64
	WMin   uint64
	QNums  uint64
	QCosts uint64
	QMax   uint64
	QMin   uint64
}

// Worker tuple.
type Worker struct {
	// session
	S *gorm.DB

	// mertric
	M *Metric

	// engine
	E string

	// xid
	XID string

	// table number
	N int
}

// CreateWorkers creates the new workers.
func CreateWorkers(conf *xcommon.Conf, threads int) []Worker {
	var workers []Worker
	var conn *gorm.DB
	var err error

	//dsn := fmt.Sprintf("%s:%d", conf.MysqlHost, conf.MysqlPort)
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.MysqlUser, conf.MysqlPassword, conf.MysqlHost, conf.MysqlPort, conf.MysqlDb)
	for i := 0; i < threads; i++ {

		if conn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		}); err != nil {
			log.Panicf("create.worker.error:%v", err)
		}

		workers = append(workers, Worker{
			S: conn,
			M: &Metric{},
			E: conf.MysqlTableEngine,
			N: conf.OltpTablesCount,
		},
		)
	}
	return workers
}

// AllWorkersMetric returns all the worker's metric.
func AllWorkersMetric(workers []Worker) *Metric {
	all := &Metric{}
	for _, worker := range workers {
		all.WNums += worker.M.WNums
		all.WCosts += worker.M.WCosts

		if all.WMax < worker.M.WMax {
			all.WMax = worker.M.WMax
		}

		if all.WMin > worker.M.WMin {
			all.WMin = worker.M.WMin
		}

		all.QNums += worker.M.QNums
		all.QCosts += worker.M.QCosts

		if all.QMax < worker.M.QMax {
			all.QMax = worker.M.QMax
		}

		if all.QMin > worker.M.QMin {
			all.QMin = worker.M.QMin
		}
	}

	return all
}

// StopWorkers used to stop all the worker.
func StopWorkers(workers []Worker) {
	for _, worker := range workers {
		sqlDB, err := worker.S.DB()
		sqlDB.Close()

		if err != nil {
			log.Panicf("close.worker.error:%v", err)
		}
	}
}
