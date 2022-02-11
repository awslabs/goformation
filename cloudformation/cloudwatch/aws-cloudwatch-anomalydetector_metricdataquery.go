package cloudwatch

import (
	"github.com/awslabs/goformation/v6/cloudformation/policies"
)

// AnomalyDetector_MetricDataQuery AWS CloudFormation Resource (AWS::CloudWatch::AnomalyDetector.MetricDataQuery)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-anomalydetector-metricdataquery.html
type AnomalyDetector_MetricDataQuery struct {

	// AccountId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-anomalydetector-metricdataquery.html#cfn-cloudwatch-anomalydetector-metricdataquery-accountid
	AccountId *string `json:"AccountId,omitempty"`

	// Expression AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-anomalydetector-metricdataquery.html#cfn-cloudwatch-anomalydetector-metricdataquery-expression
	Expression *string `json:"Expression,omitempty"`

	// Id AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-anomalydetector-metricdataquery.html#cfn-cloudwatch-anomalydetector-metricdataquery-id
	Id string `json:"Id"`

	// Label AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-anomalydetector-metricdataquery.html#cfn-cloudwatch-anomalydetector-metricdataquery-label
	Label *string `json:"Label,omitempty"`

	// MetricStat AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-anomalydetector-metricdataquery.html#cfn-cloudwatch-anomalydetector-metricdataquery-metricstat
	MetricStat *AnomalyDetector_MetricStat `json:"MetricStat,omitempty"`

	// Period AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-anomalydetector-metricdataquery.html#cfn-cloudwatch-anomalydetector-metricdataquery-period
	Period *int `json:"Period,omitempty"`

	// ReturnData AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-anomalydetector-metricdataquery.html#cfn-cloudwatch-anomalydetector-metricdataquery-returndata
	ReturnData *bool `json:"ReturnData,omitempty"`

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
func (r *AnomalyDetector_MetricDataQuery) AWSCloudFormationType() string {
	return "AWS::CloudWatch::AnomalyDetector.MetricDataQuery"
}
