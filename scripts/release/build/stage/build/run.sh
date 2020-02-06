#!/usr/bin/env bash
# shellcheck disable=2029

# Path(s) are relative to the root of the Jenkins workspace.
#ssh -i ReleaseBuildInstanceKey.pem -A ubuntu@"$(cat scripts/release/common/ec2/tmp/instance)" bash go/src/github.com/algorand/go-algorand/scripts/release/build/stage/build/task.sh "$1" "$2"
ssh -i ReleaseBuildInstanceKey.pem -A ubuntu@"$(cat scripts/release/common/ec2/tmp/instance)" bash ben-branch/scripts/release/build/stage/build/task.sh "$1" "$2"

