package resources

// AWS::StepFunctions::StateMachine AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html
type AWSStepFunctionsStateMachine struct {
    
    // DefinitionString AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-definitionstring
    DefinitionString string `json:"DefinitionString"`
    
    // RoleArn AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-rolearn
    RoleArn string `json:"RoleArn"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSStepFunctionsStateMachine) AWSCloudFormationType() string {
    return "AWS::StepFunctions::StateMachine"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSStepFunctionsStateMachine) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}