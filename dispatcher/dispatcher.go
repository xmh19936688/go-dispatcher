// Package dispatcher Usage:
//
// wg := &sync.WaitGroup{}
// ch := make(chan T)
// handler := func(data T) {}
// d := dispatcher.New[T]().
//
//	MaxCurrency(-1).
//	Handler(handler).
//	Chan(ch).
//	WaitGroup(wg)
//
// go d.Start()
//
// wg.Add(len(list))
// for _, data := range lsit { ch <- data }
// wg.Wait()
package dispatcher

import (
	"math"
	"sync"
)

type dispatcher[T any] struct {
	// 限定最大并发数
	maxCurrency int
	// 处理数据的func
	handler func(T)
	// 接收数据的管道
	ch chan T
	// 用于标记完成的wg
	wg *sync.WaitGroup
}

func New[T any]() *dispatcher[T] {
	return &dispatcher[T]{}
}

func (d *dispatcher[T]) MaxCurrency(value int) *dispatcher[T] {
	if value <= 0 {
		value = math.MaxInt
	}

	d.maxCurrency = value
	return d
}

func (d *dispatcher[T]) Handler(f func(T)) *dispatcher[T] {
	d.handler = f
	return d
}

func (d *dispatcher[T]) Chan(ch chan T) *dispatcher[T] {
	d.ch = ch
	return d
}

func (d *dispatcher[T]) WaitGroup(wg *sync.WaitGroup) *dispatcher[T] {
	d.wg = wg
	return d
}

// 调用前加go关键字
func (d *dispatcher[T]) Start() {
	if d.ch == nil {
		return
	}
	if d.handler == nil {
		return
	}
	if d.maxCurrency <= 0 {
		d.maxCurrency = math.MaxInt
	}

	for i := 0; i < d.maxCurrency; i++ {
		go func() {
			for {
				d.handler(<-d.ch)
				if d.wg != nil {
					d.wg.Done()
				}
			}
		}()
	}
}

// 串行执行，仅供测试
func (d *dispatcher[T]) Run() {
	for {
		d.handler(<-d.ch)
		if d.wg != nil {
			d.wg.Done()
		}
	}
}
