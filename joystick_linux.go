// +build linux

package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"strconv"
	"unsafe"
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

var (
	JSIOCGAXES    = _IOR('j', 0x11, 1)  /* get number of axes */
	JSIOCGBUTTONS = _IOR('j', 0x12, 1)  /* get number of buttons */
	JSIOCGNAME    = func(len int) int { /* get identifier string */
		return _IOR('j', 0x13, len)
	}
)

type Joystick struct {
	file        *os.File
	AxisCount   int
	ButtonCount int
	Name        string
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

	var axisCount uint8 = 0
	var buttCount uint8 = 0
	var buffer [256]byte

	ioerr := Ioctl(f, JSIOCGAXES, unsafe.Pointer(&axisCount))
	if ioerr != 0 {
		panic(ioerr)
	}

	ioerr = Ioctl(f, JSIOCGBUTTONS, unsafe.Pointer(&buttCount))
	if ioerr != 0 {
		panic(ioerr)
	}

	ioerr = Ioctl(f, JSIOCGNAME(len(buffer)-1), unsafe.Pointer(&buffer))
	if ioerr != 0 {
		panic(ioerr)
	}

	return &Joystick{f, int(axisCount), int(buttCount), string(buffer[:])}, nil
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
	err = binary.Read(data, binary.LittleEndian, &event)
	if err != nil {
		return JSEvent{}, err
	}
	return event, nil
}
