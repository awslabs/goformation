package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::KinesisFirehose::DeliveryStream.CopyCommand AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-redshiftdestinationconfiguration-copycommand.html
type AWSKinesisFirehoseDeliveryStream_CopyCommand struct {

	// CopyOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-redshiftdestinationconfiguration-copycommand.html#cfn-kinesisfirehose-kinesisdeliverystream-redshiftdestinationconfiguration-copycommand-copyoptions

	CopyOptions string `json:"CopyOptions"`

	// DataTableColumns AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-redshiftdestinationconfiguration-copycommand.html#cfn-kinesisfirehose-kinesisdeliverystream-redshiftdestinationconfiguration-copycommand-datatablecolumns

	DataTableColumns string `json:"DataTableColumns"`

	// DataTableName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-redshiftdestinationconfiguration-copycommand.html#cfn-kinesisfirehose-kinesisdeliverystream-redshiftdestinationconfiguration-copycommand-datatablename

	DataTableName string `json:"DataTableName"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisFirehoseDeliveryStream_CopyCommand) AWSCloudFormationType() string {
	return "AWS::KinesisFirehose::DeliveryStream.CopyCommand"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKinesisFirehoseDeliveryStream_CopyCommand) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSKinesisFirehoseDeliveryStream_CopyCommandResources retrieves all AWSKinesisFirehoseDeliveryStream_CopyCommand items from a CloudFormation template
func GetAllAWSKinesisFirehoseDeliveryStream_CopyCommand(template *Template) map[string]*AWSKinesisFirehoseDeliveryStream_CopyCommand {

	results := map[string]*AWSKinesisFirehoseDeliveryStream_CopyCommand{}
	for name, resource := range template.Resources {
		result := &AWSKinesisFirehoseDeliveryStream_CopyCommand{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSKinesisFirehoseDeliveryStream_CopyCommandWithName retrieves all AWSKinesisFirehoseDeliveryStream_CopyCommand items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSKinesisFirehoseDeliveryStream_CopyCommand(name string, template *Template) (*AWSKinesisFirehoseDeliveryStream_CopyCommand, error) {

	result := &AWSKinesisFirehoseDeliveryStream_CopyCommand{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSKinesisFirehoseDeliveryStream_CopyCommand{}, errors.New("resource not found")

}
