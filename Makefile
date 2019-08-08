CWD=$$(pwd)

build_image:
	docker build -t go_project_template .

dev:
	docker run --network host --rm -it -v "${CWD}":/go-project-template/ -w /go-project-template go_project_template

build:
	go build -o application

test:
	go test
