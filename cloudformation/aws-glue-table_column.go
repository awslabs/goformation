package cloudformation

import (
	"encoding/json"
)

// AWSGlueTable_Column AWS CloudFormation Resource (AWS::Glue::Table.Column)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-column.html
type AWSGlueTable_Column struct {

	// Comment AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-column.html#cfn-glue-table-column-comment
	Comment *Value `json:"Comment,omitempty"`

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-column.html#cfn-glue-table-column-name
	Name *Value `json:"Name,omitempty"`

	// Type AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-column.html#cfn-glue-table-column-type
	Type *Value `json:"Type,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSGlueTable_Column) AWSCloudFormationType() string {
	return "AWS::Glue::Table.Column"
}

func (r *AWSGlueTable_Column) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}
