package octane

import (
	"gitlab.com/gomidi/midi"
	"gitlab.com/gomidi/midi/reader"
	"gitlab.com/gomidi/midi/writer"

	"fmt"
	"os"
)

// Version is semver.
const Version = "0.0.4"

// TransposeKey applies a MIDI offset to a key.
func TransposeKey(key uint8, offset int) uint8 {
	return uint8((int(key) + offset) % 128)
}

// Transpose configures MIDI hooks for streaming,
// with optional note transposition.
func Transpose(midiIn midi.In, _ midi.Out, offset int) {

}

// Stream begins copying data between MIDI IN devices,
// with an optional note transposition.
func Stream(midiIn midi.In, midiOuts []midi.Out, offset int) {
	var writers []*writer.Writer

	for _, midiOut := range midiOuts {
		wr := writer.New(midiOut)
		writers = append(writers, wr)
	}

	rd := reader.New(
		reader.NoteOn(func(_ *reader.Position, channel, key, velocity uint8) {
			for _, wr := range writers {
				wr.SetChannel(channel)

				if err := writer.NoteOn(wr, TransposeKey(key, offset), velocity); err != nil {
					fmt.Fprintln(os.Stderr, err)
				}
			}
		}),
		reader.NoteOff(func(_ *reader.Position, channel, key, velocity uint8) {
			for _, wr := range writers {
				wr.SetChannel(channel)

				if err := writer.NoteOffVelocity(wr, TransposeKey(key, offset), velocity); err != nil {
					fmt.Fprintln(os.Stderr, err)
				}
			}
		}),
	)

	if err := rd.ListenTo(midiIn); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
