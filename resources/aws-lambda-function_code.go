package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Lambda::Function.Code AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-code.html
type AWSLambdaFunction_Code struct {

	// S3Bucket AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-code.html#cfn-lambda-function-code-s3bucket
	S3Bucket string `json:"S3Bucket"`

	// S3Key AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-code.html#cfn-lambda-function-code-s3key
	S3Key string `json:"S3Key"`

	// S3ObjectVersion AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-code.html#cfn-lambda-function-code-s3objectversion
	S3ObjectVersion string `json:"S3ObjectVersion"`

	// ZipFile AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-code.html#cfn-lambda-function-code-zipfile
	ZipFile string `json:"ZipFile"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSLambdaFunction_Code) AWSCloudFormationType() string {
	return "AWS::Lambda::Function.Code"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSLambdaFunction_Code) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSLambdaFunction_CodeResources retrieves all AWSLambdaFunction_Code items from a CloudFormation template
func GetAllAWSLambdaFunction_Code(template *Template) map[string]*AWSLambdaFunction_Code {

	results := map[string]*AWSLambdaFunction_Code{}
	for name, resource := range template.Resources {
		result := &AWSLambdaFunction_Code{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSLambdaFunction_CodeWithName retrieves all AWSLambdaFunction_Code items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSLambdaFunction_Code(name string, template *Template) (*AWSLambdaFunction_Code, error) {

	result := &AWSLambdaFunction_Code{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSLambdaFunction_Code{}, errors.New("resource not found")

}
