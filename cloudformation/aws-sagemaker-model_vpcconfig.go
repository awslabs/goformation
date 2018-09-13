package cloudformation

import (
	"encoding/json"
)

// AWSSageMakerModel_VpcConfig AWS CloudFormation Resource (AWS::SageMaker::Model.VpcConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-model-vpcconfig.html
type AWSSageMakerModel_VpcConfig struct {

	// SecurityGroupIds AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-model-vpcconfig.html#cfn-sagemaker-model-vpcconfig-securitygroupids
	SecurityGroupIds []*Value `json:"SecurityGroupIds,omitempty"`

	// Subnets AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-model-vpcconfig.html#cfn-sagemaker-model-vpcconfig-subnets
	Subnets []*Value `json:"Subnets,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSSageMakerModel_VpcConfig) AWSCloudFormationType() string {
	return "AWS::SageMaker::Model.VpcConfig"
}

func (r *AWSSageMakerModel_VpcConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}
