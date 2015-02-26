# ARS
New Mexico Air Airline Reservation System

## Info
This source tree is comprised of two main components:

 Go based web server
 AngularJS based Single-Page-Application


 The directory tree:

     go
     go/main.go
     ...
     web
     web/styles
     ...


## Building

 Set GOPATH to point to the 'go' directory, e.g.

     $ export GOPATH="/Users/colby/Spring2015/CS460/repos/ARS/go"

 Then, move to that directory and ensure dependencies are loaded:

     $ cd $GOPATH
     $ go get github.com/gorilla/mux

 Then, build the ARS web server binary:

     $ cd ars-server
     $ go build

 After this, it will be runnable:

     $ ./ars-server -directory="../../web" -port=3500

 A UNIXy launcher script is available in the same directory as ars-server, called 'ars-launch'

     $ ./ars-launch

 This script will do essentially the same as the longer command-line above.

