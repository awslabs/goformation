package resources

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/awslabs/goformation/cloudformation/policies"
)

// AWSCodeBuildProject AWS CloudFormation Resource (AWS::CodeBuild::Project)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html
type AWSCodeBuildProject struct {

	// Artifacts AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-artifacts
	Artifacts *AWSCodeBuildProject_Artifacts `json:"Artifacts,omitempty"`

	// BadgeEnabled AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-badgeenabled
	BadgeEnabled bool `json:"BadgeEnabled,omitempty"`

	// Cache AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-cache
	Cache *AWSCodeBuildProject_ProjectCache `json:"Cache,omitempty"`

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-description
	Description string `json:"Description,omitempty"`

	// EncryptionKey AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-encryptionkey
	EncryptionKey string `json:"EncryptionKey,omitempty"`

	// Environment AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-environment
	Environment *AWSCodeBuildProject_Environment `json:"Environment,omitempty"`

	// LogsConfig AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-logsconfig
	LogsConfig *AWSCodeBuildProject_LogsConfig `json:"LogsConfig,omitempty"`

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-name
	Name string `json:"Name,omitempty"`

	// QueuedTimeoutInMinutes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-queuedtimeoutinminutes
	QueuedTimeoutInMinutes int `json:"QueuedTimeoutInMinutes,omitempty"`

	// SecondaryArtifacts AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-secondaryartifacts
	SecondaryArtifacts []AWSCodeBuildProject_Artifacts `json:"SecondaryArtifacts,omitempty"`

	// SecondarySources AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-secondarysources
	SecondarySources []AWSCodeBuildProject_Source `json:"SecondarySources,omitempty"`

	// ServiceRole AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-servicerole
	ServiceRole string `json:"ServiceRole,omitempty"`

	// Source AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-source
	Source *AWSCodeBuildProject_Source `json:"Source,omitempty"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-tags
	Tags []Tag `json:"Tags,omitempty"`

	// TimeoutInMinutes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-timeoutinminutes
	TimeoutInMinutes int `json:"TimeoutInMinutes,omitempty"`

	// Triggers AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-triggers
	Triggers *AWSCodeBuildProject_ProjectTriggers `json:"Triggers,omitempty"`

	// VpcConfig AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-vpcconfig
	VpcConfig *AWSCodeBuildProject_VpcConfig `json:"VpcConfig,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy policies.DeletionPolicy

	// _dependsOn stores the logical ID of the resources to be created before this resource
	_dependsOn []string

	// _metadata stores structured data associated with this resource
	_metadata map[string]interface{}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodeBuildProject) AWSCloudFormationType() string {
	return "AWS::CodeBuild::Project"
}

// DependsOn returns a slice of logical ID names this resource depends on.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSCodeBuildProject) DependsOn() []string {
	return r._dependsOn
}

// SetDependsOn specify that the creation of this resource follows another.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSCodeBuildProject) SetDependsOn(dependencies []string) {
	r._dependsOn = dependencies
}

// Metadata returns the metadata associated with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSCodeBuildProject) Metadata() map[string]interface{} {
	return r._metadata
}

// SetMetadata enables you to associate structured data with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSCodeBuildProject) SetMetadata(metadata map[string]interface{}) {
	r._metadata = metadata
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSCodeBuildProject) SetDeletionPolicy(policy policies.DeletionPolicy) {
	r._deletionPolicy = policy
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r AWSCodeBuildProject) MarshalJSON() ([]byte, error) {
	type Properties AWSCodeBuildProject
	return json.Marshal(&struct {
		Type           string
		Properties     Properties
		DependsOn      []string                `json:"DependsOn,omitempty"`
		Metadata       map[string]interface{}  `json:"Metadata,omitempty"`
		DeletionPolicy policies.DeletionPolicy `json:"DeletionPolicy,omitempty"`
	}{
		Type:           r.AWSCloudFormationType(),
		Properties:     (Properties)(r),
		DependsOn:      r._dependsOn,
		Metadata:       r._metadata,
		DeletionPolicy: r._deletionPolicy,
	})
}

// UnmarshalJSON is a custom JSON unmarshalling hook that strips the outer
// AWS CloudFormation resource object, and just keeps the 'Properties' field.
func (r *AWSCodeBuildProject) UnmarshalJSON(b []byte) error {
	type Properties AWSCodeBuildProject
	res := &struct {
		Type       string
		Properties *Properties
		DependsOn  []string
		Metadata   map[string]interface{}
	}{}

	dec := json.NewDecoder(bytes.NewReader(b))
	dec.DisallowUnknownFields() // Force error if unknown field is found

	if err := dec.Decode(&res); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return err
	}

	// If the resource has no Properties set, it could be nil
	if res.Properties != nil {
		*r = AWSCodeBuildProject(*res.Properties)
	}
	if res.DependsOn != nil {
		r._dependsOn = res.DependsOn
	}
	if res.Metadata != nil {
		r._metadata = res.Metadata
	}

	return nil
}
