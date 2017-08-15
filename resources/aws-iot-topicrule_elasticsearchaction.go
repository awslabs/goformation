package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::IoT::TopicRule.ElasticsearchAction AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-elasticsearch.html
type AWSIoTTopicRule_ElasticsearchAction struct {

	// Endpoint AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-elasticsearch.html#cfn-iot-elasticsearch-endpoint
	Endpoint string `json:"Endpoint"`

	// Id AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-elasticsearch.html#cfn-iot-elasticsearch-id
	Id string `json:"Id"`

	// Index AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-elasticsearch.html#cfn-iot-elasticsearch-index
	Index string `json:"Index"`

	// RoleArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-elasticsearch.html#cfn-iot-elasticsearch-rolearn
	RoleArn string `json:"RoleArn"`

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-elasticsearch.html#cfn-iot-elasticsearch-type
	Type string `json:"Type"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIoTTopicRule_ElasticsearchAction) AWSCloudFormationType() string {
	return "AWS::IoT::TopicRule.ElasticsearchAction"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSIoTTopicRule_ElasticsearchAction) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSIoTTopicRule_ElasticsearchActionResources retrieves all AWSIoTTopicRule_ElasticsearchAction items from a CloudFormation template
func GetAllAWSIoTTopicRule_ElasticsearchAction(template *Template) map[string]*AWSIoTTopicRule_ElasticsearchAction {

	results := map[string]*AWSIoTTopicRule_ElasticsearchAction{}
	for name, resource := range template.Resources {
		result := &AWSIoTTopicRule_ElasticsearchAction{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSIoTTopicRule_ElasticsearchActionWithName retrieves all AWSIoTTopicRule_ElasticsearchAction items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSIoTTopicRule_ElasticsearchAction(name string, template *Template) (*AWSIoTTopicRule_ElasticsearchAction, error) {

	result := &AWSIoTTopicRule_ElasticsearchAction{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSIoTTopicRule_ElasticsearchAction{}, errors.New("resource not found")

}
