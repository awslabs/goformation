package resources

import "github.com/awslabs/goformation/cloudformation/policies"

// AWSGlueTrigger_Action AWS CloudFormation Resource (AWS::Glue::Trigger.Action)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-trigger-action.html
type AWSGlueTrigger_Action struct {

	// Arguments AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-trigger-action.html#cfn-glue-trigger-action-arguments
	Arguments interface{} `json:"Arguments,omitempty"`

	// CrawlerName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-trigger-action.html#cfn-glue-trigger-action-crawlername
	CrawlerName string `json:"CrawlerName,omitempty"`

	// JobName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-trigger-action.html#cfn-glue-trigger-action-jobname
	JobName string `json:"JobName,omitempty"`

	// NotificationProperty AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-trigger-action.html#cfn-glue-trigger-action-notificationproperty
	NotificationProperty *AWSGlueTrigger_NotificationProperty `json:"NotificationProperty,omitempty"`

	// SecurityConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-trigger-action.html#cfn-glue-trigger-action-securityconfiguration
	SecurityConfiguration string `json:"SecurityConfiguration,omitempty"`

	// Timeout AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-trigger-action.html#cfn-glue-trigger-action-timeout
	Timeout int `json:"Timeout,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy policies.DeletionPolicy

	// _dependsOn stores the logical ID of the resources to be created before this resource
	_dependsOn []string

	// _metadata stores structured data associated with this resource
	_metadata map[string]interface{}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSGlueTrigger_Action) AWSCloudFormationType() string {
	return "AWS::Glue::Trigger.Action"
}

// DependsOn returns a slice of logical ID names this resource depends on.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSGlueTrigger_Action) DependsOn() []string {
	return r._dependsOn
}

// SetDependsOn specify that the creation of this resource follows another.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSGlueTrigger_Action) SetDependsOn(dependencies []string) {
	r._dependsOn = dependencies
}

// Metadata returns the metadata associated with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSGlueTrigger_Action) Metadata() map[string]interface{} {
	return r._metadata
}

// SetMetadata enables you to associate structured data with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSGlueTrigger_Action) SetMetadata(metadata map[string]interface{}) {
	r._metadata = metadata
}

// DeletionPolicy returns the AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSGlueTrigger_Action) DeletionPolicy() policies.DeletionPolicy {
	return r._deletionPolicy
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSGlueTrigger_Action) SetDeletionPolicy(policy policies.DeletionPolicy) {
	r._deletionPolicy = policy
}
