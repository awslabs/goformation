package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Batch::ComputeEnvironment.ComputeResources AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-computeresources.html
type AWSBatchComputeEnvironment_ComputeResources struct {

	// BidPercentage AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-computeresources.html#cfn-batch-computeenvironment-computeresources-bidpercentage
	BidPercentage int64 `json:"BidPercentage"`

	// DesiredvCpus AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-computeresources.html#cfn-batch-computeenvironment-computeresources-desiredvcpus
	DesiredvCpus int64 `json:"DesiredvCpus"`

	// Ec2KeyPair AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-computeresources.html#cfn-batch-computeenvironment-computeresources-ec2keypair
	Ec2KeyPair string `json:"Ec2KeyPair"`

	// ImageId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-computeresources.html#cfn-batch-computeenvironment-computeresources-imageid
	ImageId string `json:"ImageId"`

	// InstanceRole AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-computeresources.html#cfn-batch-computeenvironment-computeresources-instancerole
	InstanceRole string `json:"InstanceRole"`

	// InstanceTypes AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-computeresources.html#cfn-batch-computeenvironment-computeresources-instancetypes
	InstanceTypes []string `json:"InstanceTypes"`

	// MaxvCpus AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-computeresources.html#cfn-batch-computeenvironment-computeresources-maxvcpus
	MaxvCpus int64 `json:"MaxvCpus"`

	// MinvCpus AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-computeresources.html#cfn-batch-computeenvironment-computeresources-minvcpus
	MinvCpus int64 `json:"MinvCpus"`

	// SecurityGroupIds AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-computeresources.html#cfn-batch-computeenvironment-computeresources-securitygroupids
	SecurityGroupIds []string `json:"SecurityGroupIds"`

	// SpotIamFleetRole AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-computeresources.html#cfn-batch-computeenvironment-computeresources-spotiamfleetrole
	SpotIamFleetRole string `json:"SpotIamFleetRole"`

	// Subnets AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-computeresources.html#cfn-batch-computeenvironment-computeresources-subnets
	Subnets []string `json:"Subnets"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-computeresources.html#cfn-batch-computeenvironment-computeresources-tags
	Tags interface{} `json:"Tags"`

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-computeresources.html#cfn-batch-computeenvironment-computeresources-type
	Type string `json:"Type"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSBatchComputeEnvironment_ComputeResources) AWSCloudFormationType() string {
	return "AWS::Batch::ComputeEnvironment.ComputeResources"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSBatchComputeEnvironment_ComputeResources) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSBatchComputeEnvironment_ComputeResourcesResources retrieves all AWSBatchComputeEnvironment_ComputeResources items from a CloudFormation template
func GetAllAWSBatchComputeEnvironment_ComputeResources(template *Template) map[string]*AWSBatchComputeEnvironment_ComputeResources {

	results := map[string]*AWSBatchComputeEnvironment_ComputeResources{}
	for name, resource := range template.Resources {
		result := &AWSBatchComputeEnvironment_ComputeResources{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSBatchComputeEnvironment_ComputeResourcesWithName retrieves all AWSBatchComputeEnvironment_ComputeResources items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSBatchComputeEnvironment_ComputeResources(name string, template *Template) (*AWSBatchComputeEnvironment_ComputeResources, error) {

	result := &AWSBatchComputeEnvironment_ComputeResources{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSBatchComputeEnvironment_ComputeResources{}, errors.New("resource not found")

}
