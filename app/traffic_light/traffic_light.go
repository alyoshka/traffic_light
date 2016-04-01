package trafficlight

import (
	"fmt"
	"io/ioutil"
)

const (
	// RED - command for turning red light on
	RED = "R"
	// GREEN - command for turning green light on
	GREEN = "G"
	// YELLOW - command for turning yellow light on
	YELLOW = "Y"
	// BLINK - command for blinking yellow light
	BLINK = "y"
)

// TrafficLight - interface for communicating with arduino
type TrafficLight interface {
	Red()
	Yellow()
	Green()
	Blink()
}

// TrafficLightTTY - implementation of traffic light interface
type TrafficLightTTY struct {
	tty string
}

// NewTrafficLightImpl - constructor
func NewTrafficLightTTY(tty string) TrafficLight {
	return &TrafficLightTTY{
		tty: tty,
	}
}

func (t *TrafficLightTTY) send(command []byte) {
	err := ioutil.WriteFile(t.tty, command, 0644)
	if err != nil {
		fmt.Println("Failed to write: ", err)
		panic(err)
	}
}

// Red - set the red light on
func (t *TrafficLightTTY) Red() {
	t.send([]byte(RED))
}

// Yellow - set yellow light
func (t *TrafficLightTTY) Yellow() {
	t.send([]byte(YELLOW))
}

// Green - set green light on
func (t *TrafficLightTTY) Green() {
	t.send([]byte(GREEN))
}

// Blink - blink yellow light
func (t *TrafficLightTTY) Blink() {
	t.send([]byte(BLINK))
}
