package main

import (
	"bytes"
	"encoding/binary"
	"os"
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
