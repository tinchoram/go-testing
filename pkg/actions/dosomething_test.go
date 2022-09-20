package actions

import (
	"fmt"
	"testing"
)

func TestDoSomething(t *testing.T) {

	error := DoSomething()
	if error != nil {
		fmt.Println("error:", error)
	}
}
