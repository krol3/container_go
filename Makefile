run:
	go run main.go

build:
	docker build -t krol/app-exec .

container-run:
	docker run -p 5050:5050 krol/app-exec

publish:
	docker tag krol/app-exec ghcr.io/krol3/app-exec:latest
	# echo $CR_PAT | docker login ghcr.io -u krol3 --password-stdin
	docker push ghcr.io/krol3/app-exec:latest
	docker push krol/app-exec:latest