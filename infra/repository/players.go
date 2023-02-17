package repository

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func ListTables() {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("eu-west-1"),
		Credentials: credentials.NewSharedCredentials("", "extrembcn@gmail.com"),
	})

	if err != nil {
		fmt.Printf("Error creating aws session. Cause: %v \n", err)
	}

	svc := dynamodb.New(sess)

	id := "1"
	tableName := "players"
	
	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: &tableName,
		Key:       map[string]*dynamodb.AttributeValue{"id": {S: &id}},
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(result.Item)

	

}
