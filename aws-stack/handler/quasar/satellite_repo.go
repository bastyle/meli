package main

import (
	"errors"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"fmt"
)

type SatelliteEntity struct {
	Name string `json:"name"`
}

type SatelliteDistance struct {
	Distance float64 `json:"distance"`
}

func UpdateSatellite(inputDistance float64, name string) error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-2")},
	)
	svc := dynamodb.New(sess)
	input := &dynamodb.UpdateItemInput{
		TableName: aws.String("Satel"),
		Key: map[string]*dynamodb.AttributeValue{
			"name": {
				S: aws.String("kenobi"),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set distance=:d"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":d": {N: aws.String(strconv.FormatFloat(inputDistance, 'f', -1, 64))},
		},
	}
	_, err = svc.UpdateItem(input)
	if err != nil {
		fmt.Println(err.Error())
		return errors.New("Error updating key")
	}
	return nil
}
