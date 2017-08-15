package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::WorkSpaces::Workspace AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-workspace.html
type AWSWorkSpacesWorkspace struct {

	// BundleId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-workspace.html#cfn-workspaces-workspace-bundleid

	BundleId string `json:"BundleId"`

	// DirectoryId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-workspace.html#cfn-workspaces-workspace-directoryid

	DirectoryId string `json:"DirectoryId"`

	// RootVolumeEncryptionEnabled AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-workspace.html#cfn-workspaces-workspace-rootvolumeencryptionenabled

	RootVolumeEncryptionEnabled bool `json:"RootVolumeEncryptionEnabled"`

	// UserName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-workspace.html#cfn-workspaces-workspace-username

	UserName string `json:"UserName"`

	// UserVolumeEncryptionEnabled AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-workspace.html#cfn-workspaces-workspace-uservolumeencryptionenabled

	UserVolumeEncryptionEnabled bool `json:"UserVolumeEncryptionEnabled"`

	// VolumeEncryptionKey AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-workspace.html#cfn-workspaces-workspace-volumeencryptionkey

	VolumeEncryptionKey string `json:"VolumeEncryptionKey"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWorkSpacesWorkspace) AWSCloudFormationType() string {
	return "AWS::WorkSpaces::Workspace"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWorkSpacesWorkspace) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSWorkSpacesWorkspaceResources retrieves all AWSWorkSpacesWorkspace items from a CloudFormation template
func GetAllAWSWorkSpacesWorkspace(template *Template) map[string]*AWSWorkSpacesWorkspace {

	results := map[string]*AWSWorkSpacesWorkspace{}
	for name, resource := range template.Resources {
		result := &AWSWorkSpacesWorkspace{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSWorkSpacesWorkspaceWithName retrieves all AWSWorkSpacesWorkspace items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSWorkSpacesWorkspace(name string, template *Template) (*AWSWorkSpacesWorkspace, error) {

	result := &AWSWorkSpacesWorkspace{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSWorkSpacesWorkspace{}, errors.New("resource not found")

}
