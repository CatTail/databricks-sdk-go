swagger-codegen-cli.jar:
	wget http://central.maven.org/maven2/io/swagger/swagger-codegen-cli/2.3.1/swagger-codegen-cli-2.3.1.jar -O swagger-codegen-cli.jar

build-examples:
	go build -o ./examples/cluster/main ./examples/cluster

build: swagger-codegen-cli.jar
	java -jar swagger-codegen-cli.jar generate -i swagger.yaml -l go -c config.json -o client

clean:
	rm -rf ./client

.PHONY: build clean
