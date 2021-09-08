module github.com/seldonio/seldon-operator

go 1.15

require (
	github.com/beorn7/perks v0.0.0-20180321164747-3a771d992973 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/emicklei/go-restful v2.9.3+incompatible // indirect
	github.com/fsnotify/fsnotify v1.4.7 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/go-logr/logr v0.1.0 // indirect
	github.com/go-logr/zapr v0.1.1 // indirect
	github.com/gobuffalo/envy v1.6.15 // indirect
	github.com/gogo/protobuf v1.2.1 // indirect
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b // indirect
	github.com/google/btree v1.0.0 // indirect
	github.com/google/go-cmp v0.5.1
	github.com/google/gofuzz v0.0.0-20170612174753-24818f796faf // indirect
	github.com/google/uuid v1.1.1 // indirect
	github.com/googleapis/gnostic v0.2.0 // indirect
	github.com/gregjones/httpcache v0.0.0-20190212212710-3befbb6ad0cc // indirect
	github.com/hashicorp/golang-lru v0.5.1 // indirect
	github.com/hpcloud/tail v1.0.0 // indirect
	github.com/imdario/mergo v0.3.7 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/joho/godotenv v1.3.0 // indirect
	github.com/json-iterator/go v1.1.6 // indirect
	github.com/knative/pkg v0.0.0-20190419031227-3c8c4a93547f
	github.com/markbates/inflect v1.0.4 // indirect
	github.com/mattbaird/jsonpatch v0.0.0-20171005235357-81af80346b1a // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v0.0.0-20180701023420-4b7aa43c6742 // indirect
	github.com/onsi/ginkgo v1.8.0 // indirect
	github.com/onsi/gomega v1.5.0
	github.com/pborman/uuid v0.0.0-20180906182336-adf5a7427709 // indirect
	github.com/petar/GoLLRB v0.0.0-20130427215148-53be0d36a84c // indirect
	github.com/peterbourgon/diskv v2.0.1+incompatible // indirect
	github.com/pkg/errors v0.8.1 // indirect
	github.com/prometheus/client_golang v0.9.2 // indirect
	github.com/prometheus/common v0.2.0 // indirect
	github.com/prometheus/procfs v0.0.0-20190328153300-af7bedc223fb // indirect
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cobra v0.0.3 // indirect
	github.com/spf13/pflag v1.0.3 // indirect
	go.uber.org/atomic v1.3.2 // indirect
	go.uber.org/multierr v1.1.0 // indirect
	go.uber.org/zap v1.9.1 // indirect
	golang.org/x/net v0.0.0-20200822124328-c89045814202
	golang.org/x/oauth2 v0.0.0-20210819190943-2bc19b11175f // indirect
	golang.org/x/time v0.0.0-20210723032227-1f47c861a9ac // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
	gopkg.in/yaml.v2 v2.2.2
	k8s.io/api v0.0.0-20190226173710-145d52631d00
	k8s.io/apiextensions-apiserver v0.0.0-20181126155829-0cd23ebeb688 // indirect
	k8s.io/apimachinery v0.0.0-20190221084156-01f179d85dbc
	k8s.io/client-go v0.0.0-20190226174127-78295b709ec6
	k8s.io/code-generator v0.0.0-20181128191024-b1289fc74931 // indirect
	k8s.io/gengo v0.0.0-20190327210449-e17681d19d3a // indirect
	k8s.io/klog v0.2.0 // indirect
	k8s.io/kube-openapi v0.0.0-20190320154901-5e45bb682580 // indirect
	sigs.k8s.io/controller-runtime v0.1.9
	sigs.k8s.io/controller-tools v0.1.7 // indirect
	sigs.k8s.io/testing_frameworks v0.1.1 // indirect
)

replace gopkg.in/fsnotify.v1 v1.4.7 => github.com/fsnotify/fsnotify v1.4.7

replace github.com/fsnotify/fsnotify v1.4.7 => gopkg.in/fsnotify.v1 v1.4.7
