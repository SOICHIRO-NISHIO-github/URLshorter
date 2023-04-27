PACKAGE_LIST := $(shell go list ./...)
VERSION := 0.1.2
NAME :=yubs
DIST := $(NAME)-$(VERSION)

yubs: coverage.out
	go build -o yubs $(PACKAGE_LIST)
	
coverage.out:
	go test -covermode=count -coverprofile=coverage.out $(PACKAGE_LIST)
	
test:
	go test $(PACKAGE_LIST)
	
clean:
	rm -f yubs
