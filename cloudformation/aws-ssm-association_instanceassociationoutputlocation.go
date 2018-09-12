package cloudformation

import (
	"encoding/json"
)

// AWSSSMAssociation_InstanceAssociationOutputLocation AWS CloudFormation Resource (AWS::SSM::Association.InstanceAssociationOutputLocation)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-instanceassociationoutputlocation.html
type AWSSSMAssociation_InstanceAssociationOutputLocation struct {

	// S3Location AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-instanceassociationoutputlocation.html#cfn-ssm-association-instanceassociationoutputlocation-s3location
	S3Location *AWSSSMAssociation_S3OutputLocation `json:"S3Location,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSSSMAssociation_InstanceAssociationOutputLocation) AWSCloudFormationType() string {
	return "AWS::SSM::Association.InstanceAssociationOutputLocation"
}

func (r *AWSSSMAssociation_InstanceAssociationOutputLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}
