# octane: MIDI adder

Hello, Operator? [🎵 MP3](https://raw.githubusercontent.com/mcandre/octane/master/hello-operator.mp3)

# EXAMPLE

```console
$ octane -list
Polling for MIDI devices...
MIDI IN devices:

* Arturia KeyStep 32
* SQ-1 SEQ IN

MIDI OUT devices:

* Arturia KeyStep 32
* SQ-1 MIDI OUT
* SQ-1 CTRL

$ octane -in 'SQ-1 SEQ IN' -out 'Arturia KeyStep 32' -transposeNote -48
Polling for MIDI devices...
Connected to MIDI IN device: SQ-1 SEQ IN
Connected to MIDI OUT device: Arturia KeyStep 32

(Play a sequence)

#0 [4 d:4] channel.NoteOn channel 0 key 48 velocity 64
#0 [1187 d:1187] channel.NoteOff channel 0 key 48
...
```

See `octane -help` for more options.

# TIPS

* Polling may complete faster when MIDI software such as Arturia Analog Lab is running. Though be careful about such software interfering with your jam sessions.
* In a pinch, many MIDI devices can serve as adapters to reach further MIDI devices, using "thru" options.

# DOWNLOAD

https://github.com/mcandre/octane/releases

# DOCUMENTATION

https://godoc.org/github.com/mcandre/octane

# RUNTIME REQUIREMENTS

(None)

# CONTRIBUTING

See [DEVELOPMENT.md](DEVELOPMENT.md).

# LICENSE

FreeBSD

# CREDITS

* [gomidi](https://gitlab.com/gomidi)
