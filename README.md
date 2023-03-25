A Kubernetes operator built with
[operator-builder](https://github.com/nukleros/operator-builder).

    brew tap nukleros/tap  
    brew install nukleros/tap/operator-builder

    operator-builder create api \
    --workload-config .source-manifests/workload.yaml \
    --controller \
    --resource --force


    IMG=xo66ot/acme-webstore-mgr:0.1.0
    make deploy
    kubectl apply -f config/samples/

## Local Development & Testing

To install the custom resource/s for this operator, make sure you have a
kubeconfig set up for a test cluster, then run:

    make install

To run the controller locally against a test cluster:

    make run

You can then test the operator by creating the sample manifest/s:

    kubectl apply -f config/samples

To clean up:

    make uninstall

## Deploy the Controller Manager

First, set the image:

    export IMG=myrepo/myproject:v0.1.0

Now you can build and push the image:

    make docker-build
    make docker-push

Then deploy:

    make deploy

To clean up:

    make undeploy

