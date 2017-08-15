package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::S3::Bucket.RoutingRuleCondition AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-routingrules-routingrulecondition.html
type AWSS3Bucket_RoutingRuleCondition struct {

	// HttpErrorCodeReturnedEquals AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-routingrules-routingrulecondition.html#cfn-s3-websiteconfiguration-routingrules-routingrulecondition-httperrorcodereturnedequals
	HttpErrorCodeReturnedEquals string `json:"HttpErrorCodeReturnedEquals"`

	// KeyPrefixEquals AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-routingrules-routingrulecondition.html#cfn-s3-websiteconfiguration-routingrules-routingrulecondition-keyprefixequals
	KeyPrefixEquals string `json:"KeyPrefixEquals"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSS3Bucket_RoutingRuleCondition) AWSCloudFormationType() string {
	return "AWS::S3::Bucket.RoutingRuleCondition"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSS3Bucket_RoutingRuleCondition) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSS3Bucket_RoutingRuleConditionResources retrieves all AWSS3Bucket_RoutingRuleCondition items from a CloudFormation template
func GetAllAWSS3Bucket_RoutingRuleCondition(template *Template) map[string]*AWSS3Bucket_RoutingRuleCondition {

	results := map[string]*AWSS3Bucket_RoutingRuleCondition{}
	for name, resource := range template.Resources {
		result := &AWSS3Bucket_RoutingRuleCondition{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSS3Bucket_RoutingRuleConditionWithName retrieves all AWSS3Bucket_RoutingRuleCondition items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSS3Bucket_RoutingRuleCondition(name string, template *Template) (*AWSS3Bucket_RoutingRuleCondition, error) {

	result := &AWSS3Bucket_RoutingRuleCondition{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSS3Bucket_RoutingRuleCondition{}, errors.New("resource not found")

}
