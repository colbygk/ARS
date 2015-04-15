@echo off
cd ..
set GOPATH=%CD%
cd ars-server

@echo on
start cmd /c ars-server -directory=../../web