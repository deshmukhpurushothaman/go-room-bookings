#!/bin/bash

go build -o go-room-bookings ./cmd/web
./go-room-bookings -dbname=go-room-bookings -dbuser=postgres -dbpass=deshmukh123 -cache=false -production=false
