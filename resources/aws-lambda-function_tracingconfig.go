package resources

// AWS::Lambda::Function.TracingConfig AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-tracingconfig.html
type AWSLambdaFunctionTracingConfig struct {
    
    // Mode AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-tracingconfig.html#cfn-lambda-function-tracingconfig-mode
    Mode string `json:"Mode"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSLambdaFunctionTracingConfig) AWSCloudFormationType() string {
    return "AWS::Lambda::Function.TracingConfig"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSLambdaFunctionTracingConfig) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}