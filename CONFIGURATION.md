# CONFIGURATION

octane uses CLI flags for configuration.

# -list

List available MIDI devices.

Example:

```sh
octane -list
```

# `-in <devices>`

Select MIDI device inputs.

Comma separated.

Device names vary by platform.

Example:

```sh
octane -in "mio:mio MIDI 1 24:0"
```

# `-out <devices>`

Select MIDI device outputs.

Comma separated.

Device names vary by platform.

Example:

```sh
octane -out "mio:mio MIDI 1 24:0"
```

# `-transposeNote <offset>`

Sums incoming pitches with the given offset.

Example:

```sh
octane \
    -in "mio:mio MIDI 1 24:0" \
    -out "mio:mio MIDI 1 24:0" \
    -transposeNote -48
```
