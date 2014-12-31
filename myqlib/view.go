package myqlib

import (
	"bytes"
	"strings"
)

// All Views must implement the following
type View interface {
	// outputs (write to the buffer)
	Help(b *bytes.Buffer)   // short help
	Header1(b *bytes.Buffer) // header to print above data
	Header2(b *bytes.Buffer) // header to print above data

	// A full line of output given the state
	Data(b *bytes.Buffer, state MyqState)

	// All the cols (including time col)
	Cols() []Col
}

// NormalView
type NormalView struct {
	cols []Col   // slice of columns in this view
  help string  // short description of the view
}

func (v NormalView) Help(b *bytes.Buffer) {
	b.WriteString(v.help)
}

func (v NormalView) Header1(b *bytes.Buffer) {
	// Print all col header1s for each in order
	var header1 bytes.Buffer
	for _, col := range v.Cols() {
		col.Header1(&header1)
		b.WriteString(" ")
	}
	// If the header1 buffer is all spaces, skip printing it
	hdr := strings.TrimSpace(header1.String())
	if hdr != "" {
		b.WriteString(header1.String())
		b.WriteString("\n")
	}
}

func (v NormalView) Header2(b *bytes.Buffer) {
	// Print all col header2s for each in order
	for _, col := range v.Cols() {
		col.Header2(b)
		b.WriteString(" ") // one space before each column
	}
	b.WriteString("\n")
}

func (v NormalView) Data(b *bytes.Buffer, state MyqState) {
	// Output all the col values in order based on their format
	for _, col := range v.Cols() {
		col.Data(b, state)
		b.WriteString(" ") // one space before each column
	}
	b.WriteString("\n")
}

// All columns preceeded by the time column
func (v NormalView) Cols() []Col {
	return v.cols
}