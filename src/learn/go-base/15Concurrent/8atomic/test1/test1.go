package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	value int64
}

type Result struct {
	job *Job
	sum int64
}

var wg sync.WaitGroup
var jobChan = make(chan *Job, 100)
var resultChan = make(chan *Result, 100)

func producer(job chan<- *Job) {
	defer wg.Done()
	for {
		newJob := &Job{
			value: rand.Int63(),
		}
		job <- newJob
		time.Sleep(time.Millisecond * 500)
	}
}

func consumer(jobChan <-chan *Job, resultChan chan<- *Result) {
	defer wg.Done()
	for {
		job := <-jobChan
		sum := int64(0)
		n := job.value
		for n > 0 {
			sum += n % 10
			n /= 10
		}
		newResult := &Result{
			job: job,
			sum: sum,
		}
		resultChan <- newResult
	}
}

func main() {
	// 1.开启一个goroutine循环生成int64类型的随机数，发送到jobChan
	wg.Add(1)
	go producer(jobChan)

	// 2.开启24个goroutine从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
	wg.Add(24)
	for i := 0; i < 24; i++ {
		go consumer(jobChan, resultChan)
	}

	// //3.主goroutine从resultChan取出结果并打印到终端输出
	for result := range resultChan {
		fmt.Printf("value:%d sum:%d\n", result.job.value, result.sum)
	}

	wg.Wait()
}
