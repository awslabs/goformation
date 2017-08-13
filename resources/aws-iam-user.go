package resources

// AWS::IAM::User AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-user.html
type AWSIAMUser struct {
    
    // Groups AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-user.html#cfn-iam-user-groups
    Groups []string `json:"Groups"`
    
    // LoginProfile AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-user.html#cfn-iam-user-loginprofile
    LoginProfile AWSIAMUserLoginProfile `json:"LoginProfile"`
    
    // ManagedPolicyArns AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-user.html#cfn-iam-user-managepolicyarns
    ManagedPolicyArns []string `json:"ManagedPolicyArns"`
    
    // Path AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-user.html#cfn-iam-user-path
    Path string `json:"Path"`
    
    // Policies AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-user.html#cfn-iam-user-policies
    Policies []AWSIAMUserPolicy `json:"Policies"`
    
    // UserName AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-user.html#cfn-iam-user-username
    UserName string `json:"UserName"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIAMUser) AWSCloudFormationType() string {
    return "AWS::IAM::User"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSIAMUser) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}