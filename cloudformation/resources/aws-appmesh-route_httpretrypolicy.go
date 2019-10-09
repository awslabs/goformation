package resources

import "github.com/awslabs/goformation/cloudformation/policies"

// AWSAppMeshRoute_HttpRetryPolicy AWS CloudFormation Resource (AWS::AppMesh::Route.HttpRetryPolicy)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-httpretrypolicy.html
type AWSAppMeshRoute_HttpRetryPolicy struct {

	// HttpRetryEvents AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-httpretrypolicy.html#cfn-appmesh-route-httpretrypolicy-httpretryevents
	HttpRetryEvents []string `json:"HttpRetryEvents,omitempty"`

	// MaxRetries AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-httpretrypolicy.html#cfn-appmesh-route-httpretrypolicy-maxretries
	MaxRetries int `json:"MaxRetries"`

	// PerRetryTimeout AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-httpretrypolicy.html#cfn-appmesh-route-httpretrypolicy-perretrytimeout
	PerRetryTimeout *AWSAppMeshRoute_Duration `json:"PerRetryTimeout,omitempty"`

	// TcpRetryEvents AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-httpretrypolicy.html#cfn-appmesh-route-httpretrypolicy-tcpretryevents
	TcpRetryEvents []string `json:"TcpRetryEvents,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy policies.DeletionPolicy

	// _dependsOn stores the logical ID of the resources to be created before this resource
	_dependsOn []string

	// _metadata stores structured data associated with this resource
	_metadata map[string]interface{}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSAppMeshRoute_HttpRetryPolicy) AWSCloudFormationType() string {
	return "AWS::AppMesh::Route.HttpRetryPolicy"
}

// DependsOn returns a slice of logical ID names this resource depends on.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSAppMeshRoute_HttpRetryPolicy) DependsOn() []string {
	return r._dependsOn
}

// SetDependsOn specify that the creation of this resource follows another.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSAppMeshRoute_HttpRetryPolicy) SetDependsOn(dependencies []string) {
	r._dependsOn = dependencies
}

// Metadata returns the metadata associated with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSAppMeshRoute_HttpRetryPolicy) Metadata() map[string]interface{} {
	return r._metadata
}

// SetMetadata enables you to associate structured data with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSAppMeshRoute_HttpRetryPolicy) SetMetadata(metadata map[string]interface{}) {
	r._metadata = metadata
}

// DeletionPolicy returns the AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSAppMeshRoute_HttpRetryPolicy) DeletionPolicy() policies.DeletionPolicy {
	return r._deletionPolicy
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSAppMeshRoute_HttpRetryPolicy) SetDeletionPolicy(policy policies.DeletionPolicy) {
	r._deletionPolicy = policy
}
