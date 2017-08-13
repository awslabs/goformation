package resources

// AWS::GameLift::Alias AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-alias.html
type AWSGameLiftAlias struct {
    
    // Description AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-alias.html#cfn-gamelift-alias-description
    Description string `json:"Description"`
    
    // Name AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-alias.html#cfn-gamelift-alias-name
    Name string `json:"Name"`
    
    // RoutingStrategy AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-alias.html#cfn-gamelift-alias-routingstrategy
    RoutingStrategy AWSGameLiftAliasRoutingStrategy `json:"RoutingStrategy"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSGameLiftAlias) AWSCloudFormationType() string {
    return "AWS::GameLift::Alias"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSGameLiftAlias) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}