@echo off
cd ..
set GOPATH=%CD%
cd ars-server

@echo on
go get github.com/gorilla/mux
go build
start cmd /c ars-server -directory=../../web
exit