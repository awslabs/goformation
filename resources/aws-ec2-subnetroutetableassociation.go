package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EC2::SubnetRouteTableAssociation AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet-route-table-assoc.html
type AWSEC2SubnetRouteTableAssociation struct {

	// RouteTableId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet-route-table-assoc.html#cfn-ec2-subnetroutetableassociation-routetableid
	RouteTableId string `json:"RouteTableId"`

	// SubnetId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet-route-table-assoc.html#cfn-ec2-subnetroutetableassociation-subnetid
	SubnetId string `json:"SubnetId"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2SubnetRouteTableAssociation) AWSCloudFormationType() string {
	return "AWS::EC2::SubnetRouteTableAssociation"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2SubnetRouteTableAssociation) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEC2SubnetRouteTableAssociationResources retrieves all AWSEC2SubnetRouteTableAssociation items from a CloudFormation template
func GetAllAWSEC2SubnetRouteTableAssociationResources(template *Template) map[string]*AWSEC2SubnetRouteTableAssociation {

	results := map[string]*AWSEC2SubnetRouteTableAssociation{}
	for name, resource := range template.Resources {
		result := &AWSEC2SubnetRouteTableAssociation{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEC2SubnetRouteTableAssociationWithName retrieves all AWSEC2SubnetRouteTableAssociation items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSEC2SubnetRouteTableAssociationWithName(name string, template *Template) (*AWSEC2SubnetRouteTableAssociation, error) {

	result := &AWSEC2SubnetRouteTableAssociation{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEC2SubnetRouteTableAssociation{}, errors.New("resource not found")

}
