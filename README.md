# Clipboard for Go

Provide copying and pasting to the Clipboard for Go.

Build:

    $ go get pkg.nimblebun.works/clipboard

Platforms:

- OSX
- Windows 7 (probably work on other Windows)
- Linux, Unix (requires 'xclip' or 'xsel' command to be installed)

Document:

- http://godoc.org/pkg.nimblebun.works/clipboard

Notes:

- Text string only
- UTF-8 text encoding only (no conversion)

TODO:

- Clipboard watcher(?)

## Commands:

paste shell command:

    $ go get pkg.nimblebun.works/clipboard/cmd/gopaste
    $ # example:
    $ gopaste > document.txt

copy shell command:

    $ go get pkg.nimblebun.works/clipboard/cmd/gocopy
    $ # example:
    $ cat document.txt | gocopy
