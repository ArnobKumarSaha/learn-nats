#!/bin/sh

set -x

mkdir -p trivy/db
cd trivy/db
nats -s nats:4222 object get trivy db.tar.gz
tar xvf db.tar.gz
rm db.tar.gz


/root/.cache/tv version
/bin/oras version