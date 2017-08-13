package resources

// AWS::DynamoDB::Table.Projection AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-projectionobject.html
type AWSDynamoDBTableProjection struct {
    
    // NonKeyAttributes AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-projectionobject.html#cfn-dynamodb-projectionobj-nonkeyatt
    NonKeyAttributes []string `json:"NonKeyAttributes"`
    
    // ProjectionType AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-projectionobject.html#cfn-dynamodb-projectionobj-projtype
    ProjectionType string `json:"ProjectionType"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSDynamoDBTableProjection) AWSCloudFormationType() string {
    return "AWS::DynamoDB::Table.Projection"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSDynamoDBTableProjection) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}