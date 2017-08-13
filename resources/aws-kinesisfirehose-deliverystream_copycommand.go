package resources

// AWS::KinesisFirehose::DeliveryStream.CopyCommand AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-redshiftdestinationconfiguration-copycommand.html
type AWSKinesisFirehoseDeliveryStreamCopyCommand struct {
    
    // CopyOptions AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-redshiftdestinationconfiguration-copycommand.html#cfn-kinesisfirehose-kinesisdeliverystream-redshiftdestinationconfiguration-copycommand-copyoptions
    CopyOptions string `json:"CopyOptions"`
    
    // DataTableColumns AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-redshiftdestinationconfiguration-copycommand.html#cfn-kinesisfirehose-kinesisdeliverystream-redshiftdestinationconfiguration-copycommand-datatablecolumns
    DataTableColumns string `json:"DataTableColumns"`
    
    // DataTableName AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-redshiftdestinationconfiguration-copycommand.html#cfn-kinesisfirehose-kinesisdeliverystream-redshiftdestinationconfiguration-copycommand-datatablename
    DataTableName string `json:"DataTableName"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisFirehoseDeliveryStreamCopyCommand) AWSCloudFormationType() string {
    return "AWS::KinesisFirehose::DeliveryStream.CopyCommand"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKinesisFirehoseDeliveryStreamCopyCommand) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}