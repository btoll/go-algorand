#!/usr/bin/env bash
# shellcheck disable=2029

# This is temporary (taken from `compute_branch_channel.sh`).
if [ "$1" = "master" ]; then
    CHANNEL=master
elif [ "$1" = "rel/nightly" ]; then
    CHANNEL=nightly
elif [ "$1" = "rel/stable" ]; then
    CHANNEL=stable
elif [ "$1" = "rel/beta" ]; then
    CHANNEL=beta
else
    CHANNEL=dev
fi

# Path(s) are relative to the root of the Jenkins workspace.
INSTANCE=$(cat scripts/release/common/ec2/tmp/instance)
#BUCKET="$1"
BRANCH="$1"
RELEASE="$4"

rm -rf pkg/* && mkdir -p pkg/"$FULLVERSION"
#aws s3 sync s3://"$BUCKET"/"$CHANNEL"/"$RELEASE" pkg/ --exclude "*" --include "*.deb" --include "*.rpm"
aws s3 sync s3://ben-test-2.0.3/"$CHANNEL"/"$RELEASE" pkg/ --exclude "*" --include "*.deb" --include "*.rpm"
ssh -i ReleaseBuildInstanceKey.pem -A ubuntu@"$INSTANCE" mkdir node_pkg
# Upload the packages and their signatures.
scp -i ReleaseBuildInstanceKey.pem -o StrictHostKeyChecking=no -r pkg/* ubuntu@"$INSTANCE":~/node_pkg/

aws s3 cp s3://algorand-devops-misc/tools/gnupg2.2.9_centos7_amd64.tar.bz2 .
scp -i ReleaseBuildInstanceKey.pem -o StrictHostKeyChecking=no -r gnupg2.2.9_centos7_amd64.tar.bz2 ubuntu@"$INSTANCE":

#scp -i ReleaseBuildInstanceKey.pem -o StrictHostKeyChecking=no -r scripts/release/test/stage/setup/task.sh ubuntu@"$INSTANCE":
scp -i ReleaseBuildInstanceKey.pem -o StrictHostKeyChecking=no -r scripts/release/common/setup.sh ubuntu@"$INSTANCE":
#ssh -i ReleaseBuildInstanceKey.pem -A ubuntu@"$INSTANCE" bash task.sh "$BRANCH" "$CHANNEL" "$RELEASE"
ssh -i ReleaseBuildInstanceKey.pem -A ubuntu@"$INSTANCE" bash setup.sh "$BRANCH" "$CHANNEL" "$RELEASE"

