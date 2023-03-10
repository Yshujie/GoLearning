package goroutineTest

import (
	"fmt"
	"time"
)

func Running() {

	var times int

	// 构建一个无限循环
	for {
		times++

		fmt.Println("tick: ", times)

		time.Sleep(time.Second)
	}

}
