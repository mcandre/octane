// Package main implements a CLI MIDI shifter application.
package main

import (
	"github.com/mcandre/octane"
	"gitlab.com/gomidi/midi/v2"
	"gitlab.com/gomidi/midi/v2/drivers"
	_ "gitlab.com/gomidi/midi/v2/drivers/rtmididrv"

	"flag"
	"fmt"
	"os"
	"strings"
)

var flagList = flag.Bool("list", false, "List MIDI devices")
var flagIn = flag.String("in", "", "Select comma-separated MIDI IN devices by name. Example: \"Arturia KeyStep 32,SQ-1 SEQ IN\"")
var flagOut = flag.String("out", "", "Select comma-separated MIDI OUT devices by name. Example: \"Arturia KeyStep 32,SQ-1 MIDI OUT\"")
var flagTransposeNote = flag.Int("transposeNote", 0, "Note offset. Example: -48")
var flagHelp = flag.Bool("help", false, "Show usage information")
var flagVersion = flag.Bool("version", false, "Show version information")

func main() {
	flag.Parse()

	switch {
	case *flagVersion:
		fmt.Printf("/VVV %v\n", octane.Version)
		os.Exit(0)
	case *flagHelp:
		flag.PrintDefaults()
		os.Exit(0)
	}

	defer midi.CloseDriver()

	fmt.Println("Polling for MIDI devices...")

	midiIns := midi.GetInPorts()
	midiOuts := midi.GetOutPorts()

	if *flagList {
		if len(midiIns) == 0 {
			fmt.Println("No MIDI IN devices.")
		} else {
			fmt.Printf("MIDI IN devices:\n\n")

			for _, midiIn := range midiIns {
				fmt.Printf("* %v\n", midiIn)
			}

			fmt.Println()
		}

		if len(midiOuts) == 0 {
			fmt.Println("No MIDI OUT devices.")
		} else {
			fmt.Printf("MIDI OUT devices:\n\n")

			for _, midiOut := range midiOuts {
				fmt.Printf("* %v\n", midiOut)
			}

			fmt.Println()
		}

		os.Exit(0)
	}

	var midiInWhitelist []string

	if *flagIn != "" {
		midiInWhitelist = strings.Split(*flagIn, ",")
	}

	var midiInsFiltered []drivers.In

	if len(midiInWhitelist) == 0 {
		midiInsFiltered = midiIns
	} else {
		for _, name := range midiInWhitelist {
			var foundMIDI bool

			for _, midiIn := range midiIns {
				if midiIn.String() == name {
					midiInsFiltered = append(midiInsFiltered, midiIn)
					foundMIDI = true
					break
				}
			}

			if !foundMIDI {
				fmt.Fprintf(os.Stderr, "Unable to find MIDI IN named: %v\n", name)
				os.Exit(1)
			}
		}
	}

	var midiOutWhitelist []string

	if *flagOut != "" {
		midiOutWhitelist = strings.Split(*flagOut, ",")
	}

	var midiOutsFiltered []drivers.Out

	if len(midiOutWhitelist) == 0 {
		midiOutsFiltered = midiOuts
	} else {
		for _, name := range midiOutWhitelist {
			var foundMIDI bool

			for _, midiOut := range midiOuts {
				if midiOut.String() == name {
					midiOutsFiltered = append(midiOutsFiltered, midiOut)
					foundMIDI = true
					break
				}
			}

			if !foundMIDI {
				fmt.Fprintf(os.Stderr, "Unable to find MIDI OUT named: %v\n", name)
				os.Exit(1)
			}
		}
	}

	for _, midiIn := range midiInsFiltered {
		if err2 := midiIn.Open(); err2 != nil {
			panic(err2)
		}

		defer midiIn.Close()

		fmt.Printf("Connected to MIDI IN device: %v\n", midiIn)
	}

	for _, midiOut := range midiOutsFiltered {
		if err2 := midiOut.Open(); err2 != nil {
			panic(err2)
		}

		defer midiOut.Close()

		fmt.Printf("Connected to MIDI OUT device: %v\n", midiOut)
	}

	for _, midiIn := range midiInsFiltered {
		octane.Stream(midiIn, midiOutsFiltered, *flagTransposeNote)
	}

	select {}
}
