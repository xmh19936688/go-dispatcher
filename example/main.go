package main

import (
	"fmt"
	"sync"

	"github.com/xmh19936688/go-dispatcher/dispatcher"
)

type Job struct {
	name string
}

var list = [...]Job{
	{name: "0"},
	{name: "1"},
	{name: "2"},
	{name: "3"},
	{name: "4"},
	{name: "5"},
	{name: "6"},
	{name: "7"},
	{name: "8"},
	{name: "9"},
	{name: "a"},
	{name: "b"},
	{name: "c"},
	{name: "d"},
	{name: "e"},
	{name: "f"},
}

func handler(job Job) {
	fmt.Println(job.name)
}

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan Job)

	d := dispatcher.New[Job]().MaxCurrency(-1).
		Handler(handler).Chan(ch).WaitGroup(wg)
	go d.Start()

	wg.Add(len(list))
	for _, data := range list {
		ch <- data
	}
	wg.Wait()
}
