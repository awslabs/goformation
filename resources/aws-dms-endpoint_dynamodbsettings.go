package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::DMS::Endpoint.DynamoDbSettings AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-dynamodbsettings.html
type AWSDMSEndpoint_DynamoDbSettings struct {

	// ServiceAccessRoleArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-dynamodbsettings.html#cfn-dms-endpoint-dynamodbsettings-serviceaccessrolearn

	ServiceAccessRoleArn string `json:"ServiceAccessRoleArn"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSDMSEndpoint_DynamoDbSettings) AWSCloudFormationType() string {
	return "AWS::DMS::Endpoint.DynamoDbSettings"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSDMSEndpoint_DynamoDbSettings) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSDMSEndpoint_DynamoDbSettingsResources retrieves all AWSDMSEndpoint_DynamoDbSettings items from a CloudFormation template
func GetAllAWSDMSEndpoint_DynamoDbSettings(template *Template) map[string]*AWSDMSEndpoint_DynamoDbSettings {

	results := map[string]*AWSDMSEndpoint_DynamoDbSettings{}
	for name, resource := range template.Resources {
		result := &AWSDMSEndpoint_DynamoDbSettings{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSDMSEndpoint_DynamoDbSettingsWithName retrieves all AWSDMSEndpoint_DynamoDbSettings items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSDMSEndpoint_DynamoDbSettings(name string, template *Template) (*AWSDMSEndpoint_DynamoDbSettings, error) {

	result := &AWSDMSEndpoint_DynamoDbSettings{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSDMSEndpoint_DynamoDbSettings{}, errors.New("resource not found")

}
