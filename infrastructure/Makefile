ENV_NAME:=alpha
AWS_PROFILE=sinfra
APP_VERSION:=latest

create_ecr:
	aws cloudformation create-stack --stack-name worker-management-ecr --template-body file://01_ecr.yaml --profile $(AWS_PROFILE)
create_ecs_cluster:
	aws cloudformation create-stack --stack-name $(ENV_NAME)-worker-management-ecs-cluster --template-body file://02_cluster.yaml --profile $(AWS_PROFILE) --parameters ParameterKey=EnvironmentName,ParameterValue=$(ENV_NAME)
create_ecs_roles:
	aws cloudformation create-stack --stack-name $(ENV_NAME)-worker-management-ecs-iam --template-body file://03_roles.yaml --profile $(AWS_PROFILE) --capabilities CAPABILITY_IAM --parameters ParameterKey=EnvironmentName,ParameterValue=$(ENV_NAME)
create_dynamo_table:
	aws cloudformation create-stack --stack-name $(ENV_NAME)-worker-db --template-body file://04_dynamodb.yaml --profile $(AWS_PROFILE) --parameters ParameterKey=EnvironmentName,ParameterValue=$(ENV_NAME)
create_service:
	aws cloudformation create-stack --stack-name $(ENV_NAME)-worker-management-srv --template-body file://05_service.yaml --profile $(AWS_PROFILE) --parameters ParameterKey=EnvironmentName,ParameterValue=$(ENV_NAME)
create_sqs:
	aws cloudformation create-stack --stack-name $(ENV_NAME)-worker-sqs --template-body file://06_sqs.yaml --profile $(AWS_PROFILE) --parameters ParameterKey=EnvironmentName,ParameterValue=$(ENV_NAME)

update_service:
	aws cloudformation create-stack --stack-name $(ENV_NAME)-worker-db --template-body file://04_dynamodb.yaml --profile $(AWS_PROFILE) --parameters ParameterKey=EnvironmentName,UsePreviousValue=true ParameterKey=AppVersion,ParameterValue=$(APP_VERSION)
