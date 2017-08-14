package resources

// AWS::CloudWatch::Dashboard AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-dashboard.html
type AWSCloudWatchDashboard struct {

	// DashboardBody AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-dashboard.html#cfn-cloudwatch-dashboard-dashboardbody

	DashboardBody string `json:"DashboardBody"`

	// DashboardName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-dashboard.html#cfn-cloudwatch-dashboard-dashboardname

	DashboardName string `json:"DashboardName"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCloudWatchDashboard) AWSCloudFormationType() string {
	return "AWS::CloudWatch::Dashboard"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCloudWatchDashboard) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
