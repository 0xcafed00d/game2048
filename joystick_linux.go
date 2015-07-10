// +build linux

package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"strconv"
)

const (
	JS_EVENT_BUTTON uint8 = 0x01 /* button pressed/released */
	JS_EVENT_AXIS   uint8 = 0x02 /* joystick moved */
	JS_EVENT_INIT   uint8 = 0x80

	JS_AXIS_X0 uint8 = 0
	JS_AXIS_Y0 uint8 = 1
	JS_AXIS_X1 uint8 = 2
	JS_AXIS_Y1 uint8 = 3
)

type Joystick struct {
	file *os.File
}

type JSEvent struct {
	Time   uint32 /* event timestamp in milliseconds */
	Value  int16  /* value */
	Type   uint8  /* event type */
	Number uint8  /* axis/button number */
}

func (j *JSEvent) String() string {
	var Type, Number string

	if j.Type&JS_EVENT_INIT > 0 {
		Type = "Init "
	}
	if j.Type&JS_EVENT_BUTTON > 0 {
		Type += "Button"
		Number = strconv.FormatUint(uint64(j.Number), 10)
	}
	if j.Type&JS_EVENT_AXIS > 0 {
		Type = "Axis"
		switch j.Number {
		case JS_AXIS_X0:
			Number = "Axis X0"
		case JS_AXIS_Y0:
			Number = "Axis Y0"
		case JS_AXIS_X1:
			Number = "Axis X1"
		case JS_AXIS_Y1:
			Number = "Axis Y0"
		default:
			Number = "Axis " + strconv.FormatUint(uint64(j.Number), 10)
		}
	}

	return fmt.Sprintf("[Time: %v, Type: %v, Number: %v, Value: %v]", j.Time, Type, Number, j.Value)
}

func OpenJoystick(name string) (*Joystick, error) {
	f, err := os.OpenFile(name, os.O_RDONLY, 0666)

	if err != nil {
		return nil, err
	}

	return &Joystick{f}, nil
}

func (j *Joystick) Close() error {
	return j.file.Close()
}

func (j *Joystick) GetEvent() (JSEvent, error) {
	var event JSEvent

	if j.file == nil {
		panic("file is nil")
	}

	b := make([]byte, 8)
	_, err := j.file.Read(b)
	if err != nil {
		return JSEvent{}, err
	}

	data := bytes.NewReader(b)

	err = binary.Read(data, binary.LittleEndian, &event.Time)
	if err != nil {
		return JSEvent{}, err
	}
	err = binary.Read(data, binary.LittleEndian, &event.Value)
	if err != nil {
		return JSEvent{}, err
	}
	err = binary.Read(data, binary.LittleEndian, &event.Type)
	if err != nil {
		return JSEvent{}, err
	}
	err = binary.Read(data, binary.LittleEndian, &event.Number)
	if err != nil {
		return JSEvent{}, err
	}

	return event, nil
}
