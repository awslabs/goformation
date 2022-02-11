package networkmanager

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/awslabs/goformation/v5/cloudformation/policies"
	"github.com/awslabs/goformation/v5/cloudformation/tags"
)

// Device AWS CloudFormation Resource (AWS::NetworkManager::Device)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-device.html
type Device struct {

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-device.html#cfn-networkmanager-device-description
	Description *string `json:"Description,omitempty"`

	// GlobalNetworkId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-device.html#cfn-networkmanager-device-globalnetworkid
	GlobalNetworkId string `json:"GlobalNetworkId"`

	// Location AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-device.html#cfn-networkmanager-device-location
	Location *Device_Location `json:"Location,omitempty"`

	// Model AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-device.html#cfn-networkmanager-device-model
	Model *string `json:"Model,omitempty"`

	// SerialNumber AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-device.html#cfn-networkmanager-device-serialnumber
	SerialNumber *string `json:"SerialNumber,omitempty"`

	// SiteId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-device.html#cfn-networkmanager-device-siteid
	SiteId *string `json:"SiteId,omitempty"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-device.html#cfn-networkmanager-device-tags
	Tags *[]tags.Tag `json:"Tags,omitempty"`

	// Type AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-device.html#cfn-networkmanager-device-type
	Type *string `json:"Type,omitempty"`

	// Vendor AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-device.html#cfn-networkmanager-device-vendor
	Vendor *string `json:"Vendor,omitempty"`

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
func (r *Device) AWSCloudFormationType() string {
	return "AWS::NetworkManager::Device"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r Device) MarshalJSON() ([]byte, error) {
	type Properties Device
	return json.Marshal(&struct {
		Type                string
		Properties          Properties
		DependsOn           []string                     `json:"DependsOn,omitempty"`
		Metadata            map[string]interface{}       `json:"Metadata,omitempty"`
		DeletionPolicy      policies.DeletionPolicy      `json:"DeletionPolicy,omitempty"`
		UpdateReplacePolicy policies.UpdateReplacePolicy `json:"UpdateReplacePolicy,omitempty"`
		Condition           string                       `json:"Condition,omitempty"`
	}{
		Type:                r.AWSCloudFormationType(),
		Properties:          (Properties)(r),
		DependsOn:           r.AWSCloudFormationDependsOn,
		Metadata:            r.AWSCloudFormationMetadata,
		DeletionPolicy:      r.AWSCloudFormationDeletionPolicy,
		UpdateReplacePolicy: r.AWSCloudFormationUpdateReplacePolicy,
		Condition:           r.AWSCloudFormationCondition,
	})
}

// UnmarshalJSON is a custom JSON unmarshalling hook that strips the outer
// AWS CloudFormation resource object, and just keeps the 'Properties' field.
func (r *Device) UnmarshalJSON(b []byte) error {
	type Properties Device
	res := &struct {
		Type                string
		Properties          *Properties
		DependsOn           []string
		Metadata            map[string]interface{}
		DeletionPolicy      string
		UpdateReplacePolicy string
		Condition           string
	}{}

	dec := json.NewDecoder(bytes.NewReader(b))
	dec.DisallowUnknownFields() // Force error if unknown field is found

	if err := dec.Decode(&res); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return err
	}

	// If the resource has no Properties set, it could be nil
	if res.Properties != nil {
		*r = Device(*res.Properties)
	}
	if res.DependsOn != nil {
		r.AWSCloudFormationDependsOn = res.DependsOn
	}
	if res.Metadata != nil {
		r.AWSCloudFormationMetadata = res.Metadata
	}
	if res.DeletionPolicy != "" {
		r.AWSCloudFormationDeletionPolicy = policies.DeletionPolicy(res.DeletionPolicy)
	}
	if res.UpdateReplacePolicy != "" {
		r.AWSCloudFormationUpdateReplacePolicy = policies.UpdateReplacePolicy(res.UpdateReplacePolicy)
	}
	if res.Condition != "" {
		r.AWSCloudFormationCondition = res.Condition
	}
	return nil
}
