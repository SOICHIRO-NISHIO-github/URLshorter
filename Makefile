PACKAGE_LIST := $(shell go list ./...)
yubs:
	go build -o yubs $(PACKAGE_LIST)
test:
	go test  $(PACKAGE_LIST)
clean:
	rm -f yubs
