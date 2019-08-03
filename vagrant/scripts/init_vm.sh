#!/bin/bash

echo Welcome to VM settins ...

add-apt-repository ppa:longsleep/golang-backports
apt-get update
apt-get install sysdig golang-go debootstrap -y

## boot
echo ...... debootstrap ...
LINUX_HOME=/home/vagrant/ubuntu_xenial_1604
mkdir ${LINUX_HOME}
debootstrap --arch=amd64 xenial ${LINUX_HOME} http://archive.ubuntu.com/ubuntu/