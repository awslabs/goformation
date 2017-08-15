package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EMR::InstanceFleetConfig.SpotProvisioningSpecification AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancefleetconfig-spotprovisioningspecification.html
type AWSEMRInstanceFleetConfig_SpotProvisioningSpecification struct {

	// BlockDurationMinutes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancefleetconfig-spotprovisioningspecification.html#cfn-elasticmapreduce-instancefleetconfig-spotprovisioningspecification-blockdurationminutes
	BlockDurationMinutes int64 `json:"BlockDurationMinutes"`

	// TimeoutAction AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancefleetconfig-spotprovisioningspecification.html#cfn-elasticmapreduce-instancefleetconfig-spotprovisioningspecification-timeoutaction
	TimeoutAction string `json:"TimeoutAction"`

	// TimeoutDurationMinutes AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancefleetconfig-spotprovisioningspecification.html#cfn-elasticmapreduce-instancefleetconfig-spotprovisioningspecification-timeoutdurationminutes
	TimeoutDurationMinutes int64 `json:"TimeoutDurationMinutes"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRInstanceFleetConfig_SpotProvisioningSpecification) AWSCloudFormationType() string {
	return "AWS::EMR::InstanceFleetConfig.SpotProvisioningSpecification"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRInstanceFleetConfig_SpotProvisioningSpecification) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEMRInstanceFleetConfig_SpotProvisioningSpecificationResources retrieves all AWSEMRInstanceFleetConfig_SpotProvisioningSpecification items from a CloudFormation template
func GetAllAWSEMRInstanceFleetConfig_SpotProvisioningSpecification(template *Template) map[string]*AWSEMRInstanceFleetConfig_SpotProvisioningSpecification {

	results := map[string]*AWSEMRInstanceFleetConfig_SpotProvisioningSpecification{}
	for name, resource := range template.Resources {
		result := &AWSEMRInstanceFleetConfig_SpotProvisioningSpecification{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEMRInstanceFleetConfig_SpotProvisioningSpecificationWithName retrieves all AWSEMRInstanceFleetConfig_SpotProvisioningSpecification items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEMRInstanceFleetConfig_SpotProvisioningSpecification(name string, template *Template) (*AWSEMRInstanceFleetConfig_SpotProvisioningSpecification, error) {

	result := &AWSEMRInstanceFleetConfig_SpotProvisioningSpecification{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEMRInstanceFleetConfig_SpotProvisioningSpecification{}, errors.New("resource not found")

}
