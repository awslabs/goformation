package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::CloudFront::Distribution.Logging AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-logging.html
type AWSCloudFrontDistribution_Logging struct {

	// Bucket AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-logging.html#cfn-cloudfront-logging-bucket
	Bucket string `json:"Bucket"`

	// IncludeCookies AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-logging.html#cfn-cloudfront-logging-includecookies
	IncludeCookies bool `json:"IncludeCookies"`

	// Prefix AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-logging.html#cfn-cloudfront-logging-prefix
	Prefix string `json:"Prefix"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCloudFrontDistribution_Logging) AWSCloudFormationType() string {
	return "AWS::CloudFront::Distribution.Logging"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCloudFrontDistribution_Logging) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCloudFrontDistribution_LoggingResources retrieves all AWSCloudFrontDistribution_Logging items from a CloudFormation template
func GetAllAWSCloudFrontDistribution_Logging(template *Template) map[string]*AWSCloudFrontDistribution_Logging {

	results := map[string]*AWSCloudFrontDistribution_Logging{}
	for name, resource := range template.Resources {
		result := &AWSCloudFrontDistribution_Logging{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCloudFrontDistribution_LoggingWithName retrieves all AWSCloudFrontDistribution_Logging items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSCloudFrontDistribution_Logging(name string, template *Template) (*AWSCloudFrontDistribution_Logging, error) {

	result := &AWSCloudFrontDistribution_Logging{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCloudFrontDistribution_Logging{}, errors.New("resource not found")

}
