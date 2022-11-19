package animation

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"time"
)

type Animation struct {
	Frames   []string
	IsMoving bool
	Done     chan bool

	startPadding int
	time         time.Duration
	tick         int
}

var (
	Animations = []*Animation{
		&Animation{
			Frames:   []string{"🌕", "🌖", "🌗", "🌘", "🌑", "🌒", "🌓", "🌔"},
			IsMoving: false,
			time:     1 * time.Second,
			tick:     0,
			Done:     make(chan bool),
		},
		&Animation{
			Frames:   []string{"🌎", "🌍", "🌏"},
			IsMoving: false,
			time:     1 * time.Second,
			tick:     0,
			Done:     make(chan bool),
		},
		&Animation{
			Frames:   []string{"🏎️"},
			IsMoving: true,
			time:     3 * time.Second,
			tick:     0,
			Done:     make(chan bool),
		},
	}
)

func NewAnimation(frames []string, isMoving bool, time time.Duration) (*Animation, error) {
	if len(frames) == 0 {
		return nil, errors.New("no frames provided")
	}

	a := &Animation{
		Frames:   frames,
		IsMoving: isMoving,
		time:     time,

		tick: 0,
		Done: make(chan bool),
	}

	return a, nil
}

func (a *Animation) Run() {
	a.startPadding = int(a.time / (100 * time.Millisecond))

	ticker := time.NewTicker(100 * time.Millisecond)
	timer := time.After(a.time)

	tput("civis") // hide

	for {
		select {
		case <-timer:
			ticker.Stop()
			tput("cvvis")
			a.Done <- true
			return
		case <-ticker.C:
			a.nextFrame()
		}
	}
}

func (a *Animation) nextFrame() {
	a.tick++

	var (
		frame   = a.Frames[a.tick%len(a.Frames)]
		padding = ""
	)

	ClearLine()
	if a.IsMoving {
		for i := 0; i < a.startPadding-a.tick; i++ {
			padding += " "
		}
	}
	fmt.Print(padding + frame)
}

func ClearLine() {
	fmt.Printf("\033[2K")
	fmt.Printf("\r")
}

func tput(arg string) error {
	cmd := exec.Command("tput", arg)
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
