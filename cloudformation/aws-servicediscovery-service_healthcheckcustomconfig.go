package cloudformation

import (
	"encoding/json"
)

// AWSServiceDiscoveryService_HealthCheckCustomConfig AWS CloudFormation Resource (AWS::ServiceDiscovery::Service.HealthCheckCustomConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicediscovery-service-healthcheckcustomconfig.html
type AWSServiceDiscoveryService_HealthCheckCustomConfig struct {

	// FailureThreshold AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicediscovery-service-healthcheckcustomconfig.html#cfn-servicediscovery-service-healthcheckcustomconfig-failurethreshold
	FailureThreshold *Value `json:"FailureThreshold,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSServiceDiscoveryService_HealthCheckCustomConfig) AWSCloudFormationType() string {
	return "AWS::ServiceDiscovery::Service.HealthCheckCustomConfig"
}

func (r *AWSServiceDiscoveryService_HealthCheckCustomConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}
