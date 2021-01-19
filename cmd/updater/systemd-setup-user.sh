#!/bin/bash
#
# Run this script to set up the systemd user service for algod.
# The argument is the username for whom systemd will install the service.

set -e

setup_user() {
    local user="$1"
    local userline

    if ! userline=$(getent passwd "$user"); then
        echo "[ERROR] \`$USER\' not found on system. Aborting..."
        exit 1
    else
        homedir=$(awk -F: '{ print $6 }' <<< "$userline")
    fi

    mkdir -p "$homedir/.config/systemd/user"
    cp "${SCRIPTPATH}/algorand@.service.template-user" \
        "$homedir/.config/systemd/user/algorand@.service"

    systemctl --user daemon-reload
}


if [ "$#" != 1 ]; then
    echo "Usage: $0 username"
    exit 1
fi

SCRIPTPATH="$( cd "$(dirname "$0")" ; pwd -P )"
USER="$1"

if ! id -u "${USER}"> /dev/null; then
    echo "$0 [ERROR] Username \`$USER\` does not exist on system"
    exit 1
fi

setup_user "${USER}"

