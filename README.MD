# Suma/UYUNI tools wrapper

Goal is to provide a single pain of glass to call all existing suma/Uyuni command line tools.
Currently we have several tools without any name convention and is hard to find them all and what they do.

The project provides a wrapper to call those tools, and we can have an overlook of what tools exists.

## Architecture

Declare all available tools.
The tool command will be called, and Go stdin, stdout and stderror will link with the tools stds.
With this approach the tool is running from Go but all input and output will be mapped to the underlying command.

It's agnostic of the language used in the underlying tool. It was tested with a tool in perl and another in python.

## Local tools

It also supports the development of local tools, which instead of call a remote command call a Go method inside the project.
If we need to develop a new tool it can be done inside this project.

One sample tool is available to server as example.

## Install

1. On project run `go build -o uyuni`
2. Copy the build file `uyuni` to your suma/uyuni server
3. Have fun

## TODO
- Use internationalization in the tool description

# DEMO

[![asciicast](https://asciinema.org/a/SrgBVVvoQUklbbbVSbGE3vtod.svg)](https://asciinema.org/a/SrgBVVvoQUklbbbVSbGE3vtod)