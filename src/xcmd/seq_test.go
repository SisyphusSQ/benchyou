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
	"benchyou/src/xcommon"
	"testing"
)

func TestXcmdSeq(t *testing.T) {
	mysql, cleanup := xcommon.MockMySQL()
	defer cleanup()

	cmd := NewSeqCommand()
	MockInitFlags(cmd, mysql.Addr())
	seqCommandFn(cmd, nil)
}
