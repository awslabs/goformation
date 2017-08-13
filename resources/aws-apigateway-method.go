package resources

// AWS::ApiGateway::Method AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html
type AWSApiGatewayMethod struct {
    
    // ApiKeyRequired AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-apikeyrequired
    ApiKeyRequired bool `json:"ApiKeyRequired"`
    
    // AuthorizationType AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-authorizationtype
    AuthorizationType string `json:"AuthorizationType"`
    
    // AuthorizerId AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-authorizerid
    AuthorizerId string `json:"AuthorizerId"`
    
    // HttpMethod AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-httpmethod
    HttpMethod string `json:"HttpMethod"`
    
    // Integration AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-integration
    Integration AWSApiGatewayMethodIntegration `json:"Integration"`
    
    // MethodResponses AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-methodresponses
    MethodResponses []AWSApiGatewayMethodMethodResponse `json:"MethodResponses"`
    
    // RequestModels AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-requestmodels
    RequestModels map[string]string `json:"RequestModels"`
    
    // RequestParameters AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-requestparameters
    RequestParameters map[string]bool `json:"RequestParameters"`
    
    // ResourceId AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-resourceid
    ResourceId string `json:"ResourceId"`
    
    // RestApiId AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-restapiid
    RestApiId string `json:"RestApiId"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSApiGatewayMethod) AWSCloudFormationType() string {
    return "AWS::ApiGateway::Method"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSApiGatewayMethod) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}