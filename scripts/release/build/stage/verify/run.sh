#!/usr/bin/env bash

set -ex

trap 'bash ./scripts/release/common/ec2/shutdown.sh' ERR

# Path(s) are relative to the root of the Jenkins workspace.
ssh -i ReleaseBuildInstanceKey.pem -A ubuntu@"$(cat scripts/release/common/ec2/tmp/instance)" bash ben-branch/scripts/release/build/stage/verify/task.sh

