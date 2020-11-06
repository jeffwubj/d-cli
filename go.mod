module github.com/jeffwubj/d-cli

go 1.15

require (
	cloud.google.com/go v0.57.0 // indirect
	github.com/Microsoft/go-winio v0.4.15-0.20200908182639-5b44b70ab3ab // indirect
	github.com/Microsoft/hcsshim/test v0.0.0-20200826032352-301c83a30e7c // indirect
	github.com/agl/ed25519 v0.0.0-20170116200512-5312a6153412 // indirect
	github.com/bitly/go-hostpool v0.1.0 // indirect
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/cloudflare/cfssl v1.5.0 // indirect
	github.com/containerd/cgroups v0.0.0-20200710171044-318312a37340 // indirect
	github.com/containerd/containerd v1.4.1-0.20200903181227-d4e78200d6da // indirect
	github.com/containerd/continuity v0.0.0-20200710164510-efbc4488d8fe // indirect
	github.com/containerd/go-runc v0.0.0-20201020171139-16b287bc67d0 // indirect
	github.com/davecgh/go-spew v1.1.1
	github.com/docker/cli v20.10.0-beta1.0.20201103165149-c20be83d6b34+incompatible
	github.com/docker/compose-on-kubernetes v0.4.24 // indirect
	github.com/docker/docker v20.10.0-beta1.0.20201030232932-c2cc352355d4+incompatible
	github.com/docker/docker-credential-helpers v0.6.3 // indirect
	github.com/docker/go v1.5.1-1 // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/swarmkit v1.12.1-0.20200604173159-967c829cdd33 // indirect
	github.com/fatih/color v1.10.0 // indirect
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/fvbommel/sortorder v1.0.2 // indirect
	github.com/gofrs/flock v0.7.3 // indirect
	github.com/google/go-cmp v0.4.1 // indirect
	github.com/google/shlex v0.0.0-20191202100458-e7afc7fbc510 // indirect
	github.com/googleapis/gnostic v0.2.2 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/hailocab/go-hostpool v0.0.0-20160125115350-e80d13ce29ed // indirect
	github.com/hashicorp/golang-lru v0.5.3 // indirect
	github.com/jinzhu/gorm v1.9.16 // indirect
	github.com/jmoiron/sqlx v1.2.1-0.20190826204134-d7d95172beb5 // indirect
	github.com/kr/pretty v0.2.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/logrusorgru/aurora v2.0.3+incompatible
	github.com/miekg/pkcs11 v1.0.3 // indirect
	github.com/mitchellh/mapstructure v1.3.1 // indirect
	github.com/moby/buildkit v0.7.1-0.20200718032743-4d1f260e8490 // indirect
	github.com/moby/locker v1.0.1 // indirect
	github.com/moby/term v0.0.0-20200915141129-7f0af18e79f2 // indirect
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/onsi/ginkgo v1.12.0 // indirect
	github.com/onsi/gomega v1.9.0 // indirect
	github.com/opencontainers/runc v1.0.0-rc92 //indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/pelletier/go-toml v1.8.0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.7.0
	github.com/smartystreets/assertions v1.0.0 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/cobra v1.0.0
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.7.0 // indirect
	github.com/theupdateframework/notary v0.6.1 // indirect
	github.com/tonistiigi/fsutil v0.0.0-20200724193237-c3ed55f3b481 // indirect
	github.com/tonistiigi/go-rosetta v0.0.0-20201102221648-9ba854641817 // indirect
	gitlab.eng.vmware.com/wujeff/d-cli v0.0.0-20201106060208-4aee41cdf37f // indirect
	go.etcd.io/bbolt v1.3.5 // indirect
	golang.org/x/sync v0.0.0-20200625203802-6e8e738ad208 // indirect
	golang.org/x/time v0.0.0-20200416051211-89c76fbcd5d1 // indirect
	// genproto: the actual version is replaced in replace()
	google.golang.org/genproto v0.0.0-20200527145253-8367513e4ece // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/dancannon/gorethink.v3 v3.0.5 // indirect
	gopkg.in/fatih/pool.v2 v2.0.0 // indirect
	gopkg.in/gorethink/gorethink.v3 v3.0.5 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/ini.v1 v1.56.0 // indirect
	k8s.io/api v0.19.0
	k8s.io/apimachinery v0.19.0
	k8s.io/cli-runtime v0.0.0-00010101000000-000000000000
	k8s.io/client-go v0.19.0
	k8s.io/klog v1.0.0 // indirect
	k8s.io/utils v0.0.0-20190221042446-c2654d5206da
)

replace (
	github.com/containerd/containerd v1.4.0-0 => github.com/containerd/containerd v1.4.1
	github.com/docker/docker => github.com/docker/docker v20.10.0-beta1.0.20201030232932-c2cc352355d4+incompatible
	github.com/fvbommel/sortorder => vbom.ml/util/sortorder v1.0.2
	// protobuf: corresponds to containerd
	github.com/golang/protobuf => github.com/golang/protobuf v1.3.5
	github.com/hashicorp/go-immutable-radix => github.com/tonistiigi/go-immutable-radix v0.0.0-20170803185627-826af9ccf0fe
	github.com/jaguilar/vt100 => github.com/tonistiigi/vt100 v0.0.0-20190402012908-ad4c4a574305
	// genproto: corresponds to containerd
	google.golang.org/genproto => google.golang.org/genproto v0.0.0-20200224152610-e50cd9704f63
	k8s.io/api => k8s.io/api v0.0.0-20190620084959-7cf5895f2711
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20190620085554-14e95df34f1f
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190612205821-1799e75a0719
	k8s.io/apiserver => k8s.io/apiserver v0.0.0-20190620085212-47dc9a115b18
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.0.0-20190620085706-2090e6d8f84c
	k8s.io/client-go => k8s.io/client-go v0.0.0-20190620085101-78d2af792bab
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.0.0-20190620090043-8301c0bda1f0
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.0.0-20190620090013-c9a0fc045dc1
	k8s.io/code-generator => k8s.io/code-generator v0.0.0-20190612205613-18da4a14b22b
	k8s.io/component-base => k8s.io/component-base v0.0.0-20190620085130-185d68e6e6ea
	k8s.io/cri-api => k8s.io/cri-api v0.0.0-20190531030430-6117653b35f1
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.0.0-20190620090116-299a7b270edc
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.0.0-20190620085325-f29e2b4a4f84
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.0.0-20190620085942-b7f18460b210
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.0.0-20190620085809-589f994ddf7f
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.0.0-20190620085912-4acac5405ec6
	k8s.io/kubelet => k8s.io/kubelet v0.0.0-20190620085838-f1cb295a73c9
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.0.0-20190620090156-2138f2c9de18
	k8s.io/metrics => k8s.io/metrics v0.0.0-20190620085625-3b22d835f165
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.0.0-20190620085408-1aef9010884e
	vbom.ml/util/sortorder => github.com/fvbommel/sortorder v1.0.2
)
