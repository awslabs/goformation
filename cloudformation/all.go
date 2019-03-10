package cloudformation

import (
	"fmt"
)

// AllResources fetches an iterable map all CloudFormation and SAM resources
func AllResources() map[string]Resource {
	return map[string]Resource{
		"AWS::AmazonMQ::Broker":                                       &AWSAmazonMQBroker{},
		"AWS::AmazonMQ::Configuration":                                &AWSAmazonMQConfiguration{},
		"AWS::AmazonMQ::ConfigurationAssociation":                     &AWSAmazonMQConfigurationAssociation{},
		"AWS::ApiGateway::Account":                                    &AWSApiGatewayAccount{},
		"AWS::ApiGateway::ApiKey":                                     &AWSApiGatewayApiKey{},
		"AWS::ApiGateway::Authorizer":                                 &AWSApiGatewayAuthorizer{},
		"AWS::ApiGateway::BasePathMapping":                            &AWSApiGatewayBasePathMapping{},
		"AWS::ApiGateway::ClientCertificate":                          &AWSApiGatewayClientCertificate{},
		"AWS::ApiGateway::Deployment":                                 &AWSApiGatewayDeployment{},
		"AWS::ApiGateway::DocumentationPart":                          &AWSApiGatewayDocumentationPart{},
		"AWS::ApiGateway::DocumentationVersion":                       &AWSApiGatewayDocumentationVersion{},
		"AWS::ApiGateway::DomainName":                                 &AWSApiGatewayDomainName{},
		"AWS::ApiGateway::GatewayResponse":                            &AWSApiGatewayGatewayResponse{},
		"AWS::ApiGateway::Method":                                     &AWSApiGatewayMethod{},
		"AWS::ApiGateway::Model":                                      &AWSApiGatewayModel{},
		"AWS::ApiGateway::RequestValidator":                           &AWSApiGatewayRequestValidator{},
		"AWS::ApiGateway::Resource":                                   &AWSApiGatewayResource{},
		"AWS::ApiGateway::RestApi":                                    &AWSApiGatewayRestApi{},
		"AWS::ApiGateway::Stage":                                      &AWSApiGatewayStage{},
		"AWS::ApiGateway::UsagePlan":                                  &AWSApiGatewayUsagePlan{},
		"AWS::ApiGateway::UsagePlanKey":                               &AWSApiGatewayUsagePlanKey{},
		"AWS::ApiGateway::VpcLink":                                    &AWSApiGatewayVpcLink{},
		"AWS::ApiGatewayV2::Api":                                      &AWSApiGatewayV2Api{},
		"AWS::ApiGatewayV2::Authorizer":                               &AWSApiGatewayV2Authorizer{},
		"AWS::ApiGatewayV2::Deployment":                               &AWSApiGatewayV2Deployment{},
		"AWS::ApiGatewayV2::Integration":                              &AWSApiGatewayV2Integration{},
		"AWS::ApiGatewayV2::IntegrationResponse":                      &AWSApiGatewayV2IntegrationResponse{},
		"AWS::ApiGatewayV2::Model":                                    &AWSApiGatewayV2Model{},
		"AWS::ApiGatewayV2::Route":                                    &AWSApiGatewayV2Route{},
		"AWS::ApiGatewayV2::RouteResponse":                            &AWSApiGatewayV2RouteResponse{},
		"AWS::ApiGatewayV2::Stage":                                    &AWSApiGatewayV2Stage{},
		"AWS::AppStream::DirectoryConfig":                             &AWSAppStreamDirectoryConfig{},
		"AWS::AppStream::Fleet":                                       &AWSAppStreamFleet{},
		"AWS::AppStream::ImageBuilder":                                &AWSAppStreamImageBuilder{},
		"AWS::AppStream::Stack":                                       &AWSAppStreamStack{},
		"AWS::AppStream::StackFleetAssociation":                       &AWSAppStreamStackFleetAssociation{},
		"AWS::AppStream::StackUserAssociation":                        &AWSAppStreamStackUserAssociation{},
		"AWS::AppStream::User":                                        &AWSAppStreamUser{},
		"AWS::AppSync::ApiKey":                                        &AWSAppSyncApiKey{},
		"AWS::AppSync::DataSource":                                    &AWSAppSyncDataSource{},
		"AWS::AppSync::FunctionConfiguration":                         &AWSAppSyncFunctionConfiguration{},
		"AWS::AppSync::GraphQLApi":                                    &AWSAppSyncGraphQLApi{},
		"AWS::AppSync::GraphQLSchema":                                 &AWSAppSyncGraphQLSchema{},
		"AWS::AppSync::Resolver":                                      &AWSAppSyncResolver{},
		"AWS::ApplicationAutoScaling::ScalableTarget":                 &AWSApplicationAutoScalingScalableTarget{},
		"AWS::ApplicationAutoScaling::ScalingPolicy":                  &AWSApplicationAutoScalingScalingPolicy{},
		"AWS::Athena::NamedQuery":                                     &AWSAthenaNamedQuery{},
		"AWS::AutoScaling::AutoScalingGroup":                          &AWSAutoScalingAutoScalingGroup{},
		"AWS::AutoScaling::LaunchConfiguration":                       &AWSAutoScalingLaunchConfiguration{},
		"AWS::AutoScaling::LifecycleHook":                             &AWSAutoScalingLifecycleHook{},
		"AWS::AutoScaling::ScalingPolicy":                             &AWSAutoScalingScalingPolicy{},
		"AWS::AutoScaling::ScheduledAction":                           &AWSAutoScalingScheduledAction{},
		"AWS::AutoScalingPlans::ScalingPlan":                          &AWSAutoScalingPlansScalingPlan{},
		"AWS::Batch::ComputeEnvironment":                              &AWSBatchComputeEnvironment{},
		"AWS::Batch::JobDefinition":                                   &AWSBatchJobDefinition{},
		"AWS::Batch::JobQueue":                                        &AWSBatchJobQueue{},
		"AWS::Budgets::Budget":                                        &AWSBudgetsBudget{},
		"AWS::CertificateManager::Certificate":                        &AWSCertificateManagerCertificate{},
		"AWS::Cloud9::EnvironmentEC2":                                 &AWSCloud9EnvironmentEC2{},
		"AWS::CloudFormation::CustomResource":                         &AWSCloudFormationCustomResource{},
		"AWS::CloudFormation::Macro":                                  &AWSCloudFormationMacro{},
		"AWS::CloudFormation::Stack":                                  &AWSCloudFormationStack{},
		"AWS::CloudFormation::WaitCondition":                          &AWSCloudFormationWaitCondition{},
		"AWS::CloudFormation::WaitConditionHandle":                    &AWSCloudFormationWaitConditionHandle{},
		"AWS::CloudFront::CloudFrontOriginAccessIdentity":             &AWSCloudFrontCloudFrontOriginAccessIdentity{},
		"AWS::CloudFront::Distribution":                               &AWSCloudFrontDistribution{},
		"AWS::CloudFront::StreamingDistribution":                      &AWSCloudFrontStreamingDistribution{},
		"AWS::CloudTrail::Trail":                                      &AWSCloudTrailTrail{},
		"AWS::CloudWatch::Alarm":                                      &AWSCloudWatchAlarm{},
		"AWS::CloudWatch::Dashboard":                                  &AWSCloudWatchDashboard{},
		"AWS::CodeBuild::Project":                                     &AWSCodeBuildProject{},
		"AWS::CodeCommit::Repository":                                 &AWSCodeCommitRepository{},
		"AWS::CodeDeploy::Application":                                &AWSCodeDeployApplication{},
		"AWS::CodeDeploy::DeploymentConfig":                           &AWSCodeDeployDeploymentConfig{},
		"AWS::CodeDeploy::DeploymentGroup":                            &AWSCodeDeployDeploymentGroup{},
		"AWS::CodePipeline::CustomActionType":                         &AWSCodePipelineCustomActionType{},
		"AWS::CodePipeline::Pipeline":                                 &AWSCodePipelinePipeline{},
		"AWS::CodePipeline::Webhook":                                  &AWSCodePipelineWebhook{},
		"AWS::Cognito::IdentityPool":                                  &AWSCognitoIdentityPool{},
		"AWS::Cognito::IdentityPoolRoleAttachment":                    &AWSCognitoIdentityPoolRoleAttachment{},
		"AWS::Cognito::UserPool":                                      &AWSCognitoUserPool{},
		"AWS::Cognito::UserPoolClient":                                &AWSCognitoUserPoolClient{},
		"AWS::Cognito::UserPoolGroup":                                 &AWSCognitoUserPoolGroup{},
		"AWS::Cognito::UserPoolUser":                                  &AWSCognitoUserPoolUser{},
		"AWS::Cognito::UserPoolUserToGroupAttachment":                 &AWSCognitoUserPoolUserToGroupAttachment{},
		"AWS::Config::AggregationAuthorization":                       &AWSConfigAggregationAuthorization{},
		"AWS::Config::ConfigRule":                                     &AWSConfigConfigRule{},
		"AWS::Config::ConfigurationAggregator":                        &AWSConfigConfigurationAggregator{},
		"AWS::Config::ConfigurationRecorder":                          &AWSConfigConfigurationRecorder{},
		"AWS::Config::DeliveryChannel":                                &AWSConfigDeliveryChannel{},
		"AWS::DAX::Cluster":                                           &AWSDAXCluster{},
		"AWS::DAX::ParameterGroup":                                    &AWSDAXParameterGroup{},
		"AWS::DAX::SubnetGroup":                                       &AWSDAXSubnetGroup{},
		"AWS::DLM::LifecyclePolicy":                                   &AWSDLMLifecyclePolicy{},
		"AWS::DMS::Certificate":                                       &AWSDMSCertificate{},
		"AWS::DMS::Endpoint":                                          &AWSDMSEndpoint{},
		"AWS::DMS::EventSubscription":                                 &AWSDMSEventSubscription{},
		"AWS::DMS::ReplicationInstance":                               &AWSDMSReplicationInstance{},
		"AWS::DMS::ReplicationSubnetGroup":                            &AWSDMSReplicationSubnetGroup{},
		"AWS::DMS::ReplicationTask":                                   &AWSDMSReplicationTask{},
		"AWS::DataPipeline::Pipeline":                                 &AWSDataPipelinePipeline{},
		"AWS::DirectoryService::MicrosoftAD":                          &AWSDirectoryServiceMicrosoftAD{},
		"AWS::DirectoryService::SimpleAD":                             &AWSDirectoryServiceSimpleAD{},
		"AWS::DocDB::DBCluster":                                       &AWSDocDBDBCluster{},
		"AWS::DocDB::DBClusterParameterGroup":                         &AWSDocDBDBClusterParameterGroup{},
		"AWS::DocDB::DBInstance":                                      &AWSDocDBDBInstance{},
		"AWS::DocDB::DBSubnetGroup":                                   &AWSDocDBDBSubnetGroup{},
		"AWS::DynamoDB::Table":                                        &AWSDynamoDBTable{},
		"AWS::EC2::CustomerGateway":                                   &AWSEC2CustomerGateway{},
		"AWS::EC2::DHCPOptions":                                       &AWSEC2DHCPOptions{},
		"AWS::EC2::EC2Fleet":                                          &AWSEC2EC2Fleet{},
		"AWS::EC2::EIP":                                               &AWSEC2EIP{},
		"AWS::EC2::EIPAssociation":                                    &AWSEC2EIPAssociation{},
		"AWS::EC2::EgressOnlyInternetGateway":                         &AWSEC2EgressOnlyInternetGateway{},
		"AWS::EC2::FlowLog":                                           &AWSEC2FlowLog{},
		"AWS::EC2::Host":                                              &AWSEC2Host{},
		"AWS::EC2::Instance":                                          &AWSEC2Instance{},
		"AWS::EC2::InternetGateway":                                   &AWSEC2InternetGateway{},
		"AWS::EC2::LaunchTemplate":                                    &AWSEC2LaunchTemplate{},
		"AWS::EC2::NatGateway":                                        &AWSEC2NatGateway{},
		"AWS::EC2::NetworkAcl":                                        &AWSEC2NetworkAcl{},
		"AWS::EC2::NetworkAclEntry":                                   &AWSEC2NetworkAclEntry{},
		"AWS::EC2::NetworkInterface":                                  &AWSEC2NetworkInterface{},
		"AWS::EC2::NetworkInterfaceAttachment":                        &AWSEC2NetworkInterfaceAttachment{},
		"AWS::EC2::NetworkInterfacePermission":                        &AWSEC2NetworkInterfacePermission{},
		"AWS::EC2::PlacementGroup":                                    &AWSEC2PlacementGroup{},
		"AWS::EC2::Route":                                             &AWSEC2Route{},
		"AWS::EC2::RouteTable":                                        &AWSEC2RouteTable{},
		"AWS::EC2::SecurityGroup":                                     &AWSEC2SecurityGroup{},
		"AWS::EC2::SecurityGroupEgress":                               &AWSEC2SecurityGroupEgress{},
		"AWS::EC2::SecurityGroupIngress":                              &AWSEC2SecurityGroupIngress{},
		"AWS::EC2::SpotFleet":                                         &AWSEC2SpotFleet{},
		"AWS::EC2::Subnet":                                            &AWSEC2Subnet{},
		"AWS::EC2::SubnetCidrBlock":                                   &AWSEC2SubnetCidrBlock{},
		"AWS::EC2::SubnetNetworkAclAssociation":                       &AWSEC2SubnetNetworkAclAssociation{},
		"AWS::EC2::SubnetRouteTableAssociation":                       &AWSEC2SubnetRouteTableAssociation{},
		"AWS::EC2::TransitGateway":                                    &AWSEC2TransitGateway{},
		"AWS::EC2::TransitGatewayAttachment":                          &AWSEC2TransitGatewayAttachment{},
		"AWS::EC2::TransitGatewayRoute":                               &AWSEC2TransitGatewayRoute{},
		"AWS::EC2::TransitGatewayRouteTable":                          &AWSEC2TransitGatewayRouteTable{},
		"AWS::EC2::TransitGatewayRouteTableAssociation":               &AWSEC2TransitGatewayRouteTableAssociation{},
		"AWS::EC2::TransitGatewayRouteTablePropagation":               &AWSEC2TransitGatewayRouteTablePropagation{},
		"AWS::EC2::TrunkInterfaceAssociation":                         &AWSEC2TrunkInterfaceAssociation{},
		"AWS::EC2::VPC":                                               &AWSEC2VPC{},
		"AWS::EC2::VPCCidrBlock":                                      &AWSEC2VPCCidrBlock{},
		"AWS::EC2::VPCDHCPOptionsAssociation":                         &AWSEC2VPCDHCPOptionsAssociation{},
		"AWS::EC2::VPCEndpoint":                                       &AWSEC2VPCEndpoint{},
		"AWS::EC2::VPCEndpointConnectionNotification":                 &AWSEC2VPCEndpointConnectionNotification{},
		"AWS::EC2::VPCEndpointServicePermissions":                     &AWSEC2VPCEndpointServicePermissions{},
		"AWS::EC2::VPCGatewayAttachment":                              &AWSEC2VPCGatewayAttachment{},
		"AWS::EC2::VPCPeeringConnection":                              &AWSEC2VPCPeeringConnection{},
		"AWS::EC2::VPNConnection":                                     &AWSEC2VPNConnection{},
		"AWS::EC2::VPNConnectionRoute":                                &AWSEC2VPNConnectionRoute{},
		"AWS::EC2::VPNGateway":                                        &AWSEC2VPNGateway{},
		"AWS::EC2::VPNGatewayRoutePropagation":                        &AWSEC2VPNGatewayRoutePropagation{},
		"AWS::EC2::Volume":                                            &AWSEC2Volume{},
		"AWS::EC2::VolumeAttachment":                                  &AWSEC2VolumeAttachment{},
		"AWS::ECR::Repository":                                        &AWSECRRepository{},
		"AWS::ECS::Cluster":                                           &AWSECSCluster{},
		"AWS::ECS::Service":                                           &AWSECSService{},
		"AWS::ECS::TaskDefinition":                                    &AWSECSTaskDefinition{},
		"AWS::EFS::FileSystem":                                        &AWSEFSFileSystem{},
		"AWS::EFS::MountTarget":                                       &AWSEFSMountTarget{},
		"AWS::EKS::Cluster":                                           &AWSEKSCluster{},
		"AWS::EMR::Cluster":                                           &AWSEMRCluster{},
		"AWS::EMR::InstanceFleetConfig":                               &AWSEMRInstanceFleetConfig{},
		"AWS::EMR::InstanceGroupConfig":                               &AWSEMRInstanceGroupConfig{},
		"AWS::EMR::SecurityConfiguration":                             &AWSEMRSecurityConfiguration{},
		"AWS::EMR::Step":                                              &AWSEMRStep{},
		"AWS::ElastiCache::CacheCluster":                              &AWSElastiCacheCacheCluster{},
		"AWS::ElastiCache::ParameterGroup":                            &AWSElastiCacheParameterGroup{},
		"AWS::ElastiCache::ReplicationGroup":                          &AWSElastiCacheReplicationGroup{},
		"AWS::ElastiCache::SecurityGroup":                             &AWSElastiCacheSecurityGroup{},
		"AWS::ElastiCache::SecurityGroupIngress":                      &AWSElastiCacheSecurityGroupIngress{},
		"AWS::ElastiCache::SubnetGroup":                               &AWSElastiCacheSubnetGroup{},
		"AWS::ElasticBeanstalk::Application":                          &AWSElasticBeanstalkApplication{},
		"AWS::ElasticBeanstalk::ApplicationVersion":                   &AWSElasticBeanstalkApplicationVersion{},
		"AWS::ElasticBeanstalk::ConfigurationTemplate":                &AWSElasticBeanstalkConfigurationTemplate{},
		"AWS::ElasticBeanstalk::Environment":                          &AWSElasticBeanstalkEnvironment{},
		"AWS::ElasticLoadBalancing::LoadBalancer":                     &AWSElasticLoadBalancingLoadBalancer{},
		"AWS::ElasticLoadBalancingV2::Listener":                       &AWSElasticLoadBalancingV2Listener{},
		"AWS::ElasticLoadBalancingV2::ListenerCertificate":            &AWSElasticLoadBalancingV2ListenerCertificate{},
		"AWS::ElasticLoadBalancingV2::ListenerRule":                   &AWSElasticLoadBalancingV2ListenerRule{},
		"AWS::ElasticLoadBalancingV2::LoadBalancer":                   &AWSElasticLoadBalancingV2LoadBalancer{},
		"AWS::ElasticLoadBalancingV2::TargetGroup":                    &AWSElasticLoadBalancingV2TargetGroup{},
		"AWS::Elasticsearch::Domain":                                  &AWSElasticsearchDomain{},
		"AWS::Events::EventBusPolicy":                                 &AWSEventsEventBusPolicy{},
		"AWS::Events::Rule":                                           &AWSEventsRule{},
		"AWS::FSx::FileSystem":                                        &AWSFSxFileSystem{},
		"AWS::GameLift::Alias":                                        &AWSGameLiftAlias{},
		"AWS::GameLift::Build":                                        &AWSGameLiftBuild{},
		"AWS::GameLift::Fleet":                                        &AWSGameLiftFleet{},
		"AWS::Glue::Classifier":                                       &AWSGlueClassifier{},
		"AWS::Glue::Connection":                                       &AWSGlueConnection{},
		"AWS::Glue::Crawler":                                          &AWSGlueCrawler{},
		"AWS::Glue::Database":                                         &AWSGlueDatabase{},
		"AWS::Glue::DevEndpoint":                                      &AWSGlueDevEndpoint{},
		"AWS::Glue::Job":                                              &AWSGlueJob{},
		"AWS::Glue::Partition":                                        &AWSGluePartition{},
		"AWS::Glue::Table":                                            &AWSGlueTable{},
		"AWS::Glue::Trigger":                                          &AWSGlueTrigger{},
		"AWS::GuardDuty::Detector":                                    &AWSGuardDutyDetector{},
		"AWS::GuardDuty::Filter":                                      &AWSGuardDutyFilter{},
		"AWS::GuardDuty::IPSet":                                       &AWSGuardDutyIPSet{},
		"AWS::GuardDuty::Master":                                      &AWSGuardDutyMaster{},
		"AWS::GuardDuty::Member":                                      &AWSGuardDutyMember{},
		"AWS::GuardDuty::ThreatIntelSet":                              &AWSGuardDutyThreatIntelSet{},
		"AWS::IAM::AccessKey":                                         &AWSIAMAccessKey{},
		"AWS::IAM::Group":                                             &AWSIAMGroup{},
		"AWS::IAM::InstanceProfile":                                   &AWSIAMInstanceProfile{},
		"AWS::IAM::ManagedPolicy":                                     &AWSIAMManagedPolicy{},
		"AWS::IAM::Policy":                                            &AWSIAMPolicy{},
		"AWS::IAM::Role":                                              &AWSIAMRole{},
		"AWS::IAM::ServiceLinkedRole":                                 &AWSIAMServiceLinkedRole{},
		"AWS::IAM::User":                                              &AWSIAMUser{},
		"AWS::IAM::UserToGroupAddition":                               &AWSIAMUserToGroupAddition{},
		"AWS::Inspector::AssessmentTarget":                            &AWSInspectorAssessmentTarget{},
		"AWS::Inspector::AssessmentTemplate":                          &AWSInspectorAssessmentTemplate{},
		"AWS::Inspector::ResourceGroup":                               &AWSInspectorResourceGroup{},
		"AWS::IoT1Click::Device":                                      &AWSIoT1ClickDevice{},
		"AWS::IoT1Click::Placement":                                   &AWSIoT1ClickPlacement{},
		"AWS::IoT1Click::Project":                                     &AWSIoT1ClickProject{},
		"AWS::IoT::Certificate":                                       &AWSIoTCertificate{},
		"AWS::IoT::Policy":                                            &AWSIoTPolicy{},
		"AWS::IoT::PolicyPrincipalAttachment":                         &AWSIoTPolicyPrincipalAttachment{},
		"AWS::IoT::Thing":                                             &AWSIoTThing{},
		"AWS::IoT::ThingPrincipalAttachment":                          &AWSIoTThingPrincipalAttachment{},
		"AWS::IoT::TopicRule":                                         &AWSIoTTopicRule{},
		"AWS::IoTAnalytics::Channel":                                  &AWSIoTAnalyticsChannel{},
		"AWS::IoTAnalytics::Dataset":                                  &AWSIoTAnalyticsDataset{},
		"AWS::IoTAnalytics::Datastore":                                &AWSIoTAnalyticsDatastore{},
		"AWS::IoTAnalytics::Pipeline":                                 &AWSIoTAnalyticsPipeline{},
		"AWS::KMS::Alias":                                             &AWSKMSAlias{},
		"AWS::KMS::Key":                                               &AWSKMSKey{},
		"AWS::Kinesis::Stream":                                        &AWSKinesisStream{},
		"AWS::Kinesis::StreamConsumer":                                &AWSKinesisStreamConsumer{},
		"AWS::KinesisAnalytics::Application":                          &AWSKinesisAnalyticsApplication{},
		"AWS::KinesisAnalytics::ApplicationOutput":                    &AWSKinesisAnalyticsApplicationOutput{},
		"AWS::KinesisAnalytics::ApplicationReferenceDataSource":       &AWSKinesisAnalyticsApplicationReferenceDataSource{},
		"AWS::KinesisAnalyticsV2::Application":                        &AWSKinesisAnalyticsV2Application{},
		"AWS::KinesisAnalyticsV2::ApplicationCloudWatchLoggingOption": &AWSKinesisAnalyticsV2ApplicationCloudWatchLoggingOption{},
		"AWS::KinesisAnalyticsV2::ApplicationOutput":                  &AWSKinesisAnalyticsV2ApplicationOutput{},
		"AWS::KinesisAnalyticsV2::ApplicationReferenceDataSource":     &AWSKinesisAnalyticsV2ApplicationReferenceDataSource{},
		"AWS::KinesisFirehose::DeliveryStream":                        &AWSKinesisFirehoseDeliveryStream{},
		"AWS::Lambda::Alias":                                          &AWSLambdaAlias{},
		"AWS::Lambda::EventSourceMapping":                             &AWSLambdaEventSourceMapping{},
		"AWS::Lambda::Function":                                       &AWSLambdaFunction{},
		"AWS::Lambda::LayerVersion":                                   &AWSLambdaLayerVersion{},
		"AWS::Lambda::LayerVersionPermission":                         &AWSLambdaLayerVersionPermission{},
		"AWS::Lambda::Permission":                                     &AWSLambdaPermission{},
		"AWS::Lambda::Version":                                        &AWSLambdaVersion{},
		"AWS::Logs::Destination":                                      &AWSLogsDestination{},
		"AWS::Logs::LogGroup":                                         &AWSLogsLogGroup{},
		"AWS::Logs::LogStream":                                        &AWSLogsLogStream{},
		"AWS::Logs::MetricFilter":                                     &AWSLogsMetricFilter{},
		"AWS::Logs::SubscriptionFilter":                               &AWSLogsSubscriptionFilter{},
		"AWS::Neptune::DBCluster":                                     &AWSNeptuneDBCluster{},
		"AWS::Neptune::DBClusterParameterGroup":                       &AWSNeptuneDBClusterParameterGroup{},
		"AWS::Neptune::DBInstance":                                    &AWSNeptuneDBInstance{},
		"AWS::Neptune::DBParameterGroup":                              &AWSNeptuneDBParameterGroup{},
		"AWS::Neptune::DBSubnetGroup":                                 &AWSNeptuneDBSubnetGroup{},
		"AWS::OpsWorks::App":                                          &AWSOpsWorksApp{},
		"AWS::OpsWorks::ElasticLoadBalancerAttachment":                &AWSOpsWorksElasticLoadBalancerAttachment{},
		"AWS::OpsWorks::Instance":                                     &AWSOpsWorksInstance{},
		"AWS::OpsWorks::Layer":                                        &AWSOpsWorksLayer{},
		"AWS::OpsWorks::Stack":                                        &AWSOpsWorksStack{},
		"AWS::OpsWorks::UserProfile":                                  &AWSOpsWorksUserProfile{},
		"AWS::OpsWorks::Volume":                                       &AWSOpsWorksVolume{},
		"AWS::OpsWorksCM::Server":                                     &AWSOpsWorksCMServer{},
		"AWS::RAM::ResourceShare":                                     &AWSRAMResourceShare{},
		"AWS::RDS::DBCluster":                                         &AWSRDSDBCluster{},
		"AWS::RDS::DBClusterParameterGroup":                           &AWSRDSDBClusterParameterGroup{},
		"AWS::RDS::DBInstance":                                        &AWSRDSDBInstance{},
		"AWS::RDS::DBParameterGroup":                                  &AWSRDSDBParameterGroup{},
		"AWS::RDS::DBSecurityGroup":                                   &AWSRDSDBSecurityGroup{},
		"AWS::RDS::DBSecurityGroupIngress":                            &AWSRDSDBSecurityGroupIngress{},
		"AWS::RDS::DBSubnetGroup":                                     &AWSRDSDBSubnetGroup{},
		"AWS::RDS::EventSubscription":                                 &AWSRDSEventSubscription{},
		"AWS::RDS::OptionGroup":                                       &AWSRDSOptionGroup{},
		"AWS::Redshift::Cluster":                                      &AWSRedshiftCluster{},
		"AWS::Redshift::ClusterParameterGroup":                        &AWSRedshiftClusterParameterGroup{},
		"AWS::Redshift::ClusterSecurityGroup":                         &AWSRedshiftClusterSecurityGroup{},
		"AWS::Redshift::ClusterSecurityGroupIngress":                  &AWSRedshiftClusterSecurityGroupIngress{},
		"AWS::Redshift::ClusterSubnetGroup":                           &AWSRedshiftClusterSubnetGroup{},
		"AWS::RoboMaker::Fleet":                                       &AWSRoboMakerFleet{},
		"AWS::RoboMaker::Robot":                                       &AWSRoboMakerRobot{},
		"AWS::RoboMaker::RobotApplication":                            &AWSRoboMakerRobotApplication{},
		"AWS::RoboMaker::RobotApplicationVersion":                     &AWSRoboMakerRobotApplicationVersion{},
		"AWS::RoboMaker::SimulationApplication":                       &AWSRoboMakerSimulationApplication{},
		"AWS::RoboMaker::SimulationApplicationVersion":                &AWSRoboMakerSimulationApplicationVersion{},
		"AWS::Route53::HealthCheck":                                   &AWSRoute53HealthCheck{},
		"AWS::Route53::HostedZone":                                    &AWSRoute53HostedZone{},
		"AWS::Route53::RecordSet":                                     &AWSRoute53RecordSet{},
		"AWS::Route53::RecordSetGroup":                                &AWSRoute53RecordSetGroup{},
		"AWS::Route53Resolver::ResolverEndpoint":                      &AWSRoute53ResolverResolverEndpoint{},
		"AWS::Route53Resolver::ResolverRule":                          &AWSRoute53ResolverResolverRule{},
		"AWS::Route53Resolver::ResolverRuleAssociation":               &AWSRoute53ResolverResolverRuleAssociation{},
		"AWS::S3::Bucket":                                             &AWSS3Bucket{},
		"AWS::S3::BucketPolicy":                                       &AWSS3BucketPolicy{},
		"AWS::SDB::Domain":                                            &AWSSDBDomain{},
		"AWS::SES::ConfigurationSet":                                  &AWSSESConfigurationSet{},
		"AWS::SES::ConfigurationSetEventDestination":                  &AWSSESConfigurationSetEventDestination{},
		"AWS::SES::ReceiptFilter":                                     &AWSSESReceiptFilter{},
		"AWS::SES::ReceiptRule":                                       &AWSSESReceiptRule{},
		"AWS::SES::ReceiptRuleSet":                                    &AWSSESReceiptRuleSet{},
		"AWS::SES::Template":                                          &AWSSESTemplate{},
		"AWS::SNS::Subscription":                                      &AWSSNSSubscription{},
		"AWS::SNS::Topic":                                             &AWSSNSTopic{},
		"AWS::SNS::TopicPolicy":                                       &AWSSNSTopicPolicy{},
		"AWS::SQS::Queue":                                             &AWSSQSQueue{},
		"AWS::SQS::QueuePolicy":                                       &AWSSQSQueuePolicy{},
		"AWS::SSM::Association":                                       &AWSSSMAssociation{},
		"AWS::SSM::Document":                                          &AWSSSMDocument{},
		"AWS::SSM::MaintenanceWindow":                                 &AWSSSMMaintenanceWindow{},
		"AWS::SSM::MaintenanceWindowTask":                             &AWSSSMMaintenanceWindowTask{},
		"AWS::SSM::Parameter":                                         &AWSSSMParameter{},
		"AWS::SSM::PatchBaseline":                                     &AWSSSMPatchBaseline{},
		"AWS::SSM::ResourceDataSync":                                  &AWSSSMResourceDataSync{},
		"AWS::SageMaker::Endpoint":                                    &AWSSageMakerEndpoint{},
		"AWS::SageMaker::EndpointConfig":                              &AWSSageMakerEndpointConfig{},
		"AWS::SageMaker::Model":                                       &AWSSageMakerModel{},
		"AWS::SageMaker::NotebookInstance":                            &AWSSageMakerNotebookInstance{},
		"AWS::SageMaker::NotebookInstanceLifecycleConfig":             &AWSSageMakerNotebookInstanceLifecycleConfig{},
		"AWS::SecretsManager::ResourcePolicy":                         &AWSSecretsManagerResourcePolicy{},
		"AWS::SecretsManager::RotationSchedule":                       &AWSSecretsManagerRotationSchedule{},
		"AWS::SecretsManager::Secret":                                 &AWSSecretsManagerSecret{},
		"AWS::SecretsManager::SecretTargetAttachment":                 &AWSSecretsManagerSecretTargetAttachment{},
		"AWS::Serverless::Api":                                        &AWSServerlessApi{},
		"AWS::Serverless::Application":                                &AWSServerlessApplication{},
		"AWS::Serverless::Function":                                   &AWSServerlessFunction{},
		"AWS::Serverless::LayerVersion":                               &AWSServerlessLayerVersion{},
		"AWS::Serverless::SimpleTable":                                &AWSServerlessSimpleTable{},
		"AWS::ServiceCatalog::AcceptedPortfolioShare":                 &AWSServiceCatalogAcceptedPortfolioShare{},
		"AWS::ServiceCatalog::CloudFormationProduct":                  &AWSServiceCatalogCloudFormationProduct{},
		"AWS::ServiceCatalog::CloudFormationProvisionedProduct":       &AWSServiceCatalogCloudFormationProvisionedProduct{},
		"AWS::ServiceCatalog::LaunchNotificationConstraint":           &AWSServiceCatalogLaunchNotificationConstraint{},
		"AWS::ServiceCatalog::LaunchRoleConstraint":                   &AWSServiceCatalogLaunchRoleConstraint{},
		"AWS::ServiceCatalog::LaunchTemplateConstraint":               &AWSServiceCatalogLaunchTemplateConstraint{},
		"AWS::ServiceCatalog::Portfolio":                              &AWSServiceCatalogPortfolio{},
		"AWS::ServiceCatalog::PortfolioPrincipalAssociation":          &AWSServiceCatalogPortfolioPrincipalAssociation{},
		"AWS::ServiceCatalog::PortfolioProductAssociation":            &AWSServiceCatalogPortfolioProductAssociation{},
		"AWS::ServiceCatalog::PortfolioShare":                         &AWSServiceCatalogPortfolioShare{},
		"AWS::ServiceCatalog::TagOption":                              &AWSServiceCatalogTagOption{},
		"AWS::ServiceCatalog::TagOptionAssociation":                   &AWSServiceCatalogTagOptionAssociation{},
		"AWS::ServiceDiscovery::HttpNamespace":                        &AWSServiceDiscoveryHttpNamespace{},
		"AWS::ServiceDiscovery::Instance":                             &AWSServiceDiscoveryInstance{},
		"AWS::ServiceDiscovery::PrivateDnsNamespace":                  &AWSServiceDiscoveryPrivateDnsNamespace{},
		"AWS::ServiceDiscovery::PublicDnsNamespace":                   &AWSServiceDiscoveryPublicDnsNamespace{},
		"AWS::ServiceDiscovery::Service":                              &AWSServiceDiscoveryService{},
		"AWS::StepFunctions::Activity":                                &AWSStepFunctionsActivity{},
		"AWS::StepFunctions::StateMachine":                            &AWSStepFunctionsStateMachine{},
		"AWS::WAF::ByteMatchSet":                                      &AWSWAFByteMatchSet{},
		"AWS::WAF::IPSet":                                             &AWSWAFIPSet{},
		"AWS::WAF::Rule":                                              &AWSWAFRule{},
		"AWS::WAF::SizeConstraintSet":                                 &AWSWAFSizeConstraintSet{},
		"AWS::WAF::SqlInjectionMatchSet":                              &AWSWAFSqlInjectionMatchSet{},
		"AWS::WAF::WebACL":                                            &AWSWAFWebACL{},
		"AWS::WAF::XssMatchSet":                                       &AWSWAFXssMatchSet{},
		"AWS::WAFRegional::ByteMatchSet":                              &AWSWAFRegionalByteMatchSet{},
		"AWS::WAFRegional::IPSet":                                     &AWSWAFRegionalIPSet{},
		"AWS::WAFRegional::Rule":                                      &AWSWAFRegionalRule{},
		"AWS::WAFRegional::SizeConstraintSet":                         &AWSWAFRegionalSizeConstraintSet{},
		"AWS::WAFRegional::SqlInjectionMatchSet":                      &AWSWAFRegionalSqlInjectionMatchSet{},
		"AWS::WAFRegional::WebACL":                                    &AWSWAFRegionalWebACL{},
		"AWS::WAFRegional::WebACLAssociation":                         &AWSWAFRegionalWebACLAssociation{},
		"AWS::WAFRegional::XssMatchSet":                               &AWSWAFRegionalXssMatchSet{},
		"AWS::WorkSpaces::Workspace":                                  &AWSWorkSpacesWorkspace{},
		"Alexa::ASK::Skill":                                           &AlexaASKSkill{},
	}
}

// GetAllAWSAmazonMQBrokerResources retrieves all AWSAmazonMQBroker items from an AWS CloudFormation template
func (t *Template) GetAllAWSAmazonMQBrokerResources() map[string]*AWSAmazonMQBroker {
	results := map[string]*AWSAmazonMQBroker{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSAmazonMQBroker:
			results[name] = resource
		}
	}
	return results
}

// GetAWSAmazonMQBrokerWithName retrieves all AWSAmazonMQBroker items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSAmazonMQBrokerWithName(name string) (*AWSAmazonMQBroker, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSAmazonMQBroker:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSAmazonMQBroker not found", name)
}

// GetAllAWSAmazonMQConfigurationResources retrieves all AWSAmazonMQConfiguration items from an AWS CloudFormation template
func (t *Template) GetAllAWSAmazonMQConfigurationResources() map[string]*AWSAmazonMQConfiguration {
	results := map[string]*AWSAmazonMQConfiguration{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSAmazonMQConfiguration:
			results[name] = resource
		}
	}
	return results
}

// GetAWSAmazonMQConfigurationWithName retrieves all AWSAmazonMQConfiguration items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSAmazonMQConfigurationWithName(name string) (*AWSAmazonMQConfiguration, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSAmazonMQConfiguration:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSAmazonMQConfiguration not found", name)
}

// GetAllAWSAmazonMQConfigurationAssociationResources retrieves all AWSAmazonMQConfigurationAssociation items from an AWS CloudFormation template
func (t *Template) GetAllAWSAmazonMQConfigurationAssociationResources() map[string]*AWSAmazonMQConfigurationAssociation {
	results := map[string]*AWSAmazonMQConfigurationAssociation{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSAmazonMQConfigurationAssociation:
			results[name] = resource
		}
	}
	return results
}

// GetAWSAmazonMQConfigurationAssociationWithName retrieves all AWSAmazonMQConfigurationAssociation items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSAmazonMQConfigurationAssociationWithName(name string) (*AWSAmazonMQConfigurationAssociation, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSAmazonMQConfigurationAssociation:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSAmazonMQConfigurationAssociation not found", name)
}

// GetAllAWSApiGatewayAccountResources retrieves all AWSApiGatewayAccount items from an AWS CloudFormation template
func (t *Template) GetAllAWSApiGatewayAccountResources() map[string]*AWSApiGatewayAccount {
	results := map[string]*AWSApiGatewayAccount{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSApiGatewayAccount:
			results[name] = resource
		}
	}
	return results
}

// GetAWSApiGatewayAccountWithName retrieves all AWSApiGatewayAccount items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApiGatewayAccountWithName(name string) (*AWSApiGatewayAccount, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSApiGatewayAccount:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSApiGatewayAccount not found", name)
}

// GetAllAWSApiGatewayApiKeyResources retrieves all AWSApiGatewayApiKey items from an AWS CloudFormation template
func (t *Template) GetAllAWSApiGatewayApiKeyResources() map[string]*AWSApiGatewayApiKey {
	results := map[string]*AWSApiGatewayApiKey{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSApiGatewayApiKey:
			results[name] = resource
		}
	}
	return results
}

// GetAWSApiGatewayApiKeyWithName retrieves all AWSApiGatewayApiKey items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApiGatewayApiKeyWithName(name string) (*AWSApiGatewayApiKey, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSApiGatewayApiKey:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSApiGatewayApiKey not found", name)
}

// GetAllAWSApiGatewayAuthorizerResources retrieves all AWSApiGatewayAuthorizer items from an AWS CloudFormation template
func (t *Template) GetAllAWSApiGatewayAuthorizerResources() map[string]*AWSApiGatewayAuthorizer {
	results := map[string]*AWSApiGatewayAuthorizer{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSApiGatewayAuthorizer:
			results[name] = resource
		}
	}
	return results
}

// GetAWSApiGatewayAuthorizerWithName retrieves all AWSApiGatewayAuthorizer items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApiGatewayAuthorizerWithName(name string) (*AWSApiGatewayAuthorizer, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSApiGatewayAuthorizer:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSApiGatewayAuthorizer not found", name)
}

// GetAllAWSApiGatewayBasePathMappingResources retrieves all AWSApiGatewayBasePathMapping items from an AWS CloudFormation template
func (t *Template) GetAllAWSApiGatewayBasePathMappingResources() map[string]*AWSApiGatewayBasePathMapping {
	results := map[string]*AWSApiGatewayBasePathMapping{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSApiGatewayBasePathMapping:
			results[name] = resource
		}
	}
	return results
}

// GetAWSApiGatewayBasePathMappingWithName retrieves all AWSApiGatewayBasePathMapping items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApiGatewayBasePathMappingWithName(name string) (*AWSApiGatewayBasePathMapping, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSApiGatewayBasePathMapping:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSApiGatewayBasePathMapping not found", name)
}

// GetAllAWSApiGatewayClientCertificateResources retrieves all AWSApiGatewayClientCertificate items from an AWS CloudFormation template
func (t *Template) GetAllAWSApiGatewayClientCertificateResources() map[string]*AWSApiGatewayClientCertificate {
	results := map[string]*AWSApiGatewayClientCertificate{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSApiGatewayClientCertificate:
			results[name] = resource
		}
	}
	return results
}

// GetAWSApiGatewayClientCertificateWithName retrieves all AWSApiGatewayClientCertificate items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApiGatewayClientCertificateWithName(name string) (*AWSApiGatewayClientCertificate, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSApiGatewayClientCertificate:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSApiGatewayClientCertificate not found", name)
}

// GetAllAWSApiGatewayDeploymentResources retrieves all AWSApiGatewayDeployment items from an AWS CloudFormation template
func (t *Template) GetAllAWSApiGatewayDeploymentResources() map[string]*AWSApiGatewayDeployment {
	results := map[string]*AWSApiGatewayDeployment{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSApiGatewayDeployment:
			results[name] = resource
		}
	}
	return results
}

// GetAWSApiGatewayDeploymentWithName retrieves all AWSApiGatewayDeployment items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApiGatewayDeploymentWithName(name string) (*AWSApiGatewayDeployment, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSApiGatewayDeployment:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSApiGatewayDeployment not found", name)
}

// GetAllAWSApiGatewayDocumentationPartResources retrieves all AWSApiGatewayDocumentationPart items from an AWS CloudFormation template
func (t *Template) GetAllAWSApiGatewayDocumentationPartResources() map[string]*AWSApiGatewayDocumentationPart {
	results := map[string]*AWSApiGatewayDocumentationPart{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSApiGatewayDocumentationPart:
			results[name] = resource
		}
	}
	return results
}

// GetAWSApiGatewayDocumentationPartWithName retrieves all AWSApiGatewayDocumentationPart items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApiGatewayDocumentationPartWithName(name string) (*AWSApiGatewayDocumentationPart, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSApiGatewayDocumentationPart:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSApiGatewayDocumentationPart not found", name)
}

// GetAllAWSApiGatewayDocumentationVersionResources retrieves all AWSApiGatewayDocumentationVersion items from an AWS CloudFormation template
func (t *Template) GetAllAWSApiGatewayDocumentationVersionResources() map[string]*AWSApiGatewayDocumentationVersion {
	results := map[string]*AWSApiGatewayDocumentationVersion{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSApiGatewayDocumentationVersion:
			results[name] = resource
		}
	}
	return results
}

// GetAWSApiGatewayDocumentationVersionWithName retrieves all AWSApiGatewayDocumentationVersion items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApiGatewayDocumentationVersionWithName(name string) (*AWSApiGatewayDocumentationVersion, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSApiGatewayDocumentationVersion:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSApiGatewayDocumentationVersion not found", name)
}

// GetAllAWSApiGatewayDomainNameResources retrieves all AWSApiGatewayDomainName items from an AWS CloudFormation template
func (t *Template) GetAllAWSApiGatewayDomainNameResources() map[string]*AWSApiGatewayDomainName {
	results := map[string]*AWSApiGatewayDomainName{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSApiGatewayDomainName:
			results[name] = resource
		}
	}
	return results
}

// GetAWSApiGatewayDomainNameWithName retrieves all AWSApiGatewayDomainName items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApiGatewayDomainNameWithName(name string) (*AWSApiGatewayDomainName, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSApiGatewayDomainName:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSApiGatewayDomainName not found", name)
}

// GetAllAWSApiGatewayGatewayResponseResources retrieves all AWSApiGatewayGatewayResponse items from an AWS CloudFormation template
func (t *Template) GetAllAWSApiGatewayGatewayResponseResources() map[string]*AWSApiGatewayGatewayResponse {
	results := map[string]*AWSApiGatewayGatewayResponse{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSApiGatewayGatewayResponse:
			results[name] = resource
		}
	}
	return results
}

// GetAWSApiGatewayGatewayResponseWithName retrieves all AWSApiGatewayGatewayResponse items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApiGatewayGatewayResponseWithName(name string) (*AWSApiGatewayGatewayResponse, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSApiGatewayGatewayResponse:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSApiGatewayGatewayResponse not found", name)
}

// GetAllAWSApiGatewayMethodResources retrieves all AWSApiGatewayMethod items from an AWS CloudFormation template
func (t *Template) GetAllAWSApiGatewayMethodResources() map[string]*AWSApiGatewayMethod {
	results := map[string]*AWSApiGatewayMethod{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSApiGatewayMethod:
			results[name] = resource
		}
	}
	return results
}

// GetAWSApiGatewayMethodWithName retrieves all AWSApiGatewayMethod items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApiGatewayMethodWithName(name string) (*AWSApiGatewayMethod, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSApiGatewayMethod:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSApiGatewayMethod not found", name)
}

// GetAllAWSApiGatewayModelResources retrieves all AWSApiGatewayModel items from an AWS CloudFormation template
func (t *Template) GetAllAWSApiGatewayModelResources() map[string]*AWSApiGatewayModel {
	results := map[string]*AWSApiGatewayModel{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSApiGatewayModel:
			results[name] = resource
		}
	}
	return results
}

// GetAWSApiGatewayModelWithName retrieves all AWSApiGatewayModel items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApiGatewayModelWithName(name string) (*AWSApiGatewayModel, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSApiGatewayModel:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSApiGatewayModel not found", name)
}

// GetAllAWSApiGatewayRequestValidatorResources retrieves all AWSApiGatewayRequestValidator items from an AWS CloudFormation template
func (t *Template) GetAllAWSApiGatewayRequestValidatorResources() map[string]*AWSApiGatewayRequestValidator {
	results := map[string]*AWSApiGatewayRequestValidator{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSApiGatewayRequestValidator:
			results[name] = resource
		}
	}
	return results
}

// GetAWSApiGatewayRequestValidatorWithName retrieves all AWSApiGatewayRequestValidator items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApiGatewayRequestValidatorWithName(name string) (*AWSApiGatewayRequestValidator, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSApiGatewayRequestValidator:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSApiGatewayRequestValidator not found", name)
}

// GetAllAWSApiGatewayResourceResources retrieves all AWSApiGatewayResource items from an AWS CloudFormation template
func (t *Template) GetAllAWSApiGatewayResourceResources() map[string]*AWSApiGatewayResource {
	results := map[string]*AWSApiGatewayResource{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSApiGatewayResource:
			results[name] = resource
		}
	}
	return results
}

// GetAWSApiGatewayResourceWithName retrieves all AWSApiGatewayResource items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApiGatewayResourceWithName(name string) (*AWSApiGatewayResource, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSApiGatewayResource:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSApiGatewayResource not found", name)
}

// GetAllAWSApiGatewayRestApiResources retrieves all AWSApiGatewayRestApi items from an AWS CloudFormation template
func (t *Template) GetAllAWSApiGatewayRestApiResources() map[string]*AWSApiGatewayRestApi {
	results := map[string]*AWSApiGatewayRestApi{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSApiGatewayRestApi:
			results[name] = resource
		}
	}
	return results
}

// GetAWSApiGatewayRestApiWithName retrieves all AWSApiGatewayRestApi items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApiGatewayRestApiWithName(name string) (*AWSApiGatewayRestApi, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSApiGatewayRestApi:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSApiGatewayRestApi not found", name)
}

// GetAllAWSApiGatewayStageResources retrieves all AWSApiGatewayStage items from an AWS CloudFormation template
func (t *Template) GetAllAWSApiGatewayStageResources() map[string]*AWSApiGatewayStage {
	results := map[string]*AWSApiGatewayStage{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSApiGatewayStage:
			results[name] = resource
		}
	}
	return results
}

// GetAWSApiGatewayStageWithName retrieves all AWSApiGatewayStage items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApiGatewayStageWithName(name string) (*AWSApiGatewayStage, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSApiGatewayStage:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSApiGatewayStage not found", name)
}

// GetAllAWSApiGatewayUsagePlanResources retrieves all AWSApiGatewayUsagePlan items from an AWS CloudFormation template
func (t *Template) GetAllAWSApiGatewayUsagePlanResources() map[string]*AWSApiGatewayUsagePlan {
	results := map[string]*AWSApiGatewayUsagePlan{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSApiGatewayUsagePlan:
			results[name] = resource
		}
	}
	return results
}

// GetAWSApiGatewayUsagePlanWithName retrieves all AWSApiGatewayUsagePlan items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApiGatewayUsagePlanWithName(name string) (*AWSApiGatewayUsagePlan, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSApiGatewayUsagePlan:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSApiGatewayUsagePlan not found", name)
}

// GetAllAWSApiGatewayUsagePlanKeyResources retrieves all AWSApiGatewayUsagePlanKey items from an AWS CloudFormation template
func (t *Template) GetAllAWSApiGatewayUsagePlanKeyResources() map[string]*AWSApiGatewayUsagePlanKey {
	results := map[string]*AWSApiGatewayUsagePlanKey{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSApiGatewayUsagePlanKey:
			results[name] = resource
		}
	}
	return results
}

// GetAWSApiGatewayUsagePlanKeyWithName retrieves all AWSApiGatewayUsagePlanKey items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApiGatewayUsagePlanKeyWithName(name string) (*AWSApiGatewayUsagePlanKey, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSApiGatewayUsagePlanKey:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSApiGatewayUsagePlanKey not found", name)
}

// GetAllAWSApiGatewayVpcLinkResources retrieves all AWSApiGatewayVpcLink items from an AWS CloudFormation template
func (t *Template) GetAllAWSApiGatewayVpcLinkResources() map[string]*AWSApiGatewayVpcLink {
	results := map[string]*AWSApiGatewayVpcLink{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSApiGatewayVpcLink:
			results[name] = resource
		}
	}
	return results
}

// GetAWSApiGatewayVpcLinkWithName retrieves all AWSApiGatewayVpcLink items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApiGatewayVpcLinkWithName(name string) (*AWSApiGatewayVpcLink, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSApiGatewayVpcLink:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSApiGatewayVpcLink not found", name)
}

// GetAllAWSApiGatewayV2ApiResources retrieves all AWSApiGatewayV2Api items from an AWS CloudFormation template
func (t *Template) GetAllAWSApiGatewayV2ApiResources() map[string]*AWSApiGatewayV2Api {
	results := map[string]*AWSApiGatewayV2Api{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSApiGatewayV2Api:
			results[name] = resource
		}
	}
	return results
}

// GetAWSApiGatewayV2ApiWithName retrieves all AWSApiGatewayV2Api items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApiGatewayV2ApiWithName(name string) (*AWSApiGatewayV2Api, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSApiGatewayV2Api:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSApiGatewayV2Api not found", name)
}

// GetAllAWSApiGatewayV2AuthorizerResources retrieves all AWSApiGatewayV2Authorizer items from an AWS CloudFormation template
func (t *Template) GetAllAWSApiGatewayV2AuthorizerResources() map[string]*AWSApiGatewayV2Authorizer {
	results := map[string]*AWSApiGatewayV2Authorizer{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSApiGatewayV2Authorizer:
			results[name] = resource
		}
	}
	return results
}

// GetAWSApiGatewayV2AuthorizerWithName retrieves all AWSApiGatewayV2Authorizer items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApiGatewayV2AuthorizerWithName(name string) (*AWSApiGatewayV2Authorizer, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSApiGatewayV2Authorizer:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSApiGatewayV2Authorizer not found", name)
}

// GetAllAWSApiGatewayV2DeploymentResources retrieves all AWSApiGatewayV2Deployment items from an AWS CloudFormation template
func (t *Template) GetAllAWSApiGatewayV2DeploymentResources() map[string]*AWSApiGatewayV2Deployment {
	results := map[string]*AWSApiGatewayV2Deployment{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSApiGatewayV2Deployment:
			results[name] = resource
		}
	}
	return results
}

// GetAWSApiGatewayV2DeploymentWithName retrieves all AWSApiGatewayV2Deployment items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApiGatewayV2DeploymentWithName(name string) (*AWSApiGatewayV2Deployment, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSApiGatewayV2Deployment:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSApiGatewayV2Deployment not found", name)
}

// GetAllAWSApiGatewayV2IntegrationResources retrieves all AWSApiGatewayV2Integration items from an AWS CloudFormation template
func (t *Template) GetAllAWSApiGatewayV2IntegrationResources() map[string]*AWSApiGatewayV2Integration {
	results := map[string]*AWSApiGatewayV2Integration{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSApiGatewayV2Integration:
			results[name] = resource
		}
	}
	return results
}

// GetAWSApiGatewayV2IntegrationWithName retrieves all AWSApiGatewayV2Integration items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApiGatewayV2IntegrationWithName(name string) (*AWSApiGatewayV2Integration, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSApiGatewayV2Integration:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSApiGatewayV2Integration not found", name)
}

// GetAllAWSApiGatewayV2IntegrationResponseResources retrieves all AWSApiGatewayV2IntegrationResponse items from an AWS CloudFormation template
func (t *Template) GetAllAWSApiGatewayV2IntegrationResponseResources() map[string]*AWSApiGatewayV2IntegrationResponse {
	results := map[string]*AWSApiGatewayV2IntegrationResponse{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSApiGatewayV2IntegrationResponse:
			results[name] = resource
		}
	}
	return results
}

// GetAWSApiGatewayV2IntegrationResponseWithName retrieves all AWSApiGatewayV2IntegrationResponse items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApiGatewayV2IntegrationResponseWithName(name string) (*AWSApiGatewayV2IntegrationResponse, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSApiGatewayV2IntegrationResponse:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSApiGatewayV2IntegrationResponse not found", name)
}

// GetAllAWSApiGatewayV2ModelResources retrieves all AWSApiGatewayV2Model items from an AWS CloudFormation template
func (t *Template) GetAllAWSApiGatewayV2ModelResources() map[string]*AWSApiGatewayV2Model {
	results := map[string]*AWSApiGatewayV2Model{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSApiGatewayV2Model:
			results[name] = resource
		}
	}
	return results
}

// GetAWSApiGatewayV2ModelWithName retrieves all AWSApiGatewayV2Model items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApiGatewayV2ModelWithName(name string) (*AWSApiGatewayV2Model, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSApiGatewayV2Model:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSApiGatewayV2Model not found", name)
}

// GetAllAWSApiGatewayV2RouteResources retrieves all AWSApiGatewayV2Route items from an AWS CloudFormation template
func (t *Template) GetAllAWSApiGatewayV2RouteResources() map[string]*AWSApiGatewayV2Route {
	results := map[string]*AWSApiGatewayV2Route{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSApiGatewayV2Route:
			results[name] = resource
		}
	}
	return results
}

// GetAWSApiGatewayV2RouteWithName retrieves all AWSApiGatewayV2Route items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApiGatewayV2RouteWithName(name string) (*AWSApiGatewayV2Route, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSApiGatewayV2Route:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSApiGatewayV2Route not found", name)
}

// GetAllAWSApiGatewayV2RouteResponseResources retrieves all AWSApiGatewayV2RouteResponse items from an AWS CloudFormation template
func (t *Template) GetAllAWSApiGatewayV2RouteResponseResources() map[string]*AWSApiGatewayV2RouteResponse {
	results := map[string]*AWSApiGatewayV2RouteResponse{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSApiGatewayV2RouteResponse:
			results[name] = resource
		}
	}
	return results
}

// GetAWSApiGatewayV2RouteResponseWithName retrieves all AWSApiGatewayV2RouteResponse items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApiGatewayV2RouteResponseWithName(name string) (*AWSApiGatewayV2RouteResponse, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSApiGatewayV2RouteResponse:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSApiGatewayV2RouteResponse not found", name)
}

// GetAllAWSApiGatewayV2StageResources retrieves all AWSApiGatewayV2Stage items from an AWS CloudFormation template
func (t *Template) GetAllAWSApiGatewayV2StageResources() map[string]*AWSApiGatewayV2Stage {
	results := map[string]*AWSApiGatewayV2Stage{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSApiGatewayV2Stage:
			results[name] = resource
		}
	}
	return results
}

// GetAWSApiGatewayV2StageWithName retrieves all AWSApiGatewayV2Stage items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApiGatewayV2StageWithName(name string) (*AWSApiGatewayV2Stage, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSApiGatewayV2Stage:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSApiGatewayV2Stage not found", name)
}

// GetAllAWSAppStreamDirectoryConfigResources retrieves all AWSAppStreamDirectoryConfig items from an AWS CloudFormation template
func (t *Template) GetAllAWSAppStreamDirectoryConfigResources() map[string]*AWSAppStreamDirectoryConfig {
	results := map[string]*AWSAppStreamDirectoryConfig{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSAppStreamDirectoryConfig:
			results[name] = resource
		}
	}
	return results
}

// GetAWSAppStreamDirectoryConfigWithName retrieves all AWSAppStreamDirectoryConfig items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSAppStreamDirectoryConfigWithName(name string) (*AWSAppStreamDirectoryConfig, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSAppStreamDirectoryConfig:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSAppStreamDirectoryConfig not found", name)
}

// GetAllAWSAppStreamFleetResources retrieves all AWSAppStreamFleet items from an AWS CloudFormation template
func (t *Template) GetAllAWSAppStreamFleetResources() map[string]*AWSAppStreamFleet {
	results := map[string]*AWSAppStreamFleet{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSAppStreamFleet:
			results[name] = resource
		}
	}
	return results
}

// GetAWSAppStreamFleetWithName retrieves all AWSAppStreamFleet items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSAppStreamFleetWithName(name string) (*AWSAppStreamFleet, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSAppStreamFleet:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSAppStreamFleet not found", name)
}

// GetAllAWSAppStreamImageBuilderResources retrieves all AWSAppStreamImageBuilder items from an AWS CloudFormation template
func (t *Template) GetAllAWSAppStreamImageBuilderResources() map[string]*AWSAppStreamImageBuilder {
	results := map[string]*AWSAppStreamImageBuilder{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSAppStreamImageBuilder:
			results[name] = resource
		}
	}
	return results
}

// GetAWSAppStreamImageBuilderWithName retrieves all AWSAppStreamImageBuilder items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSAppStreamImageBuilderWithName(name string) (*AWSAppStreamImageBuilder, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSAppStreamImageBuilder:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSAppStreamImageBuilder not found", name)
}

// GetAllAWSAppStreamStackResources retrieves all AWSAppStreamStack items from an AWS CloudFormation template
func (t *Template) GetAllAWSAppStreamStackResources() map[string]*AWSAppStreamStack {
	results := map[string]*AWSAppStreamStack{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSAppStreamStack:
			results[name] = resource
		}
	}
	return results
}

// GetAWSAppStreamStackWithName retrieves all AWSAppStreamStack items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSAppStreamStackWithName(name string) (*AWSAppStreamStack, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSAppStreamStack:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSAppStreamStack not found", name)
}

// GetAllAWSAppStreamStackFleetAssociationResources retrieves all AWSAppStreamStackFleetAssociation items from an AWS CloudFormation template
func (t *Template) GetAllAWSAppStreamStackFleetAssociationResources() map[string]*AWSAppStreamStackFleetAssociation {
	results := map[string]*AWSAppStreamStackFleetAssociation{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSAppStreamStackFleetAssociation:
			results[name] = resource
		}
	}
	return results
}

// GetAWSAppStreamStackFleetAssociationWithName retrieves all AWSAppStreamStackFleetAssociation items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSAppStreamStackFleetAssociationWithName(name string) (*AWSAppStreamStackFleetAssociation, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSAppStreamStackFleetAssociation:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSAppStreamStackFleetAssociation not found", name)
}

// GetAllAWSAppStreamStackUserAssociationResources retrieves all AWSAppStreamStackUserAssociation items from an AWS CloudFormation template
func (t *Template) GetAllAWSAppStreamStackUserAssociationResources() map[string]*AWSAppStreamStackUserAssociation {
	results := map[string]*AWSAppStreamStackUserAssociation{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSAppStreamStackUserAssociation:
			results[name] = resource
		}
	}
	return results
}

// GetAWSAppStreamStackUserAssociationWithName retrieves all AWSAppStreamStackUserAssociation items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSAppStreamStackUserAssociationWithName(name string) (*AWSAppStreamStackUserAssociation, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSAppStreamStackUserAssociation:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSAppStreamStackUserAssociation not found", name)
}

// GetAllAWSAppStreamUserResources retrieves all AWSAppStreamUser items from an AWS CloudFormation template
func (t *Template) GetAllAWSAppStreamUserResources() map[string]*AWSAppStreamUser {
	results := map[string]*AWSAppStreamUser{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSAppStreamUser:
			results[name] = resource
		}
	}
	return results
}

// GetAWSAppStreamUserWithName retrieves all AWSAppStreamUser items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSAppStreamUserWithName(name string) (*AWSAppStreamUser, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSAppStreamUser:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSAppStreamUser not found", name)
}

// GetAllAWSAppSyncApiKeyResources retrieves all AWSAppSyncApiKey items from an AWS CloudFormation template
func (t *Template) GetAllAWSAppSyncApiKeyResources() map[string]*AWSAppSyncApiKey {
	results := map[string]*AWSAppSyncApiKey{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSAppSyncApiKey:
			results[name] = resource
		}
	}
	return results
}

// GetAWSAppSyncApiKeyWithName retrieves all AWSAppSyncApiKey items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSAppSyncApiKeyWithName(name string) (*AWSAppSyncApiKey, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSAppSyncApiKey:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSAppSyncApiKey not found", name)
}

// GetAllAWSAppSyncDataSourceResources retrieves all AWSAppSyncDataSource items from an AWS CloudFormation template
func (t *Template) GetAllAWSAppSyncDataSourceResources() map[string]*AWSAppSyncDataSource {
	results := map[string]*AWSAppSyncDataSource{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSAppSyncDataSource:
			results[name] = resource
		}
	}
	return results
}

// GetAWSAppSyncDataSourceWithName retrieves all AWSAppSyncDataSource items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSAppSyncDataSourceWithName(name string) (*AWSAppSyncDataSource, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSAppSyncDataSource:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSAppSyncDataSource not found", name)
}

// GetAllAWSAppSyncFunctionConfigurationResources retrieves all AWSAppSyncFunctionConfiguration items from an AWS CloudFormation template
func (t *Template) GetAllAWSAppSyncFunctionConfigurationResources() map[string]*AWSAppSyncFunctionConfiguration {
	results := map[string]*AWSAppSyncFunctionConfiguration{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSAppSyncFunctionConfiguration:
			results[name] = resource
		}
	}
	return results
}

// GetAWSAppSyncFunctionConfigurationWithName retrieves all AWSAppSyncFunctionConfiguration items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSAppSyncFunctionConfigurationWithName(name string) (*AWSAppSyncFunctionConfiguration, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSAppSyncFunctionConfiguration:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSAppSyncFunctionConfiguration not found", name)
}

// GetAllAWSAppSyncGraphQLApiResources retrieves all AWSAppSyncGraphQLApi items from an AWS CloudFormation template
func (t *Template) GetAllAWSAppSyncGraphQLApiResources() map[string]*AWSAppSyncGraphQLApi {
	results := map[string]*AWSAppSyncGraphQLApi{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSAppSyncGraphQLApi:
			results[name] = resource
		}
	}
	return results
}

// GetAWSAppSyncGraphQLApiWithName retrieves all AWSAppSyncGraphQLApi items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSAppSyncGraphQLApiWithName(name string) (*AWSAppSyncGraphQLApi, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSAppSyncGraphQLApi:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSAppSyncGraphQLApi not found", name)
}

// GetAllAWSAppSyncGraphQLSchemaResources retrieves all AWSAppSyncGraphQLSchema items from an AWS CloudFormation template
func (t *Template) GetAllAWSAppSyncGraphQLSchemaResources() map[string]*AWSAppSyncGraphQLSchema {
	results := map[string]*AWSAppSyncGraphQLSchema{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSAppSyncGraphQLSchema:
			results[name] = resource
		}
	}
	return results
}

// GetAWSAppSyncGraphQLSchemaWithName retrieves all AWSAppSyncGraphQLSchema items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSAppSyncGraphQLSchemaWithName(name string) (*AWSAppSyncGraphQLSchema, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSAppSyncGraphQLSchema:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSAppSyncGraphQLSchema not found", name)
}

// GetAllAWSAppSyncResolverResources retrieves all AWSAppSyncResolver items from an AWS CloudFormation template
func (t *Template) GetAllAWSAppSyncResolverResources() map[string]*AWSAppSyncResolver {
	results := map[string]*AWSAppSyncResolver{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSAppSyncResolver:
			results[name] = resource
		}
	}
	return results
}

// GetAWSAppSyncResolverWithName retrieves all AWSAppSyncResolver items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSAppSyncResolverWithName(name string) (*AWSAppSyncResolver, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSAppSyncResolver:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSAppSyncResolver not found", name)
}

// GetAllAWSApplicationAutoScalingScalableTargetResources retrieves all AWSApplicationAutoScalingScalableTarget items from an AWS CloudFormation template
func (t *Template) GetAllAWSApplicationAutoScalingScalableTargetResources() map[string]*AWSApplicationAutoScalingScalableTarget {
	results := map[string]*AWSApplicationAutoScalingScalableTarget{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSApplicationAutoScalingScalableTarget:
			results[name] = resource
		}
	}
	return results
}

// GetAWSApplicationAutoScalingScalableTargetWithName retrieves all AWSApplicationAutoScalingScalableTarget items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApplicationAutoScalingScalableTargetWithName(name string) (*AWSApplicationAutoScalingScalableTarget, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSApplicationAutoScalingScalableTarget:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSApplicationAutoScalingScalableTarget not found", name)
}

// GetAllAWSApplicationAutoScalingScalingPolicyResources retrieves all AWSApplicationAutoScalingScalingPolicy items from an AWS CloudFormation template
func (t *Template) GetAllAWSApplicationAutoScalingScalingPolicyResources() map[string]*AWSApplicationAutoScalingScalingPolicy {
	results := map[string]*AWSApplicationAutoScalingScalingPolicy{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSApplicationAutoScalingScalingPolicy:
			results[name] = resource
		}
	}
	return results
}

// GetAWSApplicationAutoScalingScalingPolicyWithName retrieves all AWSApplicationAutoScalingScalingPolicy items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApplicationAutoScalingScalingPolicyWithName(name string) (*AWSApplicationAutoScalingScalingPolicy, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSApplicationAutoScalingScalingPolicy:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSApplicationAutoScalingScalingPolicy not found", name)
}

// GetAllAWSAthenaNamedQueryResources retrieves all AWSAthenaNamedQuery items from an AWS CloudFormation template
func (t *Template) GetAllAWSAthenaNamedQueryResources() map[string]*AWSAthenaNamedQuery {
	results := map[string]*AWSAthenaNamedQuery{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSAthenaNamedQuery:
			results[name] = resource
		}
	}
	return results
}

// GetAWSAthenaNamedQueryWithName retrieves all AWSAthenaNamedQuery items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSAthenaNamedQueryWithName(name string) (*AWSAthenaNamedQuery, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSAthenaNamedQuery:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSAthenaNamedQuery not found", name)
}

// GetAllAWSAutoScalingAutoScalingGroupResources retrieves all AWSAutoScalingAutoScalingGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSAutoScalingAutoScalingGroupResources() map[string]*AWSAutoScalingAutoScalingGroup {
	results := map[string]*AWSAutoScalingAutoScalingGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSAutoScalingAutoScalingGroup:
			results[name] = resource
		}
	}
	return results
}

// GetAWSAutoScalingAutoScalingGroupWithName retrieves all AWSAutoScalingAutoScalingGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSAutoScalingAutoScalingGroupWithName(name string) (*AWSAutoScalingAutoScalingGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSAutoScalingAutoScalingGroup:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSAutoScalingAutoScalingGroup not found", name)
}

// GetAllAWSAutoScalingLaunchConfigurationResources retrieves all AWSAutoScalingLaunchConfiguration items from an AWS CloudFormation template
func (t *Template) GetAllAWSAutoScalingLaunchConfigurationResources() map[string]*AWSAutoScalingLaunchConfiguration {
	results := map[string]*AWSAutoScalingLaunchConfiguration{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSAutoScalingLaunchConfiguration:
			results[name] = resource
		}
	}
	return results
}

// GetAWSAutoScalingLaunchConfigurationWithName retrieves all AWSAutoScalingLaunchConfiguration items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSAutoScalingLaunchConfigurationWithName(name string) (*AWSAutoScalingLaunchConfiguration, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSAutoScalingLaunchConfiguration:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSAutoScalingLaunchConfiguration not found", name)
}

// GetAllAWSAutoScalingLifecycleHookResources retrieves all AWSAutoScalingLifecycleHook items from an AWS CloudFormation template
func (t *Template) GetAllAWSAutoScalingLifecycleHookResources() map[string]*AWSAutoScalingLifecycleHook {
	results := map[string]*AWSAutoScalingLifecycleHook{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSAutoScalingLifecycleHook:
			results[name] = resource
		}
	}
	return results
}

// GetAWSAutoScalingLifecycleHookWithName retrieves all AWSAutoScalingLifecycleHook items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSAutoScalingLifecycleHookWithName(name string) (*AWSAutoScalingLifecycleHook, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSAutoScalingLifecycleHook:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSAutoScalingLifecycleHook not found", name)
}

// GetAllAWSAutoScalingScalingPolicyResources retrieves all AWSAutoScalingScalingPolicy items from an AWS CloudFormation template
func (t *Template) GetAllAWSAutoScalingScalingPolicyResources() map[string]*AWSAutoScalingScalingPolicy {
	results := map[string]*AWSAutoScalingScalingPolicy{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSAutoScalingScalingPolicy:
			results[name] = resource
		}
	}
	return results
}

// GetAWSAutoScalingScalingPolicyWithName retrieves all AWSAutoScalingScalingPolicy items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSAutoScalingScalingPolicyWithName(name string) (*AWSAutoScalingScalingPolicy, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSAutoScalingScalingPolicy:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSAutoScalingScalingPolicy not found", name)
}

// GetAllAWSAutoScalingScheduledActionResources retrieves all AWSAutoScalingScheduledAction items from an AWS CloudFormation template
func (t *Template) GetAllAWSAutoScalingScheduledActionResources() map[string]*AWSAutoScalingScheduledAction {
	results := map[string]*AWSAutoScalingScheduledAction{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSAutoScalingScheduledAction:
			results[name] = resource
		}
	}
	return results
}

// GetAWSAutoScalingScheduledActionWithName retrieves all AWSAutoScalingScheduledAction items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSAutoScalingScheduledActionWithName(name string) (*AWSAutoScalingScheduledAction, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSAutoScalingScheduledAction:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSAutoScalingScheduledAction not found", name)
}

// GetAllAWSAutoScalingPlansScalingPlanResources retrieves all AWSAutoScalingPlansScalingPlan items from an AWS CloudFormation template
func (t *Template) GetAllAWSAutoScalingPlansScalingPlanResources() map[string]*AWSAutoScalingPlansScalingPlan {
	results := map[string]*AWSAutoScalingPlansScalingPlan{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSAutoScalingPlansScalingPlan:
			results[name] = resource
		}
	}
	return results
}

// GetAWSAutoScalingPlansScalingPlanWithName retrieves all AWSAutoScalingPlansScalingPlan items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSAutoScalingPlansScalingPlanWithName(name string) (*AWSAutoScalingPlansScalingPlan, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSAutoScalingPlansScalingPlan:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSAutoScalingPlansScalingPlan not found", name)
}

// GetAllAWSBatchComputeEnvironmentResources retrieves all AWSBatchComputeEnvironment items from an AWS CloudFormation template
func (t *Template) GetAllAWSBatchComputeEnvironmentResources() map[string]*AWSBatchComputeEnvironment {
	results := map[string]*AWSBatchComputeEnvironment{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSBatchComputeEnvironment:
			results[name] = resource
		}
	}
	return results
}

// GetAWSBatchComputeEnvironmentWithName retrieves all AWSBatchComputeEnvironment items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSBatchComputeEnvironmentWithName(name string) (*AWSBatchComputeEnvironment, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSBatchComputeEnvironment:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSBatchComputeEnvironment not found", name)
}

// GetAllAWSBatchJobDefinitionResources retrieves all AWSBatchJobDefinition items from an AWS CloudFormation template
func (t *Template) GetAllAWSBatchJobDefinitionResources() map[string]*AWSBatchJobDefinition {
	results := map[string]*AWSBatchJobDefinition{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSBatchJobDefinition:
			results[name] = resource
		}
	}
	return results
}

// GetAWSBatchJobDefinitionWithName retrieves all AWSBatchJobDefinition items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSBatchJobDefinitionWithName(name string) (*AWSBatchJobDefinition, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSBatchJobDefinition:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSBatchJobDefinition not found", name)
}

// GetAllAWSBatchJobQueueResources retrieves all AWSBatchJobQueue items from an AWS CloudFormation template
func (t *Template) GetAllAWSBatchJobQueueResources() map[string]*AWSBatchJobQueue {
	results := map[string]*AWSBatchJobQueue{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSBatchJobQueue:
			results[name] = resource
		}
	}
	return results
}

// GetAWSBatchJobQueueWithName retrieves all AWSBatchJobQueue items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSBatchJobQueueWithName(name string) (*AWSBatchJobQueue, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSBatchJobQueue:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSBatchJobQueue not found", name)
}

// GetAllAWSBudgetsBudgetResources retrieves all AWSBudgetsBudget items from an AWS CloudFormation template
func (t *Template) GetAllAWSBudgetsBudgetResources() map[string]*AWSBudgetsBudget {
	results := map[string]*AWSBudgetsBudget{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSBudgetsBudget:
			results[name] = resource
		}
	}
	return results
}

// GetAWSBudgetsBudgetWithName retrieves all AWSBudgetsBudget items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSBudgetsBudgetWithName(name string) (*AWSBudgetsBudget, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSBudgetsBudget:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSBudgetsBudget not found", name)
}

// GetAllAWSCertificateManagerCertificateResources retrieves all AWSCertificateManagerCertificate items from an AWS CloudFormation template
func (t *Template) GetAllAWSCertificateManagerCertificateResources() map[string]*AWSCertificateManagerCertificate {
	results := map[string]*AWSCertificateManagerCertificate{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSCertificateManagerCertificate:
			results[name] = resource
		}
	}
	return results
}

// GetAWSCertificateManagerCertificateWithName retrieves all AWSCertificateManagerCertificate items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCertificateManagerCertificateWithName(name string) (*AWSCertificateManagerCertificate, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSCertificateManagerCertificate:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSCertificateManagerCertificate not found", name)
}

// GetAllAWSCloud9EnvironmentEC2Resources retrieves all AWSCloud9EnvironmentEC2 items from an AWS CloudFormation template
func (t *Template) GetAllAWSCloud9EnvironmentEC2Resources() map[string]*AWSCloud9EnvironmentEC2 {
	results := map[string]*AWSCloud9EnvironmentEC2{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSCloud9EnvironmentEC2:
			results[name] = resource
		}
	}
	return results
}

// GetAWSCloud9EnvironmentEC2WithName retrieves all AWSCloud9EnvironmentEC2 items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCloud9EnvironmentEC2WithName(name string) (*AWSCloud9EnvironmentEC2, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSCloud9EnvironmentEC2:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSCloud9EnvironmentEC2 not found", name)
}

// GetAllAWSCloudFormationCustomResourceResources retrieves all AWSCloudFormationCustomResource items from an AWS CloudFormation template
func (t *Template) GetAllAWSCloudFormationCustomResourceResources() map[string]*AWSCloudFormationCustomResource {
	results := map[string]*AWSCloudFormationCustomResource{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSCloudFormationCustomResource:
			results[name] = resource
		}
	}
	return results
}

// GetAWSCloudFormationCustomResourceWithName retrieves all AWSCloudFormationCustomResource items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCloudFormationCustomResourceWithName(name string) (*AWSCloudFormationCustomResource, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSCloudFormationCustomResource:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSCloudFormationCustomResource not found", name)
}

// GetAllAWSCloudFormationMacroResources retrieves all AWSCloudFormationMacro items from an AWS CloudFormation template
func (t *Template) GetAllAWSCloudFormationMacroResources() map[string]*AWSCloudFormationMacro {
	results := map[string]*AWSCloudFormationMacro{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSCloudFormationMacro:
			results[name] = resource
		}
	}
	return results
}

// GetAWSCloudFormationMacroWithName retrieves all AWSCloudFormationMacro items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCloudFormationMacroWithName(name string) (*AWSCloudFormationMacro, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSCloudFormationMacro:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSCloudFormationMacro not found", name)
}

// GetAllAWSCloudFormationStackResources retrieves all AWSCloudFormationStack items from an AWS CloudFormation template
func (t *Template) GetAllAWSCloudFormationStackResources() map[string]*AWSCloudFormationStack {
	results := map[string]*AWSCloudFormationStack{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSCloudFormationStack:
			results[name] = resource
		}
	}
	return results
}

// GetAWSCloudFormationStackWithName retrieves all AWSCloudFormationStack items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCloudFormationStackWithName(name string) (*AWSCloudFormationStack, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSCloudFormationStack:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSCloudFormationStack not found", name)
}

// GetAllAWSCloudFormationWaitConditionResources retrieves all AWSCloudFormationWaitCondition items from an AWS CloudFormation template
func (t *Template) GetAllAWSCloudFormationWaitConditionResources() map[string]*AWSCloudFormationWaitCondition {
	results := map[string]*AWSCloudFormationWaitCondition{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSCloudFormationWaitCondition:
			results[name] = resource
		}
	}
	return results
}

// GetAWSCloudFormationWaitConditionWithName retrieves all AWSCloudFormationWaitCondition items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCloudFormationWaitConditionWithName(name string) (*AWSCloudFormationWaitCondition, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSCloudFormationWaitCondition:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSCloudFormationWaitCondition not found", name)
}

// GetAllAWSCloudFormationWaitConditionHandleResources retrieves all AWSCloudFormationWaitConditionHandle items from an AWS CloudFormation template
func (t *Template) GetAllAWSCloudFormationWaitConditionHandleResources() map[string]*AWSCloudFormationWaitConditionHandle {
	results := map[string]*AWSCloudFormationWaitConditionHandle{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSCloudFormationWaitConditionHandle:
			results[name] = resource
		}
	}
	return results
}

// GetAWSCloudFormationWaitConditionHandleWithName retrieves all AWSCloudFormationWaitConditionHandle items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCloudFormationWaitConditionHandleWithName(name string) (*AWSCloudFormationWaitConditionHandle, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSCloudFormationWaitConditionHandle:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSCloudFormationWaitConditionHandle not found", name)
}

// GetAllAWSCloudFrontCloudFrontOriginAccessIdentityResources retrieves all AWSCloudFrontCloudFrontOriginAccessIdentity items from an AWS CloudFormation template
func (t *Template) GetAllAWSCloudFrontCloudFrontOriginAccessIdentityResources() map[string]*AWSCloudFrontCloudFrontOriginAccessIdentity {
	results := map[string]*AWSCloudFrontCloudFrontOriginAccessIdentity{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSCloudFrontCloudFrontOriginAccessIdentity:
			results[name] = resource
		}
	}
	return results
}

// GetAWSCloudFrontCloudFrontOriginAccessIdentityWithName retrieves all AWSCloudFrontCloudFrontOriginAccessIdentity items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCloudFrontCloudFrontOriginAccessIdentityWithName(name string) (*AWSCloudFrontCloudFrontOriginAccessIdentity, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSCloudFrontCloudFrontOriginAccessIdentity:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSCloudFrontCloudFrontOriginAccessIdentity not found", name)
}

// GetAllAWSCloudFrontDistributionResources retrieves all AWSCloudFrontDistribution items from an AWS CloudFormation template
func (t *Template) GetAllAWSCloudFrontDistributionResources() map[string]*AWSCloudFrontDistribution {
	results := map[string]*AWSCloudFrontDistribution{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSCloudFrontDistribution:
			results[name] = resource
		}
	}
	return results
}

// GetAWSCloudFrontDistributionWithName retrieves all AWSCloudFrontDistribution items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCloudFrontDistributionWithName(name string) (*AWSCloudFrontDistribution, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSCloudFrontDistribution:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSCloudFrontDistribution not found", name)
}

// GetAllAWSCloudFrontStreamingDistributionResources retrieves all AWSCloudFrontStreamingDistribution items from an AWS CloudFormation template
func (t *Template) GetAllAWSCloudFrontStreamingDistributionResources() map[string]*AWSCloudFrontStreamingDistribution {
	results := map[string]*AWSCloudFrontStreamingDistribution{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSCloudFrontStreamingDistribution:
			results[name] = resource
		}
	}
	return results
}

// GetAWSCloudFrontStreamingDistributionWithName retrieves all AWSCloudFrontStreamingDistribution items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCloudFrontStreamingDistributionWithName(name string) (*AWSCloudFrontStreamingDistribution, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSCloudFrontStreamingDistribution:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSCloudFrontStreamingDistribution not found", name)
}

// GetAllAWSCloudTrailTrailResources retrieves all AWSCloudTrailTrail items from an AWS CloudFormation template
func (t *Template) GetAllAWSCloudTrailTrailResources() map[string]*AWSCloudTrailTrail {
	results := map[string]*AWSCloudTrailTrail{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSCloudTrailTrail:
			results[name] = resource
		}
	}
	return results
}

// GetAWSCloudTrailTrailWithName retrieves all AWSCloudTrailTrail items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCloudTrailTrailWithName(name string) (*AWSCloudTrailTrail, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSCloudTrailTrail:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSCloudTrailTrail not found", name)
}

// GetAllAWSCloudWatchAlarmResources retrieves all AWSCloudWatchAlarm items from an AWS CloudFormation template
func (t *Template) GetAllAWSCloudWatchAlarmResources() map[string]*AWSCloudWatchAlarm {
	results := map[string]*AWSCloudWatchAlarm{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSCloudWatchAlarm:
			results[name] = resource
		}
	}
	return results
}

// GetAWSCloudWatchAlarmWithName retrieves all AWSCloudWatchAlarm items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCloudWatchAlarmWithName(name string) (*AWSCloudWatchAlarm, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSCloudWatchAlarm:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSCloudWatchAlarm not found", name)
}

// GetAllAWSCloudWatchDashboardResources retrieves all AWSCloudWatchDashboard items from an AWS CloudFormation template
func (t *Template) GetAllAWSCloudWatchDashboardResources() map[string]*AWSCloudWatchDashboard {
	results := map[string]*AWSCloudWatchDashboard{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSCloudWatchDashboard:
			results[name] = resource
		}
	}
	return results
}

// GetAWSCloudWatchDashboardWithName retrieves all AWSCloudWatchDashboard items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCloudWatchDashboardWithName(name string) (*AWSCloudWatchDashboard, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSCloudWatchDashboard:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSCloudWatchDashboard not found", name)
}

// GetAllAWSCodeBuildProjectResources retrieves all AWSCodeBuildProject items from an AWS CloudFormation template
func (t *Template) GetAllAWSCodeBuildProjectResources() map[string]*AWSCodeBuildProject {
	results := map[string]*AWSCodeBuildProject{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSCodeBuildProject:
			results[name] = resource
		}
	}
	return results
}

// GetAWSCodeBuildProjectWithName retrieves all AWSCodeBuildProject items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCodeBuildProjectWithName(name string) (*AWSCodeBuildProject, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSCodeBuildProject:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSCodeBuildProject not found", name)
}

// GetAllAWSCodeCommitRepositoryResources retrieves all AWSCodeCommitRepository items from an AWS CloudFormation template
func (t *Template) GetAllAWSCodeCommitRepositoryResources() map[string]*AWSCodeCommitRepository {
	results := map[string]*AWSCodeCommitRepository{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSCodeCommitRepository:
			results[name] = resource
		}
	}
	return results
}

// GetAWSCodeCommitRepositoryWithName retrieves all AWSCodeCommitRepository items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCodeCommitRepositoryWithName(name string) (*AWSCodeCommitRepository, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSCodeCommitRepository:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSCodeCommitRepository not found", name)
}

// GetAllAWSCodeDeployApplicationResources retrieves all AWSCodeDeployApplication items from an AWS CloudFormation template
func (t *Template) GetAllAWSCodeDeployApplicationResources() map[string]*AWSCodeDeployApplication {
	results := map[string]*AWSCodeDeployApplication{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSCodeDeployApplication:
			results[name] = resource
		}
	}
	return results
}

// GetAWSCodeDeployApplicationWithName retrieves all AWSCodeDeployApplication items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCodeDeployApplicationWithName(name string) (*AWSCodeDeployApplication, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSCodeDeployApplication:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSCodeDeployApplication not found", name)
}

// GetAllAWSCodeDeployDeploymentConfigResources retrieves all AWSCodeDeployDeploymentConfig items from an AWS CloudFormation template
func (t *Template) GetAllAWSCodeDeployDeploymentConfigResources() map[string]*AWSCodeDeployDeploymentConfig {
	results := map[string]*AWSCodeDeployDeploymentConfig{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSCodeDeployDeploymentConfig:
			results[name] = resource
		}
	}
	return results
}

// GetAWSCodeDeployDeploymentConfigWithName retrieves all AWSCodeDeployDeploymentConfig items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCodeDeployDeploymentConfigWithName(name string) (*AWSCodeDeployDeploymentConfig, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSCodeDeployDeploymentConfig:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSCodeDeployDeploymentConfig not found", name)
}

// GetAllAWSCodeDeployDeploymentGroupResources retrieves all AWSCodeDeployDeploymentGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSCodeDeployDeploymentGroupResources() map[string]*AWSCodeDeployDeploymentGroup {
	results := map[string]*AWSCodeDeployDeploymentGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSCodeDeployDeploymentGroup:
			results[name] = resource
		}
	}
	return results
}

// GetAWSCodeDeployDeploymentGroupWithName retrieves all AWSCodeDeployDeploymentGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCodeDeployDeploymentGroupWithName(name string) (*AWSCodeDeployDeploymentGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSCodeDeployDeploymentGroup:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSCodeDeployDeploymentGroup not found", name)
}

// GetAllAWSCodePipelineCustomActionTypeResources retrieves all AWSCodePipelineCustomActionType items from an AWS CloudFormation template
func (t *Template) GetAllAWSCodePipelineCustomActionTypeResources() map[string]*AWSCodePipelineCustomActionType {
	results := map[string]*AWSCodePipelineCustomActionType{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSCodePipelineCustomActionType:
			results[name] = resource
		}
	}
	return results
}

// GetAWSCodePipelineCustomActionTypeWithName retrieves all AWSCodePipelineCustomActionType items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCodePipelineCustomActionTypeWithName(name string) (*AWSCodePipelineCustomActionType, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSCodePipelineCustomActionType:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSCodePipelineCustomActionType not found", name)
}

// GetAllAWSCodePipelinePipelineResources retrieves all AWSCodePipelinePipeline items from an AWS CloudFormation template
func (t *Template) GetAllAWSCodePipelinePipelineResources() map[string]*AWSCodePipelinePipeline {
	results := map[string]*AWSCodePipelinePipeline{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSCodePipelinePipeline:
			results[name] = resource
		}
	}
	return results
}

// GetAWSCodePipelinePipelineWithName retrieves all AWSCodePipelinePipeline items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCodePipelinePipelineWithName(name string) (*AWSCodePipelinePipeline, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSCodePipelinePipeline:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSCodePipelinePipeline not found", name)
}

// GetAllAWSCodePipelineWebhookResources retrieves all AWSCodePipelineWebhook items from an AWS CloudFormation template
func (t *Template) GetAllAWSCodePipelineWebhookResources() map[string]*AWSCodePipelineWebhook {
	results := map[string]*AWSCodePipelineWebhook{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSCodePipelineWebhook:
			results[name] = resource
		}
	}
	return results
}

// GetAWSCodePipelineWebhookWithName retrieves all AWSCodePipelineWebhook items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCodePipelineWebhookWithName(name string) (*AWSCodePipelineWebhook, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSCodePipelineWebhook:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSCodePipelineWebhook not found", name)
}

// GetAllAWSCognitoIdentityPoolResources retrieves all AWSCognitoIdentityPool items from an AWS CloudFormation template
func (t *Template) GetAllAWSCognitoIdentityPoolResources() map[string]*AWSCognitoIdentityPool {
	results := map[string]*AWSCognitoIdentityPool{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSCognitoIdentityPool:
			results[name] = resource
		}
	}
	return results
}

// GetAWSCognitoIdentityPoolWithName retrieves all AWSCognitoIdentityPool items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCognitoIdentityPoolWithName(name string) (*AWSCognitoIdentityPool, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSCognitoIdentityPool:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSCognitoIdentityPool not found", name)
}

// GetAllAWSCognitoIdentityPoolRoleAttachmentResources retrieves all AWSCognitoIdentityPoolRoleAttachment items from an AWS CloudFormation template
func (t *Template) GetAllAWSCognitoIdentityPoolRoleAttachmentResources() map[string]*AWSCognitoIdentityPoolRoleAttachment {
	results := map[string]*AWSCognitoIdentityPoolRoleAttachment{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSCognitoIdentityPoolRoleAttachment:
			results[name] = resource
		}
	}
	return results
}

// GetAWSCognitoIdentityPoolRoleAttachmentWithName retrieves all AWSCognitoIdentityPoolRoleAttachment items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCognitoIdentityPoolRoleAttachmentWithName(name string) (*AWSCognitoIdentityPoolRoleAttachment, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSCognitoIdentityPoolRoleAttachment:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSCognitoIdentityPoolRoleAttachment not found", name)
}

// GetAllAWSCognitoUserPoolResources retrieves all AWSCognitoUserPool items from an AWS CloudFormation template
func (t *Template) GetAllAWSCognitoUserPoolResources() map[string]*AWSCognitoUserPool {
	results := map[string]*AWSCognitoUserPool{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSCognitoUserPool:
			results[name] = resource
		}
	}
	return results
}

// GetAWSCognitoUserPoolWithName retrieves all AWSCognitoUserPool items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCognitoUserPoolWithName(name string) (*AWSCognitoUserPool, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSCognitoUserPool:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSCognitoUserPool not found", name)
}

// GetAllAWSCognitoUserPoolClientResources retrieves all AWSCognitoUserPoolClient items from an AWS CloudFormation template
func (t *Template) GetAllAWSCognitoUserPoolClientResources() map[string]*AWSCognitoUserPoolClient {
	results := map[string]*AWSCognitoUserPoolClient{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSCognitoUserPoolClient:
			results[name] = resource
		}
	}
	return results
}

// GetAWSCognitoUserPoolClientWithName retrieves all AWSCognitoUserPoolClient items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCognitoUserPoolClientWithName(name string) (*AWSCognitoUserPoolClient, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSCognitoUserPoolClient:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSCognitoUserPoolClient not found", name)
}

// GetAllAWSCognitoUserPoolGroupResources retrieves all AWSCognitoUserPoolGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSCognitoUserPoolGroupResources() map[string]*AWSCognitoUserPoolGroup {
	results := map[string]*AWSCognitoUserPoolGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSCognitoUserPoolGroup:
			results[name] = resource
		}
	}
	return results
}

// GetAWSCognitoUserPoolGroupWithName retrieves all AWSCognitoUserPoolGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCognitoUserPoolGroupWithName(name string) (*AWSCognitoUserPoolGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSCognitoUserPoolGroup:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSCognitoUserPoolGroup not found", name)
}

// GetAllAWSCognitoUserPoolUserResources retrieves all AWSCognitoUserPoolUser items from an AWS CloudFormation template
func (t *Template) GetAllAWSCognitoUserPoolUserResources() map[string]*AWSCognitoUserPoolUser {
	results := map[string]*AWSCognitoUserPoolUser{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSCognitoUserPoolUser:
			results[name] = resource
		}
	}
	return results
}

// GetAWSCognitoUserPoolUserWithName retrieves all AWSCognitoUserPoolUser items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCognitoUserPoolUserWithName(name string) (*AWSCognitoUserPoolUser, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSCognitoUserPoolUser:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSCognitoUserPoolUser not found", name)
}

// GetAllAWSCognitoUserPoolUserToGroupAttachmentResources retrieves all AWSCognitoUserPoolUserToGroupAttachment items from an AWS CloudFormation template
func (t *Template) GetAllAWSCognitoUserPoolUserToGroupAttachmentResources() map[string]*AWSCognitoUserPoolUserToGroupAttachment {
	results := map[string]*AWSCognitoUserPoolUserToGroupAttachment{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSCognitoUserPoolUserToGroupAttachment:
			results[name] = resource
		}
	}
	return results
}

// GetAWSCognitoUserPoolUserToGroupAttachmentWithName retrieves all AWSCognitoUserPoolUserToGroupAttachment items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCognitoUserPoolUserToGroupAttachmentWithName(name string) (*AWSCognitoUserPoolUserToGroupAttachment, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSCognitoUserPoolUserToGroupAttachment:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSCognitoUserPoolUserToGroupAttachment not found", name)
}

// GetAllAWSConfigAggregationAuthorizationResources retrieves all AWSConfigAggregationAuthorization items from an AWS CloudFormation template
func (t *Template) GetAllAWSConfigAggregationAuthorizationResources() map[string]*AWSConfigAggregationAuthorization {
	results := map[string]*AWSConfigAggregationAuthorization{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSConfigAggregationAuthorization:
			results[name] = resource
		}
	}
	return results
}

// GetAWSConfigAggregationAuthorizationWithName retrieves all AWSConfigAggregationAuthorization items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSConfigAggregationAuthorizationWithName(name string) (*AWSConfigAggregationAuthorization, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSConfigAggregationAuthorization:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSConfigAggregationAuthorization not found", name)
}

// GetAllAWSConfigConfigRuleResources retrieves all AWSConfigConfigRule items from an AWS CloudFormation template
func (t *Template) GetAllAWSConfigConfigRuleResources() map[string]*AWSConfigConfigRule {
	results := map[string]*AWSConfigConfigRule{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSConfigConfigRule:
			results[name] = resource
		}
	}
	return results
}

// GetAWSConfigConfigRuleWithName retrieves all AWSConfigConfigRule items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSConfigConfigRuleWithName(name string) (*AWSConfigConfigRule, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSConfigConfigRule:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSConfigConfigRule not found", name)
}

// GetAllAWSConfigConfigurationAggregatorResources retrieves all AWSConfigConfigurationAggregator items from an AWS CloudFormation template
func (t *Template) GetAllAWSConfigConfigurationAggregatorResources() map[string]*AWSConfigConfigurationAggregator {
	results := map[string]*AWSConfigConfigurationAggregator{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSConfigConfigurationAggregator:
			results[name] = resource
		}
	}
	return results
}

// GetAWSConfigConfigurationAggregatorWithName retrieves all AWSConfigConfigurationAggregator items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSConfigConfigurationAggregatorWithName(name string) (*AWSConfigConfigurationAggregator, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSConfigConfigurationAggregator:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSConfigConfigurationAggregator not found", name)
}

// GetAllAWSConfigConfigurationRecorderResources retrieves all AWSConfigConfigurationRecorder items from an AWS CloudFormation template
func (t *Template) GetAllAWSConfigConfigurationRecorderResources() map[string]*AWSConfigConfigurationRecorder {
	results := map[string]*AWSConfigConfigurationRecorder{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSConfigConfigurationRecorder:
			results[name] = resource
		}
	}
	return results
}

// GetAWSConfigConfigurationRecorderWithName retrieves all AWSConfigConfigurationRecorder items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSConfigConfigurationRecorderWithName(name string) (*AWSConfigConfigurationRecorder, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSConfigConfigurationRecorder:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSConfigConfigurationRecorder not found", name)
}

// GetAllAWSConfigDeliveryChannelResources retrieves all AWSConfigDeliveryChannel items from an AWS CloudFormation template
func (t *Template) GetAllAWSConfigDeliveryChannelResources() map[string]*AWSConfigDeliveryChannel {
	results := map[string]*AWSConfigDeliveryChannel{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSConfigDeliveryChannel:
			results[name] = resource
		}
	}
	return results
}

// GetAWSConfigDeliveryChannelWithName retrieves all AWSConfigDeliveryChannel items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSConfigDeliveryChannelWithName(name string) (*AWSConfigDeliveryChannel, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSConfigDeliveryChannel:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSConfigDeliveryChannel not found", name)
}

// GetAllAWSDAXClusterResources retrieves all AWSDAXCluster items from an AWS CloudFormation template
func (t *Template) GetAllAWSDAXClusterResources() map[string]*AWSDAXCluster {
	results := map[string]*AWSDAXCluster{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSDAXCluster:
			results[name] = resource
		}
	}
	return results
}

// GetAWSDAXClusterWithName retrieves all AWSDAXCluster items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSDAXClusterWithName(name string) (*AWSDAXCluster, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSDAXCluster:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSDAXCluster not found", name)
}

// GetAllAWSDAXParameterGroupResources retrieves all AWSDAXParameterGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSDAXParameterGroupResources() map[string]*AWSDAXParameterGroup {
	results := map[string]*AWSDAXParameterGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSDAXParameterGroup:
			results[name] = resource
		}
	}
	return results
}

// GetAWSDAXParameterGroupWithName retrieves all AWSDAXParameterGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSDAXParameterGroupWithName(name string) (*AWSDAXParameterGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSDAXParameterGroup:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSDAXParameterGroup not found", name)
}

// GetAllAWSDAXSubnetGroupResources retrieves all AWSDAXSubnetGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSDAXSubnetGroupResources() map[string]*AWSDAXSubnetGroup {
	results := map[string]*AWSDAXSubnetGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSDAXSubnetGroup:
			results[name] = resource
		}
	}
	return results
}

// GetAWSDAXSubnetGroupWithName retrieves all AWSDAXSubnetGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSDAXSubnetGroupWithName(name string) (*AWSDAXSubnetGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSDAXSubnetGroup:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSDAXSubnetGroup not found", name)
}

// GetAllAWSDLMLifecyclePolicyResources retrieves all AWSDLMLifecyclePolicy items from an AWS CloudFormation template
func (t *Template) GetAllAWSDLMLifecyclePolicyResources() map[string]*AWSDLMLifecyclePolicy {
	results := map[string]*AWSDLMLifecyclePolicy{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSDLMLifecyclePolicy:
			results[name] = resource
		}
	}
	return results
}

// GetAWSDLMLifecyclePolicyWithName retrieves all AWSDLMLifecyclePolicy items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSDLMLifecyclePolicyWithName(name string) (*AWSDLMLifecyclePolicy, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSDLMLifecyclePolicy:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSDLMLifecyclePolicy not found", name)
}

// GetAllAWSDMSCertificateResources retrieves all AWSDMSCertificate items from an AWS CloudFormation template
func (t *Template) GetAllAWSDMSCertificateResources() map[string]*AWSDMSCertificate {
	results := map[string]*AWSDMSCertificate{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSDMSCertificate:
			results[name] = resource
		}
	}
	return results
}

// GetAWSDMSCertificateWithName retrieves all AWSDMSCertificate items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSDMSCertificateWithName(name string) (*AWSDMSCertificate, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSDMSCertificate:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSDMSCertificate not found", name)
}

// GetAllAWSDMSEndpointResources retrieves all AWSDMSEndpoint items from an AWS CloudFormation template
func (t *Template) GetAllAWSDMSEndpointResources() map[string]*AWSDMSEndpoint {
	results := map[string]*AWSDMSEndpoint{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSDMSEndpoint:
			results[name] = resource
		}
	}
	return results
}

// GetAWSDMSEndpointWithName retrieves all AWSDMSEndpoint items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSDMSEndpointWithName(name string) (*AWSDMSEndpoint, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSDMSEndpoint:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSDMSEndpoint not found", name)
}

// GetAllAWSDMSEventSubscriptionResources retrieves all AWSDMSEventSubscription items from an AWS CloudFormation template
func (t *Template) GetAllAWSDMSEventSubscriptionResources() map[string]*AWSDMSEventSubscription {
	results := map[string]*AWSDMSEventSubscription{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSDMSEventSubscription:
			results[name] = resource
		}
	}
	return results
}

// GetAWSDMSEventSubscriptionWithName retrieves all AWSDMSEventSubscription items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSDMSEventSubscriptionWithName(name string) (*AWSDMSEventSubscription, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSDMSEventSubscription:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSDMSEventSubscription not found", name)
}

// GetAllAWSDMSReplicationInstanceResources retrieves all AWSDMSReplicationInstance items from an AWS CloudFormation template
func (t *Template) GetAllAWSDMSReplicationInstanceResources() map[string]*AWSDMSReplicationInstance {
	results := map[string]*AWSDMSReplicationInstance{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSDMSReplicationInstance:
			results[name] = resource
		}
	}
	return results
}

// GetAWSDMSReplicationInstanceWithName retrieves all AWSDMSReplicationInstance items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSDMSReplicationInstanceWithName(name string) (*AWSDMSReplicationInstance, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSDMSReplicationInstance:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSDMSReplicationInstance not found", name)
}

// GetAllAWSDMSReplicationSubnetGroupResources retrieves all AWSDMSReplicationSubnetGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSDMSReplicationSubnetGroupResources() map[string]*AWSDMSReplicationSubnetGroup {
	results := map[string]*AWSDMSReplicationSubnetGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSDMSReplicationSubnetGroup:
			results[name] = resource
		}
	}
	return results
}

// GetAWSDMSReplicationSubnetGroupWithName retrieves all AWSDMSReplicationSubnetGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSDMSReplicationSubnetGroupWithName(name string) (*AWSDMSReplicationSubnetGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSDMSReplicationSubnetGroup:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSDMSReplicationSubnetGroup not found", name)
}

// GetAllAWSDMSReplicationTaskResources retrieves all AWSDMSReplicationTask items from an AWS CloudFormation template
func (t *Template) GetAllAWSDMSReplicationTaskResources() map[string]*AWSDMSReplicationTask {
	results := map[string]*AWSDMSReplicationTask{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSDMSReplicationTask:
			results[name] = resource
		}
	}
	return results
}

// GetAWSDMSReplicationTaskWithName retrieves all AWSDMSReplicationTask items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSDMSReplicationTaskWithName(name string) (*AWSDMSReplicationTask, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSDMSReplicationTask:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSDMSReplicationTask not found", name)
}

// GetAllAWSDataPipelinePipelineResources retrieves all AWSDataPipelinePipeline items from an AWS CloudFormation template
func (t *Template) GetAllAWSDataPipelinePipelineResources() map[string]*AWSDataPipelinePipeline {
	results := map[string]*AWSDataPipelinePipeline{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSDataPipelinePipeline:
			results[name] = resource
		}
	}
	return results
}

// GetAWSDataPipelinePipelineWithName retrieves all AWSDataPipelinePipeline items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSDataPipelinePipelineWithName(name string) (*AWSDataPipelinePipeline, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSDataPipelinePipeline:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSDataPipelinePipeline not found", name)
}

// GetAllAWSDirectoryServiceMicrosoftADResources retrieves all AWSDirectoryServiceMicrosoftAD items from an AWS CloudFormation template
func (t *Template) GetAllAWSDirectoryServiceMicrosoftADResources() map[string]*AWSDirectoryServiceMicrosoftAD {
	results := map[string]*AWSDirectoryServiceMicrosoftAD{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSDirectoryServiceMicrosoftAD:
			results[name] = resource
		}
	}
	return results
}

// GetAWSDirectoryServiceMicrosoftADWithName retrieves all AWSDirectoryServiceMicrosoftAD items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSDirectoryServiceMicrosoftADWithName(name string) (*AWSDirectoryServiceMicrosoftAD, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSDirectoryServiceMicrosoftAD:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSDirectoryServiceMicrosoftAD not found", name)
}

// GetAllAWSDirectoryServiceSimpleADResources retrieves all AWSDirectoryServiceSimpleAD items from an AWS CloudFormation template
func (t *Template) GetAllAWSDirectoryServiceSimpleADResources() map[string]*AWSDirectoryServiceSimpleAD {
	results := map[string]*AWSDirectoryServiceSimpleAD{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSDirectoryServiceSimpleAD:
			results[name] = resource
		}
	}
	return results
}

// GetAWSDirectoryServiceSimpleADWithName retrieves all AWSDirectoryServiceSimpleAD items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSDirectoryServiceSimpleADWithName(name string) (*AWSDirectoryServiceSimpleAD, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSDirectoryServiceSimpleAD:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSDirectoryServiceSimpleAD not found", name)
}

// GetAllAWSDocDBDBClusterResources retrieves all AWSDocDBDBCluster items from an AWS CloudFormation template
func (t *Template) GetAllAWSDocDBDBClusterResources() map[string]*AWSDocDBDBCluster {
	results := map[string]*AWSDocDBDBCluster{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSDocDBDBCluster:
			results[name] = resource
		}
	}
	return results
}

// GetAWSDocDBDBClusterWithName retrieves all AWSDocDBDBCluster items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSDocDBDBClusterWithName(name string) (*AWSDocDBDBCluster, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSDocDBDBCluster:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSDocDBDBCluster not found", name)
}

// GetAllAWSDocDBDBClusterParameterGroupResources retrieves all AWSDocDBDBClusterParameterGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSDocDBDBClusterParameterGroupResources() map[string]*AWSDocDBDBClusterParameterGroup {
	results := map[string]*AWSDocDBDBClusterParameterGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSDocDBDBClusterParameterGroup:
			results[name] = resource
		}
	}
	return results
}

// GetAWSDocDBDBClusterParameterGroupWithName retrieves all AWSDocDBDBClusterParameterGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSDocDBDBClusterParameterGroupWithName(name string) (*AWSDocDBDBClusterParameterGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSDocDBDBClusterParameterGroup:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSDocDBDBClusterParameterGroup not found", name)
}

// GetAllAWSDocDBDBInstanceResources retrieves all AWSDocDBDBInstance items from an AWS CloudFormation template
func (t *Template) GetAllAWSDocDBDBInstanceResources() map[string]*AWSDocDBDBInstance {
	results := map[string]*AWSDocDBDBInstance{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSDocDBDBInstance:
			results[name] = resource
		}
	}
	return results
}

// GetAWSDocDBDBInstanceWithName retrieves all AWSDocDBDBInstance items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSDocDBDBInstanceWithName(name string) (*AWSDocDBDBInstance, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSDocDBDBInstance:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSDocDBDBInstance not found", name)
}

// GetAllAWSDocDBDBSubnetGroupResources retrieves all AWSDocDBDBSubnetGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSDocDBDBSubnetGroupResources() map[string]*AWSDocDBDBSubnetGroup {
	results := map[string]*AWSDocDBDBSubnetGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSDocDBDBSubnetGroup:
			results[name] = resource
		}
	}
	return results
}

// GetAWSDocDBDBSubnetGroupWithName retrieves all AWSDocDBDBSubnetGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSDocDBDBSubnetGroupWithName(name string) (*AWSDocDBDBSubnetGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSDocDBDBSubnetGroup:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSDocDBDBSubnetGroup not found", name)
}

// GetAllAWSDynamoDBTableResources retrieves all AWSDynamoDBTable items from an AWS CloudFormation template
func (t *Template) GetAllAWSDynamoDBTableResources() map[string]*AWSDynamoDBTable {
	results := map[string]*AWSDynamoDBTable{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSDynamoDBTable:
			results[name] = resource
		}
	}
	return results
}

// GetAWSDynamoDBTableWithName retrieves all AWSDynamoDBTable items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSDynamoDBTableWithName(name string) (*AWSDynamoDBTable, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSDynamoDBTable:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSDynamoDBTable not found", name)
}

// GetAllAWSEC2CustomerGatewayResources retrieves all AWSEC2CustomerGateway items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2CustomerGatewayResources() map[string]*AWSEC2CustomerGateway {
	results := map[string]*AWSEC2CustomerGateway{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSEC2CustomerGateway:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2CustomerGatewayWithName retrieves all AWSEC2CustomerGateway items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2CustomerGatewayWithName(name string) (*AWSEC2CustomerGateway, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSEC2CustomerGateway:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2CustomerGateway not found", name)
}

// GetAllAWSEC2DHCPOptionsResources retrieves all AWSEC2DHCPOptions items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2DHCPOptionsResources() map[string]*AWSEC2DHCPOptions {
	results := map[string]*AWSEC2DHCPOptions{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSEC2DHCPOptions:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2DHCPOptionsWithName retrieves all AWSEC2DHCPOptions items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2DHCPOptionsWithName(name string) (*AWSEC2DHCPOptions, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSEC2DHCPOptions:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2DHCPOptions not found", name)
}

// GetAllAWSEC2EC2FleetResources retrieves all AWSEC2EC2Fleet items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2EC2FleetResources() map[string]*AWSEC2EC2Fleet {
	results := map[string]*AWSEC2EC2Fleet{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSEC2EC2Fleet:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2EC2FleetWithName retrieves all AWSEC2EC2Fleet items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2EC2FleetWithName(name string) (*AWSEC2EC2Fleet, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSEC2EC2Fleet:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2EC2Fleet not found", name)
}

// GetAllAWSEC2EIPResources retrieves all AWSEC2EIP items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2EIPResources() map[string]*AWSEC2EIP {
	results := map[string]*AWSEC2EIP{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSEC2EIP:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2EIPWithName retrieves all AWSEC2EIP items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2EIPWithName(name string) (*AWSEC2EIP, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSEC2EIP:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2EIP not found", name)
}

// GetAllAWSEC2EIPAssociationResources retrieves all AWSEC2EIPAssociation items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2EIPAssociationResources() map[string]*AWSEC2EIPAssociation {
	results := map[string]*AWSEC2EIPAssociation{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSEC2EIPAssociation:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2EIPAssociationWithName retrieves all AWSEC2EIPAssociation items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2EIPAssociationWithName(name string) (*AWSEC2EIPAssociation, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSEC2EIPAssociation:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2EIPAssociation not found", name)
}

// GetAllAWSEC2EgressOnlyInternetGatewayResources retrieves all AWSEC2EgressOnlyInternetGateway items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2EgressOnlyInternetGatewayResources() map[string]*AWSEC2EgressOnlyInternetGateway {
	results := map[string]*AWSEC2EgressOnlyInternetGateway{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSEC2EgressOnlyInternetGateway:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2EgressOnlyInternetGatewayWithName retrieves all AWSEC2EgressOnlyInternetGateway items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2EgressOnlyInternetGatewayWithName(name string) (*AWSEC2EgressOnlyInternetGateway, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSEC2EgressOnlyInternetGateway:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2EgressOnlyInternetGateway not found", name)
}

// GetAllAWSEC2FlowLogResources retrieves all AWSEC2FlowLog items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2FlowLogResources() map[string]*AWSEC2FlowLog {
	results := map[string]*AWSEC2FlowLog{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSEC2FlowLog:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2FlowLogWithName retrieves all AWSEC2FlowLog items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2FlowLogWithName(name string) (*AWSEC2FlowLog, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSEC2FlowLog:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2FlowLog not found", name)
}

// GetAllAWSEC2HostResources retrieves all AWSEC2Host items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2HostResources() map[string]*AWSEC2Host {
	results := map[string]*AWSEC2Host{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSEC2Host:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2HostWithName retrieves all AWSEC2Host items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2HostWithName(name string) (*AWSEC2Host, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSEC2Host:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2Host not found", name)
}

// GetAllAWSEC2InstanceResources retrieves all AWSEC2Instance items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2InstanceResources() map[string]*AWSEC2Instance {
	results := map[string]*AWSEC2Instance{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSEC2Instance:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2InstanceWithName retrieves all AWSEC2Instance items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2InstanceWithName(name string) (*AWSEC2Instance, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSEC2Instance:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2Instance not found", name)
}

// GetAllAWSEC2InternetGatewayResources retrieves all AWSEC2InternetGateway items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2InternetGatewayResources() map[string]*AWSEC2InternetGateway {
	results := map[string]*AWSEC2InternetGateway{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSEC2InternetGateway:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2InternetGatewayWithName retrieves all AWSEC2InternetGateway items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2InternetGatewayWithName(name string) (*AWSEC2InternetGateway, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSEC2InternetGateway:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2InternetGateway not found", name)
}

// GetAllAWSEC2LaunchTemplateResources retrieves all AWSEC2LaunchTemplate items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2LaunchTemplateResources() map[string]*AWSEC2LaunchTemplate {
	results := map[string]*AWSEC2LaunchTemplate{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSEC2LaunchTemplate:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2LaunchTemplateWithName retrieves all AWSEC2LaunchTemplate items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2LaunchTemplateWithName(name string) (*AWSEC2LaunchTemplate, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSEC2LaunchTemplate:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2LaunchTemplate not found", name)
}

// GetAllAWSEC2NatGatewayResources retrieves all AWSEC2NatGateway items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2NatGatewayResources() map[string]*AWSEC2NatGateway {
	results := map[string]*AWSEC2NatGateway{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSEC2NatGateway:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2NatGatewayWithName retrieves all AWSEC2NatGateway items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2NatGatewayWithName(name string) (*AWSEC2NatGateway, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSEC2NatGateway:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2NatGateway not found", name)
}

// GetAllAWSEC2NetworkAclResources retrieves all AWSEC2NetworkAcl items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2NetworkAclResources() map[string]*AWSEC2NetworkAcl {
	results := map[string]*AWSEC2NetworkAcl{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSEC2NetworkAcl:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2NetworkAclWithName retrieves all AWSEC2NetworkAcl items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2NetworkAclWithName(name string) (*AWSEC2NetworkAcl, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSEC2NetworkAcl:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2NetworkAcl not found", name)
}

// GetAllAWSEC2NetworkAclEntryResources retrieves all AWSEC2NetworkAclEntry items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2NetworkAclEntryResources() map[string]*AWSEC2NetworkAclEntry {
	results := map[string]*AWSEC2NetworkAclEntry{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSEC2NetworkAclEntry:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2NetworkAclEntryWithName retrieves all AWSEC2NetworkAclEntry items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2NetworkAclEntryWithName(name string) (*AWSEC2NetworkAclEntry, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSEC2NetworkAclEntry:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2NetworkAclEntry not found", name)
}

// GetAllAWSEC2NetworkInterfaceResources retrieves all AWSEC2NetworkInterface items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2NetworkInterfaceResources() map[string]*AWSEC2NetworkInterface {
	results := map[string]*AWSEC2NetworkInterface{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSEC2NetworkInterface:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2NetworkInterfaceWithName retrieves all AWSEC2NetworkInterface items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2NetworkInterfaceWithName(name string) (*AWSEC2NetworkInterface, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSEC2NetworkInterface:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2NetworkInterface not found", name)
}

// GetAllAWSEC2NetworkInterfaceAttachmentResources retrieves all AWSEC2NetworkInterfaceAttachment items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2NetworkInterfaceAttachmentResources() map[string]*AWSEC2NetworkInterfaceAttachment {
	results := map[string]*AWSEC2NetworkInterfaceAttachment{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSEC2NetworkInterfaceAttachment:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2NetworkInterfaceAttachmentWithName retrieves all AWSEC2NetworkInterfaceAttachment items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2NetworkInterfaceAttachmentWithName(name string) (*AWSEC2NetworkInterfaceAttachment, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSEC2NetworkInterfaceAttachment:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2NetworkInterfaceAttachment not found", name)
}

// GetAllAWSEC2NetworkInterfacePermissionResources retrieves all AWSEC2NetworkInterfacePermission items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2NetworkInterfacePermissionResources() map[string]*AWSEC2NetworkInterfacePermission {
	results := map[string]*AWSEC2NetworkInterfacePermission{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSEC2NetworkInterfacePermission:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2NetworkInterfacePermissionWithName retrieves all AWSEC2NetworkInterfacePermission items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2NetworkInterfacePermissionWithName(name string) (*AWSEC2NetworkInterfacePermission, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSEC2NetworkInterfacePermission:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2NetworkInterfacePermission not found", name)
}

// GetAllAWSEC2PlacementGroupResources retrieves all AWSEC2PlacementGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2PlacementGroupResources() map[string]*AWSEC2PlacementGroup {
	results := map[string]*AWSEC2PlacementGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSEC2PlacementGroup:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2PlacementGroupWithName retrieves all AWSEC2PlacementGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2PlacementGroupWithName(name string) (*AWSEC2PlacementGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSEC2PlacementGroup:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2PlacementGroup not found", name)
}

// GetAllAWSEC2RouteResources retrieves all AWSEC2Route items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2RouteResources() map[string]*AWSEC2Route {
	results := map[string]*AWSEC2Route{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSEC2Route:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2RouteWithName retrieves all AWSEC2Route items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2RouteWithName(name string) (*AWSEC2Route, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSEC2Route:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2Route not found", name)
}

// GetAllAWSEC2RouteTableResources retrieves all AWSEC2RouteTable items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2RouteTableResources() map[string]*AWSEC2RouteTable {
	results := map[string]*AWSEC2RouteTable{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSEC2RouteTable:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2RouteTableWithName retrieves all AWSEC2RouteTable items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2RouteTableWithName(name string) (*AWSEC2RouteTable, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSEC2RouteTable:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2RouteTable not found", name)
}

// GetAllAWSEC2SecurityGroupResources retrieves all AWSEC2SecurityGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2SecurityGroupResources() map[string]*AWSEC2SecurityGroup {
	results := map[string]*AWSEC2SecurityGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSEC2SecurityGroup:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2SecurityGroupWithName retrieves all AWSEC2SecurityGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2SecurityGroupWithName(name string) (*AWSEC2SecurityGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSEC2SecurityGroup:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2SecurityGroup not found", name)
}

// GetAllAWSEC2SecurityGroupEgressResources retrieves all AWSEC2SecurityGroupEgress items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2SecurityGroupEgressResources() map[string]*AWSEC2SecurityGroupEgress {
	results := map[string]*AWSEC2SecurityGroupEgress{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSEC2SecurityGroupEgress:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2SecurityGroupEgressWithName retrieves all AWSEC2SecurityGroupEgress items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2SecurityGroupEgressWithName(name string) (*AWSEC2SecurityGroupEgress, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSEC2SecurityGroupEgress:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2SecurityGroupEgress not found", name)
}

// GetAllAWSEC2SecurityGroupIngressResources retrieves all AWSEC2SecurityGroupIngress items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2SecurityGroupIngressResources() map[string]*AWSEC2SecurityGroupIngress {
	results := map[string]*AWSEC2SecurityGroupIngress{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSEC2SecurityGroupIngress:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2SecurityGroupIngressWithName retrieves all AWSEC2SecurityGroupIngress items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2SecurityGroupIngressWithName(name string) (*AWSEC2SecurityGroupIngress, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSEC2SecurityGroupIngress:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2SecurityGroupIngress not found", name)
}

// GetAllAWSEC2SpotFleetResources retrieves all AWSEC2SpotFleet items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2SpotFleetResources() map[string]*AWSEC2SpotFleet {
	results := map[string]*AWSEC2SpotFleet{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSEC2SpotFleet:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2SpotFleetWithName retrieves all AWSEC2SpotFleet items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2SpotFleetWithName(name string) (*AWSEC2SpotFleet, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSEC2SpotFleet:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2SpotFleet not found", name)
}

// GetAllAWSEC2SubnetResources retrieves all AWSEC2Subnet items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2SubnetResources() map[string]*AWSEC2Subnet {
	results := map[string]*AWSEC2Subnet{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSEC2Subnet:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2SubnetWithName retrieves all AWSEC2Subnet items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2SubnetWithName(name string) (*AWSEC2Subnet, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSEC2Subnet:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2Subnet not found", name)
}

// GetAllAWSEC2SubnetCidrBlockResources retrieves all AWSEC2SubnetCidrBlock items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2SubnetCidrBlockResources() map[string]*AWSEC2SubnetCidrBlock {
	results := map[string]*AWSEC2SubnetCidrBlock{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSEC2SubnetCidrBlock:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2SubnetCidrBlockWithName retrieves all AWSEC2SubnetCidrBlock items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2SubnetCidrBlockWithName(name string) (*AWSEC2SubnetCidrBlock, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSEC2SubnetCidrBlock:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2SubnetCidrBlock not found", name)
}

// GetAllAWSEC2SubnetNetworkAclAssociationResources retrieves all AWSEC2SubnetNetworkAclAssociation items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2SubnetNetworkAclAssociationResources() map[string]*AWSEC2SubnetNetworkAclAssociation {
	results := map[string]*AWSEC2SubnetNetworkAclAssociation{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSEC2SubnetNetworkAclAssociation:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2SubnetNetworkAclAssociationWithName retrieves all AWSEC2SubnetNetworkAclAssociation items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2SubnetNetworkAclAssociationWithName(name string) (*AWSEC2SubnetNetworkAclAssociation, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSEC2SubnetNetworkAclAssociation:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2SubnetNetworkAclAssociation not found", name)
}

// GetAllAWSEC2SubnetRouteTableAssociationResources retrieves all AWSEC2SubnetRouteTableAssociation items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2SubnetRouteTableAssociationResources() map[string]*AWSEC2SubnetRouteTableAssociation {
	results := map[string]*AWSEC2SubnetRouteTableAssociation{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSEC2SubnetRouteTableAssociation:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2SubnetRouteTableAssociationWithName retrieves all AWSEC2SubnetRouteTableAssociation items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2SubnetRouteTableAssociationWithName(name string) (*AWSEC2SubnetRouteTableAssociation, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSEC2SubnetRouteTableAssociation:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2SubnetRouteTableAssociation not found", name)
}

// GetAllAWSEC2TransitGatewayResources retrieves all AWSEC2TransitGateway items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2TransitGatewayResources() map[string]*AWSEC2TransitGateway {
	results := map[string]*AWSEC2TransitGateway{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSEC2TransitGateway:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2TransitGatewayWithName retrieves all AWSEC2TransitGateway items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2TransitGatewayWithName(name string) (*AWSEC2TransitGateway, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSEC2TransitGateway:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2TransitGateway not found", name)
}

// GetAllAWSEC2TransitGatewayAttachmentResources retrieves all AWSEC2TransitGatewayAttachment items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2TransitGatewayAttachmentResources() map[string]*AWSEC2TransitGatewayAttachment {
	results := map[string]*AWSEC2TransitGatewayAttachment{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSEC2TransitGatewayAttachment:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2TransitGatewayAttachmentWithName retrieves all AWSEC2TransitGatewayAttachment items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2TransitGatewayAttachmentWithName(name string) (*AWSEC2TransitGatewayAttachment, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSEC2TransitGatewayAttachment:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2TransitGatewayAttachment not found", name)
}

// GetAllAWSEC2TransitGatewayRouteResources retrieves all AWSEC2TransitGatewayRoute items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2TransitGatewayRouteResources() map[string]*AWSEC2TransitGatewayRoute {
	results := map[string]*AWSEC2TransitGatewayRoute{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSEC2TransitGatewayRoute:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2TransitGatewayRouteWithName retrieves all AWSEC2TransitGatewayRoute items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2TransitGatewayRouteWithName(name string) (*AWSEC2TransitGatewayRoute, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSEC2TransitGatewayRoute:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2TransitGatewayRoute not found", name)
}

// GetAllAWSEC2TransitGatewayRouteTableResources retrieves all AWSEC2TransitGatewayRouteTable items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2TransitGatewayRouteTableResources() map[string]*AWSEC2TransitGatewayRouteTable {
	results := map[string]*AWSEC2TransitGatewayRouteTable{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSEC2TransitGatewayRouteTable:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2TransitGatewayRouteTableWithName retrieves all AWSEC2TransitGatewayRouteTable items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2TransitGatewayRouteTableWithName(name string) (*AWSEC2TransitGatewayRouteTable, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSEC2TransitGatewayRouteTable:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2TransitGatewayRouteTable not found", name)
}

// GetAllAWSEC2TransitGatewayRouteTableAssociationResources retrieves all AWSEC2TransitGatewayRouteTableAssociation items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2TransitGatewayRouteTableAssociationResources() map[string]*AWSEC2TransitGatewayRouteTableAssociation {
	results := map[string]*AWSEC2TransitGatewayRouteTableAssociation{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSEC2TransitGatewayRouteTableAssociation:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2TransitGatewayRouteTableAssociationWithName retrieves all AWSEC2TransitGatewayRouteTableAssociation items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2TransitGatewayRouteTableAssociationWithName(name string) (*AWSEC2TransitGatewayRouteTableAssociation, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSEC2TransitGatewayRouteTableAssociation:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2TransitGatewayRouteTableAssociation not found", name)
}

// GetAllAWSEC2TransitGatewayRouteTablePropagationResources retrieves all AWSEC2TransitGatewayRouteTablePropagation items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2TransitGatewayRouteTablePropagationResources() map[string]*AWSEC2TransitGatewayRouteTablePropagation {
	results := map[string]*AWSEC2TransitGatewayRouteTablePropagation{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSEC2TransitGatewayRouteTablePropagation:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2TransitGatewayRouteTablePropagationWithName retrieves all AWSEC2TransitGatewayRouteTablePropagation items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2TransitGatewayRouteTablePropagationWithName(name string) (*AWSEC2TransitGatewayRouteTablePropagation, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSEC2TransitGatewayRouteTablePropagation:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2TransitGatewayRouteTablePropagation not found", name)
}

// GetAllAWSEC2TrunkInterfaceAssociationResources retrieves all AWSEC2TrunkInterfaceAssociation items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2TrunkInterfaceAssociationResources() map[string]*AWSEC2TrunkInterfaceAssociation {
	results := map[string]*AWSEC2TrunkInterfaceAssociation{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSEC2TrunkInterfaceAssociation:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2TrunkInterfaceAssociationWithName retrieves all AWSEC2TrunkInterfaceAssociation items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2TrunkInterfaceAssociationWithName(name string) (*AWSEC2TrunkInterfaceAssociation, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSEC2TrunkInterfaceAssociation:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2TrunkInterfaceAssociation not found", name)
}

// GetAllAWSEC2VPCResources retrieves all AWSEC2VPC items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2VPCResources() map[string]*AWSEC2VPC {
	results := map[string]*AWSEC2VPC{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSEC2VPC:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2VPCWithName retrieves all AWSEC2VPC items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2VPCWithName(name string) (*AWSEC2VPC, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSEC2VPC:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2VPC not found", name)
}

// GetAllAWSEC2VPCCidrBlockResources retrieves all AWSEC2VPCCidrBlock items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2VPCCidrBlockResources() map[string]*AWSEC2VPCCidrBlock {
	results := map[string]*AWSEC2VPCCidrBlock{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSEC2VPCCidrBlock:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2VPCCidrBlockWithName retrieves all AWSEC2VPCCidrBlock items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2VPCCidrBlockWithName(name string) (*AWSEC2VPCCidrBlock, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSEC2VPCCidrBlock:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2VPCCidrBlock not found", name)
}

// GetAllAWSEC2VPCDHCPOptionsAssociationResources retrieves all AWSEC2VPCDHCPOptionsAssociation items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2VPCDHCPOptionsAssociationResources() map[string]*AWSEC2VPCDHCPOptionsAssociation {
	results := map[string]*AWSEC2VPCDHCPOptionsAssociation{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSEC2VPCDHCPOptionsAssociation:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2VPCDHCPOptionsAssociationWithName retrieves all AWSEC2VPCDHCPOptionsAssociation items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2VPCDHCPOptionsAssociationWithName(name string) (*AWSEC2VPCDHCPOptionsAssociation, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSEC2VPCDHCPOptionsAssociation:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2VPCDHCPOptionsAssociation not found", name)
}

// GetAllAWSEC2VPCEndpointResources retrieves all AWSEC2VPCEndpoint items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2VPCEndpointResources() map[string]*AWSEC2VPCEndpoint {
	results := map[string]*AWSEC2VPCEndpoint{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSEC2VPCEndpoint:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2VPCEndpointWithName retrieves all AWSEC2VPCEndpoint items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2VPCEndpointWithName(name string) (*AWSEC2VPCEndpoint, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSEC2VPCEndpoint:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2VPCEndpoint not found", name)
}

// GetAllAWSEC2VPCEndpointConnectionNotificationResources retrieves all AWSEC2VPCEndpointConnectionNotification items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2VPCEndpointConnectionNotificationResources() map[string]*AWSEC2VPCEndpointConnectionNotification {
	results := map[string]*AWSEC2VPCEndpointConnectionNotification{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSEC2VPCEndpointConnectionNotification:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2VPCEndpointConnectionNotificationWithName retrieves all AWSEC2VPCEndpointConnectionNotification items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2VPCEndpointConnectionNotificationWithName(name string) (*AWSEC2VPCEndpointConnectionNotification, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSEC2VPCEndpointConnectionNotification:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2VPCEndpointConnectionNotification not found", name)
}

// GetAllAWSEC2VPCEndpointServicePermissionsResources retrieves all AWSEC2VPCEndpointServicePermissions items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2VPCEndpointServicePermissionsResources() map[string]*AWSEC2VPCEndpointServicePermissions {
	results := map[string]*AWSEC2VPCEndpointServicePermissions{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSEC2VPCEndpointServicePermissions:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2VPCEndpointServicePermissionsWithName retrieves all AWSEC2VPCEndpointServicePermissions items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2VPCEndpointServicePermissionsWithName(name string) (*AWSEC2VPCEndpointServicePermissions, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSEC2VPCEndpointServicePermissions:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2VPCEndpointServicePermissions not found", name)
}

// GetAllAWSEC2VPCGatewayAttachmentResources retrieves all AWSEC2VPCGatewayAttachment items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2VPCGatewayAttachmentResources() map[string]*AWSEC2VPCGatewayAttachment {
	results := map[string]*AWSEC2VPCGatewayAttachment{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSEC2VPCGatewayAttachment:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2VPCGatewayAttachmentWithName retrieves all AWSEC2VPCGatewayAttachment items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2VPCGatewayAttachmentWithName(name string) (*AWSEC2VPCGatewayAttachment, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSEC2VPCGatewayAttachment:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2VPCGatewayAttachment not found", name)
}

// GetAllAWSEC2VPCPeeringConnectionResources retrieves all AWSEC2VPCPeeringConnection items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2VPCPeeringConnectionResources() map[string]*AWSEC2VPCPeeringConnection {
	results := map[string]*AWSEC2VPCPeeringConnection{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSEC2VPCPeeringConnection:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2VPCPeeringConnectionWithName retrieves all AWSEC2VPCPeeringConnection items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2VPCPeeringConnectionWithName(name string) (*AWSEC2VPCPeeringConnection, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSEC2VPCPeeringConnection:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2VPCPeeringConnection not found", name)
}

// GetAllAWSEC2VPNConnectionResources retrieves all AWSEC2VPNConnection items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2VPNConnectionResources() map[string]*AWSEC2VPNConnection {
	results := map[string]*AWSEC2VPNConnection{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSEC2VPNConnection:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2VPNConnectionWithName retrieves all AWSEC2VPNConnection items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2VPNConnectionWithName(name string) (*AWSEC2VPNConnection, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSEC2VPNConnection:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2VPNConnection not found", name)
}

// GetAllAWSEC2VPNConnectionRouteResources retrieves all AWSEC2VPNConnectionRoute items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2VPNConnectionRouteResources() map[string]*AWSEC2VPNConnectionRoute {
	results := map[string]*AWSEC2VPNConnectionRoute{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSEC2VPNConnectionRoute:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2VPNConnectionRouteWithName retrieves all AWSEC2VPNConnectionRoute items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2VPNConnectionRouteWithName(name string) (*AWSEC2VPNConnectionRoute, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSEC2VPNConnectionRoute:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2VPNConnectionRoute not found", name)
}

// GetAllAWSEC2VPNGatewayResources retrieves all AWSEC2VPNGateway items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2VPNGatewayResources() map[string]*AWSEC2VPNGateway {
	results := map[string]*AWSEC2VPNGateway{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSEC2VPNGateway:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2VPNGatewayWithName retrieves all AWSEC2VPNGateway items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2VPNGatewayWithName(name string) (*AWSEC2VPNGateway, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSEC2VPNGateway:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2VPNGateway not found", name)
}

// GetAllAWSEC2VPNGatewayRoutePropagationResources retrieves all AWSEC2VPNGatewayRoutePropagation items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2VPNGatewayRoutePropagationResources() map[string]*AWSEC2VPNGatewayRoutePropagation {
	results := map[string]*AWSEC2VPNGatewayRoutePropagation{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSEC2VPNGatewayRoutePropagation:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2VPNGatewayRoutePropagationWithName retrieves all AWSEC2VPNGatewayRoutePropagation items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2VPNGatewayRoutePropagationWithName(name string) (*AWSEC2VPNGatewayRoutePropagation, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSEC2VPNGatewayRoutePropagation:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2VPNGatewayRoutePropagation not found", name)
}

// GetAllAWSEC2VolumeResources retrieves all AWSEC2Volume items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2VolumeResources() map[string]*AWSEC2Volume {
	results := map[string]*AWSEC2Volume{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSEC2Volume:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2VolumeWithName retrieves all AWSEC2Volume items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2VolumeWithName(name string) (*AWSEC2Volume, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSEC2Volume:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2Volume not found", name)
}

// GetAllAWSEC2VolumeAttachmentResources retrieves all AWSEC2VolumeAttachment items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2VolumeAttachmentResources() map[string]*AWSEC2VolumeAttachment {
	results := map[string]*AWSEC2VolumeAttachment{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSEC2VolumeAttachment:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2VolumeAttachmentWithName retrieves all AWSEC2VolumeAttachment items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2VolumeAttachmentWithName(name string) (*AWSEC2VolumeAttachment, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSEC2VolumeAttachment:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2VolumeAttachment not found", name)
}

// GetAllAWSECRRepositoryResources retrieves all AWSECRRepository items from an AWS CloudFormation template
func (t *Template) GetAllAWSECRRepositoryResources() map[string]*AWSECRRepository {
	results := map[string]*AWSECRRepository{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSECRRepository:
			results[name] = resource
		}
	}
	return results
}

// GetAWSECRRepositoryWithName retrieves all AWSECRRepository items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSECRRepositoryWithName(name string) (*AWSECRRepository, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSECRRepository:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSECRRepository not found", name)
}

// GetAllAWSECSClusterResources retrieves all AWSECSCluster items from an AWS CloudFormation template
func (t *Template) GetAllAWSECSClusterResources() map[string]*AWSECSCluster {
	results := map[string]*AWSECSCluster{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSECSCluster:
			results[name] = resource
		}
	}
	return results
}

// GetAWSECSClusterWithName retrieves all AWSECSCluster items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSECSClusterWithName(name string) (*AWSECSCluster, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSECSCluster:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSECSCluster not found", name)
}

// GetAllAWSECSServiceResources retrieves all AWSECSService items from an AWS CloudFormation template
func (t *Template) GetAllAWSECSServiceResources() map[string]*AWSECSService {
	results := map[string]*AWSECSService{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSECSService:
			results[name] = resource
		}
	}
	return results
}

// GetAWSECSServiceWithName retrieves all AWSECSService items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSECSServiceWithName(name string) (*AWSECSService, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSECSService:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSECSService not found", name)
}

// GetAllAWSECSTaskDefinitionResources retrieves all AWSECSTaskDefinition items from an AWS CloudFormation template
func (t *Template) GetAllAWSECSTaskDefinitionResources() map[string]*AWSECSTaskDefinition {
	results := map[string]*AWSECSTaskDefinition{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSECSTaskDefinition:
			results[name] = resource
		}
	}
	return results
}

// GetAWSECSTaskDefinitionWithName retrieves all AWSECSTaskDefinition items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSECSTaskDefinitionWithName(name string) (*AWSECSTaskDefinition, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSECSTaskDefinition:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSECSTaskDefinition not found", name)
}

// GetAllAWSEFSFileSystemResources retrieves all AWSEFSFileSystem items from an AWS CloudFormation template
func (t *Template) GetAllAWSEFSFileSystemResources() map[string]*AWSEFSFileSystem {
	results := map[string]*AWSEFSFileSystem{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSEFSFileSystem:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEFSFileSystemWithName retrieves all AWSEFSFileSystem items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEFSFileSystemWithName(name string) (*AWSEFSFileSystem, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSEFSFileSystem:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEFSFileSystem not found", name)
}

// GetAllAWSEFSMountTargetResources retrieves all AWSEFSMountTarget items from an AWS CloudFormation template
func (t *Template) GetAllAWSEFSMountTargetResources() map[string]*AWSEFSMountTarget {
	results := map[string]*AWSEFSMountTarget{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSEFSMountTarget:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEFSMountTargetWithName retrieves all AWSEFSMountTarget items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEFSMountTargetWithName(name string) (*AWSEFSMountTarget, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSEFSMountTarget:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEFSMountTarget not found", name)
}

// GetAllAWSEKSClusterResources retrieves all AWSEKSCluster items from an AWS CloudFormation template
func (t *Template) GetAllAWSEKSClusterResources() map[string]*AWSEKSCluster {
	results := map[string]*AWSEKSCluster{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSEKSCluster:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEKSClusterWithName retrieves all AWSEKSCluster items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEKSClusterWithName(name string) (*AWSEKSCluster, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSEKSCluster:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEKSCluster not found", name)
}

// GetAllAWSEMRClusterResources retrieves all AWSEMRCluster items from an AWS CloudFormation template
func (t *Template) GetAllAWSEMRClusterResources() map[string]*AWSEMRCluster {
	results := map[string]*AWSEMRCluster{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSEMRCluster:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEMRClusterWithName retrieves all AWSEMRCluster items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEMRClusterWithName(name string) (*AWSEMRCluster, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSEMRCluster:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEMRCluster not found", name)
}

// GetAllAWSEMRInstanceFleetConfigResources retrieves all AWSEMRInstanceFleetConfig items from an AWS CloudFormation template
func (t *Template) GetAllAWSEMRInstanceFleetConfigResources() map[string]*AWSEMRInstanceFleetConfig {
	results := map[string]*AWSEMRInstanceFleetConfig{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSEMRInstanceFleetConfig:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEMRInstanceFleetConfigWithName retrieves all AWSEMRInstanceFleetConfig items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEMRInstanceFleetConfigWithName(name string) (*AWSEMRInstanceFleetConfig, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSEMRInstanceFleetConfig:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEMRInstanceFleetConfig not found", name)
}

// GetAllAWSEMRInstanceGroupConfigResources retrieves all AWSEMRInstanceGroupConfig items from an AWS CloudFormation template
func (t *Template) GetAllAWSEMRInstanceGroupConfigResources() map[string]*AWSEMRInstanceGroupConfig {
	results := map[string]*AWSEMRInstanceGroupConfig{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSEMRInstanceGroupConfig:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEMRInstanceGroupConfigWithName retrieves all AWSEMRInstanceGroupConfig items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEMRInstanceGroupConfigWithName(name string) (*AWSEMRInstanceGroupConfig, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSEMRInstanceGroupConfig:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEMRInstanceGroupConfig not found", name)
}

// GetAllAWSEMRSecurityConfigurationResources retrieves all AWSEMRSecurityConfiguration items from an AWS CloudFormation template
func (t *Template) GetAllAWSEMRSecurityConfigurationResources() map[string]*AWSEMRSecurityConfiguration {
	results := map[string]*AWSEMRSecurityConfiguration{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSEMRSecurityConfiguration:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEMRSecurityConfigurationWithName retrieves all AWSEMRSecurityConfiguration items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEMRSecurityConfigurationWithName(name string) (*AWSEMRSecurityConfiguration, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSEMRSecurityConfiguration:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEMRSecurityConfiguration not found", name)
}

// GetAllAWSEMRStepResources retrieves all AWSEMRStep items from an AWS CloudFormation template
func (t *Template) GetAllAWSEMRStepResources() map[string]*AWSEMRStep {
	results := map[string]*AWSEMRStep{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSEMRStep:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEMRStepWithName retrieves all AWSEMRStep items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEMRStepWithName(name string) (*AWSEMRStep, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSEMRStep:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEMRStep not found", name)
}

// GetAllAWSElastiCacheCacheClusterResources retrieves all AWSElastiCacheCacheCluster items from an AWS CloudFormation template
func (t *Template) GetAllAWSElastiCacheCacheClusterResources() map[string]*AWSElastiCacheCacheCluster {
	results := map[string]*AWSElastiCacheCacheCluster{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSElastiCacheCacheCluster:
			results[name] = resource
		}
	}
	return results
}

// GetAWSElastiCacheCacheClusterWithName retrieves all AWSElastiCacheCacheCluster items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSElastiCacheCacheClusterWithName(name string) (*AWSElastiCacheCacheCluster, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSElastiCacheCacheCluster:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSElastiCacheCacheCluster not found", name)
}

// GetAllAWSElastiCacheParameterGroupResources retrieves all AWSElastiCacheParameterGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSElastiCacheParameterGroupResources() map[string]*AWSElastiCacheParameterGroup {
	results := map[string]*AWSElastiCacheParameterGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSElastiCacheParameterGroup:
			results[name] = resource
		}
	}
	return results
}

// GetAWSElastiCacheParameterGroupWithName retrieves all AWSElastiCacheParameterGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSElastiCacheParameterGroupWithName(name string) (*AWSElastiCacheParameterGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSElastiCacheParameterGroup:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSElastiCacheParameterGroup not found", name)
}

// GetAllAWSElastiCacheReplicationGroupResources retrieves all AWSElastiCacheReplicationGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSElastiCacheReplicationGroupResources() map[string]*AWSElastiCacheReplicationGroup {
	results := map[string]*AWSElastiCacheReplicationGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSElastiCacheReplicationGroup:
			results[name] = resource
		}
	}
	return results
}

// GetAWSElastiCacheReplicationGroupWithName retrieves all AWSElastiCacheReplicationGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSElastiCacheReplicationGroupWithName(name string) (*AWSElastiCacheReplicationGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSElastiCacheReplicationGroup:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSElastiCacheReplicationGroup not found", name)
}

// GetAllAWSElastiCacheSecurityGroupResources retrieves all AWSElastiCacheSecurityGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSElastiCacheSecurityGroupResources() map[string]*AWSElastiCacheSecurityGroup {
	results := map[string]*AWSElastiCacheSecurityGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSElastiCacheSecurityGroup:
			results[name] = resource
		}
	}
	return results
}

// GetAWSElastiCacheSecurityGroupWithName retrieves all AWSElastiCacheSecurityGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSElastiCacheSecurityGroupWithName(name string) (*AWSElastiCacheSecurityGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSElastiCacheSecurityGroup:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSElastiCacheSecurityGroup not found", name)
}

// GetAllAWSElastiCacheSecurityGroupIngressResources retrieves all AWSElastiCacheSecurityGroupIngress items from an AWS CloudFormation template
func (t *Template) GetAllAWSElastiCacheSecurityGroupIngressResources() map[string]*AWSElastiCacheSecurityGroupIngress {
	results := map[string]*AWSElastiCacheSecurityGroupIngress{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSElastiCacheSecurityGroupIngress:
			results[name] = resource
		}
	}
	return results
}

// GetAWSElastiCacheSecurityGroupIngressWithName retrieves all AWSElastiCacheSecurityGroupIngress items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSElastiCacheSecurityGroupIngressWithName(name string) (*AWSElastiCacheSecurityGroupIngress, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSElastiCacheSecurityGroupIngress:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSElastiCacheSecurityGroupIngress not found", name)
}

// GetAllAWSElastiCacheSubnetGroupResources retrieves all AWSElastiCacheSubnetGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSElastiCacheSubnetGroupResources() map[string]*AWSElastiCacheSubnetGroup {
	results := map[string]*AWSElastiCacheSubnetGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSElastiCacheSubnetGroup:
			results[name] = resource
		}
	}
	return results
}

// GetAWSElastiCacheSubnetGroupWithName retrieves all AWSElastiCacheSubnetGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSElastiCacheSubnetGroupWithName(name string) (*AWSElastiCacheSubnetGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSElastiCacheSubnetGroup:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSElastiCacheSubnetGroup not found", name)
}

// GetAllAWSElasticBeanstalkApplicationResources retrieves all AWSElasticBeanstalkApplication items from an AWS CloudFormation template
func (t *Template) GetAllAWSElasticBeanstalkApplicationResources() map[string]*AWSElasticBeanstalkApplication {
	results := map[string]*AWSElasticBeanstalkApplication{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSElasticBeanstalkApplication:
			results[name] = resource
		}
	}
	return results
}

// GetAWSElasticBeanstalkApplicationWithName retrieves all AWSElasticBeanstalkApplication items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSElasticBeanstalkApplicationWithName(name string) (*AWSElasticBeanstalkApplication, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSElasticBeanstalkApplication:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSElasticBeanstalkApplication not found", name)
}

// GetAllAWSElasticBeanstalkApplicationVersionResources retrieves all AWSElasticBeanstalkApplicationVersion items from an AWS CloudFormation template
func (t *Template) GetAllAWSElasticBeanstalkApplicationVersionResources() map[string]*AWSElasticBeanstalkApplicationVersion {
	results := map[string]*AWSElasticBeanstalkApplicationVersion{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSElasticBeanstalkApplicationVersion:
			results[name] = resource
		}
	}
	return results
}

// GetAWSElasticBeanstalkApplicationVersionWithName retrieves all AWSElasticBeanstalkApplicationVersion items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSElasticBeanstalkApplicationVersionWithName(name string) (*AWSElasticBeanstalkApplicationVersion, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSElasticBeanstalkApplicationVersion:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSElasticBeanstalkApplicationVersion not found", name)
}

// GetAllAWSElasticBeanstalkConfigurationTemplateResources retrieves all AWSElasticBeanstalkConfigurationTemplate items from an AWS CloudFormation template
func (t *Template) GetAllAWSElasticBeanstalkConfigurationTemplateResources() map[string]*AWSElasticBeanstalkConfigurationTemplate {
	results := map[string]*AWSElasticBeanstalkConfigurationTemplate{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSElasticBeanstalkConfigurationTemplate:
			results[name] = resource
		}
	}
	return results
}

// GetAWSElasticBeanstalkConfigurationTemplateWithName retrieves all AWSElasticBeanstalkConfigurationTemplate items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSElasticBeanstalkConfigurationTemplateWithName(name string) (*AWSElasticBeanstalkConfigurationTemplate, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSElasticBeanstalkConfigurationTemplate:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSElasticBeanstalkConfigurationTemplate not found", name)
}

// GetAllAWSElasticBeanstalkEnvironmentResources retrieves all AWSElasticBeanstalkEnvironment items from an AWS CloudFormation template
func (t *Template) GetAllAWSElasticBeanstalkEnvironmentResources() map[string]*AWSElasticBeanstalkEnvironment {
	results := map[string]*AWSElasticBeanstalkEnvironment{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSElasticBeanstalkEnvironment:
			results[name] = resource
		}
	}
	return results
}

// GetAWSElasticBeanstalkEnvironmentWithName retrieves all AWSElasticBeanstalkEnvironment items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSElasticBeanstalkEnvironmentWithName(name string) (*AWSElasticBeanstalkEnvironment, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSElasticBeanstalkEnvironment:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSElasticBeanstalkEnvironment not found", name)
}

// GetAllAWSElasticLoadBalancingLoadBalancerResources retrieves all AWSElasticLoadBalancingLoadBalancer items from an AWS CloudFormation template
func (t *Template) GetAllAWSElasticLoadBalancingLoadBalancerResources() map[string]*AWSElasticLoadBalancingLoadBalancer {
	results := map[string]*AWSElasticLoadBalancingLoadBalancer{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSElasticLoadBalancingLoadBalancer:
			results[name] = resource
		}
	}
	return results
}

// GetAWSElasticLoadBalancingLoadBalancerWithName retrieves all AWSElasticLoadBalancingLoadBalancer items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSElasticLoadBalancingLoadBalancerWithName(name string) (*AWSElasticLoadBalancingLoadBalancer, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSElasticLoadBalancingLoadBalancer:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSElasticLoadBalancingLoadBalancer not found", name)
}

// GetAllAWSElasticLoadBalancingV2ListenerResources retrieves all AWSElasticLoadBalancingV2Listener items from an AWS CloudFormation template
func (t *Template) GetAllAWSElasticLoadBalancingV2ListenerResources() map[string]*AWSElasticLoadBalancingV2Listener {
	results := map[string]*AWSElasticLoadBalancingV2Listener{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSElasticLoadBalancingV2Listener:
			results[name] = resource
		}
	}
	return results
}

// GetAWSElasticLoadBalancingV2ListenerWithName retrieves all AWSElasticLoadBalancingV2Listener items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSElasticLoadBalancingV2ListenerWithName(name string) (*AWSElasticLoadBalancingV2Listener, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSElasticLoadBalancingV2Listener:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSElasticLoadBalancingV2Listener not found", name)
}

// GetAllAWSElasticLoadBalancingV2ListenerCertificateResources retrieves all AWSElasticLoadBalancingV2ListenerCertificate items from an AWS CloudFormation template
func (t *Template) GetAllAWSElasticLoadBalancingV2ListenerCertificateResources() map[string]*AWSElasticLoadBalancingV2ListenerCertificate {
	results := map[string]*AWSElasticLoadBalancingV2ListenerCertificate{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSElasticLoadBalancingV2ListenerCertificate:
			results[name] = resource
		}
	}
	return results
}

// GetAWSElasticLoadBalancingV2ListenerCertificateWithName retrieves all AWSElasticLoadBalancingV2ListenerCertificate items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSElasticLoadBalancingV2ListenerCertificateWithName(name string) (*AWSElasticLoadBalancingV2ListenerCertificate, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSElasticLoadBalancingV2ListenerCertificate:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSElasticLoadBalancingV2ListenerCertificate not found", name)
}

// GetAllAWSElasticLoadBalancingV2ListenerRuleResources retrieves all AWSElasticLoadBalancingV2ListenerRule items from an AWS CloudFormation template
func (t *Template) GetAllAWSElasticLoadBalancingV2ListenerRuleResources() map[string]*AWSElasticLoadBalancingV2ListenerRule {
	results := map[string]*AWSElasticLoadBalancingV2ListenerRule{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSElasticLoadBalancingV2ListenerRule:
			results[name] = resource
		}
	}
	return results
}

// GetAWSElasticLoadBalancingV2ListenerRuleWithName retrieves all AWSElasticLoadBalancingV2ListenerRule items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSElasticLoadBalancingV2ListenerRuleWithName(name string) (*AWSElasticLoadBalancingV2ListenerRule, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSElasticLoadBalancingV2ListenerRule:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSElasticLoadBalancingV2ListenerRule not found", name)
}

// GetAllAWSElasticLoadBalancingV2LoadBalancerResources retrieves all AWSElasticLoadBalancingV2LoadBalancer items from an AWS CloudFormation template
func (t *Template) GetAllAWSElasticLoadBalancingV2LoadBalancerResources() map[string]*AWSElasticLoadBalancingV2LoadBalancer {
	results := map[string]*AWSElasticLoadBalancingV2LoadBalancer{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSElasticLoadBalancingV2LoadBalancer:
			results[name] = resource
		}
	}
	return results
}

// GetAWSElasticLoadBalancingV2LoadBalancerWithName retrieves all AWSElasticLoadBalancingV2LoadBalancer items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSElasticLoadBalancingV2LoadBalancerWithName(name string) (*AWSElasticLoadBalancingV2LoadBalancer, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSElasticLoadBalancingV2LoadBalancer:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSElasticLoadBalancingV2LoadBalancer not found", name)
}

// GetAllAWSElasticLoadBalancingV2TargetGroupResources retrieves all AWSElasticLoadBalancingV2TargetGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSElasticLoadBalancingV2TargetGroupResources() map[string]*AWSElasticLoadBalancingV2TargetGroup {
	results := map[string]*AWSElasticLoadBalancingV2TargetGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSElasticLoadBalancingV2TargetGroup:
			results[name] = resource
		}
	}
	return results
}

// GetAWSElasticLoadBalancingV2TargetGroupWithName retrieves all AWSElasticLoadBalancingV2TargetGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSElasticLoadBalancingV2TargetGroupWithName(name string) (*AWSElasticLoadBalancingV2TargetGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSElasticLoadBalancingV2TargetGroup:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSElasticLoadBalancingV2TargetGroup not found", name)
}

// GetAllAWSElasticsearchDomainResources retrieves all AWSElasticsearchDomain items from an AWS CloudFormation template
func (t *Template) GetAllAWSElasticsearchDomainResources() map[string]*AWSElasticsearchDomain {
	results := map[string]*AWSElasticsearchDomain{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSElasticsearchDomain:
			results[name] = resource
		}
	}
	return results
}

// GetAWSElasticsearchDomainWithName retrieves all AWSElasticsearchDomain items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSElasticsearchDomainWithName(name string) (*AWSElasticsearchDomain, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSElasticsearchDomain:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSElasticsearchDomain not found", name)
}

// GetAllAWSEventsEventBusPolicyResources retrieves all AWSEventsEventBusPolicy items from an AWS CloudFormation template
func (t *Template) GetAllAWSEventsEventBusPolicyResources() map[string]*AWSEventsEventBusPolicy {
	results := map[string]*AWSEventsEventBusPolicy{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSEventsEventBusPolicy:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEventsEventBusPolicyWithName retrieves all AWSEventsEventBusPolicy items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEventsEventBusPolicyWithName(name string) (*AWSEventsEventBusPolicy, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSEventsEventBusPolicy:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEventsEventBusPolicy not found", name)
}

// GetAllAWSEventsRuleResources retrieves all AWSEventsRule items from an AWS CloudFormation template
func (t *Template) GetAllAWSEventsRuleResources() map[string]*AWSEventsRule {
	results := map[string]*AWSEventsRule{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSEventsRule:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEventsRuleWithName retrieves all AWSEventsRule items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEventsRuleWithName(name string) (*AWSEventsRule, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSEventsRule:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEventsRule not found", name)
}

// GetAllAWSFSxFileSystemResources retrieves all AWSFSxFileSystem items from an AWS CloudFormation template
func (t *Template) GetAllAWSFSxFileSystemResources() map[string]*AWSFSxFileSystem {
	results := map[string]*AWSFSxFileSystem{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSFSxFileSystem:
			results[name] = resource
		}
	}
	return results
}

// GetAWSFSxFileSystemWithName retrieves all AWSFSxFileSystem items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSFSxFileSystemWithName(name string) (*AWSFSxFileSystem, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSFSxFileSystem:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSFSxFileSystem not found", name)
}

// GetAllAWSGameLiftAliasResources retrieves all AWSGameLiftAlias items from an AWS CloudFormation template
func (t *Template) GetAllAWSGameLiftAliasResources() map[string]*AWSGameLiftAlias {
	results := map[string]*AWSGameLiftAlias{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSGameLiftAlias:
			results[name] = resource
		}
	}
	return results
}

// GetAWSGameLiftAliasWithName retrieves all AWSGameLiftAlias items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSGameLiftAliasWithName(name string) (*AWSGameLiftAlias, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSGameLiftAlias:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSGameLiftAlias not found", name)
}

// GetAllAWSGameLiftBuildResources retrieves all AWSGameLiftBuild items from an AWS CloudFormation template
func (t *Template) GetAllAWSGameLiftBuildResources() map[string]*AWSGameLiftBuild {
	results := map[string]*AWSGameLiftBuild{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSGameLiftBuild:
			results[name] = resource
		}
	}
	return results
}

// GetAWSGameLiftBuildWithName retrieves all AWSGameLiftBuild items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSGameLiftBuildWithName(name string) (*AWSGameLiftBuild, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSGameLiftBuild:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSGameLiftBuild not found", name)
}

// GetAllAWSGameLiftFleetResources retrieves all AWSGameLiftFleet items from an AWS CloudFormation template
func (t *Template) GetAllAWSGameLiftFleetResources() map[string]*AWSGameLiftFleet {
	results := map[string]*AWSGameLiftFleet{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSGameLiftFleet:
			results[name] = resource
		}
	}
	return results
}

// GetAWSGameLiftFleetWithName retrieves all AWSGameLiftFleet items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSGameLiftFleetWithName(name string) (*AWSGameLiftFleet, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSGameLiftFleet:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSGameLiftFleet not found", name)
}

// GetAllAWSGlueClassifierResources retrieves all AWSGlueClassifier items from an AWS CloudFormation template
func (t *Template) GetAllAWSGlueClassifierResources() map[string]*AWSGlueClassifier {
	results := map[string]*AWSGlueClassifier{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSGlueClassifier:
			results[name] = resource
		}
	}
	return results
}

// GetAWSGlueClassifierWithName retrieves all AWSGlueClassifier items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSGlueClassifierWithName(name string) (*AWSGlueClassifier, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSGlueClassifier:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSGlueClassifier not found", name)
}

// GetAllAWSGlueConnectionResources retrieves all AWSGlueConnection items from an AWS CloudFormation template
func (t *Template) GetAllAWSGlueConnectionResources() map[string]*AWSGlueConnection {
	results := map[string]*AWSGlueConnection{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSGlueConnection:
			results[name] = resource
		}
	}
	return results
}

// GetAWSGlueConnectionWithName retrieves all AWSGlueConnection items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSGlueConnectionWithName(name string) (*AWSGlueConnection, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSGlueConnection:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSGlueConnection not found", name)
}

// GetAllAWSGlueCrawlerResources retrieves all AWSGlueCrawler items from an AWS CloudFormation template
func (t *Template) GetAllAWSGlueCrawlerResources() map[string]*AWSGlueCrawler {
	results := map[string]*AWSGlueCrawler{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSGlueCrawler:
			results[name] = resource
		}
	}
	return results
}

// GetAWSGlueCrawlerWithName retrieves all AWSGlueCrawler items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSGlueCrawlerWithName(name string) (*AWSGlueCrawler, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSGlueCrawler:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSGlueCrawler not found", name)
}

// GetAllAWSGlueDatabaseResources retrieves all AWSGlueDatabase items from an AWS CloudFormation template
func (t *Template) GetAllAWSGlueDatabaseResources() map[string]*AWSGlueDatabase {
	results := map[string]*AWSGlueDatabase{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSGlueDatabase:
			results[name] = resource
		}
	}
	return results
}

// GetAWSGlueDatabaseWithName retrieves all AWSGlueDatabase items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSGlueDatabaseWithName(name string) (*AWSGlueDatabase, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSGlueDatabase:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSGlueDatabase not found", name)
}

// GetAllAWSGlueDevEndpointResources retrieves all AWSGlueDevEndpoint items from an AWS CloudFormation template
func (t *Template) GetAllAWSGlueDevEndpointResources() map[string]*AWSGlueDevEndpoint {
	results := map[string]*AWSGlueDevEndpoint{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSGlueDevEndpoint:
			results[name] = resource
		}
	}
	return results
}

// GetAWSGlueDevEndpointWithName retrieves all AWSGlueDevEndpoint items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSGlueDevEndpointWithName(name string) (*AWSGlueDevEndpoint, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSGlueDevEndpoint:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSGlueDevEndpoint not found", name)
}

// GetAllAWSGlueJobResources retrieves all AWSGlueJob items from an AWS CloudFormation template
func (t *Template) GetAllAWSGlueJobResources() map[string]*AWSGlueJob {
	results := map[string]*AWSGlueJob{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSGlueJob:
			results[name] = resource
		}
	}
	return results
}

// GetAWSGlueJobWithName retrieves all AWSGlueJob items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSGlueJobWithName(name string) (*AWSGlueJob, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSGlueJob:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSGlueJob not found", name)
}

// GetAllAWSGluePartitionResources retrieves all AWSGluePartition items from an AWS CloudFormation template
func (t *Template) GetAllAWSGluePartitionResources() map[string]*AWSGluePartition {
	results := map[string]*AWSGluePartition{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSGluePartition:
			results[name] = resource
		}
	}
	return results
}

// GetAWSGluePartitionWithName retrieves all AWSGluePartition items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSGluePartitionWithName(name string) (*AWSGluePartition, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSGluePartition:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSGluePartition not found", name)
}

// GetAllAWSGlueTableResources retrieves all AWSGlueTable items from an AWS CloudFormation template
func (t *Template) GetAllAWSGlueTableResources() map[string]*AWSGlueTable {
	results := map[string]*AWSGlueTable{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSGlueTable:
			results[name] = resource
		}
	}
	return results
}

// GetAWSGlueTableWithName retrieves all AWSGlueTable items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSGlueTableWithName(name string) (*AWSGlueTable, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSGlueTable:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSGlueTable not found", name)
}

// GetAllAWSGlueTriggerResources retrieves all AWSGlueTrigger items from an AWS CloudFormation template
func (t *Template) GetAllAWSGlueTriggerResources() map[string]*AWSGlueTrigger {
	results := map[string]*AWSGlueTrigger{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSGlueTrigger:
			results[name] = resource
		}
	}
	return results
}

// GetAWSGlueTriggerWithName retrieves all AWSGlueTrigger items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSGlueTriggerWithName(name string) (*AWSGlueTrigger, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSGlueTrigger:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSGlueTrigger not found", name)
}

// GetAllAWSGuardDutyDetectorResources retrieves all AWSGuardDutyDetector items from an AWS CloudFormation template
func (t *Template) GetAllAWSGuardDutyDetectorResources() map[string]*AWSGuardDutyDetector {
	results := map[string]*AWSGuardDutyDetector{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSGuardDutyDetector:
			results[name] = resource
		}
	}
	return results
}

// GetAWSGuardDutyDetectorWithName retrieves all AWSGuardDutyDetector items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSGuardDutyDetectorWithName(name string) (*AWSGuardDutyDetector, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSGuardDutyDetector:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSGuardDutyDetector not found", name)
}

// GetAllAWSGuardDutyFilterResources retrieves all AWSGuardDutyFilter items from an AWS CloudFormation template
func (t *Template) GetAllAWSGuardDutyFilterResources() map[string]*AWSGuardDutyFilter {
	results := map[string]*AWSGuardDutyFilter{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSGuardDutyFilter:
			results[name] = resource
		}
	}
	return results
}

// GetAWSGuardDutyFilterWithName retrieves all AWSGuardDutyFilter items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSGuardDutyFilterWithName(name string) (*AWSGuardDutyFilter, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSGuardDutyFilter:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSGuardDutyFilter not found", name)
}

// GetAllAWSGuardDutyIPSetResources retrieves all AWSGuardDutyIPSet items from an AWS CloudFormation template
func (t *Template) GetAllAWSGuardDutyIPSetResources() map[string]*AWSGuardDutyIPSet {
	results := map[string]*AWSGuardDutyIPSet{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSGuardDutyIPSet:
			results[name] = resource
		}
	}
	return results
}

// GetAWSGuardDutyIPSetWithName retrieves all AWSGuardDutyIPSet items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSGuardDutyIPSetWithName(name string) (*AWSGuardDutyIPSet, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSGuardDutyIPSet:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSGuardDutyIPSet not found", name)
}

// GetAllAWSGuardDutyMasterResources retrieves all AWSGuardDutyMaster items from an AWS CloudFormation template
func (t *Template) GetAllAWSGuardDutyMasterResources() map[string]*AWSGuardDutyMaster {
	results := map[string]*AWSGuardDutyMaster{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSGuardDutyMaster:
			results[name] = resource
		}
	}
	return results
}

// GetAWSGuardDutyMasterWithName retrieves all AWSGuardDutyMaster items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSGuardDutyMasterWithName(name string) (*AWSGuardDutyMaster, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSGuardDutyMaster:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSGuardDutyMaster not found", name)
}

// GetAllAWSGuardDutyMemberResources retrieves all AWSGuardDutyMember items from an AWS CloudFormation template
func (t *Template) GetAllAWSGuardDutyMemberResources() map[string]*AWSGuardDutyMember {
	results := map[string]*AWSGuardDutyMember{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSGuardDutyMember:
			results[name] = resource
		}
	}
	return results
}

// GetAWSGuardDutyMemberWithName retrieves all AWSGuardDutyMember items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSGuardDutyMemberWithName(name string) (*AWSGuardDutyMember, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSGuardDutyMember:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSGuardDutyMember not found", name)
}

// GetAllAWSGuardDutyThreatIntelSetResources retrieves all AWSGuardDutyThreatIntelSet items from an AWS CloudFormation template
func (t *Template) GetAllAWSGuardDutyThreatIntelSetResources() map[string]*AWSGuardDutyThreatIntelSet {
	results := map[string]*AWSGuardDutyThreatIntelSet{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSGuardDutyThreatIntelSet:
			results[name] = resource
		}
	}
	return results
}

// GetAWSGuardDutyThreatIntelSetWithName retrieves all AWSGuardDutyThreatIntelSet items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSGuardDutyThreatIntelSetWithName(name string) (*AWSGuardDutyThreatIntelSet, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSGuardDutyThreatIntelSet:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSGuardDutyThreatIntelSet not found", name)
}

// GetAllAWSIAMAccessKeyResources retrieves all AWSIAMAccessKey items from an AWS CloudFormation template
func (t *Template) GetAllAWSIAMAccessKeyResources() map[string]*AWSIAMAccessKey {
	results := map[string]*AWSIAMAccessKey{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSIAMAccessKey:
			results[name] = resource
		}
	}
	return results
}

// GetAWSIAMAccessKeyWithName retrieves all AWSIAMAccessKey items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSIAMAccessKeyWithName(name string) (*AWSIAMAccessKey, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSIAMAccessKey:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSIAMAccessKey not found", name)
}

// GetAllAWSIAMGroupResources retrieves all AWSIAMGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSIAMGroupResources() map[string]*AWSIAMGroup {
	results := map[string]*AWSIAMGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSIAMGroup:
			results[name] = resource
		}
	}
	return results
}

// GetAWSIAMGroupWithName retrieves all AWSIAMGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSIAMGroupWithName(name string) (*AWSIAMGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSIAMGroup:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSIAMGroup not found", name)
}

// GetAllAWSIAMInstanceProfileResources retrieves all AWSIAMInstanceProfile items from an AWS CloudFormation template
func (t *Template) GetAllAWSIAMInstanceProfileResources() map[string]*AWSIAMInstanceProfile {
	results := map[string]*AWSIAMInstanceProfile{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSIAMInstanceProfile:
			results[name] = resource
		}
	}
	return results
}

// GetAWSIAMInstanceProfileWithName retrieves all AWSIAMInstanceProfile items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSIAMInstanceProfileWithName(name string) (*AWSIAMInstanceProfile, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSIAMInstanceProfile:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSIAMInstanceProfile not found", name)
}

// GetAllAWSIAMManagedPolicyResources retrieves all AWSIAMManagedPolicy items from an AWS CloudFormation template
func (t *Template) GetAllAWSIAMManagedPolicyResources() map[string]*AWSIAMManagedPolicy {
	results := map[string]*AWSIAMManagedPolicy{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSIAMManagedPolicy:
			results[name] = resource
		}
	}
	return results
}

// GetAWSIAMManagedPolicyWithName retrieves all AWSIAMManagedPolicy items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSIAMManagedPolicyWithName(name string) (*AWSIAMManagedPolicy, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSIAMManagedPolicy:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSIAMManagedPolicy not found", name)
}

// GetAllAWSIAMPolicyResources retrieves all AWSIAMPolicy items from an AWS CloudFormation template
func (t *Template) GetAllAWSIAMPolicyResources() map[string]*AWSIAMPolicy {
	results := map[string]*AWSIAMPolicy{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSIAMPolicy:
			results[name] = resource
		}
	}
	return results
}

// GetAWSIAMPolicyWithName retrieves all AWSIAMPolicy items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSIAMPolicyWithName(name string) (*AWSIAMPolicy, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSIAMPolicy:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSIAMPolicy not found", name)
}

// GetAllAWSIAMRoleResources retrieves all AWSIAMRole items from an AWS CloudFormation template
func (t *Template) GetAllAWSIAMRoleResources() map[string]*AWSIAMRole {
	results := map[string]*AWSIAMRole{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSIAMRole:
			results[name] = resource
		}
	}
	return results
}

// GetAWSIAMRoleWithName retrieves all AWSIAMRole items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSIAMRoleWithName(name string) (*AWSIAMRole, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSIAMRole:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSIAMRole not found", name)
}

// GetAllAWSIAMServiceLinkedRoleResources retrieves all AWSIAMServiceLinkedRole items from an AWS CloudFormation template
func (t *Template) GetAllAWSIAMServiceLinkedRoleResources() map[string]*AWSIAMServiceLinkedRole {
	results := map[string]*AWSIAMServiceLinkedRole{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSIAMServiceLinkedRole:
			results[name] = resource
		}
	}
	return results
}

// GetAWSIAMServiceLinkedRoleWithName retrieves all AWSIAMServiceLinkedRole items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSIAMServiceLinkedRoleWithName(name string) (*AWSIAMServiceLinkedRole, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSIAMServiceLinkedRole:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSIAMServiceLinkedRole not found", name)
}

// GetAllAWSIAMUserResources retrieves all AWSIAMUser items from an AWS CloudFormation template
func (t *Template) GetAllAWSIAMUserResources() map[string]*AWSIAMUser {
	results := map[string]*AWSIAMUser{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSIAMUser:
			results[name] = resource
		}
	}
	return results
}

// GetAWSIAMUserWithName retrieves all AWSIAMUser items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSIAMUserWithName(name string) (*AWSIAMUser, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSIAMUser:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSIAMUser not found", name)
}

// GetAllAWSIAMUserToGroupAdditionResources retrieves all AWSIAMUserToGroupAddition items from an AWS CloudFormation template
func (t *Template) GetAllAWSIAMUserToGroupAdditionResources() map[string]*AWSIAMUserToGroupAddition {
	results := map[string]*AWSIAMUserToGroupAddition{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSIAMUserToGroupAddition:
			results[name] = resource
		}
	}
	return results
}

// GetAWSIAMUserToGroupAdditionWithName retrieves all AWSIAMUserToGroupAddition items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSIAMUserToGroupAdditionWithName(name string) (*AWSIAMUserToGroupAddition, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSIAMUserToGroupAddition:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSIAMUserToGroupAddition not found", name)
}

// GetAllAWSInspectorAssessmentTargetResources retrieves all AWSInspectorAssessmentTarget items from an AWS CloudFormation template
func (t *Template) GetAllAWSInspectorAssessmentTargetResources() map[string]*AWSInspectorAssessmentTarget {
	results := map[string]*AWSInspectorAssessmentTarget{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSInspectorAssessmentTarget:
			results[name] = resource
		}
	}
	return results
}

// GetAWSInspectorAssessmentTargetWithName retrieves all AWSInspectorAssessmentTarget items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSInspectorAssessmentTargetWithName(name string) (*AWSInspectorAssessmentTarget, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSInspectorAssessmentTarget:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSInspectorAssessmentTarget not found", name)
}

// GetAllAWSInspectorAssessmentTemplateResources retrieves all AWSInspectorAssessmentTemplate items from an AWS CloudFormation template
func (t *Template) GetAllAWSInspectorAssessmentTemplateResources() map[string]*AWSInspectorAssessmentTemplate {
	results := map[string]*AWSInspectorAssessmentTemplate{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSInspectorAssessmentTemplate:
			results[name] = resource
		}
	}
	return results
}

// GetAWSInspectorAssessmentTemplateWithName retrieves all AWSInspectorAssessmentTemplate items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSInspectorAssessmentTemplateWithName(name string) (*AWSInspectorAssessmentTemplate, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSInspectorAssessmentTemplate:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSInspectorAssessmentTemplate not found", name)
}

// GetAllAWSInspectorResourceGroupResources retrieves all AWSInspectorResourceGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSInspectorResourceGroupResources() map[string]*AWSInspectorResourceGroup {
	results := map[string]*AWSInspectorResourceGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSInspectorResourceGroup:
			results[name] = resource
		}
	}
	return results
}

// GetAWSInspectorResourceGroupWithName retrieves all AWSInspectorResourceGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSInspectorResourceGroupWithName(name string) (*AWSInspectorResourceGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSInspectorResourceGroup:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSInspectorResourceGroup not found", name)
}

// GetAllAWSIoT1ClickDeviceResources retrieves all AWSIoT1ClickDevice items from an AWS CloudFormation template
func (t *Template) GetAllAWSIoT1ClickDeviceResources() map[string]*AWSIoT1ClickDevice {
	results := map[string]*AWSIoT1ClickDevice{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSIoT1ClickDevice:
			results[name] = resource
		}
	}
	return results
}

// GetAWSIoT1ClickDeviceWithName retrieves all AWSIoT1ClickDevice items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSIoT1ClickDeviceWithName(name string) (*AWSIoT1ClickDevice, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSIoT1ClickDevice:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSIoT1ClickDevice not found", name)
}

// GetAllAWSIoT1ClickPlacementResources retrieves all AWSIoT1ClickPlacement items from an AWS CloudFormation template
func (t *Template) GetAllAWSIoT1ClickPlacementResources() map[string]*AWSIoT1ClickPlacement {
	results := map[string]*AWSIoT1ClickPlacement{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSIoT1ClickPlacement:
			results[name] = resource
		}
	}
	return results
}

// GetAWSIoT1ClickPlacementWithName retrieves all AWSIoT1ClickPlacement items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSIoT1ClickPlacementWithName(name string) (*AWSIoT1ClickPlacement, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSIoT1ClickPlacement:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSIoT1ClickPlacement not found", name)
}

// GetAllAWSIoT1ClickProjectResources retrieves all AWSIoT1ClickProject items from an AWS CloudFormation template
func (t *Template) GetAllAWSIoT1ClickProjectResources() map[string]*AWSIoT1ClickProject {
	results := map[string]*AWSIoT1ClickProject{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSIoT1ClickProject:
			results[name] = resource
		}
	}
	return results
}

// GetAWSIoT1ClickProjectWithName retrieves all AWSIoT1ClickProject items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSIoT1ClickProjectWithName(name string) (*AWSIoT1ClickProject, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSIoT1ClickProject:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSIoT1ClickProject not found", name)
}

// GetAllAWSIoTCertificateResources retrieves all AWSIoTCertificate items from an AWS CloudFormation template
func (t *Template) GetAllAWSIoTCertificateResources() map[string]*AWSIoTCertificate {
	results := map[string]*AWSIoTCertificate{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSIoTCertificate:
			results[name] = resource
		}
	}
	return results
}

// GetAWSIoTCertificateWithName retrieves all AWSIoTCertificate items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSIoTCertificateWithName(name string) (*AWSIoTCertificate, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSIoTCertificate:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSIoTCertificate not found", name)
}

// GetAllAWSIoTPolicyResources retrieves all AWSIoTPolicy items from an AWS CloudFormation template
func (t *Template) GetAllAWSIoTPolicyResources() map[string]*AWSIoTPolicy {
	results := map[string]*AWSIoTPolicy{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSIoTPolicy:
			results[name] = resource
		}
	}
	return results
}

// GetAWSIoTPolicyWithName retrieves all AWSIoTPolicy items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSIoTPolicyWithName(name string) (*AWSIoTPolicy, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSIoTPolicy:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSIoTPolicy not found", name)
}

// GetAllAWSIoTPolicyPrincipalAttachmentResources retrieves all AWSIoTPolicyPrincipalAttachment items from an AWS CloudFormation template
func (t *Template) GetAllAWSIoTPolicyPrincipalAttachmentResources() map[string]*AWSIoTPolicyPrincipalAttachment {
	results := map[string]*AWSIoTPolicyPrincipalAttachment{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSIoTPolicyPrincipalAttachment:
			results[name] = resource
		}
	}
	return results
}

// GetAWSIoTPolicyPrincipalAttachmentWithName retrieves all AWSIoTPolicyPrincipalAttachment items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSIoTPolicyPrincipalAttachmentWithName(name string) (*AWSIoTPolicyPrincipalAttachment, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSIoTPolicyPrincipalAttachment:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSIoTPolicyPrincipalAttachment not found", name)
}

// GetAllAWSIoTThingResources retrieves all AWSIoTThing items from an AWS CloudFormation template
func (t *Template) GetAllAWSIoTThingResources() map[string]*AWSIoTThing {
	results := map[string]*AWSIoTThing{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSIoTThing:
			results[name] = resource
		}
	}
	return results
}

// GetAWSIoTThingWithName retrieves all AWSIoTThing items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSIoTThingWithName(name string) (*AWSIoTThing, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSIoTThing:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSIoTThing not found", name)
}

// GetAllAWSIoTThingPrincipalAttachmentResources retrieves all AWSIoTThingPrincipalAttachment items from an AWS CloudFormation template
func (t *Template) GetAllAWSIoTThingPrincipalAttachmentResources() map[string]*AWSIoTThingPrincipalAttachment {
	results := map[string]*AWSIoTThingPrincipalAttachment{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSIoTThingPrincipalAttachment:
			results[name] = resource
		}
	}
	return results
}

// GetAWSIoTThingPrincipalAttachmentWithName retrieves all AWSIoTThingPrincipalAttachment items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSIoTThingPrincipalAttachmentWithName(name string) (*AWSIoTThingPrincipalAttachment, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSIoTThingPrincipalAttachment:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSIoTThingPrincipalAttachment not found", name)
}

// GetAllAWSIoTTopicRuleResources retrieves all AWSIoTTopicRule items from an AWS CloudFormation template
func (t *Template) GetAllAWSIoTTopicRuleResources() map[string]*AWSIoTTopicRule {
	results := map[string]*AWSIoTTopicRule{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSIoTTopicRule:
			results[name] = resource
		}
	}
	return results
}

// GetAWSIoTTopicRuleWithName retrieves all AWSIoTTopicRule items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSIoTTopicRuleWithName(name string) (*AWSIoTTopicRule, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSIoTTopicRule:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSIoTTopicRule not found", name)
}

// GetAllAWSIoTAnalyticsChannelResources retrieves all AWSIoTAnalyticsChannel items from an AWS CloudFormation template
func (t *Template) GetAllAWSIoTAnalyticsChannelResources() map[string]*AWSIoTAnalyticsChannel {
	results := map[string]*AWSIoTAnalyticsChannel{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSIoTAnalyticsChannel:
			results[name] = resource
		}
	}
	return results
}

// GetAWSIoTAnalyticsChannelWithName retrieves all AWSIoTAnalyticsChannel items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSIoTAnalyticsChannelWithName(name string) (*AWSIoTAnalyticsChannel, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSIoTAnalyticsChannel:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSIoTAnalyticsChannel not found", name)
}

// GetAllAWSIoTAnalyticsDatasetResources retrieves all AWSIoTAnalyticsDataset items from an AWS CloudFormation template
func (t *Template) GetAllAWSIoTAnalyticsDatasetResources() map[string]*AWSIoTAnalyticsDataset {
	results := map[string]*AWSIoTAnalyticsDataset{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSIoTAnalyticsDataset:
			results[name] = resource
		}
	}
	return results
}

// GetAWSIoTAnalyticsDatasetWithName retrieves all AWSIoTAnalyticsDataset items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSIoTAnalyticsDatasetWithName(name string) (*AWSIoTAnalyticsDataset, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSIoTAnalyticsDataset:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSIoTAnalyticsDataset not found", name)
}

// GetAllAWSIoTAnalyticsDatastoreResources retrieves all AWSIoTAnalyticsDatastore items from an AWS CloudFormation template
func (t *Template) GetAllAWSIoTAnalyticsDatastoreResources() map[string]*AWSIoTAnalyticsDatastore {
	results := map[string]*AWSIoTAnalyticsDatastore{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSIoTAnalyticsDatastore:
			results[name] = resource
		}
	}
	return results
}

// GetAWSIoTAnalyticsDatastoreWithName retrieves all AWSIoTAnalyticsDatastore items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSIoTAnalyticsDatastoreWithName(name string) (*AWSIoTAnalyticsDatastore, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSIoTAnalyticsDatastore:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSIoTAnalyticsDatastore not found", name)
}

// GetAllAWSIoTAnalyticsPipelineResources retrieves all AWSIoTAnalyticsPipeline items from an AWS CloudFormation template
func (t *Template) GetAllAWSIoTAnalyticsPipelineResources() map[string]*AWSIoTAnalyticsPipeline {
	results := map[string]*AWSIoTAnalyticsPipeline{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSIoTAnalyticsPipeline:
			results[name] = resource
		}
	}
	return results
}

// GetAWSIoTAnalyticsPipelineWithName retrieves all AWSIoTAnalyticsPipeline items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSIoTAnalyticsPipelineWithName(name string) (*AWSIoTAnalyticsPipeline, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSIoTAnalyticsPipeline:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSIoTAnalyticsPipeline not found", name)
}

// GetAllAWSKMSAliasResources retrieves all AWSKMSAlias items from an AWS CloudFormation template
func (t *Template) GetAllAWSKMSAliasResources() map[string]*AWSKMSAlias {
	results := map[string]*AWSKMSAlias{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSKMSAlias:
			results[name] = resource
		}
	}
	return results
}

// GetAWSKMSAliasWithName retrieves all AWSKMSAlias items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSKMSAliasWithName(name string) (*AWSKMSAlias, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSKMSAlias:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSKMSAlias not found", name)
}

// GetAllAWSKMSKeyResources retrieves all AWSKMSKey items from an AWS CloudFormation template
func (t *Template) GetAllAWSKMSKeyResources() map[string]*AWSKMSKey {
	results := map[string]*AWSKMSKey{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSKMSKey:
			results[name] = resource
		}
	}
	return results
}

// GetAWSKMSKeyWithName retrieves all AWSKMSKey items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSKMSKeyWithName(name string) (*AWSKMSKey, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSKMSKey:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSKMSKey not found", name)
}

// GetAllAWSKinesisStreamResources retrieves all AWSKinesisStream items from an AWS CloudFormation template
func (t *Template) GetAllAWSKinesisStreamResources() map[string]*AWSKinesisStream {
	results := map[string]*AWSKinesisStream{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSKinesisStream:
			results[name] = resource
		}
	}
	return results
}

// GetAWSKinesisStreamWithName retrieves all AWSKinesisStream items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSKinesisStreamWithName(name string) (*AWSKinesisStream, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSKinesisStream:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSKinesisStream not found", name)
}

// GetAllAWSKinesisStreamConsumerResources retrieves all AWSKinesisStreamConsumer items from an AWS CloudFormation template
func (t *Template) GetAllAWSKinesisStreamConsumerResources() map[string]*AWSKinesisStreamConsumer {
	results := map[string]*AWSKinesisStreamConsumer{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSKinesisStreamConsumer:
			results[name] = resource
		}
	}
	return results
}

// GetAWSKinesisStreamConsumerWithName retrieves all AWSKinesisStreamConsumer items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSKinesisStreamConsumerWithName(name string) (*AWSKinesisStreamConsumer, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSKinesisStreamConsumer:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSKinesisStreamConsumer not found", name)
}

// GetAllAWSKinesisAnalyticsApplicationResources retrieves all AWSKinesisAnalyticsApplication items from an AWS CloudFormation template
func (t *Template) GetAllAWSKinesisAnalyticsApplicationResources() map[string]*AWSKinesisAnalyticsApplication {
	results := map[string]*AWSKinesisAnalyticsApplication{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSKinesisAnalyticsApplication:
			results[name] = resource
		}
	}
	return results
}

// GetAWSKinesisAnalyticsApplicationWithName retrieves all AWSKinesisAnalyticsApplication items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSKinesisAnalyticsApplicationWithName(name string) (*AWSKinesisAnalyticsApplication, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSKinesisAnalyticsApplication:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSKinesisAnalyticsApplication not found", name)
}

// GetAllAWSKinesisAnalyticsApplicationOutputResources retrieves all AWSKinesisAnalyticsApplicationOutput items from an AWS CloudFormation template
func (t *Template) GetAllAWSKinesisAnalyticsApplicationOutputResources() map[string]*AWSKinesisAnalyticsApplicationOutput {
	results := map[string]*AWSKinesisAnalyticsApplicationOutput{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSKinesisAnalyticsApplicationOutput:
			results[name] = resource
		}
	}
	return results
}

// GetAWSKinesisAnalyticsApplicationOutputWithName retrieves all AWSKinesisAnalyticsApplicationOutput items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSKinesisAnalyticsApplicationOutputWithName(name string) (*AWSKinesisAnalyticsApplicationOutput, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSKinesisAnalyticsApplicationOutput:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSKinesisAnalyticsApplicationOutput not found", name)
}

// GetAllAWSKinesisAnalyticsApplicationReferenceDataSourceResources retrieves all AWSKinesisAnalyticsApplicationReferenceDataSource items from an AWS CloudFormation template
func (t *Template) GetAllAWSKinesisAnalyticsApplicationReferenceDataSourceResources() map[string]*AWSKinesisAnalyticsApplicationReferenceDataSource {
	results := map[string]*AWSKinesisAnalyticsApplicationReferenceDataSource{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSKinesisAnalyticsApplicationReferenceDataSource:
			results[name] = resource
		}
	}
	return results
}

// GetAWSKinesisAnalyticsApplicationReferenceDataSourceWithName retrieves all AWSKinesisAnalyticsApplicationReferenceDataSource items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSKinesisAnalyticsApplicationReferenceDataSourceWithName(name string) (*AWSKinesisAnalyticsApplicationReferenceDataSource, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSKinesisAnalyticsApplicationReferenceDataSource:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSKinesisAnalyticsApplicationReferenceDataSource not found", name)
}

// GetAllAWSKinesisAnalyticsV2ApplicationResources retrieves all AWSKinesisAnalyticsV2Application items from an AWS CloudFormation template
func (t *Template) GetAllAWSKinesisAnalyticsV2ApplicationResources() map[string]*AWSKinesisAnalyticsV2Application {
	results := map[string]*AWSKinesisAnalyticsV2Application{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSKinesisAnalyticsV2Application:
			results[name] = resource
		}
	}
	return results
}

// GetAWSKinesisAnalyticsV2ApplicationWithName retrieves all AWSKinesisAnalyticsV2Application items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSKinesisAnalyticsV2ApplicationWithName(name string) (*AWSKinesisAnalyticsV2Application, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSKinesisAnalyticsV2Application:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSKinesisAnalyticsV2Application not found", name)
}

// GetAllAWSKinesisAnalyticsV2ApplicationCloudWatchLoggingOptionResources retrieves all AWSKinesisAnalyticsV2ApplicationCloudWatchLoggingOption items from an AWS CloudFormation template
func (t *Template) GetAllAWSKinesisAnalyticsV2ApplicationCloudWatchLoggingOptionResources() map[string]*AWSKinesisAnalyticsV2ApplicationCloudWatchLoggingOption {
	results := map[string]*AWSKinesisAnalyticsV2ApplicationCloudWatchLoggingOption{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSKinesisAnalyticsV2ApplicationCloudWatchLoggingOption:
			results[name] = resource
		}
	}
	return results
}

// GetAWSKinesisAnalyticsV2ApplicationCloudWatchLoggingOptionWithName retrieves all AWSKinesisAnalyticsV2ApplicationCloudWatchLoggingOption items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSKinesisAnalyticsV2ApplicationCloudWatchLoggingOptionWithName(name string) (*AWSKinesisAnalyticsV2ApplicationCloudWatchLoggingOption, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSKinesisAnalyticsV2ApplicationCloudWatchLoggingOption:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSKinesisAnalyticsV2ApplicationCloudWatchLoggingOption not found", name)
}

// GetAllAWSKinesisAnalyticsV2ApplicationOutputResources retrieves all AWSKinesisAnalyticsV2ApplicationOutput items from an AWS CloudFormation template
func (t *Template) GetAllAWSKinesisAnalyticsV2ApplicationOutputResources() map[string]*AWSKinesisAnalyticsV2ApplicationOutput {
	results := map[string]*AWSKinesisAnalyticsV2ApplicationOutput{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSKinesisAnalyticsV2ApplicationOutput:
			results[name] = resource
		}
	}
	return results
}

// GetAWSKinesisAnalyticsV2ApplicationOutputWithName retrieves all AWSKinesisAnalyticsV2ApplicationOutput items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSKinesisAnalyticsV2ApplicationOutputWithName(name string) (*AWSKinesisAnalyticsV2ApplicationOutput, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSKinesisAnalyticsV2ApplicationOutput:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSKinesisAnalyticsV2ApplicationOutput not found", name)
}

// GetAllAWSKinesisAnalyticsV2ApplicationReferenceDataSourceResources retrieves all AWSKinesisAnalyticsV2ApplicationReferenceDataSource items from an AWS CloudFormation template
func (t *Template) GetAllAWSKinesisAnalyticsV2ApplicationReferenceDataSourceResources() map[string]*AWSKinesisAnalyticsV2ApplicationReferenceDataSource {
	results := map[string]*AWSKinesisAnalyticsV2ApplicationReferenceDataSource{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSKinesisAnalyticsV2ApplicationReferenceDataSource:
			results[name] = resource
		}
	}
	return results
}

// GetAWSKinesisAnalyticsV2ApplicationReferenceDataSourceWithName retrieves all AWSKinesisAnalyticsV2ApplicationReferenceDataSource items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSKinesisAnalyticsV2ApplicationReferenceDataSourceWithName(name string) (*AWSKinesisAnalyticsV2ApplicationReferenceDataSource, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSKinesisAnalyticsV2ApplicationReferenceDataSource:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSKinesisAnalyticsV2ApplicationReferenceDataSource not found", name)
}

// GetAllAWSKinesisFirehoseDeliveryStreamResources retrieves all AWSKinesisFirehoseDeliveryStream items from an AWS CloudFormation template
func (t *Template) GetAllAWSKinesisFirehoseDeliveryStreamResources() map[string]*AWSKinesisFirehoseDeliveryStream {
	results := map[string]*AWSKinesisFirehoseDeliveryStream{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSKinesisFirehoseDeliveryStream:
			results[name] = resource
		}
	}
	return results
}

// GetAWSKinesisFirehoseDeliveryStreamWithName retrieves all AWSKinesisFirehoseDeliveryStream items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSKinesisFirehoseDeliveryStreamWithName(name string) (*AWSKinesisFirehoseDeliveryStream, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSKinesisFirehoseDeliveryStream:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSKinesisFirehoseDeliveryStream not found", name)
}

// GetAllAWSLambdaAliasResources retrieves all AWSLambdaAlias items from an AWS CloudFormation template
func (t *Template) GetAllAWSLambdaAliasResources() map[string]*AWSLambdaAlias {
	results := map[string]*AWSLambdaAlias{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSLambdaAlias:
			results[name] = resource
		}
	}
	return results
}

// GetAWSLambdaAliasWithName retrieves all AWSLambdaAlias items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSLambdaAliasWithName(name string) (*AWSLambdaAlias, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSLambdaAlias:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSLambdaAlias not found", name)
}

// GetAllAWSLambdaEventSourceMappingResources retrieves all AWSLambdaEventSourceMapping items from an AWS CloudFormation template
func (t *Template) GetAllAWSLambdaEventSourceMappingResources() map[string]*AWSLambdaEventSourceMapping {
	results := map[string]*AWSLambdaEventSourceMapping{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSLambdaEventSourceMapping:
			results[name] = resource
		}
	}
	return results
}

// GetAWSLambdaEventSourceMappingWithName retrieves all AWSLambdaEventSourceMapping items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSLambdaEventSourceMappingWithName(name string) (*AWSLambdaEventSourceMapping, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSLambdaEventSourceMapping:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSLambdaEventSourceMapping not found", name)
}

// GetAllAWSLambdaFunctionResources retrieves all AWSLambdaFunction items from an AWS CloudFormation template
func (t *Template) GetAllAWSLambdaFunctionResources() map[string]*AWSLambdaFunction {
	results := map[string]*AWSLambdaFunction{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSLambdaFunction:
			results[name] = resource
		}
	}
	return results
}

// GetAWSLambdaFunctionWithName retrieves all AWSLambdaFunction items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSLambdaFunctionWithName(name string) (*AWSLambdaFunction, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSLambdaFunction:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSLambdaFunction not found", name)
}

// GetAllAWSLambdaLayerVersionResources retrieves all AWSLambdaLayerVersion items from an AWS CloudFormation template
func (t *Template) GetAllAWSLambdaLayerVersionResources() map[string]*AWSLambdaLayerVersion {
	results := map[string]*AWSLambdaLayerVersion{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSLambdaLayerVersion:
			results[name] = resource
		}
	}
	return results
}

// GetAWSLambdaLayerVersionWithName retrieves all AWSLambdaLayerVersion items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSLambdaLayerVersionWithName(name string) (*AWSLambdaLayerVersion, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSLambdaLayerVersion:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSLambdaLayerVersion not found", name)
}

// GetAllAWSLambdaLayerVersionPermissionResources retrieves all AWSLambdaLayerVersionPermission items from an AWS CloudFormation template
func (t *Template) GetAllAWSLambdaLayerVersionPermissionResources() map[string]*AWSLambdaLayerVersionPermission {
	results := map[string]*AWSLambdaLayerVersionPermission{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSLambdaLayerVersionPermission:
			results[name] = resource
		}
	}
	return results
}

// GetAWSLambdaLayerVersionPermissionWithName retrieves all AWSLambdaLayerVersionPermission items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSLambdaLayerVersionPermissionWithName(name string) (*AWSLambdaLayerVersionPermission, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSLambdaLayerVersionPermission:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSLambdaLayerVersionPermission not found", name)
}

// GetAllAWSLambdaPermissionResources retrieves all AWSLambdaPermission items from an AWS CloudFormation template
func (t *Template) GetAllAWSLambdaPermissionResources() map[string]*AWSLambdaPermission {
	results := map[string]*AWSLambdaPermission{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSLambdaPermission:
			results[name] = resource
		}
	}
	return results
}

// GetAWSLambdaPermissionWithName retrieves all AWSLambdaPermission items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSLambdaPermissionWithName(name string) (*AWSLambdaPermission, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSLambdaPermission:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSLambdaPermission not found", name)
}

// GetAllAWSLambdaVersionResources retrieves all AWSLambdaVersion items from an AWS CloudFormation template
func (t *Template) GetAllAWSLambdaVersionResources() map[string]*AWSLambdaVersion {
	results := map[string]*AWSLambdaVersion{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSLambdaVersion:
			results[name] = resource
		}
	}
	return results
}

// GetAWSLambdaVersionWithName retrieves all AWSLambdaVersion items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSLambdaVersionWithName(name string) (*AWSLambdaVersion, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSLambdaVersion:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSLambdaVersion not found", name)
}

// GetAllAWSLogsDestinationResources retrieves all AWSLogsDestination items from an AWS CloudFormation template
func (t *Template) GetAllAWSLogsDestinationResources() map[string]*AWSLogsDestination {
	results := map[string]*AWSLogsDestination{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSLogsDestination:
			results[name] = resource
		}
	}
	return results
}

// GetAWSLogsDestinationWithName retrieves all AWSLogsDestination items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSLogsDestinationWithName(name string) (*AWSLogsDestination, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSLogsDestination:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSLogsDestination not found", name)
}

// GetAllAWSLogsLogGroupResources retrieves all AWSLogsLogGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSLogsLogGroupResources() map[string]*AWSLogsLogGroup {
	results := map[string]*AWSLogsLogGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSLogsLogGroup:
			results[name] = resource
		}
	}
	return results
}

// GetAWSLogsLogGroupWithName retrieves all AWSLogsLogGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSLogsLogGroupWithName(name string) (*AWSLogsLogGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSLogsLogGroup:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSLogsLogGroup not found", name)
}

// GetAllAWSLogsLogStreamResources retrieves all AWSLogsLogStream items from an AWS CloudFormation template
func (t *Template) GetAllAWSLogsLogStreamResources() map[string]*AWSLogsLogStream {
	results := map[string]*AWSLogsLogStream{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSLogsLogStream:
			results[name] = resource
		}
	}
	return results
}

// GetAWSLogsLogStreamWithName retrieves all AWSLogsLogStream items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSLogsLogStreamWithName(name string) (*AWSLogsLogStream, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSLogsLogStream:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSLogsLogStream not found", name)
}

// GetAllAWSLogsMetricFilterResources retrieves all AWSLogsMetricFilter items from an AWS CloudFormation template
func (t *Template) GetAllAWSLogsMetricFilterResources() map[string]*AWSLogsMetricFilter {
	results := map[string]*AWSLogsMetricFilter{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSLogsMetricFilter:
			results[name] = resource
		}
	}
	return results
}

// GetAWSLogsMetricFilterWithName retrieves all AWSLogsMetricFilter items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSLogsMetricFilterWithName(name string) (*AWSLogsMetricFilter, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSLogsMetricFilter:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSLogsMetricFilter not found", name)
}

// GetAllAWSLogsSubscriptionFilterResources retrieves all AWSLogsSubscriptionFilter items from an AWS CloudFormation template
func (t *Template) GetAllAWSLogsSubscriptionFilterResources() map[string]*AWSLogsSubscriptionFilter {
	results := map[string]*AWSLogsSubscriptionFilter{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSLogsSubscriptionFilter:
			results[name] = resource
		}
	}
	return results
}

// GetAWSLogsSubscriptionFilterWithName retrieves all AWSLogsSubscriptionFilter items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSLogsSubscriptionFilterWithName(name string) (*AWSLogsSubscriptionFilter, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSLogsSubscriptionFilter:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSLogsSubscriptionFilter not found", name)
}

// GetAllAWSNeptuneDBClusterResources retrieves all AWSNeptuneDBCluster items from an AWS CloudFormation template
func (t *Template) GetAllAWSNeptuneDBClusterResources() map[string]*AWSNeptuneDBCluster {
	results := map[string]*AWSNeptuneDBCluster{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSNeptuneDBCluster:
			results[name] = resource
		}
	}
	return results
}

// GetAWSNeptuneDBClusterWithName retrieves all AWSNeptuneDBCluster items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSNeptuneDBClusterWithName(name string) (*AWSNeptuneDBCluster, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSNeptuneDBCluster:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSNeptuneDBCluster not found", name)
}

// GetAllAWSNeptuneDBClusterParameterGroupResources retrieves all AWSNeptuneDBClusterParameterGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSNeptuneDBClusterParameterGroupResources() map[string]*AWSNeptuneDBClusterParameterGroup {
	results := map[string]*AWSNeptuneDBClusterParameterGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSNeptuneDBClusterParameterGroup:
			results[name] = resource
		}
	}
	return results
}

// GetAWSNeptuneDBClusterParameterGroupWithName retrieves all AWSNeptuneDBClusterParameterGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSNeptuneDBClusterParameterGroupWithName(name string) (*AWSNeptuneDBClusterParameterGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSNeptuneDBClusterParameterGroup:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSNeptuneDBClusterParameterGroup not found", name)
}

// GetAllAWSNeptuneDBInstanceResources retrieves all AWSNeptuneDBInstance items from an AWS CloudFormation template
func (t *Template) GetAllAWSNeptuneDBInstanceResources() map[string]*AWSNeptuneDBInstance {
	results := map[string]*AWSNeptuneDBInstance{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSNeptuneDBInstance:
			results[name] = resource
		}
	}
	return results
}

// GetAWSNeptuneDBInstanceWithName retrieves all AWSNeptuneDBInstance items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSNeptuneDBInstanceWithName(name string) (*AWSNeptuneDBInstance, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSNeptuneDBInstance:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSNeptuneDBInstance not found", name)
}

// GetAllAWSNeptuneDBParameterGroupResources retrieves all AWSNeptuneDBParameterGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSNeptuneDBParameterGroupResources() map[string]*AWSNeptuneDBParameterGroup {
	results := map[string]*AWSNeptuneDBParameterGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSNeptuneDBParameterGroup:
			results[name] = resource
		}
	}
	return results
}

// GetAWSNeptuneDBParameterGroupWithName retrieves all AWSNeptuneDBParameterGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSNeptuneDBParameterGroupWithName(name string) (*AWSNeptuneDBParameterGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSNeptuneDBParameterGroup:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSNeptuneDBParameterGroup not found", name)
}

// GetAllAWSNeptuneDBSubnetGroupResources retrieves all AWSNeptuneDBSubnetGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSNeptuneDBSubnetGroupResources() map[string]*AWSNeptuneDBSubnetGroup {
	results := map[string]*AWSNeptuneDBSubnetGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSNeptuneDBSubnetGroup:
			results[name] = resource
		}
	}
	return results
}

// GetAWSNeptuneDBSubnetGroupWithName retrieves all AWSNeptuneDBSubnetGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSNeptuneDBSubnetGroupWithName(name string) (*AWSNeptuneDBSubnetGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSNeptuneDBSubnetGroup:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSNeptuneDBSubnetGroup not found", name)
}

// GetAllAWSOpsWorksAppResources retrieves all AWSOpsWorksApp items from an AWS CloudFormation template
func (t *Template) GetAllAWSOpsWorksAppResources() map[string]*AWSOpsWorksApp {
	results := map[string]*AWSOpsWorksApp{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSOpsWorksApp:
			results[name] = resource
		}
	}
	return results
}

// GetAWSOpsWorksAppWithName retrieves all AWSOpsWorksApp items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSOpsWorksAppWithName(name string) (*AWSOpsWorksApp, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSOpsWorksApp:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSOpsWorksApp not found", name)
}

// GetAllAWSOpsWorksElasticLoadBalancerAttachmentResources retrieves all AWSOpsWorksElasticLoadBalancerAttachment items from an AWS CloudFormation template
func (t *Template) GetAllAWSOpsWorksElasticLoadBalancerAttachmentResources() map[string]*AWSOpsWorksElasticLoadBalancerAttachment {
	results := map[string]*AWSOpsWorksElasticLoadBalancerAttachment{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSOpsWorksElasticLoadBalancerAttachment:
			results[name] = resource
		}
	}
	return results
}

// GetAWSOpsWorksElasticLoadBalancerAttachmentWithName retrieves all AWSOpsWorksElasticLoadBalancerAttachment items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSOpsWorksElasticLoadBalancerAttachmentWithName(name string) (*AWSOpsWorksElasticLoadBalancerAttachment, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSOpsWorksElasticLoadBalancerAttachment:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSOpsWorksElasticLoadBalancerAttachment not found", name)
}

// GetAllAWSOpsWorksInstanceResources retrieves all AWSOpsWorksInstance items from an AWS CloudFormation template
func (t *Template) GetAllAWSOpsWorksInstanceResources() map[string]*AWSOpsWorksInstance {
	results := map[string]*AWSOpsWorksInstance{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSOpsWorksInstance:
			results[name] = resource
		}
	}
	return results
}

// GetAWSOpsWorksInstanceWithName retrieves all AWSOpsWorksInstance items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSOpsWorksInstanceWithName(name string) (*AWSOpsWorksInstance, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSOpsWorksInstance:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSOpsWorksInstance not found", name)
}

// GetAllAWSOpsWorksLayerResources retrieves all AWSOpsWorksLayer items from an AWS CloudFormation template
func (t *Template) GetAllAWSOpsWorksLayerResources() map[string]*AWSOpsWorksLayer {
	results := map[string]*AWSOpsWorksLayer{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSOpsWorksLayer:
			results[name] = resource
		}
	}
	return results
}

// GetAWSOpsWorksLayerWithName retrieves all AWSOpsWorksLayer items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSOpsWorksLayerWithName(name string) (*AWSOpsWorksLayer, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSOpsWorksLayer:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSOpsWorksLayer not found", name)
}

// GetAllAWSOpsWorksStackResources retrieves all AWSOpsWorksStack items from an AWS CloudFormation template
func (t *Template) GetAllAWSOpsWorksStackResources() map[string]*AWSOpsWorksStack {
	results := map[string]*AWSOpsWorksStack{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSOpsWorksStack:
			results[name] = resource
		}
	}
	return results
}

// GetAWSOpsWorksStackWithName retrieves all AWSOpsWorksStack items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSOpsWorksStackWithName(name string) (*AWSOpsWorksStack, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSOpsWorksStack:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSOpsWorksStack not found", name)
}

// GetAllAWSOpsWorksUserProfileResources retrieves all AWSOpsWorksUserProfile items from an AWS CloudFormation template
func (t *Template) GetAllAWSOpsWorksUserProfileResources() map[string]*AWSOpsWorksUserProfile {
	results := map[string]*AWSOpsWorksUserProfile{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSOpsWorksUserProfile:
			results[name] = resource
		}
	}
	return results
}

// GetAWSOpsWorksUserProfileWithName retrieves all AWSOpsWorksUserProfile items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSOpsWorksUserProfileWithName(name string) (*AWSOpsWorksUserProfile, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSOpsWorksUserProfile:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSOpsWorksUserProfile not found", name)
}

// GetAllAWSOpsWorksVolumeResources retrieves all AWSOpsWorksVolume items from an AWS CloudFormation template
func (t *Template) GetAllAWSOpsWorksVolumeResources() map[string]*AWSOpsWorksVolume {
	results := map[string]*AWSOpsWorksVolume{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSOpsWorksVolume:
			results[name] = resource
		}
	}
	return results
}

// GetAWSOpsWorksVolumeWithName retrieves all AWSOpsWorksVolume items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSOpsWorksVolumeWithName(name string) (*AWSOpsWorksVolume, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSOpsWorksVolume:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSOpsWorksVolume not found", name)
}

// GetAllAWSOpsWorksCMServerResources retrieves all AWSOpsWorksCMServer items from an AWS CloudFormation template
func (t *Template) GetAllAWSOpsWorksCMServerResources() map[string]*AWSOpsWorksCMServer {
	results := map[string]*AWSOpsWorksCMServer{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSOpsWorksCMServer:
			results[name] = resource
		}
	}
	return results
}

// GetAWSOpsWorksCMServerWithName retrieves all AWSOpsWorksCMServer items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSOpsWorksCMServerWithName(name string) (*AWSOpsWorksCMServer, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSOpsWorksCMServer:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSOpsWorksCMServer not found", name)
}

// GetAllAWSRAMResourceShareResources retrieves all AWSRAMResourceShare items from an AWS CloudFormation template
func (t *Template) GetAllAWSRAMResourceShareResources() map[string]*AWSRAMResourceShare {
	results := map[string]*AWSRAMResourceShare{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSRAMResourceShare:
			results[name] = resource
		}
	}
	return results
}

// GetAWSRAMResourceShareWithName retrieves all AWSRAMResourceShare items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRAMResourceShareWithName(name string) (*AWSRAMResourceShare, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSRAMResourceShare:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSRAMResourceShare not found", name)
}

// GetAllAWSRDSDBClusterResources retrieves all AWSRDSDBCluster items from an AWS CloudFormation template
func (t *Template) GetAllAWSRDSDBClusterResources() map[string]*AWSRDSDBCluster {
	results := map[string]*AWSRDSDBCluster{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSRDSDBCluster:
			results[name] = resource
		}
	}
	return results
}

// GetAWSRDSDBClusterWithName retrieves all AWSRDSDBCluster items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRDSDBClusterWithName(name string) (*AWSRDSDBCluster, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSRDSDBCluster:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSRDSDBCluster not found", name)
}

// GetAllAWSRDSDBClusterParameterGroupResources retrieves all AWSRDSDBClusterParameterGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSRDSDBClusterParameterGroupResources() map[string]*AWSRDSDBClusterParameterGroup {
	results := map[string]*AWSRDSDBClusterParameterGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSRDSDBClusterParameterGroup:
			results[name] = resource
		}
	}
	return results
}

// GetAWSRDSDBClusterParameterGroupWithName retrieves all AWSRDSDBClusterParameterGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRDSDBClusterParameterGroupWithName(name string) (*AWSRDSDBClusterParameterGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSRDSDBClusterParameterGroup:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSRDSDBClusterParameterGroup not found", name)
}

// GetAllAWSRDSDBInstanceResources retrieves all AWSRDSDBInstance items from an AWS CloudFormation template
func (t *Template) GetAllAWSRDSDBInstanceResources() map[string]*AWSRDSDBInstance {
	results := map[string]*AWSRDSDBInstance{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSRDSDBInstance:
			results[name] = resource
		}
	}
	return results
}

// GetAWSRDSDBInstanceWithName retrieves all AWSRDSDBInstance items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRDSDBInstanceWithName(name string) (*AWSRDSDBInstance, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSRDSDBInstance:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSRDSDBInstance not found", name)
}

// GetAllAWSRDSDBParameterGroupResources retrieves all AWSRDSDBParameterGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSRDSDBParameterGroupResources() map[string]*AWSRDSDBParameterGroup {
	results := map[string]*AWSRDSDBParameterGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSRDSDBParameterGroup:
			results[name] = resource
		}
	}
	return results
}

// GetAWSRDSDBParameterGroupWithName retrieves all AWSRDSDBParameterGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRDSDBParameterGroupWithName(name string) (*AWSRDSDBParameterGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSRDSDBParameterGroup:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSRDSDBParameterGroup not found", name)
}

// GetAllAWSRDSDBSecurityGroupResources retrieves all AWSRDSDBSecurityGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSRDSDBSecurityGroupResources() map[string]*AWSRDSDBSecurityGroup {
	results := map[string]*AWSRDSDBSecurityGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSRDSDBSecurityGroup:
			results[name] = resource
		}
	}
	return results
}

// GetAWSRDSDBSecurityGroupWithName retrieves all AWSRDSDBSecurityGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRDSDBSecurityGroupWithName(name string) (*AWSRDSDBSecurityGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSRDSDBSecurityGroup:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSRDSDBSecurityGroup not found", name)
}

// GetAllAWSRDSDBSecurityGroupIngressResources retrieves all AWSRDSDBSecurityGroupIngress items from an AWS CloudFormation template
func (t *Template) GetAllAWSRDSDBSecurityGroupIngressResources() map[string]*AWSRDSDBSecurityGroupIngress {
	results := map[string]*AWSRDSDBSecurityGroupIngress{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSRDSDBSecurityGroupIngress:
			results[name] = resource
		}
	}
	return results
}

// GetAWSRDSDBSecurityGroupIngressWithName retrieves all AWSRDSDBSecurityGroupIngress items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRDSDBSecurityGroupIngressWithName(name string) (*AWSRDSDBSecurityGroupIngress, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSRDSDBSecurityGroupIngress:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSRDSDBSecurityGroupIngress not found", name)
}

// GetAllAWSRDSDBSubnetGroupResources retrieves all AWSRDSDBSubnetGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSRDSDBSubnetGroupResources() map[string]*AWSRDSDBSubnetGroup {
	results := map[string]*AWSRDSDBSubnetGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSRDSDBSubnetGroup:
			results[name] = resource
		}
	}
	return results
}

// GetAWSRDSDBSubnetGroupWithName retrieves all AWSRDSDBSubnetGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRDSDBSubnetGroupWithName(name string) (*AWSRDSDBSubnetGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSRDSDBSubnetGroup:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSRDSDBSubnetGroup not found", name)
}

// GetAllAWSRDSEventSubscriptionResources retrieves all AWSRDSEventSubscription items from an AWS CloudFormation template
func (t *Template) GetAllAWSRDSEventSubscriptionResources() map[string]*AWSRDSEventSubscription {
	results := map[string]*AWSRDSEventSubscription{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSRDSEventSubscription:
			results[name] = resource
		}
	}
	return results
}

// GetAWSRDSEventSubscriptionWithName retrieves all AWSRDSEventSubscription items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRDSEventSubscriptionWithName(name string) (*AWSRDSEventSubscription, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSRDSEventSubscription:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSRDSEventSubscription not found", name)
}

// GetAllAWSRDSOptionGroupResources retrieves all AWSRDSOptionGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSRDSOptionGroupResources() map[string]*AWSRDSOptionGroup {
	results := map[string]*AWSRDSOptionGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSRDSOptionGroup:
			results[name] = resource
		}
	}
	return results
}

// GetAWSRDSOptionGroupWithName retrieves all AWSRDSOptionGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRDSOptionGroupWithName(name string) (*AWSRDSOptionGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSRDSOptionGroup:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSRDSOptionGroup not found", name)
}

// GetAllAWSRedshiftClusterResources retrieves all AWSRedshiftCluster items from an AWS CloudFormation template
func (t *Template) GetAllAWSRedshiftClusterResources() map[string]*AWSRedshiftCluster {
	results := map[string]*AWSRedshiftCluster{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSRedshiftCluster:
			results[name] = resource
		}
	}
	return results
}

// GetAWSRedshiftClusterWithName retrieves all AWSRedshiftCluster items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRedshiftClusterWithName(name string) (*AWSRedshiftCluster, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSRedshiftCluster:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSRedshiftCluster not found", name)
}

// GetAllAWSRedshiftClusterParameterGroupResources retrieves all AWSRedshiftClusterParameterGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSRedshiftClusterParameterGroupResources() map[string]*AWSRedshiftClusterParameterGroup {
	results := map[string]*AWSRedshiftClusterParameterGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSRedshiftClusterParameterGroup:
			results[name] = resource
		}
	}
	return results
}

// GetAWSRedshiftClusterParameterGroupWithName retrieves all AWSRedshiftClusterParameterGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRedshiftClusterParameterGroupWithName(name string) (*AWSRedshiftClusterParameterGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSRedshiftClusterParameterGroup:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSRedshiftClusterParameterGroup not found", name)
}

// GetAllAWSRedshiftClusterSecurityGroupResources retrieves all AWSRedshiftClusterSecurityGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSRedshiftClusterSecurityGroupResources() map[string]*AWSRedshiftClusterSecurityGroup {
	results := map[string]*AWSRedshiftClusterSecurityGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSRedshiftClusterSecurityGroup:
			results[name] = resource
		}
	}
	return results
}

// GetAWSRedshiftClusterSecurityGroupWithName retrieves all AWSRedshiftClusterSecurityGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRedshiftClusterSecurityGroupWithName(name string) (*AWSRedshiftClusterSecurityGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSRedshiftClusterSecurityGroup:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSRedshiftClusterSecurityGroup not found", name)
}

// GetAllAWSRedshiftClusterSecurityGroupIngressResources retrieves all AWSRedshiftClusterSecurityGroupIngress items from an AWS CloudFormation template
func (t *Template) GetAllAWSRedshiftClusterSecurityGroupIngressResources() map[string]*AWSRedshiftClusterSecurityGroupIngress {
	results := map[string]*AWSRedshiftClusterSecurityGroupIngress{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSRedshiftClusterSecurityGroupIngress:
			results[name] = resource
		}
	}
	return results
}

// GetAWSRedshiftClusterSecurityGroupIngressWithName retrieves all AWSRedshiftClusterSecurityGroupIngress items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRedshiftClusterSecurityGroupIngressWithName(name string) (*AWSRedshiftClusterSecurityGroupIngress, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSRedshiftClusterSecurityGroupIngress:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSRedshiftClusterSecurityGroupIngress not found", name)
}

// GetAllAWSRedshiftClusterSubnetGroupResources retrieves all AWSRedshiftClusterSubnetGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSRedshiftClusterSubnetGroupResources() map[string]*AWSRedshiftClusterSubnetGroup {
	results := map[string]*AWSRedshiftClusterSubnetGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSRedshiftClusterSubnetGroup:
			results[name] = resource
		}
	}
	return results
}

// GetAWSRedshiftClusterSubnetGroupWithName retrieves all AWSRedshiftClusterSubnetGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRedshiftClusterSubnetGroupWithName(name string) (*AWSRedshiftClusterSubnetGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSRedshiftClusterSubnetGroup:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSRedshiftClusterSubnetGroup not found", name)
}

// GetAllAWSRoboMakerFleetResources retrieves all AWSRoboMakerFleet items from an AWS CloudFormation template
func (t *Template) GetAllAWSRoboMakerFleetResources() map[string]*AWSRoboMakerFleet {
	results := map[string]*AWSRoboMakerFleet{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSRoboMakerFleet:
			results[name] = resource
		}
	}
	return results
}

// GetAWSRoboMakerFleetWithName retrieves all AWSRoboMakerFleet items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRoboMakerFleetWithName(name string) (*AWSRoboMakerFleet, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSRoboMakerFleet:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSRoboMakerFleet not found", name)
}

// GetAllAWSRoboMakerRobotResources retrieves all AWSRoboMakerRobot items from an AWS CloudFormation template
func (t *Template) GetAllAWSRoboMakerRobotResources() map[string]*AWSRoboMakerRobot {
	results := map[string]*AWSRoboMakerRobot{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSRoboMakerRobot:
			results[name] = resource
		}
	}
	return results
}

// GetAWSRoboMakerRobotWithName retrieves all AWSRoboMakerRobot items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRoboMakerRobotWithName(name string) (*AWSRoboMakerRobot, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSRoboMakerRobot:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSRoboMakerRobot not found", name)
}

// GetAllAWSRoboMakerRobotApplicationResources retrieves all AWSRoboMakerRobotApplication items from an AWS CloudFormation template
func (t *Template) GetAllAWSRoboMakerRobotApplicationResources() map[string]*AWSRoboMakerRobotApplication {
	results := map[string]*AWSRoboMakerRobotApplication{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSRoboMakerRobotApplication:
			results[name] = resource
		}
	}
	return results
}

// GetAWSRoboMakerRobotApplicationWithName retrieves all AWSRoboMakerRobotApplication items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRoboMakerRobotApplicationWithName(name string) (*AWSRoboMakerRobotApplication, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSRoboMakerRobotApplication:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSRoboMakerRobotApplication not found", name)
}

// GetAllAWSRoboMakerRobotApplicationVersionResources retrieves all AWSRoboMakerRobotApplicationVersion items from an AWS CloudFormation template
func (t *Template) GetAllAWSRoboMakerRobotApplicationVersionResources() map[string]*AWSRoboMakerRobotApplicationVersion {
	results := map[string]*AWSRoboMakerRobotApplicationVersion{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSRoboMakerRobotApplicationVersion:
			results[name] = resource
		}
	}
	return results
}

// GetAWSRoboMakerRobotApplicationVersionWithName retrieves all AWSRoboMakerRobotApplicationVersion items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRoboMakerRobotApplicationVersionWithName(name string) (*AWSRoboMakerRobotApplicationVersion, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSRoboMakerRobotApplicationVersion:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSRoboMakerRobotApplicationVersion not found", name)
}

// GetAllAWSRoboMakerSimulationApplicationResources retrieves all AWSRoboMakerSimulationApplication items from an AWS CloudFormation template
func (t *Template) GetAllAWSRoboMakerSimulationApplicationResources() map[string]*AWSRoboMakerSimulationApplication {
	results := map[string]*AWSRoboMakerSimulationApplication{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSRoboMakerSimulationApplication:
			results[name] = resource
		}
	}
	return results
}

// GetAWSRoboMakerSimulationApplicationWithName retrieves all AWSRoboMakerSimulationApplication items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRoboMakerSimulationApplicationWithName(name string) (*AWSRoboMakerSimulationApplication, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSRoboMakerSimulationApplication:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSRoboMakerSimulationApplication not found", name)
}

// GetAllAWSRoboMakerSimulationApplicationVersionResources retrieves all AWSRoboMakerSimulationApplicationVersion items from an AWS CloudFormation template
func (t *Template) GetAllAWSRoboMakerSimulationApplicationVersionResources() map[string]*AWSRoboMakerSimulationApplicationVersion {
	results := map[string]*AWSRoboMakerSimulationApplicationVersion{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSRoboMakerSimulationApplicationVersion:
			results[name] = resource
		}
	}
	return results
}

// GetAWSRoboMakerSimulationApplicationVersionWithName retrieves all AWSRoboMakerSimulationApplicationVersion items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRoboMakerSimulationApplicationVersionWithName(name string) (*AWSRoboMakerSimulationApplicationVersion, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSRoboMakerSimulationApplicationVersion:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSRoboMakerSimulationApplicationVersion not found", name)
}

// GetAllAWSRoute53HealthCheckResources retrieves all AWSRoute53HealthCheck items from an AWS CloudFormation template
func (t *Template) GetAllAWSRoute53HealthCheckResources() map[string]*AWSRoute53HealthCheck {
	results := map[string]*AWSRoute53HealthCheck{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSRoute53HealthCheck:
			results[name] = resource
		}
	}
	return results
}

// GetAWSRoute53HealthCheckWithName retrieves all AWSRoute53HealthCheck items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRoute53HealthCheckWithName(name string) (*AWSRoute53HealthCheck, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSRoute53HealthCheck:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSRoute53HealthCheck not found", name)
}

// GetAllAWSRoute53HostedZoneResources retrieves all AWSRoute53HostedZone items from an AWS CloudFormation template
func (t *Template) GetAllAWSRoute53HostedZoneResources() map[string]*AWSRoute53HostedZone {
	results := map[string]*AWSRoute53HostedZone{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSRoute53HostedZone:
			results[name] = resource
		}
	}
	return results
}

// GetAWSRoute53HostedZoneWithName retrieves all AWSRoute53HostedZone items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRoute53HostedZoneWithName(name string) (*AWSRoute53HostedZone, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSRoute53HostedZone:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSRoute53HostedZone not found", name)
}

// GetAllAWSRoute53RecordSetResources retrieves all AWSRoute53RecordSet items from an AWS CloudFormation template
func (t *Template) GetAllAWSRoute53RecordSetResources() map[string]*AWSRoute53RecordSet {
	results := map[string]*AWSRoute53RecordSet{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSRoute53RecordSet:
			results[name] = resource
		}
	}
	return results
}

// GetAWSRoute53RecordSetWithName retrieves all AWSRoute53RecordSet items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRoute53RecordSetWithName(name string) (*AWSRoute53RecordSet, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSRoute53RecordSet:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSRoute53RecordSet not found", name)
}

// GetAllAWSRoute53RecordSetGroupResources retrieves all AWSRoute53RecordSetGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSRoute53RecordSetGroupResources() map[string]*AWSRoute53RecordSetGroup {
	results := map[string]*AWSRoute53RecordSetGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSRoute53RecordSetGroup:
			results[name] = resource
		}
	}
	return results
}

// GetAWSRoute53RecordSetGroupWithName retrieves all AWSRoute53RecordSetGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRoute53RecordSetGroupWithName(name string) (*AWSRoute53RecordSetGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSRoute53RecordSetGroup:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSRoute53RecordSetGroup not found", name)
}

// GetAllAWSRoute53ResolverResolverEndpointResources retrieves all AWSRoute53ResolverResolverEndpoint items from an AWS CloudFormation template
func (t *Template) GetAllAWSRoute53ResolverResolverEndpointResources() map[string]*AWSRoute53ResolverResolverEndpoint {
	results := map[string]*AWSRoute53ResolverResolverEndpoint{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSRoute53ResolverResolverEndpoint:
			results[name] = resource
		}
	}
	return results
}

// GetAWSRoute53ResolverResolverEndpointWithName retrieves all AWSRoute53ResolverResolverEndpoint items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRoute53ResolverResolverEndpointWithName(name string) (*AWSRoute53ResolverResolverEndpoint, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSRoute53ResolverResolverEndpoint:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSRoute53ResolverResolverEndpoint not found", name)
}

// GetAllAWSRoute53ResolverResolverRuleResources retrieves all AWSRoute53ResolverResolverRule items from an AWS CloudFormation template
func (t *Template) GetAllAWSRoute53ResolverResolverRuleResources() map[string]*AWSRoute53ResolverResolverRule {
	results := map[string]*AWSRoute53ResolverResolverRule{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSRoute53ResolverResolverRule:
			results[name] = resource
		}
	}
	return results
}

// GetAWSRoute53ResolverResolverRuleWithName retrieves all AWSRoute53ResolverResolverRule items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRoute53ResolverResolverRuleWithName(name string) (*AWSRoute53ResolverResolverRule, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSRoute53ResolverResolverRule:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSRoute53ResolverResolverRule not found", name)
}

// GetAllAWSRoute53ResolverResolverRuleAssociationResources retrieves all AWSRoute53ResolverResolverRuleAssociation items from an AWS CloudFormation template
func (t *Template) GetAllAWSRoute53ResolverResolverRuleAssociationResources() map[string]*AWSRoute53ResolverResolverRuleAssociation {
	results := map[string]*AWSRoute53ResolverResolverRuleAssociation{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSRoute53ResolverResolverRuleAssociation:
			results[name] = resource
		}
	}
	return results
}

// GetAWSRoute53ResolverResolverRuleAssociationWithName retrieves all AWSRoute53ResolverResolverRuleAssociation items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRoute53ResolverResolverRuleAssociationWithName(name string) (*AWSRoute53ResolverResolverRuleAssociation, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSRoute53ResolverResolverRuleAssociation:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSRoute53ResolverResolverRuleAssociation not found", name)
}

// GetAllAWSS3BucketResources retrieves all AWSS3Bucket items from an AWS CloudFormation template
func (t *Template) GetAllAWSS3BucketResources() map[string]*AWSS3Bucket {
	results := map[string]*AWSS3Bucket{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSS3Bucket:
			results[name] = resource
		}
	}
	return results
}

// GetAWSS3BucketWithName retrieves all AWSS3Bucket items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSS3BucketWithName(name string) (*AWSS3Bucket, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSS3Bucket:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSS3Bucket not found", name)
}

// GetAllAWSS3BucketPolicyResources retrieves all AWSS3BucketPolicy items from an AWS CloudFormation template
func (t *Template) GetAllAWSS3BucketPolicyResources() map[string]*AWSS3BucketPolicy {
	results := map[string]*AWSS3BucketPolicy{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSS3BucketPolicy:
			results[name] = resource
		}
	}
	return results
}

// GetAWSS3BucketPolicyWithName retrieves all AWSS3BucketPolicy items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSS3BucketPolicyWithName(name string) (*AWSS3BucketPolicy, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSS3BucketPolicy:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSS3BucketPolicy not found", name)
}

// GetAllAWSSDBDomainResources retrieves all AWSSDBDomain items from an AWS CloudFormation template
func (t *Template) GetAllAWSSDBDomainResources() map[string]*AWSSDBDomain {
	results := map[string]*AWSSDBDomain{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSSDBDomain:
			results[name] = resource
		}
	}
	return results
}

// GetAWSSDBDomainWithName retrieves all AWSSDBDomain items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSDBDomainWithName(name string) (*AWSSDBDomain, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSSDBDomain:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSSDBDomain not found", name)
}

// GetAllAWSSESConfigurationSetResources retrieves all AWSSESConfigurationSet items from an AWS CloudFormation template
func (t *Template) GetAllAWSSESConfigurationSetResources() map[string]*AWSSESConfigurationSet {
	results := map[string]*AWSSESConfigurationSet{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSSESConfigurationSet:
			results[name] = resource
		}
	}
	return results
}

// GetAWSSESConfigurationSetWithName retrieves all AWSSESConfigurationSet items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSESConfigurationSetWithName(name string) (*AWSSESConfigurationSet, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSSESConfigurationSet:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSSESConfigurationSet not found", name)
}

// GetAllAWSSESConfigurationSetEventDestinationResources retrieves all AWSSESConfigurationSetEventDestination items from an AWS CloudFormation template
func (t *Template) GetAllAWSSESConfigurationSetEventDestinationResources() map[string]*AWSSESConfigurationSetEventDestination {
	results := map[string]*AWSSESConfigurationSetEventDestination{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSSESConfigurationSetEventDestination:
			results[name] = resource
		}
	}
	return results
}

// GetAWSSESConfigurationSetEventDestinationWithName retrieves all AWSSESConfigurationSetEventDestination items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSESConfigurationSetEventDestinationWithName(name string) (*AWSSESConfigurationSetEventDestination, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSSESConfigurationSetEventDestination:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSSESConfigurationSetEventDestination not found", name)
}

// GetAllAWSSESReceiptFilterResources retrieves all AWSSESReceiptFilter items from an AWS CloudFormation template
func (t *Template) GetAllAWSSESReceiptFilterResources() map[string]*AWSSESReceiptFilter {
	results := map[string]*AWSSESReceiptFilter{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSSESReceiptFilter:
			results[name] = resource
		}
	}
	return results
}

// GetAWSSESReceiptFilterWithName retrieves all AWSSESReceiptFilter items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSESReceiptFilterWithName(name string) (*AWSSESReceiptFilter, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSSESReceiptFilter:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSSESReceiptFilter not found", name)
}

// GetAllAWSSESReceiptRuleResources retrieves all AWSSESReceiptRule items from an AWS CloudFormation template
func (t *Template) GetAllAWSSESReceiptRuleResources() map[string]*AWSSESReceiptRule {
	results := map[string]*AWSSESReceiptRule{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSSESReceiptRule:
			results[name] = resource
		}
	}
	return results
}

// GetAWSSESReceiptRuleWithName retrieves all AWSSESReceiptRule items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSESReceiptRuleWithName(name string) (*AWSSESReceiptRule, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSSESReceiptRule:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSSESReceiptRule not found", name)
}

// GetAllAWSSESReceiptRuleSetResources retrieves all AWSSESReceiptRuleSet items from an AWS CloudFormation template
func (t *Template) GetAllAWSSESReceiptRuleSetResources() map[string]*AWSSESReceiptRuleSet {
	results := map[string]*AWSSESReceiptRuleSet{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSSESReceiptRuleSet:
			results[name] = resource
		}
	}
	return results
}

// GetAWSSESReceiptRuleSetWithName retrieves all AWSSESReceiptRuleSet items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSESReceiptRuleSetWithName(name string) (*AWSSESReceiptRuleSet, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSSESReceiptRuleSet:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSSESReceiptRuleSet not found", name)
}

// GetAllAWSSESTemplateResources retrieves all AWSSESTemplate items from an AWS CloudFormation template
func (t *Template) GetAllAWSSESTemplateResources() map[string]*AWSSESTemplate {
	results := map[string]*AWSSESTemplate{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSSESTemplate:
			results[name] = resource
		}
	}
	return results
}

// GetAWSSESTemplateWithName retrieves all AWSSESTemplate items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSESTemplateWithName(name string) (*AWSSESTemplate, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSSESTemplate:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSSESTemplate not found", name)
}

// GetAllAWSSNSSubscriptionResources retrieves all AWSSNSSubscription items from an AWS CloudFormation template
func (t *Template) GetAllAWSSNSSubscriptionResources() map[string]*AWSSNSSubscription {
	results := map[string]*AWSSNSSubscription{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSSNSSubscription:
			results[name] = resource
		}
	}
	return results
}

// GetAWSSNSSubscriptionWithName retrieves all AWSSNSSubscription items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSNSSubscriptionWithName(name string) (*AWSSNSSubscription, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSSNSSubscription:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSSNSSubscription not found", name)
}

// GetAllAWSSNSTopicResources retrieves all AWSSNSTopic items from an AWS CloudFormation template
func (t *Template) GetAllAWSSNSTopicResources() map[string]*AWSSNSTopic {
	results := map[string]*AWSSNSTopic{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSSNSTopic:
			results[name] = resource
		}
	}
	return results
}

// GetAWSSNSTopicWithName retrieves all AWSSNSTopic items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSNSTopicWithName(name string) (*AWSSNSTopic, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSSNSTopic:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSSNSTopic not found", name)
}

// GetAllAWSSNSTopicPolicyResources retrieves all AWSSNSTopicPolicy items from an AWS CloudFormation template
func (t *Template) GetAllAWSSNSTopicPolicyResources() map[string]*AWSSNSTopicPolicy {
	results := map[string]*AWSSNSTopicPolicy{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSSNSTopicPolicy:
			results[name] = resource
		}
	}
	return results
}

// GetAWSSNSTopicPolicyWithName retrieves all AWSSNSTopicPolicy items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSNSTopicPolicyWithName(name string) (*AWSSNSTopicPolicy, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSSNSTopicPolicy:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSSNSTopicPolicy not found", name)
}

// GetAllAWSSQSQueueResources retrieves all AWSSQSQueue items from an AWS CloudFormation template
func (t *Template) GetAllAWSSQSQueueResources() map[string]*AWSSQSQueue {
	results := map[string]*AWSSQSQueue{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSSQSQueue:
			results[name] = resource
		}
	}
	return results
}

// GetAWSSQSQueueWithName retrieves all AWSSQSQueue items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSQSQueueWithName(name string) (*AWSSQSQueue, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSSQSQueue:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSSQSQueue not found", name)
}

// GetAllAWSSQSQueuePolicyResources retrieves all AWSSQSQueuePolicy items from an AWS CloudFormation template
func (t *Template) GetAllAWSSQSQueuePolicyResources() map[string]*AWSSQSQueuePolicy {
	results := map[string]*AWSSQSQueuePolicy{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSSQSQueuePolicy:
			results[name] = resource
		}
	}
	return results
}

// GetAWSSQSQueuePolicyWithName retrieves all AWSSQSQueuePolicy items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSQSQueuePolicyWithName(name string) (*AWSSQSQueuePolicy, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSSQSQueuePolicy:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSSQSQueuePolicy not found", name)
}

// GetAllAWSSSMAssociationResources retrieves all AWSSSMAssociation items from an AWS CloudFormation template
func (t *Template) GetAllAWSSSMAssociationResources() map[string]*AWSSSMAssociation {
	results := map[string]*AWSSSMAssociation{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSSSMAssociation:
			results[name] = resource
		}
	}
	return results
}

// GetAWSSSMAssociationWithName retrieves all AWSSSMAssociation items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSSMAssociationWithName(name string) (*AWSSSMAssociation, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSSSMAssociation:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSSSMAssociation not found", name)
}

// GetAllAWSSSMDocumentResources retrieves all AWSSSMDocument items from an AWS CloudFormation template
func (t *Template) GetAllAWSSSMDocumentResources() map[string]*AWSSSMDocument {
	results := map[string]*AWSSSMDocument{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSSSMDocument:
			results[name] = resource
		}
	}
	return results
}

// GetAWSSSMDocumentWithName retrieves all AWSSSMDocument items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSSMDocumentWithName(name string) (*AWSSSMDocument, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSSSMDocument:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSSSMDocument not found", name)
}

// GetAllAWSSSMMaintenanceWindowResources retrieves all AWSSSMMaintenanceWindow items from an AWS CloudFormation template
func (t *Template) GetAllAWSSSMMaintenanceWindowResources() map[string]*AWSSSMMaintenanceWindow {
	results := map[string]*AWSSSMMaintenanceWindow{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSSSMMaintenanceWindow:
			results[name] = resource
		}
	}
	return results
}

// GetAWSSSMMaintenanceWindowWithName retrieves all AWSSSMMaintenanceWindow items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSSMMaintenanceWindowWithName(name string) (*AWSSSMMaintenanceWindow, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSSSMMaintenanceWindow:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSSSMMaintenanceWindow not found", name)
}

// GetAllAWSSSMMaintenanceWindowTaskResources retrieves all AWSSSMMaintenanceWindowTask items from an AWS CloudFormation template
func (t *Template) GetAllAWSSSMMaintenanceWindowTaskResources() map[string]*AWSSSMMaintenanceWindowTask {
	results := map[string]*AWSSSMMaintenanceWindowTask{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSSSMMaintenanceWindowTask:
			results[name] = resource
		}
	}
	return results
}

// GetAWSSSMMaintenanceWindowTaskWithName retrieves all AWSSSMMaintenanceWindowTask items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSSMMaintenanceWindowTaskWithName(name string) (*AWSSSMMaintenanceWindowTask, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSSSMMaintenanceWindowTask:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSSSMMaintenanceWindowTask not found", name)
}

// GetAllAWSSSMParameterResources retrieves all AWSSSMParameter items from an AWS CloudFormation template
func (t *Template) GetAllAWSSSMParameterResources() map[string]*AWSSSMParameter {
	results := map[string]*AWSSSMParameter{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSSSMParameter:
			results[name] = resource
		}
	}
	return results
}

// GetAWSSSMParameterWithName retrieves all AWSSSMParameter items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSSMParameterWithName(name string) (*AWSSSMParameter, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSSSMParameter:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSSSMParameter not found", name)
}

// GetAllAWSSSMPatchBaselineResources retrieves all AWSSSMPatchBaseline items from an AWS CloudFormation template
func (t *Template) GetAllAWSSSMPatchBaselineResources() map[string]*AWSSSMPatchBaseline {
	results := map[string]*AWSSSMPatchBaseline{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSSSMPatchBaseline:
			results[name] = resource
		}
	}
	return results
}

// GetAWSSSMPatchBaselineWithName retrieves all AWSSSMPatchBaseline items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSSMPatchBaselineWithName(name string) (*AWSSSMPatchBaseline, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSSSMPatchBaseline:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSSSMPatchBaseline not found", name)
}

// GetAllAWSSSMResourceDataSyncResources retrieves all AWSSSMResourceDataSync items from an AWS CloudFormation template
func (t *Template) GetAllAWSSSMResourceDataSyncResources() map[string]*AWSSSMResourceDataSync {
	results := map[string]*AWSSSMResourceDataSync{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSSSMResourceDataSync:
			results[name] = resource
		}
	}
	return results
}

// GetAWSSSMResourceDataSyncWithName retrieves all AWSSSMResourceDataSync items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSSMResourceDataSyncWithName(name string) (*AWSSSMResourceDataSync, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSSSMResourceDataSync:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSSSMResourceDataSync not found", name)
}

// GetAllAWSSageMakerEndpointResources retrieves all AWSSageMakerEndpoint items from an AWS CloudFormation template
func (t *Template) GetAllAWSSageMakerEndpointResources() map[string]*AWSSageMakerEndpoint {
	results := map[string]*AWSSageMakerEndpoint{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSSageMakerEndpoint:
			results[name] = resource
		}
	}
	return results
}

// GetAWSSageMakerEndpointWithName retrieves all AWSSageMakerEndpoint items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSageMakerEndpointWithName(name string) (*AWSSageMakerEndpoint, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSSageMakerEndpoint:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSSageMakerEndpoint not found", name)
}

// GetAllAWSSageMakerEndpointConfigResources retrieves all AWSSageMakerEndpointConfig items from an AWS CloudFormation template
func (t *Template) GetAllAWSSageMakerEndpointConfigResources() map[string]*AWSSageMakerEndpointConfig {
	results := map[string]*AWSSageMakerEndpointConfig{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSSageMakerEndpointConfig:
			results[name] = resource
		}
	}
	return results
}

// GetAWSSageMakerEndpointConfigWithName retrieves all AWSSageMakerEndpointConfig items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSageMakerEndpointConfigWithName(name string) (*AWSSageMakerEndpointConfig, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSSageMakerEndpointConfig:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSSageMakerEndpointConfig not found", name)
}

// GetAllAWSSageMakerModelResources retrieves all AWSSageMakerModel items from an AWS CloudFormation template
func (t *Template) GetAllAWSSageMakerModelResources() map[string]*AWSSageMakerModel {
	results := map[string]*AWSSageMakerModel{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSSageMakerModel:
			results[name] = resource
		}
	}
	return results
}

// GetAWSSageMakerModelWithName retrieves all AWSSageMakerModel items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSageMakerModelWithName(name string) (*AWSSageMakerModel, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSSageMakerModel:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSSageMakerModel not found", name)
}

// GetAllAWSSageMakerNotebookInstanceResources retrieves all AWSSageMakerNotebookInstance items from an AWS CloudFormation template
func (t *Template) GetAllAWSSageMakerNotebookInstanceResources() map[string]*AWSSageMakerNotebookInstance {
	results := map[string]*AWSSageMakerNotebookInstance{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSSageMakerNotebookInstance:
			results[name] = resource
		}
	}
	return results
}

// GetAWSSageMakerNotebookInstanceWithName retrieves all AWSSageMakerNotebookInstance items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSageMakerNotebookInstanceWithName(name string) (*AWSSageMakerNotebookInstance, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSSageMakerNotebookInstance:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSSageMakerNotebookInstance not found", name)
}

// GetAllAWSSageMakerNotebookInstanceLifecycleConfigResources retrieves all AWSSageMakerNotebookInstanceLifecycleConfig items from an AWS CloudFormation template
func (t *Template) GetAllAWSSageMakerNotebookInstanceLifecycleConfigResources() map[string]*AWSSageMakerNotebookInstanceLifecycleConfig {
	results := map[string]*AWSSageMakerNotebookInstanceLifecycleConfig{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSSageMakerNotebookInstanceLifecycleConfig:
			results[name] = resource
		}
	}
	return results
}

// GetAWSSageMakerNotebookInstanceLifecycleConfigWithName retrieves all AWSSageMakerNotebookInstanceLifecycleConfig items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSageMakerNotebookInstanceLifecycleConfigWithName(name string) (*AWSSageMakerNotebookInstanceLifecycleConfig, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSSageMakerNotebookInstanceLifecycleConfig:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSSageMakerNotebookInstanceLifecycleConfig not found", name)
}

// GetAllAWSSecretsManagerResourcePolicyResources retrieves all AWSSecretsManagerResourcePolicy items from an AWS CloudFormation template
func (t *Template) GetAllAWSSecretsManagerResourcePolicyResources() map[string]*AWSSecretsManagerResourcePolicy {
	results := map[string]*AWSSecretsManagerResourcePolicy{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSSecretsManagerResourcePolicy:
			results[name] = resource
		}
	}
	return results
}

// GetAWSSecretsManagerResourcePolicyWithName retrieves all AWSSecretsManagerResourcePolicy items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSecretsManagerResourcePolicyWithName(name string) (*AWSSecretsManagerResourcePolicy, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSSecretsManagerResourcePolicy:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSSecretsManagerResourcePolicy not found", name)
}

// GetAllAWSSecretsManagerRotationScheduleResources retrieves all AWSSecretsManagerRotationSchedule items from an AWS CloudFormation template
func (t *Template) GetAllAWSSecretsManagerRotationScheduleResources() map[string]*AWSSecretsManagerRotationSchedule {
	results := map[string]*AWSSecretsManagerRotationSchedule{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSSecretsManagerRotationSchedule:
			results[name] = resource
		}
	}
	return results
}

// GetAWSSecretsManagerRotationScheduleWithName retrieves all AWSSecretsManagerRotationSchedule items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSecretsManagerRotationScheduleWithName(name string) (*AWSSecretsManagerRotationSchedule, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSSecretsManagerRotationSchedule:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSSecretsManagerRotationSchedule not found", name)
}

// GetAllAWSSecretsManagerSecretResources retrieves all AWSSecretsManagerSecret items from an AWS CloudFormation template
func (t *Template) GetAllAWSSecretsManagerSecretResources() map[string]*AWSSecretsManagerSecret {
	results := map[string]*AWSSecretsManagerSecret{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSSecretsManagerSecret:
			results[name] = resource
		}
	}
	return results
}

// GetAWSSecretsManagerSecretWithName retrieves all AWSSecretsManagerSecret items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSecretsManagerSecretWithName(name string) (*AWSSecretsManagerSecret, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSSecretsManagerSecret:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSSecretsManagerSecret not found", name)
}

// GetAllAWSSecretsManagerSecretTargetAttachmentResources retrieves all AWSSecretsManagerSecretTargetAttachment items from an AWS CloudFormation template
func (t *Template) GetAllAWSSecretsManagerSecretTargetAttachmentResources() map[string]*AWSSecretsManagerSecretTargetAttachment {
	results := map[string]*AWSSecretsManagerSecretTargetAttachment{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSSecretsManagerSecretTargetAttachment:
			results[name] = resource
		}
	}
	return results
}

// GetAWSSecretsManagerSecretTargetAttachmentWithName retrieves all AWSSecretsManagerSecretTargetAttachment items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSecretsManagerSecretTargetAttachmentWithName(name string) (*AWSSecretsManagerSecretTargetAttachment, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSSecretsManagerSecretTargetAttachment:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSSecretsManagerSecretTargetAttachment not found", name)
}

// GetAllAWSServerlessApiResources retrieves all AWSServerlessApi items from an AWS CloudFormation template
func (t *Template) GetAllAWSServerlessApiResources() map[string]*AWSServerlessApi {
	results := map[string]*AWSServerlessApi{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSServerlessApi:
			results[name] = resource
		}
	}
	return results
}

// GetAWSServerlessApiWithName retrieves all AWSServerlessApi items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSServerlessApiWithName(name string) (*AWSServerlessApi, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSServerlessApi:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSServerlessApi not found", name)
}

// GetAllAWSServerlessApplicationResources retrieves all AWSServerlessApplication items from an AWS CloudFormation template
func (t *Template) GetAllAWSServerlessApplicationResources() map[string]*AWSServerlessApplication {
	results := map[string]*AWSServerlessApplication{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSServerlessApplication:
			results[name] = resource
		}
	}
	return results
}

// GetAWSServerlessApplicationWithName retrieves all AWSServerlessApplication items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSServerlessApplicationWithName(name string) (*AWSServerlessApplication, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSServerlessApplication:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSServerlessApplication not found", name)
}

// GetAllAWSServerlessFunctionResources retrieves all AWSServerlessFunction items from an AWS CloudFormation template
func (t *Template) GetAllAWSServerlessFunctionResources() map[string]*AWSServerlessFunction {
	results := map[string]*AWSServerlessFunction{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSServerlessFunction:
			results[name] = resource
		}
	}
	return results
}

// GetAWSServerlessFunctionWithName retrieves all AWSServerlessFunction items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSServerlessFunctionWithName(name string) (*AWSServerlessFunction, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSServerlessFunction:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSServerlessFunction not found", name)
}

// GetAllAWSServerlessLayerVersionResources retrieves all AWSServerlessLayerVersion items from an AWS CloudFormation template
func (t *Template) GetAllAWSServerlessLayerVersionResources() map[string]*AWSServerlessLayerVersion {
	results := map[string]*AWSServerlessLayerVersion{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSServerlessLayerVersion:
			results[name] = resource
		}
	}
	return results
}

// GetAWSServerlessLayerVersionWithName retrieves all AWSServerlessLayerVersion items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSServerlessLayerVersionWithName(name string) (*AWSServerlessLayerVersion, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSServerlessLayerVersion:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSServerlessLayerVersion not found", name)
}

// GetAllAWSServerlessSimpleTableResources retrieves all AWSServerlessSimpleTable items from an AWS CloudFormation template
func (t *Template) GetAllAWSServerlessSimpleTableResources() map[string]*AWSServerlessSimpleTable {
	results := map[string]*AWSServerlessSimpleTable{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSServerlessSimpleTable:
			results[name] = resource
		}
	}
	return results
}

// GetAWSServerlessSimpleTableWithName retrieves all AWSServerlessSimpleTable items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSServerlessSimpleTableWithName(name string) (*AWSServerlessSimpleTable, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSServerlessSimpleTable:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSServerlessSimpleTable not found", name)
}

// GetAllAWSServiceCatalogAcceptedPortfolioShareResources retrieves all AWSServiceCatalogAcceptedPortfolioShare items from an AWS CloudFormation template
func (t *Template) GetAllAWSServiceCatalogAcceptedPortfolioShareResources() map[string]*AWSServiceCatalogAcceptedPortfolioShare {
	results := map[string]*AWSServiceCatalogAcceptedPortfolioShare{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSServiceCatalogAcceptedPortfolioShare:
			results[name] = resource
		}
	}
	return results
}

// GetAWSServiceCatalogAcceptedPortfolioShareWithName retrieves all AWSServiceCatalogAcceptedPortfolioShare items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSServiceCatalogAcceptedPortfolioShareWithName(name string) (*AWSServiceCatalogAcceptedPortfolioShare, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSServiceCatalogAcceptedPortfolioShare:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSServiceCatalogAcceptedPortfolioShare not found", name)
}

// GetAllAWSServiceCatalogCloudFormationProductResources retrieves all AWSServiceCatalogCloudFormationProduct items from an AWS CloudFormation template
func (t *Template) GetAllAWSServiceCatalogCloudFormationProductResources() map[string]*AWSServiceCatalogCloudFormationProduct {
	results := map[string]*AWSServiceCatalogCloudFormationProduct{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSServiceCatalogCloudFormationProduct:
			results[name] = resource
		}
	}
	return results
}

// GetAWSServiceCatalogCloudFormationProductWithName retrieves all AWSServiceCatalogCloudFormationProduct items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSServiceCatalogCloudFormationProductWithName(name string) (*AWSServiceCatalogCloudFormationProduct, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSServiceCatalogCloudFormationProduct:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSServiceCatalogCloudFormationProduct not found", name)
}

// GetAllAWSServiceCatalogCloudFormationProvisionedProductResources retrieves all AWSServiceCatalogCloudFormationProvisionedProduct items from an AWS CloudFormation template
func (t *Template) GetAllAWSServiceCatalogCloudFormationProvisionedProductResources() map[string]*AWSServiceCatalogCloudFormationProvisionedProduct {
	results := map[string]*AWSServiceCatalogCloudFormationProvisionedProduct{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSServiceCatalogCloudFormationProvisionedProduct:
			results[name] = resource
		}
	}
	return results
}

// GetAWSServiceCatalogCloudFormationProvisionedProductWithName retrieves all AWSServiceCatalogCloudFormationProvisionedProduct items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSServiceCatalogCloudFormationProvisionedProductWithName(name string) (*AWSServiceCatalogCloudFormationProvisionedProduct, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSServiceCatalogCloudFormationProvisionedProduct:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSServiceCatalogCloudFormationProvisionedProduct not found", name)
}

// GetAllAWSServiceCatalogLaunchNotificationConstraintResources retrieves all AWSServiceCatalogLaunchNotificationConstraint items from an AWS CloudFormation template
func (t *Template) GetAllAWSServiceCatalogLaunchNotificationConstraintResources() map[string]*AWSServiceCatalogLaunchNotificationConstraint {
	results := map[string]*AWSServiceCatalogLaunchNotificationConstraint{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSServiceCatalogLaunchNotificationConstraint:
			results[name] = resource
		}
	}
	return results
}

// GetAWSServiceCatalogLaunchNotificationConstraintWithName retrieves all AWSServiceCatalogLaunchNotificationConstraint items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSServiceCatalogLaunchNotificationConstraintWithName(name string) (*AWSServiceCatalogLaunchNotificationConstraint, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSServiceCatalogLaunchNotificationConstraint:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSServiceCatalogLaunchNotificationConstraint not found", name)
}

// GetAllAWSServiceCatalogLaunchRoleConstraintResources retrieves all AWSServiceCatalogLaunchRoleConstraint items from an AWS CloudFormation template
func (t *Template) GetAllAWSServiceCatalogLaunchRoleConstraintResources() map[string]*AWSServiceCatalogLaunchRoleConstraint {
	results := map[string]*AWSServiceCatalogLaunchRoleConstraint{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSServiceCatalogLaunchRoleConstraint:
			results[name] = resource
		}
	}
	return results
}

// GetAWSServiceCatalogLaunchRoleConstraintWithName retrieves all AWSServiceCatalogLaunchRoleConstraint items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSServiceCatalogLaunchRoleConstraintWithName(name string) (*AWSServiceCatalogLaunchRoleConstraint, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSServiceCatalogLaunchRoleConstraint:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSServiceCatalogLaunchRoleConstraint not found", name)
}

// GetAllAWSServiceCatalogLaunchTemplateConstraintResources retrieves all AWSServiceCatalogLaunchTemplateConstraint items from an AWS CloudFormation template
func (t *Template) GetAllAWSServiceCatalogLaunchTemplateConstraintResources() map[string]*AWSServiceCatalogLaunchTemplateConstraint {
	results := map[string]*AWSServiceCatalogLaunchTemplateConstraint{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSServiceCatalogLaunchTemplateConstraint:
			results[name] = resource
		}
	}
	return results
}

// GetAWSServiceCatalogLaunchTemplateConstraintWithName retrieves all AWSServiceCatalogLaunchTemplateConstraint items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSServiceCatalogLaunchTemplateConstraintWithName(name string) (*AWSServiceCatalogLaunchTemplateConstraint, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSServiceCatalogLaunchTemplateConstraint:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSServiceCatalogLaunchTemplateConstraint not found", name)
}

// GetAllAWSServiceCatalogPortfolioResources retrieves all AWSServiceCatalogPortfolio items from an AWS CloudFormation template
func (t *Template) GetAllAWSServiceCatalogPortfolioResources() map[string]*AWSServiceCatalogPortfolio {
	results := map[string]*AWSServiceCatalogPortfolio{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSServiceCatalogPortfolio:
			results[name] = resource
		}
	}
	return results
}

// GetAWSServiceCatalogPortfolioWithName retrieves all AWSServiceCatalogPortfolio items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSServiceCatalogPortfolioWithName(name string) (*AWSServiceCatalogPortfolio, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSServiceCatalogPortfolio:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSServiceCatalogPortfolio not found", name)
}

// GetAllAWSServiceCatalogPortfolioPrincipalAssociationResources retrieves all AWSServiceCatalogPortfolioPrincipalAssociation items from an AWS CloudFormation template
func (t *Template) GetAllAWSServiceCatalogPortfolioPrincipalAssociationResources() map[string]*AWSServiceCatalogPortfolioPrincipalAssociation {
	results := map[string]*AWSServiceCatalogPortfolioPrincipalAssociation{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSServiceCatalogPortfolioPrincipalAssociation:
			results[name] = resource
		}
	}
	return results
}

// GetAWSServiceCatalogPortfolioPrincipalAssociationWithName retrieves all AWSServiceCatalogPortfolioPrincipalAssociation items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSServiceCatalogPortfolioPrincipalAssociationWithName(name string) (*AWSServiceCatalogPortfolioPrincipalAssociation, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSServiceCatalogPortfolioPrincipalAssociation:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSServiceCatalogPortfolioPrincipalAssociation not found", name)
}

// GetAllAWSServiceCatalogPortfolioProductAssociationResources retrieves all AWSServiceCatalogPortfolioProductAssociation items from an AWS CloudFormation template
func (t *Template) GetAllAWSServiceCatalogPortfolioProductAssociationResources() map[string]*AWSServiceCatalogPortfolioProductAssociation {
	results := map[string]*AWSServiceCatalogPortfolioProductAssociation{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSServiceCatalogPortfolioProductAssociation:
			results[name] = resource
		}
	}
	return results
}

// GetAWSServiceCatalogPortfolioProductAssociationWithName retrieves all AWSServiceCatalogPortfolioProductAssociation items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSServiceCatalogPortfolioProductAssociationWithName(name string) (*AWSServiceCatalogPortfolioProductAssociation, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSServiceCatalogPortfolioProductAssociation:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSServiceCatalogPortfolioProductAssociation not found", name)
}

// GetAllAWSServiceCatalogPortfolioShareResources retrieves all AWSServiceCatalogPortfolioShare items from an AWS CloudFormation template
func (t *Template) GetAllAWSServiceCatalogPortfolioShareResources() map[string]*AWSServiceCatalogPortfolioShare {
	results := map[string]*AWSServiceCatalogPortfolioShare{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSServiceCatalogPortfolioShare:
			results[name] = resource
		}
	}
	return results
}

// GetAWSServiceCatalogPortfolioShareWithName retrieves all AWSServiceCatalogPortfolioShare items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSServiceCatalogPortfolioShareWithName(name string) (*AWSServiceCatalogPortfolioShare, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSServiceCatalogPortfolioShare:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSServiceCatalogPortfolioShare not found", name)
}

// GetAllAWSServiceCatalogTagOptionResources retrieves all AWSServiceCatalogTagOption items from an AWS CloudFormation template
func (t *Template) GetAllAWSServiceCatalogTagOptionResources() map[string]*AWSServiceCatalogTagOption {
	results := map[string]*AWSServiceCatalogTagOption{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSServiceCatalogTagOption:
			results[name] = resource
		}
	}
	return results
}

// GetAWSServiceCatalogTagOptionWithName retrieves all AWSServiceCatalogTagOption items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSServiceCatalogTagOptionWithName(name string) (*AWSServiceCatalogTagOption, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSServiceCatalogTagOption:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSServiceCatalogTagOption not found", name)
}

// GetAllAWSServiceCatalogTagOptionAssociationResources retrieves all AWSServiceCatalogTagOptionAssociation items from an AWS CloudFormation template
func (t *Template) GetAllAWSServiceCatalogTagOptionAssociationResources() map[string]*AWSServiceCatalogTagOptionAssociation {
	results := map[string]*AWSServiceCatalogTagOptionAssociation{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSServiceCatalogTagOptionAssociation:
			results[name] = resource
		}
	}
	return results
}

// GetAWSServiceCatalogTagOptionAssociationWithName retrieves all AWSServiceCatalogTagOptionAssociation items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSServiceCatalogTagOptionAssociationWithName(name string) (*AWSServiceCatalogTagOptionAssociation, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSServiceCatalogTagOptionAssociation:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSServiceCatalogTagOptionAssociation not found", name)
}

// GetAllAWSServiceDiscoveryHttpNamespaceResources retrieves all AWSServiceDiscoveryHttpNamespace items from an AWS CloudFormation template
func (t *Template) GetAllAWSServiceDiscoveryHttpNamespaceResources() map[string]*AWSServiceDiscoveryHttpNamespace {
	results := map[string]*AWSServiceDiscoveryHttpNamespace{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSServiceDiscoveryHttpNamespace:
			results[name] = resource
		}
	}
	return results
}

// GetAWSServiceDiscoveryHttpNamespaceWithName retrieves all AWSServiceDiscoveryHttpNamespace items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSServiceDiscoveryHttpNamespaceWithName(name string) (*AWSServiceDiscoveryHttpNamespace, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSServiceDiscoveryHttpNamespace:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSServiceDiscoveryHttpNamespace not found", name)
}

// GetAllAWSServiceDiscoveryInstanceResources retrieves all AWSServiceDiscoveryInstance items from an AWS CloudFormation template
func (t *Template) GetAllAWSServiceDiscoveryInstanceResources() map[string]*AWSServiceDiscoveryInstance {
	results := map[string]*AWSServiceDiscoveryInstance{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSServiceDiscoveryInstance:
			results[name] = resource
		}
	}
	return results
}

// GetAWSServiceDiscoveryInstanceWithName retrieves all AWSServiceDiscoveryInstance items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSServiceDiscoveryInstanceWithName(name string) (*AWSServiceDiscoveryInstance, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSServiceDiscoveryInstance:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSServiceDiscoveryInstance not found", name)
}

// GetAllAWSServiceDiscoveryPrivateDnsNamespaceResources retrieves all AWSServiceDiscoveryPrivateDnsNamespace items from an AWS CloudFormation template
func (t *Template) GetAllAWSServiceDiscoveryPrivateDnsNamespaceResources() map[string]*AWSServiceDiscoveryPrivateDnsNamespace {
	results := map[string]*AWSServiceDiscoveryPrivateDnsNamespace{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSServiceDiscoveryPrivateDnsNamespace:
			results[name] = resource
		}
	}
	return results
}

// GetAWSServiceDiscoveryPrivateDnsNamespaceWithName retrieves all AWSServiceDiscoveryPrivateDnsNamespace items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSServiceDiscoveryPrivateDnsNamespaceWithName(name string) (*AWSServiceDiscoveryPrivateDnsNamespace, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSServiceDiscoveryPrivateDnsNamespace:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSServiceDiscoveryPrivateDnsNamespace not found", name)
}

// GetAllAWSServiceDiscoveryPublicDnsNamespaceResources retrieves all AWSServiceDiscoveryPublicDnsNamespace items from an AWS CloudFormation template
func (t *Template) GetAllAWSServiceDiscoveryPublicDnsNamespaceResources() map[string]*AWSServiceDiscoveryPublicDnsNamespace {
	results := map[string]*AWSServiceDiscoveryPublicDnsNamespace{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSServiceDiscoveryPublicDnsNamespace:
			results[name] = resource
		}
	}
	return results
}

// GetAWSServiceDiscoveryPublicDnsNamespaceWithName retrieves all AWSServiceDiscoveryPublicDnsNamespace items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSServiceDiscoveryPublicDnsNamespaceWithName(name string) (*AWSServiceDiscoveryPublicDnsNamespace, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSServiceDiscoveryPublicDnsNamespace:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSServiceDiscoveryPublicDnsNamespace not found", name)
}

// GetAllAWSServiceDiscoveryServiceResources retrieves all AWSServiceDiscoveryService items from an AWS CloudFormation template
func (t *Template) GetAllAWSServiceDiscoveryServiceResources() map[string]*AWSServiceDiscoveryService {
	results := map[string]*AWSServiceDiscoveryService{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSServiceDiscoveryService:
			results[name] = resource
		}
	}
	return results
}

// GetAWSServiceDiscoveryServiceWithName retrieves all AWSServiceDiscoveryService items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSServiceDiscoveryServiceWithName(name string) (*AWSServiceDiscoveryService, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSServiceDiscoveryService:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSServiceDiscoveryService not found", name)
}

// GetAllAWSStepFunctionsActivityResources retrieves all AWSStepFunctionsActivity items from an AWS CloudFormation template
func (t *Template) GetAllAWSStepFunctionsActivityResources() map[string]*AWSStepFunctionsActivity {
	results := map[string]*AWSStepFunctionsActivity{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSStepFunctionsActivity:
			results[name] = resource
		}
	}
	return results
}

// GetAWSStepFunctionsActivityWithName retrieves all AWSStepFunctionsActivity items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSStepFunctionsActivityWithName(name string) (*AWSStepFunctionsActivity, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSStepFunctionsActivity:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSStepFunctionsActivity not found", name)
}

// GetAllAWSStepFunctionsStateMachineResources retrieves all AWSStepFunctionsStateMachine items from an AWS CloudFormation template
func (t *Template) GetAllAWSStepFunctionsStateMachineResources() map[string]*AWSStepFunctionsStateMachine {
	results := map[string]*AWSStepFunctionsStateMachine{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSStepFunctionsStateMachine:
			results[name] = resource
		}
	}
	return results
}

// GetAWSStepFunctionsStateMachineWithName retrieves all AWSStepFunctionsStateMachine items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSStepFunctionsStateMachineWithName(name string) (*AWSStepFunctionsStateMachine, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSStepFunctionsStateMachine:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSStepFunctionsStateMachine not found", name)
}

// GetAllAWSWAFByteMatchSetResources retrieves all AWSWAFByteMatchSet items from an AWS CloudFormation template
func (t *Template) GetAllAWSWAFByteMatchSetResources() map[string]*AWSWAFByteMatchSet {
	results := map[string]*AWSWAFByteMatchSet{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSWAFByteMatchSet:
			results[name] = resource
		}
	}
	return results
}

// GetAWSWAFByteMatchSetWithName retrieves all AWSWAFByteMatchSet items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSWAFByteMatchSetWithName(name string) (*AWSWAFByteMatchSet, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSWAFByteMatchSet:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSWAFByteMatchSet not found", name)
}

// GetAllAWSWAFIPSetResources retrieves all AWSWAFIPSet items from an AWS CloudFormation template
func (t *Template) GetAllAWSWAFIPSetResources() map[string]*AWSWAFIPSet {
	results := map[string]*AWSWAFIPSet{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSWAFIPSet:
			results[name] = resource
		}
	}
	return results
}

// GetAWSWAFIPSetWithName retrieves all AWSWAFIPSet items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSWAFIPSetWithName(name string) (*AWSWAFIPSet, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSWAFIPSet:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSWAFIPSet not found", name)
}

// GetAllAWSWAFRuleResources retrieves all AWSWAFRule items from an AWS CloudFormation template
func (t *Template) GetAllAWSWAFRuleResources() map[string]*AWSWAFRule {
	results := map[string]*AWSWAFRule{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSWAFRule:
			results[name] = resource
		}
	}
	return results
}

// GetAWSWAFRuleWithName retrieves all AWSWAFRule items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSWAFRuleWithName(name string) (*AWSWAFRule, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSWAFRule:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSWAFRule not found", name)
}

// GetAllAWSWAFSizeConstraintSetResources retrieves all AWSWAFSizeConstraintSet items from an AWS CloudFormation template
func (t *Template) GetAllAWSWAFSizeConstraintSetResources() map[string]*AWSWAFSizeConstraintSet {
	results := map[string]*AWSWAFSizeConstraintSet{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSWAFSizeConstraintSet:
			results[name] = resource
		}
	}
	return results
}

// GetAWSWAFSizeConstraintSetWithName retrieves all AWSWAFSizeConstraintSet items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSWAFSizeConstraintSetWithName(name string) (*AWSWAFSizeConstraintSet, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSWAFSizeConstraintSet:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSWAFSizeConstraintSet not found", name)
}

// GetAllAWSWAFSqlInjectionMatchSetResources retrieves all AWSWAFSqlInjectionMatchSet items from an AWS CloudFormation template
func (t *Template) GetAllAWSWAFSqlInjectionMatchSetResources() map[string]*AWSWAFSqlInjectionMatchSet {
	results := map[string]*AWSWAFSqlInjectionMatchSet{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSWAFSqlInjectionMatchSet:
			results[name] = resource
		}
	}
	return results
}

// GetAWSWAFSqlInjectionMatchSetWithName retrieves all AWSWAFSqlInjectionMatchSet items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSWAFSqlInjectionMatchSetWithName(name string) (*AWSWAFSqlInjectionMatchSet, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSWAFSqlInjectionMatchSet:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSWAFSqlInjectionMatchSet not found", name)
}

// GetAllAWSWAFWebACLResources retrieves all AWSWAFWebACL items from an AWS CloudFormation template
func (t *Template) GetAllAWSWAFWebACLResources() map[string]*AWSWAFWebACL {
	results := map[string]*AWSWAFWebACL{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSWAFWebACL:
			results[name] = resource
		}
	}
	return results
}

// GetAWSWAFWebACLWithName retrieves all AWSWAFWebACL items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSWAFWebACLWithName(name string) (*AWSWAFWebACL, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSWAFWebACL:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSWAFWebACL not found", name)
}

// GetAllAWSWAFXssMatchSetResources retrieves all AWSWAFXssMatchSet items from an AWS CloudFormation template
func (t *Template) GetAllAWSWAFXssMatchSetResources() map[string]*AWSWAFXssMatchSet {
	results := map[string]*AWSWAFXssMatchSet{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSWAFXssMatchSet:
			results[name] = resource
		}
	}
	return results
}

// GetAWSWAFXssMatchSetWithName retrieves all AWSWAFXssMatchSet items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSWAFXssMatchSetWithName(name string) (*AWSWAFXssMatchSet, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSWAFXssMatchSet:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSWAFXssMatchSet not found", name)
}

// GetAllAWSWAFRegionalByteMatchSetResources retrieves all AWSWAFRegionalByteMatchSet items from an AWS CloudFormation template
func (t *Template) GetAllAWSWAFRegionalByteMatchSetResources() map[string]*AWSWAFRegionalByteMatchSet {
	results := map[string]*AWSWAFRegionalByteMatchSet{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSWAFRegionalByteMatchSet:
			results[name] = resource
		}
	}
	return results
}

// GetAWSWAFRegionalByteMatchSetWithName retrieves all AWSWAFRegionalByteMatchSet items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSWAFRegionalByteMatchSetWithName(name string) (*AWSWAFRegionalByteMatchSet, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSWAFRegionalByteMatchSet:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSWAFRegionalByteMatchSet not found", name)
}

// GetAllAWSWAFRegionalIPSetResources retrieves all AWSWAFRegionalIPSet items from an AWS CloudFormation template
func (t *Template) GetAllAWSWAFRegionalIPSetResources() map[string]*AWSWAFRegionalIPSet {
	results := map[string]*AWSWAFRegionalIPSet{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSWAFRegionalIPSet:
			results[name] = resource
		}
	}
	return results
}

// GetAWSWAFRegionalIPSetWithName retrieves all AWSWAFRegionalIPSet items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSWAFRegionalIPSetWithName(name string) (*AWSWAFRegionalIPSet, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSWAFRegionalIPSet:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSWAFRegionalIPSet not found", name)
}

// GetAllAWSWAFRegionalRuleResources retrieves all AWSWAFRegionalRule items from an AWS CloudFormation template
func (t *Template) GetAllAWSWAFRegionalRuleResources() map[string]*AWSWAFRegionalRule {
	results := map[string]*AWSWAFRegionalRule{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSWAFRegionalRule:
			results[name] = resource
		}
	}
	return results
}

// GetAWSWAFRegionalRuleWithName retrieves all AWSWAFRegionalRule items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSWAFRegionalRuleWithName(name string) (*AWSWAFRegionalRule, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSWAFRegionalRule:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSWAFRegionalRule not found", name)
}

// GetAllAWSWAFRegionalSizeConstraintSetResources retrieves all AWSWAFRegionalSizeConstraintSet items from an AWS CloudFormation template
func (t *Template) GetAllAWSWAFRegionalSizeConstraintSetResources() map[string]*AWSWAFRegionalSizeConstraintSet {
	results := map[string]*AWSWAFRegionalSizeConstraintSet{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSWAFRegionalSizeConstraintSet:
			results[name] = resource
		}
	}
	return results
}

// GetAWSWAFRegionalSizeConstraintSetWithName retrieves all AWSWAFRegionalSizeConstraintSet items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSWAFRegionalSizeConstraintSetWithName(name string) (*AWSWAFRegionalSizeConstraintSet, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSWAFRegionalSizeConstraintSet:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSWAFRegionalSizeConstraintSet not found", name)
}

// GetAllAWSWAFRegionalSqlInjectionMatchSetResources retrieves all AWSWAFRegionalSqlInjectionMatchSet items from an AWS CloudFormation template
func (t *Template) GetAllAWSWAFRegionalSqlInjectionMatchSetResources() map[string]*AWSWAFRegionalSqlInjectionMatchSet {
	results := map[string]*AWSWAFRegionalSqlInjectionMatchSet{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSWAFRegionalSqlInjectionMatchSet:
			results[name] = resource
		}
	}
	return results
}

// GetAWSWAFRegionalSqlInjectionMatchSetWithName retrieves all AWSWAFRegionalSqlInjectionMatchSet items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSWAFRegionalSqlInjectionMatchSetWithName(name string) (*AWSWAFRegionalSqlInjectionMatchSet, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSWAFRegionalSqlInjectionMatchSet:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSWAFRegionalSqlInjectionMatchSet not found", name)
}

// GetAllAWSWAFRegionalWebACLResources retrieves all AWSWAFRegionalWebACL items from an AWS CloudFormation template
func (t *Template) GetAllAWSWAFRegionalWebACLResources() map[string]*AWSWAFRegionalWebACL {
	results := map[string]*AWSWAFRegionalWebACL{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSWAFRegionalWebACL:
			results[name] = resource
		}
	}
	return results
}

// GetAWSWAFRegionalWebACLWithName retrieves all AWSWAFRegionalWebACL items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSWAFRegionalWebACLWithName(name string) (*AWSWAFRegionalWebACL, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSWAFRegionalWebACL:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSWAFRegionalWebACL not found", name)
}

// GetAllAWSWAFRegionalWebACLAssociationResources retrieves all AWSWAFRegionalWebACLAssociation items from an AWS CloudFormation template
func (t *Template) GetAllAWSWAFRegionalWebACLAssociationResources() map[string]*AWSWAFRegionalWebACLAssociation {
	results := map[string]*AWSWAFRegionalWebACLAssociation{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSWAFRegionalWebACLAssociation:
			results[name] = resource
		}
	}
	return results
}

// GetAWSWAFRegionalWebACLAssociationWithName retrieves all AWSWAFRegionalWebACLAssociation items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSWAFRegionalWebACLAssociationWithName(name string) (*AWSWAFRegionalWebACLAssociation, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSWAFRegionalWebACLAssociation:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSWAFRegionalWebACLAssociation not found", name)
}

// GetAllAWSWAFRegionalXssMatchSetResources retrieves all AWSWAFRegionalXssMatchSet items from an AWS CloudFormation template
func (t *Template) GetAllAWSWAFRegionalXssMatchSetResources() map[string]*AWSWAFRegionalXssMatchSet {
	results := map[string]*AWSWAFRegionalXssMatchSet{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSWAFRegionalXssMatchSet:
			results[name] = resource
		}
	}
	return results
}

// GetAWSWAFRegionalXssMatchSetWithName retrieves all AWSWAFRegionalXssMatchSet items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSWAFRegionalXssMatchSetWithName(name string) (*AWSWAFRegionalXssMatchSet, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSWAFRegionalXssMatchSet:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSWAFRegionalXssMatchSet not found", name)
}

// GetAllAWSWorkSpacesWorkspaceResources retrieves all AWSWorkSpacesWorkspace items from an AWS CloudFormation template
func (t *Template) GetAllAWSWorkSpacesWorkspaceResources() map[string]*AWSWorkSpacesWorkspace {
	results := map[string]*AWSWorkSpacesWorkspace{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AWSWorkSpacesWorkspace:
			results[name] = resource
		}
	}
	return results
}

// GetAWSWorkSpacesWorkspaceWithName retrieves all AWSWorkSpacesWorkspace items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSWorkSpacesWorkspaceWithName(name string) (*AWSWorkSpacesWorkspace, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AWSWorkSpacesWorkspace:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSWorkSpacesWorkspace not found", name)
}

// GetAllAlexaASKSkillResources retrieves all AlexaASKSkill items from an AWS CloudFormation template
func (t *Template) GetAllAlexaASKSkillResources() map[string]*AlexaASKSkill {
	results := map[string]*AlexaASKSkill{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *AlexaASKSkill:
			results[name] = resource
		}
	}
	return results
}

// GetAlexaASKSkillWithName retrieves all AlexaASKSkill items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAlexaASKSkillWithName(name string) (*AlexaASKSkill, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *AlexaASKSkill:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AlexaASKSkill not found", name)
}
