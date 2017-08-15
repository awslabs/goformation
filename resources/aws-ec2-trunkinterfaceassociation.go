package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSEC2TrunkInterfaceAssociation AWS CloudFormation Resource (AWS::EC2::TrunkInterfaceAssociation)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trunkinterfaceassociation.html
type AWSEC2TrunkInterfaceAssociation struct {

	// BranchInterfaceId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trunkinterfaceassociation.html#cfn-ec2-trunkinterfaceassociation-branchinterfaceid
	BranchInterfaceId string `json:"BranchInterfaceId"`

	// GREKey AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trunkinterfaceassociation.html#cfn-ec2-trunkinterfaceassociation-grekey
	GREKey int64 `json:"GREKey"`

	// TrunkInterfaceId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trunkinterfaceassociation.html#cfn-ec2-trunkinterfaceassociation-trunkinterfaceid
	TrunkInterfaceId string `json:"TrunkInterfaceId"`

	// VLANId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trunkinterfaceassociation.html#cfn-ec2-trunkinterfaceassociation-vlanid
	VLANId int64 `json:"VLANId"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2TrunkInterfaceAssociation) AWSCloudFormationType() string {
	return "AWS::EC2::TrunkInterfaceAssociation"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2TrunkInterfaceAssociation) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEC2TrunkInterfaceAssociationResources retrieves all AWSEC2TrunkInterfaceAssociation items from a CloudFormation template
func GetAllAWSEC2TrunkInterfaceAssociationResources(template *Template) map[string]*AWSEC2TrunkInterfaceAssociation {

	results := map[string]*AWSEC2TrunkInterfaceAssociation{}
	for name, resource := range template.Resources {
		result := &AWSEC2TrunkInterfaceAssociation{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEC2TrunkInterfaceAssociationWithName retrieves all AWSEC2TrunkInterfaceAssociation items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSEC2TrunkInterfaceAssociationWithName(name string, template *Template) (*AWSEC2TrunkInterfaceAssociation, error) {

	result := &AWSEC2TrunkInterfaceAssociation{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEC2TrunkInterfaceAssociation{}, errors.New("resource not found")

}
