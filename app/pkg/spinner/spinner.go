package spinner

import (
	"fmt"
	"time"
)

type Spinner struct {
	symbols []string
	index   int
	done    chan bool
}

func NewSpinner() *Spinner {
	spinner := &Spinner{
		symbols: []string{"|", "/", "-", "\\"},
		index:   0,
		done:    make(chan bool),
	}

	go spinner.start()
	return spinner
}

func (s *Spinner) start() {
	for {
		select {
		case <-s.done:
			return
		default:
			fmt.Printf("\r%s ", s.symbols[s.index])
			s.index = (s.index + 1) % len(s.symbols)
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func (s *Spinner) Stop() {
	s.done <- true
	fmt.Printf("\rDone!  \n")
}
