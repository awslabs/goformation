package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Redshift::Cluster.LoggingProperties AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-cluster-loggingproperties.html
type AWSRedshiftCluster_LoggingProperties struct {

	// BucketName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-cluster-loggingproperties.html#cfn-redshift-cluster-loggingproperties-bucketname
	BucketName string `json:"BucketName"`

	// S3KeyPrefix AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-cluster-loggingproperties.html#cfn-redshift-cluster-loggingproperties-s3keyprefix
	S3KeyPrefix string `json:"S3KeyPrefix"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSRedshiftCluster_LoggingProperties) AWSCloudFormationType() string {
	return "AWS::Redshift::Cluster.LoggingProperties"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSRedshiftCluster_LoggingProperties) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSRedshiftCluster_LoggingPropertiesResources retrieves all AWSRedshiftCluster_LoggingProperties items from a CloudFormation template
func GetAllAWSRedshiftCluster_LoggingProperties(template *Template) map[string]*AWSRedshiftCluster_LoggingProperties {

	results := map[string]*AWSRedshiftCluster_LoggingProperties{}
	for name, resource := range template.Resources {
		result := &AWSRedshiftCluster_LoggingProperties{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSRedshiftCluster_LoggingPropertiesWithName retrieves all AWSRedshiftCluster_LoggingProperties items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSRedshiftCluster_LoggingProperties(name string, template *Template) (*AWSRedshiftCluster_LoggingProperties, error) {

	result := &AWSRedshiftCluster_LoggingProperties{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSRedshiftCluster_LoggingProperties{}, errors.New("resource not found")

}
