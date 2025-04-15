package main

import (
	"bytes"
	"reflect"
	"testing"
)

type SpyCountdownOperations struct {
	calls []string
}

const (
	write = "write"
	sleep = "sleep"
)

func (s *SpyCountdownOperations) Sleep() {
	s.calls = append(s.calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.calls = append(s.calls, write)
	return
}

type SpySleeper struct {
	calls int
}

func (s *SpySleeper) Sleep() {
	s.calls++
}

func TestWithSpies(t *testing.T) {

}

func TestCountdown(t *testing.T) {

	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpySleeper{}

		Countdown(buffer, spySleeper)

		got := buffer.String()
		want :=
			`3
2
1
Go!`

		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}

		if spySleeper.calls != 3 {
			t.Errorf("not enough calls to sleeper, want 3 got %d",
				spySleeper.calls)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spyCountDownSleepWriter := SpyCountdownOperations{}
		Countdown(&spyCountDownSleepWriter, &spyCountDownSleepWriter)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		got := spyCountDownSleepWriter.calls

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v but want %v", got, want)
		}
	})

}
