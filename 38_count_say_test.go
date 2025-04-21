package daily_challenge

import "testing"

func TestCountSay(t *testing.T) {
	res := countAndSay(4)
	t.Logf("res is %s\n", res)
}
