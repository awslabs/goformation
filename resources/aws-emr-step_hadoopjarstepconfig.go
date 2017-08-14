package resources

// AWS::EMR::Step.HadoopJarStepConfig AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-step-hadoopjarstepconfig.html
type AWSEMRStepHadoopJarStepConfig struct {

	// Args AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-step-hadoopjarstepconfig.html#cfn-elasticmapreduce-step-hadoopjarstepconfig-args
	Args []string `json:"Args"`

	// Jar AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-step-hadoopjarstepconfig.html#cfn-elasticmapreduce-step-hadoopjarstepconfig-jar
	Jar string `json:"Jar"`

	// MainClass AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-step-hadoopjarstepconfig.html#cfn-elasticmapreduce-step-hadoopjarstepconfig-mainclass
	MainClass string `json:"MainClass"`

	// StepProperties AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-step-hadoopjarstepconfig.html#cfn-elasticmapreduce-step-hadoopjarstepconfig-stepproperties
	StepProperties []AWSEMRStepHadoopJarStepConfigKeyValue `json:"StepProperties"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRStepHadoopJarStepConfig) AWSCloudFormationType() string {
	return "AWS::EMR::Step.HadoopJarStepConfig"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRStepHadoopJarStepConfig) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
