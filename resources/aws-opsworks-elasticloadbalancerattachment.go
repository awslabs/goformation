package resources

// AWS::OpsWorks::ElasticLoadBalancerAttachment AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-elbattachment.html
type AWSOpsWorksElasticLoadBalancerAttachment struct {
    
    // ElasticLoadBalancerName AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-elbattachment.html#cfn-opsworks-elbattachment-elbname
    ElasticLoadBalancerName string `json:"ElasticLoadBalancerName"`
    
    // LayerId AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-elbattachment.html#cfn-opsworks-elbattachment-layerid
    LayerId string `json:"LayerId"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSOpsWorksElasticLoadBalancerAttachment) AWSCloudFormationType() string {
    return "AWS::OpsWorks::ElasticLoadBalancerAttachment"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSOpsWorksElasticLoadBalancerAttachment) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}