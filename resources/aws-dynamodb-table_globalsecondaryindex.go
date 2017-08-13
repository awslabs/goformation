package resources

// AWS::DynamoDB::Table.GlobalSecondaryIndex AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-gsi.html
type AWSDynamoDBTableGlobalSecondaryIndex struct {
    
    // IndexName AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-gsi.html#cfn-dynamodb-gsi-indexname
    IndexName string `json:"IndexName"`
    
    // KeySchema AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-gsi.html#cfn-dynamodb-gsi-keyschema
    KeySchema []AWSDynamoDBTableGlobalSecondaryIndexKeySchema `json:"KeySchema"`
    
    // Projection AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-gsi.html#cfn-dynamodb-gsi-projection
    Projection AWSDynamoDBTableGlobalSecondaryIndexProjection `json:"Projection"`
    
    // ProvisionedThroughput AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-gsi.html#cfn-dynamodb-gsi-provisionedthroughput
    ProvisionedThroughput AWSDynamoDBTableGlobalSecondaryIndexProvisionedThroughput `json:"ProvisionedThroughput"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSDynamoDBTableGlobalSecondaryIndex) AWSCloudFormationType() string {
    return "AWS::DynamoDB::Table.GlobalSecondaryIndex"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSDynamoDBTableGlobalSecondaryIndex) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}