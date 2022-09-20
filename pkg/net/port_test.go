package net

import (
	"fmt"
	"testing"
)

func TestScanNet2(t *testing.T) {
	// test: go test pkg/net/* -v -args --site=scanme.nmap.org test
	error := ScanNet()
	if error != nil {
		fmt.Println("error: ", error)
	}
}
