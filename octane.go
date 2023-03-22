package octane

import (
	"gitlab.com/gomidi/midi/v2"
	"gitlab.com/gomidi/midi/v2/drivers"

	"fmt"
	"os"
)

// TransposeKey applies a MIDI offset to a key.
func TransposeKey(key uint8, offset int) uint8 {
	return uint8((int(key) + offset) % 128)
}

// Stream begins copying data between MIDI IN devices,
// with an optional note transposition.
func Stream(midiIn drivers.In, midiOuts []drivers.Out, offset int) {
	var senders []func(msg midi.Message) error

	for _, midiOut := range midiOuts {
		sender, err := midi.SendTo(midiOut)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			break
		}

		senders = append(senders, sender)
	}

	var channel uint8
	var key uint8
	var velocity uint8

	react := func(msg midi.Message, timestampms int32) {
		switch {
		case msg.GetNoteStart(&channel, &key, &velocity):
			for _, sender := range senders {
				if err := sender(midi.NoteOn(channel, TransposeKey(key, offset), velocity)); err != nil {
					fmt.Fprintln(os.Stderr, err)
				}
			}
		case msg.GetNoteEnd(&channel, &key):
			for _, sender := range senders {
				if err := sender(midi.NoteOff(channel, TransposeKey(key, offset))); err != nil {
					fmt.Fprintln(os.Stderr, err)
				}
			}
		}
	}

	if _, err := midi.ListenTo(midiIn, react); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
