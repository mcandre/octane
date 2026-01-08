//go:build !js
// +build !js

package rtmididrv

import (
	"fmt"

	"gitlab.com/gomidi/midi/v2/drivers"
	"gitlab.com/gomidi/midi/v2/drivers/rtmididrv/imported/rtmidi"
)

func newOut(driver *Driver, number int, name string) drivers.Out {
	o := &out{driver: driver, number: number, name: name}
	return o
}

type out struct {
	number int
	//sync.RWMutex
	driver  *Driver
	name    string
	midiOut rtmidi.MIDIOut
}

// IsOpen returns wether the port is open
func (me *out) IsOpen() (open bool) {
	//	o.RLock()
	open = me.midiOut != nil
	//	o.RUnlock()
	return
}

// Send writes a MIDI sysex message to the outut port
func (me *out) SendSysEx(data []byte) error {
	//fmt.Printf("try to send sysex\n")

	if me.midiOut == nil {
		//o.RUnlock()
		return drivers.ErrPortClosed
	}
	//o.mx.RUnlock()

	// since we always open the outputstream with a latency of 0
	// the timestamp is ignored
	//var ts portmidi.Timestamp // or portmidi.Time()

	//o.mx.Lock()
	//	defer o.mx.Unlock()
	//fmt.Printf("sending sysex % X\n", data)
	//err := o.stream.WriteSysExBytes(ts, data)
	err := me.midiOut.SendMessage(data)
	if err != nil {
		return fmt.Errorf("could not send sysex message to MIDI out %v (%s): %v", me.Number(), me, err)
	}
	return nil
}

func (me *out) Send(b []byte) error {
	if me.midiOut == nil {
		//o.RUnlock()
		return drivers.ErrPortClosed
	}
	//	o.RUnlock()

	//fmt.Printf("send % X\n", m.Data)
	/*
		var bt []byte

		switch {
		case b[2] == 0 && b[1] == 0:
			bt = []byte{b[0]}
			//	case b[2] == 0:
		//	bt = []byte{b[0], b[1]}
		default:
			bt = []byte{b[0], b[1], b[2]}
		}

		//bt := []byte{b[0], b[1], b[2]}
		err := o.midiOut.SendMessage(bt)
	*/
	err := me.midiOut.SendMessage(b)
	if err != nil {
		return fmt.Errorf("could not send message to MIDI out %v (%s): %v", me.number, me, err)
	}
	return nil
}

/*
// Send writes a MIDI message to the MIDI output port
// If the output port is closed, it returns midi.ErrClosed
func (o *out) send(bt []byte) error {
	//o.RLock()
	o.Lock()
	defer o.Unlock()
	if o.midiOut == nil {
		//o.RUnlock()
		return drivers.ErrPortClosed
	}
	//	o.RUnlock()

	//fmt.Printf("send % X\n", m.Data)
	err := o.midiOut.SendMessage(bt)
	if err != nil {
		return fmt.Errorf("could not send message to MIDI out %v (%s): %v", o.number, o, err)
	}
	return nil
}
*/

// Underlying returns the underlying rtmidi.MIDIOut. Use it with type casting:
//
//	rtOut := o.Underlying().(rtmidi.MIDIOut)
func (me *out) Underlying() interface{} {
	return me.midiOut
}

// Number returns the number of the MIDI out port.
// Note that with rtmidi, out and in ports are counted separately.
// That means there might exists out ports and an in ports that share the same number
func (me *out) Number() int {
	return me.number
}

// String returns the name of the MIDI out port.
func (me *out) String() string {
	return me.name
}

// Close closes the MIDI out port
func (me *out) Close() (err error) {
	if !me.IsOpen() {
		return nil
	}
	//o.Lock()
	//defer o.Unlock()

	err = me.midiOut.Close()
	me.midiOut = nil

	if err != nil {
		err = fmt.Errorf("can't close MIDI out %v (%s): %v", me.number, me, err)
	}

	return
}

// Open opens the MIDI out port
func (me *out) Open() (err error) {
	if me.IsOpen() {
		return nil
	}
	//	o.Lock()
	//defer o.Unlock()
	me.midiOut, err = rtmidi.NewMIDIOutDefault()
	if err != nil {
		me.midiOut = nil
		return fmt.Errorf("can't open default MIDI out: %v", err)
	}

	err = me.midiOut.OpenPort(me.number, "")
	if err != nil {
		me.midiOut = nil
		return fmt.Errorf("can't open MIDI out port %v (%s): %v", me.number, me, err)
	}

	//	o.driver.Lock()
	me.driver.opened = append(me.driver.opened, me)
	//	o.driver.Unlock()

	return nil
}
