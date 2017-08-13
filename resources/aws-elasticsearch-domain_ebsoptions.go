package resources

// AWS::Elasticsearch::Domain.EBSOptions AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticsearch-domain-ebsoptions.html
type AWSElasticsearchDomainEBSOptions struct {
    
    // EBSEnabled AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticsearch-domain-ebsoptions.html#cfn-elasticsearch-domain-ebsoptions-ebsenabled
    EBSEnabled bool `json:"EBSEnabled"`
    
    // Iops AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticsearch-domain-ebsoptions.html#cfn-elasticsearch-domain-ebsoptions-iops
    Iops int64 `json:"Iops"`
    
    // VolumeSize AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticsearch-domain-ebsoptions.html#cfn-elasticsearch-domain-ebsoptions-volumesize
    VolumeSize int64 `json:"VolumeSize"`
    
    // VolumeType AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticsearch-domain-ebsoptions.html#cfn-elasticsearch-domain-ebsoptions-volumetype
    VolumeType string `json:"VolumeType"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticsearchDomainEBSOptions) AWSCloudFormationType() string {
    return "AWS::Elasticsearch::Domain.EBSOptions"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSElasticsearchDomainEBSOptions) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}