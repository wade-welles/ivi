// Copyright (c) 2017-2020 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

/*
Package key33220 implements the IVI Instrument driver for the Agilent 33220A
function generator.

State Caching: Not implemented
*/
package key33220

import (
	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/fgen"
)

// Ag33220 provides the IVI driver for an Agilent 33220A or 33210A
// function generator.
type Ag33220 struct {
	inst     ivi.Instrument
	Channels []Channel
	ivi.Inherent
}

// New creates a new Ag33220 IVI Instrument.
func New(inst ivi.Instrument, reset bool) (Ag33220, error) {
	channelNames := []string{
		"Output",
	}
	outputCount := len(channelNames)
	channels := make([]Channel, outputCount)
	for i, ch := range channelNames {
		baseChannel := fgen.NewChannel(i, ch, inst)
		channels[i] = Channel{baseChannel}
	}
	inherentBase := ivi.InherentBase{
		ClassSpecMajorVersion: 4,
		ClassSpecMinorVersion: 3,
		ClassSpecRevision:     "5.2",
		GroupCapabilities: []string{
			"IviFgenBase",
			"IviFgenStdfunc",
			"IviFgenTrigger",
			"IviFgenInternalTrigger",
			"IviFgenBurst",
		},
		SupportedInstrumentModels: []string{
			"33220A",
			"33210A",
		},
	}
	inherent := ivi.NewInherent(inst, inherentBase)
	driver := Ag33220{
		inst:     inst,
		Channels: channels,
		Inherent: inherent,
	}
	if reset {
		err := driver.Reset()
		return driver, err
	}
	return driver, nil
}

// Channel represents a repeated capability of an output channel for the
// function generator.
type Channel struct {
	fgen.Channel
}
