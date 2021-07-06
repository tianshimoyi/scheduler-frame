
.PHONY: all

Binary:="daiwa"

all: clean swag build run

Docker: swag buildDocker runDocker

swag:
	@swag init

build:
	@go build -o ${Binary} .

run:
	@./${Binary}

clean:
	@ rm -rf ${Binary}

docker: buildDocker pushDocker

buildDocker:
	@docker build -t tianshimoyi/my-scheduler:v1 .

pushDocker:
	@docker push tianshimoyi/my-scheduler:v1