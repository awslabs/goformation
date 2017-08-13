package resources

// AWS::CodeBuild::Project.EnvironmentVariable AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-environmentvariable.html
type AWSCodeBuildProjectEnvironmentVariable struct {
    
    // Name AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-environmentvariable.html#cfn-codebuild-project-environmentvariable-name
    Name string `json:"Name"`
    
    // Value AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-environmentvariable.html#cfn-codebuild-project-environmentvariable-value
    Value string `json:"Value"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodeBuildProjectEnvironmentVariable) AWSCloudFormationType() string {
    return "AWS::CodeBuild::Project.EnvironmentVariable"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCodeBuildProjectEnvironmentVariable) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}