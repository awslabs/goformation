package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Config::DeliveryChannel.ConfigSnapshotDeliveryProperties AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-deliverychannel-configsnapshotdeliveryproperties.html
type AWSConfigDeliveryChannel_ConfigSnapshotDeliveryProperties struct {

	// DeliveryFrequency AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-deliverychannel-configsnapshotdeliveryproperties.html#cfn-config-deliverychannel-configsnapshotdeliveryproperties-deliveryfrequency

	DeliveryFrequency string `json:"DeliveryFrequency"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSConfigDeliveryChannel_ConfigSnapshotDeliveryProperties) AWSCloudFormationType() string {
	return "AWS::Config::DeliveryChannel.ConfigSnapshotDeliveryProperties"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSConfigDeliveryChannel_ConfigSnapshotDeliveryProperties) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSConfigDeliveryChannel_ConfigSnapshotDeliveryPropertiesResources retrieves all AWSConfigDeliveryChannel_ConfigSnapshotDeliveryProperties items from a CloudFormation template
func GetAllAWSConfigDeliveryChannel_ConfigSnapshotDeliveryProperties(template *Template) map[string]*AWSConfigDeliveryChannel_ConfigSnapshotDeliveryProperties {

	results := map[string]*AWSConfigDeliveryChannel_ConfigSnapshotDeliveryProperties{}
	for name, resource := range template.Resources {
		result := &AWSConfigDeliveryChannel_ConfigSnapshotDeliveryProperties{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSConfigDeliveryChannel_ConfigSnapshotDeliveryPropertiesWithName retrieves all AWSConfigDeliveryChannel_ConfigSnapshotDeliveryProperties items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSConfigDeliveryChannel_ConfigSnapshotDeliveryProperties(name string, template *Template) (*AWSConfigDeliveryChannel_ConfigSnapshotDeliveryProperties, error) {

	result := &AWSConfigDeliveryChannel_ConfigSnapshotDeliveryProperties{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSConfigDeliveryChannel_ConfigSnapshotDeliveryProperties{}, errors.New("resource not found")

}
