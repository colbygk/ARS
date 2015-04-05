#!/usr/bin/env bash

sudo apt-get update 2> /dev/null
sudo apt-get install -y make 2> /dev/null
sudo apt-get install -y vim 2> /dev/null
sudo apt-get install -y openssl 2> /dev/null
sudo apt-get install -y git 2> /dev/null

sudo debconf-set-selections <<< 'golang-go golang-go/dashboard boolean false' 
sudo apt-get install -y golang-go 2> /dev/null

useradd ars

sudo debconf-set-selections <<< 'mysql-server mysql-server/root_password password @RSdev'
sudo debconf-set-selections <<< 'mysql-server mysql-server/root_password_again password @RSdev'
sudo apt-get install -y mysql-server 2> /dev/null
sudo apt-get install -y mysql-client 2> /dev/null

if [ ! -f /var/log/dbinstalled ];
then
    echo "CREATE USER 'ars'@'localhost' IDENTIFIED BY 'ARSePassW0rd'" | mysql -uroot -p@RSdev
    echo "CREATE DATABASE ARSdb" | mysql -uroot -p@RSdev
    echo "GRANT ALL ON ARSdb.* TO 'ars'@'localhost'" | mysql -uroot -p@RSdev
    echo "flush privileges" | mysql -uroot -p@RSdev
    touch /var/log/dbinstalled
fi

export TM8DEV=/ARS
if [ -d ${TM8DEV} ];
then
  cd ${TM8DEV}
  if [ -f schema/ars.sql ];
  then
    mysql -uroot -p@RSdev ARSdb < schema/ars.sql
  fi
  if [ ! -d go/src/github.com ];
  then
    export GOPATH=${TM8DEV}/go
    echo "export GOPATH=${GOPATH}" >> ~vagrant/.bashrc
    cd $GOPATH
    go get github.com/gorilla/mux
    go get github.com/go-sql-driver/mysql
    cd ars-server
    go build
    ./ars-launch >& ars-server.log &
  fi
fi
