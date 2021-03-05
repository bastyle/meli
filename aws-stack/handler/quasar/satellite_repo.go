package main

import (
	"errors"
	"strconv"

	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

const awsRegion = "us-east-2"
const tableN = "Satel"

/*type SatelliteEntity struct {
	Name string `json:"name"`
}

type SatelliteDistance struct {
	Distance float64 `json:"distance"`
}*/

type SatEntity struct {
	Name     string   `json:"name"`
	Distance float32  `json:"distance"`
	Message  []string `json:"message"`
	X        float32  `json:"x"`
	Y        float32  `json:"y"`
}

func UpdateSatellite(name string, inputDistance float64, messages []string) error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(awsRegion)},
	)
	svc := dynamodb.New(sess)
	//var msgAtt []*dynamodb.AttributeValue
	msgAtt, err := getListAttribute(messages)
	if err != nil {
		return err
	}
	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(tableN),
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

func getListAttribute(messages []string) ([]*dynamodb.AttributeValue, error) {
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

func GetDataSatell(satName string) (SatEntity, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(awsRegion)},
	)
	svc := dynamodb.New(sess)
	satResult, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableN),
		Key: map[string]*dynamodb.AttributeValue{
			"name": {
				S: aws.String(satName),
			},
		},
	})
	item := SatEntity{}
	err = dynamodbattribute.UnmarshalMap(satResult.Item, &item)
	//fmt.Println("item: ", item)
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

func GetAllDataSatell() ([]SatEntity, error) {
	var items []SatEntity
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(awsRegion)},
	)
	svc := dynamodb.New(sess)
	proj := expression.NamesList(expression.Name("name"), expression.Name("distance"), expression.Name("message"), expression.Name("x"), expression.Name("y"))
	filt := expression.Name("name").NotEqual(expression.Value(""))
	//filt := expression.Name("name").NotEqual(expression.Value("kenobi"))
	expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()
	if err != nil {
		fmt.Println("Got error building expression:")
		fmt.Println(err.Error())
		return items, err
	}
	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(tableN),
	}
	result, err := svc.Scan(params)
	if err != nil {
		fmt.Println("Query API call failed:")
		fmt.Println((err.Error()))
		return items, err
	}
	num_items := len(result.Items)
	fmt.Println("cant items:: ", num_items)
	for _, i := range result.Items {
		item := SatEntity{}
		err = dynamodbattribute.UnmarshalMap(i, &item)
		if err != nil {
			fmt.Println("Got error unmarshalling:")
			fmt.Println(err.Error())
			return items, err
		}
		items = append(items, item)
	}
	return items, nil
}

func ResetSatellDynamicData() error {
	allSatellites, err := GetAllDataSatell()
	if err != nil {
		return err
	}
	blankArray := [1]string{""}
	for _, sat := range allSatellites {
		if err := UpdateSatellite(sat.Name, 0, blankArray[:]); err != nil {
			fmt.Printf("Cannot reset satellite %v.\n", sat.Name)
		}
	}
	return nil
}

/*func getSession() (DynamoDB, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(awsRegion)},
	)
	if err != nil {
		return nil, err
	} else {
		return dynamodb.New(sess), nil
	}
}

func UpdateDistanceSatellite(inputDistance float64, name string) error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(awsRegion)},
	)
	svc := dynamodb.New(sess)
	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(tableN),
		Key: map[string]*dynamodb.AttributeValue{
			"name": {
				S: aws.String(name),
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

func GetList(messages []string) (string, error) {
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
		TableName: aws.String(tableN),
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
