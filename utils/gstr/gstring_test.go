package gstr

import "testing"

func TestDeleteHttp(t *testing.T) {
	url := "http://www.baidu.com/s?ie=utf-8&f=8&rsv_bp=1"
	hh := DeleteHttp(url)

	t.Log(hh)
}
