package resources

// AWS::Lambda::Function.DeadLetterConfig AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-deadletterconfig.html
type AWSLambdaFunctionDeadLetterConfig struct {
    
    // TargetArn AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-deadletterconfig.html#cfn-lambda-function-deadletterconfig-targetarn
    TargetArn string `json:"TargetArn"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSLambdaFunctionDeadLetterConfig) AWSCloudFormationType() string {
    return "AWS::Lambda::Function.DeadLetterConfig"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSLambdaFunctionDeadLetterConfig) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}