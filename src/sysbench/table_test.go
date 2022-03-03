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
	"benchyou/src/xcommon"
	"benchyou/src/xworker"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSysbenchTable(t *testing.T) {
	mysql, cleanup := xcommon.MockMySQL()
	defer cleanup()

	conf := xcommon.MockConf(mysql.Addr())

	workers := xworker.CreateWorkers(conf, 2)
	assert.NotNil(t, workers)

	job := NewTable(workers)
	job.Prepare()
	job.Cleanup()
}

func TestPrepareTable(t *testing.T) {
	conf := &xcommon.Conf{
		MysqlHost:        "127.0.0.1",
		MysqlUser:        "root",
		MysqlPassword:    "root",
		MysqlPort:        3306,
		MysqlDb:          "sbtest",
		MysqlTableEngine: "innodb",
	}

	// worker
	workers := xworker.CreateWorkers(conf, 1)
	table := NewTable(workers)
	table.Prepare()
}

func TestCleanUpTable(t *testing.T) {
	conf := &xcommon.Conf{
		MysqlHost:        "127.0.0.1",
		MysqlUser:        "root",
		MysqlPassword:    "root",
		MysqlPort:        3306,
		MysqlDb:          "sbtest",
		MysqlTableEngine: "innodb",
		OltpTablesCount:  64,
	}

	// worker
	workers := xworker.CreateWorkers(conf, 1)
	table := NewTable(workers)
	table.Cleanup()
}
