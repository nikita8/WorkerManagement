create-ecr:
	aws cloudformation create-stack --stack-name worker-management-ecr --template-body file://infrastructure/01_ecr.yaml --profile team0workermgmt --region us-east-1
create-ecs-cluster:
	aws cloudformation create-stack --stack-name worker-management-ecs-cluster --template-body file://infrastructure/02_cluster.yaml --profile team0workermgmt --region us-east-1
create_ecs_iam:
	aws cloudformation create-stack --stack-name worker-management-ecs-iam --template-body file://03_roles_and_groups.yaml --profile team0workermgmt --region us-east-1 --capabilities CAPABILITY_IAM
create_service:
	aws cloudformation create-stack --stack-name worker-management-ecs-service --template-body file://04_service.yaml --profile team0workermgmt --region us-east-1
