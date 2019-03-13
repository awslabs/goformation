package resources

import (
	"sort"

	"encoding/json"
)

// AWSServerlessFunction_Properties is a helper struct that can hold either a S3Event, SNSEvent, SQSEvent, KinesisEvent, DynamoDBEvent, ApiEvent, ScheduleEvent, CloudWatchEventEvent, IoTRuleEvent, or AlexaSkillEvent value
type AWSServerlessFunction_Properties struct {
	S3Event              *AWSServerlessFunction_S3Event
	SNSEvent             *AWSServerlessFunction_SNSEvent
	SQSEvent             *AWSServerlessFunction_SQSEvent
	KinesisEvent         *AWSServerlessFunction_KinesisEvent
	DynamoDBEvent        *AWSServerlessFunction_DynamoDBEvent
	ApiEvent             *AWSServerlessFunction_ApiEvent
	ScheduleEvent        *AWSServerlessFunction_ScheduleEvent
	CloudWatchEventEvent *AWSServerlessFunction_CloudWatchEventEvent
	IoTRuleEvent         *AWSServerlessFunction_IoTRuleEvent
	AlexaSkillEvent      *AWSServerlessFunction_AlexaSkillEvent
}

func (r AWSServerlessFunction_Properties) value() interface{} {

	ret := []interface{}{}

	if r.S3Event != nil {
		ret = append(ret, *r.S3Event)
	}

	if r.SNSEvent != nil {
		ret = append(ret, *r.SNSEvent)
	}

	if r.SQSEvent != nil {
		ret = append(ret, *r.SQSEvent)
	}

	if r.KinesisEvent != nil {
		ret = append(ret, *r.KinesisEvent)
	}

	if r.DynamoDBEvent != nil {
		ret = append(ret, *r.DynamoDBEvent)
	}

	if r.ApiEvent != nil {
		ret = append(ret, *r.ApiEvent)
	}

	if r.ScheduleEvent != nil {
		ret = append(ret, *r.ScheduleEvent)
	}

	if r.CloudWatchEventEvent != nil {
		ret = append(ret, *r.CloudWatchEventEvent)
	}

	if r.IoTRuleEvent != nil {
		ret = append(ret, *r.IoTRuleEvent)
	}

	if r.AlexaSkillEvent != nil {
		ret = append(ret, *r.AlexaSkillEvent)
	}

	sort.Sort(byJSONLength(ret))
	if len(ret) > 0 {
		return ret[0]
	}

	return nil

}

func (r AWSServerlessFunction_Properties) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.value())
}

// Hook into the marshaller
func (r *AWSServerlessFunction_Properties) UnmarshalJSON(b []byte) error {

	// Unmarshal into interface{} to check it's type
	var typecheck interface{}
	if err := json.Unmarshal(b, &typecheck); err != nil {
		return err
	}

	switch val := typecheck.(type) {

	case map[string]interface{}:
		val = val // This ensures val is used to stop an error

		json.Unmarshal(b, &r.S3Event)

		json.Unmarshal(b, &r.SNSEvent)

		json.Unmarshal(b, &r.SQSEvent)

		json.Unmarshal(b, &r.KinesisEvent)

		json.Unmarshal(b, &r.DynamoDBEvent)

		json.Unmarshal(b, &r.ApiEvent)

		json.Unmarshal(b, &r.ScheduleEvent)

		json.Unmarshal(b, &r.CloudWatchEventEvent)

		json.Unmarshal(b, &r.IoTRuleEvent)

		json.Unmarshal(b, &r.AlexaSkillEvent)

	case []interface{}:

	}

	return nil
}
