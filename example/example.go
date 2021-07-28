package main

import (
	"fmt"
	"github.com/spl0i7/pipeline"
	"log"
	"time"
)

type MultiplyTenSlow struct{}

func (m MultiplyTenSlow) Process(result pipeline.Message) ([]pipeline.Message, error) {
	time.Sleep(1* time.Second)
	number := result.(int)
	return []pipeline.Message{number * 10, number * 10}, nil
}

type MultiplyHundredSlow struct{}

func (m MultiplyHundredSlow) Process(result pipeline.Message) ([]pipeline.Message, error) {
	time.Sleep(time.Duration(1 * time.Second))
	number := result.(int)
	return []pipeline.Message{number * 100, number * 100}, nil
}

type DivideThreeSlow struct{}

func (m DivideThreeSlow) Process(result pipeline.Message) ([]pipeline.Message, error) {
	time.Sleep(time.Duration(1 * time.Second))
	number := result.(int)
	return []pipeline.Message{number / 3}, nil
}


func main() {

	p := pipeline.NewConcurrentPipeline()

	p.AddPipe(MultiplyHundredSlow{}, &pipeline.PipelineOpts{
		Concurrency: 5,
	})
	p.AddPipe(MultiplyTenSlow{}, &pipeline.PipelineOpts{
		Concurrency: 5,
	})
	p.AddPipe(DivideThreeSlow{}, &pipeline.PipelineOpts{
		Concurrency: 5,
	})

	if err := p.Start(); err != nil {
		log.Println(err)
	}

	for i := 1; i <= 3; i++ {
		p.Input() <- i
	}




	go func() {
		count := 0
		for number := range p.Output() {
			fmt.Println(number)
			count++
		}
	}()

	p.Stop()

}