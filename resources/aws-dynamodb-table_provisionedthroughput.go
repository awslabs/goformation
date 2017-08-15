package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::DynamoDB::Table.ProvisionedThroughput AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-provisionedthroughput.html
type AWSDynamoDBTable_ProvisionedThroughput struct {

	// ReadCapacityUnits AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-provisionedthroughput.html#cfn-dynamodb-provisionedthroughput-readcapacityunits
	ReadCapacityUnits int64 `json:"ReadCapacityUnits"`

	// WriteCapacityUnits AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-provisionedthroughput.html#cfn-dynamodb-provisionedthroughput-writecapacityunits
	WriteCapacityUnits int64 `json:"WriteCapacityUnits"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSDynamoDBTable_ProvisionedThroughput) AWSCloudFormationType() string {
	return "AWS::DynamoDB::Table.ProvisionedThroughput"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSDynamoDBTable_ProvisionedThroughput) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSDynamoDBTable_ProvisionedThroughputResources retrieves all AWSDynamoDBTable_ProvisionedThroughput items from a CloudFormation template
func GetAllAWSDynamoDBTable_ProvisionedThroughput(template *Template) map[string]*AWSDynamoDBTable_ProvisionedThroughput {

	results := map[string]*AWSDynamoDBTable_ProvisionedThroughput{}
	for name, resource := range template.Resources {
		result := &AWSDynamoDBTable_ProvisionedThroughput{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSDynamoDBTable_ProvisionedThroughputWithName retrieves all AWSDynamoDBTable_ProvisionedThroughput items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSDynamoDBTable_ProvisionedThroughput(name string, template *Template) (*AWSDynamoDBTable_ProvisionedThroughput, error) {

	result := &AWSDynamoDBTable_ProvisionedThroughput{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSDynamoDBTable_ProvisionedThroughput{}, errors.New("resource not found")

}
