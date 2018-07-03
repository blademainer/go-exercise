package main

import (
	"fmt"
	"math/rand"
	"time"
)

type (
	ProducerTask struct {
		Name string
		Url  string
	}

	TaskData struct {
		data string
	}

	Producer interface {
		Produce() *TaskData
	}
)

func (*ProducerTask) Produce() *TaskData {
	time.Sleep(100 * time.Millisecond)
	return &TaskData{}
}

var restartChan = make(chan bool)

func restart() {
	restartChan <- true
}

func reloadConfigs() []*ProducerTask {
	size := rand.Intn(100) + 100
	producerTasks := make([]*ProducerTask, 0)
	for i := 0; i < size; i++ {
		producerTasks = append(producerTasks, &ProducerTask{})
	}
	return producerTasks
}

func start(incomingMessages chan *TaskData) {
	go func() {
		for {
			configs := reloadConfigs()
			fmt.Println("Starting tasks....")
			stopped := false
			exited := make(chan bool)
			size := len(configs)
			for i, producer := range configs {
				go func(index int) {
					for !stopped {
						select {
						default:
							taskData := producer.Produce()
							incomingMessages <- taskData
						}
					}
					fmt.Printf("Producer: %v %d exited!! \n!", producer, index)
					exited <- true
				}(i)

			}
			fmt.Println("Started tasks and wait for done....")
			// waiting to restart
			<-restartChan
			fmt.Println("Processing restart events!")
			stopped = true
			for i := 0; i < size; i++ {
				<-exited
			}
			fmt.Println("All exited!")
		}
	}()
}

func main() {
	incomingMessages := make(chan *TaskData)
	start(incomingMessages)
	go func() {
		for {
			data := <-incomingMessages
			fmt.Println("Got data: ", data)
		}
	}()
	time.Sleep(1 * time.Second)
	restart()
	time.Sleep(1 * time.Second)
	restart()
	time.Sleep(1 * time.Second)
}
