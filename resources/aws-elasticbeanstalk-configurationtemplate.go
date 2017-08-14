package resources

// AWS::ElasticBeanstalk::ConfigurationTemplate AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-beanstalk-configurationtemplate.html
type AWSElasticBeanstalkConfigurationTemplate struct {

	// ApplicationName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-beanstalk-configurationtemplate.html#cfn-elasticbeanstalk-configurationtemplate-applicationname

	ApplicationName string `json:"ApplicationName"`

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-beanstalk-configurationtemplate.html#cfn-elasticbeanstalk-configurationtemplate-description

	Description string `json:"Description"`

	// EnvironmentId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-beanstalk-configurationtemplate.html#cfn-elasticbeanstalk-configurationtemplate-environmentid

	EnvironmentId string `json:"EnvironmentId"`

	// OptionSettings AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-beanstalk-configurationtemplate.html#cfn-elasticbeanstalk-configurationtemplate-optionsettings

	OptionSettings []AWSElasticBeanstalkConfigurationTemplate_ConfigurationOptionSetting `json:"OptionSettings"`

	// SolutionStackName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-beanstalk-configurationtemplate.html#cfn-elasticbeanstalk-configurationtemplate-solutionstackname

	SolutionStackName string `json:"SolutionStackName"`

	// SourceConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-beanstalk-configurationtemplate.html#cfn-elasticbeanstalk-configurationtemplate-sourceconfiguration

	SourceConfiguration AWSElasticBeanstalkConfigurationTemplate_SourceConfiguration `json:"SourceConfiguration"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticBeanstalkConfigurationTemplate) AWSCloudFormationType() string {
	return "AWS::ElasticBeanstalk::ConfigurationTemplate"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSElasticBeanstalkConfigurationTemplate) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
