@echo off
set CGO_CXXFLAGS="-I%cd%\libs\webview2\build\native\include"
go build -ldflags="-s -w -H windowsgui" -o build/basic.exe main.go
