package resources

// AWS::ApiGateway::ApiKey AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-apikey.html
type AWSApiGatewayApiKey struct {
    
    // Description AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-apikey.html#cfn-apigateway-apigateway-apikey-description
    Description string `json:"Description"`
    
    // Enabled AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-apikey.html#cfn-apigateway-apigateway-apikey-enabled
    Enabled bool `json:"Enabled"`
    
    // Name AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-apikey.html#cfn-apigateway-apigateway-apikey-name
    Name string `json:"Name"`
    
    // StageKeys AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-apikey.html#cfn-apigateway-apigateway-apikey-stagekeys
    StageKeys []AWSApiGatewayApiKeyStageKey `json:"StageKeys"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSApiGatewayApiKey) AWSCloudFormationType() string {
    return "AWS::ApiGateway::ApiKey"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSApiGatewayApiKey) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}