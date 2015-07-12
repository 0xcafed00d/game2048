package main

const (
	MaxAxisCount   = 6
	MaxButtonCount = 32
)

type JoystickInfo struct {
	AxisData [MaxAxisCount]int
	Buttons  uint32
}

type Joystick interface {
	AxisCount() int
	ButtonCount() int
	Name() string
	Read() JoystickInfo
	Close()
}
