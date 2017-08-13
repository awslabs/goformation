package resources

// AWS::CloudTrail::Trail AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail.html
type AWSCloudTrailTrail struct {
    
    // CloudWatchLogsLogGroupArn AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail.html#aws-cloudtrail-trail-cloudwatchlogsloggroupaem
    CloudWatchLogsLogGroupArn string `json:"CloudWatchLogsLogGroupArn"`
    
    // CloudWatchLogsRoleArn AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail.html#aws-cloudtrail-trail-cloudwatchlogsrolearn
    CloudWatchLogsRoleArn string `json:"CloudWatchLogsRoleArn"`
    
    // EnableLogFileValidation AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail.html#aws-cloudtrail-trail-enablelogfilevalidation
    EnableLogFileValidation bool `json:"EnableLogFileValidation"`
    
    // IncludeGlobalServiceEvents AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail.html#aws-cloudtrail-trail-includeglobalserviceevents
    IncludeGlobalServiceEvents bool `json:"IncludeGlobalServiceEvents"`
    
    // IsLogging AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail.html#aws-cloudtrail-trail-islogging
    IsLogging bool `json:"IsLogging"`
    
    // IsMultiRegionTrail AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail.html#aws-cloudtrail-trail-ismultiregiontrail
    IsMultiRegionTrail bool `json:"IsMultiRegionTrail"`
    
    // KMSKeyId AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail.html#aws-cloudtrail-trail-kmskeyid
    KMSKeyId string `json:"KMSKeyId"`
    
    // S3BucketName AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail.html#aws-cloudtrail-trail-s3bucketname
    S3BucketName string `json:"S3BucketName"`
    
    // S3KeyPrefix AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail.html#aws-cloudtrail-trail-s3keyprefix
    S3KeyPrefix string `json:"S3KeyPrefix"`
    
    // SnsTopicName AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail.html#aws-cloudtrail-trail-snstopicname
    SnsTopicName string `json:"SnsTopicName"`
    
    // Tags AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail.html#cfn-cloudtrail-trail-tags
    Tags []AWSCloudTrailTrailTag `json:"Tags"`
    
    // TrailName AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail.html#aws-cloudtrail-trail-trailname
    TrailName string `json:"TrailName"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCloudTrailTrail) AWSCloudFormationType() string {
    return "AWS::CloudTrail::Trail"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCloudTrailTrail) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}