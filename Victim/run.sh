#!/bin/bash

GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-H windowsgui" -o norvoinc.exe main.go
