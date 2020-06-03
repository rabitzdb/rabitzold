package ingestion

import (
	"fmt"
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	//"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, kinesisEvent events.KinesisEvent) error {
	for _, record := range kinesisEvent.Records {
		kinesisRecord := record.Kinesis
		dataBytes := kinesisRecord.Data
		dataText := string(dataBytes)
		fmt.Printf("%s Data = %s \n", record.EventName, dataText)
		var result map[string]interface{}
		//Parse Json
		json.Unmarshal(dataBytes, &result)
		//Parámetros:
		// “timestamp” : Timestamp from data
		// “precision” : ns, u, ms, s, m, h. default ns
		// “data” : [{ <name> : <value>}]
		//TODO: Validate Structure

		//TODO: Save data
	}
	return nil
}

/*
Validate Json Structure
 */
func validateStructure(){
	//return nil
}

/*
Write data to dataspace
 */
func write(){
	//return nil
}
