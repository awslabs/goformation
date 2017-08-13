package resources

// AWS::DMS::ReplicationTask AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationtask.html
type AWSDMSReplicationTask struct {
    
    // CdcStartTime AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationtask.html#cfn-dms-replicationtask-cdcstarttime
    CdcStartTime float64 `json:"CdcStartTime"`
    
    // MigrationType AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationtask.html#cfn-dms-replicationtask-migrationtype
    MigrationType string `json:"MigrationType"`
    
    // ReplicationInstanceArn AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationtask.html#cfn-dms-replicationtask-replicationinstancearn
    ReplicationInstanceArn string `json:"ReplicationInstanceArn"`
    
    // ReplicationTaskIdentifier AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationtask.html#cfn-dms-replicationtask-replicationtaskidentifier
    ReplicationTaskIdentifier string `json:"ReplicationTaskIdentifier"`
    
    // ReplicationTaskSettings AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationtask.html#cfn-dms-replicationtask-replicationtasksettings
    ReplicationTaskSettings string `json:"ReplicationTaskSettings"`
    
    // SourceEndpointArn AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationtask.html#cfn-dms-replicationtask-sourceendpointarn
    SourceEndpointArn string `json:"SourceEndpointArn"`
    
    // TableMappings AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationtask.html#cfn-dms-replicationtask-tablemappings
    TableMappings string `json:"TableMappings"`
    
    // Tags AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationtask.html#cfn-dms-replicationtask-tags
    Tags []AWSDMSReplicationTaskTag `json:"Tags"`
    
    // TargetEndpointArn AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationtask.html#cfn-dms-replicationtask-targetendpointarn
    TargetEndpointArn string `json:"TargetEndpointArn"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSDMSReplicationTask) AWSCloudFormationType() string {
    return "AWS::DMS::ReplicationTask"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSDMSReplicationTask) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}