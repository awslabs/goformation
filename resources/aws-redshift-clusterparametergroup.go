package resources

// AWS::Redshift::ClusterParameterGroup AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clusterparametergroup.html
type AWSRedshiftClusterParameterGroup struct {
    
    // Description AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clusterparametergroup.html#cfn-redshift-clusterparametergroup-description
    Description string `json:"Description"`
    
    // ParameterGroupFamily AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clusterparametergroup.html#cfn-redshift-clusterparametergroup-parametergroupfamily
    ParameterGroupFamily string `json:"ParameterGroupFamily"`
    
    // Parameters AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clusterparametergroup.html#cfn-redshift-clusterparametergroup-parameters
    Parameters []AWSRedshiftClusterParameterGroupParameter `json:"Parameters"`
    
    // Tags AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clusterparametergroup.html#cfn-redshift-clusterparametergroup-tags
    Tags []AWSRedshiftClusterParameterGroupTag `json:"Tags"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSRedshiftClusterParameterGroup) AWSCloudFormationType() string {
    return "AWS::Redshift::ClusterParameterGroup"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSRedshiftClusterParameterGroup) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}