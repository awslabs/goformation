package resources

// AWS::Lambda::Function.Environment AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-environment.html
type AWSLambdaFunction_Environment struct {

	// Variables AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-environment.html#cfn-lambda-function-environment-variables
	Variables map[string]string `json:"Variables"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSLambdaFunction_Environment) AWSCloudFormationType() string {
	return "AWS::Lambda::Function.Environment"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSLambdaFunction_Environment) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
