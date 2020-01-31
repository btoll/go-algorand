#!/usr/bin/env bash
# shellcheck disable=1090

echo
date "+build_release begin PACKAGE stage %Y%m%d_%H%M%S"
echo

set -ex

. "${HOME}/build_env"

# Order matters (deb -> rpm)!
"${HOME}"/go/src/github.com/algorand/go-algorand/scripts/release/deb/package.sh

sg docker "docker run --rm --env-file ${HOME}/build_env_docker --mount type=bind,src=${HOME},dst=/root/subhome algocentosbuild /root/subhome/go/src/github.com/algorand/go-algorand/scripts/release/rpm/package.sh"

echo
date "+build_release end PACKAGE stage %Y%m%d_%H%M%S"
echo

