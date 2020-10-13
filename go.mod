module github.com/pbar1/gq

go 1.15

require (
	github.com/Masterminds/goutils v1.1.0 // indirect
	github.com/Masterminds/semver v1.5.0 // indirect
	github.com/Masterminds/sprig v2.22.0+incompatible
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/google/uuid v1.1.2 // indirect
	github.com/hashicorp/hcl v1.0.0
	github.com/huandu/xstrings v1.3.2 // indirect
	github.com/imdario/mergo v0.3.11 // indirect
	github.com/json-iterator/go v1.1.10
	github.com/mitchellh/copystructure v1.0.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.1 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/pelletier/go-toml v1.8.0
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.6.1 // indirect
	golang.org/x/crypto v0.0.0-20200820211705-5c72a883971a // indirect
	gopkg.in/yaml.v2 v2.3.0
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/kubernetes v1.19.2 // indirect
)

replace k8s.io/api => k8s.io/api v0.19.2

replace k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.19.2

replace k8s.io/apimachinery => k8s.io/apimachinery v0.19.3-rc.0

replace k8s.io/apiserver => k8s.io/apiserver v0.19.2

replace k8s.io/cli-runtime => k8s.io/cli-runtime v0.19.2

replace k8s.io/client-go => k8s.io/client-go v0.19.2

replace k8s.io/cloud-provider => k8s.io/cloud-provider v0.19.2

replace k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.19.2

replace k8s.io/code-generator => k8s.io/code-generator v0.19.3-rc.0

replace k8s.io/component-base => k8s.io/component-base v0.19.2

replace k8s.io/controller-manager => k8s.io/controller-manager v0.19.3-rc.0

replace k8s.io/cri-api => k8s.io/cri-api v0.19.3-rc.0

replace k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.19.2

replace k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.19.2

replace k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.19.2

replace k8s.io/kube-proxy => k8s.io/kube-proxy v0.19.2

replace k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.19.2

replace k8s.io/kubectl => k8s.io/kubectl v0.19.2

replace k8s.io/kubelet => k8s.io/kubelet v0.19.2

replace k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.19.2

replace k8s.io/metrics => k8s.io/metrics v0.19.2

replace k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.19.2

replace k8s.io/sample-cli-plugin => k8s.io/sample-cli-plugin v0.19.2

replace k8s.io/sample-controller => k8s.io/sample-controller v0.19.2
