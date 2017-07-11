// Copyright (c) 2017 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

/*
Package agilente36xx implements the IVI driver for the Agilent/Keysight E3600
series of power supplies.

IVI Instrument Class: IviDCPwr
IVI Class Specification: IVI-4.4
IVI Specification Revision: 3.0
IVI Specification Edition: September 24, 2015
Capability Groups Supported (specification section):
   4. IviDCPwrBase
	 5. IviDCPwrTrigger
	 6. IviDCPwrSoftwareTrigger
	 7. IviDCPwrMeasurement

Hardware Information:
  Instrument Manufacturer:          Keysight Technologies
	Supported Instrument Models:      E3631A

State Caching: Not implemented
*/
package agilente36xx

import "github.com/gotmc/ivi"

// AgilentE36xx provides the IVI driver for the Agilent/Keysight E3600 series
// of power supplies.
type AgilentE36xx struct {
	inst        ivi.Instrument
	outputCount int
	Channels    []Channel
}

// OutputCount returns the number of available output channels. OutputCount is
// the getter for the read-only IviDCPwrBase Attribute Output Channel Count
// described in Section 4.2.7 of IVI-4.4: IviDCPwr Class Specification.
func (dcpwr *AgilentE36xx) OutputCount() int {
	return dcpwr.outputCount
}

// New creates a new AgilentE36xx IVI Instrument.
func New(inst ivi.Instrument) (*AgilentE36xx, error) {
	outputCount := 3
	p6v := Channel{
		id:   0,
		name: "P6V",
		inst: inst,
	}
	p25v := Channel{
		id:   1,
		name: "P25V",
		inst: inst,
	}
	n25v := Channel{
		id:   2,
		name: "N25V",
		inst: inst,
	}
	channels := make([]Channel, outputCount)
	channels[0] = p6v
	channels[1] = p25v
	channels[2] = n25v
	dcpwr := AgilentE36xx{
		inst:        inst,
		outputCount: outputCount,
		Channels:    channels,
	}
	return &dcpwr, nil
}
