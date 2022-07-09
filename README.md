# elgato-light-control

Can build or install such that the console doesn't show up:
```
go build -ldflags "-H=windowsgui" ./cmd/elgato-lighttoggle
go install -ldflags "-H=windowsgui" ./cmd/elgato-lighttoggle
```
