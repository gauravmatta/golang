// This sample program demonstrates how to use a buffered channel to work on multiple taskctl
//with a predefined number of goroutines.

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Task struct {
	Id        int
	JobId     int
	Status    string
	CreatedOn time.Time
}

func (t *Task) Run() {

	sleep := rand.Int63n(1000)
	// Delaying the execution for the sake of example
	time.Sleep(time.Duration(sleep) * time.Millisecond)
	t.Status = "Completed"
}

// wg is used to wait for the program to finish.
var wg sync.WaitGroup

// Arbitrary number of workers for the sake of example It can be big in a real-world scenario
const noOfWorkers = 3

func main() {
	// Create a buffered channel to manage the task queue.
	taskQueue := make(chan *Task, 10)

	// Launch goroutines to handle the work.
	// Since we have 2 workers we will launch 2 goroutines to process the tasks from the channel.
	//The worker process is distributing with the value of noOfWorkers.
	wg.Add(noOfWorkers * 2)
	for gr := 1; gr <= noOfWorkers; gr++ {
		go workerForRange(taskQueue, gr)
		go workerFor(taskQueue, gr)
	}

	// Add Tasks into Buffered channel.
	for i := 1; i <= 10; i++ {
		taskQueue <- &Task{
			Id:        i,
			JobId:     100 + i,
			CreatedOn: time.Now(),
		}
	}

	// Once all the tasks are added into the channel, we can close the channel to indicate that no more tasks will be added.
	close(taskQueue)
	wg.Wait()
}

func workerFor(taskQueue <-chan *Task, workerId int) {
	defer wg.Done()
	for {
		// Read the value from channel
		v, exists := <-taskQueue
		if !exists {
			// channel was closed and that exit from goroutine
			return
		}
		fmt.Printf("WorkerFor%d: received request for Task:%d - Job:%d\n", workerId, v.Id, v.JobId)
		v.Run()
		// Display we finished the work.
		fmt.Printf("WorkerFor%d: Status:%s for Task:%d - Job:%d\n", workerId, v.Status, v.Id, v.JobId)
	}
	/* Three behaviors of v, exists := <-taskQueue
	   1: exists==true
		Data received from the channel.
	   2: exists==false
		Channel was closed, no more data left in the channel
		Receive default value
	   3: Blocking code because send is not happened
		  No data available in the channel
		  Waiting for send operation */
}

// worker is launched as a goroutine to process Tasks from the buffered channel.
func workerForRange(taskQueue <-chan *Task, workerId int) {
	// Schedule the call to Done method of WaitGroup.
	defer wg.Done()
	for v := range taskQueue {
		fmt.Printf("WorkerForRange%d: received request for Task:%d - Job:%d\n", workerId, v.Id, v.JobId)
		v.Run()
		// Display we finished the work.
		fmt.Printf("WorkerForRange%d: Status:%s for Task:%d - Job:%d\n", workerId, v.Status, v.Id, v.JobId)
	}
}
