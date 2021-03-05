include .env

clean:
		@rm -rf dist
		@mkdir -p dist

build: clean
		@for dir in `ls handler`; do \
			GOOS=linux go build -o dist/handler/$$dir ./handler/$$dir; \
		done

run:
		aws-sam-local local start-api

install:
		go get github.com/aws/aws-lambda-go/events
		go get github.com/aws/aws-lambda-go/lambda
		go get github.com/aws/aws-sdk-go/aws
		go get github.com/aws/aws-sdk-go/aws/session
		go get github.com/aws/aws-sdk-go/service/dynamodb
		go get github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute

install-dev:
		go get github.com/awslabs/aws-sam-local

test:
		go test -v ./... --cover

configure:
		aws s3api create-bucket \
			--bucket $(AWS_BUCKET_NAME) \
			--region $(AWS_REGION) \
			--create-bucket-configuration LocationConstraint=$(AWS_REGION)

put-start-items:
		aws dynamodb put-item \
			--table-name Satel \
			--item '{
				"name": {"S": "kenobi"},
				"distance": {"N": "0"},
				"x": {"N": "-500"},
				"y": {"N": "-200"}
			}' \
			--return-consumed-capacity TOTAL \
			--region $(AWS_REGION)
		
		aws dynamodb put-item \
			--table-name Satel \
			--item '{
				"name": {"S": "skywalker"},
				"distance": {"N": "0"},
				"x": {"N": "100"},
				"y": {"N": "-100"}
			}' \
			--return-consumed-capacity TOTAL \
			--region $(AWS_REGION)

		aws dynamodb put-item \
			--table-name Satel \
			--item '{
				"name": {"S": "sato"},
				"distance": {"N": "0"},
				"x": {"N": "500"},
				"y": {"N": "100"}
			}' \
			--return-consumed-capacity TOTAL \
			--region $(AWS_REGION)

package: build
		@aws cloudformation package \
			--template-file template.yml \
			--s3-bucket $(AWS_BUCKET_NAME) \
			--region $(AWS_REGION) \
			--output-template-file package.yml

deploy:
		@aws cloudformation deploy \
			--template-file package.yml \
			--region $(AWS_REGION) \
			--capabilities CAPABILITY_IAM \
			--stack-name $(AWS_STACK_NAME)

describe:
		@aws cloudformation describe-stacks \
			--region $(AWS_REGION) \
			--stack-name $(AWS_STACK_NAME) \

outputs:
		@make describe | jq -r '.Stacks[0].Outputs'

url:
		@make describe | jq -r ".Stacks[0].Outputs[0].OutputValue" -j
