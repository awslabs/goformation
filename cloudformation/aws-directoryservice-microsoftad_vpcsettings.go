package cloudformation

import (
	"encoding/json"
)

// AWSDirectoryServiceMicrosoftAD_VpcSettings AWS CloudFormation Resource (AWS::DirectoryService::MicrosoftAD.VpcSettings)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-directoryservice-microsoftad-vpcsettings.html
type AWSDirectoryServiceMicrosoftAD_VpcSettings struct {

	// SubnetIds AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-directoryservice-microsoftad-vpcsettings.html#cfn-directoryservice-microsoftad-vpcsettings-subnetids
	SubnetIds []*Value `json:"SubnetIds,omitempty"`

	// VpcId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-directoryservice-microsoftad-vpcsettings.html#cfn-directoryservice-microsoftad-vpcsettings-vpcid
	VpcId *Value `json:"VpcId,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSDirectoryServiceMicrosoftAD_VpcSettings) AWSCloudFormationType() string {
	return "AWS::DirectoryService::MicrosoftAD.VpcSettings"
}

func (r *AWSDirectoryServiceMicrosoftAD_VpcSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}
