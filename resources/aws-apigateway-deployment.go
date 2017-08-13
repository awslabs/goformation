package resources

// AWS::ApiGateway::Deployment AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-deployment.html
type AWSApiGatewayDeployment struct {
    
    // Description AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-deployment.html#cfn-apigateway-deployment-description
    Description string `json:"Description"`
    
    // RestApiId AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-deployment.html#cfn-apigateway-deployment-restapiid
    RestApiId string `json:"RestApiId"`
    
    // StageDescription AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-deployment.html#cfn-apigateway-deployment-stagedescription
    StageDescription AWSApiGatewayDeploymentStageDescription `json:"StageDescription"`
    
    // StageName AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-deployment.html#cfn-apigateway-deployment-stagename
    StageName string `json:"StageName"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSApiGatewayDeployment) AWSCloudFormationType() string {
    return "AWS::ApiGateway::Deployment"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSApiGatewayDeployment) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}