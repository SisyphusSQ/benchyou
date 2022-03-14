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
	"mybenchx/src/xworker"
	"fmt"
	"log"
)

// Table tuple.
type Table struct {
	workers []xworker.Worker
}

// NewTable creates the new table.
func NewTable(workers []xworker.Worker) *Table {
	return &Table{workers}
}

// Prepare used to prepare the tables.
func (t *Table) Prepare() {
	session := t.workers[0].S
	count := t.workers[0].N
	engine := t.workers[0].E

	for i := 0; i < count; i++ {
		sql := fmt.Sprintf(`create table mybenchx%d (
							id bigint unsigned not null auto_increment,
							k int not null default '0',
							c varchar(120) not null default '',
							pad varchar(60) not null default '',
							created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
							unix_stamp bigint not null default '0',
							primary key (id),
							key idx_k_1 (k),
							key idx_created_at (created_at),
							key idx_unix_stamp (unix_stamp)
							) engine=%s`, i, engine)

		if err := session.Exec(sql).Error; err != nil {
			log.Panicf("creata.table.error[%v]", err)
		}
		log.Printf("create table mybenchx%d(engine=%v) finished...\n", i, engine)
	}
}

// Cleanup used to cleanup the tables.
func (t *Table) Cleanup() {
	session := t.workers[0].S
	//count := t.workers[0].N

	// for test
	count := 64

	for i := 0; i < count; i++ {
		sql := fmt.Sprintf(`drop table mybenchx%d;`, i)

		if err := session.Exec(sql).Error; err != nil {
			log.Panicf("drop.table.error[%v]", err)
		}
		log.Printf("drop table mybenchx%d finished...\n", i)
	}
}
