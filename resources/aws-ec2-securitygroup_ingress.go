package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EC2::SecurityGroup.Ingress AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-rule.html
type AWSEC2SecurityGroup_Ingress struct {

	// CidrIp AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-rule.html#cfn-ec2-security-group-rule-cidrip
	CidrIp string `json:"CidrIp"`

	// CidrIpv6 AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-rule.html#cfn-ec2-security-group-rule-cidripv6
	CidrIpv6 string `json:"CidrIpv6"`

	// FromPort AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-rule.html#cfn-ec2-security-group-rule-fromport
	FromPort int64 `json:"FromPort"`

	// IpProtocol AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-rule.html#cfn-ec2-security-group-rule-ipprotocol
	IpProtocol string `json:"IpProtocol"`

	// SourceSecurityGroupId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-rule.html#cfn-ec2-security-group-rule-sourcesecuritygroupid
	SourceSecurityGroupId string `json:"SourceSecurityGroupId"`

	// SourceSecurityGroupName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-rule.html#cfn-ec2-security-group-rule-sourcesecuritygroupname
	SourceSecurityGroupName string `json:"SourceSecurityGroupName"`

	// SourceSecurityGroupOwnerId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-rule.html#cfn-ec2-security-group-rule-sourcesecuritygroupownerid
	SourceSecurityGroupOwnerId string `json:"SourceSecurityGroupOwnerId"`

	// ToPort AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-rule.html#cfn-ec2-security-group-rule-toport
	ToPort int64 `json:"ToPort"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2SecurityGroup_Ingress) AWSCloudFormationType() string {
	return "AWS::EC2::SecurityGroup.Ingress"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2SecurityGroup_Ingress) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEC2SecurityGroup_IngressResources retrieves all AWSEC2SecurityGroup_Ingress items from a CloudFormation template
func GetAllAWSEC2SecurityGroup_Ingress(template *Template) map[string]*AWSEC2SecurityGroup_Ingress {

	results := map[string]*AWSEC2SecurityGroup_Ingress{}
	for name, resource := range template.Resources {
		result := &AWSEC2SecurityGroup_Ingress{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEC2SecurityGroup_IngressWithName retrieves all AWSEC2SecurityGroup_Ingress items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEC2SecurityGroup_Ingress(name string, template *Template) (*AWSEC2SecurityGroup_Ingress, error) {

	result := &AWSEC2SecurityGroup_Ingress{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEC2SecurityGroup_Ingress{}, errors.New("resource not found")

}
