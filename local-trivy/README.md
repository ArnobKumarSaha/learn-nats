# install nats

```
> helm upgrade -i nats nats/nats --set nats.jetstream.enabled=true
> k get svc
> k port-forward svc/nats 4222:4222

# from another terminal
> nats account info
```

# upload trivy to nats

- https://aquasecurity.github.io/trivy/v0.18.3/air-gap/
- https://docs.nats.io/nats-concepts/jetstream/obj_store/obj_walkthrough

```
> nats object add trivy || true

> wget https://github.com/aquasecurity/trivy-db/releases/latest/download/trivy-offline.db.tgz
> mv trivy-offline.db.tgz db.tar.gz
> nats object put trivy db.tar.gz
> nats object ls trivy
> nats object info trivy
```

# upload trivy to file-server (of the operator filesystem)
```
Need to run these from `local-trivy` folder.
> wget https://github.com/aquasecurity/trivy-db/releases/latest/download/trivy-offline.db.tgz
> mv trivy-offline.db.tgz db.tar.gz
> kubectl curl -k -X POST -F file=@/home/arnob/go/src/github.com/Arnobkumarsaha/learn-nats/local-trivy/db.tar.gz  https://scanner-0:8443/files/trivy -n kubeops
```

# Build natscli

```
docker build -t appscode/natscli -f natscli-dockerfile .
kind load docker-image appscode/natscli
```

# Upload trivydb to nats

```
k apply -f refresh.yaml
```

- Run this using a cron job periodically

# Scan a local image

```
k apply -f scan.yaml
```
