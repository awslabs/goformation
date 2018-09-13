package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSEKSCluster AWS CloudFormation Resource (AWS::EKS::Cluster)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-cluster.html
type AWSEKSCluster struct {

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-cluster.html#cfn-eks-cluster-name
	Name *Value `json:"Name,omitempty"`

	// ResourcesVpcConfig AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-cluster.html#cfn-eks-cluster-resourcesvpcconfig
	ResourcesVpcConfig *AWSEKSCluster_ResourcesVpcConfig `json:"ResourcesVpcConfig,omitempty"`

	// RoleArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-cluster.html#cfn-eks-cluster-rolearn
	RoleArn *Value `json:"RoleArn,omitempty"`

	// Version AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-cluster.html#cfn-eks-cluster-version
	Version *Value `json:"Version,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEKSCluster) AWSCloudFormationType() string {
	return "AWS::EKS::Cluster"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSEKSCluster) MarshalJSON() ([]byte, error) {
	type Properties AWSEKSCluster
	return json.Marshal(&struct {
		Type       string
		Properties Properties
	}{
		Type:       r.AWSCloudFormationType(),
		Properties: (Properties)(*r),
	})
}

// UnmarshalJSON is a custom JSON unmarshalling hook that strips the outer
// AWS CloudFormation resource object, and just keeps the 'Properties' field.
func (r *AWSEKSCluster) UnmarshalJSON(b []byte) error {
	type Properties AWSEKSCluster
	res := &struct {
		Type       string
		Properties *Properties
	}{}
	if err := json.Unmarshal(b, &res); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return err
	}

	// If the resource has no Properties set, it could be nil
	if res.Properties != nil {
		*r = AWSEKSCluster(*res.Properties)
	}

	return nil
}

// GetAllAWSEKSClusterResources retrieves all AWSEKSCluster items from an AWS CloudFormation template
func (t *Template) GetAllAWSEKSClusterResources() map[string]AWSEKSCluster {
	results := map[string]AWSEKSCluster{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSEKSCluster:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::EKS::Cluster" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						result := &AWSEKSCluster{}
						if err := result.UnmarshalJSON(b); err == nil {
							results[name] = *result
						}
					}
				}
			}
		}
	}
	return results
}

// GetAWSEKSClusterWithName retrieves all AWSEKSCluster items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEKSClusterWithName(name string) (AWSEKSCluster, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSEKSCluster:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::EKS::Cluster" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						result := &AWSEKSCluster{}
						if err := result.UnmarshalJSON(b); err == nil {
							return *result, nil
						}
					}
				}
			}
		}
	}
	return AWSEKSCluster{}, errors.New("resource not found")
}
