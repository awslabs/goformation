package resources

// AWS::EC2::Instance.AssociationParameter AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-ssmassociations-associationparameters.html
type AWSEC2InstanceAssociationParameter struct {

	// Key AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-ssmassociations-associationparameters.html#cfn-ec2-instance-ssmassociations-associationparameters-key
	Key string `json:"Key"`

	// Value AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-ssmassociations-associationparameters.html#cfn-ec2-instance-ssmassociations-associationparameters-value
	Value []AWSEC2InstanceAssociationParameterstring `json:"Value"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2InstanceAssociationParameter) AWSCloudFormationType() string {
	return "AWS::EC2::Instance.AssociationParameter"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2InstanceAssociationParameter) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
