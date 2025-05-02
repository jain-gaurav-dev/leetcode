package daily_challenge

import (
	"testing"
)

func TestPushDomino(t *testing.T) {
	res := pushDominoes("RR.L")
	t.Logf("res is %s\n", res)
}
