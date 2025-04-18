package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalWorld     = "Go!"
	countdownStart = 3
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct {
}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}

	fmt.Fprint(out, finalWorld)
}

func main() {
	Countdown(os.Stdout, &DefaultSleeper{})
}
