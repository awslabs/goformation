package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EC2::VPNConnection AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-connection.html
type AWSEC2VPNConnection struct {

	// CustomerGatewayId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-connection.html#cfn-ec2-vpnconnection-customergatewayid
	CustomerGatewayId string `json:"CustomerGatewayId"`

	// StaticRoutesOnly AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-connection.html#cfn-ec2-vpnconnection-StaticRoutesOnly
	StaticRoutesOnly bool `json:"StaticRoutesOnly"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-connection.html#cfn-ec2-vpnconnection-tags
	Tags []Tag `json:"Tags"`

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-connection.html#cfn-ec2-vpnconnection-type
	Type string `json:"Type"`

	// VpnGatewayId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-connection.html#cfn-ec2-vpnconnection-vpngatewayid
	VpnGatewayId string `json:"VpnGatewayId"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2VPNConnection) AWSCloudFormationType() string {
	return "AWS::EC2::VPNConnection"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2VPNConnection) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEC2VPNConnectionResources retrieves all AWSEC2VPNConnection items from a CloudFormation template
func GetAllAWSEC2VPNConnectionResources(template *Template) map[string]*AWSEC2VPNConnection {

	results := map[string]*AWSEC2VPNConnection{}
	for name, resource := range template.Resources {
		result := &AWSEC2VPNConnection{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEC2VPNConnectionWithName retrieves all AWSEC2VPNConnection items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSEC2VPNConnectionWithName(name string, template *Template) (*AWSEC2VPNConnection, error) {

	result := &AWSEC2VPNConnection{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEC2VPNConnection{}, errors.New("resource not found")

}
