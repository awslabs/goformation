package groundstation

import (
	"github.com/awslabs/goformation/v5/cloudformation/policies"
)

// Config_UplinkEchoConfig AWS CloudFormation Resource (AWS::GroundStation::Config.UplinkEchoConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-uplinkechoconfig.html
type Config_UplinkEchoConfig struct {

	// AntennaUplinkConfigArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-uplinkechoconfig.html#cfn-groundstation-config-uplinkechoconfig-antennauplinkconfigarn
	AntennaUplinkConfigArn string `json:"AntennaUplinkConfigArn,omitempty"`

	// Enabled AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-uplinkechoconfig.html#cfn-groundstation-config-uplinkechoconfig-enabled
	Enabled bool `json:"Enabled,omitempty"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationUpdateReplacePolicy represents a CloudFormation UpdateReplacePolicy
	AWSCloudFormationUpdateReplacePolicy policies.UpdateReplacePolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`

	// AWSCloudFormationCondition stores the logical ID of the condition that must be satisfied for this resource to be created
	AWSCloudFormationCondition string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *Config_UplinkEchoConfig) AWSCloudFormationType() string {
	return "AWS::GroundStation::Config.UplinkEchoConfig"
}
