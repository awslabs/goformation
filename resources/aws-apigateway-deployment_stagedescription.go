package resources

// AWS::ApiGateway::Deployment.StageDescription AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription.html
type AWSApiGatewayDeploymentStageDescription struct {
    
    // CacheClusterEnabled AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-cacheclusterenabled
    CacheClusterEnabled bool `json:"CacheClusterEnabled"`
    
    // CacheClusterSize AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-cacheclustersize
    CacheClusterSize string `json:"CacheClusterSize"`
    
    // CacheDataEncrypted AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-cachedataencrypted
    CacheDataEncrypted bool `json:"CacheDataEncrypted"`
    
    // CacheTtlInSeconds AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-cachettlinseconds
    CacheTtlInSeconds int64 `json:"CacheTtlInSeconds"`
    
    // CachingEnabled AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-cachingenabled
    CachingEnabled bool `json:"CachingEnabled"`
    
    // ClientCertificateId AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-clientcertificateid
    ClientCertificateId string `json:"ClientCertificateId"`
    
    // DataTraceEnabled AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-datatraceenabled
    DataTraceEnabled bool `json:"DataTraceEnabled"`
    
    // Description AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-description
    Description string `json:"Description"`
    
    // LoggingLevel AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-logginglevel
    LoggingLevel string `json:"LoggingLevel"`
    
    // MethodSettings AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-methodsettings
    MethodSettings []AWSApiGatewayDeploymentStageDescriptionMethodSetting `json:"MethodSettings"`
    
    // MetricsEnabled AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-metricsenabled
    MetricsEnabled bool `json:"MetricsEnabled"`
    
    // StageName AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-stagename
    StageName string `json:"StageName"`
    
    // ThrottlingBurstLimit AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-throttlingburstlimit
    ThrottlingBurstLimit int64 `json:"ThrottlingBurstLimit"`
    
    // ThrottlingRateLimit AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-throttlingratelimit
    ThrottlingRateLimit float64 `json:"ThrottlingRateLimit"`
    
    // Variables AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-variables
    Variables map[string]string `json:"Variables"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSApiGatewayDeploymentStageDescription) AWSCloudFormationType() string {
    return "AWS::ApiGateway::Deployment.StageDescription"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSApiGatewayDeploymentStageDescription) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}