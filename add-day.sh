#!/bin/bash

# Currently only supports creating days from go template

day="go/$1"

go_mod_init() {
  mkdir -p "$1"
  cp template.go "$1/main.go"
}

go_mod_init "$day/p1"
go_mod_init "$day/p2"
