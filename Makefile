build: bin/swagger-codegen-cli.jar clean
	java -jar bin/swagger-codegen-cli.jar generate -i swagger.yaml -l go -c config.json -o databricks

bin/swagger-codegen-cli.jar:
	wget http://central.maven.org/maven2/io/swagger/swagger-codegen-cli/2.3.1/swagger-codegen-cli-2.3.1.jar -O bin/swagger-codegen-cli.jar

build-examples:
	go build -o ./bin/example ./example.go

run-examples: build-examples
	./bin/example

clean:
	rm -rf ./databricks

.PHONY: build build-examples run-examples clean
