package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EMR::Cluster.SpotProvisioningSpecification AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-spotprovisioningspecification.html
type AWSEMRCluster_SpotProvisioningSpecification struct {

	// BlockDurationMinutes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-spotprovisioningspecification.html#cfn-elasticmapreduce-cluster-spotprovisioningspecification-blockdurationminutes
	BlockDurationMinutes int64 `json:"BlockDurationMinutes"`

	// TimeoutAction AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-spotprovisioningspecification.html#cfn-elasticmapreduce-cluster-spotprovisioningspecification-timeoutaction
	TimeoutAction string `json:"TimeoutAction"`

	// TimeoutDurationMinutes AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-spotprovisioningspecification.html#cfn-elasticmapreduce-cluster-spotprovisioningspecification-timeoutdurationminutes
	TimeoutDurationMinutes int64 `json:"TimeoutDurationMinutes"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRCluster_SpotProvisioningSpecification) AWSCloudFormationType() string {
	return "AWS::EMR::Cluster.SpotProvisioningSpecification"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRCluster_SpotProvisioningSpecification) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEMRCluster_SpotProvisioningSpecificationResources retrieves all AWSEMRCluster_SpotProvisioningSpecification items from a CloudFormation template
func GetAllAWSEMRCluster_SpotProvisioningSpecification(template *Template) map[string]*AWSEMRCluster_SpotProvisioningSpecification {

	results := map[string]*AWSEMRCluster_SpotProvisioningSpecification{}
	for name, resource := range template.Resources {
		result := &AWSEMRCluster_SpotProvisioningSpecification{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEMRCluster_SpotProvisioningSpecificationWithName retrieves all AWSEMRCluster_SpotProvisioningSpecification items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEMRCluster_SpotProvisioningSpecification(name string, template *Template) (*AWSEMRCluster_SpotProvisioningSpecification, error) {

	result := &AWSEMRCluster_SpotProvisioningSpecification{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEMRCluster_SpotProvisioningSpecification{}, errors.New("resource not found")

}
