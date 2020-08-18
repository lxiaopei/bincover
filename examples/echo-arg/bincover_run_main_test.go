// +build testbincover

package main

import (
	"testing"

	"github.com/lxiaopei/bincover"
)

func TestBincoverRunMain(t *testing.T) {
	bincover.RunTest(main)
}
