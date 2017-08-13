package resources

// AWS::CodeDeploy::DeploymentGroup.RevisionLocation AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-deployment-revision.html
type AWSCodeDeployDeploymentGroupRevisionLocation struct {
    
    // GitHubLocation AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-deployment-revision.html#cfn-properties-codedeploy-deploymentgroup-deployment-revision-githublocation
    GitHubLocation AWSCodeDeployDeploymentGroupRevisionLocationGitHubLocation `json:"GitHubLocation"`
    
    // RevisionType AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-deployment-revision.html#cfn-properties-codedeploy-deploymentgroup-deployment-revision-revisiontype
    RevisionType string `json:"RevisionType"`
    
    // S3Location AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-deployment-revision.html#cfn-properties-codedeploy-deploymentgroup-deployment-revision-s3location
    S3Location AWSCodeDeployDeploymentGroupRevisionLocationS3Location `json:"S3Location"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodeDeployDeploymentGroupRevisionLocation) AWSCloudFormationType() string {
    return "AWS::CodeDeploy::DeploymentGroup.RevisionLocation"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCodeDeployDeploymentGroupRevisionLocation) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}