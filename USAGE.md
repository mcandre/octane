# USAGE GUIDE

We provide a rich set of features.

# LIST AVAILABLE MIDI DEVICES

`-list`

Example:

```sh
octane -list
```

# SELECT MIDI DEVICE INPUTS

`-in <devices>`

Comma separated.

Device names vary by platform.

Example:

```sh
octane -in "mio:mio MIDI 1 24:0"
```

# SELECT MIDI DEVICE OUTPUTS

`-out <devices>`

Comma separated.

Device names vary by platform.

Example:

```sh
octane -out "mio:mio MIDI 1 24:0"
```

# TRANSPOTES NOTES

`-transposeNote <offset>`

Sums incoming pitches with the given offset.

Example:

```sh
octane \
    -in "mio:mio MIDI 1 24:0" \
    -out "mio:mio MIDI 1 24:0" \
    -transposeNote -48
```
