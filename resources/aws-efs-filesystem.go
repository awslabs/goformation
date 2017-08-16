package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSEFSFileSystem AWS CloudFormation Resource (AWS::EFS::FileSystem)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-filesystem.html
type AWSEFSFileSystem struct {

	// FileSystemTags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-filesystem.html#cfn-efs-filesystem-filesystemtags
	FileSystemTags []AWSEFSFileSystem_ElasticFileSystemTag `json:"FileSystemTags"`

	// PerformanceMode AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-filesystem.html#cfn-efs-filesystem-performancemode
	PerformanceMode string `json:"PerformanceMode"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEFSFileSystem) AWSCloudFormationType() string {
	return "AWS::EFS::FileSystem"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEFSFileSystem) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEFSFileSystemResources retrieves all AWSEFSFileSystem items from a CloudFormation template
func (t *Template) GetAllAWSEFSFileSystemResources() map[string]*AWSEFSFileSystem {

	results := map[string]*AWSEFSFileSystem{}
	for name, resource := range t.Resources {
		result := &AWSEFSFileSystem{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEFSFileSystemWithName retrieves all AWSEFSFileSystem items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEFSFileSystemWithName(name string) (*AWSEFSFileSystem, error) {

	result := &AWSEFSFileSystem{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEFSFileSystem{}, errors.New("resource not found")

}
