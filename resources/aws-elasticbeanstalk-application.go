package resources

// AWS::ElasticBeanstalk::Application AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk.html
type AWSElasticBeanstalkApplication struct {
    
    // ApplicationName AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk.html#cfn-elasticbeanstalk-application-name
    ApplicationName string `json:"ApplicationName"`
    
    // Description AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk.html#cfn-elasticbeanstalk-application-description
    Description string `json:"Description"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticBeanstalkApplication) AWSCloudFormationType() string {
    return "AWS::ElasticBeanstalk::Application"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSElasticBeanstalkApplication) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}