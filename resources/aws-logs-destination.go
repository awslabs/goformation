package resources

// AWS::Logs::Destination AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-destination.html
type AWSLogsDestination struct {
    
    // DestinationName AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-destination.html#cfn-logs-destination-destinationname
    DestinationName string `json:"DestinationName"`
    
    // DestinationPolicy AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-destination.html#cfn-logs-destination-destinationpolicy
    DestinationPolicy string `json:"DestinationPolicy"`
    
    // RoleArn AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-destination.html#cfn-logs-destination-rolearn
    RoleArn string `json:"RoleArn"`
    
    // TargetArn AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-destination.html#cfn-logs-destination-targetarn
    TargetArn string `json:"TargetArn"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSLogsDestination) AWSCloudFormationType() string {
    return "AWS::Logs::Destination"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSLogsDestination) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}