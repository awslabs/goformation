package cloudformation

import (
	"encoding/json"
)

// AWSServiceDiscoveryService_DnsConfig AWS CloudFormation Resource (AWS::ServiceDiscovery::Service.DnsConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicediscovery-service-dnsconfig.html
type AWSServiceDiscoveryService_DnsConfig struct {

	// DnsRecords AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicediscovery-service-dnsconfig.html#cfn-servicediscovery-service-dnsconfig-dnsrecords
	DnsRecords []AWSServiceDiscoveryService_DnsRecord `json:"DnsRecords,omitempty"`

	// NamespaceId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicediscovery-service-dnsconfig.html#cfn-servicediscovery-service-dnsconfig-namespaceid
	NamespaceId *Value `json:"NamespaceId,omitempty"`

	// RoutingPolicy AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicediscovery-service-dnsconfig.html#cfn-servicediscovery-service-dnsconfig-routingpolicy
	RoutingPolicy *Value `json:"RoutingPolicy,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSServiceDiscoveryService_DnsConfig) AWSCloudFormationType() string {
	return "AWS::ServiceDiscovery::Service.DnsConfig"
}

func (r *AWSServiceDiscoveryService_DnsConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}
