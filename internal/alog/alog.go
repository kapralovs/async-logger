package alog

import (
	"fmt"
	"sync"
	"time"
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
		header := formatHeaderOutput(time.Now())
		fmt.Printf("%s %s\n", header, fmt.Sprint(v...))
	}()
	l.wg.Wait()
}

func formatHeaderOutput(t time.Time) string {
	header := ""
	year, month, day := t.Date()
	formattedMonth := itoa(int(month))
	formattedDay := itoa(day)
	hours, minutes, seconds := t.Clock()
	formattedHours := itoa(hours)
	formattedMinutes := itoa(minutes)
	formattedSeconds := itoa(seconds)
	date := fmt.Sprintf("%d/%s/%s", year, formattedMonth, formattedDay)
	clock := fmt.Sprintf("%s:%s:%s", formattedHours, formattedMinutes, formattedSeconds)
	header = fmt.Sprintf("%s %s", date, clock)
	return header
}

func itoa(value int) string {
	strValue := ""
	if value < 10 {
		strValue = fmt.Sprintf("0%d", value)
	} else {
		strValue = fmt.Sprint(value)
	}
	return strValue
}
