#!/bin/bash

gofmt -w .
STUB_MODE=true SECRET_TOKEN=123456 go run .
