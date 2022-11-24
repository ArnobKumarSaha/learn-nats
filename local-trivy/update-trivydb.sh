#!/bin/sh

set -x

# https://aquasecurity.github.io/trivy/v0.23.0/advanced/air-gap/
# busybox is missing ca-certificates
oras pull ghcr.io/aquasecurity/trivy-db:latest --insecure

# oras pulls in tar format. File will be named `db.tar.gz`

nats -s nats:4222 object add trivy || true
nats -s nats:4222 object put trivy db.tar.gz -f
nats -s nats:4222 object ls trivy
nats -s nats:4222 object info trivy
