package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::S3::Bucket.RoutingRule AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-routingrules.html
type AWSS3Bucket_RoutingRule struct {

	// RedirectRule AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-routingrules.html#cfn-s3-websiteconfiguration-routingrules-redirectrule
	RedirectRule AWSS3Bucket_RedirectRule `json:"RedirectRule"`

	// RoutingRuleCondition AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-routingrules.html#cfn-s3-websiteconfiguration-routingrules-routingrulecondition
	RoutingRuleCondition AWSS3Bucket_RoutingRuleCondition `json:"RoutingRuleCondition"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSS3Bucket_RoutingRule) AWSCloudFormationType() string {
	return "AWS::S3::Bucket.RoutingRule"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSS3Bucket_RoutingRule) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSS3Bucket_RoutingRuleResources retrieves all AWSS3Bucket_RoutingRule items from a CloudFormation template
func GetAllAWSS3Bucket_RoutingRule(template *Template) map[string]*AWSS3Bucket_RoutingRule {

	results := map[string]*AWSS3Bucket_RoutingRule{}
	for name, resource := range template.Resources {
		result := &AWSS3Bucket_RoutingRule{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSS3Bucket_RoutingRuleWithName retrieves all AWSS3Bucket_RoutingRule items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSS3Bucket_RoutingRule(name string, template *Template) (*AWSS3Bucket_RoutingRule, error) {

	result := &AWSS3Bucket_RoutingRule{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSS3Bucket_RoutingRule{}, errors.New("resource not found")

}
