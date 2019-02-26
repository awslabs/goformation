package cloudformation

// AWSEC2LaunchTemplate_CapacityReservationTarget AWS CloudFormation Resource (AWS::EC2::LaunchTemplate.CapacityReservationTarget)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-capacityreservationtarget.html
type AWSEC2LaunchTemplate_CapacityReservationTarget struct {

	// CapacityReservationId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-capacityreservationtarget.html#cfn-ec2-launchtemplate-capacityreservationtarget-capacityreservationid
	CapacityReservationId string `json:"CapacityReservationId,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2LaunchTemplate_CapacityReservationTarget) AWSCloudFormationType() string {
	return "AWS::EC2::LaunchTemplate.CapacityReservationTarget"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSEC2LaunchTemplate_CapacityReservationTarget) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
