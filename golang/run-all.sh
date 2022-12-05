#!/usr/bin/env bash

parent_path=$( cd "$(dirname "${BASH_SOURCE[0]}")" ; pwd -P )

for number in {1..5}
do
    echo "========== Part ${number} =========="
    go run ${parent_path}/day${number}/main.go
    echo
done
