package myjob

import(
	"fmt"
	"time"
)

type jobFunc func() string

func Run(intervalSecond int, job jobFunc){
	fmt.Printf("start loop...\n")

	loopIndex := 1

	for{
		fmt.Printf("loop: %d\n", loopIndex)
		res := job()
		fmt.Printf(res)
		loopIndex ++
		time.Sleep( time.Duration(intervalSecond) * time.Second )
	}
}



