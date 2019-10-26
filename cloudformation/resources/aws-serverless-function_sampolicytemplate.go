package resources

import "github.com/awslabs/goformation/v2/cloudformation/policies"

// AWSServerlessFunction_SAMPolicyTemplate AWS CloudFormation Resource (AWS::Serverless::Function.SAMPolicyTemplate)
// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
type AWSServerlessFunction_SAMPolicyTemplate struct {

	// AMIDescribePolicy AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	AMIDescribePolicy *AWSServerlessFunction_EmptySAMPT `json:"AMIDescribePolicy,omitempty"`

	// CloudFormationDescribeStacksPolicy AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	CloudFormationDescribeStacksPolicy *AWSServerlessFunction_EmptySAMPT `json:"CloudFormationDescribeStacksPolicy,omitempty"`

	// CloudWatchPutMetricPolicy AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	CloudWatchPutMetricPolicy *AWSServerlessFunction_EmptySAMPT `json:"CloudWatchPutMetricPolicy,omitempty"`

	// DynamoDBCrudPolicy AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	DynamoDBCrudPolicy *AWSServerlessFunction_TableSAMPT `json:"DynamoDBCrudPolicy,omitempty"`

	// DynamoDBReadPolicy AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	DynamoDBReadPolicy *AWSServerlessFunction_TableSAMPT `json:"DynamoDBReadPolicy,omitempty"`

	// DynamoDBStreamReadPolicy AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	DynamoDBStreamReadPolicy *AWSServerlessFunction_TableStreamSAMPT `json:"DynamoDBStreamReadPolicy,omitempty"`

	// EC2DescribePolicy AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	EC2DescribePolicy *AWSServerlessFunction_EmptySAMPT `json:"EC2DescribePolicy,omitempty"`

	// ElasticsearchHttpPostPolicy AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	ElasticsearchHttpPostPolicy *AWSServerlessFunction_DomainSAMPT `json:"ElasticsearchHttpPostPolicy,omitempty"`

	// FilterLogEventsPolicy AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	FilterLogEventsPolicy *AWSServerlessFunction_LogGroupSAMPT `json:"FilterLogEventsPolicy,omitempty"`

	// KMSDecryptPolicy AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	KMSDecryptPolicy *AWSServerlessFunction_KeySAMPT `json:"KMSDecryptPolicy,omitempty"`

	// KinesisCrudPolicy AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	KinesisCrudPolicy *AWSServerlessFunction_StreamSAMPT `json:"KinesisCrudPolicy,omitempty"`

	// KinesisStreamReadPolicy AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	KinesisStreamReadPolicy *AWSServerlessFunction_StreamSAMPT `json:"KinesisStreamReadPolicy,omitempty"`

	// LambdaInvokePolicy AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	LambdaInvokePolicy *AWSServerlessFunction_FunctionSAMPT `json:"LambdaInvokePolicy,omitempty"`

	// RekognitionDetectOnlyPolicy AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	RekognitionDetectOnlyPolicy *AWSServerlessFunction_EmptySAMPT `json:"RekognitionDetectOnlyPolicy,omitempty"`

	// RekognitionLabelsPolicy AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	RekognitionLabelsPolicy *AWSServerlessFunction_EmptySAMPT `json:"RekognitionLabelsPolicy,omitempty"`

	// RekognitionNoDataAccessPolicy AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	RekognitionNoDataAccessPolicy *AWSServerlessFunction_CollectionSAMPT `json:"RekognitionNoDataAccessPolicy,omitempty"`

	// RekognitionReadPolicy AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	RekognitionReadPolicy *AWSServerlessFunction_CollectionSAMPT `json:"RekognitionReadPolicy,omitempty"`

	// RekognitionWriteOnlyAccessPolicy AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	RekognitionWriteOnlyAccessPolicy *AWSServerlessFunction_CollectionSAMPT `json:"RekognitionWriteOnlyAccessPolicy,omitempty"`

	// S3CrudPolicy AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	S3CrudPolicy *AWSServerlessFunction_BucketSAMPT `json:"S3CrudPolicy,omitempty"`

	// S3ReadPolicy AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	S3ReadPolicy *AWSServerlessFunction_BucketSAMPT `json:"S3ReadPolicy,omitempty"`

	// SESBulkTemplatedCrudPolicy AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	SESBulkTemplatedCrudPolicy *AWSServerlessFunction_IdentitySAMPT `json:"SESBulkTemplatedCrudPolicy,omitempty"`

	// SESCrudPolicy AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	SESCrudPolicy *AWSServerlessFunction_IdentitySAMPT `json:"SESCrudPolicy,omitempty"`

	// SESEmailTemplateCrudPolicy AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	SESEmailTemplateCrudPolicy *AWSServerlessFunction_EmptySAMPT `json:"SESEmailTemplateCrudPolicy,omitempty"`

	// SESSendBouncePolicy AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	SESSendBouncePolicy *AWSServerlessFunction_IdentitySAMPT `json:"SESSendBouncePolicy,omitempty"`

	// SNSCrudPolicy AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	SNSCrudPolicy *AWSServerlessFunction_TopicSAMPT `json:"SNSCrudPolicy,omitempty"`

	// SNSPublishMessagePolicy AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	SNSPublishMessagePolicy *AWSServerlessFunction_TopicSAMPT `json:"SNSPublishMessagePolicy,omitempty"`

	// SQSPollerPolicy AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	SQSPollerPolicy *AWSServerlessFunction_QueueSAMPT `json:"SQSPollerPolicy,omitempty"`

	// SQSSendMessagePolicy AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	SQSSendMessagePolicy *AWSServerlessFunction_QueueSAMPT `json:"SQSSendMessagePolicy,omitempty"`

	// StepFunctionsExecutionPolicy AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	StepFunctionsExecutionPolicy *AWSServerlessFunction_StateMachineSAMPT `json:"StepFunctionsExecutionPolicy,omitempty"`

	// VPCAccessPolicy AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	VPCAccessPolicy *AWSServerlessFunction_EmptySAMPT `json:"VPCAccessPolicy,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy policies.DeletionPolicy

	// _dependsOn stores the logical ID of the resources to be created before this resource
	_dependsOn []string

	// _metadata stores structured data associated with this resource
	_metadata map[string]interface{}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSServerlessFunction_SAMPolicyTemplate) AWSCloudFormationType() string {
	return "AWS::Serverless::Function.SAMPolicyTemplate"
}

// DependsOn returns a slice of logical ID names this resource depends on.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSServerlessFunction_SAMPolicyTemplate) DependsOn() []string {
	return r._dependsOn
}

// SetDependsOn specify that the creation of this resource follows another.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSServerlessFunction_SAMPolicyTemplate) SetDependsOn(dependencies []string) {
	r._dependsOn = dependencies
}

// Metadata returns the metadata associated with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSServerlessFunction_SAMPolicyTemplate) Metadata() map[string]interface{} {
	return r._metadata
}

// SetMetadata enables you to associate structured data with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSServerlessFunction_SAMPolicyTemplate) SetMetadata(metadata map[string]interface{}) {
	r._metadata = metadata
}

// DeletionPolicy returns the AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSServerlessFunction_SAMPolicyTemplate) DeletionPolicy() policies.DeletionPolicy {
	return r._deletionPolicy
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSServerlessFunction_SAMPolicyTemplate) SetDeletionPolicy(policy policies.DeletionPolicy) {
	r._deletionPolicy = policy
}
