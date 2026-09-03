# vim-server

A Vim wrapper that opens files in a single Vim instance.

## Setup

Builds binary and install it to `~/bin`

    make

## Usage

```
Usage: vim-server [VIM_SERVER_OPTIONS] [OPTIONS] [FILE...]

VIM_SERVER_OPTIONS

  -vs-auto, --vs-auto    connect automatically if only one server runned
  -vs-help, --vs-help    show this help message

OPTIONS & FILES

  all other flags and arguments are passed directly to the vim
```
