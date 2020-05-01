# bhd - backwards hexdump

## What?

Hexdumps a file, with the bytes ordered from right to left instead of left to right.

## Seriously, what?

The use case is when the file is a binary file full of little-endian fields.
By reversing the byte order so it matches the order of digits within each byte,
it becomes easier to read the field values by eye.

The `--group` option allows bytes to be grouped together in a consistent way,
to make reading highly structured files easier.

## Installation

[Install a Go compiler](https://golang.org/dl/) if you don't have one already.

    $ make check
    $ sudo make install

## Documentation

    $ man bhd
    $ bhd --help

## Copying

© 2020 Richard Kettlewell

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see [http://www.gnu.org/licenses/](http://www.gnu.org/licenses/).
