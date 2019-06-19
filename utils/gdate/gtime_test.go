package gdate

import (
	"testing"
)

func TestCurrentDay(t *testing.T) {
	test := TimeFormat(Day24())
	println(test)
}
