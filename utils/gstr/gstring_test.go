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

func TestDeleteHttp(t *testing.T) {
	url := "http://www.baidu.com/s?ie=utf-8&f=8&rsv_bp=1"
	hh := DeleteHttp(url)

	t.Log(hh)
}
