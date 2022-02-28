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

func TestXcmdRandom(t *testing.T) {
	mysql, cleanup := xcommon.MockMySQL()
	defer cleanup()

	cmd := NewRandomCommand()
	MockInitFlags(cmd, mysql.Addr())
	randomCommandFn(cmd, nil)
}
