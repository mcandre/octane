package octane

import (
	"gitlab.com/gomidi/midi/mid"
	"gitlab.com/gomidi/rtmididrv/imported/rtmidi"
)

// Version is semver.
const Version = "0.0.1"

// NOPCallback does nothing.
func NOPCallback(_ rtmidi.MIDIIn, _ []byte, _ float64) {}

// RegisterHooks connects callbacks to MIDI IN devices.
func RegisterHooks(midiIn mid.In) {
	reader := mid.NewReader()
	go mid.ConnectIn(midiIn, reader)
}
