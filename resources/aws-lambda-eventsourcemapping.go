package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSLambdaEventSourceMapping AWS CloudFormation Resource (AWS::Lambda::EventSourceMapping)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html
type AWSLambdaEventSourceMapping struct {

	// BatchSize AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-batchsize
	BatchSize int `json:"BatchSize"`

	// Enabled AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-enabled
	Enabled bool `json:"Enabled"`

	// EventSourceArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-eventsourcearn
	EventSourceArn string `json:"EventSourceArn"`

	// FunctionName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-functionname
	FunctionName string `json:"FunctionName"`

	// StartingPosition AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-startingposition
	StartingPosition string `json:"StartingPosition"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSLambdaEventSourceMapping) AWSCloudFormationType() string {
	return "AWS::Lambda::EventSourceMapping"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSLambdaEventSourceMapping) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSLambdaEventSourceMappingResources retrieves all AWSLambdaEventSourceMapping items from a CloudFormation template
func (t *Template) GetAllAWSLambdaEventSourceMappingResources() map[string]*AWSLambdaEventSourceMapping {

	results := map[string]*AWSLambdaEventSourceMapping{}
	for name, resource := range t.Resources {
		result := &AWSLambdaEventSourceMapping{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSLambdaEventSourceMappingWithName retrieves all AWSLambdaEventSourceMapping items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSLambdaEventSourceMappingWithName(name string) (*AWSLambdaEventSourceMapping, error) {

	result := &AWSLambdaEventSourceMapping{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSLambdaEventSourceMapping{}, errors.New("resource not found")

}
