package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSEC2VPCGatewayAttachment AWS CloudFormation Resource (AWS::EC2::VPCGatewayAttachment)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpc-gateway-attachment.html
type AWSEC2VPCGatewayAttachment struct {

	// InternetGatewayId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpc-gateway-attachment.html#cfn-ec2-vpcgatewayattachment-internetgatewayid
	InternetGatewayId string `json:"InternetGatewayId"`

	// VpcId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpc-gateway-attachment.html#cfn-ec2-vpcgatewayattachment-vpcid
	VpcId string `json:"VpcId"`

	// VpnGatewayId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpc-gateway-attachment.html#cfn-ec2-vpcgatewayattachment-vpngatewayid
	VpnGatewayId string `json:"VpnGatewayId"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2VPCGatewayAttachment) AWSCloudFormationType() string {
	return "AWS::EC2::VPCGatewayAttachment"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2VPCGatewayAttachment) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEC2VPCGatewayAttachmentResources retrieves all AWSEC2VPCGatewayAttachment items from a CloudFormation template
func GetAllAWSEC2VPCGatewayAttachmentResources(template *Template) map[string]*AWSEC2VPCGatewayAttachment {

	results := map[string]*AWSEC2VPCGatewayAttachment{}
	for name, resource := range template.Resources {
		result := &AWSEC2VPCGatewayAttachment{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEC2VPCGatewayAttachmentWithName retrieves all AWSEC2VPCGatewayAttachment items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSEC2VPCGatewayAttachmentWithName(name string, template *Template) (*AWSEC2VPCGatewayAttachment, error) {

	result := &AWSEC2VPCGatewayAttachment{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEC2VPCGatewayAttachment{}, errors.New("resource not found")

}
