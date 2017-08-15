package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EMR::InstanceFleetConfig.InstanceTypeConfig AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancefleetconfig-instancetypeconfig.html
type AWSEMRInstanceFleetConfig_InstanceTypeConfig struct {

	// BidPrice AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancefleetconfig-instancetypeconfig.html#cfn-elasticmapreduce-instancefleetconfig-instancetypeconfig-bidprice
	BidPrice string `json:"BidPrice"`

	// BidPriceAsPercentageOfOnDemandPrice AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancefleetconfig-instancetypeconfig.html#cfn-elasticmapreduce-instancefleetconfig-instancetypeconfig-bidpriceaspercentageofondemandprice
	BidPriceAsPercentageOfOnDemandPrice float64 `json:"BidPriceAsPercentageOfOnDemandPrice"`

	// Configurations AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancefleetconfig-instancetypeconfig.html#cfn-elasticmapreduce-instancefleetconfig-instancetypeconfig-configurations
	Configurations []AWSEMRInstanceFleetConfig_Configuration `json:"Configurations"`

	// EbsConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancefleetconfig-instancetypeconfig.html#cfn-elasticmapreduce-instancefleetconfig-instancetypeconfig-ebsconfiguration
	EbsConfiguration AWSEMRInstanceFleetConfig_EbsConfiguration `json:"EbsConfiguration"`

	// InstanceType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancefleetconfig-instancetypeconfig.html#cfn-elasticmapreduce-instancefleetconfig-instancetypeconfig-instancetype
	InstanceType string `json:"InstanceType"`

	// WeightedCapacity AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancefleetconfig-instancetypeconfig.html#cfn-elasticmapreduce-instancefleetconfig-instancetypeconfig-weightedcapacity
	WeightedCapacity int64 `json:"WeightedCapacity"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRInstanceFleetConfig_InstanceTypeConfig) AWSCloudFormationType() string {
	return "AWS::EMR::InstanceFleetConfig.InstanceTypeConfig"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRInstanceFleetConfig_InstanceTypeConfig) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEMRInstanceFleetConfig_InstanceTypeConfigResources retrieves all AWSEMRInstanceFleetConfig_InstanceTypeConfig items from a CloudFormation template
func GetAllAWSEMRInstanceFleetConfig_InstanceTypeConfig(template *Template) map[string]*AWSEMRInstanceFleetConfig_InstanceTypeConfig {

	results := map[string]*AWSEMRInstanceFleetConfig_InstanceTypeConfig{}
	for name, resource := range template.Resources {
		result := &AWSEMRInstanceFleetConfig_InstanceTypeConfig{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEMRInstanceFleetConfig_InstanceTypeConfigWithName retrieves all AWSEMRInstanceFleetConfig_InstanceTypeConfig items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEMRInstanceFleetConfig_InstanceTypeConfig(name string, template *Template) (*AWSEMRInstanceFleetConfig_InstanceTypeConfig, error) {

	result := &AWSEMRInstanceFleetConfig_InstanceTypeConfig{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEMRInstanceFleetConfig_InstanceTypeConfig{}, errors.New("resource not found")

}
