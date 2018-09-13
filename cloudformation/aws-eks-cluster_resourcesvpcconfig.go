package cloudformation

import (
	"encoding/json"
)

// AWSEKSCluster_ResourcesVpcConfig AWS CloudFormation Resource (AWS::EKS::Cluster.ResourcesVpcConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-resourcesvpcconfig.html
type AWSEKSCluster_ResourcesVpcConfig struct {

	// SecurityGroupIds AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-resourcesvpcconfig.html#cfn-eks-cluster-resourcesvpcconfig-securitygroupids
	SecurityGroupIds []*Value `json:"SecurityGroupIds,omitempty"`

	// SubnetIds AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-resourcesvpcconfig.html#cfn-eks-cluster-resourcesvpcconfig-subnetids
	SubnetIds []*Value `json:"SubnetIds,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEKSCluster_ResourcesVpcConfig) AWSCloudFormationType() string {
	return "AWS::EKS::Cluster.ResourcesVpcConfig"
}

func (r *AWSEKSCluster_ResourcesVpcConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}
