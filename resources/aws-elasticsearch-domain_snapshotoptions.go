package resources

// AWS::Elasticsearch::Domain.SnapshotOptions AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticsearch-domain-snapshotoptions.html
type AWSElasticsearchDomainSnapshotOptions struct {
    
    // AutomatedSnapshotStartHour AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticsearch-domain-snapshotoptions.html#cfn-elasticsearch-domain-snapshotoptions-automatedsnapshotstarthour
    AutomatedSnapshotStartHour int64 `json:"AutomatedSnapshotStartHour"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticsearchDomainSnapshotOptions) AWSCloudFormationType() string {
    return "AWS::Elasticsearch::Domain.SnapshotOptions"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSElasticsearchDomainSnapshotOptions) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}