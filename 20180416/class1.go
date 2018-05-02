package main

import "fmt"

type Log struct {
}

func (log *Log) print(str string) {
	fmt.Println("Here we have", str)
}

type Job struct {
	Log
	Command string
}

func (j *Job) print() {
	j.Log.print(j.Command)
}

func main() {
	my_job := &Job{Command: "My First Job"}
	my_job.print()
}
