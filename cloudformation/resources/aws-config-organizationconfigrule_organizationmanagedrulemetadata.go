package resources

import "github.com/awslabs/goformation/cloudformation/policies"

// AWSConfigOrganizationConfigRule_OrganizationManagedRuleMetadata AWS CloudFormation Resource (AWS::Config::OrganizationConfigRule.OrganizationManagedRuleMetadata)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-organizationconfigrule-organizationmanagedrulemetadata.html
type AWSConfigOrganizationConfigRule_OrganizationManagedRuleMetadata struct {

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-organizationconfigrule-organizationmanagedrulemetadata.html#cfn-config-organizationconfigrule-organizationmanagedrulemetadata-description
	Description string `json:"Description,omitempty"`

	// InputParameters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-organizationconfigrule-organizationmanagedrulemetadata.html#cfn-config-organizationconfigrule-organizationmanagedrulemetadata-inputparameters
	InputParameters string `json:"InputParameters,omitempty"`

	// MaximumExecutionFrequency AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-organizationconfigrule-organizationmanagedrulemetadata.html#cfn-config-organizationconfigrule-organizationmanagedrulemetadata-maximumexecutionfrequency
	MaximumExecutionFrequency string `json:"MaximumExecutionFrequency,omitempty"`

	// ResourceIdScope AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-organizationconfigrule-organizationmanagedrulemetadata.html#cfn-config-organizationconfigrule-organizationmanagedrulemetadata-resourceidscope
	ResourceIdScope string `json:"ResourceIdScope,omitempty"`

	// ResourceTypesScope AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-organizationconfigrule-organizationmanagedrulemetadata.html#cfn-config-organizationconfigrule-organizationmanagedrulemetadata-resourcetypesscope
	ResourceTypesScope []string `json:"ResourceTypesScope,omitempty"`

	// RuleIdentifier AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-organizationconfigrule-organizationmanagedrulemetadata.html#cfn-config-organizationconfigrule-organizationmanagedrulemetadata-ruleidentifier
	RuleIdentifier string `json:"RuleIdentifier,omitempty"`

	// TagKeyScope AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-organizationconfigrule-organizationmanagedrulemetadata.html#cfn-config-organizationconfigrule-organizationmanagedrulemetadata-tagkeyscope
	TagKeyScope string `json:"TagKeyScope,omitempty"`

	// TagValueScope AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-organizationconfigrule-organizationmanagedrulemetadata.html#cfn-config-organizationconfigrule-organizationmanagedrulemetadata-tagvaluescope
	TagValueScope string `json:"TagValueScope,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy policies.DeletionPolicy

	// _dependsOn stores the logical ID of the resources to be created before this resource
	_dependsOn []string

	// _metadata stores structured data associated with this resource
	_metadata map[string]interface{}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSConfigOrganizationConfigRule_OrganizationManagedRuleMetadata) AWSCloudFormationType() string {
	return "AWS::Config::OrganizationConfigRule.OrganizationManagedRuleMetadata"
}

// DependsOn returns a slice of logical ID names this resource depends on.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSConfigOrganizationConfigRule_OrganizationManagedRuleMetadata) DependsOn() []string {
	return r._dependsOn
}

// SetDependsOn specify that the creation of this resource follows another.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSConfigOrganizationConfigRule_OrganizationManagedRuleMetadata) SetDependsOn(dependencies []string) {
	r._dependsOn = dependencies
}

// Metadata returns the metadata associated with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSConfigOrganizationConfigRule_OrganizationManagedRuleMetadata) Metadata() map[string]interface{} {
	return r._metadata
}

// SetMetadata enables you to associate structured data with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSConfigOrganizationConfigRule_OrganizationManagedRuleMetadata) SetMetadata(metadata map[string]interface{}) {
	r._metadata = metadata
}

// DeletionPolicy returns the AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSConfigOrganizationConfigRule_OrganizationManagedRuleMetadata) DeletionPolicy() policies.DeletionPolicy {
	return r._deletionPolicy
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSConfigOrganizationConfigRule_OrganizationManagedRuleMetadata) SetDeletionPolicy(policy policies.DeletionPolicy) {
	r._deletionPolicy = policy
}
