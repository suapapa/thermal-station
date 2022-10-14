#!/bin/bash
set -e
GOOS=linux GOARCH=arm go build
scp thermal-station orangepi@opi-hangulclock.local:~/
