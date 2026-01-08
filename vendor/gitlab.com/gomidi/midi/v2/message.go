package midi

import (
	"bytes"
	"fmt"

	"gitlab.com/gomidi/midi/v2/internal/utils"
)

type Messages []Message

// Bytes returns all bytes of all the messages
func (me Messages) Bytes() (all []byte) {
	for _, msg := range me {
		all = append(all, msg.Bytes()...)
	}
	return
}

// Message is a complete midi message (not including meta messages)
type Message []byte

// Bytes returns the underlying bytes of the message.
func (me Message) Bytes() []byte {
	return []byte(me)
}

// IsPlayable returns, if the message can be send to an instrument.
func (me Message) IsPlayable() bool {
	if me.Type() <= UnknownMsg {
		return false
	}

	return me.Type() < firstMetaMsg
}

// IsOneOf returns true, if the message has one of the given types.
func (me Message) IsOneOf(checkers ...Type) bool {
	for _, checker := range checkers {
		if me.Is(checker) {
			return true
		}
	}
	return false
}

// Type returns the type of the message.
func (me Message) Type() Type {
	return getType(me)
}

// Is returns true, if the message is of the given type.
func (me Message) Is(t Type) bool {
	return me.Type().Is(t)
}

// GetNoteOn returns true if (and only if) the message is a NoteOnMsg.
// Then it also extracts the data to the given arguments.
// Only arguments that are not nil are parsed and filled.
func (me Message) GetNoteOn(channel, key, velocity *uint8) (is bool) {
	if !me.Is(NoteOnMsg) {
		return false
	}

	if len(me) != 3 {
		return false
	}

	if channel != nil {
		_, *channel = utils.ParseStatus(me[0])
	}

	if key != nil || velocity != nil {
		_key, _velocity := utils.ParseTwoUint7(me[1], me[2])

		if key != nil {
			*key = _key
		}

		if velocity != nil {
			*velocity = _velocity
		}
	}

	return true
}

// GetNoteStart returns true if (and only if) the message is a NoteOnMsg with a velocity > 0.
// Then it also extracts the data to the given arguments.
// Only arguments that are not nil are parsed and filled.
func (me Message) GetNoteStart(channel, key, velocity *uint8) (is bool) {
	var vel uint8

	if !me.GetNoteOn(channel, key, &vel) || vel == 0 {
		return false
	}

	if velocity != nil {
		*velocity = vel
	}
	return true
}

// GetNoteOff returns true if (and only if) the message is a NoteOffMsg.
// Then it also extracts the data to the given arguments.
// Only arguments that are not nil are parsed and filled.
func (me Message) GetNoteOff(channel, key, velocity *uint8) (is bool) {
	if !me.Is(NoteOffMsg) {
		return false
	}

	if len(me) != 3 {
		return false
	}

	if channel != nil {
		_, *channel = utils.ParseStatus(me[0])
	}

	if key != nil || velocity != nil {
		_key, _velocity := utils.ParseTwoUint7(me[1], me[2])

		if key != nil {
			*key = _key
		}

		if velocity != nil {
			*velocity = _velocity
		}
	}

	return true
}

// GetChannel returns true if (and only if) the message is a ChannelMsg.
// Then it also extracts the channel to the given argument.
// Only arguments that are not nil are parsed and filled.
func (me Message) GetChannel(channel *uint8) (is bool) {
	if !me.Is(ChannelMsg) {
		return false
	}

	if len(me) < 1 {
		return false
	}

	if channel != nil {
		_, *channel = utils.ParseStatus(me[0])
	}
	return true
}

// GetNoteEnd returns true if (and only if) the message is a NoteOnMsg with a velocity == 0 or a NoteOffMsg.
// Then it also extracts the data to the given arguments.
// Only arguments that are not nil are parsed and filled.
func (me Message) GetNoteEnd(channel, key *uint8) (is bool) {
	if !me.Is(NoteOnMsg) && !me.Is(NoteOffMsg) {
		return false
	}

	var vel uint8
	var ch uint8
	var k uint8

	is = false

	switch {
	case me.GetNoteOn(&ch, &k, &vel):
		is = vel == 0
	case me.GetNoteOff(&ch, &k, &vel):
		is = true
	}

	if !is {
		return false
	}

	if channel != nil {
		*channel = ch
	}

	if key != nil {
		*key = k
	}

	return true
}

// GetPolyAfterTouch returns true if (and only if) the message is a PolyAfterTouchMsg.
// Then it also extracts the data to the given arguments.
// Only arguments that are not nil are parsed and filled.
func (me Message) GetPolyAfterTouch(channel, key, pressure *uint8) (is bool) {
	if !me.Is(PolyAfterTouchMsg) {
		return false
	}

	if len(me) != 3 {
		return false
	}

	if channel != nil {
		_, *channel = utils.ParseStatus(me[0])
	}

	if key != nil || pressure != nil {
		var _key, _pressure = utils.ParseTwoUint7(me[1], me[2])

		if key != nil {
			*key = _key
		}

		if pressure != nil {
			*pressure = _pressure
		}
	}
	return true
}

// GetAfterTouch returns true if (and only if) the message is a AfterTouchMsg.
// Then it also extracts the data to the given arguments.
// Only arguments that are not nil are parsed and filled.
func (me Message) GetAfterTouch(channel, pressure *uint8) (is bool) {
	if !me.Is(AfterTouchMsg) {
		return false
	}

	if len(me) != 2 {
		return false
	}

	if channel != nil {
		_, *channel = utils.ParseStatus(me[0])
	}

	if pressure != nil {
		*pressure = utils.ParseUint7(me[1])
	}
	return true
}

// GetProgramChange returns true if (and only if) the message is a ProgramChangeMsg.
// Then it also extracts the data to the given arguments.
// Only arguments that are not nil are parsed and filled.
func (me Message) GetProgramChange(channel, program *uint8) (is bool) {
	if !me.Is(ProgramChangeMsg) {
		return false
	}

	if len(me) != 2 {
		return false
	}

	if channel != nil {
		_, *channel = utils.ParseStatus(me[0])
	}

	if program != nil {
		*program = utils.ParseUint7(me[1])
	}
	return true
}

// GetPitchBend returns true if (and only if) the message is a PitchBendMsg.
// Then it also extracts the data to the given arguments.
// Either relative or absolute may be nil, if not needed.
// Only arguments that are not nil are parsed and filled.
func (me Message) GetPitchBend(channel *uint8, relative *int16, absolute *uint16) (is bool) {
	if !me.Is(PitchBendMsg) {
		return false
	}

	if len(me) != 3 {
		return false
	}

	if channel != nil {
		_, *channel = utils.ParseStatus(me[0])
	}

	rel, abs := utils.ParsePitchWheelVals(me[1], me[2])
	if relative != nil {
		*relative = rel
	}
	if absolute != nil {
		*absolute = abs
	}
	return true
}

// GetControlChange returns true if (and only if) the message is a ControlChangeMsg.
// Then it also extracts the data to the given arguments.
// Only arguments that are not nil are parsed and filled.
func (me Message) GetControlChange(channel, controller, value *uint8) (is bool) {
	if !me.Is(ControlChangeMsg) {
		return false
	}

	if len(me) != 3 {
		return false
	}

	if channel != nil {
		_, *channel = utils.ParseStatus(me[0])
	}

	if controller != nil || value != nil {
		var _controller, _value uint8

		_controller, _value = utils.ParseTwoUint7(me[1], me[2])

		if controller != nil {
			*controller = _controller
		}

		if value != nil {
			*value = _value
		}
	}

	return true
}

/*
MTC Quarter Frame

These are the MTC (i.e. SMPTE based) equivalent of the F8 Timing Clock messages,
though offer much higher resolution, as they are sent at a rate of 96 to 120 times
a second (depending on the SMPTE frame rate). Each Quarter Frame message provides
partial timecode information, 8 sequential messages being required to fully
describe a timecode instant. The reconstituted timecode refers to when the first
partial was received. The most significant nibble of the data byte indicates the
partial (aka Message Type).

Partial	Data byte	Usage
1	0000 bcde	Frame number LSBs 	abcde = Frame number (0 to frameRate-1)
2	0001 000a	Frame number MSB
3	0010 cdef	Seconds LSBs 	abcdef = Seconds (0-59)
4	0011 00ab	Seconds MSBs
5	0100 cdef	Minutes LSBs 	abcdef = Minutes (0-59)
6	0101 00ab	Minutes MSBs
7	0110 defg	Hours LSBs 	ab = Frame Rate (00 = 24, 01 = 25, 10 = 30drop, 11 = 30nondrop)
cdefg = Hours (0-23)
8	0111 0abc	Frame Rate, and Hours MSB
*/

// GetMTC returns true if (and only if) the message is a MTCMsg.
// Then it also extracts the data to the given arguments.
// Only arguments that are not nil are parsed and filled.
func (me Message) GetMTC(quarterframe *uint8) (is bool) {
	if !me.Is(MTCMsg) {
		return false
	}

	if len(me) != 2 {
		return false
	}

	if quarterframe != nil {
		*quarterframe = utils.ParseUint7(me[1])
	}

	return true
}

// GetSongSelect returns true if (and only if) the message is a SongSelectMsg.
// Then it also extracts the song number to the given argument.
// Only arguments that are not nil are parsed and filled.
func (me Message) GetSongSelect(song *uint8) (is bool) {
	if !me.Is(SongSelectMsg) {
		return false
	}

	if len(me) != 2 {
		return false
	}

	if song != nil {
		*song = utils.ParseUint7(me[1])
	}

	return true
}

// GetSPP returns true if (and only if) the message is a SPPMsg.
// Then it also extracts the spp to the given argument.
// Only arguments that are not nil are parsed and filled.
func (me Message) GetSPP(spp *uint16) (is bool) {
	if !me.Is(SPPMsg) {
		return false
	}

	if len(me) != 3 {
		return false
	}

	if spp != nil {
		_, *spp = utils.ParsePitchWheelVals(me[2], me[1])
	}

	return true
}

// String represents the Message as a string that contains the Type and its properties.
func (me Message) String() string {
	var bf bytes.Buffer
	fmt.Fprint(&bf, me.Type().String())

	var channel, val1, val2 uint8
	var pitchabs uint16
	var pitchrel int16
	var spp uint16
	var sysex []byte

	switch {
	case me.GetNoteOn(&channel, &val1, &val2):
		fmt.Fprintf(&bf, " channel: %v key: %v velocity: %v", channel, val1, val2)
	case me.GetNoteOff(&channel, &val1, &val2):
		if val2 > 0 {
			fmt.Fprintf(&bf, " channel: %v key: %v velocity: %v", channel, val1, val2)
		} else {
			fmt.Fprintf(&bf, " channel: %v key: %v", channel, val1)
		}
	case me.GetPolyAfterTouch(&channel, &val1, &val2):
		fmt.Fprintf(&bf, " channel: %v key: %v pressure: %v", channel, val1, val2)
	case me.GetAfterTouch(&channel, &val1):
		fmt.Fprintf(&bf, " channel: %v pressure: %v", channel, val1)
	case me.GetControlChange(&channel, &val1, &val2):
		fmt.Fprintf(&bf, " channel: %v controller: %v value: %v", channel, val1, val2)
	case me.GetProgramChange(&channel, &val1):
		fmt.Fprintf(&bf, " channel: %v program: %v", channel, val1)
	case me.GetPitchBend(&channel, &pitchrel, &pitchabs):
		fmt.Fprintf(&bf, " channel: %v pitch: %v (%v)", channel, pitchrel, pitchabs)
	case me.GetMTC(&val1):
		fmt.Fprintf(&bf, " mtc: %v", val1)
	case me.GetSPP(&spp):
		fmt.Fprintf(&bf, " spp: %v", spp)
	case me.GetSongSelect(&val1):
		fmt.Fprintf(&bf, " song: %v", val1)
	case me.GetSysEx(&sysex):
		fmt.Fprintf(&bf, " data: % X", sysex)
	default:
	}

	return bf.String()
}

// GetSysEx returns true, if the message is a sysex message.
// Then it extracts the inner bytes to the given slice.
func (me Message) GetSysEx(bt *[]byte) bool {
	if len(me) < 3 {
		return false
	}

	if !me.Is(SysExMsg) {
		return false
	}

	if me[0] == 0xF0 && me[len(me)-1] == 0xF7 {
		*bt = me[1 : len(me)-1]
		return true
	}

	return false
}
