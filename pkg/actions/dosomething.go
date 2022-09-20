package actions

import (
	"fmt"
	"time"
)

func DoSomething() error {
	fmt.Println("Doing...👀")
	time.Sleep(time.Second * 2)
	fmt.Println("Done!")
	return nil
}
