package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EFS::FileSystem.ElasticFileSystemTag AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-filesystem-filesystemtags.html
type AWSEFSFileSystem_ElasticFileSystemTag struct {

	// Key AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-filesystem-filesystemtags.html#cfn-efs-filesystem-filesystemtags-key

	Key string `json:"Key"`

	// Value AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-filesystem-filesystemtags.html#cfn-efs-filesystem-filesystemtags-value

	Value string `json:"Value"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEFSFileSystem_ElasticFileSystemTag) AWSCloudFormationType() string {
	return "AWS::EFS::FileSystem.ElasticFileSystemTag"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEFSFileSystem_ElasticFileSystemTag) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEFSFileSystem_ElasticFileSystemTagResources retrieves all AWSEFSFileSystem_ElasticFileSystemTag items from a CloudFormation template
func GetAllAWSEFSFileSystem_ElasticFileSystemTag(template *Template) map[string]*AWSEFSFileSystem_ElasticFileSystemTag {

	results := map[string]*AWSEFSFileSystem_ElasticFileSystemTag{}
	for name, resource := range template.Resources {
		result := &AWSEFSFileSystem_ElasticFileSystemTag{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEFSFileSystem_ElasticFileSystemTagWithName retrieves all AWSEFSFileSystem_ElasticFileSystemTag items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEFSFileSystem_ElasticFileSystemTag(name string, template *Template) (*AWSEFSFileSystem_ElasticFileSystemTag, error) {

	result := &AWSEFSFileSystem_ElasticFileSystemTag{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEFSFileSystem_ElasticFileSystemTag{}, errors.New("resource not found")

}
