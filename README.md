# octane: MIDI adder

# EXAMPLE

```console
$ octane -list
Polling for MIDI devices...
MIDI IN devices:

* SQ-1 SEQ IN

MIDI OUT devices:

* SQ-1 MIDI OUT
* SQ-1 CTRL

$ octane -in "SQ-1 SEQ IN"
Polling for MIDI devices...
Connected to MIDI IN device: SQ-1 SEQ IN

(Play a sequence)

#0 [4 d:4] channel.NoteOn channel 0 key 48 velocity 64
#0 [1187 d:1187] channel.NoteOff channel 0 key 48
...
```

See `octane -help` for more options.

# TIPS

* Polling may complete faster when MIDI software such as Arturia Analog Lab is running. Though be careful about such software interfering with your jam sessions.

# TODO

* Hook up MIDI OUT device (e.g. Bastl microGranny using a USB MIDI adapter)
* Implement -transposeOctave <signed integer>
* Check whether MIDI OUT devices note-hang when disconnected between note on/off events.
* Implement -transposeNote <signed integer>
