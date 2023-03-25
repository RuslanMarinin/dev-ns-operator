#! /bin/sh


# operator-builder init-config standalone -p .source-manifests/workload.yaml --force;

# operator-builder init --workload-config .source-manifests/workload.yaml;

operator-builder create api \
    --workload-config .source-manifests/workload.yaml \
    --controller \
    --resource --force

IMG=xo66ot/acme-webstore-mgr:0.1.0

make deploy

kubectl apply -f config/samples/