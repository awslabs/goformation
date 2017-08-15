package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EMR::Cluster.InstanceFleetProvisioningSpecifications AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-instancefleetprovisioningspecifications.html
type AWSEMRCluster_InstanceFleetProvisioningSpecifications struct {

	// SpotSpecification AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-instancefleetprovisioningspecifications.html#cfn-elasticmapreduce-cluster-instancefleetprovisioningspecifications-spotspecification
	SpotSpecification AWSEMRCluster_SpotProvisioningSpecification `json:"SpotSpecification"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRCluster_InstanceFleetProvisioningSpecifications) AWSCloudFormationType() string {
	return "AWS::EMR::Cluster.InstanceFleetProvisioningSpecifications"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRCluster_InstanceFleetProvisioningSpecifications) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEMRCluster_InstanceFleetProvisioningSpecificationsResources retrieves all AWSEMRCluster_InstanceFleetProvisioningSpecifications items from a CloudFormation template
func GetAllAWSEMRCluster_InstanceFleetProvisioningSpecifications(template *Template) map[string]*AWSEMRCluster_InstanceFleetProvisioningSpecifications {

	results := map[string]*AWSEMRCluster_InstanceFleetProvisioningSpecifications{}
	for name, resource := range template.Resources {
		result := &AWSEMRCluster_InstanceFleetProvisioningSpecifications{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEMRCluster_InstanceFleetProvisioningSpecificationsWithName retrieves all AWSEMRCluster_InstanceFleetProvisioningSpecifications items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEMRCluster_InstanceFleetProvisioningSpecifications(name string, template *Template) (*AWSEMRCluster_InstanceFleetProvisioningSpecifications, error) {

	result := &AWSEMRCluster_InstanceFleetProvisioningSpecifications{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEMRCluster_InstanceFleetProvisioningSpecifications{}, errors.New("resource not found")

}
