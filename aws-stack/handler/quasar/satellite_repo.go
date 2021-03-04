package main

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"fmt"
)

type SatelliteEntity struct {
	Name string `json:"name"`
}

type SatelliteDistance struct {
	Distance float64 `json:"distance"`
}

func UpdateSatellite(distance float64, name string) error {
	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)

	// Create DynamoDB client

	satDistance := SatelliteDistance{Distance: distance}

	svc := dynamodb.New(sess)
	data, err := dynamodbattribute.MarshalMap(satDistance)
	if err != nil {
		fmt.Println("Got error marshalling satDistance:")
		fmt.Println(err.Error())
		return Errors.new("Error al transformar entidad")
	}

	satToSearch := SatelliteEntity{Name: name}
	key, err := dynamodbattribute.MarshalMap(satToSearch)
	if err != nil {
		fmt.Println("Got error marshalling entity:")
		fmt.Println(err.Error())
		return errors.New("Error transformatting key")
	}

	// Update item in table Movies
	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: data,
		TableName:                 aws.String("Satel"),
		Key:                       key,
		ReturnValues:              aws.String("UPDATED_NEW"),
		UpdateExpression:          aws.String("set satDistance.distance = :r"),
	}

	_, err = svc.UpdateItem(input)
	if err != nil {
		fmt.Println(err.Error())
		return errors.New("Error updating key")
	}

}
