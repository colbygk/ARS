# ARS
New Mexico Air Airline Reservation System

This repository covers the development efforts of
Team 8 in CS 460 Spring 2015.

This repository is current with Week 11 of the
class and provides a separate development environment
in the form of a [Vagrant](http://vagrantup.com) vm setup.

## General Info
This source tree is comprised of four main components:

* Vagrant Development Environment Configuration
* MySQL schema and test data
* Go based web server
* AngularJS based Single-Page-Application

The directory tree:

    Vagrantfile
    bootstrap.sh
    ...
    schema
    schema/ars.sql
    ...
    go
    go/main.go
    ...
    web
    web/css
    ...

## Using Vagrant to build an isolated development Environment

You should install (for your OS):

  [VirtualBox](http://virtualbox.org) and [Vagrant](http://vagrantup.org)

Once you have cloned this repository, cd into the main directory 'ARS',
then,
    
    $ vagrant up

This should start up vagrant, which will download a virtual machine
instance of Ubuntu 14.04, install mysql, go, and various other
required tools in that instance. Depending on your network connection
this should take between 2-10 minutes. Note, this will consume
about ~350MB of diskspace and once running will present a virual
machine with 512MB of RAM and 1 processor.

After vagrantup completes, you should see the message
*ARS - Airline Reservation System Development Install*

After the VM has started, you should be able to access
http://localhost:3500 on the host computer, which will be
pointing at the 'web' directory of the repository.

To access the VM, you can do,

    $ vagrant ssh

This should connect you to the VM. You can then do:

    $ cd /ARS

Now you will be in the directory that contains the repository you
cloned and where you launched this instance of Vagrant.  A copy of
MySQL is running and testing copy of the database has been loaded.

### Developing from within the Vagrant Instance

The repo in /ARS will be the one that you cloned earlier. However,
once connected to the Vagrant instance, it will not have access
to your local ssh keys, nor your particular user information.

Therefore, you should make sure that your user info has been set:

    $ git config --global user.name "First Last"
    $ git config --global user.email your@email.com

I recommend that you then [add an ssh-key](https://help.github.com/articles/generating-ssh-keys/)
to allow easy git operations or if you used HTTPS when cloning, [caching
the credentials](https://help.github.com/articles/caching-your-github-password-in-git/)
to ease updating.

## MySQL Testing Database

Within the VM, MySQL has been loaded with test data in the
database named ARSdb. This consists of two tables, 'flights'
and 'airports'.  Note that the schema/ars.sql file will first
attempt to drop those tables, create them and then load the
test data.

The test data is based off of the JSON files web/flights.json
and web/airports.json.  These files were converted into
SQL insert statements via the script web/convert_json_to_sql.js
This can be re-run using nodejs.

Note that the testing flights data are currently only
covering 04/27/2015-04/28/2015

## Building ars-server

If you wish to rebuild the ars-server, you should have GOPATH
set (already done for you under the Vagrant VM, i.e. you should
only need to do this if you're not using Vagrant).
e.g.

    $ export GOPATH="/Users/colby/Spring2015/CS460/repos/ARS/go"

Then, move to that directory and ensure dependencies are loaded (already
loaded under Vagrant):

     $ cd $GOPATH
     $ go get github.com/gorilla/mux
     $ go get github.com/go-sql-driver/mysql
     $ go get gopkg.in/gorp.v1

Then, build the ARS web server binary:

     $ cd ars-server
     $ go build

After this, it will be runnable:

     $ ./ars-server -directory="../../web" -port=3500

A UNIXy launcher script is available in the same directory as ars-server, called 'ars-launch'

     $ ./ars-launch

A Windowsy launcher script is also available .

This script will do essentially the same as the longer command-line above.

