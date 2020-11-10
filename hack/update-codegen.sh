#!/usr/bin/env bash

# This script is quick and ugly!

go mod vendor

chmod +x vendor/k8s.io/code-generator/generate-groups.sh

# Replace polarsquad:v1 with your own API name
# generate-groups.sh works better when working in $GOPATH
vendor/k8s.io/code-generator/generate-groups.sh \
	all \
	github.com/topikettunen/polar-controller/pkg/client \
	github.com/topikettunen/polar-controller/pkg/apis \
	polarsquad:v1 \
	--go-header-file hack/custom-boilerplate.go.txt

echo "Copying generated files under pkg/..."
cp -r github.com/topikettunen/polar-controller/pkg/client pkg
cp github.com/topikettunen/polar-controller/pkg/apis/polarsquad/v1/zz_generated.deepcopy.go pkg/apis/polarsquad/v1

echo "Remove unnecessary files..."

rm -r github.com/ vendor
