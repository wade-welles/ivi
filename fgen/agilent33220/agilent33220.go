// Copyright (c) 2017 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package agilent33220

import (
	"fmt"
	"io"
	"strconv"

	"github.com/gotmc/ivi"
)

type StandardWaveform int

const (
	Sine StandardWaveform = iota
	Square
	Triangle
	RampUp
	RampDown
	DC
)

type StandardFunction struct {
	amplitude     float64
	dCOffset      float64
	dutyCycleHigh float64
	frequency     float64
	startPhase    float64
	waveform      StandardWaveform
}

func (stdFunc *StandardFunction) ConfigureWaveform(w io.Writer) {

}

type Agilent33220 struct {
	inst     ivi.Instrument
	Channels *[]Channel
}

type Channel struct {
}

func (fgen *Agilent33220) OutputCount() int {
	return 1
}

func (fgen *Agilent33220) GetAmplitude(ch int) (float64, error) {
	return getAmplitude(fgen.inst, ch)
}

func getAmplitude(inst ivi.Instrument, ch int) (float64, error) {
	if ch != 0 {
		return 0, fmt.Errorf("Channel doesn't exist: %d", ch)
	}
	value, err := inst.Query("VOLT?")
	if err != nil {
		return 0, err
	}
	return strconv.ParseFloat(value, 64)
}

func (fgen *Agilent33220) Close() error {
	return fgen.inst.Close()
}

func New(inst ivi.Instrument) (Agilent33220, error) {
	var fgen Agilent33220
	fgen.inst = inst
	return fgen, nil
}
