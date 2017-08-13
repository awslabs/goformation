package resources

// AWS::CodeDeploy::DeploymentGroup.TagFilter AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-onpremisesinstancetagfilters.html
type AWSCodeDeployDeploymentGroupTagFilter struct {
    
    // Key AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-onpremisesinstancetagfilters.html#cfn-properties-codedeploy-deploymentgroup-onpremisesinstancetagfilters-key
    Key string `json:"Key"`
    
    // Type AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-onpremisesinstancetagfilters.html#cfn-properties-codedeploy-deploymentgroup-onpremisesinstancetagfilters-type
    Type string `json:"Type"`
    
    // Value AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-onpremisesinstancetagfilters.html#cfn-properties-codedeploy-deploymentgroup-onpremisesinstancetagfilters-value
    Value string `json:"Value"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodeDeployDeploymentGroupTagFilter) AWSCloudFormationType() string {
    return "AWS::CodeDeploy::DeploymentGroup.TagFilter"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCodeDeployDeploymentGroupTagFilter) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}