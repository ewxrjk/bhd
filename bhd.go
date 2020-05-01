package main

import (
	"fmt"
	"io"
	"os"

	"bufio"

	"github.com/spf13/cobra"
)

var width int
var group int
var forward bool
var format []string

func init() {
	bhdCmd.PersistentFlags().IntVarP(&width, "width", "w", 16, "width of each line in bytes")
	bhdCmd.PersistentFlags().IntVarP(&group, "group", "g", 1, "byte group size")
	bhdCmd.PersistentFlags().BoolVarP(&forward, "forward", "f", false, "forward hexdump")
	bhdCmd.PersistentFlags().StringSliceVarP(&format, "format", "F", []string{}, "line format")
}

func main() {
	if err := bhdCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		os.Exit(1)
	}
}

var bhdCmd = cobra.Command{
	Use:   "bhd [PATH...]",
	Short: "Backwards hexdump",
	Long: `Backwards hexdump

Displays a hexdump of the input files, with bytes ordered from right to left,
making little-endian fields easier to read by eye.
• Use --group to group bytes together to easily read multi-byte fields.
• Use --forward for a normal hexdump.

--format is a comma-separated list of tokens:
• hex is the hexdump
• text is the plain text
• offset is the offset within the file
`,
	SilenceUsage:  true,
	SilenceErrors: true,
	Version:       Version,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		w := bufio.NewWriter(os.Stdout)
		if len(args) == 0 {
			if err = HexdumpFile("-", w, format); err != nil {
				return
			}
		} else {
			for _, path := range args {
				if err = HexdumpFile(path, w, format); err != nil {
					return
				}
			}
		}
		if err = w.Flush(); err != nil {
			return
		}
		return
	},
}

// HexdumpFile dumps the contents of path to a writer, as a (possibly backwards) hexdump
func HexdumpFile(path string, w io.Writer, format []string) (err error) {
	if path == "-" {
		if err = HexdumpReader(os.Stdin, w, format); err != nil {
			return
		}
	} else {
		var f *os.File
		if f, err = os.Open(path); err != nil {
			return
		}
		defer f.Close()
		if err = HexdumpReader(f, w, format); err != nil {
			return
		}
	}
	return
}

// HexdumpReader dumps bytes read from r to a writer, as a (possibly backwards) hexdump
func HexdumpReader(r io.Reader, w io.Writer, format []string) (err error) {
	b := bufio.NewReader(r)
	var n int
	buf := make([]byte, width)
	offset := uint64(0)
	for {
		if n, err = io.ReadFull(b, buf); n == 0 {
			break
		}
		var s string
		if s, err = Convert(offset, buf, n, width, group, format, !forward); err != nil {
			return
		}
		if _, err = fmt.Fprintf(w, "%s\n", s); err != nil {
			return
		}
		offset += uint64(n)
	}
	if err == io.EOF || err == io.ErrUnexpectedEOF {
		err = nil
	}
	return
}
