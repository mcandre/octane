package octane

import (
	"gitlab.com/gomidi/midi/mid"
	"gitlab.com/gomidi/rtmididrv/imported/rtmidi"
)

// Version is semver.
const Version = "0.0.1"

// NOPCallback does nothing.
func NOPCallback(_ rtmidi.MIDIIn, _ []byte, _ float64) {}

// RegisterTranspose configures MIDI hooks for streaming,
// with optional note transposition.
func RegisterTranspose(reader *mid.Reader, midiOut mid.Out, offset int) {
	writer := mid.ConnectOut(midiOut)

	reader.Msg.Channel.NoteOn = func(_ *mid.Position, channel uint8, key uint8, velocity uint8) {
		keyTransposed := uint8(int(key) + offset)
		writer.SetChannel(channel)
		writer.NoteOn(keyTransposed, velocity)
	}

	reader.Msg.Channel.NoteOff = func(_ *mid.Position, channel uint8, key uint8, velocity uint8) {
		keyTransposed := uint8(int(key) + offset)
		writer.SetChannel(channel)
		writer.NoteOffVelocity(keyTransposed, velocity)
	}
}

// Stream begins copying data between MIDI IN devices,
// with an optional note transposition.
func Stream(midiIn mid.In, midiOuts []mid.Out, offset int) {
	reader := mid.NewReader()

	for _, midiOut := range midiOuts {
		RegisterTranspose(reader, midiOut, offset)
	}

	go mid.ConnectIn(midiIn, reader)
}
