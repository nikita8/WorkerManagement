build:
	docker build -t worker-management .
run:
	docker run -it -p 3000:3000 worker-management:latest
ecr_push:
	make build
	eval `aws ecr get-login --no-include-email --region us-east-1 --profile sinfra`
	docker tag worker-management:latest 574574226067.dkr.ecr.us-east-1.amazonaws.com/worker-management:$(ver)
	docker push 574574226067.dkr.ecr.us-east-1.amazonaws.com/worker-management:$(ver)

