package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::KinesisFirehose::DeliveryStream.RedshiftDestinationConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-redshiftdestinationconfiguration.html
type AWSKinesisFirehoseDeliveryStream_RedshiftDestinationConfiguration struct {

	// CloudWatchLoggingOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-redshiftdestinationconfiguration.html#cfn-kinesisfirehose-kinesisdeliverystream-redshiftdestinationconfiguration-cloudwatchloggingoptions

	CloudWatchLoggingOptions AWSKinesisFirehoseDeliveryStream_CloudWatchLoggingOptions `json:"CloudWatchLoggingOptions"`

	// ClusterJDBCURL AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-redshiftdestinationconfiguration.html#cfn-kinesisfirehose-kinesisdeliverystream-redshiftdestinationconfiguration-clusterjdbcurl

	ClusterJDBCURL string `json:"ClusterJDBCURL"`

	// CopyCommand AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-redshiftdestinationconfiguration.html#cfn-kinesisfirehose-kinesisdeliverystream-redshiftdestinationconfiguration-copycommand

	CopyCommand AWSKinesisFirehoseDeliveryStream_CopyCommand `json:"CopyCommand"`

	// Password AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-redshiftdestinationconfiguration.html#cfn-kinesisfirehose-kinesisdeliverystream-redshiftdestinationconfiguration-password

	Password string `json:"Password"`

	// RoleARN AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-redshiftdestinationconfiguration.html#cfn-kinesisfirehose-kinesisdeliverystream-redshiftdestinationconfiguration-rolearn

	RoleARN string `json:"RoleARN"`

	// S3Configuration AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-redshiftdestinationconfiguration.html#cfn-kinesisfirehose-kinesisdeliverystream-redshiftdestinationconfiguration-s3configuration

	S3Configuration AWSKinesisFirehoseDeliveryStream_S3DestinationConfiguration `json:"S3Configuration"`

	// Username AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-redshiftdestinationconfiguration.html#cfn-kinesisfirehose-kinesisdeliverystream-redshiftdestinationconfiguration-usename

	Username string `json:"Username"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisFirehoseDeliveryStream_RedshiftDestinationConfiguration) AWSCloudFormationType() string {
	return "AWS::KinesisFirehose::DeliveryStream.RedshiftDestinationConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKinesisFirehoseDeliveryStream_RedshiftDestinationConfiguration) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSKinesisFirehoseDeliveryStream_RedshiftDestinationConfigurationResources retrieves all AWSKinesisFirehoseDeliveryStream_RedshiftDestinationConfiguration items from a CloudFormation template
func GetAllAWSKinesisFirehoseDeliveryStream_RedshiftDestinationConfiguration(template *Template) map[string]*AWSKinesisFirehoseDeliveryStream_RedshiftDestinationConfiguration {

	results := map[string]*AWSKinesisFirehoseDeliveryStream_RedshiftDestinationConfiguration{}
	for name, resource := range template.Resources {
		result := &AWSKinesisFirehoseDeliveryStream_RedshiftDestinationConfiguration{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSKinesisFirehoseDeliveryStream_RedshiftDestinationConfigurationWithName retrieves all AWSKinesisFirehoseDeliveryStream_RedshiftDestinationConfiguration items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSKinesisFirehoseDeliveryStream_RedshiftDestinationConfiguration(name string, template *Template) (*AWSKinesisFirehoseDeliveryStream_RedshiftDestinationConfiguration, error) {

	result := &AWSKinesisFirehoseDeliveryStream_RedshiftDestinationConfiguration{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSKinesisFirehoseDeliveryStream_RedshiftDestinationConfiguration{}, errors.New("resource not found")

}
