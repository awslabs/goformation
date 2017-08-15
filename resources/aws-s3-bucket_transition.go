package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::S3::Bucket.Transition AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule-transition.html
type AWSS3Bucket_Transition struct {

	// StorageClass AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule-transition.html#cfn-s3-bucket-lifecycleconfig-rule-transition-storageclass
	StorageClass string `json:"StorageClass"`

	// TransitionDate AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule-transition.html#cfn-s3-bucket-lifecycleconfig-rule-transition-transitiondate
	TransitionDate string `json:"TransitionDate"`

	// TransitionInDays AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule-transition.html#cfn-s3-bucket-lifecycleconfig-rule-transition-transitionindays
	TransitionInDays int64 `json:"TransitionInDays"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSS3Bucket_Transition) AWSCloudFormationType() string {
	return "AWS::S3::Bucket.Transition"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSS3Bucket_Transition) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSS3Bucket_TransitionResources retrieves all AWSS3Bucket_Transition items from a CloudFormation template
func GetAllAWSS3Bucket_Transition(template *Template) map[string]*AWSS3Bucket_Transition {

	results := map[string]*AWSS3Bucket_Transition{}
	for name, resource := range template.Resources {
		result := &AWSS3Bucket_Transition{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSS3Bucket_TransitionWithName retrieves all AWSS3Bucket_Transition items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSS3Bucket_Transition(name string, template *Template) (*AWSS3Bucket_Transition, error) {

	result := &AWSS3Bucket_Transition{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSS3Bucket_Transition{}, errors.New("resource not found")

}
