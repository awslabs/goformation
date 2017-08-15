package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::S3::Bucket.CorsRule AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-cors-corsrule.html
type AWSS3Bucket_CorsRule struct {

	// AllowedHeaders AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-cors-corsrule.html#cfn-s3-bucket-cors-corsrule-allowedheaders
	AllowedHeaders []string `json:"AllowedHeaders"`

	// AllowedMethods AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-cors-corsrule.html#cfn-s3-bucket-cors-corsrule-allowedmethods
	AllowedMethods []string `json:"AllowedMethods"`

	// AllowedOrigins AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-cors-corsrule.html#cfn-s3-bucket-cors-corsrule-allowedorigins
	AllowedOrigins []string `json:"AllowedOrigins"`

	// ExposedHeaders AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-cors-corsrule.html#cfn-s3-bucket-cors-corsrule-exposedheaders
	ExposedHeaders []string `json:"ExposedHeaders"`

	// Id AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-cors-corsrule.html#cfn-s3-bucket-cors-corsrule-id
	Id string `json:"Id"`

	// MaxAge AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-cors-corsrule.html#cfn-s3-bucket-cors-corsrule-maxage
	MaxAge int64 `json:"MaxAge"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSS3Bucket_CorsRule) AWSCloudFormationType() string {
	return "AWS::S3::Bucket.CorsRule"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSS3Bucket_CorsRule) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSS3Bucket_CorsRuleResources retrieves all AWSS3Bucket_CorsRule items from a CloudFormation template
func GetAllAWSS3Bucket_CorsRule(template *Template) map[string]*AWSS3Bucket_CorsRule {

	results := map[string]*AWSS3Bucket_CorsRule{}
	for name, resource := range template.Resources {
		result := &AWSS3Bucket_CorsRule{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSS3Bucket_CorsRuleWithName retrieves all AWSS3Bucket_CorsRule items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSS3Bucket_CorsRule(name string, template *Template) (*AWSS3Bucket_CorsRule, error) {

	result := &AWSS3Bucket_CorsRule{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSS3Bucket_CorsRule{}, errors.New("resource not found")

}
