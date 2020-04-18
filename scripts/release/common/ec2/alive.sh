#!/usr/bin/env bash

GREEN_FG=$(echo -en "\e[32m")
YELLOW_FG=$(echo -en "\e[33m")
END_FG_COLOR=$(echo -en "\e[39m")

echo "$YELLOW_FG[$0]$END_FG_COLOR: Waiting for SSH connection"
end=$((SECONDS+90))
while [ $SECONDS -lt $end ]
do
    if ssh -i algorand_baseline.pem -o "StrictHostKeyChecking no" "ubuntu@$1" "uname"
    then
        echo "$GREEN_FG[$0]$END_FG_COLOR: SSH connection ready"
        exit 0
    fi

    sleep 1s
done

