#!/bin/bash
# This is a demo script that runs the test and curls commands, run this after installation.

echo "Running test..."
#go test ./...
go test -v -bench . -benchmem -cover -race -parallel 10000 .

echo ""

echo "Getting the sysinfo of the service..."
curl -X POST -H "Content-Type: application/json" \
-d '{"type": "sysinfo"}' \
http://localhost:8080/execute

echo ""

echo "Pinging Google..."
curl -X POST -H "Content-Type: application/json" \
-d '{"type": "ping", "payload": "google.com"}' \
http://localhost:8080/execute