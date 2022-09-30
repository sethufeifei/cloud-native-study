package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Queue struct {
	queue []string
	cond  *sync.Cond
}

func main() {

	q := Queue{
		queue: []string{},
		cond:  sync.NewCond(&sync.Mutex{}),
	}

	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			for {
				time.Sleep(time.Second * time.Duration(rand.Intn(10)))
				q.produce("a", index)
			}
		}(i)
	}

	for i := 0; i < 7; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			for {
				time.Sleep(time.Second * 3)
				q.consumer(index)
			}
		}(i)
	}

	wg.Wait()

}

func (q *Queue) produce(item string, index int) {

	q.cond.L.Lock()
	defer q.cond.L.Unlock()

	q.queue = append(q.queue, item)
	fmt.Printf("[produce-%d] data = %s \n", index, item)
	q.cond.Signal()
}

func (q *Queue) consumer(index int) string {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()

	// 存在虚假唤醒
	for len(q.queue) <= 0 {
		fmt.Printf("[consumer-%d] no data available, wait \n", index)
		q.cond.Wait()
	}

	result := q.queue[0]
	fmt.Printf("[consumer-%d] data = %s \n", index, result)
	q.queue = q.queue[1:]
	return result
}
