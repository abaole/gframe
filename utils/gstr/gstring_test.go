package gstr

import "testing"

func TestWxRate(t *testing.T) {
	te1 := 0.60
	r1 := WxRate(te1)
	t.Log(r1)

	te2 := 0.39
	r2 := WxRate(te2)
	t.Log(r2)
}
