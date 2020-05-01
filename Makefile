INSTALL=install
prefix=/usr/local
exec_prefix=$(prefix)
bindir=$(exec_prefix)/bin
datadir=$(prefix)/share
mandir=$(datadir)/man
man1dir=$(mandir)/man1

all: bhd

bhd: bhd.go convert.go version.go
	go build -o bhd .

check: all
	go test ./...

version.go: bhd.go convert.go Makefile make-version
	./make-version $@

install:
	$(INSTALL) -m 755 bhd $(bindir)/bhd
	$(INSTALL) -m 644 bhd.1 $(man1dir)/bhd.1

uninstall:
	rm -f $(bindir)/bhd
	rm -f $(man1dir)/bhd.1

clean distclean maintainer-clean:
	rm -f version.go
	rm -f bhd
