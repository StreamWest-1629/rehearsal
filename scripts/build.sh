#!/bin/sh

cd /src
go build -o ./bin/rehearsal ./cmd/rehearsal

./bin/rehearsal
