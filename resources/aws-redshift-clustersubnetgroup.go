package resources

// AWS::Redshift::ClusterSubnetGroup AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clustersubnetgroup.html
type AWSRedshiftClusterSubnetGroup struct {

	// Description AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clustersubnetgroup.html#cfn-redshift-clustersubnetgroup-description
	Description string `json:"Description"`

	// SubnetIds AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clustersubnetgroup.html#cfn-redshift-clustersubnetgroup-subnetids
	SubnetIds []AWSRedshiftClusterSubnetGroupstring `json:"SubnetIds"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clustersubnetgroup.html#cfn-redshift-clustersubnetgroup-tags
	Tags []AWSRedshiftClusterSubnetGroupTag `json:"Tags"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSRedshiftClusterSubnetGroup) AWSCloudFormationType() string {
	return "AWS::Redshift::ClusterSubnetGroup"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSRedshiftClusterSubnetGroup) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
