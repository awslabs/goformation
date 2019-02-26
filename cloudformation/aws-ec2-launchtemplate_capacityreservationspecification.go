package cloudformation

// AWSEC2LaunchTemplate_CapacityReservationSpecification AWS CloudFormation Resource (AWS::EC2::LaunchTemplate.CapacityReservationSpecification)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata-capacityreservationspecification.html
type AWSEC2LaunchTemplate_CapacityReservationSpecification struct {

	// CapacityReservationPreference AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata-capacityreservationspecification.html#cfn-ec2-launchtemplate-launchtemplatedata-capacityreservationspecification-capacityreservationpreference
	CapacityReservationPreference *AWSEC2LaunchTemplate_CapacityReservationPreference `json:"CapacityReservationPreference,omitempty"`

	// CapacityReservationTarget AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata-capacityreservationspecification.html#cfn-ec2-launchtemplate-launchtemplatedata-capacityreservationspecification-capacityreservationtarget
	CapacityReservationTarget *AWSEC2LaunchTemplate_CapacityReservationTarget `json:"CapacityReservationTarget,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2LaunchTemplate_CapacityReservationSpecification) AWSCloudFormationType() string {
	return "AWS::EC2::LaunchTemplate.CapacityReservationSpecification"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSEC2LaunchTemplate_CapacityReservationSpecification) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
