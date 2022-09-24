package main

import (
	"fmt"
	"time"
)

func main() {

	// 课后练习 1.1
	// 编写一个小程序：
	// 给定一个字符串数组
	// [“I”,“am”,“stupid”,“and”,“weak”]
	// 用 for 循环遍历该数组并修改为
	// [“I”,“am”,“smart”,“and”,“strong”]
	var strs = [5]string{"i", "am", "stupid", "and", "weak"}
	strs = convertStr(strs)
	fmt.Println(strs)

	// 课后练习 1.2
	// 基于 Channel 编写一个简单的单线程生产者消费者模型：
	//
	// * 队列：
	// 队列长度 10，队列元素类型为 int
	// * 生产者：
	// 每 1 秒往队列中放入一个类型为 int 的元素，队列满时生产者可以阻塞
	// * 消费者：
	// 每一秒从队列中获取一个元素并打印，队列为空时消费者阻塞

	ch := make(chan int, 10)
	go func() {
		produce(ch)
	}()

	go consumer(ch, "consumer1")
	consumer(ch, "consumer2")

}

func convertStr(strs [5]string) (strs2 [5]string) {
	for i, str := range strs {
		if str == "stupid" {
			strs[i] = "smart"
		} else if str == "weak" {
			strs[i] = "strong"
		}
	}
	return strs
}

func produce(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
		time.Sleep(1 * time.Second)
	}
	close(ch)
}

func consumer(ch chan int, consumerName string) {
	for value := range ch {
		fmt.Printf("consumerName: %s, value: %d\n", consumerName, value)
	}
}
