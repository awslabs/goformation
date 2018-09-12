package cloudformation

import (
	"encoding/json"
)

// AWSGlueTable_SkewedInfo AWS CloudFormation Resource (AWS::Glue::Table.SkewedInfo)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-skewedinfo.html
type AWSGlueTable_SkewedInfo struct {

	// SkewedColumnNames AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-skewedinfo.html#cfn-glue-table-skewedinfo-skewedcolumnnames
	SkewedColumnNames []*Value `json:"SkewedColumnNames,omitempty"`

	// SkewedColumnValueLocationMaps AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-skewedinfo.html#cfn-glue-table-skewedinfo-skewedcolumnvaluelocationmaps
	SkewedColumnValueLocationMaps interface{} `json:"SkewedColumnValueLocationMaps,omitempty"`

	// SkewedColumnValues AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-skewedinfo.html#cfn-glue-table-skewedinfo-skewedcolumnvalues
	SkewedColumnValues []*Value `json:"SkewedColumnValues,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSGlueTable_SkewedInfo) AWSCloudFormationType() string {
	return "AWS::Glue::Table.SkewedInfo"
}

func (r *AWSGlueTable_SkewedInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}
