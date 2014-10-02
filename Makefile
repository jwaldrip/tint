test: install-deps
	@clear
	go test

install-deps:
	go get github.com/onsi/ginkgo
	go get github.com/onsi/gomega

install: install-deps
	go install
