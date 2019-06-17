package utils

import (
	"fmt"
	"testing"
	"time"
)

func TestGenNo(t *testing.T) {
	for i:=0; i<10; i++  {
		fmt.Println(GenNo())
	}
}

func TestTradeNo(t *testing.T) {
	for i:=0; i<10; i++  {
		t1 := time.Now().UnixNano()
		fmt.Println(CreateTradeNo(12))
		fmt.Println(time.Now().UnixNano() - t1)
	}
}

func TestGenTrxNo(t *testing.T) {
	for i:=0; i<10; i++  {
		t1 := time.Now().UnixNano()
		fmt.Println(GenTrxNo())
		fmt.Println(time.Now().UnixNano() - t1)
	}
}

func TestGenPaySecret(t *testing.T) {
	for i:=0; i<10; i++  {
		t1 := time.Now().UnixNano()
		fmt.Println(GenPaySecret())
		fmt.Println(time.Now().UnixNano() - t1)
	}
}

func TestGenPayKey(t *testing.T) {
	for i:=0; i<10; i++  {
		t1 := time.Now().UnixNano()
		fmt.Println(GenPayKey())
		fmt.Println(time.Now().UnixNano() - t1)
	}
}
