package cloudformation

// AWSWorkSpacesWorkspace_WorkspaceProperties AWS CloudFormation Resource (AWS::WorkSpaces::Workspace.WorkspaceProperties)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-workspace-workspaceproperties.html
type AWSWorkSpacesWorkspace_WorkspaceProperties struct {

	// ComputeTypeName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-workspace-workspaceproperties.html#cfn-workspaces-workspace-workspaceproperties-computetypename
	ComputeTypeName string `json:"ComputeTypeName,omitempty"`

	// RootVolumeSizeGib AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-workspace-workspaceproperties.html#cfn-workspaces-workspace-workspaceproperties-rootvolumesizegib
	RootVolumeSizeGib int `json:"RootVolumeSizeGib,omitempty"`

	// RunningMode AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-workspace-workspaceproperties.html#cfn-workspaces-workspace-workspaceproperties-runningmode
	RunningMode string `json:"RunningMode,omitempty"`

	// RunningModeAutoStopTimeoutInMinutes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-workspace-workspaceproperties.html#cfn-workspaces-workspace-workspaceproperties-runningmodeautostoptimeoutinminutes
	RunningModeAutoStopTimeoutInMinutes int `json:"RunningModeAutoStopTimeoutInMinutes,omitempty"`

	// UserVolumeSizeGib AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-workspace-workspaceproperties.html#cfn-workspaces-workspace-workspaceproperties-uservolumesizegib
	UserVolumeSizeGib int `json:"UserVolumeSizeGib,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWorkSpacesWorkspace_WorkspaceProperties) AWSCloudFormationType() string {
	return "AWS::WorkSpaces::Workspace.WorkspaceProperties"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSWorkSpacesWorkspace_WorkspaceProperties) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
