#!/bin/sh


pwd

CGO_LDFLAGS="-static -w -s -Wl,--dynamic-linker=/vorteil/ld-linux-x86-64.so.2 -Wl,-rpath,/vorteil" /usr/local/go/bin/go build -v -tags netgo -o /tcpdumpout/tcpdump
