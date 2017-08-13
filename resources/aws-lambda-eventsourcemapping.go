package resources

// AWS::Lambda::EventSourceMapping AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html
type AWSLambdaEventSourceMapping struct {
    
    // BatchSize AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-batchsize
    BatchSize int64 `json:"BatchSize"`
    
    // Enabled AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-enabled
    Enabled bool `json:"Enabled"`
    
    // EventSourceArn AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-eventsourcearn
    EventSourceArn string `json:"EventSourceArn"`
    
    // FunctionName AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-functionname
    FunctionName string `json:"FunctionName"`
    
    // StartingPosition AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-startingposition
    StartingPosition string `json:"StartingPosition"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSLambdaEventSourceMapping) AWSCloudFormationType() string {
    return "AWS::Lambda::EventSourceMapping"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSLambdaEventSourceMapping) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}