package drivers

import (
	"fmt"

	midilib "gitlab.com/gomidi/midi/v2/internal/utils"
)

type readerState int

const (
	readerStateClean                readerState = 0
	readerStateWithinChannelMessage readerState = 1
	readerStateWithinSysCommon      readerState = 2
	readerStateInSysEx              readerState = 3
	readerStateWithinUnknown        readerState = 4
)

const (
	byteMIDITimingCodeMessage  = byte(0xF1)
	byteSysSongPositionPointer = byte(0xF2)
	byteSysSongSelect          = byte(0xF3)
	byteSysTuneRequest         = byte(0xF6)
)

const (
	byteProgramChange         = 0xC
	byteChannelPressure       = 0xD
	byteNoteOff               = 0x8
	byteNoteOn                = 0x9
	bytePolyphonicKeyPressure = 0xA
	byteControlChange         = 0xB
	bytePitchWheel            = 0xE
)

type Reader struct {
	//	maxlenSysex     int
	sysexBf  []byte
	sysexlen int

	ts_ms      int32
	sysexTS    int32
	state      readerState
	statusByte uint8
	issetBf    bool
	bf         byte
	typ        uint8

	SysExBufferSize uint32
	OnMsg           func([]byte, int32)
	HandleSysex     bool
	OnErr           func(error)
}

func (me *Reader) withinChannelMessage(b byte) {
	//fmt.Println("withinChannelMessage")
	switch me.typ {
	case byteChannelPressure:
		me.issetBf = false
		me.state = readerStateClean
		//p.receiver.Receive(Channel(p.channel).Aftertouch(b), p.timestamp)
		me.OnMsg([]byte{me.statusByte, b}, me.ts_ms)
	case byteProgramChange:
		me.issetBf = false // first: is set, second: the byte
		me.state = readerStateClean
		//p.receiver.Receive(Channel(p.channel).ProgramChange(b), p.timestamp)
		me.OnMsg([]byte{me.statusByte, b}, me.ts_ms)
	case byteControlChange:
		if me.issetBf {
			me.issetBf = false // first: is set, second: the byte
			me.state = readerStateClean
			//p.receiver.Receive(Channel(p.channel).ControlChange(p.getBf(), b), p.timestamp)
			me.OnMsg([]byte{me.statusByte, me.bf, b}, me.ts_ms)
		} else {
			me.issetBf = true
			me.bf = b
		}
	case byteNoteOn:
		if me.issetBf {
			me.issetBf = false // first: is set, second: the byte
			me.state = readerStateClean
			//p.receiver.Receive(Channel(p.channel).NoteOn(p.getBf(), b), p.timestamp)
			me.OnMsg([]byte{me.statusByte, me.bf, b}, me.ts_ms)
		} else {
			me.issetBf = true
			me.bf = b
		}
	case byteNoteOff:
		if me.issetBf {
			me.issetBf = false // first: is set, second: the byte
			me.state = readerStateClean
			//p.receiver.Receive(Channel(p.channel).NoteOffVelocity(p.getBf(), b), p.timestamp)
			me.OnMsg([]byte{me.statusByte, me.bf, b}, me.ts_ms)
		} else {
			me.issetBf = true
			me.bf = b
		}
	case bytePolyphonicKeyPressure:
		if me.issetBf {
			me.issetBf = false // first: is set, second: the byte
			me.state = readerStateClean
			//p.receiver.Receive(Channel(p.channel).PolyAftertouch(p.getBf(), b), p.timestamp)
			me.OnMsg([]byte{me.statusByte, me.bf, b}, me.ts_ms)
		} else {
			me.issetBf = true
			me.bf = b
		}
	case bytePitchWheel:
		if me.issetBf {
			//rel, abs := midilib.ParsePitchWheelVals(bf, b)
			//_ = abs
			me.issetBf = false // first: is set, second: the byte
			me.state = readerStateClean
			//p.receiver.Receive(Channel(p.channel).Pitchbend(rel), p.timestamp)
			me.OnMsg([]byte{me.statusByte, me.bf, b}, me.ts_ms)
		} else {
			me.issetBf = true
			me.bf = b
		}
	default:
		panic("unknown typ")
	}
}

func (me *Reader) cleanState(b byte) {
	//fmt.Println("clean state")
	switch {

	/* start sysex */
	case b == 0xF0:
		me.statusByte = 0
		me.sysexBf = make([]byte, me.SysExBufferSize)
		//sysexBf.Reset()
		//sysexBf.WriteByte(b)
		me.sysexBf[0] = b
		me.sysexlen = 1
		me.sysexTS = me.ts_ms
		me.state = readerStateInSysEx
	// end sysex
	// [MIDI] permits 0xF7 octets that are not part of a (0xF0, 0xF7) pair
	// to appear on a MIDI 1.0 DIN cable.  Unpaired 0xF7 octets have no
	// semantic meaning in MIDI apart from cancelling running status.
	case b == 0xF7:
		me.sysexBf = nil
		me.sysexlen = 0
		me.statusByte = 0
		me.OnMsg([]byte{b}, me.ts_ms)

	// here we clear for System Common Category messages
	case b > 0xF0 && b < 0xF7:
		me.statusByte = 0
		me.issetBf = false // reset buffer
		//fmt.Printf("sys common msg started\n")
		switch b {
		case byteMIDITimingCodeMessage, byteSysSongPositionPointer, byteSysSongSelect:
			me.state = readerStateWithinSysCommon
			me.typ = b
		case byteSysTuneRequest:
			me.OnMsg([]byte{b}, me.ts_ms)
			/*
				if p.syscommonHander != nil {
					p.syscommonHander(Tune(), p.timestamp)
				}
			*/
			return
		default:
			// 0xF4, 0xF5, or 0xFD
			me.state = readerStateWithinUnknown
			return
		}

	// channel message with status byte
	case b >= 0x80 && b <= 0xEF:
		//fmt.Println("channel message")
		me.statusByte = b
		me.issetBf = false // reset buffer
		//typ, channel = midilib.ParseStatus(statusByte)
		me.typ, _ = midilib.ParseStatus(me.statusByte)
		me.state = readerStateWithinChannelMessage
	default:
		if me.statusByte != 0 {
			me.state = readerStateWithinChannelMessage
			me.withinChannelMessage(b)
		}
	}
}

func (me *Reader) eachByte(b byte) {
	if b >= 0xF8 {
		//r.OnMsg([]byte{b, 0, 0}, r.ts_ms)
		me.OnMsg([]byte{b}, me.ts_ms)
		return
	}

	//fmt.Printf("state: %v\n", p.state)

	switch me.state {
	case readerStateInSysEx:
		//fmt.Println("readerStateInSysEx")
		/* interrupted sysex, discard old data */
		if b == 0xF0 {
			me.statusByte = 0
			me.sysexBf = make([]byte, me.SysExBufferSize)
			me.sysexTS = me.ts_ms
			//sysexBf.Reset()
			//sysexBf.WriteByte(b)
			me.sysexBf[0] = b
			me.sysexlen = 1
			me.state = readerStateInSysEx
			return
		}

		if b == 0xF7 {
			/*
				if p.sysexHandler != nil {
					p.sysexBf.WriteByte(b)
					bt := p.sysexBf.Bytes()
					p.sysexBf.Reset()
					p.sysexHandler(bt, p.timestamp)
				}
			*/
			me.state = readerStateClean
			if me.HandleSysex {
				me.sysexBf[me.sysexlen] = b
				me.sysexlen++
				//go
				func(bb []byte, l int) {
					var _bt = make([]byte, l)

					for i := 0; i < l; i++ {
						_bt[i] = bb[i]
					}
					me.OnMsg(_bt, me.sysexTS)
				}(me.sysexBf, me.sysexlen)
			}
			me.sysexBf = nil
			me.sysexlen = 0
			return
		}
		if midilib.IsStatusByte(b) {
			//p.sysexBf.Reset()
			me.sysexBf = nil
			me.sysexlen = 0
			me.state = readerStateClean
			me.cleanState(b)
			return
		}

		if me.HandleSysex {
			me.sysexBf[me.sysexlen] = b
			me.sysexlen++
		}

		/*
			if p.sysexHandler != nil {
				p.sysexBf.WriteByte(b)
			}
		*/
	case readerStateClean:
		//fmt.Println("readerStateClean")
		me.cleanState(b)
	case readerStateWithinUnknown:
		//fmt.Println("readerStateWithinUnknown")
		//p.withinUnknown(b)
		if midilib.IsStatusByte(b) {
			me.state = readerStateClean
			me.cleanState(b)
		}
	case readerStateWithinSysCommon:
		//fmt.Println("readerStateWithinSysCommon")
		switch me.typ {
		case byteMIDITimingCodeMessage:
			/*
				if p.syscommonHander != nil {
					p.syscommonHander(MTC(b), p.timestamp)
				}
			*/
			me.issetBf = false
			me.state = readerStateClean
			me.OnMsg([]byte{me.typ, b}, me.ts_ms)
		case byteSysSongPositionPointer:
			if me.issetBf {
				/*
					if p.syscommonHander != nil {
						_, abs := midilib.ParsePitchWheelVals(p.getBf(), b)
						p.syscommonHander(SPP(abs), p.timestamp)
					}
				*/
				me.issetBf = false
				me.state = readerStateClean
				me.OnMsg([]byte{me.typ, me.bf, b}, me.ts_ms)
			} else {
				me.issetBf = true
				me.bf = b
			}
		case byteSysSongSelect:
			/*
				if p.syscommonHander != nil {
					p.syscommonHander(SongSelect(b), p.timestamp)
				}
			*/
			me.issetBf = false
			me.state = readerStateClean
			me.OnMsg([]byte{me.typ, b}, me.ts_ms)
		case byteSysTuneRequest:
			//panic("must not be handled here, but within clean state")
		default:
			if me.OnErr != nil {
				me.OnErr(fmt.Errorf("unknown syscommon message: % X", b))
			}
			//panic("unknown syscommon")
		}
	case readerStateWithinChannelMessage:
		//fmt.Println("readerStateWithinChannelMessage")
		me.withinChannelMessage(b)
	default:
		panic(fmt.Sprintf("unknown state %v, must not happen", me.state))
	}
}

func (me *Reader) Reset() {
	//fmt.Println("reset")

	if me.SysExBufferSize == 0 {
		me.SysExBufferSize = 1024
	}

	me.sysexBf = make([]byte, me.SysExBufferSize)
	me.sysexlen = 0
	me.ts_ms = 0
	me.statusByte = 0
	me.issetBf = false
	me.state = readerStateClean
}

func NewReader(config ListenConfig, onMsg func([]byte, int32)) *Reader {
	var r Reader
	r.OnMsg = onMsg
	r.OnErr = config.OnErr
	//r.OnSysEx = config.OnSysEx
	r.HandleSysex = config.SysEx
	r.SysExBufferSize = config.SysExBufferSize
	r.Reset()
	return &r
}

func (me *Reader) setDelta(deltaMilliSeconds int32) {
	me.ts_ms += deltaMilliSeconds
}

func (me *Reader) resetStatus() {
	me.statusByte = 0
	me.issetBf = false // first: is set, second: the byte
}

// func (r *Reader) EachMessage(bt []byte, deltaSeconds float64) {
func (me *Reader) EachMessage(bt []byte, deltaMilliSeconds int32) {

	// TODO: verify
	// assume that each call is without running state
	//r.ResetStatus()

	me.setDelta(deltaMilliSeconds) // int32(math.Round(deltaSeconds * 1000))

	//fmt.Printf("got % X\n", bt)

	for _, b := range bt {
		// => realtime message
		me.eachByte(b)

	}

}
