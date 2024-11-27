#!/bin/bash

# Load environment variables from .env file
export $(cat .env | xargs)

# Convert each env variable into an ldflag for Go build
LD_FLAGS=""
for var in $(cat .env | cut -d= -f1); do
  LD_FLAGS="$LD_FLAGS -X main.$var=${!var}"
done

export LD_FLAGS