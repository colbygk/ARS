@echo off
cd ..
set GOPATH=%CD%
cd ars-interface

@echo on
start cmd /c ars-interface -directory=../../web