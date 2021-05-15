VERSION=0.0.1
all: build install

test:
	TF_LOG=WARN TF_ACC=true go test -count=1 -v ./...

build:
	mkdir -p ./bin
	go build -o ./bin/terraform-provider-bip39_$(VERSION) .

install:
	mkdir -p ~/.terraform.d/plugins
	cp -f ./bin/terraform-provider-bip39_$(VERSION) ~/.terraform.d/plugins/terraform-provider-bip39_$(VERSION)