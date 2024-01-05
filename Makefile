update-codegen:
	go mod vendor
	chmod +x ./hack/update-codegen.sh
	chmod +x /root/code/gva-controller/vendor/k8s.io/code-generator/generate-internal-groups.sh
	./hack/update-codegen.sh

verify-codegen:
	chmod +x ./hack/verify-codegen.sh
	./hack/verify-codegen.sh

clean:
	rm -rf vendor  ./pkg/client  ./pkg/apis/deploy/v1/zz_generated.deepcopy.go
