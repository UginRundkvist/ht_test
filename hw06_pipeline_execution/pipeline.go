package hw06pipelineexecution

import (
	"fmt"
	"time"
)

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	start := time.Now()
	time.Sleep(1000)
	new1 := make(In)
	new2 := make(In)
	new3 := make(In)
	select {
	case <-done:
		return make(In)
	default:
		new1 = stages[0](in)
	}

	select {
	case <-done:
		return make(In)
	default:
		new2 = stages[1](new1)
	}

	select {
	case <-done:
		return make(In)
	default:
		new3 = stages[2](new2)
	}

	select {
	case <-done:
		return make(In)
	default:
		elapsed := time.Since(start)
		fmt.Println("Время выполнения: ", elapsed)
		return (stages[3](new3))
	}

}
