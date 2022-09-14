package main

import (
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/kapralovs/alog_test/alog"
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
			for j := 0; j < numOfLogs; j++ {
				asyncLogger.Println("Thread: ", threadNum, ",Log: ", j+1)
			}
		}(i + 1)
	}
	wg.Wait()
	log.Println("Program finished")
}
