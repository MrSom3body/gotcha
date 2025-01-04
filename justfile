alias b := build

default:
    @just --list

build:
    CGO_ENABLED=0 go build -ldflags="-s -w -extldflags '-static'"

bench:
    nix build
    hyperfine -N -w 50 gotcha ./result/bin/gotcha
