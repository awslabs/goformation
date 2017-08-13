package resources

// AWS::ApiGateway::UsagePlan.ApiStage AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-apistage.html
type AWSApiGatewayUsagePlanApiStage struct {
    
    // ApiId AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-apistage.html#cfn-apigateway-usageplan-apistage-apiid
    ApiId string `json:"ApiId"`
    
    // Stage AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-apistage.html#cfn-apigateway-usageplan-apistage-stage
    Stage string `json:"Stage"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSApiGatewayUsagePlanApiStage) AWSCloudFormationType() string {
    return "AWS::ApiGateway::UsagePlan.ApiStage"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSApiGatewayUsagePlanApiStage) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}