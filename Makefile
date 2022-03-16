PROJECT="deamo_exporter"
MAIN_PATH="deamo_exporter.go"

all: build

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${PROJECT} ${MAIN_PATH}

clean:
	@if [ -f ${PROJECT} ] ; then rm ${PROJECT} ; fi