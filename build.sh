#!/bin/bash
GOOS=darwin GOARCH=amd64 go build -o ftx-auto-lend-mac
GOOS=linux GOARCH=386 go build -o ftx-auto-lend-linux
GOOS=windows GOARCH=386 go build -o ftx-auto-lend-win.exe

zip ftx-auto-lend-mac.zip ftx-auto-lend-mac
zip ftx-auto-lend-linux.zip ftx-auto-lend-linux
zip ftx-auto-lend-win.zip ftx-auto-lend-win.exe

rm ftx-auto-lend-mac ftx-auto-lend-linux ftx-auto-lend-win.exe