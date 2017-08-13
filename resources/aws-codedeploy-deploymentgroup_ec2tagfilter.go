package resources

// AWS::CodeDeploy::DeploymentGroup.EC2TagFilter AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-ec2tagfilters.html
type AWSCodeDeployDeploymentGroupEC2TagFilter struct {
    
    // Key AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-ec2tagfilters.html#cfn-properties-codedeploy-deploymentgroup-ec2tagfilters-key
    Key string `json:"Key"`
    
    // Type AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-ec2tagfilters.html#cfn-properties-codedeploy-deploymentgroup-ec2tagfilters-type
    Type string `json:"Type"`
    
    // Value AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-ec2tagfilters.html#cfn-properties-codedeploy-deploymentgroup-ec2tagfilters-value
    Value string `json:"Value"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodeDeployDeploymentGroupEC2TagFilter) AWSCloudFormationType() string {
    return "AWS::CodeDeploy::DeploymentGroup.EC2TagFilter"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCodeDeployDeploymentGroupEC2TagFilter) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}