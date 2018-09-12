package cloudformation

import (
	"encoding/json"
)

// AWSEC2SpotFleet_GroupIdentifier AWS CloudFormation Resource (AWS::EC2::SpotFleet.GroupIdentifier)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-securitygroups.html
type AWSEC2SpotFleet_GroupIdentifier struct {

	// GroupId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-securitygroups.html#cfn-ec2-spotfleet-groupidentifier-groupid
	GroupId *Value `json:"GroupId,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2SpotFleet_GroupIdentifier) AWSCloudFormationType() string {
	return "AWS::EC2::SpotFleet.GroupIdentifier"
}

func (r *AWSEC2SpotFleet_GroupIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}
