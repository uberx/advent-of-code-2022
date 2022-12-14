#!/usr/bin/env bash

parent_path=$(
    cd "$(dirname "${BASH_SOURCE[0]}")"
    pwd -P
)

(
    cd ${parent_path} &&
        (
            echo -e "ADVENT OF CODE 2022\n"
            for number in {1..23}; do
                echo "========== Part ${number} =========="
                go run day${number}/main.go
                echo
            done
        )
)
