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
		// now := time.Now()
		// year, month, day := now.Date()
		// formattedMonth := itoa(int(month))
		// formattedDay := itoa(day)
		// hours, minutes, seconds := now.Clock()
		// formattedHours := itoa(hours)
		// formattedMinutes := itoa(minutes)
		// formattedSeconds := itoa(seconds)
		// date := fmt.Sprintf("%d/%s/%s", year, formattedMonth, formattedDay)
		// clock := fmt.Sprintf("%s:%s:%s", formattedHours, formattedMinutes, formattedSeconds)
		// fmt.Printf("%s %s %s\n", date, clock, fmt.Sprint(v...))
	}()
	l.wg.Wait()
}

func formatHeaderOutput(t time.Time) string {
	formatted := ""
	year, month, day := t.Date()
	formattedMonth := itoa(int(month))
	formattedDay := itoa(day)
	hours, minutes, seconds := t.Clock()
	formattedHours := itoa(hours)
	formattedMinutes := itoa(minutes)
	formattedSeconds := itoa(seconds)
	date := fmt.Sprintf("%d/%s/%s", year, formattedMonth, formattedDay)
	clock := fmt.Sprintf("%s:%s:%s", formattedHours, formattedMinutes, formattedSeconds)
	fmt.Printf("%s %s %s\n", date, clock, fmt.Sprint(v...))
	//TODO: Дописать логику форматирования хэдера
	return formatted
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
