#!/bin/bash
./http autoupdate.v3 &
cp autoupdate.v2 autoupdate
./autoupdate update
echo "updated!!!"
echo "start new version!"
./autoupdate

