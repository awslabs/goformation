package cloudformation

import (
	"encoding/json"
)

// AWSGlueTable_SerdeInfo AWS CloudFormation Resource (AWS::Glue::Table.SerdeInfo)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-serdeinfo.html
type AWSGlueTable_SerdeInfo struct {

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-serdeinfo.html#cfn-glue-table-serdeinfo-name
	Name *Value `json:"Name,omitempty"`

	// Parameters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-serdeinfo.html#cfn-glue-table-serdeinfo-parameters
	Parameters interface{} `json:"Parameters,omitempty"`

	// SerializationLibrary AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-serdeinfo.html#cfn-glue-table-serdeinfo-serializationlibrary
	SerializationLibrary *Value `json:"SerializationLibrary,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSGlueTable_SerdeInfo) AWSCloudFormationType() string {
	return "AWS::Glue::Table.SerdeInfo"
}

func (r *AWSGlueTable_SerdeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}
