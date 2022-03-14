/*
 * benchyou
 * xelabs.org
 *
 * Copyright (c) XeLabs
 * GPL License
 *
 */

package xcmd

import (
	"mybenchx/src/sysbench"
	"mybenchx/src/xworker"
	"github.com/spf13/cobra"
)

// NewPrepareCommand creates the new cmd.
func NewPrepareCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "prepare",
		Run: prepareCommandFn,
	}

	return cmd
}

func prepareCommandFn(cmd *cobra.Command, args []string) {
	conf, err := parseConf(cmd)
	if err != nil {
		panic(err)
	}

	// worker
	workers := xworker.CreateWorkers(conf, 1)
	table := sysbench.NewTable(workers)
	table.Prepare()

	// fill up table with oltp-table-size
	if conf.OltpTableSize >= 0 {
		preWorkers := xworker.CreateWorkers(conf, conf.OltpTablesCount)
		preInsert := sysbench.NewPreInsert(conf, preWorkers)
		preInsert.Run()

		// wait for complete
		preInsert.Stop()
	}
}

// NewCleanupCommand creates the new cmd.
func NewCleanupCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "cleanup",
		Run: cleanupCommandFn,
	}

	return cmd
}

func cleanupCommandFn(cmd *cobra.Command, args []string) {
	conf, err := parseConf(cmd)
	if err != nil {
		panic(err)
	}

	// worker
	workers := xworker.CreateWorkers(conf, 1)
	table := sysbench.NewTable(workers)
	table.Cleanup()
}
