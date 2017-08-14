package resources

// AWS::CodePipeline::CustomActionType.Settings AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-settings.html
type AWSCodePipelineCustomActionType_Settings struct {

	// EntityUrlTemplate AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-settings.html#cfn-codepipeline-customactiontype-settings-entityurltemplate

	EntityUrlTemplate string `json:"EntityUrlTemplate"`

	// ExecutionUrlTemplate AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-settings.html#cfn-codepipeline-customactiontype-settings-executionurltemplate

	ExecutionUrlTemplate string `json:"ExecutionUrlTemplate"`

	// RevisionUrlTemplate AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-settings.html#cfn-codepipeline-customactiontype-settings-revisionurltemplate

	RevisionUrlTemplate string `json:"RevisionUrlTemplate"`

	// ThirdPartyConfigurationUrl AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-settings.html#cfn-codepipeline-customactiontype-settings-thirdpartyconfigurationurl

	ThirdPartyConfigurationUrl string `json:"ThirdPartyConfigurationUrl"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodePipelineCustomActionType_Settings) AWSCloudFormationType() string {
	return "AWS::CodePipeline::CustomActionType.Settings"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCodePipelineCustomActionType_Settings) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
