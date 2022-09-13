package alog

import (
	"fmt"
	"sync"
)

type Logger struct {
	wg *sync.WaitGroup
	mu *sync.Mutex
}

func New() *Logger {
	return &Logger{
		wg: &sync.WaitGroup{},
		mu: &sync.Mutex{},
	}
}

func (l *Logger) Println(v ...any) {
	defer l.mu.Unlock()
	l.mu.Lock()
	l.wg.Add(1)
	go func() {
		defer l.wg.Done()
		fmt.Println(v...)
	}()
	l.wg.Wait()
}
