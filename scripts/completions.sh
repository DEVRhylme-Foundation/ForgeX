#!/bin/sh
set -e

mkdir -p completions
for sh in bash zsh fish; do
  go run main.go completion "$sh" > "completions/forgex.$sh"
done
