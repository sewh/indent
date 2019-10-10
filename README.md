indent
======

A simple command line tool to indent input from stdin and output it to stdout. It was written for use in the acme text editor.

## Installation

```bash
go get github.com/sewh/indent
go install github.com/sewh/indent

# ensure $(go env GOPATH)/bin is in your path...

indent
```

## Usage

```bash
# To indent:

indent +

# To unindent

indent -

# You can also add a number to the end to use spaces instead of tabs. The number you provide
# is the amount of spaces that are added or removed.

indent + 4

indent - 2
```

