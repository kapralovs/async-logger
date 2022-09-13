package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/kapralovs/async-logger/internal/alog"
)

func main() {
	numOfThreads, _ := strconv.Atoi(os.Args[1])
	numOfLogs, _ := strconv.Atoi(os.Args[2])
	asyncLogger := alog.New()
	wg := &sync.WaitGroup{}
	for i := 0; i < numOfThreads; i++ {
		wg.Add(1)
		go func(threadNum int) {
			defer wg.Done()
			fmt.Println("Do some work BEFORE logs")
			for j := 0; j < numOfLogs; j++ {
				asyncLogger.Println(time.Now(), " - Thread: ", threadNum, ",Log: ", j+1)
				fmt.Println("Do some work RIGHT AFTER log")
			}
			fmt.Println("Do some work AFTER logs")
		}(i + 1)
	}
	wg.Wait()
	asyncLogger.Println("Program finished")
}
