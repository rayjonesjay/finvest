#!/usr/bin/env bash

echo "hello world"
flag=$1

# Check if the flag is provided
if [ -z "$flag" ]; then
    printf "Please provide a flag: -l for login, -s for signup\n"
    exit 1
else
    # Check the value of the flag
    if [ "$flag" == "-l" ]; then
        curl localhost:8080/login
    elif [ "$flag" == "-s" ]; then
        curl localhost:8080/signup
    else
        printf "Invalid flag. Use -l for login or -s for signup\n"
        exit 1
    fi
fi

