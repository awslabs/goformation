package cloudformation

import (
	"encoding/json"
)

// AWSEC2Instance_AssociationParameter AWS CloudFormation Resource (AWS::EC2::Instance.AssociationParameter)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-ssmassociations-associationparameters.html
type AWSEC2Instance_AssociationParameter struct {

	// Key AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-ssmassociations-associationparameters.html#cfn-ec2-instance-ssmassociations-associationparameters-key
	Key *Value `json:"Key,omitempty"`

	// Value AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-ssmassociations-associationparameters.html#cfn-ec2-instance-ssmassociations-associationparameters-value
	Value []*Value `json:"Value,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2Instance_AssociationParameter) AWSCloudFormationType() string {
	return "AWS::EC2::Instance.AssociationParameter"
}

func (r *AWSEC2Instance_AssociationParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}
