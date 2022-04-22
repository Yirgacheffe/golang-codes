#!/bin/sh

go build -o producer ./cmd/producer/main.go
go build -o consumer ./cmd/consumer/main.go

go build -o nisemono ./cmd/nisemono/main.go
