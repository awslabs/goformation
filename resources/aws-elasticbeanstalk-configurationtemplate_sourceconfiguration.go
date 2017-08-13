package resources

// AWS::ElasticBeanstalk::ConfigurationTemplate.SourceConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-configurationtemplate-sourceconfiguration.html
type AWSElasticBeanstalkConfigurationTemplateSourceConfiguration struct {
    
    // ApplicationName AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-configurationtemplate-sourceconfiguration.html#cfn-beanstalk-configurationtemplate-sourceconfiguration-applicationname
    ApplicationName string `json:"ApplicationName"`
    
    // TemplateName AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-configurationtemplate-sourceconfiguration.html#cfn-beanstalk-configurationtemplate-sourceconfiguration-templatename
    TemplateName string `json:"TemplateName"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticBeanstalkConfigurationTemplateSourceConfiguration) AWSCloudFormationType() string {
    return "AWS::ElasticBeanstalk::ConfigurationTemplate.SourceConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSElasticBeanstalkConfigurationTemplateSourceConfiguration) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}