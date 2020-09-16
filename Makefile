run-runner:
	go build -buildmode=plugin -o pipes/another-sample/anothersample.so pipes/another-sample/anothersample.go
	go build -buildmode=plugin -o pipes/sample-task/sampletask.so pipes/sample-task/sampletask.go
	go run ./cmd/runner/runner.go
