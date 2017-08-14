package resources

// AWS::ApiGateway::Account AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-account.html
type AWSApiGatewayAccount struct {

	// CloudWatchRoleArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-account.html#cfn-apigateway-account-cloudwatchrolearn
	CloudWatchRoleArn string `json:"CloudWatchRoleArn"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSApiGatewayAccount) AWSCloudFormationType() string {
	return "AWS::ApiGateway::Account"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSApiGatewayAccount) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
