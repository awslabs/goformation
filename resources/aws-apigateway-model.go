package resources

// AWS::ApiGateway::Model AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-model.html
type AWSApiGatewayModel struct {
    
    // ContentType AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-model.html#cfn-apigateway-model-contenttype
    ContentType string `json:"ContentType"`
    
    // Description AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-model.html#cfn-apigateway-model-description
    Description string `json:"Description"`
    
    // Name AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-model.html#cfn-apigateway-model-name
    Name string `json:"Name"`
    
    // RestApiId AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-model.html#cfn-apigateway-model-restapiid
    RestApiId string `json:"RestApiId"`
    
    // Schema AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-model.html#cfn-apigateway-model-schema
    Schema string `json:"Schema"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSApiGatewayModel) AWSCloudFormationType() string {
    return "AWS::ApiGateway::Model"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSApiGatewayModel) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}