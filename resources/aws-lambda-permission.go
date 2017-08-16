package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSLambdaPermission AWS CloudFormation Resource (AWS::Lambda::Permission)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-permission.html
type AWSLambdaPermission struct {

	// Action AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-permission.html#cfn-lambda-permission-action
	Action string `json:"Action"`

	// FunctionName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-permission.html#cfn-lambda-permission-functionname
	FunctionName string `json:"FunctionName"`

	// Principal AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-permission.html#cfn-lambda-permission-principal
	Principal string `json:"Principal"`

	// SourceAccount AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-permission.html#cfn-lambda-permission-sourceaccount
	SourceAccount string `json:"SourceAccount"`

	// SourceArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-permission.html#cfn-lambda-permission-sourcearn
	SourceArn string `json:"SourceArn"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSLambdaPermission) AWSCloudFormationType() string {
	return "AWS::Lambda::Permission"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSLambdaPermission) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSLambdaPermissionResources retrieves all AWSLambdaPermission items from a CloudFormation template
func (t *Template) GetAllAWSLambdaPermissionResources() map[string]*AWSLambdaPermission {

	results := map[string]*AWSLambdaPermission{}
	for name, resource := range t.Resources {
		result := &AWSLambdaPermission{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSLambdaPermissionWithName retrieves all AWSLambdaPermission items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSLambdaPermissionWithName(name string) (*AWSLambdaPermission, error) {

	result := &AWSLambdaPermission{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSLambdaPermission{}, errors.New("resource not found")

}
