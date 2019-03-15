#!/bin/sh

echo "Rebuild the application..."
go get -d -v ./...
go install -v ./...
echo "Exit container."
exit
