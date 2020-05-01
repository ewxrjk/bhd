package main

import (
	"errors"
	"fmt"
	"strings"
)

var ErrUnknownFormatString = errors.New("unrecognized format string")

// Convert turns a byte sequence into a (possibly backwards) hexdump.
func Convert(offset uint64, buf []byte, n int, width int, group int, format []string, reverse bool) (s string, err error) {
	var ifirst, idelta, ilast, groupend int
	if reverse {
		ifirst = width - 1
		ilast = -1
		idelta = -1
		groupend = 0
		if len(format) == 0 {
			format = []string{"hex", "offset", "text"}
		}
	} else {
		ifirst = 0
		ilast = width
		idelta = 1
		groupend = group - 1
		if len(format) == 0 {
			format = []string{"offset", "hex", "text"}
		}
	}
	w := &strings.Builder{}
	for j, f := range format {
		if j != 0 {
			fmt.Fprintf(w, " | ")
		}
		switch f {
		case "offset":
			fmt.Fprintf(w, "%8x", offset)
		case "hex":
			for i := ifirst; i != ilast; i += idelta {
				if i < n {
					fmt.Fprintf(w, "%02x", buf[i])
				} else {
					fmt.Fprintf(w, "  ")
				}
				if i%group == groupend && i+idelta != ilast {
					fmt.Fprintf(w, " ")
				}
			}
		case "text":
			for i := ifirst; i != ilast; i += idelta {
				var ch byte
				if i < n {
					ch = buf[i]
					if ch < 32 || ch >= 127 {
						ch = '.'
					}
				} else {
					ch = ' '
				}
				fmt.Fprintf(w, "%c", ch)
			}
		default:
			err = ErrUnknownFormatString
			return
		}
	}
	s = w.String()
	return
}
