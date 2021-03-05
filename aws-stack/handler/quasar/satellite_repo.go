package main

import (
	"errors"
	"strconv"

	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

const awsRegion = "us-east-2"

/*type SatelliteEntity struct {
	Name string `json:"name"`
}

type SatelliteDistance struct {
	Distance float64 `json:"distance"`
}*/

type DataSat struct {
	Distance float32  `json:"distance"`
	Message  []string `json:"message"`
}

func UpdateSatellite(name string, inputDistance float64, messages []string) error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(awsRegion)},
	)
	svc := dynamodb.New(sess)
	//var msgAtt []*dynamodb.AttributeValue
	msgAtt, err := GetListAttribute(messages)
	if err != nil {
		return err
	}
	input := &dynamodb.UpdateItemInput{
		TableName: aws.String("Satel"),
		Key: map[string]*dynamodb.AttributeValue{
			"name": {
				S: aws.String(name),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set distance=:d, message=:m"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":d": {N: aws.String(strconv.FormatFloat(inputDistance, 'f', -1, 64))},
			":m": {
				L: msgAtt,
			},
		},
	}
	_, err = svc.UpdateItem(input)
	if err != nil {
		fmt.Println(err.Error())
		return errors.New("Error updating key")
	}
	return nil
}

func GetListAttribute(messages []string) ([]*dynamodb.AttributeValue, error) {
	var msgAtt []*dynamodb.AttributeValue
	for _, v := range messages {
		part := &dynamodb.AttributeValue{
			S: aws.String(v),
		}
		msgAtt = append(msgAtt, part)
	}
	if len(msgAtt) == 0 {
		return nil, errors.New("void array")
	}
	return msgAtt, nil
}

func UpdateDistanceSatellite(inputDistance float64, name string) error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(awsRegion)},
	)
	svc := dynamodb.New(sess)
	input := &dynamodb.UpdateItemInput{
		TableName: aws.String("Satel"),
		Key: map[string]*dynamodb.AttributeValue{
			"name": {
				S: aws.String("kenobi"), //TODO cambiar a name
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

//TODO terminar esta función debe retornar array de mensaje y distancia
func GetDataSatell(satName string) (DataSat, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(awsRegion)},
	)
	svc := dynamodb.New(sess)
	satResult, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("Satel"),
		Key: map[string]*dynamodb.AttributeValue{
			"name": {
				S: aws.String(satName),
			},
		},
	})
	item := DataSat{}
	err = dynamodbattribute.UnmarshalMap(satResult.Item, &item)
	fmt.Println("item: ", item)
	//fmt.Println("sat result: ", satResult.Item["distance"])
	if err != nil {
		fmt.Println(err.Error())
		return item, err
	}
	if err = dynamodbattribute.UnmarshalMap(satResult.Item, &item); err != nil {
		return item, err
	} else {
		return item, nil
	}
}

/*func GetList(messages []string) (string, error) {
	var msgAttValArray []*dynamodb.AttributeValue
	for k, v := range messages {
		av := &dynamodb.AttributeValue{
			S: aws.String(v),
		}
		msgAttValArray = append(msgAttValArray, av)
	}
	fmt.Println("msgAttValArray: ", msgAttValArray)
	if err := UpdateMessage(msgAttValArray, "kenobi"); err != nil {
		return "ERROR", errors.New("Error updating message")
	}
	return "", nil
}

func UpdateMessage(msgAttValArray []*dynamodb.AttributeValue, name string) error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(awsRegion)},
	)
	svc := dynamodb.New(sess)
	input := &dynamodb.UpdateItemInput{
		TableName: aws.String("Satel"),
		Key: map[string]*dynamodb.AttributeValue{
			"name": {
				S: aws.String(name),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set message=:m"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":m": {
				L: msgAttValArray,
			},
		},
	}
	_, err = svc.UpdateItem(input)
	if err != nil {
		fmt.Println(err.Error())
		return errors.New("Error updating key")
	}
	return nil
}*/
