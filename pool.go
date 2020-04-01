package main

import (
	"fmt"
	"math/rand"
	"runtime"
)

type Job struct {
	Id          int
	RuseltMatch int
}

type Result struct {
	job *Job
	sum int
}

func main() {
	runtime.GOMAXPROCS(4)
	jobChan := make(chan *Job, 128)
	resultChan := make(chan *Result, 128)

	createPool(64, jobChan, resultChan)

	go func(resultChan chan *Result) {
		for result := range resultChan {
			fmt.Printf("job id:%v randnum: %v result: %d\n", result.job.Id, result.job.RuseltMatch, result.sum)
		}
	}(resultChan)
	var id int = 0
	for {
		id++
		r_num := rand.Int()
		job := &Job{
			Id:          id,
			RuseltMatch: r_num,
		}

		jobChan <- job
		if job.Id == 100 {
			fmt.Println(job.Id)
			break
		}
	}
}

func createPool(num int, jobChan chan *Job, resultChan chan *Result) {
	for i := 0; i < num; i++ {
		go func(jobChan chan *Job, resultChan chan *Result) {
			for job := range jobChan { // 使用range 避免pince
				r_num := job.RuseltMatch

				var sum int
				for r_num != 0 {
					temp := r_num % 10
					sum += temp
					r_num /= 10
				}
				r := &Result{
					job: job,
					sum: sum,
				}

				resultChan <- r
			}
		}(jobChan, resultChan)
	}
}
