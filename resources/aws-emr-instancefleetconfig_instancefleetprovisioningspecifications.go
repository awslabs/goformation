package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EMR::InstanceFleetConfig.InstanceFleetProvisioningSpecifications AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancefleetconfig-instancefleetprovisioningspecifications.html
type AWSEMRInstanceFleetConfig_InstanceFleetProvisioningSpecifications struct {

	// SpotSpecification AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancefleetconfig-instancefleetprovisioningspecifications.html#cfn-elasticmapreduce-instancefleetconfig-instancefleetprovisioningspecifications-spotspecification

	SpotSpecification AWSEMRInstanceFleetConfig_SpotProvisioningSpecification `json:"SpotSpecification"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRInstanceFleetConfig_InstanceFleetProvisioningSpecifications) AWSCloudFormationType() string {
	return "AWS::EMR::InstanceFleetConfig.InstanceFleetProvisioningSpecifications"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRInstanceFleetConfig_InstanceFleetProvisioningSpecifications) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEMRInstanceFleetConfig_InstanceFleetProvisioningSpecificationsResources retrieves all AWSEMRInstanceFleetConfig_InstanceFleetProvisioningSpecifications items from a CloudFormation template
func GetAllAWSEMRInstanceFleetConfig_InstanceFleetProvisioningSpecifications(template *Template) map[string]*AWSEMRInstanceFleetConfig_InstanceFleetProvisioningSpecifications {

	results := map[string]*AWSEMRInstanceFleetConfig_InstanceFleetProvisioningSpecifications{}
	for name, resource := range template.Resources {
		result := &AWSEMRInstanceFleetConfig_InstanceFleetProvisioningSpecifications{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEMRInstanceFleetConfig_InstanceFleetProvisioningSpecificationsWithName retrieves all AWSEMRInstanceFleetConfig_InstanceFleetProvisioningSpecifications items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEMRInstanceFleetConfig_InstanceFleetProvisioningSpecifications(name string, template *Template) (*AWSEMRInstanceFleetConfig_InstanceFleetProvisioningSpecifications, error) {

	result := &AWSEMRInstanceFleetConfig_InstanceFleetProvisioningSpecifications{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEMRInstanceFleetConfig_InstanceFleetProvisioningSpecifications{}, errors.New("resource not found")

}
