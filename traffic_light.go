package tl

import (
	"fmt"
	"io/ioutil"
)

// RED - command for turning red light on
const RED = "R"

// GREEN - command for turning green light on
const GREEN = "G"

// YELLOW - command for turning yellow light on
const YELLOW = "Y"

// BLINK - command for blinking yellow light
const BLINK = "y"

// TrafficLight - interface for communicating with arduino
type TrafficLight interface {
	Red()
	Yellow()
	Green()
	Blink()
}

// TrafficLightImpl - implementation of traffic light interface
type TrafficLightImpl struct {
	tty string
}

// NewTrafficLightImpl - constructor
func NewTrafficLightImpl(tty string) TrafficLight {
	return &TrafficLightImpl{
		tty: tty,
	}
}

func (t *TrafficLightImpl) send(command []byte) {
	err := ioutil.WriteFile(t.tty, command, 0644)
	if err != nil {
		fmt.Println("Failed to write: ", err)
		panic(err)
	}
}

// Red - set the red light on
func (t *TrafficLightImpl) Red() {
	t.send([]byte(RED))
}

// Yellow - set yellow light
func (t *TrafficLightImpl) Yellow() {
	t.send([]byte(YELLOW))
}

// Green - set green light on
func (t *TrafficLightImpl) Green() {
	t.send([]byte(GREEN))
}

// Blink - blink yellow light
func (t *TrafficLightImpl) Blink() {
	t.send([]byte(BLINK))
}
