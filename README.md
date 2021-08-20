# octane: MIDI adder

Hello, Operator? [🎵 MP3](https://raw.githubusercontent.com/mcandre/octane/master/hello-operator.mp3)

# ABOUT

Octane ferries signals between MIDI devices, with optional transformations on the data. For example, the incoming pitch from one device can be shifted up or down before it arrives at another device. This is particularly useful when some devices have large gaps in note compatibility, or use extreme pitches to indicate CC-like signals.

By default, octane copies data from all available IN devices to all available OUT devices. Optional `-in`, `-out` flags can narrow the mapping.

Octane is free and open source: fork it to introduce your own creative MIDI tweaks!

# DOWNLOAD

https://github.com/mcandre/octane/releases

# DOCUMENTATION

https://godoc.org/github.com/mcandre/octane

# RUNTIME REQUIREMENTS

* macOS or Linux (no Windows or WSL support at this time)

# INSTALL FROM SOURCE

```console
$ go install github.com/mcandre/octane/cmd/octane@latest
```

# CONTRIBUTING

See [DEVELOPMENT.md](DEVELOPMENT.md).

# LICENSE

FreeBSD

# TIPS

* Polling may complete faster when MIDI software such as Arturia Analog Lab is running. Though be careful about such software interfering with your jam sessions.
* Polling may complete faster on Linux.
* Any USB MIDI adapter hubs may obfuscate or alter device names.
* MIDI device names may differ between operating systems.
* In a pinch, many MIDI devices can serve as adapters to reach further MIDI devices, using "thru" options.

# WE JAMMIN'

A quick hardware example triggers funky beats.

Equipment:

* Speaker (e.g., [Anker Sondcore Motion+](https://us.soundcore.com/products/a3116011))
* Sampler (e.g., [Bastl microGranny v2](https://bastl-instruments.com/instruments/microgranny) + [2GB microSD card](https://www.amazon.com/dp/B081NR485T/) + [microSD card reader](https://www.amazon.com/dp/B07H4VQ4BZ/))
* Sequencer (e.g., [KORG SQ-1](https://www.korg.com/us/products/dj/sq_1/) + [KORG MIDI TRS DIN adapter](https://www.amazon.com/dp/B0797SG8RS))
* MIDI controller (e.g., [Arturia KeyStep](https://www.arturia.com/keystep/overview))
* USB to MIDI DIN adapter (e.g., [iConnectivity mio 1 in 1 out](https://www.iconnectivity.com/products/midi/mio))
* PC (e.g., [Apple MacBook Pro](https://www.apple.com/macbook-pro/))
* audio cable (e.g. [3.5mm TRS male to male](https://www.amazon.com/dp/B00NO73Q84/))
* DIN cable (e.g., [5-pin male to male](https://www.amazon.com/dp/B093SW8ZNX/))
* assorted USB A/B/C/micro adapter cables
* assorted power supplies, batteries

Hardware Configuration:

1. Ensure speaker is powered on.
2. Connect sampler to speaker with an audio cable.
3. Ensure sampler is powered on.
4. Connect MIDI controller to sampler with a DIN cable.
5. Connect MIDI controller to PC with a USB cable.
6. Ensure MIDI controller is powered on.
7. Set the MIDI controller clock to internal.
8. Enable Hold function on MIDI controller, if you have one.
9. Connect sequencer MIDI _OUT_ port to DIN adapter.
10. Connect sequencer DIN adapter to MIDI _IN_ port of USB adapter.
11. Connect MIDI USB adapter to PC.
12. Ensure sequencer is powered on.
13. Set sequencer pattern to linear, left to right through 8 + 8 = 16 rows.
14. Randomize sequencer pitch knobs.
15. Set sequencer mode to Step Jump if you have one.
16. Ensure PC is powered on.

Software configuration:

1. Enumerate PC MIDI devices:

```
$ octane -list
Polling for MIDI devices...
MIDI IN devices:

* SQ-1 SEQ IN
* mio
* Arturia KeyStep 32

MIDI OUT devices:

* SQ-1 MIDI OUT
* SQ-1 CTRL
* mio
* Arturia KeyStep 32
```

Normally, the KORG SQ-1 sequencer emits notes too high for Bastl microGranny to understand. We will transpose the notes down four octaves, so that the sequencer can correctly trigger instructions to select different sample sounds.

2. Set `-transposeNote` to subtract 48 from each sequencer note. Test octane with different inputs and outputs for your particular setup.

```
$ octane \
    -transposeNote -48 \
    -in "SQ-1 SEQ IN" \
    -out "Arturia KeyStep 32"

Polling for MIDI devices...
Connected to MIDI IN device: SQ-1 SEQ IN
Connected to MIDI OUT device: Arturia KeyStep 32

(Play a sequence)

#0 [4 d:4] channel.NoteOn channel 0 key 48 velocity 64
#0 [1187 d:1187] channel.NoteOff channel 0 key 48
...
```

3. Set a sample going with the MIDI controller piano keys.
4. Start the sequencer playing.
5. Jam.

# CREDITS

* [gomidi](https://gitlab.com/gomidi)
