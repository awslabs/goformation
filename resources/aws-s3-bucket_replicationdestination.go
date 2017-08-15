package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::S3::Bucket.ReplicationDestination AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationconfiguration-rules-destination.html
type AWSS3Bucket_ReplicationDestination struct {

	// Bucket AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationconfiguration-rules-destination.html#cfn-s3-bucket-replicationconfiguration-rules-destination-bucket
	Bucket string `json:"Bucket"`

	// StorageClass AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationconfiguration-rules-destination.html#cfn-s3-bucket-replicationconfiguration-rules-destination-storageclass
	StorageClass string `json:"StorageClass"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSS3Bucket_ReplicationDestination) AWSCloudFormationType() string {
	return "AWS::S3::Bucket.ReplicationDestination"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSS3Bucket_ReplicationDestination) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSS3Bucket_ReplicationDestinationResources retrieves all AWSS3Bucket_ReplicationDestination items from a CloudFormation template
func GetAllAWSS3Bucket_ReplicationDestination(template *Template) map[string]*AWSS3Bucket_ReplicationDestination {

	results := map[string]*AWSS3Bucket_ReplicationDestination{}
	for name, resource := range template.Resources {
		result := &AWSS3Bucket_ReplicationDestination{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSS3Bucket_ReplicationDestinationWithName retrieves all AWSS3Bucket_ReplicationDestination items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSS3Bucket_ReplicationDestination(name string, template *Template) (*AWSS3Bucket_ReplicationDestination, error) {

	result := &AWSS3Bucket_ReplicationDestination{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSS3Bucket_ReplicationDestination{}, errors.New("resource not found")

}
