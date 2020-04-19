#!/usr/bin/env bash

if [ -z "$KEYFILE" ] || [ -z "$JENKINS" ]
then
    echo Missing KEYFILE or JENKINS environment variables.
    exit 1
fi

INSTANCE=$(ssh -i "$KEYFILE" "$JENKINS" sudo cat /opt/jenkins/workspace/build-packages/scripts/release/common/ec2/tmp/instance)
gpgp=$(find /usr/lib/gnupg{2,,1} -type f -name gpg-preset-passphrase 2> /dev/null)

# Here we need to grab the signing subkey, hence `tail -1`.
KEYGRIP=$(gpg -K --with-keygrip --textmode dev@algorand.com | grep Keygrip | tail -1 | awk '{ print $3 }')
echo "enter dev@ password"
$gpgp --verbose --preset "$KEYGRIP"

KEYGRIP=$(gpg -K --with-keygrip --textmode rpm@algorand.com | grep Keygrip | head -1 | awk '{ print $3 }')
echo "enter rpm@ password"
$gpgp --verbose --preset "$KEYGRIP"

REMOTE_GPG_SOCKET=$(ssh -o StrictHostKeyChecking=no -i ReleaseBuildInstanceKey.pem ubuntu@"$INSTANCE" gpgbin/remote_gpg_socket)
LOCAL_GPG_SOCKET=$(gpgconf --list-dirs | grep agent-socket | awk -F: '{ print $2 }')

gpg --export -a dev@algorand.com > /tmp/dev.pub
gpg --export -a rpm@algorand.com > /tmp/rpm.pub

scp -o StrictHostKeyChecking=no -i ReleaseBuildInstanceKey.pem -p /tmp/{dev,rpm}.pub ubuntu@"$INSTANCE":~/keys/
ssh -o StrictHostKeyChecking=no -i ReleaseBuildInstanceKey.pem ubuntu@"$INSTANCE" << EOF
    gpg --import keys/dev.pub
    gpg --import keys/rpm.pub
    echo "SIGNING_KEY_ADDR=dev@algorand.com" >> build_env
EOF

ssh -o StrictHostKeyChecking=no -i ReleaseBuildInstanceKey.pem -A -R "$REMOTE_GPG_SOCKET:$LOCAL_GPG_SOCKET" ubuntu@"$INSTANCE"

