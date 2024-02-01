package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
)
// go mod init
// go mod tidy
//  get github.com/aws/aws-lambda-go/lambda


type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, event *MyEvent) (*string, error) {
	if event == nil {
		return nil, fmt.Errorf("received nil event")
	}
	msg := fmt.Sprintf("Hellos %s!", event.Name)
	return &msg, nil
}

/*
access to metadata about their environment and the invocation request.  at
 Package context. Should your handler include context.Context as a parameter, Lambda will insert
  information about your function into the context's Value property. Note that you need to import
   the lambdacontext library to access the contents of the context.Context object.
*/
func CognitoHandler(ctx context.Context) {
	lc, _ := lambdacontext.FromContext(ctx)
	log.Print(lc.Identity.CognitoIdentityPoolID)
}
func main() {
	lambda.Start(HandleRequest)
	lambda.Start(CognitoHandler)
}