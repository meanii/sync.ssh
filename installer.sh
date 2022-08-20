#!/bin/bash

#
# Copyright © 2022 Anil Chauhan <https://github.com/meanii>
#
# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with this program. If not, see <http://www.gnu.org/licenses/>.

## Installing sync.ssh from the source
echo "Installing sync.ssh"
go install .

## Checking, if golang weather installed or not
if ! command -v go &> /dev/null
then
    echo "golang could not be found!"
    exit
fi

## Checking got GOPATH ENV
if [[ -z "${GOPATH}" ]];
then
  echo "GOPATH ENV couldn't found!"
  exit
fi

## Preparing daemon file
echo "setting up daemon services"
SYNCSSH_PATH=/usr/local/bin/sync.ssh
SEVICE_FILE_BIN=service/sync_bin.service
SEVICE_FILE=service/sync.service
TIMER_FILE=service/sync.timer

SYSTEMD_PATH=/lib/systemd/system/

USERNAME=$(whoami)
echo "adding daemon for $USERNAME user!"

sudo cat $SEVICE_FILE_BIN | sed "s|SYNCSSH_PATH|$SYNCSSH_PATH|g; s|USERNAME|$USERNAME|g;" > $SEVICE_FILE

## Coping Prepared Daemon file to the /lib/systemc/system
echo "copying demon service files"
sudo cp $SEVICE_FILE $SYSTEMD_PATH
sudo cp $TIMER_FILE $SYSTEMD_PATH

## Reloading, Enabling and Starting the daemon services
sudo systemctl daemon-reload
sudo systemctl enable sync
sudo systemctl enable sync.timer

sudo systemctl start sync
sudo systemctl start sync.timer

## Checking, if sync.ssh is installed properly or not
if ! command -v sync.ssh &> /dev/null
then
    echo "something went wrong while installing sync.ssh!"
    exit
else
  echo "sync.ssh installation has been done!"
  echo "try sync.ssh --help"
fi