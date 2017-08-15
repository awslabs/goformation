package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Config::ConfigRule.Scope AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-scope.html
type AWSConfigConfigRule_Scope struct {

	// ComplianceResourceId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-scope.html#cfn-config-configrule-scope-complianceresourceid
	ComplianceResourceId string `json:"ComplianceResourceId"`

	// ComplianceResourceTypes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-scope.html#cfn-config-configrule-scope-complianceresourcetypes
	ComplianceResourceTypes []string `json:"ComplianceResourceTypes"`

	// TagKey AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-scope.html#cfn-config-configrule-scope-tagkey
	TagKey string `json:"TagKey"`

	// TagValue AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-scope.html#cfn-config-configrule-scope-tagvalue
	TagValue string `json:"TagValue"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSConfigConfigRule_Scope) AWSCloudFormationType() string {
	return "AWS::Config::ConfigRule.Scope"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSConfigConfigRule_Scope) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
