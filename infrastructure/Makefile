create_ecr:
	aws cloudformation create-stack --stack-name worker-management-ecr --template-body file://01_ecr.yaml --profile sinfra --region us-east-1
create_ecs_cluster:
	aws cloudformation create-stack --stack-name worker-management-ecs-cluster --template-body file://02_cluster.yaml --profile sinfra --region us-east-1
create_ecs_roles:
	aws cloudformation create-stack --stack-name worker-management-ecs-iam --template-body file://03_roles.yaml --profile sinfra --region us-east-1 --capabilities CAPABILITY_IAM
create_service:
	aws cloudformation create-stack --stack-name worker-management-information-sys --template-body file://05_service.yaml --profile sinfra --region us-east-1
update_service:
	aws cloudformation update-stack --stack-name worker-management-information-sys --template-body file://05_service.yaml --profile sinfra --region us-east-1
create_dynamo_table:
	aws cloudformation create-stack --stack-name worker-db --template-body file://04_dynamodb.yaml --profile sinfra --region us-east-1
update_dynamo_table:
	aws cloudformation update-stack --stack-name worker-db --template-body file://04_dynamodb.yaml --profile sinfra --region us-east-1
