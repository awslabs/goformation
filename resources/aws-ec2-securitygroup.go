package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EC2::SecurityGroup AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group.html
type AWSEC2SecurityGroup struct {

	// GroupDescription AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group.html#cfn-ec2-securitygroup-groupdescription
	GroupDescription string `json:"GroupDescription"`

	// GroupName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group.html#cfn-ec2-securitygroup-groupname
	GroupName string `json:"GroupName"`

	// SecurityGroupEgress AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group.html#cfn-ec2-securitygroup-securitygroupegress
	SecurityGroupEgress []AWSEC2SecurityGroup_Egress `json:"SecurityGroupEgress"`

	// SecurityGroupIngress AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group.html#cfn-ec2-securitygroup-securitygroupingress
	SecurityGroupIngress []AWSEC2SecurityGroup_Ingress `json:"SecurityGroupIngress"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group.html#cfn-ec2-securitygroup-tags
	Tags []Tag `json:"Tags"`

	// VpcId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group.html#cfn-ec2-securitygroup-vpcid
	VpcId string `json:"VpcId"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2SecurityGroup) AWSCloudFormationType() string {
	return "AWS::EC2::SecurityGroup"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2SecurityGroup) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEC2SecurityGroupResources retrieves all AWSEC2SecurityGroup items from a CloudFormation template
func GetAllAWSEC2SecurityGroup(template *Template) map[string]*AWSEC2SecurityGroup {

	results := map[string]*AWSEC2SecurityGroup{}
	for name, resource := range template.Resources {
		result := &AWSEC2SecurityGroup{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEC2SecurityGroupWithName retrieves all AWSEC2SecurityGroup items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEC2SecurityGroup(name string, template *Template) (*AWSEC2SecurityGroup, error) {

	result := &AWSEC2SecurityGroup{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEC2SecurityGroup{}, errors.New("resource not found")

}
