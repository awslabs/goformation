package lightsail

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/awslabs/goformation/v5/cloudformation/policies"
	"github.com/awslabs/goformation/v5/cloudformation/tags"
)

// Distribution AWS CloudFormation Resource (AWS::Lightsail::Distribution)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-distribution.html
type Distribution struct {

	// BundleId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-distribution.html#cfn-lightsail-distribution-bundleid
	BundleId string `json:"BundleId,omitempty"`

	// CacheBehaviorSettings AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-distribution.html#cfn-lightsail-distribution-cachebehaviorsettings
	CacheBehaviorSettings *Distribution_CacheSettings `json:"CacheBehaviorSettings,omitempty"`

	// CacheBehaviors AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-distribution.html#cfn-lightsail-distribution-cachebehaviors
	CacheBehaviors []Distribution_CacheBehaviorPerPath `json:"CacheBehaviors,omitempty"`

	// CertificateName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-distribution.html#cfn-lightsail-distribution-certificatename
	CertificateName string `json:"CertificateName,omitempty"`

	// DefaultCacheBehavior AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-distribution.html#cfn-lightsail-distribution-defaultcachebehavior
	DefaultCacheBehavior *Distribution_CacheBehavior `json:"DefaultCacheBehavior,omitempty"`

	// DistributionName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-distribution.html#cfn-lightsail-distribution-distributionname
	DistributionName string `json:"DistributionName,omitempty"`

	// IpAddressType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-distribution.html#cfn-lightsail-distribution-ipaddresstype
	IpAddressType string `json:"IpAddressType,omitempty"`

	// IsEnabled AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-distribution.html#cfn-lightsail-distribution-isenabled
	IsEnabled bool `json:"IsEnabled,omitempty"`

	// Origin AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-distribution.html#cfn-lightsail-distribution-origin
	Origin *Distribution_InputOrigin `json:"Origin,omitempty"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-distribution.html#cfn-lightsail-distribution-tags
	Tags []tags.Tag `json:"Tags,omitempty"`

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
func (r *Distribution) AWSCloudFormationType() string {
	return "AWS::Lightsail::Distribution"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r Distribution) MarshalJSON() ([]byte, error) {
	type Properties Distribution
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
func (r *Distribution) UnmarshalJSON(b []byte) error {
	type Properties Distribution
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
		*r = Distribution(*res.Properties)
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
