package gamepad

import "github.com/0xcafed00d/joystick"

type GamePad struct {
	js joystick.Joystick
}

func NewGamepad(id int) (*GamePad, error) {
	js, err := joystick.Open(id)
	if err != nil {
		return nil, err
	}
	return &GamePad{js: js}, nil
}

func (g *GamePad) State() (State, error) {
	state, err := g.js.Read()
	return State{state: state}, err
}
