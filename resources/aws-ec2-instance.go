package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EC2::Instance AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html
type AWSEC2Instance struct {

	// AdditionalInfo AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-additionalinfo
	AdditionalInfo string `json:"AdditionalInfo"`

	// Affinity AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-affinity
	Affinity string `json:"Affinity"`

	// AvailabilityZone AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-availabilityzone
	AvailabilityZone string `json:"AvailabilityZone"`

	// BlockDeviceMappings AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-blockdevicemappings
	BlockDeviceMappings []AWSEC2Instance_BlockDeviceMapping `json:"BlockDeviceMappings"`

	// DisableApiTermination AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-disableapitermination
	DisableApiTermination bool `json:"DisableApiTermination"`

	// EbsOptimized AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-ebsoptimized
	EbsOptimized bool `json:"EbsOptimized"`

	// HostId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-hostid
	HostId string `json:"HostId"`

	// IamInstanceProfile AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-iaminstanceprofile
	IamInstanceProfile string `json:"IamInstanceProfile"`

	// ImageId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-imageid
	ImageId string `json:"ImageId"`

	// InstanceInitiatedShutdownBehavior AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-instanceinitiatedshutdownbehavior
	InstanceInitiatedShutdownBehavior string `json:"InstanceInitiatedShutdownBehavior"`

	// InstanceType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-instancetype
	InstanceType string `json:"InstanceType"`

	// Ipv6AddressCount AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-ipv6addresscount
	Ipv6AddressCount int64 `json:"Ipv6AddressCount"`

	// Ipv6Addresses AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-ipv6addresses
	Ipv6Addresses []AWSEC2Instance_InstanceIpv6Address `json:"Ipv6Addresses"`

	// KernelId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-kernelid
	KernelId string `json:"KernelId"`

	// KeyName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-keyname
	KeyName string `json:"KeyName"`

	// Monitoring AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-monitoring
	Monitoring bool `json:"Monitoring"`

	// NetworkInterfaces AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-networkinterfaces
	NetworkInterfaces []AWSEC2Instance_NetworkInterface `json:"NetworkInterfaces"`

	// PlacementGroupName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-placementgroupname
	PlacementGroupName string `json:"PlacementGroupName"`

	// PrivateIpAddress AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-privateipaddress
	PrivateIpAddress string `json:"PrivateIpAddress"`

	// RamdiskId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-ramdiskid
	RamdiskId string `json:"RamdiskId"`

	// SecurityGroupIds AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-securitygroupids
	SecurityGroupIds []string `json:"SecurityGroupIds"`

	// SecurityGroups AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-securitygroups
	SecurityGroups []string `json:"SecurityGroups"`

	// SourceDestCheck AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-sourcedestcheck
	SourceDestCheck bool `json:"SourceDestCheck"`

	// SsmAssociations AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-ssmassociations
	SsmAssociations []AWSEC2Instance_SsmAssociation `json:"SsmAssociations"`

	// SubnetId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-subnetid
	SubnetId string `json:"SubnetId"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-tags
	Tags []Tag `json:"Tags"`

	// Tenancy AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-tenancy
	Tenancy string `json:"Tenancy"`

	// UserData AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-userdata
	UserData string `json:"UserData"`

	// Volumes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-volumes
	Volumes []AWSEC2Instance_Volume `json:"Volumes"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2Instance) AWSCloudFormationType() string {
	return "AWS::EC2::Instance"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2Instance) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEC2InstanceResources retrieves all AWSEC2Instance items from a CloudFormation template
func GetAllAWSEC2InstanceResources(template *Template) map[string]*AWSEC2Instance {

	results := map[string]*AWSEC2Instance{}
	for name, resource := range template.Resources {
		result := &AWSEC2Instance{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEC2InstanceWithName retrieves all AWSEC2Instance items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSEC2InstanceWithName(name string, template *Template) (*AWSEC2Instance, error) {

	result := &AWSEC2Instance{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEC2Instance{}, errors.New("resource not found")

}
