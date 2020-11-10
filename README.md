# polar-controller

Getting started with controllers.

## Generating new controller from empty project

*Note: Following are meant only for generating new controller from empty*
*project. You can also use k8s.io/sample-controller for a skeleton*

Assume `GO111MODULE=on` (although `k8s.io/code-generator` works better when
working in $GOPATH.

``` go
go mod init github.com/<username/<repo>

go get k8s.io/code-generator
```

`k8s.io/code-generator` can be downloaded to `vendor/` on empty repo with
temporarily adding this to e.g. `main.go` (I have added it to `hack/tools.go`:

```
import (
	_ "k8s.io/code-generator"
)
```

`k8s.io/code-generator` requires following files to be set before running:

```
pkg
├── apis
│   └── polarsquad
│       ├── register.go
│       └── v1
│           ├── doc.go
│           ├── register.go
│           └── types.go
```

These files basically tell the code-generator what kind of API objects to create.
For now you should comment the reference to non-existent package (if you're working
with a fresh project) from `pkg/apis/polarsquad/v1/register.go`.

After that we can run `go mod vendor`.

You can then change your resources to whatever name you want and proceed to run
`hack/update-codegen.sh`.

	k8s.io/code-generator is not very user-friendly especially when working on
	fresh projcets

It's also recommended to copy the signals package from `pkg/signals` for your
controller (optional).

	This is used for gracefully handling first shutdown signal

Run `chmod +x vendor/k8s.io/code-generator/generate-groups.sh` and
`hack/update-codegen.sh` generates the files under
`github.com/<username>/<controller>/pkg` so I usually move those back to current
repository, unless you want to write controller under that.

After the necessary files have been generated we can continue on writing our
controller, see `main.go`, `controller.go`, `controller_test.go`.

## Usage

```
# assumes you have a working kubeconfig, not required if operating in-cluster
go build -o polar-controller .
./polar-controller -kubeconfig=$HOME/.kube/config
```

After you have started your controller you should see following messages:

```
E1110 13:14:35.366900    6710 reflector.go:127] pkg/mod/k8s.io/client-go@v0.19.3/tools/cache/reflector.go:156: Failed to watch *v1.PolarDeployment: failed to list *v1.PolarDeployment: the server could not find the requested resource (get polardeployments.polarsquad.com)
```

Since cluster doesn't have resource created for your freshly new custom resource.

```
# create a CustomResourceDefinition
kubectl create -f deploy/crd.yaml
```

After this you should see:

```
I1110 13:15:14.754126    6710 controller.go:148] Starting workers
I1110 13:15:14.754164    6710 controller.go:154] Started workers
```

```
# create a custom resource of type Foo
kubectl create -f deploy/example-polardeployment.yaml

# check custom resources
kubectl get pd

# check deployments created through the custom resource
kubectl get deployments
```

If everything was set correctly you should see something like this in your
controller:

```
I1110 13:16:53.636937    6835 controller.go:212] Successfully synced 'default/example'
I1110 13:16:53.637031    6835 event.go:291] "Event occurred" object="default/example" kind="PolarDeployment" apiVersion="polarsquad.com/v1" type="Normal" reason="Synced" message="PolarDeployment synced successfully"
```

Cool stuff! Everything should be good to go for further hacking.

## Resources

- https://www.openshift.com/blog/kubernetes-deep-dive-code-generation-customresources

- https://trstringer.com/extending-k8s-custom-controllers/#controller-custom-resources
