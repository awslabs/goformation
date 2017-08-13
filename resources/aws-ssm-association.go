package resources

// AWS::SSM::Association AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html
type AWSSSMAssociation struct {
    
    // DocumentVersion AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-documentversion
    DocumentVersion string `json:"DocumentVersion"`
    
    // InstanceId AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-instanceid
    InstanceId string `json:"InstanceId"`
    
    // Name AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-name
    Name string `json:"Name"`
    
    // Parameters AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-parameters
    Parameters map[string]AWSSSMAssociationParameterValues `json:"Parameters"`
    
    // ScheduleExpression AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-scheduleexpression
    ScheduleExpression string `json:"ScheduleExpression"`
    
    // Targets AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-targets
    Targets []AWSSSMAssociationTarget `json:"Targets"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSSSMAssociation) AWSCloudFormationType() string {
    return "AWS::SSM::Association"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSSSMAssociation) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}