package resources

// AWS::CodeBuild::Project.Source AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-source.html
type AWSCodeBuildProjectSource struct {
    
    // Auth AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-source.html#cfn-codebuild-project-source-auth
    Auth AWSCodeBuildProjectSourceSourceAuth `json:"Auth"`
    
    // BuildSpec AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-source.html#cfn-codebuild-project-source-buildspec
    BuildSpec string `json:"BuildSpec"`
    
    // Location AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-source.html#cfn-codebuild-project-source-location
    Location string `json:"Location"`
    
    // Type AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-source.html#cfn-codebuild-project-source-type
    Type string `json:"Type"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodeBuildProjectSource) AWSCloudFormationType() string {
    return "AWS::CodeBuild::Project.Source"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCodeBuildProjectSource) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}