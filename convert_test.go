package main

import "testing"

func TestConvert(t *testing.T) {
	type args struct {
		offset  uint64
		buf     []byte
		n       int
		width   int
		group   int
		format  []string
		reverse bool
	}
	tests := []struct {
		name  string
		args  args
		wantS string
	}{
		{
			name: "fwd-full-text",
			args: args{
				offset: 0,
				buf:    []byte{1, 2, 3, 4, 5, 6, 7, 8},
				n:      8,
				width:  8,
				group:  1,
				format: DefaultForwardFormat,
			},
			wantS: "       0 | 01 02 03 04 05 06 07 08 | ........",
		},
		{
			name: "fwd-full-group",
			args: args{
				offset: 0,
				buf:    []byte{1, 2, 3, 4, 5, 6, 7, 8},
				n:      8,
				width:  8,
				group:  2,
				format: DefaultForwardFormat,
			},
			wantS: "       0 | 0102 0304 0506 0708 | ........",
		},
		{
			name: "fwd-full-text-offset",
			args: args{
				offset: 64,
				buf:    []byte{1, 2, 3, 4, 5, 6, 7, 8},
				n:      8,
				width:  8,
				group:  1,
				format: DefaultForwardFormat,
			},
			wantS: "      40 | 01 02 03 04 05 06 07 08 | ........",
		},
		{
			name: "fwd-full-notext",
			args: args{
				offset: 0,
				buf:    []byte{1, 2, 3, 4, 5, 6, 7, 8},
				n:      8,
				width:  8,
				group:  1,
				format: []string{"offset", "fhex"},
			},
			wantS: "       0 | 01 02 03 04 05 06 07 08",
		},
		{
			name: "fwd-partial-text",
			args: args{
				offset: 0,
				buf:    []byte{1, 2, 3, 4, 5, 6, 7, 8},
				n:      3,
				width:  8,
				group:  1,
				format: DefaultForwardFormat,
			},
			wantS: "       0 | 01 02 03                | ...     ",
		},
		{
			name: "fwd-empty-text",
			args: args{
				offset: 0,
				buf:    []byte{1, 2, 3, 4, 5, 6, 7, 8},
				n:      0,
				width:  8,
				group:  1,
				format: DefaultForwardFormat,
			},
			wantS: "       0 |                         |         ",
		},
		{
			name: "back-full-text",
			args: args{
				offset: 0,
				buf:    []byte{1, 2, 3, 4, 5, 6, 7, 8},
				n:      8,
				width:  8,
				group:  1,
				format: DefaultBackwardFormat,
			},
			wantS: "08 07 06 05 04 03 02 01 |        0 | ........",
		},
		{
			name: "back-full-notext",
			args: args{
				offset:  0,
				buf:     []byte{1, 2, 3, 4, 5, 6, 7, 8},
				n:       8,
				width:   8,
				group:   1,
				format:  []string{"bhex", "offset"},
				reverse: true,
			},
			wantS: "08 07 06 05 04 03 02 01 |        0",
		},
		{
			name: "back-full-text-group",
			args: args{
				offset: 0,
				buf:    []byte{1, 2, 3, 4, 5, 6, 7, 8},
				n:      8,
				width:  8,
				group:  2,
				format: DefaultBackwardFormat,
			},
			wantS: "0807 0605 0403 0201 |        0 | ........",
		},
		{
			name: "back-partial-text",
			args: args{
				offset: 0,
				buf:    []byte{1, 2, 3, 4, 5, 6, 7, 8},
				n:      3,
				width:  8,
				group:  1,
				format: DefaultBackwardFormat,
			},
			wantS: "               03 02 01 |        0 |      ...",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotS, gotErr := Convert(tt.args.offset, tt.args.buf, tt.args.n, tt.args.width, tt.args.group, tt.args.format)
			if gotErr != nil {
				t.Errorf("Convert() err = %v, want %v", gotErr, nil)
			}
			if gotS != tt.wantS {
				t.Errorf("Convert()\n  got s = %q\n  want  = %q", gotS, tt.wantS)
			}
		})
	}
}
