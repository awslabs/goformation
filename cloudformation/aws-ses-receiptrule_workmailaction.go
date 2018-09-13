package cloudformation

import (
	"encoding/json"
)

// AWSSESReceiptRule_WorkmailAction AWS CloudFormation Resource (AWS::SES::ReceiptRule.WorkmailAction)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-workmailaction.html
type AWSSESReceiptRule_WorkmailAction struct {

	// OrganizationArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-workmailaction.html#cfn-ses-receiptrule-workmailaction-organizationarn
	OrganizationArn *Value `json:"OrganizationArn,omitempty"`

	// TopicArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-workmailaction.html#cfn-ses-receiptrule-workmailaction-topicarn
	TopicArn *Value `json:"TopicArn,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSSESReceiptRule_WorkmailAction) AWSCloudFormationType() string {
	return "AWS::SES::ReceiptRule.WorkmailAction"
}

func (r *AWSSESReceiptRule_WorkmailAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}
