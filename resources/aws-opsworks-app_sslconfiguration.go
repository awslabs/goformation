package resources

// AWS::OpsWorks::App.SslConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-app-sslconfiguration.html
type AWSOpsWorksAppSslConfiguration struct {
    
    // Certificate AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-app-sslconfiguration.html#cfn-opsworks-app-sslconfig-certificate
    Certificate string `json:"Certificate"`
    
    // Chain AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-app-sslconfiguration.html#cfn-opsworks-app-sslconfig-chain
    Chain string `json:"Chain"`
    
    // PrivateKey AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-app-sslconfiguration.html#cfn-opsworks-app-sslconfig-privatekey
    PrivateKey string `json:"PrivateKey"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSOpsWorksAppSslConfiguration) AWSCloudFormationType() string {
    return "AWS::OpsWorks::App.SslConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSOpsWorksAppSslConfiguration) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}