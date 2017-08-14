package resources

// AWS::Redshift::ClusterParameterGroup.Parameter AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-property-redshift-clusterparametergroup-parameter.html
type AWSRedshiftClusterParameterGroup_Parameter struct {

	// ParameterName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-property-redshift-clusterparametergroup-parameter.html#cfn-redshift-clusterparametergroup-parameter-parametername

	ParameterName string `json:"ParameterName"`

	// ParameterValue AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-property-redshift-clusterparametergroup-parameter.html#cfn-redshift-clusterparametergroup-parameter-parametervalue

	ParameterValue string `json:"ParameterValue"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSRedshiftClusterParameterGroup_Parameter) AWSCloudFormationType() string {
	return "AWS::Redshift::ClusterParameterGroup.Parameter"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSRedshiftClusterParameterGroup_Parameter) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
