#!/bin/bash
killall dlv
go build -gcflags='-N -l' ../ && dlv --listen=:2345 --headless=true --api-version=2 exec ./surebetSearch
