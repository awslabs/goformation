package serverless

import (
	"github.com/awslabs/goformation/v3/cloudformation/policies"
)

// Function_SAMPolicyTemplate AWS CloudFormation Resource (AWS::Serverless::Function.SAMPolicyTemplate)
// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
type Function_SAMPolicyTemplate struct {

	// AMIDescribePolicy AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	AMIDescribePolicy *Function_EmptySAMPT `json:"AMIDescribePolicy,omitempty"`

	// CloudFormationDescribeStacksPolicy AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	CloudFormationDescribeStacksPolicy *Function_EmptySAMPT `json:"CloudFormationDescribeStacksPolicy,omitempty"`

	// CloudWatchPutMetricPolicy AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	CloudWatchPutMetricPolicy *Function_EmptySAMPT `json:"CloudWatchPutMetricPolicy,omitempty"`

	// DynamoDBCrudPolicy AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	DynamoDBCrudPolicy *Function_TableSAMPT `json:"DynamoDBCrudPolicy,omitempty"`

	// DynamoDBReadPolicy AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	DynamoDBReadPolicy *Function_TableSAMPT `json:"DynamoDBReadPolicy,omitempty"`

	// DynamoDBStreamReadPolicy AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	DynamoDBStreamReadPolicy *Function_TableStreamSAMPT `json:"DynamoDBStreamReadPolicy,omitempty"`

	// EC2DescribePolicy AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	EC2DescribePolicy *Function_EmptySAMPT `json:"EC2DescribePolicy,omitempty"`

	// ElasticsearchHttpPostPolicy AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	ElasticsearchHttpPostPolicy *Function_DomainSAMPT `json:"ElasticsearchHttpPostPolicy,omitempty"`

	// FilterLogEventsPolicy AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	FilterLogEventsPolicy *Function_LogGroupSAMPT `json:"FilterLogEventsPolicy,omitempty"`

	// KMSDecryptPolicy AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	KMSDecryptPolicy *Function_KeySAMPT `json:"KMSDecryptPolicy,omitempty"`

	// KinesisCrudPolicy AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	KinesisCrudPolicy *Function_StreamSAMPT `json:"KinesisCrudPolicy,omitempty"`

	// KinesisStreamReadPolicy AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	KinesisStreamReadPolicy *Function_StreamSAMPT `json:"KinesisStreamReadPolicy,omitempty"`

	// LambdaInvokePolicy AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	LambdaInvokePolicy *Function_FunctionSAMPT `json:"LambdaInvokePolicy,omitempty"`

	// RekognitionDetectOnlyPolicy AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	RekognitionDetectOnlyPolicy *Function_EmptySAMPT `json:"RekognitionDetectOnlyPolicy,omitempty"`

	// RekognitionLabelsPolicy AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	RekognitionLabelsPolicy *Function_EmptySAMPT `json:"RekognitionLabelsPolicy,omitempty"`

	// RekognitionNoDataAccessPolicy AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	RekognitionNoDataAccessPolicy *Function_CollectionSAMPT `json:"RekognitionNoDataAccessPolicy,omitempty"`

	// RekognitionReadPolicy AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	RekognitionReadPolicy *Function_CollectionSAMPT `json:"RekognitionReadPolicy,omitempty"`

	// RekognitionWriteOnlyAccessPolicy AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	RekognitionWriteOnlyAccessPolicy *Function_CollectionSAMPT `json:"RekognitionWriteOnlyAccessPolicy,omitempty"`

	// S3CrudPolicy AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	S3CrudPolicy *Function_BucketSAMPT `json:"S3CrudPolicy,omitempty"`

	// S3ReadPolicy AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	S3ReadPolicy *Function_BucketSAMPT `json:"S3ReadPolicy,omitempty"`

	// SESBulkTemplatedCrudPolicy AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	SESBulkTemplatedCrudPolicy *Function_IdentitySAMPT `json:"SESBulkTemplatedCrudPolicy,omitempty"`

	// SESCrudPolicy AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	SESCrudPolicy *Function_IdentitySAMPT `json:"SESCrudPolicy,omitempty"`

	// SESEmailTemplateCrudPolicy AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	SESEmailTemplateCrudPolicy *Function_EmptySAMPT `json:"SESEmailTemplateCrudPolicy,omitempty"`

	// SESSendBouncePolicy AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	SESSendBouncePolicy *Function_IdentitySAMPT `json:"SESSendBouncePolicy,omitempty"`

	// SNSCrudPolicy AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	SNSCrudPolicy *Function_TopicSAMPT `json:"SNSCrudPolicy,omitempty"`

	// SNSPublishMessagePolicy AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	SNSPublishMessagePolicy *Function_TopicSAMPT `json:"SNSPublishMessagePolicy,omitempty"`

	// SQSPollerPolicy AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	SQSPollerPolicy *Function_QueueSAMPT `json:"SQSPollerPolicy,omitempty"`

	// SQSSendMessagePolicy AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	SQSSendMessagePolicy *Function_QueueSAMPT `json:"SQSSendMessagePolicy,omitempty"`

	// StepFunctionsExecutionPolicy AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	StepFunctionsExecutionPolicy *Function_StateMachineSAMPT `json:"StepFunctionsExecutionPolicy,omitempty"`

	// VPCAccessPolicy AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	VPCAccessPolicy *Function_EmptySAMPT `json:"VPCAccessPolicy,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy policies.DeletionPolicy

	// _dependsOn stores the logical ID of the resources to be created before this resource
	_dependsOn []string

	// _metadata stores structured data associated with this resource
	_metadata map[string]interface{}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *Function_SAMPolicyTemplate) AWSCloudFormationType() string {
	return "AWS::Serverless::Function.SAMPolicyTemplate"
}

// DependsOn returns a slice of logical ID names this resource depends on.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *Function_SAMPolicyTemplate) DependsOn() []string {
	return r._dependsOn
}

// SetDependsOn specify that the creation of this resource follows another.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *Function_SAMPolicyTemplate) SetDependsOn(dependencies []string) {
	r._dependsOn = dependencies
}

// Metadata returns the metadata associated with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *Function_SAMPolicyTemplate) Metadata() map[string]interface{} {
	return r._metadata
}

// SetMetadata enables you to associate structured data with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *Function_SAMPolicyTemplate) SetMetadata(metadata map[string]interface{}) {
	r._metadata = metadata
}

// DeletionPolicy returns the AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *Function_SAMPolicyTemplate) DeletionPolicy() policies.DeletionPolicy {
	return r._deletionPolicy
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *Function_SAMPolicyTemplate) SetDeletionPolicy(policy policies.DeletionPolicy) {
	r._deletionPolicy = policy
}
