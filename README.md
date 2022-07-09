# Elgato Light Control

A simple Go library to control Elgato lights.

```
go get github.com/seanpfeifer/elgato-light-control
```

The sample at `./cmd/elgato-lighttoggle` will find a ring light on the network and toggle it on/off. To try this sample, install via:
```
go install ./cmd/elgato-lighttoggle
```

## Installation

For Windows, you can build or install such that the console doesn't show up when the sample application runs, if you want to eg run it via a macro or Stream Deck executable.
```
go build -ldflags "-H=windowsgui" ./cmd/elgato-lighttoggle
go install -ldflags "-H=windowsgui" ./cmd/elgato-lighttoggle
```

## Why?

For fun. I wanted to try a relatively minimalist way of controlling my lights without requiring Elgato Control Center.
