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
	Distance float64  `json:"distance"`
	Message  []string `json:"message"`
	X        float64  `json:"x"`
	Y        float64  `json:"y"`
}

func existSat(satName string) bool {
	if item, err := GetDataSatell(satName); err != nil {
		return false
	} else if item.Name == satName {
		return true
	} else {
		return false
	}
}

func UpdateSatellite(name string, inputDistance float64, messages []string) error {
	if !existSat(name) {
		return errors.New("satellite dosn't exists.")
	}
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(awsRegion)},
	)
	svc := dynamodb.New(sess)
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
		fmt.Printf("err: %v \n", err.Error())
		return errors.New("Error processing the request.")
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
		return nil, errors.New("void message")
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
		return items, errors.New("error getting satellites")
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
		return items, errors.New("error getting satellites")
	}
	//num_items := len(result.Items)
	//fmt.Println("cant items:: ", num_items)
	for _, i := range result.Items {
		item := SatEntity{}
		err = dynamodbattribute.UnmarshalMap(i, &item)
		if err != nil {
			fmt.Println("Got error unmarshalling:")
			fmt.Println(err.Error())
			return items, errors.New("error getting satellites")
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
