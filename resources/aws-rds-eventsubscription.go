package resources

// AWS::RDS::EventSubscription AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-eventsubscription.html
type AWSRDSEventSubscription struct {
    
    // Enabled AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-eventsubscription.html#cfn-rds-eventsubscription-enabled
    Enabled bool `json:"Enabled"`
    
    // EventCategories AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-eventsubscription.html#cfn-rds-eventsubscription-eventcategories
    EventCategories []string `json:"EventCategories"`
    
    // SnsTopicArn AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-eventsubscription.html#cfn-rds-eventsubscription-snstopicarn
    SnsTopicArn string `json:"SnsTopicArn"`
    
    // SourceIds AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-eventsubscription.html#cfn-rds-eventsubscription-sourceids
    SourceIds []string `json:"SourceIds"`
    
    // SourceType AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-eventsubscription.html#cfn-rds-eventsubscription-sourcetype
    SourceType string `json:"SourceType"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSRDSEventSubscription) AWSCloudFormationType() string {
    return "AWS::RDS::EventSubscription"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSRDSEventSubscription) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}