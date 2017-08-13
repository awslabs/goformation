package resources

// AWS::ECS::Service.PlacementConstraint AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-placementconstraint.html
type AWSECSServicePlacementConstraint struct {
    
    // Expression AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-placementconstraint.html#cfn-ecs-service-placementconstraint-expression
    Expression string `json:"Expression"`
    
    // Type AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-placementconstraint.html#cfn-ecs-service-placementconstraint-type
    Type string `json:"Type"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSECSServicePlacementConstraint) AWSCloudFormationType() string {
    return "AWS::ECS::Service.PlacementConstraint"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSECSServicePlacementConstraint) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}