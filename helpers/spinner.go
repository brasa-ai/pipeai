package helpers

import (
	"fmt"
	"time"
)

type Spinner struct {
	chars     []string
	stop      chan struct{}
	stopped   bool
	startTime time.Time
}

func NewSpinner() *Spinner {
	return &Spinner{
		chars: []string{"⠋", "⠙", "⠚", "⠞", "⠖", "⠦", "⠧", "⠇", "⠏", "⠍"},
		stop:  make(chan struct{}),
	}
}

func (s *Spinner) Start() {
	if s.stopped {
		return
	}
	s.startTime = time.Now()
	go func() {
		i := 0
		for {
			select {
			case <-s.stop:
				fmt.Print("\r\033[K") // Clear the line
				return
			default:
				elapsed := time.Since(s.startTime)
				timeStr := formatDuration(elapsed)
				fmt.Printf("\r%s %s", s.chars[i], timeStr)
				time.Sleep(100 * time.Millisecond)
				i = (i + 1) % len(s.chars)
			}
		}
	}()
}

func (s *Spinner) Stop() {
	if s.stopped {
		return
	}
	s.stopped = true
	s.stop <- struct{}{}
	close(s.stop)
}

func formatDuration(d time.Duration) string {
	totalMs := d.Milliseconds()

	if totalMs < 1000 {
		return fmt.Sprintf("%dms", totalMs)
	}

	h := totalMs / (1000 * 60 * 60)
	m := (totalMs % (1000 * 60 * 60)) / (1000 * 60)
	s := (totalMs % (1000 * 60)) / 1000
	ms := totalMs % 1000

	if h > 0 {
		return fmt.Sprintf("%dh/%dm", h, m)
	}
	if m > 0 {
		return fmt.Sprintf("%dm/%ds", m, s)
	}
	return fmt.Sprintf("%ds/%dms", s, ms)
}
