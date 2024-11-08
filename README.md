# Command Executor
A basic Ubunut command execution application for executing `ping` and retrieving system info.

## Installation
```Go
go build -o cmd-executor

```bash
sudo ./install.sh

```bash
sudo dpkg -i cmd-executor_1.0_amd64.deb

```Testing
go test ./...

```Curl
curl -X POST -H "Content-Type: application/json" \
-d '{"type": "sysinfo"}' \
http://localhost:8080/execute

curl -X POST -H "Content-Type: application/json" \
-d '{"type": "ping", "payload": "google.com"}' \
http://localhost:8080/execute

```bash
./demo.sh







