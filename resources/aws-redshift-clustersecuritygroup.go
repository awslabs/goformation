package resources

// AWS::Redshift::ClusterSecurityGroup AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clustersecuritygroup.html
type AWSRedshiftClusterSecurityGroup struct {
    
    // Description AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clustersecuritygroup.html#cfn-redshift-clustersecuritygroup-description
    Description string `json:"Description"`
    
    // Tags AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clustersecuritygroup.html#cfn-redshift-clustersecuritygroup-tags
    Tags []AWSRedshiftClusterSecurityGroupTag `json:"Tags"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSRedshiftClusterSecurityGroup) AWSCloudFormationType() string {
    return "AWS::Redshift::ClusterSecurityGroup"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSRedshiftClusterSecurityGroup) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}