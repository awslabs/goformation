package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSEC2RouteTable AWS CloudFormation Resource (AWS::EC2::RouteTable)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-route-table.html
type AWSEC2RouteTable struct {

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-route-table.html#cfn-ec2-routetable-tags
	Tags []Tag `json:"Tags"`

	// VpcId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-route-table.html#cfn-ec2-routetable-vpcid
	VpcId string `json:"VpcId"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2RouteTable) AWSCloudFormationType() string {
	return "AWS::EC2::RouteTable"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2RouteTable) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEC2RouteTableResources retrieves all AWSEC2RouteTable items from a CloudFormation template
func GetAllAWSEC2RouteTableResources(template *Template) map[string]*AWSEC2RouteTable {

	results := map[string]*AWSEC2RouteTable{}
	for name, resource := range template.Resources {
		result := &AWSEC2RouteTable{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEC2RouteTableWithName retrieves all AWSEC2RouteTable items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSEC2RouteTableWithName(name string, template *Template) (*AWSEC2RouteTable, error) {

	result := &AWSEC2RouteTable{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEC2RouteTable{}, errors.New("resource not found")

}
