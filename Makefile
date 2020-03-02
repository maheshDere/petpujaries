.PHONY: all
all: copy-config test service-run 

copy-config:
	cp application.yaml.sample application.yaml
	cp application.yaml.sample application-test.yaml

service-run: 
	ENVIRONMENT=development go run main.go

test:
	ENVIRONMENT=test go run main.go

