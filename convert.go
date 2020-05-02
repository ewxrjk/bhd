package main

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

// ErrUnknownFormatString is returned when an unrecognized format token is used.
var ErrUnknownFormatString = errors.New("unrecognized format string")

var DefaultForwardFormat = []string{"offset", "fhex", "ftext"}

var DefaultBackwardFormat = []string{"bhex", "offset", "btext"}

type conversionParameters struct {
	w      io.Writer
	offset uint64
	buf    []byte
	n      int
	width  int
	group  int
}

func convertOffset(cp *conversionParameters) {
	fmt.Fprintf(cp.w, "%8x", cp.offset)
	return
}

func convertHexBackward(cp *conversionParameters) {
	convertHex(cp, cp.width-1, -1, -1, 0)
}

func convertHexForward(cp *conversionParameters) {
	convertHex(cp, 0, cp.width, 1, cp.group-1)
}

func convertHex(cp *conversionParameters, ifirst, ilast, idelta, groupend int) {
	for i := ifirst; i != ilast; i += idelta {
		if i < cp.n {
			fmt.Fprintf(cp.w, "%02x", cp.buf[i])
		} else {
			fmt.Fprintf(cp.w, "  ")
		}
		if i%cp.group == groupend && i+idelta != ilast {
			fmt.Fprintf(cp.w, " ")
		}
	}
}

func convertTextBackward(cp *conversionParameters) {
	convertText(cp, cp.width-1, -1, -1, 0)
}

func convertTextForward(cp *conversionParameters) {
	convertText(cp, 0, cp.width, 1, cp.group-1)
}

func convertText(cp *conversionParameters, ifirst, ilast, idelta, groupend int) {
	for i := ifirst; i != ilast; i += idelta {
		var ch byte
		if i < cp.n {
			ch = cp.buf[i]
			if ch < 32 || ch >= 127 {
				ch = '.'
			}
		} else {
			ch = ' '
		}
		fmt.Fprintf(cp.w, "%c", ch)
	}
}

var conversions = map[string]func(*conversionParameters){
	"offset": convertOffset,
	"ftext":  convertTextForward,
	"btext":  convertTextBackward,
	"fhex":   convertHexForward,
	"bhex":   convertHexBackward,
}

// Convert turns a byte sequence into a hexdump, according to a format.
func Convert(offset uint64, buf []byte, n int, width int, group int, format []string) (s string, err error) {
	w := &strings.Builder{}
	cp := &conversionParameters{
		w:      w,
		offset: offset,
		buf:    buf,
		n:      n,
		width:  width,
		group:  group,
	}
	for j, f := range format {
		if j != 0 {
			fmt.Fprintf(w, " | ")
		}
		if fn, ok := conversions[f]; ok {
			fn(cp)
		} else {
			err = ErrUnknownFormatString
			return
		}
	}
	s = w.String()
	return
}
