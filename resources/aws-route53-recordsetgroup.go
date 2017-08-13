package resources

// AWS::Route53::RecordSetGroup AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-recordsetgroup.html
type AWSRoute53RecordSetGroup struct {
    
    // Comment AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-recordsetgroup.html#cfn-route53-recordsetgroup-comment
    Comment string `json:"Comment"`
    
    // HostedZoneId AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-recordsetgroup.html#cfn-route53-recordsetgroup-hostedzoneid
    HostedZoneId string `json:"HostedZoneId"`
    
    // HostedZoneName AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-recordsetgroup.html#cfn-route53-recordsetgroup-hostedzonename
    HostedZoneName string `json:"HostedZoneName"`
    
    // RecordSets AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-recordsetgroup.html#cfn-route53-recordsetgroup-recordsets
    RecordSets []AWSRoute53RecordSetGroupRecordSet `json:"RecordSets"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSRoute53RecordSetGroup) AWSCloudFormationType() string {
    return "AWS::Route53::RecordSetGroup"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSRoute53RecordSetGroup) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}