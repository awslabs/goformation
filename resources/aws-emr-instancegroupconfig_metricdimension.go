package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EMR::InstanceGroupConfig.MetricDimension AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-metricdimension.html
type AWSEMRInstanceGroupConfig_MetricDimension struct {

	// Key AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-metricdimension.html#cfn-elasticmapreduce-instancegroupconfig-metricdimension-key
	Key string `json:"Key"`

	// Value AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-metricdimension.html#cfn-elasticmapreduce-instancegroupconfig-metricdimension-value
	Value string `json:"Value"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRInstanceGroupConfig_MetricDimension) AWSCloudFormationType() string {
	return "AWS::EMR::InstanceGroupConfig.MetricDimension"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRInstanceGroupConfig_MetricDimension) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEMRInstanceGroupConfig_MetricDimensionResources retrieves all AWSEMRInstanceGroupConfig_MetricDimension items from a CloudFormation template
func GetAllAWSEMRInstanceGroupConfig_MetricDimension(template *Template) map[string]*AWSEMRInstanceGroupConfig_MetricDimension {

	results := map[string]*AWSEMRInstanceGroupConfig_MetricDimension{}
	for name, resource := range template.Resources {
		result := &AWSEMRInstanceGroupConfig_MetricDimension{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEMRInstanceGroupConfig_MetricDimensionWithName retrieves all AWSEMRInstanceGroupConfig_MetricDimension items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEMRInstanceGroupConfig_MetricDimension(name string, template *Template) (*AWSEMRInstanceGroupConfig_MetricDimension, error) {

	result := &AWSEMRInstanceGroupConfig_MetricDimension{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEMRInstanceGroupConfig_MetricDimension{}, errors.New("resource not found")

}
