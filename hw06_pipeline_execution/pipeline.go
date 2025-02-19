package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	for _, stage := range stages {
		in = execStage(in, done, stage)
	}
	return in
}

func execStage(in In, done In, stage Stage) In {
	go func() {
		select {
		case <-done:
		default:
			in = stage(in)
		}
	}()
	return in
}
