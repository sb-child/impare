#!/bin/sh
rm libpare/main.go
rust2go-cli --src src/libpare_sys/sys.rs --dst libpare/main.go
