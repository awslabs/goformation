package resources

// AWS::ApiGateway::RestApi.S3Location AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-restapi-bodys3location.html
type AWSApiGatewayRestApiS3Location struct {
    
    // Bucket AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-restapi-bodys3location.html#cfn-apigateway-restapi-s3location-bucket
    Bucket string `json:"Bucket"`
    
    // ETag AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-restapi-bodys3location.html#cfn-apigateway-restapi-s3location-etag
    ETag string `json:"ETag"`
    
    // Key AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-restapi-bodys3location.html#cfn-apigateway-restapi-s3location-key
    Key string `json:"Key"`
    
    // Version AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-restapi-bodys3location.html#cfn-apigateway-restapi-s3location-version
    Version string `json:"Version"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSApiGatewayRestApiS3Location) AWSCloudFormationType() string {
    return "AWS::ApiGateway::RestApi.S3Location"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSApiGatewayRestApiS3Location) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}