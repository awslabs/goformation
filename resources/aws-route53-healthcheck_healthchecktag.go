package resources

// AWS::Route53::HealthCheck.HealthCheckTag AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-healthchecktag.html
type AWSRoute53HealthCheckHealthCheckTag struct {
    
    // Key AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-healthchecktag.html#cfn-route53-healthchecktags-key
    Key string `json:"Key"`
    
    // Value AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-healthchecktag.html#cfn-route53-healthchecktags-value
    Value string `json:"Value"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSRoute53HealthCheckHealthCheckTag) AWSCloudFormationType() string {
    return "AWS::Route53::HealthCheck.HealthCheckTag"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSRoute53HealthCheckHealthCheckTag) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}