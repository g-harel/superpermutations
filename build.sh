#!/bin/sh

OUTDIR=build

rm -rf ./$OUTDIR

export GOARCH=amd64
GOOS=linux  ; go build -o $OUTDIR/superpermutations-linux-amd64
GOOS=darwin ; go build -o $OUTDIR/superpermutations-darwin-amd64
GOOS=windows; go build -o $OUTDIR/superpermutations-windows-amd64.exe
