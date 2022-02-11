package applicationinsights

import (
	"github.com/awslabs/goformation/v6/cloudformation/policies"
)

// Application_HANAPrometheusExporter AWS CloudFormation Resource (AWS::ApplicationInsights::Application.HANAPrometheusExporter)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-hanaprometheusexporter.html
type Application_HANAPrometheusExporter struct {

	// AgreeToInstallHANADBClient AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-hanaprometheusexporter.html#cfn-applicationinsights-application-hanaprometheusexporter-agreetoinstallhanadbclient
	AgreeToInstallHANADBClient bool `json:"AgreeToInstallHANADBClient"`

	// HANAPort AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-hanaprometheusexporter.html#cfn-applicationinsights-application-hanaprometheusexporter-hanaport
	HANAPort string `json:"HANAPort"`

	// HANASID AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-hanaprometheusexporter.html#cfn-applicationinsights-application-hanaprometheusexporter-hanasid
	HANASID string `json:"HANASID"`

	// HANASecretName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-hanaprometheusexporter.html#cfn-applicationinsights-application-hanaprometheusexporter-hanasecretname
	HANASecretName string `json:"HANASecretName"`

	// PrometheusPort AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-hanaprometheusexporter.html#cfn-applicationinsights-application-hanaprometheusexporter-prometheusport
	PrometheusPort *string `json:"PrometheusPort,omitempty"`

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
func (r *Application_HANAPrometheusExporter) AWSCloudFormationType() string {
	return "AWS::ApplicationInsights::Application.HANAPrometheusExporter"
}
