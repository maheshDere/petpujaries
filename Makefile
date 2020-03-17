.PHONY: all
all: copy-config test service-run 

copy-config:
	cp application.yaml.sample application.yaml
	cp application.yaml.sample application-test.yaml

service-run: 
	ENVIRONMENT=development go build
	ENVIRONMENT=development ./petpujaris

test-run:
	ENVIRONMENT=test go test -v -p=1 -cover ./...  

clear-test-cache:
	go clean -testcache  
