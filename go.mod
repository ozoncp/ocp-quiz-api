module github.com/ozoncp/ocp-quiz-api

go 1.16

require (
	github.com/golang/mock v1.6.0
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/onsi/ginkgo v1.16.4
	github.com/onsi/gomega v1.10.1
	golang.org/x/text v0.3.5 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace github.com/ozoncp/ocp-quiz-api/pkg/ocp-quiz-api => ./pkg/ocp-quiz-api
