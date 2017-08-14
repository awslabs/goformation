package resources

// AWS::Lambda::Function.VpcConfig AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-vpcconfig.html
type AWSLambdaFunctionVpcConfig struct {

	// SecurityGroupIds AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-vpcconfig.html#cfn-lambda-function-vpcconfig-securitygroupids
	SecurityGroupIds []AWSLambdaFunctionVpcConfigstring `json:"SecurityGroupIds"`

	// SubnetIds AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-vpcconfig.html#cfn-lambda-function-vpcconfig-subnetids
	SubnetIds []AWSLambdaFunctionVpcConfigstring `json:"SubnetIds"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSLambdaFunctionVpcConfig) AWSCloudFormationType() string {
	return "AWS::Lambda::Function.VpcConfig"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSLambdaFunctionVpcConfig) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
