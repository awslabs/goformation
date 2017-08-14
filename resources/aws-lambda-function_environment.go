package resources

// AWS::Lambda::Function.Environment AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-environment.html
type AWSLambdaFunctionEnvironment struct {

	// Variables AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-environment.html#cfn-lambda-function-environment-variables
	Variables map[string]AWSLambdaFunctionEnvironmentstring `json:"Variables"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSLambdaFunctionEnvironment) AWSCloudFormationType() string {
	return "AWS::Lambda::Function.Environment"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSLambdaFunctionEnvironment) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
