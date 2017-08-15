package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Config::DeliveryChannel AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-deliverychannel.html
type AWSConfigDeliveryChannel struct {

	// ConfigSnapshotDeliveryProperties AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-deliverychannel.html#cfn-config-deliverychannel-configsnapshotdeliveryproperties
	ConfigSnapshotDeliveryProperties AWSConfigDeliveryChannel_ConfigSnapshotDeliveryProperties `json:"ConfigSnapshotDeliveryProperties"`

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-deliverychannel.html#cfn-config-deliverychannel-name
	Name string `json:"Name"`

	// S3BucketName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-deliverychannel.html#cfn-config-deliverychannel-s3bucketname
	S3BucketName string `json:"S3BucketName"`

	// S3KeyPrefix AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-deliverychannel.html#cfn-config-deliverychannel-s3keyprefix
	S3KeyPrefix string `json:"S3KeyPrefix"`

	// SnsTopicARN AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-deliverychannel.html#cfn-config-deliverychannel-snstopicarn
	SnsTopicARN string `json:"SnsTopicARN"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSConfigDeliveryChannel) AWSCloudFormationType() string {
	return "AWS::Config::DeliveryChannel"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSConfigDeliveryChannel) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSConfigDeliveryChannelResources retrieves all AWSConfigDeliveryChannel items from a CloudFormation template
func GetAllAWSConfigDeliveryChannel(template *Template) map[string]*AWSConfigDeliveryChannel {

	results := map[string]*AWSConfigDeliveryChannel{}
	for name, resource := range template.Resources {
		result := &AWSConfigDeliveryChannel{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSConfigDeliveryChannelWithName retrieves all AWSConfigDeliveryChannel items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSConfigDeliveryChannel(name string, template *Template) (*AWSConfigDeliveryChannel, error) {

	result := &AWSConfigDeliveryChannel{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSConfigDeliveryChannel{}, errors.New("resource not found")

}
