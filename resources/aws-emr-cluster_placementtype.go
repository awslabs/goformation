package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EMR::Cluster.PlacementType AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig-placementtype.html
type AWSEMRCluster_PlacementType struct {

	// AvailabilityZone AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig-placementtype.html#aws-properties-emr-cluster-jobflowinstancesconfig-placementtype

	AvailabilityZone string `json:"AvailabilityZone"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRCluster_PlacementType) AWSCloudFormationType() string {
	return "AWS::EMR::Cluster.PlacementType"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRCluster_PlacementType) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEMRCluster_PlacementTypeResources retrieves all AWSEMRCluster_PlacementType items from a CloudFormation template
func GetAllAWSEMRCluster_PlacementType(template *Template) map[string]*AWSEMRCluster_PlacementType {

	results := map[string]*AWSEMRCluster_PlacementType{}
	for name, resource := range template.Resources {
		result := &AWSEMRCluster_PlacementType{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEMRCluster_PlacementTypeWithName retrieves all AWSEMRCluster_PlacementType items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEMRCluster_PlacementType(name string, template *Template) (*AWSEMRCluster_PlacementType, error) {

	result := &AWSEMRCluster_PlacementType{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEMRCluster_PlacementType{}, errors.New("resource not found")

}
