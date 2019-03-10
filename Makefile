GOCMD=go
GOBUILD=$(GOCMD) build
BINARY_NAME=apk_exporter
all: aarch64 armhf 386 amd64
aarch64:
	rm -r out/aarch64
	GOOS=linux GOARCH=arm64 $(GOBUILD) -o out/aarch64/$(BINARY_NAME)
	cd out/aarch64/ && tar -zcf ../$(BINARY_NAME)_aarch64.tar.gz $(BINARY_NAME)
armhf:
	rm -r out/armhf
	GOOS=linux GOARCH=arm $(GOBUILD) -o out/armhf/$(BINARY_NAME)
	cd out/armhf/ && tar -zcf ../$(BINARY_NAME)_armhf.tar.gz $(BINARY_NAME)
386:
	rm -r out/386
	GOOS=linux GOARCH=386 $(GOBUILD) -o out/386/$(BINARY_NAME)
	cd out/386/ && tar -zcf ../$(BINARY_NAME)_386.tar.gz $(BINARY_NAME)
amd64:
	rm -r out/amd64
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o out/amd64/$(BINARY_NAME)
	cd out/amd64/ && tar -zcf ../$(BINARY_NAME)_amd64.tar.gz $(BINARY_NAME)