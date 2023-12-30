Building a lambda function using golang, interacting with aws SDK for golang.
Steps:
 Create a directory for this project.
 import the lambda package by running:
        
        go get -u github.com/aws/aws-lambda-go
        go get "github.com/aws/aws-lambda-go/lambda"

Initialize a module using:

        go mod init "github.com/marviigrey/hello-world-lambda"
