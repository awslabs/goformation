package cloudformation

import (
	"fmt"
	"github.com/awslabs/goformation/cloudformation/resources"
)

// AllResources fetches an iterable map all CloudFormation and SAM resources
func AllResources() map[string]Resource {
	return map[string]Resource{
		"AWS::AmazonMQ::Broker":                                       &resources.AWSAmazonMQBroker{},
		"AWS::AmazonMQ::Configuration":                                &resources.AWSAmazonMQConfiguration{},
		"AWS::AmazonMQ::ConfigurationAssociation":                     &resources.AWSAmazonMQConfigurationAssociation{},
		"AWS::ApiGateway::Account":                                    &resources.AWSApiGatewayAccount{},
		"AWS::ApiGateway::ApiKey":                                     &resources.AWSApiGatewayApiKey{},
		"AWS::ApiGateway::Authorizer":                                 &resources.AWSApiGatewayAuthorizer{},
		"AWS::ApiGateway::BasePathMapping":                            &resources.AWSApiGatewayBasePathMapping{},
		"AWS::ApiGateway::ClientCertificate":                          &resources.AWSApiGatewayClientCertificate{},
		"AWS::ApiGateway::Deployment":                                 &resources.AWSApiGatewayDeployment{},
		"AWS::ApiGateway::DocumentationPart":                          &resources.AWSApiGatewayDocumentationPart{},
		"AWS::ApiGateway::DocumentationVersion":                       &resources.AWSApiGatewayDocumentationVersion{},
		"AWS::ApiGateway::DomainName":                                 &resources.AWSApiGatewayDomainName{},
		"AWS::ApiGateway::GatewayResponse":                            &resources.AWSApiGatewayGatewayResponse{},
		"AWS::ApiGateway::Method":                                     &resources.AWSApiGatewayMethod{},
		"AWS::ApiGateway::Model":                                      &resources.AWSApiGatewayModel{},
		"AWS::ApiGateway::RequestValidator":                           &resources.AWSApiGatewayRequestValidator{},
		"AWS::ApiGateway::Resource":                                   &resources.AWSApiGatewayResource{},
		"AWS::ApiGateway::RestApi":                                    &resources.AWSApiGatewayRestApi{},
		"AWS::ApiGateway::Stage":                                      &resources.AWSApiGatewayStage{},
		"AWS::ApiGateway::UsagePlan":                                  &resources.AWSApiGatewayUsagePlan{},
		"AWS::ApiGateway::UsagePlanKey":                               &resources.AWSApiGatewayUsagePlanKey{},
		"AWS::ApiGateway::VpcLink":                                    &resources.AWSApiGatewayVpcLink{},
		"AWS::ApiGatewayV2::Api":                                      &resources.AWSApiGatewayV2Api{},
		"AWS::ApiGatewayV2::Authorizer":                               &resources.AWSApiGatewayV2Authorizer{},
		"AWS::ApiGatewayV2::Deployment":                               &resources.AWSApiGatewayV2Deployment{},
		"AWS::ApiGatewayV2::Integration":                              &resources.AWSApiGatewayV2Integration{},
		"AWS::ApiGatewayV2::IntegrationResponse":                      &resources.AWSApiGatewayV2IntegrationResponse{},
		"AWS::ApiGatewayV2::Model":                                    &resources.AWSApiGatewayV2Model{},
		"AWS::ApiGatewayV2::Route":                                    &resources.AWSApiGatewayV2Route{},
		"AWS::ApiGatewayV2::RouteResponse":                            &resources.AWSApiGatewayV2RouteResponse{},
		"AWS::ApiGatewayV2::Stage":                                    &resources.AWSApiGatewayV2Stage{},
		"AWS::AppStream::DirectoryConfig":                             &resources.AWSAppStreamDirectoryConfig{},
		"AWS::AppStream::Fleet":                                       &resources.AWSAppStreamFleet{},
		"AWS::AppStream::ImageBuilder":                                &resources.AWSAppStreamImageBuilder{},
		"AWS::AppStream::Stack":                                       &resources.AWSAppStreamStack{},
		"AWS::AppStream::StackFleetAssociation":                       &resources.AWSAppStreamStackFleetAssociation{},
		"AWS::AppStream::StackUserAssociation":                        &resources.AWSAppStreamStackUserAssociation{},
		"AWS::AppStream::User":                                        &resources.AWSAppStreamUser{},
		"AWS::AppSync::ApiKey":                                        &resources.AWSAppSyncApiKey{},
		"AWS::AppSync::DataSource":                                    &resources.AWSAppSyncDataSource{},
		"AWS::AppSync::FunctionConfiguration":                         &resources.AWSAppSyncFunctionConfiguration{},
		"AWS::AppSync::GraphQLApi":                                    &resources.AWSAppSyncGraphQLApi{},
		"AWS::AppSync::GraphQLSchema":                                 &resources.AWSAppSyncGraphQLSchema{},
		"AWS::AppSync::Resolver":                                      &resources.AWSAppSyncResolver{},
		"AWS::ApplicationAutoScaling::ScalableTarget":                 &resources.AWSApplicationAutoScalingScalableTarget{},
		"AWS::ApplicationAutoScaling::ScalingPolicy":                  &resources.AWSApplicationAutoScalingScalingPolicy{},
		"AWS::Athena::NamedQuery":                                     &resources.AWSAthenaNamedQuery{},
		"AWS::AutoScaling::AutoScalingGroup":                          &resources.AWSAutoScalingAutoScalingGroup{},
		"AWS::AutoScaling::LaunchConfiguration":                       &resources.AWSAutoScalingLaunchConfiguration{},
		"AWS::AutoScaling::LifecycleHook":                             &resources.AWSAutoScalingLifecycleHook{},
		"AWS::AutoScaling::ScalingPolicy":                             &resources.AWSAutoScalingScalingPolicy{},
		"AWS::AutoScaling::ScheduledAction":                           &resources.AWSAutoScalingScheduledAction{},
		"AWS::AutoScalingPlans::ScalingPlan":                          &resources.AWSAutoScalingPlansScalingPlan{},
		"AWS::Batch::ComputeEnvironment":                              &resources.AWSBatchComputeEnvironment{},
		"AWS::Batch::JobDefinition":                                   &resources.AWSBatchJobDefinition{},
		"AWS::Batch::JobQueue":                                        &resources.AWSBatchJobQueue{},
		"AWS::Budgets::Budget":                                        &resources.AWSBudgetsBudget{},
		"AWS::CertificateManager::Certificate":                        &resources.AWSCertificateManagerCertificate{},
		"AWS::Cloud9::EnvironmentEC2":                                 &resources.AWSCloud9EnvironmentEC2{},
		"AWS::CloudFormation::CustomResource":                         &resources.AWSCloudFormationCustomResource{},
		"AWS::CloudFormation::Macro":                                  &resources.AWSCloudFormationMacro{},
		"AWS::CloudFormation::Stack":                                  &resources.AWSCloudFormationStack{},
		"AWS::CloudFormation::WaitCondition":                          &resources.AWSCloudFormationWaitCondition{},
		"AWS::CloudFormation::WaitConditionHandle":                    &resources.AWSCloudFormationWaitConditionHandle{},
		"AWS::CloudFront::CloudFrontOriginAccessIdentity":             &resources.AWSCloudFrontCloudFrontOriginAccessIdentity{},
		"AWS::CloudFront::Distribution":                               &resources.AWSCloudFrontDistribution{},
		"AWS::CloudFront::StreamingDistribution":                      &resources.AWSCloudFrontStreamingDistribution{},
		"AWS::CloudTrail::Trail":                                      &resources.AWSCloudTrailTrail{},
		"AWS::CloudWatch::Alarm":                                      &resources.AWSCloudWatchAlarm{},
		"AWS::CloudWatch::Dashboard":                                  &resources.AWSCloudWatchDashboard{},
		"AWS::CodeBuild::Project":                                     &resources.AWSCodeBuildProject{},
		"AWS::CodeCommit::Repository":                                 &resources.AWSCodeCommitRepository{},
		"AWS::CodeDeploy::Application":                                &resources.AWSCodeDeployApplication{},
		"AWS::CodeDeploy::DeploymentConfig":                           &resources.AWSCodeDeployDeploymentConfig{},
		"AWS::CodeDeploy::DeploymentGroup":                            &resources.AWSCodeDeployDeploymentGroup{},
		"AWS::CodePipeline::CustomActionType":                         &resources.AWSCodePipelineCustomActionType{},
		"AWS::CodePipeline::Pipeline":                                 &resources.AWSCodePipelinePipeline{},
		"AWS::CodePipeline::Webhook":                                  &resources.AWSCodePipelineWebhook{},
		"AWS::Cognito::IdentityPool":                                  &resources.AWSCognitoIdentityPool{},
		"AWS::Cognito::IdentityPoolRoleAttachment":                    &resources.AWSCognitoIdentityPoolRoleAttachment{},
		"AWS::Cognito::UserPool":                                      &resources.AWSCognitoUserPool{},
		"AWS::Cognito::UserPoolClient":                                &resources.AWSCognitoUserPoolClient{},
		"AWS::Cognito::UserPoolGroup":                                 &resources.AWSCognitoUserPoolGroup{},
		"AWS::Cognito::UserPoolUser":                                  &resources.AWSCognitoUserPoolUser{},
		"AWS::Cognito::UserPoolUserToGroupAttachment":                 &resources.AWSCognitoUserPoolUserToGroupAttachment{},
		"AWS::Config::AggregationAuthorization":                       &resources.AWSConfigAggregationAuthorization{},
		"AWS::Config::ConfigRule":                                     &resources.AWSConfigConfigRule{},
		"AWS::Config::ConfigurationAggregator":                        &resources.AWSConfigConfigurationAggregator{},
		"AWS::Config::ConfigurationRecorder":                          &resources.AWSConfigConfigurationRecorder{},
		"AWS::Config::DeliveryChannel":                                &resources.AWSConfigDeliveryChannel{},
		"AWS::DAX::Cluster":                                           &resources.AWSDAXCluster{},
		"AWS::DAX::ParameterGroup":                                    &resources.AWSDAXParameterGroup{},
		"AWS::DAX::SubnetGroup":                                       &resources.AWSDAXSubnetGroup{},
		"AWS::DLM::LifecyclePolicy":                                   &resources.AWSDLMLifecyclePolicy{},
		"AWS::DMS::Certificate":                                       &resources.AWSDMSCertificate{},
		"AWS::DMS::Endpoint":                                          &resources.AWSDMSEndpoint{},
		"AWS::DMS::EventSubscription":                                 &resources.AWSDMSEventSubscription{},
		"AWS::DMS::ReplicationInstance":                               &resources.AWSDMSReplicationInstance{},
		"AWS::DMS::ReplicationSubnetGroup":                            &resources.AWSDMSReplicationSubnetGroup{},
		"AWS::DMS::ReplicationTask":                                   &resources.AWSDMSReplicationTask{},
		"AWS::DataPipeline::Pipeline":                                 &resources.AWSDataPipelinePipeline{},
		"AWS::DirectoryService::MicrosoftAD":                          &resources.AWSDirectoryServiceMicrosoftAD{},
		"AWS::DirectoryService::SimpleAD":                             &resources.AWSDirectoryServiceSimpleAD{},
		"AWS::DocDB::DBCluster":                                       &resources.AWSDocDBDBCluster{},
		"AWS::DocDB::DBClusterParameterGroup":                         &resources.AWSDocDBDBClusterParameterGroup{},
		"AWS::DocDB::DBInstance":                                      &resources.AWSDocDBDBInstance{},
		"AWS::DocDB::DBSubnetGroup":                                   &resources.AWSDocDBDBSubnetGroup{},
		"AWS::DynamoDB::Table":                                        &resources.AWSDynamoDBTable{},
		"AWS::EC2::CustomerGateway":                                   &resources.AWSEC2CustomerGateway{},
		"AWS::EC2::DHCPOptions":                                       &resources.AWSEC2DHCPOptions{},
		"AWS::EC2::EC2Fleet":                                          &resources.AWSEC2EC2Fleet{},
		"AWS::EC2::EIP":                                               &resources.AWSEC2EIP{},
		"AWS::EC2::EIPAssociation":                                    &resources.AWSEC2EIPAssociation{},
		"AWS::EC2::EgressOnlyInternetGateway":                         &resources.AWSEC2EgressOnlyInternetGateway{},
		"AWS::EC2::FlowLog":                                           &resources.AWSEC2FlowLog{},
		"AWS::EC2::Host":                                              &resources.AWSEC2Host{},
		"AWS::EC2::Instance":                                          &resources.AWSEC2Instance{},
		"AWS::EC2::InternetGateway":                                   &resources.AWSEC2InternetGateway{},
		"AWS::EC2::LaunchTemplate":                                    &resources.AWSEC2LaunchTemplate{},
		"AWS::EC2::NatGateway":                                        &resources.AWSEC2NatGateway{},
		"AWS::EC2::NetworkAcl":                                        &resources.AWSEC2NetworkAcl{},
		"AWS::EC2::NetworkAclEntry":                                   &resources.AWSEC2NetworkAclEntry{},
		"AWS::EC2::NetworkInterface":                                  &resources.AWSEC2NetworkInterface{},
		"AWS::EC2::NetworkInterfaceAttachment":                        &resources.AWSEC2NetworkInterfaceAttachment{},
		"AWS::EC2::NetworkInterfacePermission":                        &resources.AWSEC2NetworkInterfacePermission{},
		"AWS::EC2::PlacementGroup":                                    &resources.AWSEC2PlacementGroup{},
		"AWS::EC2::Route":                                             &resources.AWSEC2Route{},
		"AWS::EC2::RouteTable":                                        &resources.AWSEC2RouteTable{},
		"AWS::EC2::SecurityGroup":                                     &resources.AWSEC2SecurityGroup{},
		"AWS::EC2::SecurityGroupEgress":                               &resources.AWSEC2SecurityGroupEgress{},
		"AWS::EC2::SecurityGroupIngress":                              &resources.AWSEC2SecurityGroupIngress{},
		"AWS::EC2::SpotFleet":                                         &resources.AWSEC2SpotFleet{},
		"AWS::EC2::Subnet":                                            &resources.AWSEC2Subnet{},
		"AWS::EC2::SubnetCidrBlock":                                   &resources.AWSEC2SubnetCidrBlock{},
		"AWS::EC2::SubnetNetworkAclAssociation":                       &resources.AWSEC2SubnetNetworkAclAssociation{},
		"AWS::EC2::SubnetRouteTableAssociation":                       &resources.AWSEC2SubnetRouteTableAssociation{},
		"AWS::EC2::TransitGateway":                                    &resources.AWSEC2TransitGateway{},
		"AWS::EC2::TransitGatewayAttachment":                          &resources.AWSEC2TransitGatewayAttachment{},
		"AWS::EC2::TransitGatewayRoute":                               &resources.AWSEC2TransitGatewayRoute{},
		"AWS::EC2::TransitGatewayRouteTable":                          &resources.AWSEC2TransitGatewayRouteTable{},
		"AWS::EC2::TransitGatewayRouteTableAssociation":               &resources.AWSEC2TransitGatewayRouteTableAssociation{},
		"AWS::EC2::TransitGatewayRouteTablePropagation":               &resources.AWSEC2TransitGatewayRouteTablePropagation{},
		"AWS::EC2::TrunkInterfaceAssociation":                         &resources.AWSEC2TrunkInterfaceAssociation{},
		"AWS::EC2::VPC":                                               &resources.AWSEC2VPC{},
		"AWS::EC2::VPCCidrBlock":                                      &resources.AWSEC2VPCCidrBlock{},
		"AWS::EC2::VPCDHCPOptionsAssociation":                         &resources.AWSEC2VPCDHCPOptionsAssociation{},
		"AWS::EC2::VPCEndpoint":                                       &resources.AWSEC2VPCEndpoint{},
		"AWS::EC2::VPCEndpointConnectionNotification":                 &resources.AWSEC2VPCEndpointConnectionNotification{},
		"AWS::EC2::VPCEndpointServicePermissions":                     &resources.AWSEC2VPCEndpointServicePermissions{},
		"AWS::EC2::VPCGatewayAttachment":                              &resources.AWSEC2VPCGatewayAttachment{},
		"AWS::EC2::VPCPeeringConnection":                              &resources.AWSEC2VPCPeeringConnection{},
		"AWS::EC2::VPNConnection":                                     &resources.AWSEC2VPNConnection{},
		"AWS::EC2::VPNConnectionRoute":                                &resources.AWSEC2VPNConnectionRoute{},
		"AWS::EC2::VPNGateway":                                        &resources.AWSEC2VPNGateway{},
		"AWS::EC2::VPNGatewayRoutePropagation":                        &resources.AWSEC2VPNGatewayRoutePropagation{},
		"AWS::EC2::Volume":                                            &resources.AWSEC2Volume{},
		"AWS::EC2::VolumeAttachment":                                  &resources.AWSEC2VolumeAttachment{},
		"AWS::ECR::Repository":                                        &resources.AWSECRRepository{},
		"AWS::ECS::Cluster":                                           &resources.AWSECSCluster{},
		"AWS::ECS::Service":                                           &resources.AWSECSService{},
		"AWS::ECS::TaskDefinition":                                    &resources.AWSECSTaskDefinition{},
		"AWS::EFS::FileSystem":                                        &resources.AWSEFSFileSystem{},
		"AWS::EFS::MountTarget":                                       &resources.AWSEFSMountTarget{},
		"AWS::EKS::Cluster":                                           &resources.AWSEKSCluster{},
		"AWS::EMR::Cluster":                                           &resources.AWSEMRCluster{},
		"AWS::EMR::InstanceFleetConfig":                               &resources.AWSEMRInstanceFleetConfig{},
		"AWS::EMR::InstanceGroupConfig":                               &resources.AWSEMRInstanceGroupConfig{},
		"AWS::EMR::SecurityConfiguration":                             &resources.AWSEMRSecurityConfiguration{},
		"AWS::EMR::Step":                                              &resources.AWSEMRStep{},
		"AWS::ElastiCache::CacheCluster":                              &resources.AWSElastiCacheCacheCluster{},
		"AWS::ElastiCache::ParameterGroup":                            &resources.AWSElastiCacheParameterGroup{},
		"AWS::ElastiCache::ReplicationGroup":                          &resources.AWSElastiCacheReplicationGroup{},
		"AWS::ElastiCache::SecurityGroup":                             &resources.AWSElastiCacheSecurityGroup{},
		"AWS::ElastiCache::SecurityGroupIngress":                      &resources.AWSElastiCacheSecurityGroupIngress{},
		"AWS::ElastiCache::SubnetGroup":                               &resources.AWSElastiCacheSubnetGroup{},
		"AWS::ElasticBeanstalk::Application":                          &resources.AWSElasticBeanstalkApplication{},
		"AWS::ElasticBeanstalk::ApplicationVersion":                   &resources.AWSElasticBeanstalkApplicationVersion{},
		"AWS::ElasticBeanstalk::ConfigurationTemplate":                &resources.AWSElasticBeanstalkConfigurationTemplate{},
		"AWS::ElasticBeanstalk::Environment":                          &resources.AWSElasticBeanstalkEnvironment{},
		"AWS::ElasticLoadBalancing::LoadBalancer":                     &resources.AWSElasticLoadBalancingLoadBalancer{},
		"AWS::ElasticLoadBalancingV2::Listener":                       &resources.AWSElasticLoadBalancingV2Listener{},
		"AWS::ElasticLoadBalancingV2::ListenerCertificate":            &resources.AWSElasticLoadBalancingV2ListenerCertificate{},
		"AWS::ElasticLoadBalancingV2::ListenerRule":                   &resources.AWSElasticLoadBalancingV2ListenerRule{},
		"AWS::ElasticLoadBalancingV2::LoadBalancer":                   &resources.AWSElasticLoadBalancingV2LoadBalancer{},
		"AWS::ElasticLoadBalancingV2::TargetGroup":                    &resources.AWSElasticLoadBalancingV2TargetGroup{},
		"AWS::Elasticsearch::Domain":                                  &resources.AWSElasticsearchDomain{},
		"AWS::Events::EventBusPolicy":                                 &resources.AWSEventsEventBusPolicy{},
		"AWS::Events::Rule":                                           &resources.AWSEventsRule{},
		"AWS::FSx::FileSystem":                                        &resources.AWSFSxFileSystem{},
		"AWS::GameLift::Alias":                                        &resources.AWSGameLiftAlias{},
		"AWS::GameLift::Build":                                        &resources.AWSGameLiftBuild{},
		"AWS::GameLift::Fleet":                                        &resources.AWSGameLiftFleet{},
		"AWS::Glue::Classifier":                                       &resources.AWSGlueClassifier{},
		"AWS::Glue::Connection":                                       &resources.AWSGlueConnection{},
		"AWS::Glue::Crawler":                                          &resources.AWSGlueCrawler{},
		"AWS::Glue::Database":                                         &resources.AWSGlueDatabase{},
		"AWS::Glue::DevEndpoint":                                      &resources.AWSGlueDevEndpoint{},
		"AWS::Glue::Job":                                              &resources.AWSGlueJob{},
		"AWS::Glue::Partition":                                        &resources.AWSGluePartition{},
		"AWS::Glue::Table":                                            &resources.AWSGlueTable{},
		"AWS::Glue::Trigger":                                          &resources.AWSGlueTrigger{},
		"AWS::GuardDuty::Detector":                                    &resources.AWSGuardDutyDetector{},
		"AWS::GuardDuty::Filter":                                      &resources.AWSGuardDutyFilter{},
		"AWS::GuardDuty::IPSet":                                       &resources.AWSGuardDutyIPSet{},
		"AWS::GuardDuty::Master":                                      &resources.AWSGuardDutyMaster{},
		"AWS::GuardDuty::Member":                                      &resources.AWSGuardDutyMember{},
		"AWS::GuardDuty::ThreatIntelSet":                              &resources.AWSGuardDutyThreatIntelSet{},
		"AWS::IAM::AccessKey":                                         &resources.AWSIAMAccessKey{},
		"AWS::IAM::Group":                                             &resources.AWSIAMGroup{},
		"AWS::IAM::InstanceProfile":                                   &resources.AWSIAMInstanceProfile{},
		"AWS::IAM::ManagedPolicy":                                     &resources.AWSIAMManagedPolicy{},
		"AWS::IAM::Policy":                                            &resources.AWSIAMPolicy{},
		"AWS::IAM::Role":                                              &resources.AWSIAMRole{},
		"AWS::IAM::ServiceLinkedRole":                                 &resources.AWSIAMServiceLinkedRole{},
		"AWS::IAM::User":                                              &resources.AWSIAMUser{},
		"AWS::IAM::UserToGroupAddition":                               &resources.AWSIAMUserToGroupAddition{},
		"AWS::Inspector::AssessmentTarget":                            &resources.AWSInspectorAssessmentTarget{},
		"AWS::Inspector::AssessmentTemplate":                          &resources.AWSInspectorAssessmentTemplate{},
		"AWS::Inspector::ResourceGroup":                               &resources.AWSInspectorResourceGroup{},
		"AWS::IoT1Click::Device":                                      &resources.AWSIoT1ClickDevice{},
		"AWS::IoT1Click::Placement":                                   &resources.AWSIoT1ClickPlacement{},
		"AWS::IoT1Click::Project":                                     &resources.AWSIoT1ClickProject{},
		"AWS::IoT::Certificate":                                       &resources.AWSIoTCertificate{},
		"AWS::IoT::Policy":                                            &resources.AWSIoTPolicy{},
		"AWS::IoT::PolicyPrincipalAttachment":                         &resources.AWSIoTPolicyPrincipalAttachment{},
		"AWS::IoT::Thing":                                             &resources.AWSIoTThing{},
		"AWS::IoT::ThingPrincipalAttachment":                          &resources.AWSIoTThingPrincipalAttachment{},
		"AWS::IoT::TopicRule":                                         &resources.AWSIoTTopicRule{},
		"AWS::IoTAnalytics::Channel":                                  &resources.AWSIoTAnalyticsChannel{},
		"AWS::IoTAnalytics::Dataset":                                  &resources.AWSIoTAnalyticsDataset{},
		"AWS::IoTAnalytics::Datastore":                                &resources.AWSIoTAnalyticsDatastore{},
		"AWS::IoTAnalytics::Pipeline":                                 &resources.AWSIoTAnalyticsPipeline{},
		"AWS::KMS::Alias":                                             &resources.AWSKMSAlias{},
		"AWS::KMS::Key":                                               &resources.AWSKMSKey{},
		"AWS::Kinesis::Stream":                                        &resources.AWSKinesisStream{},
		"AWS::Kinesis::StreamConsumer":                                &resources.AWSKinesisStreamConsumer{},
		"AWS::KinesisAnalytics::Application":                          &resources.AWSKinesisAnalyticsApplication{},
		"AWS::KinesisAnalytics::ApplicationOutput":                    &resources.AWSKinesisAnalyticsApplicationOutput{},
		"AWS::KinesisAnalytics::ApplicationReferenceDataSource":       &resources.AWSKinesisAnalyticsApplicationReferenceDataSource{},
		"AWS::KinesisAnalyticsV2::Application":                        &resources.AWSKinesisAnalyticsV2Application{},
		"AWS::KinesisAnalyticsV2::ApplicationCloudWatchLoggingOption": &resources.AWSKinesisAnalyticsV2ApplicationCloudWatchLoggingOption{},
		"AWS::KinesisAnalyticsV2::ApplicationOutput":                  &resources.AWSKinesisAnalyticsV2ApplicationOutput{},
		"AWS::KinesisAnalyticsV2::ApplicationReferenceDataSource":     &resources.AWSKinesisAnalyticsV2ApplicationReferenceDataSource{},
		"AWS::KinesisFirehose::DeliveryStream":                        &resources.AWSKinesisFirehoseDeliveryStream{},
		"AWS::Lambda::Alias":                                          &resources.AWSLambdaAlias{},
		"AWS::Lambda::EventSourceMapping":                             &resources.AWSLambdaEventSourceMapping{},
		"AWS::Lambda::Function":                                       &resources.AWSLambdaFunction{},
		"AWS::Lambda::LayerVersion":                                   &resources.AWSLambdaLayerVersion{},
		"AWS::Lambda::LayerVersionPermission":                         &resources.AWSLambdaLayerVersionPermission{},
		"AWS::Lambda::Permission":                                     &resources.AWSLambdaPermission{},
		"AWS::Lambda::Version":                                        &resources.AWSLambdaVersion{},
		"AWS::Logs::Destination":                                      &resources.AWSLogsDestination{},
		"AWS::Logs::LogGroup":                                         &resources.AWSLogsLogGroup{},
		"AWS::Logs::LogStream":                                        &resources.AWSLogsLogStream{},
		"AWS::Logs::MetricFilter":                                     &resources.AWSLogsMetricFilter{},
		"AWS::Logs::SubscriptionFilter":                               &resources.AWSLogsSubscriptionFilter{},
		"AWS::Neptune::DBCluster":                                     &resources.AWSNeptuneDBCluster{},
		"AWS::Neptune::DBClusterParameterGroup":                       &resources.AWSNeptuneDBClusterParameterGroup{},
		"AWS::Neptune::DBInstance":                                    &resources.AWSNeptuneDBInstance{},
		"AWS::Neptune::DBParameterGroup":                              &resources.AWSNeptuneDBParameterGroup{},
		"AWS::Neptune::DBSubnetGroup":                                 &resources.AWSNeptuneDBSubnetGroup{},
		"AWS::OpsWorks::App":                                          &resources.AWSOpsWorksApp{},
		"AWS::OpsWorks::ElasticLoadBalancerAttachment":                &resources.AWSOpsWorksElasticLoadBalancerAttachment{},
		"AWS::OpsWorks::Instance":                                     &resources.AWSOpsWorksInstance{},
		"AWS::OpsWorks::Layer":                                        &resources.AWSOpsWorksLayer{},
		"AWS::OpsWorks::Stack":                                        &resources.AWSOpsWorksStack{},
		"AWS::OpsWorks::UserProfile":                                  &resources.AWSOpsWorksUserProfile{},
		"AWS::OpsWorks::Volume":                                       &resources.AWSOpsWorksVolume{},
		"AWS::OpsWorksCM::Server":                                     &resources.AWSOpsWorksCMServer{},
		"AWS::RAM::ResourceShare":                                     &resources.AWSRAMResourceShare{},
		"AWS::RDS::DBCluster":                                         &resources.AWSRDSDBCluster{},
		"AWS::RDS::DBClusterParameterGroup":                           &resources.AWSRDSDBClusterParameterGroup{},
		"AWS::RDS::DBInstance":                                        &resources.AWSRDSDBInstance{},
		"AWS::RDS::DBParameterGroup":                                  &resources.AWSRDSDBParameterGroup{},
		"AWS::RDS::DBSecurityGroup":                                   &resources.AWSRDSDBSecurityGroup{},
		"AWS::RDS::DBSecurityGroupIngress":                            &resources.AWSRDSDBSecurityGroupIngress{},
		"AWS::RDS::DBSubnetGroup":                                     &resources.AWSRDSDBSubnetGroup{},
		"AWS::RDS::EventSubscription":                                 &resources.AWSRDSEventSubscription{},
		"AWS::RDS::OptionGroup":                                       &resources.AWSRDSOptionGroup{},
		"AWS::Redshift::Cluster":                                      &resources.AWSRedshiftCluster{},
		"AWS::Redshift::ClusterParameterGroup":                        &resources.AWSRedshiftClusterParameterGroup{},
		"AWS::Redshift::ClusterSecurityGroup":                         &resources.AWSRedshiftClusterSecurityGroup{},
		"AWS::Redshift::ClusterSecurityGroupIngress":                  &resources.AWSRedshiftClusterSecurityGroupIngress{},
		"AWS::Redshift::ClusterSubnetGroup":                           &resources.AWSRedshiftClusterSubnetGroup{},
		"AWS::RoboMaker::Fleet":                                       &resources.AWSRoboMakerFleet{},
		"AWS::RoboMaker::Robot":                                       &resources.AWSRoboMakerRobot{},
		"AWS::RoboMaker::RobotApplication":                            &resources.AWSRoboMakerRobotApplication{},
		"AWS::RoboMaker::RobotApplicationVersion":                     &resources.AWSRoboMakerRobotApplicationVersion{},
		"AWS::RoboMaker::SimulationApplication":                       &resources.AWSRoboMakerSimulationApplication{},
		"AWS::RoboMaker::SimulationApplicationVersion":                &resources.AWSRoboMakerSimulationApplicationVersion{},
		"AWS::Route53::HealthCheck":                                   &resources.AWSRoute53HealthCheck{},
		"AWS::Route53::HostedZone":                                    &resources.AWSRoute53HostedZone{},
		"AWS::Route53::RecordSet":                                     &resources.AWSRoute53RecordSet{},
		"AWS::Route53::RecordSetGroup":                                &resources.AWSRoute53RecordSetGroup{},
		"AWS::Route53Resolver::ResolverEndpoint":                      &resources.AWSRoute53ResolverResolverEndpoint{},
		"AWS::Route53Resolver::ResolverRule":                          &resources.AWSRoute53ResolverResolverRule{},
		"AWS::Route53Resolver::ResolverRuleAssociation":               &resources.AWSRoute53ResolverResolverRuleAssociation{},
		"AWS::S3::Bucket":                                             &resources.AWSS3Bucket{},
		"AWS::S3::BucketPolicy":                                       &resources.AWSS3BucketPolicy{},
		"AWS::SDB::Domain":                                            &resources.AWSSDBDomain{},
		"AWS::SES::ConfigurationSet":                                  &resources.AWSSESConfigurationSet{},
		"AWS::SES::ConfigurationSetEventDestination":                  &resources.AWSSESConfigurationSetEventDestination{},
		"AWS::SES::ReceiptFilter":                                     &resources.AWSSESReceiptFilter{},
		"AWS::SES::ReceiptRule":                                       &resources.AWSSESReceiptRule{},
		"AWS::SES::ReceiptRuleSet":                                    &resources.AWSSESReceiptRuleSet{},
		"AWS::SES::Template":                                          &resources.AWSSESTemplate{},
		"AWS::SNS::Subscription":                                      &resources.AWSSNSSubscription{},
		"AWS::SNS::Topic":                                             &resources.AWSSNSTopic{},
		"AWS::SNS::TopicPolicy":                                       &resources.AWSSNSTopicPolicy{},
		"AWS::SQS::Queue":                                             &resources.AWSSQSQueue{},
		"AWS::SQS::QueuePolicy":                                       &resources.AWSSQSQueuePolicy{},
		"AWS::SSM::Association":                                       &resources.AWSSSMAssociation{},
		"AWS::SSM::Document":                                          &resources.AWSSSMDocument{},
		"AWS::SSM::MaintenanceWindow":                                 &resources.AWSSSMMaintenanceWindow{},
		"AWS::SSM::MaintenanceWindowTask":                             &resources.AWSSSMMaintenanceWindowTask{},
		"AWS::SSM::Parameter":                                         &resources.AWSSSMParameter{},
		"AWS::SSM::PatchBaseline":                                     &resources.AWSSSMPatchBaseline{},
		"AWS::SSM::ResourceDataSync":                                  &resources.AWSSSMResourceDataSync{},
		"AWS::SageMaker::Endpoint":                                    &resources.AWSSageMakerEndpoint{},
		"AWS::SageMaker::EndpointConfig":                              &resources.AWSSageMakerEndpointConfig{},
		"AWS::SageMaker::Model":                                       &resources.AWSSageMakerModel{},
		"AWS::SageMaker::NotebookInstance":                            &resources.AWSSageMakerNotebookInstance{},
		"AWS::SageMaker::NotebookInstanceLifecycleConfig":             &resources.AWSSageMakerNotebookInstanceLifecycleConfig{},
		"AWS::SecretsManager::ResourcePolicy":                         &resources.AWSSecretsManagerResourcePolicy{},
		"AWS::SecretsManager::RotationSchedule":                       &resources.AWSSecretsManagerRotationSchedule{},
		"AWS::SecretsManager::Secret":                                 &resources.AWSSecretsManagerSecret{},
		"AWS::SecretsManager::SecretTargetAttachment":                 &resources.AWSSecretsManagerSecretTargetAttachment{},
		"AWS::Serverless::Api":                                        &resources.AWSServerlessApi{},
		"AWS::Serverless::Application":                                &resources.AWSServerlessApplication{},
		"AWS::Serverless::Function":                                   &resources.AWSServerlessFunction{},
		"AWS::Serverless::LayerVersion":                               &resources.AWSServerlessLayerVersion{},
		"AWS::Serverless::SimpleTable":                                &resources.AWSServerlessSimpleTable{},
		"AWS::ServiceCatalog::AcceptedPortfolioShare":                 &resources.AWSServiceCatalogAcceptedPortfolioShare{},
		"AWS::ServiceCatalog::CloudFormationProduct":                  &resources.AWSServiceCatalogCloudFormationProduct{},
		"AWS::ServiceCatalog::CloudFormationProvisionedProduct":       &resources.AWSServiceCatalogCloudFormationProvisionedProduct{},
		"AWS::ServiceCatalog::LaunchNotificationConstraint":           &resources.AWSServiceCatalogLaunchNotificationConstraint{},
		"AWS::ServiceCatalog::LaunchRoleConstraint":                   &resources.AWSServiceCatalogLaunchRoleConstraint{},
		"AWS::ServiceCatalog::LaunchTemplateConstraint":               &resources.AWSServiceCatalogLaunchTemplateConstraint{},
		"AWS::ServiceCatalog::Portfolio":                              &resources.AWSServiceCatalogPortfolio{},
		"AWS::ServiceCatalog::PortfolioPrincipalAssociation":          &resources.AWSServiceCatalogPortfolioPrincipalAssociation{},
		"AWS::ServiceCatalog::PortfolioProductAssociation":            &resources.AWSServiceCatalogPortfolioProductAssociation{},
		"AWS::ServiceCatalog::PortfolioShare":                         &resources.AWSServiceCatalogPortfolioShare{},
		"AWS::ServiceCatalog::TagOption":                              &resources.AWSServiceCatalogTagOption{},
		"AWS::ServiceCatalog::TagOptionAssociation":                   &resources.AWSServiceCatalogTagOptionAssociation{},
		"AWS::ServiceDiscovery::HttpNamespace":                        &resources.AWSServiceDiscoveryHttpNamespace{},
		"AWS::ServiceDiscovery::Instance":                             &resources.AWSServiceDiscoveryInstance{},
		"AWS::ServiceDiscovery::PrivateDnsNamespace":                  &resources.AWSServiceDiscoveryPrivateDnsNamespace{},
		"AWS::ServiceDiscovery::PublicDnsNamespace":                   &resources.AWSServiceDiscoveryPublicDnsNamespace{},
		"AWS::ServiceDiscovery::Service":                              &resources.AWSServiceDiscoveryService{},
		"AWS::StepFunctions::Activity":                                &resources.AWSStepFunctionsActivity{},
		"AWS::StepFunctions::StateMachine":                            &resources.AWSStepFunctionsStateMachine{},
		"AWS::WAF::ByteMatchSet":                                      &resources.AWSWAFByteMatchSet{},
		"AWS::WAF::IPSet":                                             &resources.AWSWAFIPSet{},
		"AWS::WAF::Rule":                                              &resources.AWSWAFRule{},
		"AWS::WAF::SizeConstraintSet":                                 &resources.AWSWAFSizeConstraintSet{},
		"AWS::WAF::SqlInjectionMatchSet":                              &resources.AWSWAFSqlInjectionMatchSet{},
		"AWS::WAF::WebACL":                                            &resources.AWSWAFWebACL{},
		"AWS::WAF::XssMatchSet":                                       &resources.AWSWAFXssMatchSet{},
		"AWS::WAFRegional::ByteMatchSet":                              &resources.AWSWAFRegionalByteMatchSet{},
		"AWS::WAFRegional::IPSet":                                     &resources.AWSWAFRegionalIPSet{},
		"AWS::WAFRegional::Rule":                                      &resources.AWSWAFRegionalRule{},
		"AWS::WAFRegional::SizeConstraintSet":                         &resources.AWSWAFRegionalSizeConstraintSet{},
		"AWS::WAFRegional::SqlInjectionMatchSet":                      &resources.AWSWAFRegionalSqlInjectionMatchSet{},
		"AWS::WAFRegional::WebACL":                                    &resources.AWSWAFRegionalWebACL{},
		"AWS::WAFRegional::WebACLAssociation":                         &resources.AWSWAFRegionalWebACLAssociation{},
		"AWS::WAFRegional::XssMatchSet":                               &resources.AWSWAFRegionalXssMatchSet{},
		"AWS::WorkSpaces::Workspace":                                  &resources.AWSWorkSpacesWorkspace{},
		"Alexa::ASK::Skill":                                           &resources.AlexaASKSkill{},
	}
}

// GetAllAWSAmazonMQBrokerResources retrieves all AWSAmazonMQBroker items from an AWS CloudFormation template
func (t *Template) GetAllAWSAmazonMQBrokerResources() map[string]*resources.AWSAmazonMQBroker {
	results := map[string]*resources.AWSAmazonMQBroker{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSAmazonMQBroker:
			results[name] = resource
		}
	}
	return results
}

// GetAWSAmazonMQBrokerWithName retrieves all AWSAmazonMQBroker items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSAmazonMQBrokerWithName(name string) (*resources.AWSAmazonMQBroker, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSAmazonMQBroker:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSAmazonMQBroker not found", name)
}

// GetAllAWSAmazonMQConfigurationResources retrieves all AWSAmazonMQConfiguration items from an AWS CloudFormation template
func (t *Template) GetAllAWSAmazonMQConfigurationResources() map[string]*resources.AWSAmazonMQConfiguration {
	results := map[string]*resources.AWSAmazonMQConfiguration{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSAmazonMQConfiguration:
			results[name] = resource
		}
	}
	return results
}

// GetAWSAmazonMQConfigurationWithName retrieves all AWSAmazonMQConfiguration items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSAmazonMQConfigurationWithName(name string) (*resources.AWSAmazonMQConfiguration, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSAmazonMQConfiguration:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSAmazonMQConfiguration not found", name)
}

// GetAllAWSAmazonMQConfigurationAssociationResources retrieves all AWSAmazonMQConfigurationAssociation items from an AWS CloudFormation template
func (t *Template) GetAllAWSAmazonMQConfigurationAssociationResources() map[string]*resources.AWSAmazonMQConfigurationAssociation {
	results := map[string]*resources.AWSAmazonMQConfigurationAssociation{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSAmazonMQConfigurationAssociation:
			results[name] = resource
		}
	}
	return results
}

// GetAWSAmazonMQConfigurationAssociationWithName retrieves all AWSAmazonMQConfigurationAssociation items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSAmazonMQConfigurationAssociationWithName(name string) (*resources.AWSAmazonMQConfigurationAssociation, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSAmazonMQConfigurationAssociation:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSAmazonMQConfigurationAssociation not found", name)
}

// GetAllAWSApiGatewayAccountResources retrieves all AWSApiGatewayAccount items from an AWS CloudFormation template
func (t *Template) GetAllAWSApiGatewayAccountResources() map[string]*resources.AWSApiGatewayAccount {
	results := map[string]*resources.AWSApiGatewayAccount{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSApiGatewayAccount:
			results[name] = resource
		}
	}
	return results
}

// GetAWSApiGatewayAccountWithName retrieves all AWSApiGatewayAccount items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApiGatewayAccountWithName(name string) (*resources.AWSApiGatewayAccount, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSApiGatewayAccount:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSApiGatewayAccount not found", name)
}

// GetAllAWSApiGatewayApiKeyResources retrieves all AWSApiGatewayApiKey items from an AWS CloudFormation template
func (t *Template) GetAllAWSApiGatewayApiKeyResources() map[string]*resources.AWSApiGatewayApiKey {
	results := map[string]*resources.AWSApiGatewayApiKey{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSApiGatewayApiKey:
			results[name] = resource
		}
	}
	return results
}

// GetAWSApiGatewayApiKeyWithName retrieves all AWSApiGatewayApiKey items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApiGatewayApiKeyWithName(name string) (*resources.AWSApiGatewayApiKey, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSApiGatewayApiKey:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSApiGatewayApiKey not found", name)
}

// GetAllAWSApiGatewayAuthorizerResources retrieves all AWSApiGatewayAuthorizer items from an AWS CloudFormation template
func (t *Template) GetAllAWSApiGatewayAuthorizerResources() map[string]*resources.AWSApiGatewayAuthorizer {
	results := map[string]*resources.AWSApiGatewayAuthorizer{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSApiGatewayAuthorizer:
			results[name] = resource
		}
	}
	return results
}

// GetAWSApiGatewayAuthorizerWithName retrieves all AWSApiGatewayAuthorizer items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApiGatewayAuthorizerWithName(name string) (*resources.AWSApiGatewayAuthorizer, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSApiGatewayAuthorizer:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSApiGatewayAuthorizer not found", name)
}

// GetAllAWSApiGatewayBasePathMappingResources retrieves all AWSApiGatewayBasePathMapping items from an AWS CloudFormation template
func (t *Template) GetAllAWSApiGatewayBasePathMappingResources() map[string]*resources.AWSApiGatewayBasePathMapping {
	results := map[string]*resources.AWSApiGatewayBasePathMapping{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSApiGatewayBasePathMapping:
			results[name] = resource
		}
	}
	return results
}

// GetAWSApiGatewayBasePathMappingWithName retrieves all AWSApiGatewayBasePathMapping items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApiGatewayBasePathMappingWithName(name string) (*resources.AWSApiGatewayBasePathMapping, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSApiGatewayBasePathMapping:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSApiGatewayBasePathMapping not found", name)
}

// GetAllAWSApiGatewayClientCertificateResources retrieves all AWSApiGatewayClientCertificate items from an AWS CloudFormation template
func (t *Template) GetAllAWSApiGatewayClientCertificateResources() map[string]*resources.AWSApiGatewayClientCertificate {
	results := map[string]*resources.AWSApiGatewayClientCertificate{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSApiGatewayClientCertificate:
			results[name] = resource
		}
	}
	return results
}

// GetAWSApiGatewayClientCertificateWithName retrieves all AWSApiGatewayClientCertificate items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApiGatewayClientCertificateWithName(name string) (*resources.AWSApiGatewayClientCertificate, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSApiGatewayClientCertificate:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSApiGatewayClientCertificate not found", name)
}

// GetAllAWSApiGatewayDeploymentResources retrieves all AWSApiGatewayDeployment items from an AWS CloudFormation template
func (t *Template) GetAllAWSApiGatewayDeploymentResources() map[string]*resources.AWSApiGatewayDeployment {
	results := map[string]*resources.AWSApiGatewayDeployment{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSApiGatewayDeployment:
			results[name] = resource
		}
	}
	return results
}

// GetAWSApiGatewayDeploymentWithName retrieves all AWSApiGatewayDeployment items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApiGatewayDeploymentWithName(name string) (*resources.AWSApiGatewayDeployment, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSApiGatewayDeployment:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSApiGatewayDeployment not found", name)
}

// GetAllAWSApiGatewayDocumentationPartResources retrieves all AWSApiGatewayDocumentationPart items from an AWS CloudFormation template
func (t *Template) GetAllAWSApiGatewayDocumentationPartResources() map[string]*resources.AWSApiGatewayDocumentationPart {
	results := map[string]*resources.AWSApiGatewayDocumentationPart{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSApiGatewayDocumentationPart:
			results[name] = resource
		}
	}
	return results
}

// GetAWSApiGatewayDocumentationPartWithName retrieves all AWSApiGatewayDocumentationPart items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApiGatewayDocumentationPartWithName(name string) (*resources.AWSApiGatewayDocumentationPart, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSApiGatewayDocumentationPart:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSApiGatewayDocumentationPart not found", name)
}

// GetAllAWSApiGatewayDocumentationVersionResources retrieves all AWSApiGatewayDocumentationVersion items from an AWS CloudFormation template
func (t *Template) GetAllAWSApiGatewayDocumentationVersionResources() map[string]*resources.AWSApiGatewayDocumentationVersion {
	results := map[string]*resources.AWSApiGatewayDocumentationVersion{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSApiGatewayDocumentationVersion:
			results[name] = resource
		}
	}
	return results
}

// GetAWSApiGatewayDocumentationVersionWithName retrieves all AWSApiGatewayDocumentationVersion items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApiGatewayDocumentationVersionWithName(name string) (*resources.AWSApiGatewayDocumentationVersion, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSApiGatewayDocumentationVersion:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSApiGatewayDocumentationVersion not found", name)
}

// GetAllAWSApiGatewayDomainNameResources retrieves all AWSApiGatewayDomainName items from an AWS CloudFormation template
func (t *Template) GetAllAWSApiGatewayDomainNameResources() map[string]*resources.AWSApiGatewayDomainName {
	results := map[string]*resources.AWSApiGatewayDomainName{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSApiGatewayDomainName:
			results[name] = resource
		}
	}
	return results
}

// GetAWSApiGatewayDomainNameWithName retrieves all AWSApiGatewayDomainName items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApiGatewayDomainNameWithName(name string) (*resources.AWSApiGatewayDomainName, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSApiGatewayDomainName:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSApiGatewayDomainName not found", name)
}

// GetAllAWSApiGatewayGatewayResponseResources retrieves all AWSApiGatewayGatewayResponse items from an AWS CloudFormation template
func (t *Template) GetAllAWSApiGatewayGatewayResponseResources() map[string]*resources.AWSApiGatewayGatewayResponse {
	results := map[string]*resources.AWSApiGatewayGatewayResponse{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSApiGatewayGatewayResponse:
			results[name] = resource
		}
	}
	return results
}

// GetAWSApiGatewayGatewayResponseWithName retrieves all AWSApiGatewayGatewayResponse items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApiGatewayGatewayResponseWithName(name string) (*resources.AWSApiGatewayGatewayResponse, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSApiGatewayGatewayResponse:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSApiGatewayGatewayResponse not found", name)
}

// GetAllAWSApiGatewayMethodResources retrieves all AWSApiGatewayMethod items from an AWS CloudFormation template
func (t *Template) GetAllAWSApiGatewayMethodResources() map[string]*resources.AWSApiGatewayMethod {
	results := map[string]*resources.AWSApiGatewayMethod{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSApiGatewayMethod:
			results[name] = resource
		}
	}
	return results
}

// GetAWSApiGatewayMethodWithName retrieves all AWSApiGatewayMethod items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApiGatewayMethodWithName(name string) (*resources.AWSApiGatewayMethod, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSApiGatewayMethod:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSApiGatewayMethod not found", name)
}

// GetAllAWSApiGatewayModelResources retrieves all AWSApiGatewayModel items from an AWS CloudFormation template
func (t *Template) GetAllAWSApiGatewayModelResources() map[string]*resources.AWSApiGatewayModel {
	results := map[string]*resources.AWSApiGatewayModel{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSApiGatewayModel:
			results[name] = resource
		}
	}
	return results
}

// GetAWSApiGatewayModelWithName retrieves all AWSApiGatewayModel items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApiGatewayModelWithName(name string) (*resources.AWSApiGatewayModel, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSApiGatewayModel:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSApiGatewayModel not found", name)
}

// GetAllAWSApiGatewayRequestValidatorResources retrieves all AWSApiGatewayRequestValidator items from an AWS CloudFormation template
func (t *Template) GetAllAWSApiGatewayRequestValidatorResources() map[string]*resources.AWSApiGatewayRequestValidator {
	results := map[string]*resources.AWSApiGatewayRequestValidator{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSApiGatewayRequestValidator:
			results[name] = resource
		}
	}
	return results
}

// GetAWSApiGatewayRequestValidatorWithName retrieves all AWSApiGatewayRequestValidator items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApiGatewayRequestValidatorWithName(name string) (*resources.AWSApiGatewayRequestValidator, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSApiGatewayRequestValidator:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSApiGatewayRequestValidator not found", name)
}

// GetAllAWSApiGatewayResourceResources retrieves all AWSApiGatewayResource items from an AWS CloudFormation template
func (t *Template) GetAllAWSApiGatewayResourceResources() map[string]*resources.AWSApiGatewayResource {
	results := map[string]*resources.AWSApiGatewayResource{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSApiGatewayResource:
			results[name] = resource
		}
	}
	return results
}

// GetAWSApiGatewayResourceWithName retrieves all AWSApiGatewayResource items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApiGatewayResourceWithName(name string) (*resources.AWSApiGatewayResource, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSApiGatewayResource:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSApiGatewayResource not found", name)
}

// GetAllAWSApiGatewayRestApiResources retrieves all AWSApiGatewayRestApi items from an AWS CloudFormation template
func (t *Template) GetAllAWSApiGatewayRestApiResources() map[string]*resources.AWSApiGatewayRestApi {
	results := map[string]*resources.AWSApiGatewayRestApi{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSApiGatewayRestApi:
			results[name] = resource
		}
	}
	return results
}

// GetAWSApiGatewayRestApiWithName retrieves all AWSApiGatewayRestApi items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApiGatewayRestApiWithName(name string) (*resources.AWSApiGatewayRestApi, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSApiGatewayRestApi:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSApiGatewayRestApi not found", name)
}

// GetAllAWSApiGatewayStageResources retrieves all AWSApiGatewayStage items from an AWS CloudFormation template
func (t *Template) GetAllAWSApiGatewayStageResources() map[string]*resources.AWSApiGatewayStage {
	results := map[string]*resources.AWSApiGatewayStage{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSApiGatewayStage:
			results[name] = resource
		}
	}
	return results
}

// GetAWSApiGatewayStageWithName retrieves all AWSApiGatewayStage items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApiGatewayStageWithName(name string) (*resources.AWSApiGatewayStage, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSApiGatewayStage:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSApiGatewayStage not found", name)
}

// GetAllAWSApiGatewayUsagePlanResources retrieves all AWSApiGatewayUsagePlan items from an AWS CloudFormation template
func (t *Template) GetAllAWSApiGatewayUsagePlanResources() map[string]*resources.AWSApiGatewayUsagePlan {
	results := map[string]*resources.AWSApiGatewayUsagePlan{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSApiGatewayUsagePlan:
			results[name] = resource
		}
	}
	return results
}

// GetAWSApiGatewayUsagePlanWithName retrieves all AWSApiGatewayUsagePlan items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApiGatewayUsagePlanWithName(name string) (*resources.AWSApiGatewayUsagePlan, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSApiGatewayUsagePlan:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSApiGatewayUsagePlan not found", name)
}

// GetAllAWSApiGatewayUsagePlanKeyResources retrieves all AWSApiGatewayUsagePlanKey items from an AWS CloudFormation template
func (t *Template) GetAllAWSApiGatewayUsagePlanKeyResources() map[string]*resources.AWSApiGatewayUsagePlanKey {
	results := map[string]*resources.AWSApiGatewayUsagePlanKey{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSApiGatewayUsagePlanKey:
			results[name] = resource
		}
	}
	return results
}

// GetAWSApiGatewayUsagePlanKeyWithName retrieves all AWSApiGatewayUsagePlanKey items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApiGatewayUsagePlanKeyWithName(name string) (*resources.AWSApiGatewayUsagePlanKey, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSApiGatewayUsagePlanKey:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSApiGatewayUsagePlanKey not found", name)
}

// GetAllAWSApiGatewayVpcLinkResources retrieves all AWSApiGatewayVpcLink items from an AWS CloudFormation template
func (t *Template) GetAllAWSApiGatewayVpcLinkResources() map[string]*resources.AWSApiGatewayVpcLink {
	results := map[string]*resources.AWSApiGatewayVpcLink{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSApiGatewayVpcLink:
			results[name] = resource
		}
	}
	return results
}

// GetAWSApiGatewayVpcLinkWithName retrieves all AWSApiGatewayVpcLink items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApiGatewayVpcLinkWithName(name string) (*resources.AWSApiGatewayVpcLink, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSApiGatewayVpcLink:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSApiGatewayVpcLink not found", name)
}

// GetAllAWSApiGatewayV2ApiResources retrieves all AWSApiGatewayV2Api items from an AWS CloudFormation template
func (t *Template) GetAllAWSApiGatewayV2ApiResources() map[string]*resources.AWSApiGatewayV2Api {
	results := map[string]*resources.AWSApiGatewayV2Api{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSApiGatewayV2Api:
			results[name] = resource
		}
	}
	return results
}

// GetAWSApiGatewayV2ApiWithName retrieves all AWSApiGatewayV2Api items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApiGatewayV2ApiWithName(name string) (*resources.AWSApiGatewayV2Api, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSApiGatewayV2Api:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSApiGatewayV2Api not found", name)
}

// GetAllAWSApiGatewayV2AuthorizerResources retrieves all AWSApiGatewayV2Authorizer items from an AWS CloudFormation template
func (t *Template) GetAllAWSApiGatewayV2AuthorizerResources() map[string]*resources.AWSApiGatewayV2Authorizer {
	results := map[string]*resources.AWSApiGatewayV2Authorizer{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSApiGatewayV2Authorizer:
			results[name] = resource
		}
	}
	return results
}

// GetAWSApiGatewayV2AuthorizerWithName retrieves all AWSApiGatewayV2Authorizer items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApiGatewayV2AuthorizerWithName(name string) (*resources.AWSApiGatewayV2Authorizer, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSApiGatewayV2Authorizer:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSApiGatewayV2Authorizer not found", name)
}

// GetAllAWSApiGatewayV2DeploymentResources retrieves all AWSApiGatewayV2Deployment items from an AWS CloudFormation template
func (t *Template) GetAllAWSApiGatewayV2DeploymentResources() map[string]*resources.AWSApiGatewayV2Deployment {
	results := map[string]*resources.AWSApiGatewayV2Deployment{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSApiGatewayV2Deployment:
			results[name] = resource
		}
	}
	return results
}

// GetAWSApiGatewayV2DeploymentWithName retrieves all AWSApiGatewayV2Deployment items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApiGatewayV2DeploymentWithName(name string) (*resources.AWSApiGatewayV2Deployment, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSApiGatewayV2Deployment:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSApiGatewayV2Deployment not found", name)
}

// GetAllAWSApiGatewayV2IntegrationResources retrieves all AWSApiGatewayV2Integration items from an AWS CloudFormation template
func (t *Template) GetAllAWSApiGatewayV2IntegrationResources() map[string]*resources.AWSApiGatewayV2Integration {
	results := map[string]*resources.AWSApiGatewayV2Integration{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSApiGatewayV2Integration:
			results[name] = resource
		}
	}
	return results
}

// GetAWSApiGatewayV2IntegrationWithName retrieves all AWSApiGatewayV2Integration items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApiGatewayV2IntegrationWithName(name string) (*resources.AWSApiGatewayV2Integration, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSApiGatewayV2Integration:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSApiGatewayV2Integration not found", name)
}

// GetAllAWSApiGatewayV2IntegrationResponseResources retrieves all AWSApiGatewayV2IntegrationResponse items from an AWS CloudFormation template
func (t *Template) GetAllAWSApiGatewayV2IntegrationResponseResources() map[string]*resources.AWSApiGatewayV2IntegrationResponse {
	results := map[string]*resources.AWSApiGatewayV2IntegrationResponse{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSApiGatewayV2IntegrationResponse:
			results[name] = resource
		}
	}
	return results
}

// GetAWSApiGatewayV2IntegrationResponseWithName retrieves all AWSApiGatewayV2IntegrationResponse items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApiGatewayV2IntegrationResponseWithName(name string) (*resources.AWSApiGatewayV2IntegrationResponse, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSApiGatewayV2IntegrationResponse:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSApiGatewayV2IntegrationResponse not found", name)
}

// GetAllAWSApiGatewayV2ModelResources retrieves all AWSApiGatewayV2Model items from an AWS CloudFormation template
func (t *Template) GetAllAWSApiGatewayV2ModelResources() map[string]*resources.AWSApiGatewayV2Model {
	results := map[string]*resources.AWSApiGatewayV2Model{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSApiGatewayV2Model:
			results[name] = resource
		}
	}
	return results
}

// GetAWSApiGatewayV2ModelWithName retrieves all AWSApiGatewayV2Model items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApiGatewayV2ModelWithName(name string) (*resources.AWSApiGatewayV2Model, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSApiGatewayV2Model:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSApiGatewayV2Model not found", name)
}

// GetAllAWSApiGatewayV2RouteResources retrieves all AWSApiGatewayV2Route items from an AWS CloudFormation template
func (t *Template) GetAllAWSApiGatewayV2RouteResources() map[string]*resources.AWSApiGatewayV2Route {
	results := map[string]*resources.AWSApiGatewayV2Route{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSApiGatewayV2Route:
			results[name] = resource
		}
	}
	return results
}

// GetAWSApiGatewayV2RouteWithName retrieves all AWSApiGatewayV2Route items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApiGatewayV2RouteWithName(name string) (*resources.AWSApiGatewayV2Route, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSApiGatewayV2Route:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSApiGatewayV2Route not found", name)
}

// GetAllAWSApiGatewayV2RouteResponseResources retrieves all AWSApiGatewayV2RouteResponse items from an AWS CloudFormation template
func (t *Template) GetAllAWSApiGatewayV2RouteResponseResources() map[string]*resources.AWSApiGatewayV2RouteResponse {
	results := map[string]*resources.AWSApiGatewayV2RouteResponse{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSApiGatewayV2RouteResponse:
			results[name] = resource
		}
	}
	return results
}

// GetAWSApiGatewayV2RouteResponseWithName retrieves all AWSApiGatewayV2RouteResponse items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApiGatewayV2RouteResponseWithName(name string) (*resources.AWSApiGatewayV2RouteResponse, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSApiGatewayV2RouteResponse:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSApiGatewayV2RouteResponse not found", name)
}

// GetAllAWSApiGatewayV2StageResources retrieves all AWSApiGatewayV2Stage items from an AWS CloudFormation template
func (t *Template) GetAllAWSApiGatewayV2StageResources() map[string]*resources.AWSApiGatewayV2Stage {
	results := map[string]*resources.AWSApiGatewayV2Stage{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSApiGatewayV2Stage:
			results[name] = resource
		}
	}
	return results
}

// GetAWSApiGatewayV2StageWithName retrieves all AWSApiGatewayV2Stage items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApiGatewayV2StageWithName(name string) (*resources.AWSApiGatewayV2Stage, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSApiGatewayV2Stage:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSApiGatewayV2Stage not found", name)
}

// GetAllAWSAppStreamDirectoryConfigResources retrieves all AWSAppStreamDirectoryConfig items from an AWS CloudFormation template
func (t *Template) GetAllAWSAppStreamDirectoryConfigResources() map[string]*resources.AWSAppStreamDirectoryConfig {
	results := map[string]*resources.AWSAppStreamDirectoryConfig{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSAppStreamDirectoryConfig:
			results[name] = resource
		}
	}
	return results
}

// GetAWSAppStreamDirectoryConfigWithName retrieves all AWSAppStreamDirectoryConfig items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSAppStreamDirectoryConfigWithName(name string) (*resources.AWSAppStreamDirectoryConfig, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSAppStreamDirectoryConfig:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSAppStreamDirectoryConfig not found", name)
}

// GetAllAWSAppStreamFleetResources retrieves all AWSAppStreamFleet items from an AWS CloudFormation template
func (t *Template) GetAllAWSAppStreamFleetResources() map[string]*resources.AWSAppStreamFleet {
	results := map[string]*resources.AWSAppStreamFleet{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSAppStreamFleet:
			results[name] = resource
		}
	}
	return results
}

// GetAWSAppStreamFleetWithName retrieves all AWSAppStreamFleet items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSAppStreamFleetWithName(name string) (*resources.AWSAppStreamFleet, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSAppStreamFleet:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSAppStreamFleet not found", name)
}

// GetAllAWSAppStreamImageBuilderResources retrieves all AWSAppStreamImageBuilder items from an AWS CloudFormation template
func (t *Template) GetAllAWSAppStreamImageBuilderResources() map[string]*resources.AWSAppStreamImageBuilder {
	results := map[string]*resources.AWSAppStreamImageBuilder{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSAppStreamImageBuilder:
			results[name] = resource
		}
	}
	return results
}

// GetAWSAppStreamImageBuilderWithName retrieves all AWSAppStreamImageBuilder items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSAppStreamImageBuilderWithName(name string) (*resources.AWSAppStreamImageBuilder, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSAppStreamImageBuilder:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSAppStreamImageBuilder not found", name)
}

// GetAllAWSAppStreamStackResources retrieves all AWSAppStreamStack items from an AWS CloudFormation template
func (t *Template) GetAllAWSAppStreamStackResources() map[string]*resources.AWSAppStreamStack {
	results := map[string]*resources.AWSAppStreamStack{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSAppStreamStack:
			results[name] = resource
		}
	}
	return results
}

// GetAWSAppStreamStackWithName retrieves all AWSAppStreamStack items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSAppStreamStackWithName(name string) (*resources.AWSAppStreamStack, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSAppStreamStack:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSAppStreamStack not found", name)
}

// GetAllAWSAppStreamStackFleetAssociationResources retrieves all AWSAppStreamStackFleetAssociation items from an AWS CloudFormation template
func (t *Template) GetAllAWSAppStreamStackFleetAssociationResources() map[string]*resources.AWSAppStreamStackFleetAssociation {
	results := map[string]*resources.AWSAppStreamStackFleetAssociation{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSAppStreamStackFleetAssociation:
			results[name] = resource
		}
	}
	return results
}

// GetAWSAppStreamStackFleetAssociationWithName retrieves all AWSAppStreamStackFleetAssociation items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSAppStreamStackFleetAssociationWithName(name string) (*resources.AWSAppStreamStackFleetAssociation, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSAppStreamStackFleetAssociation:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSAppStreamStackFleetAssociation not found", name)
}

// GetAllAWSAppStreamStackUserAssociationResources retrieves all AWSAppStreamStackUserAssociation items from an AWS CloudFormation template
func (t *Template) GetAllAWSAppStreamStackUserAssociationResources() map[string]*resources.AWSAppStreamStackUserAssociation {
	results := map[string]*resources.AWSAppStreamStackUserAssociation{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSAppStreamStackUserAssociation:
			results[name] = resource
		}
	}
	return results
}

// GetAWSAppStreamStackUserAssociationWithName retrieves all AWSAppStreamStackUserAssociation items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSAppStreamStackUserAssociationWithName(name string) (*resources.AWSAppStreamStackUserAssociation, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSAppStreamStackUserAssociation:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSAppStreamStackUserAssociation not found", name)
}

// GetAllAWSAppStreamUserResources retrieves all AWSAppStreamUser items from an AWS CloudFormation template
func (t *Template) GetAllAWSAppStreamUserResources() map[string]*resources.AWSAppStreamUser {
	results := map[string]*resources.AWSAppStreamUser{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSAppStreamUser:
			results[name] = resource
		}
	}
	return results
}

// GetAWSAppStreamUserWithName retrieves all AWSAppStreamUser items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSAppStreamUserWithName(name string) (*resources.AWSAppStreamUser, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSAppStreamUser:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSAppStreamUser not found", name)
}

// GetAllAWSAppSyncApiKeyResources retrieves all AWSAppSyncApiKey items from an AWS CloudFormation template
func (t *Template) GetAllAWSAppSyncApiKeyResources() map[string]*resources.AWSAppSyncApiKey {
	results := map[string]*resources.AWSAppSyncApiKey{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSAppSyncApiKey:
			results[name] = resource
		}
	}
	return results
}

// GetAWSAppSyncApiKeyWithName retrieves all AWSAppSyncApiKey items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSAppSyncApiKeyWithName(name string) (*resources.AWSAppSyncApiKey, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSAppSyncApiKey:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSAppSyncApiKey not found", name)
}

// GetAllAWSAppSyncDataSourceResources retrieves all AWSAppSyncDataSource items from an AWS CloudFormation template
func (t *Template) GetAllAWSAppSyncDataSourceResources() map[string]*resources.AWSAppSyncDataSource {
	results := map[string]*resources.AWSAppSyncDataSource{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSAppSyncDataSource:
			results[name] = resource
		}
	}
	return results
}

// GetAWSAppSyncDataSourceWithName retrieves all AWSAppSyncDataSource items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSAppSyncDataSourceWithName(name string) (*resources.AWSAppSyncDataSource, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSAppSyncDataSource:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSAppSyncDataSource not found", name)
}

// GetAllAWSAppSyncFunctionConfigurationResources retrieves all AWSAppSyncFunctionConfiguration items from an AWS CloudFormation template
func (t *Template) GetAllAWSAppSyncFunctionConfigurationResources() map[string]*resources.AWSAppSyncFunctionConfiguration {
	results := map[string]*resources.AWSAppSyncFunctionConfiguration{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSAppSyncFunctionConfiguration:
			results[name] = resource
		}
	}
	return results
}

// GetAWSAppSyncFunctionConfigurationWithName retrieves all AWSAppSyncFunctionConfiguration items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSAppSyncFunctionConfigurationWithName(name string) (*resources.AWSAppSyncFunctionConfiguration, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSAppSyncFunctionConfiguration:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSAppSyncFunctionConfiguration not found", name)
}

// GetAllAWSAppSyncGraphQLApiResources retrieves all AWSAppSyncGraphQLApi items from an AWS CloudFormation template
func (t *Template) GetAllAWSAppSyncGraphQLApiResources() map[string]*resources.AWSAppSyncGraphQLApi {
	results := map[string]*resources.AWSAppSyncGraphQLApi{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSAppSyncGraphQLApi:
			results[name] = resource
		}
	}
	return results
}

// GetAWSAppSyncGraphQLApiWithName retrieves all AWSAppSyncGraphQLApi items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSAppSyncGraphQLApiWithName(name string) (*resources.AWSAppSyncGraphQLApi, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSAppSyncGraphQLApi:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSAppSyncGraphQLApi not found", name)
}

// GetAllAWSAppSyncGraphQLSchemaResources retrieves all AWSAppSyncGraphQLSchema items from an AWS CloudFormation template
func (t *Template) GetAllAWSAppSyncGraphQLSchemaResources() map[string]*resources.AWSAppSyncGraphQLSchema {
	results := map[string]*resources.AWSAppSyncGraphQLSchema{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSAppSyncGraphQLSchema:
			results[name] = resource
		}
	}
	return results
}

// GetAWSAppSyncGraphQLSchemaWithName retrieves all AWSAppSyncGraphQLSchema items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSAppSyncGraphQLSchemaWithName(name string) (*resources.AWSAppSyncGraphQLSchema, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSAppSyncGraphQLSchema:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSAppSyncGraphQLSchema not found", name)
}

// GetAllAWSAppSyncResolverResources retrieves all AWSAppSyncResolver items from an AWS CloudFormation template
func (t *Template) GetAllAWSAppSyncResolverResources() map[string]*resources.AWSAppSyncResolver {
	results := map[string]*resources.AWSAppSyncResolver{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSAppSyncResolver:
			results[name] = resource
		}
	}
	return results
}

// GetAWSAppSyncResolverWithName retrieves all AWSAppSyncResolver items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSAppSyncResolverWithName(name string) (*resources.AWSAppSyncResolver, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSAppSyncResolver:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSAppSyncResolver not found", name)
}

// GetAllAWSApplicationAutoScalingScalableTargetResources retrieves all AWSApplicationAutoScalingScalableTarget items from an AWS CloudFormation template
func (t *Template) GetAllAWSApplicationAutoScalingScalableTargetResources() map[string]*resources.AWSApplicationAutoScalingScalableTarget {
	results := map[string]*resources.AWSApplicationAutoScalingScalableTarget{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSApplicationAutoScalingScalableTarget:
			results[name] = resource
		}
	}
	return results
}

// GetAWSApplicationAutoScalingScalableTargetWithName retrieves all AWSApplicationAutoScalingScalableTarget items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApplicationAutoScalingScalableTargetWithName(name string) (*resources.AWSApplicationAutoScalingScalableTarget, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSApplicationAutoScalingScalableTarget:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSApplicationAutoScalingScalableTarget not found", name)
}

// GetAllAWSApplicationAutoScalingScalingPolicyResources retrieves all AWSApplicationAutoScalingScalingPolicy items from an AWS CloudFormation template
func (t *Template) GetAllAWSApplicationAutoScalingScalingPolicyResources() map[string]*resources.AWSApplicationAutoScalingScalingPolicy {
	results := map[string]*resources.AWSApplicationAutoScalingScalingPolicy{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSApplicationAutoScalingScalingPolicy:
			results[name] = resource
		}
	}
	return results
}

// GetAWSApplicationAutoScalingScalingPolicyWithName retrieves all AWSApplicationAutoScalingScalingPolicy items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApplicationAutoScalingScalingPolicyWithName(name string) (*resources.AWSApplicationAutoScalingScalingPolicy, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSApplicationAutoScalingScalingPolicy:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSApplicationAutoScalingScalingPolicy not found", name)
}

// GetAllAWSAthenaNamedQueryResources retrieves all AWSAthenaNamedQuery items from an AWS CloudFormation template
func (t *Template) GetAllAWSAthenaNamedQueryResources() map[string]*resources.AWSAthenaNamedQuery {
	results := map[string]*resources.AWSAthenaNamedQuery{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSAthenaNamedQuery:
			results[name] = resource
		}
	}
	return results
}

// GetAWSAthenaNamedQueryWithName retrieves all AWSAthenaNamedQuery items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSAthenaNamedQueryWithName(name string) (*resources.AWSAthenaNamedQuery, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSAthenaNamedQuery:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSAthenaNamedQuery not found", name)
}

// GetAllAWSAutoScalingAutoScalingGroupResources retrieves all AWSAutoScalingAutoScalingGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSAutoScalingAutoScalingGroupResources() map[string]*resources.AWSAutoScalingAutoScalingGroup {
	results := map[string]*resources.AWSAutoScalingAutoScalingGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSAutoScalingAutoScalingGroup:
			results[name] = resource
		}
	}
	return results
}

// GetAWSAutoScalingAutoScalingGroupWithName retrieves all AWSAutoScalingAutoScalingGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSAutoScalingAutoScalingGroupWithName(name string) (*resources.AWSAutoScalingAutoScalingGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSAutoScalingAutoScalingGroup:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSAutoScalingAutoScalingGroup not found", name)
}

// GetAllAWSAutoScalingLaunchConfigurationResources retrieves all AWSAutoScalingLaunchConfiguration items from an AWS CloudFormation template
func (t *Template) GetAllAWSAutoScalingLaunchConfigurationResources() map[string]*resources.AWSAutoScalingLaunchConfiguration {
	results := map[string]*resources.AWSAutoScalingLaunchConfiguration{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSAutoScalingLaunchConfiguration:
			results[name] = resource
		}
	}
	return results
}

// GetAWSAutoScalingLaunchConfigurationWithName retrieves all AWSAutoScalingLaunchConfiguration items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSAutoScalingLaunchConfigurationWithName(name string) (*resources.AWSAutoScalingLaunchConfiguration, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSAutoScalingLaunchConfiguration:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSAutoScalingLaunchConfiguration not found", name)
}

// GetAllAWSAutoScalingLifecycleHookResources retrieves all AWSAutoScalingLifecycleHook items from an AWS CloudFormation template
func (t *Template) GetAllAWSAutoScalingLifecycleHookResources() map[string]*resources.AWSAutoScalingLifecycleHook {
	results := map[string]*resources.AWSAutoScalingLifecycleHook{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSAutoScalingLifecycleHook:
			results[name] = resource
		}
	}
	return results
}

// GetAWSAutoScalingLifecycleHookWithName retrieves all AWSAutoScalingLifecycleHook items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSAutoScalingLifecycleHookWithName(name string) (*resources.AWSAutoScalingLifecycleHook, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSAutoScalingLifecycleHook:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSAutoScalingLifecycleHook not found", name)
}

// GetAllAWSAutoScalingScalingPolicyResources retrieves all AWSAutoScalingScalingPolicy items from an AWS CloudFormation template
func (t *Template) GetAllAWSAutoScalingScalingPolicyResources() map[string]*resources.AWSAutoScalingScalingPolicy {
	results := map[string]*resources.AWSAutoScalingScalingPolicy{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSAutoScalingScalingPolicy:
			results[name] = resource
		}
	}
	return results
}

// GetAWSAutoScalingScalingPolicyWithName retrieves all AWSAutoScalingScalingPolicy items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSAutoScalingScalingPolicyWithName(name string) (*resources.AWSAutoScalingScalingPolicy, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSAutoScalingScalingPolicy:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSAutoScalingScalingPolicy not found", name)
}

// GetAllAWSAutoScalingScheduledActionResources retrieves all AWSAutoScalingScheduledAction items from an AWS CloudFormation template
func (t *Template) GetAllAWSAutoScalingScheduledActionResources() map[string]*resources.AWSAutoScalingScheduledAction {
	results := map[string]*resources.AWSAutoScalingScheduledAction{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSAutoScalingScheduledAction:
			results[name] = resource
		}
	}
	return results
}

// GetAWSAutoScalingScheduledActionWithName retrieves all AWSAutoScalingScheduledAction items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSAutoScalingScheduledActionWithName(name string) (*resources.AWSAutoScalingScheduledAction, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSAutoScalingScheduledAction:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSAutoScalingScheduledAction not found", name)
}

// GetAllAWSAutoScalingPlansScalingPlanResources retrieves all AWSAutoScalingPlansScalingPlan items from an AWS CloudFormation template
func (t *Template) GetAllAWSAutoScalingPlansScalingPlanResources() map[string]*resources.AWSAutoScalingPlansScalingPlan {
	results := map[string]*resources.AWSAutoScalingPlansScalingPlan{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSAutoScalingPlansScalingPlan:
			results[name] = resource
		}
	}
	return results
}

// GetAWSAutoScalingPlansScalingPlanWithName retrieves all AWSAutoScalingPlansScalingPlan items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSAutoScalingPlansScalingPlanWithName(name string) (*resources.AWSAutoScalingPlansScalingPlan, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSAutoScalingPlansScalingPlan:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSAutoScalingPlansScalingPlan not found", name)
}

// GetAllAWSBatchComputeEnvironmentResources retrieves all AWSBatchComputeEnvironment items from an AWS CloudFormation template
func (t *Template) GetAllAWSBatchComputeEnvironmentResources() map[string]*resources.AWSBatchComputeEnvironment {
	results := map[string]*resources.AWSBatchComputeEnvironment{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSBatchComputeEnvironment:
			results[name] = resource
		}
	}
	return results
}

// GetAWSBatchComputeEnvironmentWithName retrieves all AWSBatchComputeEnvironment items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSBatchComputeEnvironmentWithName(name string) (*resources.AWSBatchComputeEnvironment, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSBatchComputeEnvironment:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSBatchComputeEnvironment not found", name)
}

// GetAllAWSBatchJobDefinitionResources retrieves all AWSBatchJobDefinition items from an AWS CloudFormation template
func (t *Template) GetAllAWSBatchJobDefinitionResources() map[string]*resources.AWSBatchJobDefinition {
	results := map[string]*resources.AWSBatchJobDefinition{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSBatchJobDefinition:
			results[name] = resource
		}
	}
	return results
}

// GetAWSBatchJobDefinitionWithName retrieves all AWSBatchJobDefinition items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSBatchJobDefinitionWithName(name string) (*resources.AWSBatchJobDefinition, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSBatchJobDefinition:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSBatchJobDefinition not found", name)
}

// GetAllAWSBatchJobQueueResources retrieves all AWSBatchJobQueue items from an AWS CloudFormation template
func (t *Template) GetAllAWSBatchJobQueueResources() map[string]*resources.AWSBatchJobQueue {
	results := map[string]*resources.AWSBatchJobQueue{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSBatchJobQueue:
			results[name] = resource
		}
	}
	return results
}

// GetAWSBatchJobQueueWithName retrieves all AWSBatchJobQueue items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSBatchJobQueueWithName(name string) (*resources.AWSBatchJobQueue, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSBatchJobQueue:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSBatchJobQueue not found", name)
}

// GetAllAWSBudgetsBudgetResources retrieves all AWSBudgetsBudget items from an AWS CloudFormation template
func (t *Template) GetAllAWSBudgetsBudgetResources() map[string]*resources.AWSBudgetsBudget {
	results := map[string]*resources.AWSBudgetsBudget{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSBudgetsBudget:
			results[name] = resource
		}
	}
	return results
}

// GetAWSBudgetsBudgetWithName retrieves all AWSBudgetsBudget items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSBudgetsBudgetWithName(name string) (*resources.AWSBudgetsBudget, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSBudgetsBudget:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSBudgetsBudget not found", name)
}

// GetAllAWSCertificateManagerCertificateResources retrieves all AWSCertificateManagerCertificate items from an AWS CloudFormation template
func (t *Template) GetAllAWSCertificateManagerCertificateResources() map[string]*resources.AWSCertificateManagerCertificate {
	results := map[string]*resources.AWSCertificateManagerCertificate{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSCertificateManagerCertificate:
			results[name] = resource
		}
	}
	return results
}

// GetAWSCertificateManagerCertificateWithName retrieves all AWSCertificateManagerCertificate items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCertificateManagerCertificateWithName(name string) (*resources.AWSCertificateManagerCertificate, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSCertificateManagerCertificate:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSCertificateManagerCertificate not found", name)
}

// GetAllAWSCloud9EnvironmentEC2Resources retrieves all AWSCloud9EnvironmentEC2 items from an AWS CloudFormation template
func (t *Template) GetAllAWSCloud9EnvironmentEC2Resources() map[string]*resources.AWSCloud9EnvironmentEC2 {
	results := map[string]*resources.AWSCloud9EnvironmentEC2{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSCloud9EnvironmentEC2:
			results[name] = resource
		}
	}
	return results
}

// GetAWSCloud9EnvironmentEC2WithName retrieves all AWSCloud9EnvironmentEC2 items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCloud9EnvironmentEC2WithName(name string) (*resources.AWSCloud9EnvironmentEC2, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSCloud9EnvironmentEC2:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSCloud9EnvironmentEC2 not found", name)
}

// GetAllAWSCloudFormationCustomResourceResources retrieves all AWSCloudFormationCustomResource items from an AWS CloudFormation template
func (t *Template) GetAllAWSCloudFormationCustomResourceResources() map[string]*resources.AWSCloudFormationCustomResource {
	results := map[string]*resources.AWSCloudFormationCustomResource{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSCloudFormationCustomResource:
			results[name] = resource
		}
	}
	return results
}

// GetAWSCloudFormationCustomResourceWithName retrieves all AWSCloudFormationCustomResource items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCloudFormationCustomResourceWithName(name string) (*resources.AWSCloudFormationCustomResource, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSCloudFormationCustomResource:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSCloudFormationCustomResource not found", name)
}

// GetAllAWSCloudFormationMacroResources retrieves all AWSCloudFormationMacro items from an AWS CloudFormation template
func (t *Template) GetAllAWSCloudFormationMacroResources() map[string]*resources.AWSCloudFormationMacro {
	results := map[string]*resources.AWSCloudFormationMacro{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSCloudFormationMacro:
			results[name] = resource
		}
	}
	return results
}

// GetAWSCloudFormationMacroWithName retrieves all AWSCloudFormationMacro items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCloudFormationMacroWithName(name string) (*resources.AWSCloudFormationMacro, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSCloudFormationMacro:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSCloudFormationMacro not found", name)
}

// GetAllAWSCloudFormationStackResources retrieves all AWSCloudFormationStack items from an AWS CloudFormation template
func (t *Template) GetAllAWSCloudFormationStackResources() map[string]*resources.AWSCloudFormationStack {
	results := map[string]*resources.AWSCloudFormationStack{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSCloudFormationStack:
			results[name] = resource
		}
	}
	return results
}

// GetAWSCloudFormationStackWithName retrieves all AWSCloudFormationStack items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCloudFormationStackWithName(name string) (*resources.AWSCloudFormationStack, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSCloudFormationStack:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSCloudFormationStack not found", name)
}

// GetAllAWSCloudFormationWaitConditionResources retrieves all AWSCloudFormationWaitCondition items from an AWS CloudFormation template
func (t *Template) GetAllAWSCloudFormationWaitConditionResources() map[string]*resources.AWSCloudFormationWaitCondition {
	results := map[string]*resources.AWSCloudFormationWaitCondition{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSCloudFormationWaitCondition:
			results[name] = resource
		}
	}
	return results
}

// GetAWSCloudFormationWaitConditionWithName retrieves all AWSCloudFormationWaitCondition items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCloudFormationWaitConditionWithName(name string) (*resources.AWSCloudFormationWaitCondition, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSCloudFormationWaitCondition:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSCloudFormationWaitCondition not found", name)
}

// GetAllAWSCloudFormationWaitConditionHandleResources retrieves all AWSCloudFormationWaitConditionHandle items from an AWS CloudFormation template
func (t *Template) GetAllAWSCloudFormationWaitConditionHandleResources() map[string]*resources.AWSCloudFormationWaitConditionHandle {
	results := map[string]*resources.AWSCloudFormationWaitConditionHandle{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSCloudFormationWaitConditionHandle:
			results[name] = resource
		}
	}
	return results
}

// GetAWSCloudFormationWaitConditionHandleWithName retrieves all AWSCloudFormationWaitConditionHandle items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCloudFormationWaitConditionHandleWithName(name string) (*resources.AWSCloudFormationWaitConditionHandle, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSCloudFormationWaitConditionHandle:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSCloudFormationWaitConditionHandle not found", name)
}

// GetAllAWSCloudFrontCloudFrontOriginAccessIdentityResources retrieves all AWSCloudFrontCloudFrontOriginAccessIdentity items from an AWS CloudFormation template
func (t *Template) GetAllAWSCloudFrontCloudFrontOriginAccessIdentityResources() map[string]*resources.AWSCloudFrontCloudFrontOriginAccessIdentity {
	results := map[string]*resources.AWSCloudFrontCloudFrontOriginAccessIdentity{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSCloudFrontCloudFrontOriginAccessIdentity:
			results[name] = resource
		}
	}
	return results
}

// GetAWSCloudFrontCloudFrontOriginAccessIdentityWithName retrieves all AWSCloudFrontCloudFrontOriginAccessIdentity items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCloudFrontCloudFrontOriginAccessIdentityWithName(name string) (*resources.AWSCloudFrontCloudFrontOriginAccessIdentity, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSCloudFrontCloudFrontOriginAccessIdentity:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSCloudFrontCloudFrontOriginAccessIdentity not found", name)
}

// GetAllAWSCloudFrontDistributionResources retrieves all AWSCloudFrontDistribution items from an AWS CloudFormation template
func (t *Template) GetAllAWSCloudFrontDistributionResources() map[string]*resources.AWSCloudFrontDistribution {
	results := map[string]*resources.AWSCloudFrontDistribution{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSCloudFrontDistribution:
			results[name] = resource
		}
	}
	return results
}

// GetAWSCloudFrontDistributionWithName retrieves all AWSCloudFrontDistribution items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCloudFrontDistributionWithName(name string) (*resources.AWSCloudFrontDistribution, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSCloudFrontDistribution:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSCloudFrontDistribution not found", name)
}

// GetAllAWSCloudFrontStreamingDistributionResources retrieves all AWSCloudFrontStreamingDistribution items from an AWS CloudFormation template
func (t *Template) GetAllAWSCloudFrontStreamingDistributionResources() map[string]*resources.AWSCloudFrontStreamingDistribution {
	results := map[string]*resources.AWSCloudFrontStreamingDistribution{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSCloudFrontStreamingDistribution:
			results[name] = resource
		}
	}
	return results
}

// GetAWSCloudFrontStreamingDistributionWithName retrieves all AWSCloudFrontStreamingDistribution items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCloudFrontStreamingDistributionWithName(name string) (*resources.AWSCloudFrontStreamingDistribution, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSCloudFrontStreamingDistribution:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSCloudFrontStreamingDistribution not found", name)
}

// GetAllAWSCloudTrailTrailResources retrieves all AWSCloudTrailTrail items from an AWS CloudFormation template
func (t *Template) GetAllAWSCloudTrailTrailResources() map[string]*resources.AWSCloudTrailTrail {
	results := map[string]*resources.AWSCloudTrailTrail{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSCloudTrailTrail:
			results[name] = resource
		}
	}
	return results
}

// GetAWSCloudTrailTrailWithName retrieves all AWSCloudTrailTrail items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCloudTrailTrailWithName(name string) (*resources.AWSCloudTrailTrail, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSCloudTrailTrail:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSCloudTrailTrail not found", name)
}

// GetAllAWSCloudWatchAlarmResources retrieves all AWSCloudWatchAlarm items from an AWS CloudFormation template
func (t *Template) GetAllAWSCloudWatchAlarmResources() map[string]*resources.AWSCloudWatchAlarm {
	results := map[string]*resources.AWSCloudWatchAlarm{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSCloudWatchAlarm:
			results[name] = resource
		}
	}
	return results
}

// GetAWSCloudWatchAlarmWithName retrieves all AWSCloudWatchAlarm items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCloudWatchAlarmWithName(name string) (*resources.AWSCloudWatchAlarm, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSCloudWatchAlarm:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSCloudWatchAlarm not found", name)
}

// GetAllAWSCloudWatchDashboardResources retrieves all AWSCloudWatchDashboard items from an AWS CloudFormation template
func (t *Template) GetAllAWSCloudWatchDashboardResources() map[string]*resources.AWSCloudWatchDashboard {
	results := map[string]*resources.AWSCloudWatchDashboard{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSCloudWatchDashboard:
			results[name] = resource
		}
	}
	return results
}

// GetAWSCloudWatchDashboardWithName retrieves all AWSCloudWatchDashboard items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCloudWatchDashboardWithName(name string) (*resources.AWSCloudWatchDashboard, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSCloudWatchDashboard:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSCloudWatchDashboard not found", name)
}

// GetAllAWSCodeBuildProjectResources retrieves all AWSCodeBuildProject items from an AWS CloudFormation template
func (t *Template) GetAllAWSCodeBuildProjectResources() map[string]*resources.AWSCodeBuildProject {
	results := map[string]*resources.AWSCodeBuildProject{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSCodeBuildProject:
			results[name] = resource
		}
	}
	return results
}

// GetAWSCodeBuildProjectWithName retrieves all AWSCodeBuildProject items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCodeBuildProjectWithName(name string) (*resources.AWSCodeBuildProject, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSCodeBuildProject:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSCodeBuildProject not found", name)
}

// GetAllAWSCodeCommitRepositoryResources retrieves all AWSCodeCommitRepository items from an AWS CloudFormation template
func (t *Template) GetAllAWSCodeCommitRepositoryResources() map[string]*resources.AWSCodeCommitRepository {
	results := map[string]*resources.AWSCodeCommitRepository{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSCodeCommitRepository:
			results[name] = resource
		}
	}
	return results
}

// GetAWSCodeCommitRepositoryWithName retrieves all AWSCodeCommitRepository items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCodeCommitRepositoryWithName(name string) (*resources.AWSCodeCommitRepository, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSCodeCommitRepository:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSCodeCommitRepository not found", name)
}

// GetAllAWSCodeDeployApplicationResources retrieves all AWSCodeDeployApplication items from an AWS CloudFormation template
func (t *Template) GetAllAWSCodeDeployApplicationResources() map[string]*resources.AWSCodeDeployApplication {
	results := map[string]*resources.AWSCodeDeployApplication{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSCodeDeployApplication:
			results[name] = resource
		}
	}
	return results
}

// GetAWSCodeDeployApplicationWithName retrieves all AWSCodeDeployApplication items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCodeDeployApplicationWithName(name string) (*resources.AWSCodeDeployApplication, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSCodeDeployApplication:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSCodeDeployApplication not found", name)
}

// GetAllAWSCodeDeployDeploymentConfigResources retrieves all AWSCodeDeployDeploymentConfig items from an AWS CloudFormation template
func (t *Template) GetAllAWSCodeDeployDeploymentConfigResources() map[string]*resources.AWSCodeDeployDeploymentConfig {
	results := map[string]*resources.AWSCodeDeployDeploymentConfig{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSCodeDeployDeploymentConfig:
			results[name] = resource
		}
	}
	return results
}

// GetAWSCodeDeployDeploymentConfigWithName retrieves all AWSCodeDeployDeploymentConfig items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCodeDeployDeploymentConfigWithName(name string) (*resources.AWSCodeDeployDeploymentConfig, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSCodeDeployDeploymentConfig:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSCodeDeployDeploymentConfig not found", name)
}

// GetAllAWSCodeDeployDeploymentGroupResources retrieves all AWSCodeDeployDeploymentGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSCodeDeployDeploymentGroupResources() map[string]*resources.AWSCodeDeployDeploymentGroup {
	results := map[string]*resources.AWSCodeDeployDeploymentGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSCodeDeployDeploymentGroup:
			results[name] = resource
		}
	}
	return results
}

// GetAWSCodeDeployDeploymentGroupWithName retrieves all AWSCodeDeployDeploymentGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCodeDeployDeploymentGroupWithName(name string) (*resources.AWSCodeDeployDeploymentGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSCodeDeployDeploymentGroup:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSCodeDeployDeploymentGroup not found", name)
}

// GetAllAWSCodePipelineCustomActionTypeResources retrieves all AWSCodePipelineCustomActionType items from an AWS CloudFormation template
func (t *Template) GetAllAWSCodePipelineCustomActionTypeResources() map[string]*resources.AWSCodePipelineCustomActionType {
	results := map[string]*resources.AWSCodePipelineCustomActionType{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSCodePipelineCustomActionType:
			results[name] = resource
		}
	}
	return results
}

// GetAWSCodePipelineCustomActionTypeWithName retrieves all AWSCodePipelineCustomActionType items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCodePipelineCustomActionTypeWithName(name string) (*resources.AWSCodePipelineCustomActionType, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSCodePipelineCustomActionType:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSCodePipelineCustomActionType not found", name)
}

// GetAllAWSCodePipelinePipelineResources retrieves all AWSCodePipelinePipeline items from an AWS CloudFormation template
func (t *Template) GetAllAWSCodePipelinePipelineResources() map[string]*resources.AWSCodePipelinePipeline {
	results := map[string]*resources.AWSCodePipelinePipeline{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSCodePipelinePipeline:
			results[name] = resource
		}
	}
	return results
}

// GetAWSCodePipelinePipelineWithName retrieves all AWSCodePipelinePipeline items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCodePipelinePipelineWithName(name string) (*resources.AWSCodePipelinePipeline, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSCodePipelinePipeline:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSCodePipelinePipeline not found", name)
}

// GetAllAWSCodePipelineWebhookResources retrieves all AWSCodePipelineWebhook items from an AWS CloudFormation template
func (t *Template) GetAllAWSCodePipelineWebhookResources() map[string]*resources.AWSCodePipelineWebhook {
	results := map[string]*resources.AWSCodePipelineWebhook{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSCodePipelineWebhook:
			results[name] = resource
		}
	}
	return results
}

// GetAWSCodePipelineWebhookWithName retrieves all AWSCodePipelineWebhook items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCodePipelineWebhookWithName(name string) (*resources.AWSCodePipelineWebhook, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSCodePipelineWebhook:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSCodePipelineWebhook not found", name)
}

// GetAllAWSCognitoIdentityPoolResources retrieves all AWSCognitoIdentityPool items from an AWS CloudFormation template
func (t *Template) GetAllAWSCognitoIdentityPoolResources() map[string]*resources.AWSCognitoIdentityPool {
	results := map[string]*resources.AWSCognitoIdentityPool{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSCognitoIdentityPool:
			results[name] = resource
		}
	}
	return results
}

// GetAWSCognitoIdentityPoolWithName retrieves all AWSCognitoIdentityPool items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCognitoIdentityPoolWithName(name string) (*resources.AWSCognitoIdentityPool, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSCognitoIdentityPool:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSCognitoIdentityPool not found", name)
}

// GetAllAWSCognitoIdentityPoolRoleAttachmentResources retrieves all AWSCognitoIdentityPoolRoleAttachment items from an AWS CloudFormation template
func (t *Template) GetAllAWSCognitoIdentityPoolRoleAttachmentResources() map[string]*resources.AWSCognitoIdentityPoolRoleAttachment {
	results := map[string]*resources.AWSCognitoIdentityPoolRoleAttachment{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSCognitoIdentityPoolRoleAttachment:
			results[name] = resource
		}
	}
	return results
}

// GetAWSCognitoIdentityPoolRoleAttachmentWithName retrieves all AWSCognitoIdentityPoolRoleAttachment items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCognitoIdentityPoolRoleAttachmentWithName(name string) (*resources.AWSCognitoIdentityPoolRoleAttachment, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSCognitoIdentityPoolRoleAttachment:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSCognitoIdentityPoolRoleAttachment not found", name)
}

// GetAllAWSCognitoUserPoolResources retrieves all AWSCognitoUserPool items from an AWS CloudFormation template
func (t *Template) GetAllAWSCognitoUserPoolResources() map[string]*resources.AWSCognitoUserPool {
	results := map[string]*resources.AWSCognitoUserPool{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSCognitoUserPool:
			results[name] = resource
		}
	}
	return results
}

// GetAWSCognitoUserPoolWithName retrieves all AWSCognitoUserPool items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCognitoUserPoolWithName(name string) (*resources.AWSCognitoUserPool, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSCognitoUserPool:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSCognitoUserPool not found", name)
}

// GetAllAWSCognitoUserPoolClientResources retrieves all AWSCognitoUserPoolClient items from an AWS CloudFormation template
func (t *Template) GetAllAWSCognitoUserPoolClientResources() map[string]*resources.AWSCognitoUserPoolClient {
	results := map[string]*resources.AWSCognitoUserPoolClient{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSCognitoUserPoolClient:
			results[name] = resource
		}
	}
	return results
}

// GetAWSCognitoUserPoolClientWithName retrieves all AWSCognitoUserPoolClient items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCognitoUserPoolClientWithName(name string) (*resources.AWSCognitoUserPoolClient, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSCognitoUserPoolClient:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSCognitoUserPoolClient not found", name)
}

// GetAllAWSCognitoUserPoolGroupResources retrieves all AWSCognitoUserPoolGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSCognitoUserPoolGroupResources() map[string]*resources.AWSCognitoUserPoolGroup {
	results := map[string]*resources.AWSCognitoUserPoolGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSCognitoUserPoolGroup:
			results[name] = resource
		}
	}
	return results
}

// GetAWSCognitoUserPoolGroupWithName retrieves all AWSCognitoUserPoolGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCognitoUserPoolGroupWithName(name string) (*resources.AWSCognitoUserPoolGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSCognitoUserPoolGroup:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSCognitoUserPoolGroup not found", name)
}

// GetAllAWSCognitoUserPoolUserResources retrieves all AWSCognitoUserPoolUser items from an AWS CloudFormation template
func (t *Template) GetAllAWSCognitoUserPoolUserResources() map[string]*resources.AWSCognitoUserPoolUser {
	results := map[string]*resources.AWSCognitoUserPoolUser{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSCognitoUserPoolUser:
			results[name] = resource
		}
	}
	return results
}

// GetAWSCognitoUserPoolUserWithName retrieves all AWSCognitoUserPoolUser items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCognitoUserPoolUserWithName(name string) (*resources.AWSCognitoUserPoolUser, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSCognitoUserPoolUser:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSCognitoUserPoolUser not found", name)
}

// GetAllAWSCognitoUserPoolUserToGroupAttachmentResources retrieves all AWSCognitoUserPoolUserToGroupAttachment items from an AWS CloudFormation template
func (t *Template) GetAllAWSCognitoUserPoolUserToGroupAttachmentResources() map[string]*resources.AWSCognitoUserPoolUserToGroupAttachment {
	results := map[string]*resources.AWSCognitoUserPoolUserToGroupAttachment{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSCognitoUserPoolUserToGroupAttachment:
			results[name] = resource
		}
	}
	return results
}

// GetAWSCognitoUserPoolUserToGroupAttachmentWithName retrieves all AWSCognitoUserPoolUserToGroupAttachment items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCognitoUserPoolUserToGroupAttachmentWithName(name string) (*resources.AWSCognitoUserPoolUserToGroupAttachment, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSCognitoUserPoolUserToGroupAttachment:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSCognitoUserPoolUserToGroupAttachment not found", name)
}

// GetAllAWSConfigAggregationAuthorizationResources retrieves all AWSConfigAggregationAuthorization items from an AWS CloudFormation template
func (t *Template) GetAllAWSConfigAggregationAuthorizationResources() map[string]*resources.AWSConfigAggregationAuthorization {
	results := map[string]*resources.AWSConfigAggregationAuthorization{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSConfigAggregationAuthorization:
			results[name] = resource
		}
	}
	return results
}

// GetAWSConfigAggregationAuthorizationWithName retrieves all AWSConfigAggregationAuthorization items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSConfigAggregationAuthorizationWithName(name string) (*resources.AWSConfigAggregationAuthorization, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSConfigAggregationAuthorization:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSConfigAggregationAuthorization not found", name)
}

// GetAllAWSConfigConfigRuleResources retrieves all AWSConfigConfigRule items from an AWS CloudFormation template
func (t *Template) GetAllAWSConfigConfigRuleResources() map[string]*resources.AWSConfigConfigRule {
	results := map[string]*resources.AWSConfigConfigRule{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSConfigConfigRule:
			results[name] = resource
		}
	}
	return results
}

// GetAWSConfigConfigRuleWithName retrieves all AWSConfigConfigRule items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSConfigConfigRuleWithName(name string) (*resources.AWSConfigConfigRule, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSConfigConfigRule:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSConfigConfigRule not found", name)
}

// GetAllAWSConfigConfigurationAggregatorResources retrieves all AWSConfigConfigurationAggregator items from an AWS CloudFormation template
func (t *Template) GetAllAWSConfigConfigurationAggregatorResources() map[string]*resources.AWSConfigConfigurationAggregator {
	results := map[string]*resources.AWSConfigConfigurationAggregator{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSConfigConfigurationAggregator:
			results[name] = resource
		}
	}
	return results
}

// GetAWSConfigConfigurationAggregatorWithName retrieves all AWSConfigConfigurationAggregator items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSConfigConfigurationAggregatorWithName(name string) (*resources.AWSConfigConfigurationAggregator, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSConfigConfigurationAggregator:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSConfigConfigurationAggregator not found", name)
}

// GetAllAWSConfigConfigurationRecorderResources retrieves all AWSConfigConfigurationRecorder items from an AWS CloudFormation template
func (t *Template) GetAllAWSConfigConfigurationRecorderResources() map[string]*resources.AWSConfigConfigurationRecorder {
	results := map[string]*resources.AWSConfigConfigurationRecorder{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSConfigConfigurationRecorder:
			results[name] = resource
		}
	}
	return results
}

// GetAWSConfigConfigurationRecorderWithName retrieves all AWSConfigConfigurationRecorder items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSConfigConfigurationRecorderWithName(name string) (*resources.AWSConfigConfigurationRecorder, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSConfigConfigurationRecorder:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSConfigConfigurationRecorder not found", name)
}

// GetAllAWSConfigDeliveryChannelResources retrieves all AWSConfigDeliveryChannel items from an AWS CloudFormation template
func (t *Template) GetAllAWSConfigDeliveryChannelResources() map[string]*resources.AWSConfigDeliveryChannel {
	results := map[string]*resources.AWSConfigDeliveryChannel{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSConfigDeliveryChannel:
			results[name] = resource
		}
	}
	return results
}

// GetAWSConfigDeliveryChannelWithName retrieves all AWSConfigDeliveryChannel items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSConfigDeliveryChannelWithName(name string) (*resources.AWSConfigDeliveryChannel, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSConfigDeliveryChannel:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSConfigDeliveryChannel not found", name)
}

// GetAllAWSDAXClusterResources retrieves all AWSDAXCluster items from an AWS CloudFormation template
func (t *Template) GetAllAWSDAXClusterResources() map[string]*resources.AWSDAXCluster {
	results := map[string]*resources.AWSDAXCluster{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSDAXCluster:
			results[name] = resource
		}
	}
	return results
}

// GetAWSDAXClusterWithName retrieves all AWSDAXCluster items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSDAXClusterWithName(name string) (*resources.AWSDAXCluster, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSDAXCluster:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSDAXCluster not found", name)
}

// GetAllAWSDAXParameterGroupResources retrieves all AWSDAXParameterGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSDAXParameterGroupResources() map[string]*resources.AWSDAXParameterGroup {
	results := map[string]*resources.AWSDAXParameterGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSDAXParameterGroup:
			results[name] = resource
		}
	}
	return results
}

// GetAWSDAXParameterGroupWithName retrieves all AWSDAXParameterGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSDAXParameterGroupWithName(name string) (*resources.AWSDAXParameterGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSDAXParameterGroup:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSDAXParameterGroup not found", name)
}

// GetAllAWSDAXSubnetGroupResources retrieves all AWSDAXSubnetGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSDAXSubnetGroupResources() map[string]*resources.AWSDAXSubnetGroup {
	results := map[string]*resources.AWSDAXSubnetGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSDAXSubnetGroup:
			results[name] = resource
		}
	}
	return results
}

// GetAWSDAXSubnetGroupWithName retrieves all AWSDAXSubnetGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSDAXSubnetGroupWithName(name string) (*resources.AWSDAXSubnetGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSDAXSubnetGroup:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSDAXSubnetGroup not found", name)
}

// GetAllAWSDLMLifecyclePolicyResources retrieves all AWSDLMLifecyclePolicy items from an AWS CloudFormation template
func (t *Template) GetAllAWSDLMLifecyclePolicyResources() map[string]*resources.AWSDLMLifecyclePolicy {
	results := map[string]*resources.AWSDLMLifecyclePolicy{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSDLMLifecyclePolicy:
			results[name] = resource
		}
	}
	return results
}

// GetAWSDLMLifecyclePolicyWithName retrieves all AWSDLMLifecyclePolicy items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSDLMLifecyclePolicyWithName(name string) (*resources.AWSDLMLifecyclePolicy, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSDLMLifecyclePolicy:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSDLMLifecyclePolicy not found", name)
}

// GetAllAWSDMSCertificateResources retrieves all AWSDMSCertificate items from an AWS CloudFormation template
func (t *Template) GetAllAWSDMSCertificateResources() map[string]*resources.AWSDMSCertificate {
	results := map[string]*resources.AWSDMSCertificate{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSDMSCertificate:
			results[name] = resource
		}
	}
	return results
}

// GetAWSDMSCertificateWithName retrieves all AWSDMSCertificate items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSDMSCertificateWithName(name string) (*resources.AWSDMSCertificate, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSDMSCertificate:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSDMSCertificate not found", name)
}

// GetAllAWSDMSEndpointResources retrieves all AWSDMSEndpoint items from an AWS CloudFormation template
func (t *Template) GetAllAWSDMSEndpointResources() map[string]*resources.AWSDMSEndpoint {
	results := map[string]*resources.AWSDMSEndpoint{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSDMSEndpoint:
			results[name] = resource
		}
	}
	return results
}

// GetAWSDMSEndpointWithName retrieves all AWSDMSEndpoint items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSDMSEndpointWithName(name string) (*resources.AWSDMSEndpoint, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSDMSEndpoint:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSDMSEndpoint not found", name)
}

// GetAllAWSDMSEventSubscriptionResources retrieves all AWSDMSEventSubscription items from an AWS CloudFormation template
func (t *Template) GetAllAWSDMSEventSubscriptionResources() map[string]*resources.AWSDMSEventSubscription {
	results := map[string]*resources.AWSDMSEventSubscription{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSDMSEventSubscription:
			results[name] = resource
		}
	}
	return results
}

// GetAWSDMSEventSubscriptionWithName retrieves all AWSDMSEventSubscription items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSDMSEventSubscriptionWithName(name string) (*resources.AWSDMSEventSubscription, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSDMSEventSubscription:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSDMSEventSubscription not found", name)
}

// GetAllAWSDMSReplicationInstanceResources retrieves all AWSDMSReplicationInstance items from an AWS CloudFormation template
func (t *Template) GetAllAWSDMSReplicationInstanceResources() map[string]*resources.AWSDMSReplicationInstance {
	results := map[string]*resources.AWSDMSReplicationInstance{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSDMSReplicationInstance:
			results[name] = resource
		}
	}
	return results
}

// GetAWSDMSReplicationInstanceWithName retrieves all AWSDMSReplicationInstance items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSDMSReplicationInstanceWithName(name string) (*resources.AWSDMSReplicationInstance, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSDMSReplicationInstance:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSDMSReplicationInstance not found", name)
}

// GetAllAWSDMSReplicationSubnetGroupResources retrieves all AWSDMSReplicationSubnetGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSDMSReplicationSubnetGroupResources() map[string]*resources.AWSDMSReplicationSubnetGroup {
	results := map[string]*resources.AWSDMSReplicationSubnetGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSDMSReplicationSubnetGroup:
			results[name] = resource
		}
	}
	return results
}

// GetAWSDMSReplicationSubnetGroupWithName retrieves all AWSDMSReplicationSubnetGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSDMSReplicationSubnetGroupWithName(name string) (*resources.AWSDMSReplicationSubnetGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSDMSReplicationSubnetGroup:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSDMSReplicationSubnetGroup not found", name)
}

// GetAllAWSDMSReplicationTaskResources retrieves all AWSDMSReplicationTask items from an AWS CloudFormation template
func (t *Template) GetAllAWSDMSReplicationTaskResources() map[string]*resources.AWSDMSReplicationTask {
	results := map[string]*resources.AWSDMSReplicationTask{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSDMSReplicationTask:
			results[name] = resource
		}
	}
	return results
}

// GetAWSDMSReplicationTaskWithName retrieves all AWSDMSReplicationTask items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSDMSReplicationTaskWithName(name string) (*resources.AWSDMSReplicationTask, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSDMSReplicationTask:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSDMSReplicationTask not found", name)
}

// GetAllAWSDataPipelinePipelineResources retrieves all AWSDataPipelinePipeline items from an AWS CloudFormation template
func (t *Template) GetAllAWSDataPipelinePipelineResources() map[string]*resources.AWSDataPipelinePipeline {
	results := map[string]*resources.AWSDataPipelinePipeline{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSDataPipelinePipeline:
			results[name] = resource
		}
	}
	return results
}

// GetAWSDataPipelinePipelineWithName retrieves all AWSDataPipelinePipeline items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSDataPipelinePipelineWithName(name string) (*resources.AWSDataPipelinePipeline, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSDataPipelinePipeline:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSDataPipelinePipeline not found", name)
}

// GetAllAWSDirectoryServiceMicrosoftADResources retrieves all AWSDirectoryServiceMicrosoftAD items from an AWS CloudFormation template
func (t *Template) GetAllAWSDirectoryServiceMicrosoftADResources() map[string]*resources.AWSDirectoryServiceMicrosoftAD {
	results := map[string]*resources.AWSDirectoryServiceMicrosoftAD{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSDirectoryServiceMicrosoftAD:
			results[name] = resource
		}
	}
	return results
}

// GetAWSDirectoryServiceMicrosoftADWithName retrieves all AWSDirectoryServiceMicrosoftAD items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSDirectoryServiceMicrosoftADWithName(name string) (*resources.AWSDirectoryServiceMicrosoftAD, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSDirectoryServiceMicrosoftAD:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSDirectoryServiceMicrosoftAD not found", name)
}

// GetAllAWSDirectoryServiceSimpleADResources retrieves all AWSDirectoryServiceSimpleAD items from an AWS CloudFormation template
func (t *Template) GetAllAWSDirectoryServiceSimpleADResources() map[string]*resources.AWSDirectoryServiceSimpleAD {
	results := map[string]*resources.AWSDirectoryServiceSimpleAD{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSDirectoryServiceSimpleAD:
			results[name] = resource
		}
	}
	return results
}

// GetAWSDirectoryServiceSimpleADWithName retrieves all AWSDirectoryServiceSimpleAD items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSDirectoryServiceSimpleADWithName(name string) (*resources.AWSDirectoryServiceSimpleAD, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSDirectoryServiceSimpleAD:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSDirectoryServiceSimpleAD not found", name)
}

// GetAllAWSDocDBDBClusterResources retrieves all AWSDocDBDBCluster items from an AWS CloudFormation template
func (t *Template) GetAllAWSDocDBDBClusterResources() map[string]*resources.AWSDocDBDBCluster {
	results := map[string]*resources.AWSDocDBDBCluster{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSDocDBDBCluster:
			results[name] = resource
		}
	}
	return results
}

// GetAWSDocDBDBClusterWithName retrieves all AWSDocDBDBCluster items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSDocDBDBClusterWithName(name string) (*resources.AWSDocDBDBCluster, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSDocDBDBCluster:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSDocDBDBCluster not found", name)
}

// GetAllAWSDocDBDBClusterParameterGroupResources retrieves all AWSDocDBDBClusterParameterGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSDocDBDBClusterParameterGroupResources() map[string]*resources.AWSDocDBDBClusterParameterGroup {
	results := map[string]*resources.AWSDocDBDBClusterParameterGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSDocDBDBClusterParameterGroup:
			results[name] = resource
		}
	}
	return results
}

// GetAWSDocDBDBClusterParameterGroupWithName retrieves all AWSDocDBDBClusterParameterGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSDocDBDBClusterParameterGroupWithName(name string) (*resources.AWSDocDBDBClusterParameterGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSDocDBDBClusterParameterGroup:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSDocDBDBClusterParameterGroup not found", name)
}

// GetAllAWSDocDBDBInstanceResources retrieves all AWSDocDBDBInstance items from an AWS CloudFormation template
func (t *Template) GetAllAWSDocDBDBInstanceResources() map[string]*resources.AWSDocDBDBInstance {
	results := map[string]*resources.AWSDocDBDBInstance{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSDocDBDBInstance:
			results[name] = resource
		}
	}
	return results
}

// GetAWSDocDBDBInstanceWithName retrieves all AWSDocDBDBInstance items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSDocDBDBInstanceWithName(name string) (*resources.AWSDocDBDBInstance, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSDocDBDBInstance:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSDocDBDBInstance not found", name)
}

// GetAllAWSDocDBDBSubnetGroupResources retrieves all AWSDocDBDBSubnetGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSDocDBDBSubnetGroupResources() map[string]*resources.AWSDocDBDBSubnetGroup {
	results := map[string]*resources.AWSDocDBDBSubnetGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSDocDBDBSubnetGroup:
			results[name] = resource
		}
	}
	return results
}

// GetAWSDocDBDBSubnetGroupWithName retrieves all AWSDocDBDBSubnetGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSDocDBDBSubnetGroupWithName(name string) (*resources.AWSDocDBDBSubnetGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSDocDBDBSubnetGroup:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSDocDBDBSubnetGroup not found", name)
}

// GetAllAWSDynamoDBTableResources retrieves all AWSDynamoDBTable items from an AWS CloudFormation template
func (t *Template) GetAllAWSDynamoDBTableResources() map[string]*resources.AWSDynamoDBTable {
	results := map[string]*resources.AWSDynamoDBTable{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSDynamoDBTable:
			results[name] = resource
		}
	}
	return results
}

// GetAWSDynamoDBTableWithName retrieves all AWSDynamoDBTable items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSDynamoDBTableWithName(name string) (*resources.AWSDynamoDBTable, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSDynamoDBTable:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSDynamoDBTable not found", name)
}

// GetAllAWSEC2CustomerGatewayResources retrieves all AWSEC2CustomerGateway items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2CustomerGatewayResources() map[string]*resources.AWSEC2CustomerGateway {
	results := map[string]*resources.AWSEC2CustomerGateway{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSEC2CustomerGateway:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2CustomerGatewayWithName retrieves all AWSEC2CustomerGateway items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2CustomerGatewayWithName(name string) (*resources.AWSEC2CustomerGateway, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSEC2CustomerGateway:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2CustomerGateway not found", name)
}

// GetAllAWSEC2DHCPOptionsResources retrieves all AWSEC2DHCPOptions items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2DHCPOptionsResources() map[string]*resources.AWSEC2DHCPOptions {
	results := map[string]*resources.AWSEC2DHCPOptions{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSEC2DHCPOptions:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2DHCPOptionsWithName retrieves all AWSEC2DHCPOptions items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2DHCPOptionsWithName(name string) (*resources.AWSEC2DHCPOptions, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSEC2DHCPOptions:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2DHCPOptions not found", name)
}

// GetAllAWSEC2EC2FleetResources retrieves all AWSEC2EC2Fleet items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2EC2FleetResources() map[string]*resources.AWSEC2EC2Fleet {
	results := map[string]*resources.AWSEC2EC2Fleet{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSEC2EC2Fleet:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2EC2FleetWithName retrieves all AWSEC2EC2Fleet items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2EC2FleetWithName(name string) (*resources.AWSEC2EC2Fleet, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSEC2EC2Fleet:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2EC2Fleet not found", name)
}

// GetAllAWSEC2EIPResources retrieves all AWSEC2EIP items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2EIPResources() map[string]*resources.AWSEC2EIP {
	results := map[string]*resources.AWSEC2EIP{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSEC2EIP:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2EIPWithName retrieves all AWSEC2EIP items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2EIPWithName(name string) (*resources.AWSEC2EIP, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSEC2EIP:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2EIP not found", name)
}

// GetAllAWSEC2EIPAssociationResources retrieves all AWSEC2EIPAssociation items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2EIPAssociationResources() map[string]*resources.AWSEC2EIPAssociation {
	results := map[string]*resources.AWSEC2EIPAssociation{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSEC2EIPAssociation:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2EIPAssociationWithName retrieves all AWSEC2EIPAssociation items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2EIPAssociationWithName(name string) (*resources.AWSEC2EIPAssociation, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSEC2EIPAssociation:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2EIPAssociation not found", name)
}

// GetAllAWSEC2EgressOnlyInternetGatewayResources retrieves all AWSEC2EgressOnlyInternetGateway items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2EgressOnlyInternetGatewayResources() map[string]*resources.AWSEC2EgressOnlyInternetGateway {
	results := map[string]*resources.AWSEC2EgressOnlyInternetGateway{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSEC2EgressOnlyInternetGateway:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2EgressOnlyInternetGatewayWithName retrieves all AWSEC2EgressOnlyInternetGateway items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2EgressOnlyInternetGatewayWithName(name string) (*resources.AWSEC2EgressOnlyInternetGateway, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSEC2EgressOnlyInternetGateway:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2EgressOnlyInternetGateway not found", name)
}

// GetAllAWSEC2FlowLogResources retrieves all AWSEC2FlowLog items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2FlowLogResources() map[string]*resources.AWSEC2FlowLog {
	results := map[string]*resources.AWSEC2FlowLog{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSEC2FlowLog:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2FlowLogWithName retrieves all AWSEC2FlowLog items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2FlowLogWithName(name string) (*resources.AWSEC2FlowLog, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSEC2FlowLog:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2FlowLog not found", name)
}

// GetAllAWSEC2HostResources retrieves all AWSEC2Host items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2HostResources() map[string]*resources.AWSEC2Host {
	results := map[string]*resources.AWSEC2Host{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSEC2Host:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2HostWithName retrieves all AWSEC2Host items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2HostWithName(name string) (*resources.AWSEC2Host, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSEC2Host:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2Host not found", name)
}

// GetAllAWSEC2InstanceResources retrieves all AWSEC2Instance items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2InstanceResources() map[string]*resources.AWSEC2Instance {
	results := map[string]*resources.AWSEC2Instance{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSEC2Instance:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2InstanceWithName retrieves all AWSEC2Instance items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2InstanceWithName(name string) (*resources.AWSEC2Instance, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSEC2Instance:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2Instance not found", name)
}

// GetAllAWSEC2InternetGatewayResources retrieves all AWSEC2InternetGateway items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2InternetGatewayResources() map[string]*resources.AWSEC2InternetGateway {
	results := map[string]*resources.AWSEC2InternetGateway{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSEC2InternetGateway:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2InternetGatewayWithName retrieves all AWSEC2InternetGateway items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2InternetGatewayWithName(name string) (*resources.AWSEC2InternetGateway, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSEC2InternetGateway:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2InternetGateway not found", name)
}

// GetAllAWSEC2LaunchTemplateResources retrieves all AWSEC2LaunchTemplate items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2LaunchTemplateResources() map[string]*resources.AWSEC2LaunchTemplate {
	results := map[string]*resources.AWSEC2LaunchTemplate{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSEC2LaunchTemplate:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2LaunchTemplateWithName retrieves all AWSEC2LaunchTemplate items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2LaunchTemplateWithName(name string) (*resources.AWSEC2LaunchTemplate, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSEC2LaunchTemplate:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2LaunchTemplate not found", name)
}

// GetAllAWSEC2NatGatewayResources retrieves all AWSEC2NatGateway items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2NatGatewayResources() map[string]*resources.AWSEC2NatGateway {
	results := map[string]*resources.AWSEC2NatGateway{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSEC2NatGateway:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2NatGatewayWithName retrieves all AWSEC2NatGateway items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2NatGatewayWithName(name string) (*resources.AWSEC2NatGateway, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSEC2NatGateway:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2NatGateway not found", name)
}

// GetAllAWSEC2NetworkAclResources retrieves all AWSEC2NetworkAcl items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2NetworkAclResources() map[string]*resources.AWSEC2NetworkAcl {
	results := map[string]*resources.AWSEC2NetworkAcl{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSEC2NetworkAcl:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2NetworkAclWithName retrieves all AWSEC2NetworkAcl items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2NetworkAclWithName(name string) (*resources.AWSEC2NetworkAcl, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSEC2NetworkAcl:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2NetworkAcl not found", name)
}

// GetAllAWSEC2NetworkAclEntryResources retrieves all AWSEC2NetworkAclEntry items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2NetworkAclEntryResources() map[string]*resources.AWSEC2NetworkAclEntry {
	results := map[string]*resources.AWSEC2NetworkAclEntry{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSEC2NetworkAclEntry:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2NetworkAclEntryWithName retrieves all AWSEC2NetworkAclEntry items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2NetworkAclEntryWithName(name string) (*resources.AWSEC2NetworkAclEntry, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSEC2NetworkAclEntry:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2NetworkAclEntry not found", name)
}

// GetAllAWSEC2NetworkInterfaceResources retrieves all AWSEC2NetworkInterface items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2NetworkInterfaceResources() map[string]*resources.AWSEC2NetworkInterface {
	results := map[string]*resources.AWSEC2NetworkInterface{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSEC2NetworkInterface:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2NetworkInterfaceWithName retrieves all AWSEC2NetworkInterface items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2NetworkInterfaceWithName(name string) (*resources.AWSEC2NetworkInterface, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSEC2NetworkInterface:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2NetworkInterface not found", name)
}

// GetAllAWSEC2NetworkInterfaceAttachmentResources retrieves all AWSEC2NetworkInterfaceAttachment items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2NetworkInterfaceAttachmentResources() map[string]*resources.AWSEC2NetworkInterfaceAttachment {
	results := map[string]*resources.AWSEC2NetworkInterfaceAttachment{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSEC2NetworkInterfaceAttachment:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2NetworkInterfaceAttachmentWithName retrieves all AWSEC2NetworkInterfaceAttachment items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2NetworkInterfaceAttachmentWithName(name string) (*resources.AWSEC2NetworkInterfaceAttachment, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSEC2NetworkInterfaceAttachment:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2NetworkInterfaceAttachment not found", name)
}

// GetAllAWSEC2NetworkInterfacePermissionResources retrieves all AWSEC2NetworkInterfacePermission items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2NetworkInterfacePermissionResources() map[string]*resources.AWSEC2NetworkInterfacePermission {
	results := map[string]*resources.AWSEC2NetworkInterfacePermission{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSEC2NetworkInterfacePermission:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2NetworkInterfacePermissionWithName retrieves all AWSEC2NetworkInterfacePermission items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2NetworkInterfacePermissionWithName(name string) (*resources.AWSEC2NetworkInterfacePermission, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSEC2NetworkInterfacePermission:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2NetworkInterfacePermission not found", name)
}

// GetAllAWSEC2PlacementGroupResources retrieves all AWSEC2PlacementGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2PlacementGroupResources() map[string]*resources.AWSEC2PlacementGroup {
	results := map[string]*resources.AWSEC2PlacementGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSEC2PlacementGroup:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2PlacementGroupWithName retrieves all AWSEC2PlacementGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2PlacementGroupWithName(name string) (*resources.AWSEC2PlacementGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSEC2PlacementGroup:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2PlacementGroup not found", name)
}

// GetAllAWSEC2RouteResources retrieves all AWSEC2Route items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2RouteResources() map[string]*resources.AWSEC2Route {
	results := map[string]*resources.AWSEC2Route{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSEC2Route:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2RouteWithName retrieves all AWSEC2Route items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2RouteWithName(name string) (*resources.AWSEC2Route, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSEC2Route:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2Route not found", name)
}

// GetAllAWSEC2RouteTableResources retrieves all AWSEC2RouteTable items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2RouteTableResources() map[string]*resources.AWSEC2RouteTable {
	results := map[string]*resources.AWSEC2RouteTable{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSEC2RouteTable:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2RouteTableWithName retrieves all AWSEC2RouteTable items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2RouteTableWithName(name string) (*resources.AWSEC2RouteTable, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSEC2RouteTable:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2RouteTable not found", name)
}

// GetAllAWSEC2SecurityGroupResources retrieves all AWSEC2SecurityGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2SecurityGroupResources() map[string]*resources.AWSEC2SecurityGroup {
	results := map[string]*resources.AWSEC2SecurityGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSEC2SecurityGroup:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2SecurityGroupWithName retrieves all AWSEC2SecurityGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2SecurityGroupWithName(name string) (*resources.AWSEC2SecurityGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSEC2SecurityGroup:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2SecurityGroup not found", name)
}

// GetAllAWSEC2SecurityGroupEgressResources retrieves all AWSEC2SecurityGroupEgress items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2SecurityGroupEgressResources() map[string]*resources.AWSEC2SecurityGroupEgress {
	results := map[string]*resources.AWSEC2SecurityGroupEgress{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSEC2SecurityGroupEgress:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2SecurityGroupEgressWithName retrieves all AWSEC2SecurityGroupEgress items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2SecurityGroupEgressWithName(name string) (*resources.AWSEC2SecurityGroupEgress, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSEC2SecurityGroupEgress:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2SecurityGroupEgress not found", name)
}

// GetAllAWSEC2SecurityGroupIngressResources retrieves all AWSEC2SecurityGroupIngress items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2SecurityGroupIngressResources() map[string]*resources.AWSEC2SecurityGroupIngress {
	results := map[string]*resources.AWSEC2SecurityGroupIngress{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSEC2SecurityGroupIngress:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2SecurityGroupIngressWithName retrieves all AWSEC2SecurityGroupIngress items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2SecurityGroupIngressWithName(name string) (*resources.AWSEC2SecurityGroupIngress, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSEC2SecurityGroupIngress:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2SecurityGroupIngress not found", name)
}

// GetAllAWSEC2SpotFleetResources retrieves all AWSEC2SpotFleet items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2SpotFleetResources() map[string]*resources.AWSEC2SpotFleet {
	results := map[string]*resources.AWSEC2SpotFleet{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSEC2SpotFleet:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2SpotFleetWithName retrieves all AWSEC2SpotFleet items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2SpotFleetWithName(name string) (*resources.AWSEC2SpotFleet, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSEC2SpotFleet:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2SpotFleet not found", name)
}

// GetAllAWSEC2SubnetResources retrieves all AWSEC2Subnet items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2SubnetResources() map[string]*resources.AWSEC2Subnet {
	results := map[string]*resources.AWSEC2Subnet{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSEC2Subnet:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2SubnetWithName retrieves all AWSEC2Subnet items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2SubnetWithName(name string) (*resources.AWSEC2Subnet, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSEC2Subnet:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2Subnet not found", name)
}

// GetAllAWSEC2SubnetCidrBlockResources retrieves all AWSEC2SubnetCidrBlock items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2SubnetCidrBlockResources() map[string]*resources.AWSEC2SubnetCidrBlock {
	results := map[string]*resources.AWSEC2SubnetCidrBlock{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSEC2SubnetCidrBlock:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2SubnetCidrBlockWithName retrieves all AWSEC2SubnetCidrBlock items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2SubnetCidrBlockWithName(name string) (*resources.AWSEC2SubnetCidrBlock, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSEC2SubnetCidrBlock:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2SubnetCidrBlock not found", name)
}

// GetAllAWSEC2SubnetNetworkAclAssociationResources retrieves all AWSEC2SubnetNetworkAclAssociation items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2SubnetNetworkAclAssociationResources() map[string]*resources.AWSEC2SubnetNetworkAclAssociation {
	results := map[string]*resources.AWSEC2SubnetNetworkAclAssociation{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSEC2SubnetNetworkAclAssociation:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2SubnetNetworkAclAssociationWithName retrieves all AWSEC2SubnetNetworkAclAssociation items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2SubnetNetworkAclAssociationWithName(name string) (*resources.AWSEC2SubnetNetworkAclAssociation, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSEC2SubnetNetworkAclAssociation:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2SubnetNetworkAclAssociation not found", name)
}

// GetAllAWSEC2SubnetRouteTableAssociationResources retrieves all AWSEC2SubnetRouteTableAssociation items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2SubnetRouteTableAssociationResources() map[string]*resources.AWSEC2SubnetRouteTableAssociation {
	results := map[string]*resources.AWSEC2SubnetRouteTableAssociation{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSEC2SubnetRouteTableAssociation:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2SubnetRouteTableAssociationWithName retrieves all AWSEC2SubnetRouteTableAssociation items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2SubnetRouteTableAssociationWithName(name string) (*resources.AWSEC2SubnetRouteTableAssociation, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSEC2SubnetRouteTableAssociation:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2SubnetRouteTableAssociation not found", name)
}

// GetAllAWSEC2TransitGatewayResources retrieves all AWSEC2TransitGateway items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2TransitGatewayResources() map[string]*resources.AWSEC2TransitGateway {
	results := map[string]*resources.AWSEC2TransitGateway{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSEC2TransitGateway:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2TransitGatewayWithName retrieves all AWSEC2TransitGateway items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2TransitGatewayWithName(name string) (*resources.AWSEC2TransitGateway, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSEC2TransitGateway:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2TransitGateway not found", name)
}

// GetAllAWSEC2TransitGatewayAttachmentResources retrieves all AWSEC2TransitGatewayAttachment items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2TransitGatewayAttachmentResources() map[string]*resources.AWSEC2TransitGatewayAttachment {
	results := map[string]*resources.AWSEC2TransitGatewayAttachment{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSEC2TransitGatewayAttachment:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2TransitGatewayAttachmentWithName retrieves all AWSEC2TransitGatewayAttachment items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2TransitGatewayAttachmentWithName(name string) (*resources.AWSEC2TransitGatewayAttachment, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSEC2TransitGatewayAttachment:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2TransitGatewayAttachment not found", name)
}

// GetAllAWSEC2TransitGatewayRouteResources retrieves all AWSEC2TransitGatewayRoute items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2TransitGatewayRouteResources() map[string]*resources.AWSEC2TransitGatewayRoute {
	results := map[string]*resources.AWSEC2TransitGatewayRoute{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSEC2TransitGatewayRoute:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2TransitGatewayRouteWithName retrieves all AWSEC2TransitGatewayRoute items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2TransitGatewayRouteWithName(name string) (*resources.AWSEC2TransitGatewayRoute, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSEC2TransitGatewayRoute:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2TransitGatewayRoute not found", name)
}

// GetAllAWSEC2TransitGatewayRouteTableResources retrieves all AWSEC2TransitGatewayRouteTable items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2TransitGatewayRouteTableResources() map[string]*resources.AWSEC2TransitGatewayRouteTable {
	results := map[string]*resources.AWSEC2TransitGatewayRouteTable{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSEC2TransitGatewayRouteTable:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2TransitGatewayRouteTableWithName retrieves all AWSEC2TransitGatewayRouteTable items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2TransitGatewayRouteTableWithName(name string) (*resources.AWSEC2TransitGatewayRouteTable, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSEC2TransitGatewayRouteTable:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2TransitGatewayRouteTable not found", name)
}

// GetAllAWSEC2TransitGatewayRouteTableAssociationResources retrieves all AWSEC2TransitGatewayRouteTableAssociation items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2TransitGatewayRouteTableAssociationResources() map[string]*resources.AWSEC2TransitGatewayRouteTableAssociation {
	results := map[string]*resources.AWSEC2TransitGatewayRouteTableAssociation{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSEC2TransitGatewayRouteTableAssociation:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2TransitGatewayRouteTableAssociationWithName retrieves all AWSEC2TransitGatewayRouteTableAssociation items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2TransitGatewayRouteTableAssociationWithName(name string) (*resources.AWSEC2TransitGatewayRouteTableAssociation, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSEC2TransitGatewayRouteTableAssociation:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2TransitGatewayRouteTableAssociation not found", name)
}

// GetAllAWSEC2TransitGatewayRouteTablePropagationResources retrieves all AWSEC2TransitGatewayRouteTablePropagation items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2TransitGatewayRouteTablePropagationResources() map[string]*resources.AWSEC2TransitGatewayRouteTablePropagation {
	results := map[string]*resources.AWSEC2TransitGatewayRouteTablePropagation{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSEC2TransitGatewayRouteTablePropagation:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2TransitGatewayRouteTablePropagationWithName retrieves all AWSEC2TransitGatewayRouteTablePropagation items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2TransitGatewayRouteTablePropagationWithName(name string) (*resources.AWSEC2TransitGatewayRouteTablePropagation, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSEC2TransitGatewayRouteTablePropagation:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2TransitGatewayRouteTablePropagation not found", name)
}

// GetAllAWSEC2TrunkInterfaceAssociationResources retrieves all AWSEC2TrunkInterfaceAssociation items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2TrunkInterfaceAssociationResources() map[string]*resources.AWSEC2TrunkInterfaceAssociation {
	results := map[string]*resources.AWSEC2TrunkInterfaceAssociation{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSEC2TrunkInterfaceAssociation:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2TrunkInterfaceAssociationWithName retrieves all AWSEC2TrunkInterfaceAssociation items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2TrunkInterfaceAssociationWithName(name string) (*resources.AWSEC2TrunkInterfaceAssociation, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSEC2TrunkInterfaceAssociation:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2TrunkInterfaceAssociation not found", name)
}

// GetAllAWSEC2VPCResources retrieves all AWSEC2VPC items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2VPCResources() map[string]*resources.AWSEC2VPC {
	results := map[string]*resources.AWSEC2VPC{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSEC2VPC:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2VPCWithName retrieves all AWSEC2VPC items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2VPCWithName(name string) (*resources.AWSEC2VPC, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSEC2VPC:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2VPC not found", name)
}

// GetAllAWSEC2VPCCidrBlockResources retrieves all AWSEC2VPCCidrBlock items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2VPCCidrBlockResources() map[string]*resources.AWSEC2VPCCidrBlock {
	results := map[string]*resources.AWSEC2VPCCidrBlock{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSEC2VPCCidrBlock:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2VPCCidrBlockWithName retrieves all AWSEC2VPCCidrBlock items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2VPCCidrBlockWithName(name string) (*resources.AWSEC2VPCCidrBlock, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSEC2VPCCidrBlock:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2VPCCidrBlock not found", name)
}

// GetAllAWSEC2VPCDHCPOptionsAssociationResources retrieves all AWSEC2VPCDHCPOptionsAssociation items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2VPCDHCPOptionsAssociationResources() map[string]*resources.AWSEC2VPCDHCPOptionsAssociation {
	results := map[string]*resources.AWSEC2VPCDHCPOptionsAssociation{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSEC2VPCDHCPOptionsAssociation:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2VPCDHCPOptionsAssociationWithName retrieves all AWSEC2VPCDHCPOptionsAssociation items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2VPCDHCPOptionsAssociationWithName(name string) (*resources.AWSEC2VPCDHCPOptionsAssociation, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSEC2VPCDHCPOptionsAssociation:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2VPCDHCPOptionsAssociation not found", name)
}

// GetAllAWSEC2VPCEndpointResources retrieves all AWSEC2VPCEndpoint items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2VPCEndpointResources() map[string]*resources.AWSEC2VPCEndpoint {
	results := map[string]*resources.AWSEC2VPCEndpoint{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSEC2VPCEndpoint:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2VPCEndpointWithName retrieves all AWSEC2VPCEndpoint items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2VPCEndpointWithName(name string) (*resources.AWSEC2VPCEndpoint, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSEC2VPCEndpoint:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2VPCEndpoint not found", name)
}

// GetAllAWSEC2VPCEndpointConnectionNotificationResources retrieves all AWSEC2VPCEndpointConnectionNotification items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2VPCEndpointConnectionNotificationResources() map[string]*resources.AWSEC2VPCEndpointConnectionNotification {
	results := map[string]*resources.AWSEC2VPCEndpointConnectionNotification{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSEC2VPCEndpointConnectionNotification:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2VPCEndpointConnectionNotificationWithName retrieves all AWSEC2VPCEndpointConnectionNotification items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2VPCEndpointConnectionNotificationWithName(name string) (*resources.AWSEC2VPCEndpointConnectionNotification, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSEC2VPCEndpointConnectionNotification:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2VPCEndpointConnectionNotification not found", name)
}

// GetAllAWSEC2VPCEndpointServicePermissionsResources retrieves all AWSEC2VPCEndpointServicePermissions items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2VPCEndpointServicePermissionsResources() map[string]*resources.AWSEC2VPCEndpointServicePermissions {
	results := map[string]*resources.AWSEC2VPCEndpointServicePermissions{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSEC2VPCEndpointServicePermissions:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2VPCEndpointServicePermissionsWithName retrieves all AWSEC2VPCEndpointServicePermissions items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2VPCEndpointServicePermissionsWithName(name string) (*resources.AWSEC2VPCEndpointServicePermissions, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSEC2VPCEndpointServicePermissions:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2VPCEndpointServicePermissions not found", name)
}

// GetAllAWSEC2VPCGatewayAttachmentResources retrieves all AWSEC2VPCGatewayAttachment items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2VPCGatewayAttachmentResources() map[string]*resources.AWSEC2VPCGatewayAttachment {
	results := map[string]*resources.AWSEC2VPCGatewayAttachment{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSEC2VPCGatewayAttachment:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2VPCGatewayAttachmentWithName retrieves all AWSEC2VPCGatewayAttachment items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2VPCGatewayAttachmentWithName(name string) (*resources.AWSEC2VPCGatewayAttachment, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSEC2VPCGatewayAttachment:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2VPCGatewayAttachment not found", name)
}

// GetAllAWSEC2VPCPeeringConnectionResources retrieves all AWSEC2VPCPeeringConnection items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2VPCPeeringConnectionResources() map[string]*resources.AWSEC2VPCPeeringConnection {
	results := map[string]*resources.AWSEC2VPCPeeringConnection{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSEC2VPCPeeringConnection:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2VPCPeeringConnectionWithName retrieves all AWSEC2VPCPeeringConnection items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2VPCPeeringConnectionWithName(name string) (*resources.AWSEC2VPCPeeringConnection, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSEC2VPCPeeringConnection:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2VPCPeeringConnection not found", name)
}

// GetAllAWSEC2VPNConnectionResources retrieves all AWSEC2VPNConnection items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2VPNConnectionResources() map[string]*resources.AWSEC2VPNConnection {
	results := map[string]*resources.AWSEC2VPNConnection{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSEC2VPNConnection:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2VPNConnectionWithName retrieves all AWSEC2VPNConnection items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2VPNConnectionWithName(name string) (*resources.AWSEC2VPNConnection, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSEC2VPNConnection:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2VPNConnection not found", name)
}

// GetAllAWSEC2VPNConnectionRouteResources retrieves all AWSEC2VPNConnectionRoute items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2VPNConnectionRouteResources() map[string]*resources.AWSEC2VPNConnectionRoute {
	results := map[string]*resources.AWSEC2VPNConnectionRoute{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSEC2VPNConnectionRoute:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2VPNConnectionRouteWithName retrieves all AWSEC2VPNConnectionRoute items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2VPNConnectionRouteWithName(name string) (*resources.AWSEC2VPNConnectionRoute, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSEC2VPNConnectionRoute:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2VPNConnectionRoute not found", name)
}

// GetAllAWSEC2VPNGatewayResources retrieves all AWSEC2VPNGateway items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2VPNGatewayResources() map[string]*resources.AWSEC2VPNGateway {
	results := map[string]*resources.AWSEC2VPNGateway{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSEC2VPNGateway:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2VPNGatewayWithName retrieves all AWSEC2VPNGateway items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2VPNGatewayWithName(name string) (*resources.AWSEC2VPNGateway, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSEC2VPNGateway:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2VPNGateway not found", name)
}

// GetAllAWSEC2VPNGatewayRoutePropagationResources retrieves all AWSEC2VPNGatewayRoutePropagation items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2VPNGatewayRoutePropagationResources() map[string]*resources.AWSEC2VPNGatewayRoutePropagation {
	results := map[string]*resources.AWSEC2VPNGatewayRoutePropagation{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSEC2VPNGatewayRoutePropagation:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2VPNGatewayRoutePropagationWithName retrieves all AWSEC2VPNGatewayRoutePropagation items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2VPNGatewayRoutePropagationWithName(name string) (*resources.AWSEC2VPNGatewayRoutePropagation, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSEC2VPNGatewayRoutePropagation:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2VPNGatewayRoutePropagation not found", name)
}

// GetAllAWSEC2VolumeResources retrieves all AWSEC2Volume items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2VolumeResources() map[string]*resources.AWSEC2Volume {
	results := map[string]*resources.AWSEC2Volume{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSEC2Volume:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2VolumeWithName retrieves all AWSEC2Volume items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2VolumeWithName(name string) (*resources.AWSEC2Volume, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSEC2Volume:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2Volume not found", name)
}

// GetAllAWSEC2VolumeAttachmentResources retrieves all AWSEC2VolumeAttachment items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2VolumeAttachmentResources() map[string]*resources.AWSEC2VolumeAttachment {
	results := map[string]*resources.AWSEC2VolumeAttachment{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSEC2VolumeAttachment:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEC2VolumeAttachmentWithName retrieves all AWSEC2VolumeAttachment items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2VolumeAttachmentWithName(name string) (*resources.AWSEC2VolumeAttachment, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSEC2VolumeAttachment:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEC2VolumeAttachment not found", name)
}

// GetAllAWSECRRepositoryResources retrieves all AWSECRRepository items from an AWS CloudFormation template
func (t *Template) GetAllAWSECRRepositoryResources() map[string]*resources.AWSECRRepository {
	results := map[string]*resources.AWSECRRepository{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSECRRepository:
			results[name] = resource
		}
	}
	return results
}

// GetAWSECRRepositoryWithName retrieves all AWSECRRepository items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSECRRepositoryWithName(name string) (*resources.AWSECRRepository, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSECRRepository:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSECRRepository not found", name)
}

// GetAllAWSECSClusterResources retrieves all AWSECSCluster items from an AWS CloudFormation template
func (t *Template) GetAllAWSECSClusterResources() map[string]*resources.AWSECSCluster {
	results := map[string]*resources.AWSECSCluster{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSECSCluster:
			results[name] = resource
		}
	}
	return results
}

// GetAWSECSClusterWithName retrieves all AWSECSCluster items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSECSClusterWithName(name string) (*resources.AWSECSCluster, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSECSCluster:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSECSCluster not found", name)
}

// GetAllAWSECSServiceResources retrieves all AWSECSService items from an AWS CloudFormation template
func (t *Template) GetAllAWSECSServiceResources() map[string]*resources.AWSECSService {
	results := map[string]*resources.AWSECSService{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSECSService:
			results[name] = resource
		}
	}
	return results
}

// GetAWSECSServiceWithName retrieves all AWSECSService items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSECSServiceWithName(name string) (*resources.AWSECSService, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSECSService:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSECSService not found", name)
}

// GetAllAWSECSTaskDefinitionResources retrieves all AWSECSTaskDefinition items from an AWS CloudFormation template
func (t *Template) GetAllAWSECSTaskDefinitionResources() map[string]*resources.AWSECSTaskDefinition {
	results := map[string]*resources.AWSECSTaskDefinition{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSECSTaskDefinition:
			results[name] = resource
		}
	}
	return results
}

// GetAWSECSTaskDefinitionWithName retrieves all AWSECSTaskDefinition items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSECSTaskDefinitionWithName(name string) (*resources.AWSECSTaskDefinition, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSECSTaskDefinition:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSECSTaskDefinition not found", name)
}

// GetAllAWSEFSFileSystemResources retrieves all AWSEFSFileSystem items from an AWS CloudFormation template
func (t *Template) GetAllAWSEFSFileSystemResources() map[string]*resources.AWSEFSFileSystem {
	results := map[string]*resources.AWSEFSFileSystem{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSEFSFileSystem:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEFSFileSystemWithName retrieves all AWSEFSFileSystem items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEFSFileSystemWithName(name string) (*resources.AWSEFSFileSystem, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSEFSFileSystem:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEFSFileSystem not found", name)
}

// GetAllAWSEFSMountTargetResources retrieves all AWSEFSMountTarget items from an AWS CloudFormation template
func (t *Template) GetAllAWSEFSMountTargetResources() map[string]*resources.AWSEFSMountTarget {
	results := map[string]*resources.AWSEFSMountTarget{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSEFSMountTarget:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEFSMountTargetWithName retrieves all AWSEFSMountTarget items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEFSMountTargetWithName(name string) (*resources.AWSEFSMountTarget, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSEFSMountTarget:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEFSMountTarget not found", name)
}

// GetAllAWSEKSClusterResources retrieves all AWSEKSCluster items from an AWS CloudFormation template
func (t *Template) GetAllAWSEKSClusterResources() map[string]*resources.AWSEKSCluster {
	results := map[string]*resources.AWSEKSCluster{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSEKSCluster:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEKSClusterWithName retrieves all AWSEKSCluster items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEKSClusterWithName(name string) (*resources.AWSEKSCluster, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSEKSCluster:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEKSCluster not found", name)
}

// GetAllAWSEMRClusterResources retrieves all AWSEMRCluster items from an AWS CloudFormation template
func (t *Template) GetAllAWSEMRClusterResources() map[string]*resources.AWSEMRCluster {
	results := map[string]*resources.AWSEMRCluster{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSEMRCluster:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEMRClusterWithName retrieves all AWSEMRCluster items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEMRClusterWithName(name string) (*resources.AWSEMRCluster, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSEMRCluster:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEMRCluster not found", name)
}

// GetAllAWSEMRInstanceFleetConfigResources retrieves all AWSEMRInstanceFleetConfig items from an AWS CloudFormation template
func (t *Template) GetAllAWSEMRInstanceFleetConfigResources() map[string]*resources.AWSEMRInstanceFleetConfig {
	results := map[string]*resources.AWSEMRInstanceFleetConfig{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSEMRInstanceFleetConfig:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEMRInstanceFleetConfigWithName retrieves all AWSEMRInstanceFleetConfig items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEMRInstanceFleetConfigWithName(name string) (*resources.AWSEMRInstanceFleetConfig, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSEMRInstanceFleetConfig:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEMRInstanceFleetConfig not found", name)
}

// GetAllAWSEMRInstanceGroupConfigResources retrieves all AWSEMRInstanceGroupConfig items from an AWS CloudFormation template
func (t *Template) GetAllAWSEMRInstanceGroupConfigResources() map[string]*resources.AWSEMRInstanceGroupConfig {
	results := map[string]*resources.AWSEMRInstanceGroupConfig{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSEMRInstanceGroupConfig:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEMRInstanceGroupConfigWithName retrieves all AWSEMRInstanceGroupConfig items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEMRInstanceGroupConfigWithName(name string) (*resources.AWSEMRInstanceGroupConfig, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSEMRInstanceGroupConfig:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEMRInstanceGroupConfig not found", name)
}

// GetAllAWSEMRSecurityConfigurationResources retrieves all AWSEMRSecurityConfiguration items from an AWS CloudFormation template
func (t *Template) GetAllAWSEMRSecurityConfigurationResources() map[string]*resources.AWSEMRSecurityConfiguration {
	results := map[string]*resources.AWSEMRSecurityConfiguration{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSEMRSecurityConfiguration:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEMRSecurityConfigurationWithName retrieves all AWSEMRSecurityConfiguration items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEMRSecurityConfigurationWithName(name string) (*resources.AWSEMRSecurityConfiguration, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSEMRSecurityConfiguration:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEMRSecurityConfiguration not found", name)
}

// GetAllAWSEMRStepResources retrieves all AWSEMRStep items from an AWS CloudFormation template
func (t *Template) GetAllAWSEMRStepResources() map[string]*resources.AWSEMRStep {
	results := map[string]*resources.AWSEMRStep{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSEMRStep:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEMRStepWithName retrieves all AWSEMRStep items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEMRStepWithName(name string) (*resources.AWSEMRStep, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSEMRStep:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEMRStep not found", name)
}

// GetAllAWSElastiCacheCacheClusterResources retrieves all AWSElastiCacheCacheCluster items from an AWS CloudFormation template
func (t *Template) GetAllAWSElastiCacheCacheClusterResources() map[string]*resources.AWSElastiCacheCacheCluster {
	results := map[string]*resources.AWSElastiCacheCacheCluster{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSElastiCacheCacheCluster:
			results[name] = resource
		}
	}
	return results
}

// GetAWSElastiCacheCacheClusterWithName retrieves all AWSElastiCacheCacheCluster items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSElastiCacheCacheClusterWithName(name string) (*resources.AWSElastiCacheCacheCluster, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSElastiCacheCacheCluster:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSElastiCacheCacheCluster not found", name)
}

// GetAllAWSElastiCacheParameterGroupResources retrieves all AWSElastiCacheParameterGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSElastiCacheParameterGroupResources() map[string]*resources.AWSElastiCacheParameterGroup {
	results := map[string]*resources.AWSElastiCacheParameterGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSElastiCacheParameterGroup:
			results[name] = resource
		}
	}
	return results
}

// GetAWSElastiCacheParameterGroupWithName retrieves all AWSElastiCacheParameterGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSElastiCacheParameterGroupWithName(name string) (*resources.AWSElastiCacheParameterGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSElastiCacheParameterGroup:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSElastiCacheParameterGroup not found", name)
}

// GetAllAWSElastiCacheReplicationGroupResources retrieves all AWSElastiCacheReplicationGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSElastiCacheReplicationGroupResources() map[string]*resources.AWSElastiCacheReplicationGroup {
	results := map[string]*resources.AWSElastiCacheReplicationGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSElastiCacheReplicationGroup:
			results[name] = resource
		}
	}
	return results
}

// GetAWSElastiCacheReplicationGroupWithName retrieves all AWSElastiCacheReplicationGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSElastiCacheReplicationGroupWithName(name string) (*resources.AWSElastiCacheReplicationGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSElastiCacheReplicationGroup:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSElastiCacheReplicationGroup not found", name)
}

// GetAllAWSElastiCacheSecurityGroupResources retrieves all AWSElastiCacheSecurityGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSElastiCacheSecurityGroupResources() map[string]*resources.AWSElastiCacheSecurityGroup {
	results := map[string]*resources.AWSElastiCacheSecurityGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSElastiCacheSecurityGroup:
			results[name] = resource
		}
	}
	return results
}

// GetAWSElastiCacheSecurityGroupWithName retrieves all AWSElastiCacheSecurityGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSElastiCacheSecurityGroupWithName(name string) (*resources.AWSElastiCacheSecurityGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSElastiCacheSecurityGroup:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSElastiCacheSecurityGroup not found", name)
}

// GetAllAWSElastiCacheSecurityGroupIngressResources retrieves all AWSElastiCacheSecurityGroupIngress items from an AWS CloudFormation template
func (t *Template) GetAllAWSElastiCacheSecurityGroupIngressResources() map[string]*resources.AWSElastiCacheSecurityGroupIngress {
	results := map[string]*resources.AWSElastiCacheSecurityGroupIngress{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSElastiCacheSecurityGroupIngress:
			results[name] = resource
		}
	}
	return results
}

// GetAWSElastiCacheSecurityGroupIngressWithName retrieves all AWSElastiCacheSecurityGroupIngress items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSElastiCacheSecurityGroupIngressWithName(name string) (*resources.AWSElastiCacheSecurityGroupIngress, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSElastiCacheSecurityGroupIngress:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSElastiCacheSecurityGroupIngress not found", name)
}

// GetAllAWSElastiCacheSubnetGroupResources retrieves all AWSElastiCacheSubnetGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSElastiCacheSubnetGroupResources() map[string]*resources.AWSElastiCacheSubnetGroup {
	results := map[string]*resources.AWSElastiCacheSubnetGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSElastiCacheSubnetGroup:
			results[name] = resource
		}
	}
	return results
}

// GetAWSElastiCacheSubnetGroupWithName retrieves all AWSElastiCacheSubnetGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSElastiCacheSubnetGroupWithName(name string) (*resources.AWSElastiCacheSubnetGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSElastiCacheSubnetGroup:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSElastiCacheSubnetGroup not found", name)
}

// GetAllAWSElasticBeanstalkApplicationResources retrieves all AWSElasticBeanstalkApplication items from an AWS CloudFormation template
func (t *Template) GetAllAWSElasticBeanstalkApplicationResources() map[string]*resources.AWSElasticBeanstalkApplication {
	results := map[string]*resources.AWSElasticBeanstalkApplication{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSElasticBeanstalkApplication:
			results[name] = resource
		}
	}
	return results
}

// GetAWSElasticBeanstalkApplicationWithName retrieves all AWSElasticBeanstalkApplication items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSElasticBeanstalkApplicationWithName(name string) (*resources.AWSElasticBeanstalkApplication, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSElasticBeanstalkApplication:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSElasticBeanstalkApplication not found", name)
}

// GetAllAWSElasticBeanstalkApplicationVersionResources retrieves all AWSElasticBeanstalkApplicationVersion items from an AWS CloudFormation template
func (t *Template) GetAllAWSElasticBeanstalkApplicationVersionResources() map[string]*resources.AWSElasticBeanstalkApplicationVersion {
	results := map[string]*resources.AWSElasticBeanstalkApplicationVersion{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSElasticBeanstalkApplicationVersion:
			results[name] = resource
		}
	}
	return results
}

// GetAWSElasticBeanstalkApplicationVersionWithName retrieves all AWSElasticBeanstalkApplicationVersion items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSElasticBeanstalkApplicationVersionWithName(name string) (*resources.AWSElasticBeanstalkApplicationVersion, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSElasticBeanstalkApplicationVersion:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSElasticBeanstalkApplicationVersion not found", name)
}

// GetAllAWSElasticBeanstalkConfigurationTemplateResources retrieves all AWSElasticBeanstalkConfigurationTemplate items from an AWS CloudFormation template
func (t *Template) GetAllAWSElasticBeanstalkConfigurationTemplateResources() map[string]*resources.AWSElasticBeanstalkConfigurationTemplate {
	results := map[string]*resources.AWSElasticBeanstalkConfigurationTemplate{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSElasticBeanstalkConfigurationTemplate:
			results[name] = resource
		}
	}
	return results
}

// GetAWSElasticBeanstalkConfigurationTemplateWithName retrieves all AWSElasticBeanstalkConfigurationTemplate items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSElasticBeanstalkConfigurationTemplateWithName(name string) (*resources.AWSElasticBeanstalkConfigurationTemplate, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSElasticBeanstalkConfigurationTemplate:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSElasticBeanstalkConfigurationTemplate not found", name)
}

// GetAllAWSElasticBeanstalkEnvironmentResources retrieves all AWSElasticBeanstalkEnvironment items from an AWS CloudFormation template
func (t *Template) GetAllAWSElasticBeanstalkEnvironmentResources() map[string]*resources.AWSElasticBeanstalkEnvironment {
	results := map[string]*resources.AWSElasticBeanstalkEnvironment{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSElasticBeanstalkEnvironment:
			results[name] = resource
		}
	}
	return results
}

// GetAWSElasticBeanstalkEnvironmentWithName retrieves all AWSElasticBeanstalkEnvironment items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSElasticBeanstalkEnvironmentWithName(name string) (*resources.AWSElasticBeanstalkEnvironment, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSElasticBeanstalkEnvironment:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSElasticBeanstalkEnvironment not found", name)
}

// GetAllAWSElasticLoadBalancingLoadBalancerResources retrieves all AWSElasticLoadBalancingLoadBalancer items from an AWS CloudFormation template
func (t *Template) GetAllAWSElasticLoadBalancingLoadBalancerResources() map[string]*resources.AWSElasticLoadBalancingLoadBalancer {
	results := map[string]*resources.AWSElasticLoadBalancingLoadBalancer{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSElasticLoadBalancingLoadBalancer:
			results[name] = resource
		}
	}
	return results
}

// GetAWSElasticLoadBalancingLoadBalancerWithName retrieves all AWSElasticLoadBalancingLoadBalancer items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSElasticLoadBalancingLoadBalancerWithName(name string) (*resources.AWSElasticLoadBalancingLoadBalancer, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSElasticLoadBalancingLoadBalancer:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSElasticLoadBalancingLoadBalancer not found", name)
}

// GetAllAWSElasticLoadBalancingV2ListenerResources retrieves all AWSElasticLoadBalancingV2Listener items from an AWS CloudFormation template
func (t *Template) GetAllAWSElasticLoadBalancingV2ListenerResources() map[string]*resources.AWSElasticLoadBalancingV2Listener {
	results := map[string]*resources.AWSElasticLoadBalancingV2Listener{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSElasticLoadBalancingV2Listener:
			results[name] = resource
		}
	}
	return results
}

// GetAWSElasticLoadBalancingV2ListenerWithName retrieves all AWSElasticLoadBalancingV2Listener items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSElasticLoadBalancingV2ListenerWithName(name string) (*resources.AWSElasticLoadBalancingV2Listener, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSElasticLoadBalancingV2Listener:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSElasticLoadBalancingV2Listener not found", name)
}

// GetAllAWSElasticLoadBalancingV2ListenerCertificateResources retrieves all AWSElasticLoadBalancingV2ListenerCertificate items from an AWS CloudFormation template
func (t *Template) GetAllAWSElasticLoadBalancingV2ListenerCertificateResources() map[string]*resources.AWSElasticLoadBalancingV2ListenerCertificate {
	results := map[string]*resources.AWSElasticLoadBalancingV2ListenerCertificate{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSElasticLoadBalancingV2ListenerCertificate:
			results[name] = resource
		}
	}
	return results
}

// GetAWSElasticLoadBalancingV2ListenerCertificateWithName retrieves all AWSElasticLoadBalancingV2ListenerCertificate items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSElasticLoadBalancingV2ListenerCertificateWithName(name string) (*resources.AWSElasticLoadBalancingV2ListenerCertificate, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSElasticLoadBalancingV2ListenerCertificate:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSElasticLoadBalancingV2ListenerCertificate not found", name)
}

// GetAllAWSElasticLoadBalancingV2ListenerRuleResources retrieves all AWSElasticLoadBalancingV2ListenerRule items from an AWS CloudFormation template
func (t *Template) GetAllAWSElasticLoadBalancingV2ListenerRuleResources() map[string]*resources.AWSElasticLoadBalancingV2ListenerRule {
	results := map[string]*resources.AWSElasticLoadBalancingV2ListenerRule{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSElasticLoadBalancingV2ListenerRule:
			results[name] = resource
		}
	}
	return results
}

// GetAWSElasticLoadBalancingV2ListenerRuleWithName retrieves all AWSElasticLoadBalancingV2ListenerRule items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSElasticLoadBalancingV2ListenerRuleWithName(name string) (*resources.AWSElasticLoadBalancingV2ListenerRule, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSElasticLoadBalancingV2ListenerRule:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSElasticLoadBalancingV2ListenerRule not found", name)
}

// GetAllAWSElasticLoadBalancingV2LoadBalancerResources retrieves all AWSElasticLoadBalancingV2LoadBalancer items from an AWS CloudFormation template
func (t *Template) GetAllAWSElasticLoadBalancingV2LoadBalancerResources() map[string]*resources.AWSElasticLoadBalancingV2LoadBalancer {
	results := map[string]*resources.AWSElasticLoadBalancingV2LoadBalancer{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSElasticLoadBalancingV2LoadBalancer:
			results[name] = resource
		}
	}
	return results
}

// GetAWSElasticLoadBalancingV2LoadBalancerWithName retrieves all AWSElasticLoadBalancingV2LoadBalancer items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSElasticLoadBalancingV2LoadBalancerWithName(name string) (*resources.AWSElasticLoadBalancingV2LoadBalancer, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSElasticLoadBalancingV2LoadBalancer:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSElasticLoadBalancingV2LoadBalancer not found", name)
}

// GetAllAWSElasticLoadBalancingV2TargetGroupResources retrieves all AWSElasticLoadBalancingV2TargetGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSElasticLoadBalancingV2TargetGroupResources() map[string]*resources.AWSElasticLoadBalancingV2TargetGroup {
	results := map[string]*resources.AWSElasticLoadBalancingV2TargetGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSElasticLoadBalancingV2TargetGroup:
			results[name] = resource
		}
	}
	return results
}

// GetAWSElasticLoadBalancingV2TargetGroupWithName retrieves all AWSElasticLoadBalancingV2TargetGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSElasticLoadBalancingV2TargetGroupWithName(name string) (*resources.AWSElasticLoadBalancingV2TargetGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSElasticLoadBalancingV2TargetGroup:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSElasticLoadBalancingV2TargetGroup not found", name)
}

// GetAllAWSElasticsearchDomainResources retrieves all AWSElasticsearchDomain items from an AWS CloudFormation template
func (t *Template) GetAllAWSElasticsearchDomainResources() map[string]*resources.AWSElasticsearchDomain {
	results := map[string]*resources.AWSElasticsearchDomain{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSElasticsearchDomain:
			results[name] = resource
		}
	}
	return results
}

// GetAWSElasticsearchDomainWithName retrieves all AWSElasticsearchDomain items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSElasticsearchDomainWithName(name string) (*resources.AWSElasticsearchDomain, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSElasticsearchDomain:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSElasticsearchDomain not found", name)
}

// GetAllAWSEventsEventBusPolicyResources retrieves all AWSEventsEventBusPolicy items from an AWS CloudFormation template
func (t *Template) GetAllAWSEventsEventBusPolicyResources() map[string]*resources.AWSEventsEventBusPolicy {
	results := map[string]*resources.AWSEventsEventBusPolicy{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSEventsEventBusPolicy:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEventsEventBusPolicyWithName retrieves all AWSEventsEventBusPolicy items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEventsEventBusPolicyWithName(name string) (*resources.AWSEventsEventBusPolicy, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSEventsEventBusPolicy:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEventsEventBusPolicy not found", name)
}

// GetAllAWSEventsRuleResources retrieves all AWSEventsRule items from an AWS CloudFormation template
func (t *Template) GetAllAWSEventsRuleResources() map[string]*resources.AWSEventsRule {
	results := map[string]*resources.AWSEventsRule{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSEventsRule:
			results[name] = resource
		}
	}
	return results
}

// GetAWSEventsRuleWithName retrieves all AWSEventsRule items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEventsRuleWithName(name string) (*resources.AWSEventsRule, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSEventsRule:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSEventsRule not found", name)
}

// GetAllAWSFSxFileSystemResources retrieves all AWSFSxFileSystem items from an AWS CloudFormation template
func (t *Template) GetAllAWSFSxFileSystemResources() map[string]*resources.AWSFSxFileSystem {
	results := map[string]*resources.AWSFSxFileSystem{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSFSxFileSystem:
			results[name] = resource
		}
	}
	return results
}

// GetAWSFSxFileSystemWithName retrieves all AWSFSxFileSystem items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSFSxFileSystemWithName(name string) (*resources.AWSFSxFileSystem, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSFSxFileSystem:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSFSxFileSystem not found", name)
}

// GetAllAWSGameLiftAliasResources retrieves all AWSGameLiftAlias items from an AWS CloudFormation template
func (t *Template) GetAllAWSGameLiftAliasResources() map[string]*resources.AWSGameLiftAlias {
	results := map[string]*resources.AWSGameLiftAlias{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSGameLiftAlias:
			results[name] = resource
		}
	}
	return results
}

// GetAWSGameLiftAliasWithName retrieves all AWSGameLiftAlias items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSGameLiftAliasWithName(name string) (*resources.AWSGameLiftAlias, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSGameLiftAlias:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSGameLiftAlias not found", name)
}

// GetAllAWSGameLiftBuildResources retrieves all AWSGameLiftBuild items from an AWS CloudFormation template
func (t *Template) GetAllAWSGameLiftBuildResources() map[string]*resources.AWSGameLiftBuild {
	results := map[string]*resources.AWSGameLiftBuild{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSGameLiftBuild:
			results[name] = resource
		}
	}
	return results
}

// GetAWSGameLiftBuildWithName retrieves all AWSGameLiftBuild items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSGameLiftBuildWithName(name string) (*resources.AWSGameLiftBuild, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSGameLiftBuild:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSGameLiftBuild not found", name)
}

// GetAllAWSGameLiftFleetResources retrieves all AWSGameLiftFleet items from an AWS CloudFormation template
func (t *Template) GetAllAWSGameLiftFleetResources() map[string]*resources.AWSGameLiftFleet {
	results := map[string]*resources.AWSGameLiftFleet{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSGameLiftFleet:
			results[name] = resource
		}
	}
	return results
}

// GetAWSGameLiftFleetWithName retrieves all AWSGameLiftFleet items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSGameLiftFleetWithName(name string) (*resources.AWSGameLiftFleet, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSGameLiftFleet:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSGameLiftFleet not found", name)
}

// GetAllAWSGlueClassifierResources retrieves all AWSGlueClassifier items from an AWS CloudFormation template
func (t *Template) GetAllAWSGlueClassifierResources() map[string]*resources.AWSGlueClassifier {
	results := map[string]*resources.AWSGlueClassifier{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSGlueClassifier:
			results[name] = resource
		}
	}
	return results
}

// GetAWSGlueClassifierWithName retrieves all AWSGlueClassifier items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSGlueClassifierWithName(name string) (*resources.AWSGlueClassifier, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSGlueClassifier:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSGlueClassifier not found", name)
}

// GetAllAWSGlueConnectionResources retrieves all AWSGlueConnection items from an AWS CloudFormation template
func (t *Template) GetAllAWSGlueConnectionResources() map[string]*resources.AWSGlueConnection {
	results := map[string]*resources.AWSGlueConnection{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSGlueConnection:
			results[name] = resource
		}
	}
	return results
}

// GetAWSGlueConnectionWithName retrieves all AWSGlueConnection items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSGlueConnectionWithName(name string) (*resources.AWSGlueConnection, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSGlueConnection:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSGlueConnection not found", name)
}

// GetAllAWSGlueCrawlerResources retrieves all AWSGlueCrawler items from an AWS CloudFormation template
func (t *Template) GetAllAWSGlueCrawlerResources() map[string]*resources.AWSGlueCrawler {
	results := map[string]*resources.AWSGlueCrawler{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSGlueCrawler:
			results[name] = resource
		}
	}
	return results
}

// GetAWSGlueCrawlerWithName retrieves all AWSGlueCrawler items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSGlueCrawlerWithName(name string) (*resources.AWSGlueCrawler, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSGlueCrawler:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSGlueCrawler not found", name)
}

// GetAllAWSGlueDatabaseResources retrieves all AWSGlueDatabase items from an AWS CloudFormation template
func (t *Template) GetAllAWSGlueDatabaseResources() map[string]*resources.AWSGlueDatabase {
	results := map[string]*resources.AWSGlueDatabase{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSGlueDatabase:
			results[name] = resource
		}
	}
	return results
}

// GetAWSGlueDatabaseWithName retrieves all AWSGlueDatabase items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSGlueDatabaseWithName(name string) (*resources.AWSGlueDatabase, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSGlueDatabase:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSGlueDatabase not found", name)
}

// GetAllAWSGlueDevEndpointResources retrieves all AWSGlueDevEndpoint items from an AWS CloudFormation template
func (t *Template) GetAllAWSGlueDevEndpointResources() map[string]*resources.AWSGlueDevEndpoint {
	results := map[string]*resources.AWSGlueDevEndpoint{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSGlueDevEndpoint:
			results[name] = resource
		}
	}
	return results
}

// GetAWSGlueDevEndpointWithName retrieves all AWSGlueDevEndpoint items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSGlueDevEndpointWithName(name string) (*resources.AWSGlueDevEndpoint, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSGlueDevEndpoint:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSGlueDevEndpoint not found", name)
}

// GetAllAWSGlueJobResources retrieves all AWSGlueJob items from an AWS CloudFormation template
func (t *Template) GetAllAWSGlueJobResources() map[string]*resources.AWSGlueJob {
	results := map[string]*resources.AWSGlueJob{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSGlueJob:
			results[name] = resource
		}
	}
	return results
}

// GetAWSGlueJobWithName retrieves all AWSGlueJob items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSGlueJobWithName(name string) (*resources.AWSGlueJob, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSGlueJob:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSGlueJob not found", name)
}

// GetAllAWSGluePartitionResources retrieves all AWSGluePartition items from an AWS CloudFormation template
func (t *Template) GetAllAWSGluePartitionResources() map[string]*resources.AWSGluePartition {
	results := map[string]*resources.AWSGluePartition{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSGluePartition:
			results[name] = resource
		}
	}
	return results
}

// GetAWSGluePartitionWithName retrieves all AWSGluePartition items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSGluePartitionWithName(name string) (*resources.AWSGluePartition, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSGluePartition:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSGluePartition not found", name)
}

// GetAllAWSGlueTableResources retrieves all AWSGlueTable items from an AWS CloudFormation template
func (t *Template) GetAllAWSGlueTableResources() map[string]*resources.AWSGlueTable {
	results := map[string]*resources.AWSGlueTable{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSGlueTable:
			results[name] = resource
		}
	}
	return results
}

// GetAWSGlueTableWithName retrieves all AWSGlueTable items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSGlueTableWithName(name string) (*resources.AWSGlueTable, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSGlueTable:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSGlueTable not found", name)
}

// GetAllAWSGlueTriggerResources retrieves all AWSGlueTrigger items from an AWS CloudFormation template
func (t *Template) GetAllAWSGlueTriggerResources() map[string]*resources.AWSGlueTrigger {
	results := map[string]*resources.AWSGlueTrigger{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSGlueTrigger:
			results[name] = resource
		}
	}
	return results
}

// GetAWSGlueTriggerWithName retrieves all AWSGlueTrigger items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSGlueTriggerWithName(name string) (*resources.AWSGlueTrigger, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSGlueTrigger:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSGlueTrigger not found", name)
}

// GetAllAWSGuardDutyDetectorResources retrieves all AWSGuardDutyDetector items from an AWS CloudFormation template
func (t *Template) GetAllAWSGuardDutyDetectorResources() map[string]*resources.AWSGuardDutyDetector {
	results := map[string]*resources.AWSGuardDutyDetector{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSGuardDutyDetector:
			results[name] = resource
		}
	}
	return results
}

// GetAWSGuardDutyDetectorWithName retrieves all AWSGuardDutyDetector items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSGuardDutyDetectorWithName(name string) (*resources.AWSGuardDutyDetector, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSGuardDutyDetector:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSGuardDutyDetector not found", name)
}

// GetAllAWSGuardDutyFilterResources retrieves all AWSGuardDutyFilter items from an AWS CloudFormation template
func (t *Template) GetAllAWSGuardDutyFilterResources() map[string]*resources.AWSGuardDutyFilter {
	results := map[string]*resources.AWSGuardDutyFilter{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSGuardDutyFilter:
			results[name] = resource
		}
	}
	return results
}

// GetAWSGuardDutyFilterWithName retrieves all AWSGuardDutyFilter items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSGuardDutyFilterWithName(name string) (*resources.AWSGuardDutyFilter, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSGuardDutyFilter:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSGuardDutyFilter not found", name)
}

// GetAllAWSGuardDutyIPSetResources retrieves all AWSGuardDutyIPSet items from an AWS CloudFormation template
func (t *Template) GetAllAWSGuardDutyIPSetResources() map[string]*resources.AWSGuardDutyIPSet {
	results := map[string]*resources.AWSGuardDutyIPSet{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSGuardDutyIPSet:
			results[name] = resource
		}
	}
	return results
}

// GetAWSGuardDutyIPSetWithName retrieves all AWSGuardDutyIPSet items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSGuardDutyIPSetWithName(name string) (*resources.AWSGuardDutyIPSet, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSGuardDutyIPSet:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSGuardDutyIPSet not found", name)
}

// GetAllAWSGuardDutyMasterResources retrieves all AWSGuardDutyMaster items from an AWS CloudFormation template
func (t *Template) GetAllAWSGuardDutyMasterResources() map[string]*resources.AWSGuardDutyMaster {
	results := map[string]*resources.AWSGuardDutyMaster{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSGuardDutyMaster:
			results[name] = resource
		}
	}
	return results
}

// GetAWSGuardDutyMasterWithName retrieves all AWSGuardDutyMaster items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSGuardDutyMasterWithName(name string) (*resources.AWSGuardDutyMaster, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSGuardDutyMaster:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSGuardDutyMaster not found", name)
}

// GetAllAWSGuardDutyMemberResources retrieves all AWSGuardDutyMember items from an AWS CloudFormation template
func (t *Template) GetAllAWSGuardDutyMemberResources() map[string]*resources.AWSGuardDutyMember {
	results := map[string]*resources.AWSGuardDutyMember{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSGuardDutyMember:
			results[name] = resource
		}
	}
	return results
}

// GetAWSGuardDutyMemberWithName retrieves all AWSGuardDutyMember items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSGuardDutyMemberWithName(name string) (*resources.AWSGuardDutyMember, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSGuardDutyMember:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSGuardDutyMember not found", name)
}

// GetAllAWSGuardDutyThreatIntelSetResources retrieves all AWSGuardDutyThreatIntelSet items from an AWS CloudFormation template
func (t *Template) GetAllAWSGuardDutyThreatIntelSetResources() map[string]*resources.AWSGuardDutyThreatIntelSet {
	results := map[string]*resources.AWSGuardDutyThreatIntelSet{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSGuardDutyThreatIntelSet:
			results[name] = resource
		}
	}
	return results
}

// GetAWSGuardDutyThreatIntelSetWithName retrieves all AWSGuardDutyThreatIntelSet items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSGuardDutyThreatIntelSetWithName(name string) (*resources.AWSGuardDutyThreatIntelSet, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSGuardDutyThreatIntelSet:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSGuardDutyThreatIntelSet not found", name)
}

// GetAllAWSIAMAccessKeyResources retrieves all AWSIAMAccessKey items from an AWS CloudFormation template
func (t *Template) GetAllAWSIAMAccessKeyResources() map[string]*resources.AWSIAMAccessKey {
	results := map[string]*resources.AWSIAMAccessKey{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSIAMAccessKey:
			results[name] = resource
		}
	}
	return results
}

// GetAWSIAMAccessKeyWithName retrieves all AWSIAMAccessKey items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSIAMAccessKeyWithName(name string) (*resources.AWSIAMAccessKey, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSIAMAccessKey:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSIAMAccessKey not found", name)
}

// GetAllAWSIAMGroupResources retrieves all AWSIAMGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSIAMGroupResources() map[string]*resources.AWSIAMGroup {
	results := map[string]*resources.AWSIAMGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSIAMGroup:
			results[name] = resource
		}
	}
	return results
}

// GetAWSIAMGroupWithName retrieves all AWSIAMGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSIAMGroupWithName(name string) (*resources.AWSIAMGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSIAMGroup:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSIAMGroup not found", name)
}

// GetAllAWSIAMInstanceProfileResources retrieves all AWSIAMInstanceProfile items from an AWS CloudFormation template
func (t *Template) GetAllAWSIAMInstanceProfileResources() map[string]*resources.AWSIAMInstanceProfile {
	results := map[string]*resources.AWSIAMInstanceProfile{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSIAMInstanceProfile:
			results[name] = resource
		}
	}
	return results
}

// GetAWSIAMInstanceProfileWithName retrieves all AWSIAMInstanceProfile items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSIAMInstanceProfileWithName(name string) (*resources.AWSIAMInstanceProfile, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSIAMInstanceProfile:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSIAMInstanceProfile not found", name)
}

// GetAllAWSIAMManagedPolicyResources retrieves all AWSIAMManagedPolicy items from an AWS CloudFormation template
func (t *Template) GetAllAWSIAMManagedPolicyResources() map[string]*resources.AWSIAMManagedPolicy {
	results := map[string]*resources.AWSIAMManagedPolicy{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSIAMManagedPolicy:
			results[name] = resource
		}
	}
	return results
}

// GetAWSIAMManagedPolicyWithName retrieves all AWSIAMManagedPolicy items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSIAMManagedPolicyWithName(name string) (*resources.AWSIAMManagedPolicy, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSIAMManagedPolicy:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSIAMManagedPolicy not found", name)
}

// GetAllAWSIAMPolicyResources retrieves all AWSIAMPolicy items from an AWS CloudFormation template
func (t *Template) GetAllAWSIAMPolicyResources() map[string]*resources.AWSIAMPolicy {
	results := map[string]*resources.AWSIAMPolicy{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSIAMPolicy:
			results[name] = resource
		}
	}
	return results
}

// GetAWSIAMPolicyWithName retrieves all AWSIAMPolicy items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSIAMPolicyWithName(name string) (*resources.AWSIAMPolicy, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSIAMPolicy:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSIAMPolicy not found", name)
}

// GetAllAWSIAMRoleResources retrieves all AWSIAMRole items from an AWS CloudFormation template
func (t *Template) GetAllAWSIAMRoleResources() map[string]*resources.AWSIAMRole {
	results := map[string]*resources.AWSIAMRole{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSIAMRole:
			results[name] = resource
		}
	}
	return results
}

// GetAWSIAMRoleWithName retrieves all AWSIAMRole items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSIAMRoleWithName(name string) (*resources.AWSIAMRole, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSIAMRole:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSIAMRole not found", name)
}

// GetAllAWSIAMServiceLinkedRoleResources retrieves all AWSIAMServiceLinkedRole items from an AWS CloudFormation template
func (t *Template) GetAllAWSIAMServiceLinkedRoleResources() map[string]*resources.AWSIAMServiceLinkedRole {
	results := map[string]*resources.AWSIAMServiceLinkedRole{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSIAMServiceLinkedRole:
			results[name] = resource
		}
	}
	return results
}

// GetAWSIAMServiceLinkedRoleWithName retrieves all AWSIAMServiceLinkedRole items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSIAMServiceLinkedRoleWithName(name string) (*resources.AWSIAMServiceLinkedRole, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSIAMServiceLinkedRole:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSIAMServiceLinkedRole not found", name)
}

// GetAllAWSIAMUserResources retrieves all AWSIAMUser items from an AWS CloudFormation template
func (t *Template) GetAllAWSIAMUserResources() map[string]*resources.AWSIAMUser {
	results := map[string]*resources.AWSIAMUser{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSIAMUser:
			results[name] = resource
		}
	}
	return results
}

// GetAWSIAMUserWithName retrieves all AWSIAMUser items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSIAMUserWithName(name string) (*resources.AWSIAMUser, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSIAMUser:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSIAMUser not found", name)
}

// GetAllAWSIAMUserToGroupAdditionResources retrieves all AWSIAMUserToGroupAddition items from an AWS CloudFormation template
func (t *Template) GetAllAWSIAMUserToGroupAdditionResources() map[string]*resources.AWSIAMUserToGroupAddition {
	results := map[string]*resources.AWSIAMUserToGroupAddition{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSIAMUserToGroupAddition:
			results[name] = resource
		}
	}
	return results
}

// GetAWSIAMUserToGroupAdditionWithName retrieves all AWSIAMUserToGroupAddition items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSIAMUserToGroupAdditionWithName(name string) (*resources.AWSIAMUserToGroupAddition, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSIAMUserToGroupAddition:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSIAMUserToGroupAddition not found", name)
}

// GetAllAWSInspectorAssessmentTargetResources retrieves all AWSInspectorAssessmentTarget items from an AWS CloudFormation template
func (t *Template) GetAllAWSInspectorAssessmentTargetResources() map[string]*resources.AWSInspectorAssessmentTarget {
	results := map[string]*resources.AWSInspectorAssessmentTarget{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSInspectorAssessmentTarget:
			results[name] = resource
		}
	}
	return results
}

// GetAWSInspectorAssessmentTargetWithName retrieves all AWSInspectorAssessmentTarget items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSInspectorAssessmentTargetWithName(name string) (*resources.AWSInspectorAssessmentTarget, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSInspectorAssessmentTarget:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSInspectorAssessmentTarget not found", name)
}

// GetAllAWSInspectorAssessmentTemplateResources retrieves all AWSInspectorAssessmentTemplate items from an AWS CloudFormation template
func (t *Template) GetAllAWSInspectorAssessmentTemplateResources() map[string]*resources.AWSInspectorAssessmentTemplate {
	results := map[string]*resources.AWSInspectorAssessmentTemplate{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSInspectorAssessmentTemplate:
			results[name] = resource
		}
	}
	return results
}

// GetAWSInspectorAssessmentTemplateWithName retrieves all AWSInspectorAssessmentTemplate items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSInspectorAssessmentTemplateWithName(name string) (*resources.AWSInspectorAssessmentTemplate, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSInspectorAssessmentTemplate:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSInspectorAssessmentTemplate not found", name)
}

// GetAllAWSInspectorResourceGroupResources retrieves all AWSInspectorResourceGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSInspectorResourceGroupResources() map[string]*resources.AWSInspectorResourceGroup {
	results := map[string]*resources.AWSInspectorResourceGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSInspectorResourceGroup:
			results[name] = resource
		}
	}
	return results
}

// GetAWSInspectorResourceGroupWithName retrieves all AWSInspectorResourceGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSInspectorResourceGroupWithName(name string) (*resources.AWSInspectorResourceGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSInspectorResourceGroup:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSInspectorResourceGroup not found", name)
}

// GetAllAWSIoT1ClickDeviceResources retrieves all AWSIoT1ClickDevice items from an AWS CloudFormation template
func (t *Template) GetAllAWSIoT1ClickDeviceResources() map[string]*resources.AWSIoT1ClickDevice {
	results := map[string]*resources.AWSIoT1ClickDevice{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSIoT1ClickDevice:
			results[name] = resource
		}
	}
	return results
}

// GetAWSIoT1ClickDeviceWithName retrieves all AWSIoT1ClickDevice items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSIoT1ClickDeviceWithName(name string) (*resources.AWSIoT1ClickDevice, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSIoT1ClickDevice:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSIoT1ClickDevice not found", name)
}

// GetAllAWSIoT1ClickPlacementResources retrieves all AWSIoT1ClickPlacement items from an AWS CloudFormation template
func (t *Template) GetAllAWSIoT1ClickPlacementResources() map[string]*resources.AWSIoT1ClickPlacement {
	results := map[string]*resources.AWSIoT1ClickPlacement{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSIoT1ClickPlacement:
			results[name] = resource
		}
	}
	return results
}

// GetAWSIoT1ClickPlacementWithName retrieves all AWSIoT1ClickPlacement items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSIoT1ClickPlacementWithName(name string) (*resources.AWSIoT1ClickPlacement, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSIoT1ClickPlacement:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSIoT1ClickPlacement not found", name)
}

// GetAllAWSIoT1ClickProjectResources retrieves all AWSIoT1ClickProject items from an AWS CloudFormation template
func (t *Template) GetAllAWSIoT1ClickProjectResources() map[string]*resources.AWSIoT1ClickProject {
	results := map[string]*resources.AWSIoT1ClickProject{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSIoT1ClickProject:
			results[name] = resource
		}
	}
	return results
}

// GetAWSIoT1ClickProjectWithName retrieves all AWSIoT1ClickProject items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSIoT1ClickProjectWithName(name string) (*resources.AWSIoT1ClickProject, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSIoT1ClickProject:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSIoT1ClickProject not found", name)
}

// GetAllAWSIoTCertificateResources retrieves all AWSIoTCertificate items from an AWS CloudFormation template
func (t *Template) GetAllAWSIoTCertificateResources() map[string]*resources.AWSIoTCertificate {
	results := map[string]*resources.AWSIoTCertificate{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSIoTCertificate:
			results[name] = resource
		}
	}
	return results
}

// GetAWSIoTCertificateWithName retrieves all AWSIoTCertificate items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSIoTCertificateWithName(name string) (*resources.AWSIoTCertificate, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSIoTCertificate:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSIoTCertificate not found", name)
}

// GetAllAWSIoTPolicyResources retrieves all AWSIoTPolicy items from an AWS CloudFormation template
func (t *Template) GetAllAWSIoTPolicyResources() map[string]*resources.AWSIoTPolicy {
	results := map[string]*resources.AWSIoTPolicy{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSIoTPolicy:
			results[name] = resource
		}
	}
	return results
}

// GetAWSIoTPolicyWithName retrieves all AWSIoTPolicy items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSIoTPolicyWithName(name string) (*resources.AWSIoTPolicy, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSIoTPolicy:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSIoTPolicy not found", name)
}

// GetAllAWSIoTPolicyPrincipalAttachmentResources retrieves all AWSIoTPolicyPrincipalAttachment items from an AWS CloudFormation template
func (t *Template) GetAllAWSIoTPolicyPrincipalAttachmentResources() map[string]*resources.AWSIoTPolicyPrincipalAttachment {
	results := map[string]*resources.AWSIoTPolicyPrincipalAttachment{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSIoTPolicyPrincipalAttachment:
			results[name] = resource
		}
	}
	return results
}

// GetAWSIoTPolicyPrincipalAttachmentWithName retrieves all AWSIoTPolicyPrincipalAttachment items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSIoTPolicyPrincipalAttachmentWithName(name string) (*resources.AWSIoTPolicyPrincipalAttachment, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSIoTPolicyPrincipalAttachment:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSIoTPolicyPrincipalAttachment not found", name)
}

// GetAllAWSIoTThingResources retrieves all AWSIoTThing items from an AWS CloudFormation template
func (t *Template) GetAllAWSIoTThingResources() map[string]*resources.AWSIoTThing {
	results := map[string]*resources.AWSIoTThing{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSIoTThing:
			results[name] = resource
		}
	}
	return results
}

// GetAWSIoTThingWithName retrieves all AWSIoTThing items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSIoTThingWithName(name string) (*resources.AWSIoTThing, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSIoTThing:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSIoTThing not found", name)
}

// GetAllAWSIoTThingPrincipalAttachmentResources retrieves all AWSIoTThingPrincipalAttachment items from an AWS CloudFormation template
func (t *Template) GetAllAWSIoTThingPrincipalAttachmentResources() map[string]*resources.AWSIoTThingPrincipalAttachment {
	results := map[string]*resources.AWSIoTThingPrincipalAttachment{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSIoTThingPrincipalAttachment:
			results[name] = resource
		}
	}
	return results
}

// GetAWSIoTThingPrincipalAttachmentWithName retrieves all AWSIoTThingPrincipalAttachment items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSIoTThingPrincipalAttachmentWithName(name string) (*resources.AWSIoTThingPrincipalAttachment, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSIoTThingPrincipalAttachment:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSIoTThingPrincipalAttachment not found", name)
}

// GetAllAWSIoTTopicRuleResources retrieves all AWSIoTTopicRule items from an AWS CloudFormation template
func (t *Template) GetAllAWSIoTTopicRuleResources() map[string]*resources.AWSIoTTopicRule {
	results := map[string]*resources.AWSIoTTopicRule{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSIoTTopicRule:
			results[name] = resource
		}
	}
	return results
}

// GetAWSIoTTopicRuleWithName retrieves all AWSIoTTopicRule items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSIoTTopicRuleWithName(name string) (*resources.AWSIoTTopicRule, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSIoTTopicRule:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSIoTTopicRule not found", name)
}

// GetAllAWSIoTAnalyticsChannelResources retrieves all AWSIoTAnalyticsChannel items from an AWS CloudFormation template
func (t *Template) GetAllAWSIoTAnalyticsChannelResources() map[string]*resources.AWSIoTAnalyticsChannel {
	results := map[string]*resources.AWSIoTAnalyticsChannel{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSIoTAnalyticsChannel:
			results[name] = resource
		}
	}
	return results
}

// GetAWSIoTAnalyticsChannelWithName retrieves all AWSIoTAnalyticsChannel items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSIoTAnalyticsChannelWithName(name string) (*resources.AWSIoTAnalyticsChannel, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSIoTAnalyticsChannel:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSIoTAnalyticsChannel not found", name)
}

// GetAllAWSIoTAnalyticsDatasetResources retrieves all AWSIoTAnalyticsDataset items from an AWS CloudFormation template
func (t *Template) GetAllAWSIoTAnalyticsDatasetResources() map[string]*resources.AWSIoTAnalyticsDataset {
	results := map[string]*resources.AWSIoTAnalyticsDataset{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSIoTAnalyticsDataset:
			results[name] = resource
		}
	}
	return results
}

// GetAWSIoTAnalyticsDatasetWithName retrieves all AWSIoTAnalyticsDataset items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSIoTAnalyticsDatasetWithName(name string) (*resources.AWSIoTAnalyticsDataset, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSIoTAnalyticsDataset:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSIoTAnalyticsDataset not found", name)
}

// GetAllAWSIoTAnalyticsDatastoreResources retrieves all AWSIoTAnalyticsDatastore items from an AWS CloudFormation template
func (t *Template) GetAllAWSIoTAnalyticsDatastoreResources() map[string]*resources.AWSIoTAnalyticsDatastore {
	results := map[string]*resources.AWSIoTAnalyticsDatastore{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSIoTAnalyticsDatastore:
			results[name] = resource
		}
	}
	return results
}

// GetAWSIoTAnalyticsDatastoreWithName retrieves all AWSIoTAnalyticsDatastore items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSIoTAnalyticsDatastoreWithName(name string) (*resources.AWSIoTAnalyticsDatastore, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSIoTAnalyticsDatastore:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSIoTAnalyticsDatastore not found", name)
}

// GetAllAWSIoTAnalyticsPipelineResources retrieves all AWSIoTAnalyticsPipeline items from an AWS CloudFormation template
func (t *Template) GetAllAWSIoTAnalyticsPipelineResources() map[string]*resources.AWSIoTAnalyticsPipeline {
	results := map[string]*resources.AWSIoTAnalyticsPipeline{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSIoTAnalyticsPipeline:
			results[name] = resource
		}
	}
	return results
}

// GetAWSIoTAnalyticsPipelineWithName retrieves all AWSIoTAnalyticsPipeline items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSIoTAnalyticsPipelineWithName(name string) (*resources.AWSIoTAnalyticsPipeline, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSIoTAnalyticsPipeline:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSIoTAnalyticsPipeline not found", name)
}

// GetAllAWSKMSAliasResources retrieves all AWSKMSAlias items from an AWS CloudFormation template
func (t *Template) GetAllAWSKMSAliasResources() map[string]*resources.AWSKMSAlias {
	results := map[string]*resources.AWSKMSAlias{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSKMSAlias:
			results[name] = resource
		}
	}
	return results
}

// GetAWSKMSAliasWithName retrieves all AWSKMSAlias items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSKMSAliasWithName(name string) (*resources.AWSKMSAlias, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSKMSAlias:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSKMSAlias not found", name)
}

// GetAllAWSKMSKeyResources retrieves all AWSKMSKey items from an AWS CloudFormation template
func (t *Template) GetAllAWSKMSKeyResources() map[string]*resources.AWSKMSKey {
	results := map[string]*resources.AWSKMSKey{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSKMSKey:
			results[name] = resource
		}
	}
	return results
}

// GetAWSKMSKeyWithName retrieves all AWSKMSKey items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSKMSKeyWithName(name string) (*resources.AWSKMSKey, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSKMSKey:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSKMSKey not found", name)
}

// GetAllAWSKinesisStreamResources retrieves all AWSKinesisStream items from an AWS CloudFormation template
func (t *Template) GetAllAWSKinesisStreamResources() map[string]*resources.AWSKinesisStream {
	results := map[string]*resources.AWSKinesisStream{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSKinesisStream:
			results[name] = resource
		}
	}
	return results
}

// GetAWSKinesisStreamWithName retrieves all AWSKinesisStream items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSKinesisStreamWithName(name string) (*resources.AWSKinesisStream, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSKinesisStream:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSKinesisStream not found", name)
}

// GetAllAWSKinesisStreamConsumerResources retrieves all AWSKinesisStreamConsumer items from an AWS CloudFormation template
func (t *Template) GetAllAWSKinesisStreamConsumerResources() map[string]*resources.AWSKinesisStreamConsumer {
	results := map[string]*resources.AWSKinesisStreamConsumer{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSKinesisStreamConsumer:
			results[name] = resource
		}
	}
	return results
}

// GetAWSKinesisStreamConsumerWithName retrieves all AWSKinesisStreamConsumer items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSKinesisStreamConsumerWithName(name string) (*resources.AWSKinesisStreamConsumer, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSKinesisStreamConsumer:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSKinesisStreamConsumer not found", name)
}

// GetAllAWSKinesisAnalyticsApplicationResources retrieves all AWSKinesisAnalyticsApplication items from an AWS CloudFormation template
func (t *Template) GetAllAWSKinesisAnalyticsApplicationResources() map[string]*resources.AWSKinesisAnalyticsApplication {
	results := map[string]*resources.AWSKinesisAnalyticsApplication{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSKinesisAnalyticsApplication:
			results[name] = resource
		}
	}
	return results
}

// GetAWSKinesisAnalyticsApplicationWithName retrieves all AWSKinesisAnalyticsApplication items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSKinesisAnalyticsApplicationWithName(name string) (*resources.AWSKinesisAnalyticsApplication, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSKinesisAnalyticsApplication:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSKinesisAnalyticsApplication not found", name)
}

// GetAllAWSKinesisAnalyticsApplicationOutputResources retrieves all AWSKinesisAnalyticsApplicationOutput items from an AWS CloudFormation template
func (t *Template) GetAllAWSKinesisAnalyticsApplicationOutputResources() map[string]*resources.AWSKinesisAnalyticsApplicationOutput {
	results := map[string]*resources.AWSKinesisAnalyticsApplicationOutput{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSKinesisAnalyticsApplicationOutput:
			results[name] = resource
		}
	}
	return results
}

// GetAWSKinesisAnalyticsApplicationOutputWithName retrieves all AWSKinesisAnalyticsApplicationOutput items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSKinesisAnalyticsApplicationOutputWithName(name string) (*resources.AWSKinesisAnalyticsApplicationOutput, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSKinesisAnalyticsApplicationOutput:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSKinesisAnalyticsApplicationOutput not found", name)
}

// GetAllAWSKinesisAnalyticsApplicationReferenceDataSourceResources retrieves all AWSKinesisAnalyticsApplicationReferenceDataSource items from an AWS CloudFormation template
func (t *Template) GetAllAWSKinesisAnalyticsApplicationReferenceDataSourceResources() map[string]*resources.AWSKinesisAnalyticsApplicationReferenceDataSource {
	results := map[string]*resources.AWSKinesisAnalyticsApplicationReferenceDataSource{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSKinesisAnalyticsApplicationReferenceDataSource:
			results[name] = resource
		}
	}
	return results
}

// GetAWSKinesisAnalyticsApplicationReferenceDataSourceWithName retrieves all AWSKinesisAnalyticsApplicationReferenceDataSource items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSKinesisAnalyticsApplicationReferenceDataSourceWithName(name string) (*resources.AWSKinesisAnalyticsApplicationReferenceDataSource, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSKinesisAnalyticsApplicationReferenceDataSource:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSKinesisAnalyticsApplicationReferenceDataSource not found", name)
}

// GetAllAWSKinesisAnalyticsV2ApplicationResources retrieves all AWSKinesisAnalyticsV2Application items from an AWS CloudFormation template
func (t *Template) GetAllAWSKinesisAnalyticsV2ApplicationResources() map[string]*resources.AWSKinesisAnalyticsV2Application {
	results := map[string]*resources.AWSKinesisAnalyticsV2Application{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSKinesisAnalyticsV2Application:
			results[name] = resource
		}
	}
	return results
}

// GetAWSKinesisAnalyticsV2ApplicationWithName retrieves all AWSKinesisAnalyticsV2Application items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSKinesisAnalyticsV2ApplicationWithName(name string) (*resources.AWSKinesisAnalyticsV2Application, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSKinesisAnalyticsV2Application:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSKinesisAnalyticsV2Application not found", name)
}

// GetAllAWSKinesisAnalyticsV2ApplicationCloudWatchLoggingOptionResources retrieves all AWSKinesisAnalyticsV2ApplicationCloudWatchLoggingOption items from an AWS CloudFormation template
func (t *Template) GetAllAWSKinesisAnalyticsV2ApplicationCloudWatchLoggingOptionResources() map[string]*resources.AWSKinesisAnalyticsV2ApplicationCloudWatchLoggingOption {
	results := map[string]*resources.AWSKinesisAnalyticsV2ApplicationCloudWatchLoggingOption{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSKinesisAnalyticsV2ApplicationCloudWatchLoggingOption:
			results[name] = resource
		}
	}
	return results
}

// GetAWSKinesisAnalyticsV2ApplicationCloudWatchLoggingOptionWithName retrieves all AWSKinesisAnalyticsV2ApplicationCloudWatchLoggingOption items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSKinesisAnalyticsV2ApplicationCloudWatchLoggingOptionWithName(name string) (*resources.AWSKinesisAnalyticsV2ApplicationCloudWatchLoggingOption, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSKinesisAnalyticsV2ApplicationCloudWatchLoggingOption:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSKinesisAnalyticsV2ApplicationCloudWatchLoggingOption not found", name)
}

// GetAllAWSKinesisAnalyticsV2ApplicationOutputResources retrieves all AWSKinesisAnalyticsV2ApplicationOutput items from an AWS CloudFormation template
func (t *Template) GetAllAWSKinesisAnalyticsV2ApplicationOutputResources() map[string]*resources.AWSKinesisAnalyticsV2ApplicationOutput {
	results := map[string]*resources.AWSKinesisAnalyticsV2ApplicationOutput{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSKinesisAnalyticsV2ApplicationOutput:
			results[name] = resource
		}
	}
	return results
}

// GetAWSKinesisAnalyticsV2ApplicationOutputWithName retrieves all AWSKinesisAnalyticsV2ApplicationOutput items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSKinesisAnalyticsV2ApplicationOutputWithName(name string) (*resources.AWSKinesisAnalyticsV2ApplicationOutput, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSKinesisAnalyticsV2ApplicationOutput:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSKinesisAnalyticsV2ApplicationOutput not found", name)
}

// GetAllAWSKinesisAnalyticsV2ApplicationReferenceDataSourceResources retrieves all AWSKinesisAnalyticsV2ApplicationReferenceDataSource items from an AWS CloudFormation template
func (t *Template) GetAllAWSKinesisAnalyticsV2ApplicationReferenceDataSourceResources() map[string]*resources.AWSKinesisAnalyticsV2ApplicationReferenceDataSource {
	results := map[string]*resources.AWSKinesisAnalyticsV2ApplicationReferenceDataSource{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSKinesisAnalyticsV2ApplicationReferenceDataSource:
			results[name] = resource
		}
	}
	return results
}

// GetAWSKinesisAnalyticsV2ApplicationReferenceDataSourceWithName retrieves all AWSKinesisAnalyticsV2ApplicationReferenceDataSource items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSKinesisAnalyticsV2ApplicationReferenceDataSourceWithName(name string) (*resources.AWSKinesisAnalyticsV2ApplicationReferenceDataSource, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSKinesisAnalyticsV2ApplicationReferenceDataSource:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSKinesisAnalyticsV2ApplicationReferenceDataSource not found", name)
}

// GetAllAWSKinesisFirehoseDeliveryStreamResources retrieves all AWSKinesisFirehoseDeliveryStream items from an AWS CloudFormation template
func (t *Template) GetAllAWSKinesisFirehoseDeliveryStreamResources() map[string]*resources.AWSKinesisFirehoseDeliveryStream {
	results := map[string]*resources.AWSKinesisFirehoseDeliveryStream{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSKinesisFirehoseDeliveryStream:
			results[name] = resource
		}
	}
	return results
}

// GetAWSKinesisFirehoseDeliveryStreamWithName retrieves all AWSKinesisFirehoseDeliveryStream items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSKinesisFirehoseDeliveryStreamWithName(name string) (*resources.AWSKinesisFirehoseDeliveryStream, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSKinesisFirehoseDeliveryStream:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSKinesisFirehoseDeliveryStream not found", name)
}

// GetAllAWSLambdaAliasResources retrieves all AWSLambdaAlias items from an AWS CloudFormation template
func (t *Template) GetAllAWSLambdaAliasResources() map[string]*resources.AWSLambdaAlias {
	results := map[string]*resources.AWSLambdaAlias{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSLambdaAlias:
			results[name] = resource
		}
	}
	return results
}

// GetAWSLambdaAliasWithName retrieves all AWSLambdaAlias items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSLambdaAliasWithName(name string) (*resources.AWSLambdaAlias, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSLambdaAlias:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSLambdaAlias not found", name)
}

// GetAllAWSLambdaEventSourceMappingResources retrieves all AWSLambdaEventSourceMapping items from an AWS CloudFormation template
func (t *Template) GetAllAWSLambdaEventSourceMappingResources() map[string]*resources.AWSLambdaEventSourceMapping {
	results := map[string]*resources.AWSLambdaEventSourceMapping{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSLambdaEventSourceMapping:
			results[name] = resource
		}
	}
	return results
}

// GetAWSLambdaEventSourceMappingWithName retrieves all AWSLambdaEventSourceMapping items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSLambdaEventSourceMappingWithName(name string) (*resources.AWSLambdaEventSourceMapping, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSLambdaEventSourceMapping:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSLambdaEventSourceMapping not found", name)
}

// GetAllAWSLambdaFunctionResources retrieves all AWSLambdaFunction items from an AWS CloudFormation template
func (t *Template) GetAllAWSLambdaFunctionResources() map[string]*resources.AWSLambdaFunction {
	results := map[string]*resources.AWSLambdaFunction{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSLambdaFunction:
			results[name] = resource
		}
	}
	return results
}

// GetAWSLambdaFunctionWithName retrieves all AWSLambdaFunction items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSLambdaFunctionWithName(name string) (*resources.AWSLambdaFunction, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSLambdaFunction:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSLambdaFunction not found", name)
}

// GetAllAWSLambdaLayerVersionResources retrieves all AWSLambdaLayerVersion items from an AWS CloudFormation template
func (t *Template) GetAllAWSLambdaLayerVersionResources() map[string]*resources.AWSLambdaLayerVersion {
	results := map[string]*resources.AWSLambdaLayerVersion{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSLambdaLayerVersion:
			results[name] = resource
		}
	}
	return results
}

// GetAWSLambdaLayerVersionWithName retrieves all AWSLambdaLayerVersion items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSLambdaLayerVersionWithName(name string) (*resources.AWSLambdaLayerVersion, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSLambdaLayerVersion:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSLambdaLayerVersion not found", name)
}

// GetAllAWSLambdaLayerVersionPermissionResources retrieves all AWSLambdaLayerVersionPermission items from an AWS CloudFormation template
func (t *Template) GetAllAWSLambdaLayerVersionPermissionResources() map[string]*resources.AWSLambdaLayerVersionPermission {
	results := map[string]*resources.AWSLambdaLayerVersionPermission{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSLambdaLayerVersionPermission:
			results[name] = resource
		}
	}
	return results
}

// GetAWSLambdaLayerVersionPermissionWithName retrieves all AWSLambdaLayerVersionPermission items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSLambdaLayerVersionPermissionWithName(name string) (*resources.AWSLambdaLayerVersionPermission, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSLambdaLayerVersionPermission:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSLambdaLayerVersionPermission not found", name)
}

// GetAllAWSLambdaPermissionResources retrieves all AWSLambdaPermission items from an AWS CloudFormation template
func (t *Template) GetAllAWSLambdaPermissionResources() map[string]*resources.AWSLambdaPermission {
	results := map[string]*resources.AWSLambdaPermission{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSLambdaPermission:
			results[name] = resource
		}
	}
	return results
}

// GetAWSLambdaPermissionWithName retrieves all AWSLambdaPermission items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSLambdaPermissionWithName(name string) (*resources.AWSLambdaPermission, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSLambdaPermission:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSLambdaPermission not found", name)
}

// GetAllAWSLambdaVersionResources retrieves all AWSLambdaVersion items from an AWS CloudFormation template
func (t *Template) GetAllAWSLambdaVersionResources() map[string]*resources.AWSLambdaVersion {
	results := map[string]*resources.AWSLambdaVersion{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSLambdaVersion:
			results[name] = resource
		}
	}
	return results
}

// GetAWSLambdaVersionWithName retrieves all AWSLambdaVersion items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSLambdaVersionWithName(name string) (*resources.AWSLambdaVersion, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSLambdaVersion:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSLambdaVersion not found", name)
}

// GetAllAWSLogsDestinationResources retrieves all AWSLogsDestination items from an AWS CloudFormation template
func (t *Template) GetAllAWSLogsDestinationResources() map[string]*resources.AWSLogsDestination {
	results := map[string]*resources.AWSLogsDestination{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSLogsDestination:
			results[name] = resource
		}
	}
	return results
}

// GetAWSLogsDestinationWithName retrieves all AWSLogsDestination items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSLogsDestinationWithName(name string) (*resources.AWSLogsDestination, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSLogsDestination:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSLogsDestination not found", name)
}

// GetAllAWSLogsLogGroupResources retrieves all AWSLogsLogGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSLogsLogGroupResources() map[string]*resources.AWSLogsLogGroup {
	results := map[string]*resources.AWSLogsLogGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSLogsLogGroup:
			results[name] = resource
		}
	}
	return results
}

// GetAWSLogsLogGroupWithName retrieves all AWSLogsLogGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSLogsLogGroupWithName(name string) (*resources.AWSLogsLogGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSLogsLogGroup:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSLogsLogGroup not found", name)
}

// GetAllAWSLogsLogStreamResources retrieves all AWSLogsLogStream items from an AWS CloudFormation template
func (t *Template) GetAllAWSLogsLogStreamResources() map[string]*resources.AWSLogsLogStream {
	results := map[string]*resources.AWSLogsLogStream{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSLogsLogStream:
			results[name] = resource
		}
	}
	return results
}

// GetAWSLogsLogStreamWithName retrieves all AWSLogsLogStream items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSLogsLogStreamWithName(name string) (*resources.AWSLogsLogStream, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSLogsLogStream:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSLogsLogStream not found", name)
}

// GetAllAWSLogsMetricFilterResources retrieves all AWSLogsMetricFilter items from an AWS CloudFormation template
func (t *Template) GetAllAWSLogsMetricFilterResources() map[string]*resources.AWSLogsMetricFilter {
	results := map[string]*resources.AWSLogsMetricFilter{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSLogsMetricFilter:
			results[name] = resource
		}
	}
	return results
}

// GetAWSLogsMetricFilterWithName retrieves all AWSLogsMetricFilter items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSLogsMetricFilterWithName(name string) (*resources.AWSLogsMetricFilter, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSLogsMetricFilter:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSLogsMetricFilter not found", name)
}

// GetAllAWSLogsSubscriptionFilterResources retrieves all AWSLogsSubscriptionFilter items from an AWS CloudFormation template
func (t *Template) GetAllAWSLogsSubscriptionFilterResources() map[string]*resources.AWSLogsSubscriptionFilter {
	results := map[string]*resources.AWSLogsSubscriptionFilter{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSLogsSubscriptionFilter:
			results[name] = resource
		}
	}
	return results
}

// GetAWSLogsSubscriptionFilterWithName retrieves all AWSLogsSubscriptionFilter items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSLogsSubscriptionFilterWithName(name string) (*resources.AWSLogsSubscriptionFilter, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSLogsSubscriptionFilter:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSLogsSubscriptionFilter not found", name)
}

// GetAllAWSNeptuneDBClusterResources retrieves all AWSNeptuneDBCluster items from an AWS CloudFormation template
func (t *Template) GetAllAWSNeptuneDBClusterResources() map[string]*resources.AWSNeptuneDBCluster {
	results := map[string]*resources.AWSNeptuneDBCluster{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSNeptuneDBCluster:
			results[name] = resource
		}
	}
	return results
}

// GetAWSNeptuneDBClusterWithName retrieves all AWSNeptuneDBCluster items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSNeptuneDBClusterWithName(name string) (*resources.AWSNeptuneDBCluster, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSNeptuneDBCluster:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSNeptuneDBCluster not found", name)
}

// GetAllAWSNeptuneDBClusterParameterGroupResources retrieves all AWSNeptuneDBClusterParameterGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSNeptuneDBClusterParameterGroupResources() map[string]*resources.AWSNeptuneDBClusterParameterGroup {
	results := map[string]*resources.AWSNeptuneDBClusterParameterGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSNeptuneDBClusterParameterGroup:
			results[name] = resource
		}
	}
	return results
}

// GetAWSNeptuneDBClusterParameterGroupWithName retrieves all AWSNeptuneDBClusterParameterGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSNeptuneDBClusterParameterGroupWithName(name string) (*resources.AWSNeptuneDBClusterParameterGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSNeptuneDBClusterParameterGroup:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSNeptuneDBClusterParameterGroup not found", name)
}

// GetAllAWSNeptuneDBInstanceResources retrieves all AWSNeptuneDBInstance items from an AWS CloudFormation template
func (t *Template) GetAllAWSNeptuneDBInstanceResources() map[string]*resources.AWSNeptuneDBInstance {
	results := map[string]*resources.AWSNeptuneDBInstance{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSNeptuneDBInstance:
			results[name] = resource
		}
	}
	return results
}

// GetAWSNeptuneDBInstanceWithName retrieves all AWSNeptuneDBInstance items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSNeptuneDBInstanceWithName(name string) (*resources.AWSNeptuneDBInstance, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSNeptuneDBInstance:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSNeptuneDBInstance not found", name)
}

// GetAllAWSNeptuneDBParameterGroupResources retrieves all AWSNeptuneDBParameterGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSNeptuneDBParameterGroupResources() map[string]*resources.AWSNeptuneDBParameterGroup {
	results := map[string]*resources.AWSNeptuneDBParameterGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSNeptuneDBParameterGroup:
			results[name] = resource
		}
	}
	return results
}

// GetAWSNeptuneDBParameterGroupWithName retrieves all AWSNeptuneDBParameterGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSNeptuneDBParameterGroupWithName(name string) (*resources.AWSNeptuneDBParameterGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSNeptuneDBParameterGroup:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSNeptuneDBParameterGroup not found", name)
}

// GetAllAWSNeptuneDBSubnetGroupResources retrieves all AWSNeptuneDBSubnetGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSNeptuneDBSubnetGroupResources() map[string]*resources.AWSNeptuneDBSubnetGroup {
	results := map[string]*resources.AWSNeptuneDBSubnetGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSNeptuneDBSubnetGroup:
			results[name] = resource
		}
	}
	return results
}

// GetAWSNeptuneDBSubnetGroupWithName retrieves all AWSNeptuneDBSubnetGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSNeptuneDBSubnetGroupWithName(name string) (*resources.AWSNeptuneDBSubnetGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSNeptuneDBSubnetGroup:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSNeptuneDBSubnetGroup not found", name)
}

// GetAllAWSOpsWorksAppResources retrieves all AWSOpsWorksApp items from an AWS CloudFormation template
func (t *Template) GetAllAWSOpsWorksAppResources() map[string]*resources.AWSOpsWorksApp {
	results := map[string]*resources.AWSOpsWorksApp{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSOpsWorksApp:
			results[name] = resource
		}
	}
	return results
}

// GetAWSOpsWorksAppWithName retrieves all AWSOpsWorksApp items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSOpsWorksAppWithName(name string) (*resources.AWSOpsWorksApp, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSOpsWorksApp:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSOpsWorksApp not found", name)
}

// GetAllAWSOpsWorksElasticLoadBalancerAttachmentResources retrieves all AWSOpsWorksElasticLoadBalancerAttachment items from an AWS CloudFormation template
func (t *Template) GetAllAWSOpsWorksElasticLoadBalancerAttachmentResources() map[string]*resources.AWSOpsWorksElasticLoadBalancerAttachment {
	results := map[string]*resources.AWSOpsWorksElasticLoadBalancerAttachment{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSOpsWorksElasticLoadBalancerAttachment:
			results[name] = resource
		}
	}
	return results
}

// GetAWSOpsWorksElasticLoadBalancerAttachmentWithName retrieves all AWSOpsWorksElasticLoadBalancerAttachment items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSOpsWorksElasticLoadBalancerAttachmentWithName(name string) (*resources.AWSOpsWorksElasticLoadBalancerAttachment, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSOpsWorksElasticLoadBalancerAttachment:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSOpsWorksElasticLoadBalancerAttachment not found", name)
}

// GetAllAWSOpsWorksInstanceResources retrieves all AWSOpsWorksInstance items from an AWS CloudFormation template
func (t *Template) GetAllAWSOpsWorksInstanceResources() map[string]*resources.AWSOpsWorksInstance {
	results := map[string]*resources.AWSOpsWorksInstance{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSOpsWorksInstance:
			results[name] = resource
		}
	}
	return results
}

// GetAWSOpsWorksInstanceWithName retrieves all AWSOpsWorksInstance items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSOpsWorksInstanceWithName(name string) (*resources.AWSOpsWorksInstance, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSOpsWorksInstance:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSOpsWorksInstance not found", name)
}

// GetAllAWSOpsWorksLayerResources retrieves all AWSOpsWorksLayer items from an AWS CloudFormation template
func (t *Template) GetAllAWSOpsWorksLayerResources() map[string]*resources.AWSOpsWorksLayer {
	results := map[string]*resources.AWSOpsWorksLayer{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSOpsWorksLayer:
			results[name] = resource
		}
	}
	return results
}

// GetAWSOpsWorksLayerWithName retrieves all AWSOpsWorksLayer items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSOpsWorksLayerWithName(name string) (*resources.AWSOpsWorksLayer, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSOpsWorksLayer:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSOpsWorksLayer not found", name)
}

// GetAllAWSOpsWorksStackResources retrieves all AWSOpsWorksStack items from an AWS CloudFormation template
func (t *Template) GetAllAWSOpsWorksStackResources() map[string]*resources.AWSOpsWorksStack {
	results := map[string]*resources.AWSOpsWorksStack{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSOpsWorksStack:
			results[name] = resource
		}
	}
	return results
}

// GetAWSOpsWorksStackWithName retrieves all AWSOpsWorksStack items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSOpsWorksStackWithName(name string) (*resources.AWSOpsWorksStack, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSOpsWorksStack:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSOpsWorksStack not found", name)
}

// GetAllAWSOpsWorksUserProfileResources retrieves all AWSOpsWorksUserProfile items from an AWS CloudFormation template
func (t *Template) GetAllAWSOpsWorksUserProfileResources() map[string]*resources.AWSOpsWorksUserProfile {
	results := map[string]*resources.AWSOpsWorksUserProfile{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSOpsWorksUserProfile:
			results[name] = resource
		}
	}
	return results
}

// GetAWSOpsWorksUserProfileWithName retrieves all AWSOpsWorksUserProfile items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSOpsWorksUserProfileWithName(name string) (*resources.AWSOpsWorksUserProfile, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSOpsWorksUserProfile:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSOpsWorksUserProfile not found", name)
}

// GetAllAWSOpsWorksVolumeResources retrieves all AWSOpsWorksVolume items from an AWS CloudFormation template
func (t *Template) GetAllAWSOpsWorksVolumeResources() map[string]*resources.AWSOpsWorksVolume {
	results := map[string]*resources.AWSOpsWorksVolume{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSOpsWorksVolume:
			results[name] = resource
		}
	}
	return results
}

// GetAWSOpsWorksVolumeWithName retrieves all AWSOpsWorksVolume items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSOpsWorksVolumeWithName(name string) (*resources.AWSOpsWorksVolume, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSOpsWorksVolume:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSOpsWorksVolume not found", name)
}

// GetAllAWSOpsWorksCMServerResources retrieves all AWSOpsWorksCMServer items from an AWS CloudFormation template
func (t *Template) GetAllAWSOpsWorksCMServerResources() map[string]*resources.AWSOpsWorksCMServer {
	results := map[string]*resources.AWSOpsWorksCMServer{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSOpsWorksCMServer:
			results[name] = resource
		}
	}
	return results
}

// GetAWSOpsWorksCMServerWithName retrieves all AWSOpsWorksCMServer items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSOpsWorksCMServerWithName(name string) (*resources.AWSOpsWorksCMServer, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSOpsWorksCMServer:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSOpsWorksCMServer not found", name)
}

// GetAllAWSRAMResourceShareResources retrieves all AWSRAMResourceShare items from an AWS CloudFormation template
func (t *Template) GetAllAWSRAMResourceShareResources() map[string]*resources.AWSRAMResourceShare {
	results := map[string]*resources.AWSRAMResourceShare{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSRAMResourceShare:
			results[name] = resource
		}
	}
	return results
}

// GetAWSRAMResourceShareWithName retrieves all AWSRAMResourceShare items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRAMResourceShareWithName(name string) (*resources.AWSRAMResourceShare, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSRAMResourceShare:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSRAMResourceShare not found", name)
}

// GetAllAWSRDSDBClusterResources retrieves all AWSRDSDBCluster items from an AWS CloudFormation template
func (t *Template) GetAllAWSRDSDBClusterResources() map[string]*resources.AWSRDSDBCluster {
	results := map[string]*resources.AWSRDSDBCluster{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSRDSDBCluster:
			results[name] = resource
		}
	}
	return results
}

// GetAWSRDSDBClusterWithName retrieves all AWSRDSDBCluster items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRDSDBClusterWithName(name string) (*resources.AWSRDSDBCluster, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSRDSDBCluster:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSRDSDBCluster not found", name)
}

// GetAllAWSRDSDBClusterParameterGroupResources retrieves all AWSRDSDBClusterParameterGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSRDSDBClusterParameterGroupResources() map[string]*resources.AWSRDSDBClusterParameterGroup {
	results := map[string]*resources.AWSRDSDBClusterParameterGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSRDSDBClusterParameterGroup:
			results[name] = resource
		}
	}
	return results
}

// GetAWSRDSDBClusterParameterGroupWithName retrieves all AWSRDSDBClusterParameterGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRDSDBClusterParameterGroupWithName(name string) (*resources.AWSRDSDBClusterParameterGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSRDSDBClusterParameterGroup:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSRDSDBClusterParameterGroup not found", name)
}

// GetAllAWSRDSDBInstanceResources retrieves all AWSRDSDBInstance items from an AWS CloudFormation template
func (t *Template) GetAllAWSRDSDBInstanceResources() map[string]*resources.AWSRDSDBInstance {
	results := map[string]*resources.AWSRDSDBInstance{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSRDSDBInstance:
			results[name] = resource
		}
	}
	return results
}

// GetAWSRDSDBInstanceWithName retrieves all AWSRDSDBInstance items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRDSDBInstanceWithName(name string) (*resources.AWSRDSDBInstance, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSRDSDBInstance:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSRDSDBInstance not found", name)
}

// GetAllAWSRDSDBParameterGroupResources retrieves all AWSRDSDBParameterGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSRDSDBParameterGroupResources() map[string]*resources.AWSRDSDBParameterGroup {
	results := map[string]*resources.AWSRDSDBParameterGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSRDSDBParameterGroup:
			results[name] = resource
		}
	}
	return results
}

// GetAWSRDSDBParameterGroupWithName retrieves all AWSRDSDBParameterGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRDSDBParameterGroupWithName(name string) (*resources.AWSRDSDBParameterGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSRDSDBParameterGroup:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSRDSDBParameterGroup not found", name)
}

// GetAllAWSRDSDBSecurityGroupResources retrieves all AWSRDSDBSecurityGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSRDSDBSecurityGroupResources() map[string]*resources.AWSRDSDBSecurityGroup {
	results := map[string]*resources.AWSRDSDBSecurityGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSRDSDBSecurityGroup:
			results[name] = resource
		}
	}
	return results
}

// GetAWSRDSDBSecurityGroupWithName retrieves all AWSRDSDBSecurityGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRDSDBSecurityGroupWithName(name string) (*resources.AWSRDSDBSecurityGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSRDSDBSecurityGroup:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSRDSDBSecurityGroup not found", name)
}

// GetAllAWSRDSDBSecurityGroupIngressResources retrieves all AWSRDSDBSecurityGroupIngress items from an AWS CloudFormation template
func (t *Template) GetAllAWSRDSDBSecurityGroupIngressResources() map[string]*resources.AWSRDSDBSecurityGroupIngress {
	results := map[string]*resources.AWSRDSDBSecurityGroupIngress{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSRDSDBSecurityGroupIngress:
			results[name] = resource
		}
	}
	return results
}

// GetAWSRDSDBSecurityGroupIngressWithName retrieves all AWSRDSDBSecurityGroupIngress items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRDSDBSecurityGroupIngressWithName(name string) (*resources.AWSRDSDBSecurityGroupIngress, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSRDSDBSecurityGroupIngress:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSRDSDBSecurityGroupIngress not found", name)
}

// GetAllAWSRDSDBSubnetGroupResources retrieves all AWSRDSDBSubnetGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSRDSDBSubnetGroupResources() map[string]*resources.AWSRDSDBSubnetGroup {
	results := map[string]*resources.AWSRDSDBSubnetGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSRDSDBSubnetGroup:
			results[name] = resource
		}
	}
	return results
}

// GetAWSRDSDBSubnetGroupWithName retrieves all AWSRDSDBSubnetGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRDSDBSubnetGroupWithName(name string) (*resources.AWSRDSDBSubnetGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSRDSDBSubnetGroup:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSRDSDBSubnetGroup not found", name)
}

// GetAllAWSRDSEventSubscriptionResources retrieves all AWSRDSEventSubscription items from an AWS CloudFormation template
func (t *Template) GetAllAWSRDSEventSubscriptionResources() map[string]*resources.AWSRDSEventSubscription {
	results := map[string]*resources.AWSRDSEventSubscription{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSRDSEventSubscription:
			results[name] = resource
		}
	}
	return results
}

// GetAWSRDSEventSubscriptionWithName retrieves all AWSRDSEventSubscription items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRDSEventSubscriptionWithName(name string) (*resources.AWSRDSEventSubscription, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSRDSEventSubscription:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSRDSEventSubscription not found", name)
}

// GetAllAWSRDSOptionGroupResources retrieves all AWSRDSOptionGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSRDSOptionGroupResources() map[string]*resources.AWSRDSOptionGroup {
	results := map[string]*resources.AWSRDSOptionGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSRDSOptionGroup:
			results[name] = resource
		}
	}
	return results
}

// GetAWSRDSOptionGroupWithName retrieves all AWSRDSOptionGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRDSOptionGroupWithName(name string) (*resources.AWSRDSOptionGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSRDSOptionGroup:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSRDSOptionGroup not found", name)
}

// GetAllAWSRedshiftClusterResources retrieves all AWSRedshiftCluster items from an AWS CloudFormation template
func (t *Template) GetAllAWSRedshiftClusterResources() map[string]*resources.AWSRedshiftCluster {
	results := map[string]*resources.AWSRedshiftCluster{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSRedshiftCluster:
			results[name] = resource
		}
	}
	return results
}

// GetAWSRedshiftClusterWithName retrieves all AWSRedshiftCluster items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRedshiftClusterWithName(name string) (*resources.AWSRedshiftCluster, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSRedshiftCluster:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSRedshiftCluster not found", name)
}

// GetAllAWSRedshiftClusterParameterGroupResources retrieves all AWSRedshiftClusterParameterGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSRedshiftClusterParameterGroupResources() map[string]*resources.AWSRedshiftClusterParameterGroup {
	results := map[string]*resources.AWSRedshiftClusterParameterGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSRedshiftClusterParameterGroup:
			results[name] = resource
		}
	}
	return results
}

// GetAWSRedshiftClusterParameterGroupWithName retrieves all AWSRedshiftClusterParameterGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRedshiftClusterParameterGroupWithName(name string) (*resources.AWSRedshiftClusterParameterGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSRedshiftClusterParameterGroup:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSRedshiftClusterParameterGroup not found", name)
}

// GetAllAWSRedshiftClusterSecurityGroupResources retrieves all AWSRedshiftClusterSecurityGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSRedshiftClusterSecurityGroupResources() map[string]*resources.AWSRedshiftClusterSecurityGroup {
	results := map[string]*resources.AWSRedshiftClusterSecurityGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSRedshiftClusterSecurityGroup:
			results[name] = resource
		}
	}
	return results
}

// GetAWSRedshiftClusterSecurityGroupWithName retrieves all AWSRedshiftClusterSecurityGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRedshiftClusterSecurityGroupWithName(name string) (*resources.AWSRedshiftClusterSecurityGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSRedshiftClusterSecurityGroup:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSRedshiftClusterSecurityGroup not found", name)
}

// GetAllAWSRedshiftClusterSecurityGroupIngressResources retrieves all AWSRedshiftClusterSecurityGroupIngress items from an AWS CloudFormation template
func (t *Template) GetAllAWSRedshiftClusterSecurityGroupIngressResources() map[string]*resources.AWSRedshiftClusterSecurityGroupIngress {
	results := map[string]*resources.AWSRedshiftClusterSecurityGroupIngress{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSRedshiftClusterSecurityGroupIngress:
			results[name] = resource
		}
	}
	return results
}

// GetAWSRedshiftClusterSecurityGroupIngressWithName retrieves all AWSRedshiftClusterSecurityGroupIngress items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRedshiftClusterSecurityGroupIngressWithName(name string) (*resources.AWSRedshiftClusterSecurityGroupIngress, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSRedshiftClusterSecurityGroupIngress:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSRedshiftClusterSecurityGroupIngress not found", name)
}

// GetAllAWSRedshiftClusterSubnetGroupResources retrieves all AWSRedshiftClusterSubnetGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSRedshiftClusterSubnetGroupResources() map[string]*resources.AWSRedshiftClusterSubnetGroup {
	results := map[string]*resources.AWSRedshiftClusterSubnetGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSRedshiftClusterSubnetGroup:
			results[name] = resource
		}
	}
	return results
}

// GetAWSRedshiftClusterSubnetGroupWithName retrieves all AWSRedshiftClusterSubnetGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRedshiftClusterSubnetGroupWithName(name string) (*resources.AWSRedshiftClusterSubnetGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSRedshiftClusterSubnetGroup:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSRedshiftClusterSubnetGroup not found", name)
}

// GetAllAWSRoboMakerFleetResources retrieves all AWSRoboMakerFleet items from an AWS CloudFormation template
func (t *Template) GetAllAWSRoboMakerFleetResources() map[string]*resources.AWSRoboMakerFleet {
	results := map[string]*resources.AWSRoboMakerFleet{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSRoboMakerFleet:
			results[name] = resource
		}
	}
	return results
}

// GetAWSRoboMakerFleetWithName retrieves all AWSRoboMakerFleet items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRoboMakerFleetWithName(name string) (*resources.AWSRoboMakerFleet, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSRoboMakerFleet:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSRoboMakerFleet not found", name)
}

// GetAllAWSRoboMakerRobotResources retrieves all AWSRoboMakerRobot items from an AWS CloudFormation template
func (t *Template) GetAllAWSRoboMakerRobotResources() map[string]*resources.AWSRoboMakerRobot {
	results := map[string]*resources.AWSRoboMakerRobot{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSRoboMakerRobot:
			results[name] = resource
		}
	}
	return results
}

// GetAWSRoboMakerRobotWithName retrieves all AWSRoboMakerRobot items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRoboMakerRobotWithName(name string) (*resources.AWSRoboMakerRobot, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSRoboMakerRobot:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSRoboMakerRobot not found", name)
}

// GetAllAWSRoboMakerRobotApplicationResources retrieves all AWSRoboMakerRobotApplication items from an AWS CloudFormation template
func (t *Template) GetAllAWSRoboMakerRobotApplicationResources() map[string]*resources.AWSRoboMakerRobotApplication {
	results := map[string]*resources.AWSRoboMakerRobotApplication{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSRoboMakerRobotApplication:
			results[name] = resource
		}
	}
	return results
}

// GetAWSRoboMakerRobotApplicationWithName retrieves all AWSRoboMakerRobotApplication items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRoboMakerRobotApplicationWithName(name string) (*resources.AWSRoboMakerRobotApplication, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSRoboMakerRobotApplication:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSRoboMakerRobotApplication not found", name)
}

// GetAllAWSRoboMakerRobotApplicationVersionResources retrieves all AWSRoboMakerRobotApplicationVersion items from an AWS CloudFormation template
func (t *Template) GetAllAWSRoboMakerRobotApplicationVersionResources() map[string]*resources.AWSRoboMakerRobotApplicationVersion {
	results := map[string]*resources.AWSRoboMakerRobotApplicationVersion{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSRoboMakerRobotApplicationVersion:
			results[name] = resource
		}
	}
	return results
}

// GetAWSRoboMakerRobotApplicationVersionWithName retrieves all AWSRoboMakerRobotApplicationVersion items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRoboMakerRobotApplicationVersionWithName(name string) (*resources.AWSRoboMakerRobotApplicationVersion, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSRoboMakerRobotApplicationVersion:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSRoboMakerRobotApplicationVersion not found", name)
}

// GetAllAWSRoboMakerSimulationApplicationResources retrieves all AWSRoboMakerSimulationApplication items from an AWS CloudFormation template
func (t *Template) GetAllAWSRoboMakerSimulationApplicationResources() map[string]*resources.AWSRoboMakerSimulationApplication {
	results := map[string]*resources.AWSRoboMakerSimulationApplication{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSRoboMakerSimulationApplication:
			results[name] = resource
		}
	}
	return results
}

// GetAWSRoboMakerSimulationApplicationWithName retrieves all AWSRoboMakerSimulationApplication items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRoboMakerSimulationApplicationWithName(name string) (*resources.AWSRoboMakerSimulationApplication, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSRoboMakerSimulationApplication:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSRoboMakerSimulationApplication not found", name)
}

// GetAllAWSRoboMakerSimulationApplicationVersionResources retrieves all AWSRoboMakerSimulationApplicationVersion items from an AWS CloudFormation template
func (t *Template) GetAllAWSRoboMakerSimulationApplicationVersionResources() map[string]*resources.AWSRoboMakerSimulationApplicationVersion {
	results := map[string]*resources.AWSRoboMakerSimulationApplicationVersion{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSRoboMakerSimulationApplicationVersion:
			results[name] = resource
		}
	}
	return results
}

// GetAWSRoboMakerSimulationApplicationVersionWithName retrieves all AWSRoboMakerSimulationApplicationVersion items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRoboMakerSimulationApplicationVersionWithName(name string) (*resources.AWSRoboMakerSimulationApplicationVersion, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSRoboMakerSimulationApplicationVersion:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSRoboMakerSimulationApplicationVersion not found", name)
}

// GetAllAWSRoute53HealthCheckResources retrieves all AWSRoute53HealthCheck items from an AWS CloudFormation template
func (t *Template) GetAllAWSRoute53HealthCheckResources() map[string]*resources.AWSRoute53HealthCheck {
	results := map[string]*resources.AWSRoute53HealthCheck{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSRoute53HealthCheck:
			results[name] = resource
		}
	}
	return results
}

// GetAWSRoute53HealthCheckWithName retrieves all AWSRoute53HealthCheck items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRoute53HealthCheckWithName(name string) (*resources.AWSRoute53HealthCheck, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSRoute53HealthCheck:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSRoute53HealthCheck not found", name)
}

// GetAllAWSRoute53HostedZoneResources retrieves all AWSRoute53HostedZone items from an AWS CloudFormation template
func (t *Template) GetAllAWSRoute53HostedZoneResources() map[string]*resources.AWSRoute53HostedZone {
	results := map[string]*resources.AWSRoute53HostedZone{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSRoute53HostedZone:
			results[name] = resource
		}
	}
	return results
}

// GetAWSRoute53HostedZoneWithName retrieves all AWSRoute53HostedZone items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRoute53HostedZoneWithName(name string) (*resources.AWSRoute53HostedZone, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSRoute53HostedZone:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSRoute53HostedZone not found", name)
}

// GetAllAWSRoute53RecordSetResources retrieves all AWSRoute53RecordSet items from an AWS CloudFormation template
func (t *Template) GetAllAWSRoute53RecordSetResources() map[string]*resources.AWSRoute53RecordSet {
	results := map[string]*resources.AWSRoute53RecordSet{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSRoute53RecordSet:
			results[name] = resource
		}
	}
	return results
}

// GetAWSRoute53RecordSetWithName retrieves all AWSRoute53RecordSet items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRoute53RecordSetWithName(name string) (*resources.AWSRoute53RecordSet, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSRoute53RecordSet:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSRoute53RecordSet not found", name)
}

// GetAllAWSRoute53RecordSetGroupResources retrieves all AWSRoute53RecordSetGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSRoute53RecordSetGroupResources() map[string]*resources.AWSRoute53RecordSetGroup {
	results := map[string]*resources.AWSRoute53RecordSetGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSRoute53RecordSetGroup:
			results[name] = resource
		}
	}
	return results
}

// GetAWSRoute53RecordSetGroupWithName retrieves all AWSRoute53RecordSetGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRoute53RecordSetGroupWithName(name string) (*resources.AWSRoute53RecordSetGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSRoute53RecordSetGroup:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSRoute53RecordSetGroup not found", name)
}

// GetAllAWSRoute53ResolverResolverEndpointResources retrieves all AWSRoute53ResolverResolverEndpoint items from an AWS CloudFormation template
func (t *Template) GetAllAWSRoute53ResolverResolverEndpointResources() map[string]*resources.AWSRoute53ResolverResolverEndpoint {
	results := map[string]*resources.AWSRoute53ResolverResolverEndpoint{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSRoute53ResolverResolverEndpoint:
			results[name] = resource
		}
	}
	return results
}

// GetAWSRoute53ResolverResolverEndpointWithName retrieves all AWSRoute53ResolverResolverEndpoint items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRoute53ResolverResolverEndpointWithName(name string) (*resources.AWSRoute53ResolverResolverEndpoint, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSRoute53ResolverResolverEndpoint:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSRoute53ResolverResolverEndpoint not found", name)
}

// GetAllAWSRoute53ResolverResolverRuleResources retrieves all AWSRoute53ResolverResolverRule items from an AWS CloudFormation template
func (t *Template) GetAllAWSRoute53ResolverResolverRuleResources() map[string]*resources.AWSRoute53ResolverResolverRule {
	results := map[string]*resources.AWSRoute53ResolverResolverRule{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSRoute53ResolverResolverRule:
			results[name] = resource
		}
	}
	return results
}

// GetAWSRoute53ResolverResolverRuleWithName retrieves all AWSRoute53ResolverResolverRule items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRoute53ResolverResolverRuleWithName(name string) (*resources.AWSRoute53ResolverResolverRule, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSRoute53ResolverResolverRule:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSRoute53ResolverResolverRule not found", name)
}

// GetAllAWSRoute53ResolverResolverRuleAssociationResources retrieves all AWSRoute53ResolverResolverRuleAssociation items from an AWS CloudFormation template
func (t *Template) GetAllAWSRoute53ResolverResolverRuleAssociationResources() map[string]*resources.AWSRoute53ResolverResolverRuleAssociation {
	results := map[string]*resources.AWSRoute53ResolverResolverRuleAssociation{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSRoute53ResolverResolverRuleAssociation:
			results[name] = resource
		}
	}
	return results
}

// GetAWSRoute53ResolverResolverRuleAssociationWithName retrieves all AWSRoute53ResolverResolverRuleAssociation items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRoute53ResolverResolverRuleAssociationWithName(name string) (*resources.AWSRoute53ResolverResolverRuleAssociation, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSRoute53ResolverResolverRuleAssociation:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSRoute53ResolverResolverRuleAssociation not found", name)
}

// GetAllAWSS3BucketResources retrieves all AWSS3Bucket items from an AWS CloudFormation template
func (t *Template) GetAllAWSS3BucketResources() map[string]*resources.AWSS3Bucket {
	results := map[string]*resources.AWSS3Bucket{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSS3Bucket:
			results[name] = resource
		}
	}
	return results
}

// GetAWSS3BucketWithName retrieves all AWSS3Bucket items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSS3BucketWithName(name string) (*resources.AWSS3Bucket, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSS3Bucket:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSS3Bucket not found", name)
}

// GetAllAWSS3BucketPolicyResources retrieves all AWSS3BucketPolicy items from an AWS CloudFormation template
func (t *Template) GetAllAWSS3BucketPolicyResources() map[string]*resources.AWSS3BucketPolicy {
	results := map[string]*resources.AWSS3BucketPolicy{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSS3BucketPolicy:
			results[name] = resource
		}
	}
	return results
}

// GetAWSS3BucketPolicyWithName retrieves all AWSS3BucketPolicy items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSS3BucketPolicyWithName(name string) (*resources.AWSS3BucketPolicy, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSS3BucketPolicy:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSS3BucketPolicy not found", name)
}

// GetAllAWSSDBDomainResources retrieves all AWSSDBDomain items from an AWS CloudFormation template
func (t *Template) GetAllAWSSDBDomainResources() map[string]*resources.AWSSDBDomain {
	results := map[string]*resources.AWSSDBDomain{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSSDBDomain:
			results[name] = resource
		}
	}
	return results
}

// GetAWSSDBDomainWithName retrieves all AWSSDBDomain items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSDBDomainWithName(name string) (*resources.AWSSDBDomain, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSSDBDomain:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSSDBDomain not found", name)
}

// GetAllAWSSESConfigurationSetResources retrieves all AWSSESConfigurationSet items from an AWS CloudFormation template
func (t *Template) GetAllAWSSESConfigurationSetResources() map[string]*resources.AWSSESConfigurationSet {
	results := map[string]*resources.AWSSESConfigurationSet{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSSESConfigurationSet:
			results[name] = resource
		}
	}
	return results
}

// GetAWSSESConfigurationSetWithName retrieves all AWSSESConfigurationSet items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSESConfigurationSetWithName(name string) (*resources.AWSSESConfigurationSet, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSSESConfigurationSet:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSSESConfigurationSet not found", name)
}

// GetAllAWSSESConfigurationSetEventDestinationResources retrieves all AWSSESConfigurationSetEventDestination items from an AWS CloudFormation template
func (t *Template) GetAllAWSSESConfigurationSetEventDestinationResources() map[string]*resources.AWSSESConfigurationSetEventDestination {
	results := map[string]*resources.AWSSESConfigurationSetEventDestination{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSSESConfigurationSetEventDestination:
			results[name] = resource
		}
	}
	return results
}

// GetAWSSESConfigurationSetEventDestinationWithName retrieves all AWSSESConfigurationSetEventDestination items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSESConfigurationSetEventDestinationWithName(name string) (*resources.AWSSESConfigurationSetEventDestination, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSSESConfigurationSetEventDestination:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSSESConfigurationSetEventDestination not found", name)
}

// GetAllAWSSESReceiptFilterResources retrieves all AWSSESReceiptFilter items from an AWS CloudFormation template
func (t *Template) GetAllAWSSESReceiptFilterResources() map[string]*resources.AWSSESReceiptFilter {
	results := map[string]*resources.AWSSESReceiptFilter{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSSESReceiptFilter:
			results[name] = resource
		}
	}
	return results
}

// GetAWSSESReceiptFilterWithName retrieves all AWSSESReceiptFilter items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSESReceiptFilterWithName(name string) (*resources.AWSSESReceiptFilter, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSSESReceiptFilter:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSSESReceiptFilter not found", name)
}

// GetAllAWSSESReceiptRuleResources retrieves all AWSSESReceiptRule items from an AWS CloudFormation template
func (t *Template) GetAllAWSSESReceiptRuleResources() map[string]*resources.AWSSESReceiptRule {
	results := map[string]*resources.AWSSESReceiptRule{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSSESReceiptRule:
			results[name] = resource
		}
	}
	return results
}

// GetAWSSESReceiptRuleWithName retrieves all AWSSESReceiptRule items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSESReceiptRuleWithName(name string) (*resources.AWSSESReceiptRule, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSSESReceiptRule:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSSESReceiptRule not found", name)
}

// GetAllAWSSESReceiptRuleSetResources retrieves all AWSSESReceiptRuleSet items from an AWS CloudFormation template
func (t *Template) GetAllAWSSESReceiptRuleSetResources() map[string]*resources.AWSSESReceiptRuleSet {
	results := map[string]*resources.AWSSESReceiptRuleSet{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSSESReceiptRuleSet:
			results[name] = resource
		}
	}
	return results
}

// GetAWSSESReceiptRuleSetWithName retrieves all AWSSESReceiptRuleSet items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSESReceiptRuleSetWithName(name string) (*resources.AWSSESReceiptRuleSet, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSSESReceiptRuleSet:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSSESReceiptRuleSet not found", name)
}

// GetAllAWSSESTemplateResources retrieves all AWSSESTemplate items from an AWS CloudFormation template
func (t *Template) GetAllAWSSESTemplateResources() map[string]*resources.AWSSESTemplate {
	results := map[string]*resources.AWSSESTemplate{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSSESTemplate:
			results[name] = resource
		}
	}
	return results
}

// GetAWSSESTemplateWithName retrieves all AWSSESTemplate items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSESTemplateWithName(name string) (*resources.AWSSESTemplate, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSSESTemplate:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSSESTemplate not found", name)
}

// GetAllAWSSNSSubscriptionResources retrieves all AWSSNSSubscription items from an AWS CloudFormation template
func (t *Template) GetAllAWSSNSSubscriptionResources() map[string]*resources.AWSSNSSubscription {
	results := map[string]*resources.AWSSNSSubscription{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSSNSSubscription:
			results[name] = resource
		}
	}
	return results
}

// GetAWSSNSSubscriptionWithName retrieves all AWSSNSSubscription items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSNSSubscriptionWithName(name string) (*resources.AWSSNSSubscription, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSSNSSubscription:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSSNSSubscription not found", name)
}

// GetAllAWSSNSTopicResources retrieves all AWSSNSTopic items from an AWS CloudFormation template
func (t *Template) GetAllAWSSNSTopicResources() map[string]*resources.AWSSNSTopic {
	results := map[string]*resources.AWSSNSTopic{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSSNSTopic:
			results[name] = resource
		}
	}
	return results
}

// GetAWSSNSTopicWithName retrieves all AWSSNSTopic items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSNSTopicWithName(name string) (*resources.AWSSNSTopic, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSSNSTopic:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSSNSTopic not found", name)
}

// GetAllAWSSNSTopicPolicyResources retrieves all AWSSNSTopicPolicy items from an AWS CloudFormation template
func (t *Template) GetAllAWSSNSTopicPolicyResources() map[string]*resources.AWSSNSTopicPolicy {
	results := map[string]*resources.AWSSNSTopicPolicy{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSSNSTopicPolicy:
			results[name] = resource
		}
	}
	return results
}

// GetAWSSNSTopicPolicyWithName retrieves all AWSSNSTopicPolicy items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSNSTopicPolicyWithName(name string) (*resources.AWSSNSTopicPolicy, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSSNSTopicPolicy:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSSNSTopicPolicy not found", name)
}

// GetAllAWSSQSQueueResources retrieves all AWSSQSQueue items from an AWS CloudFormation template
func (t *Template) GetAllAWSSQSQueueResources() map[string]*resources.AWSSQSQueue {
	results := map[string]*resources.AWSSQSQueue{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSSQSQueue:
			results[name] = resource
		}
	}
	return results
}

// GetAWSSQSQueueWithName retrieves all AWSSQSQueue items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSQSQueueWithName(name string) (*resources.AWSSQSQueue, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSSQSQueue:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSSQSQueue not found", name)
}

// GetAllAWSSQSQueuePolicyResources retrieves all AWSSQSQueuePolicy items from an AWS CloudFormation template
func (t *Template) GetAllAWSSQSQueuePolicyResources() map[string]*resources.AWSSQSQueuePolicy {
	results := map[string]*resources.AWSSQSQueuePolicy{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSSQSQueuePolicy:
			results[name] = resource
		}
	}
	return results
}

// GetAWSSQSQueuePolicyWithName retrieves all AWSSQSQueuePolicy items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSQSQueuePolicyWithName(name string) (*resources.AWSSQSQueuePolicy, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSSQSQueuePolicy:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSSQSQueuePolicy not found", name)
}

// GetAllAWSSSMAssociationResources retrieves all AWSSSMAssociation items from an AWS CloudFormation template
func (t *Template) GetAllAWSSSMAssociationResources() map[string]*resources.AWSSSMAssociation {
	results := map[string]*resources.AWSSSMAssociation{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSSSMAssociation:
			results[name] = resource
		}
	}
	return results
}

// GetAWSSSMAssociationWithName retrieves all AWSSSMAssociation items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSSMAssociationWithName(name string) (*resources.AWSSSMAssociation, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSSSMAssociation:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSSSMAssociation not found", name)
}

// GetAllAWSSSMDocumentResources retrieves all AWSSSMDocument items from an AWS CloudFormation template
func (t *Template) GetAllAWSSSMDocumentResources() map[string]*resources.AWSSSMDocument {
	results := map[string]*resources.AWSSSMDocument{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSSSMDocument:
			results[name] = resource
		}
	}
	return results
}

// GetAWSSSMDocumentWithName retrieves all AWSSSMDocument items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSSMDocumentWithName(name string) (*resources.AWSSSMDocument, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSSSMDocument:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSSSMDocument not found", name)
}

// GetAllAWSSSMMaintenanceWindowResources retrieves all AWSSSMMaintenanceWindow items from an AWS CloudFormation template
func (t *Template) GetAllAWSSSMMaintenanceWindowResources() map[string]*resources.AWSSSMMaintenanceWindow {
	results := map[string]*resources.AWSSSMMaintenanceWindow{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSSSMMaintenanceWindow:
			results[name] = resource
		}
	}
	return results
}

// GetAWSSSMMaintenanceWindowWithName retrieves all AWSSSMMaintenanceWindow items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSSMMaintenanceWindowWithName(name string) (*resources.AWSSSMMaintenanceWindow, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSSSMMaintenanceWindow:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSSSMMaintenanceWindow not found", name)
}

// GetAllAWSSSMMaintenanceWindowTaskResources retrieves all AWSSSMMaintenanceWindowTask items from an AWS CloudFormation template
func (t *Template) GetAllAWSSSMMaintenanceWindowTaskResources() map[string]*resources.AWSSSMMaintenanceWindowTask {
	results := map[string]*resources.AWSSSMMaintenanceWindowTask{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSSSMMaintenanceWindowTask:
			results[name] = resource
		}
	}
	return results
}

// GetAWSSSMMaintenanceWindowTaskWithName retrieves all AWSSSMMaintenanceWindowTask items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSSMMaintenanceWindowTaskWithName(name string) (*resources.AWSSSMMaintenanceWindowTask, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSSSMMaintenanceWindowTask:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSSSMMaintenanceWindowTask not found", name)
}

// GetAllAWSSSMParameterResources retrieves all AWSSSMParameter items from an AWS CloudFormation template
func (t *Template) GetAllAWSSSMParameterResources() map[string]*resources.AWSSSMParameter {
	results := map[string]*resources.AWSSSMParameter{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSSSMParameter:
			results[name] = resource
		}
	}
	return results
}

// GetAWSSSMParameterWithName retrieves all AWSSSMParameter items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSSMParameterWithName(name string) (*resources.AWSSSMParameter, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSSSMParameter:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSSSMParameter not found", name)
}

// GetAllAWSSSMPatchBaselineResources retrieves all AWSSSMPatchBaseline items from an AWS CloudFormation template
func (t *Template) GetAllAWSSSMPatchBaselineResources() map[string]*resources.AWSSSMPatchBaseline {
	results := map[string]*resources.AWSSSMPatchBaseline{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSSSMPatchBaseline:
			results[name] = resource
		}
	}
	return results
}

// GetAWSSSMPatchBaselineWithName retrieves all AWSSSMPatchBaseline items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSSMPatchBaselineWithName(name string) (*resources.AWSSSMPatchBaseline, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSSSMPatchBaseline:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSSSMPatchBaseline not found", name)
}

// GetAllAWSSSMResourceDataSyncResources retrieves all AWSSSMResourceDataSync items from an AWS CloudFormation template
func (t *Template) GetAllAWSSSMResourceDataSyncResources() map[string]*resources.AWSSSMResourceDataSync {
	results := map[string]*resources.AWSSSMResourceDataSync{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSSSMResourceDataSync:
			results[name] = resource
		}
	}
	return results
}

// GetAWSSSMResourceDataSyncWithName retrieves all AWSSSMResourceDataSync items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSSMResourceDataSyncWithName(name string) (*resources.AWSSSMResourceDataSync, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSSSMResourceDataSync:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSSSMResourceDataSync not found", name)
}

// GetAllAWSSageMakerEndpointResources retrieves all AWSSageMakerEndpoint items from an AWS CloudFormation template
func (t *Template) GetAllAWSSageMakerEndpointResources() map[string]*resources.AWSSageMakerEndpoint {
	results := map[string]*resources.AWSSageMakerEndpoint{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSSageMakerEndpoint:
			results[name] = resource
		}
	}
	return results
}

// GetAWSSageMakerEndpointWithName retrieves all AWSSageMakerEndpoint items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSageMakerEndpointWithName(name string) (*resources.AWSSageMakerEndpoint, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSSageMakerEndpoint:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSSageMakerEndpoint not found", name)
}

// GetAllAWSSageMakerEndpointConfigResources retrieves all AWSSageMakerEndpointConfig items from an AWS CloudFormation template
func (t *Template) GetAllAWSSageMakerEndpointConfigResources() map[string]*resources.AWSSageMakerEndpointConfig {
	results := map[string]*resources.AWSSageMakerEndpointConfig{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSSageMakerEndpointConfig:
			results[name] = resource
		}
	}
	return results
}

// GetAWSSageMakerEndpointConfigWithName retrieves all AWSSageMakerEndpointConfig items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSageMakerEndpointConfigWithName(name string) (*resources.AWSSageMakerEndpointConfig, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSSageMakerEndpointConfig:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSSageMakerEndpointConfig not found", name)
}

// GetAllAWSSageMakerModelResources retrieves all AWSSageMakerModel items from an AWS CloudFormation template
func (t *Template) GetAllAWSSageMakerModelResources() map[string]*resources.AWSSageMakerModel {
	results := map[string]*resources.AWSSageMakerModel{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSSageMakerModel:
			results[name] = resource
		}
	}
	return results
}

// GetAWSSageMakerModelWithName retrieves all AWSSageMakerModel items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSageMakerModelWithName(name string) (*resources.AWSSageMakerModel, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSSageMakerModel:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSSageMakerModel not found", name)
}

// GetAllAWSSageMakerNotebookInstanceResources retrieves all AWSSageMakerNotebookInstance items from an AWS CloudFormation template
func (t *Template) GetAllAWSSageMakerNotebookInstanceResources() map[string]*resources.AWSSageMakerNotebookInstance {
	results := map[string]*resources.AWSSageMakerNotebookInstance{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSSageMakerNotebookInstance:
			results[name] = resource
		}
	}
	return results
}

// GetAWSSageMakerNotebookInstanceWithName retrieves all AWSSageMakerNotebookInstance items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSageMakerNotebookInstanceWithName(name string) (*resources.AWSSageMakerNotebookInstance, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSSageMakerNotebookInstance:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSSageMakerNotebookInstance not found", name)
}

// GetAllAWSSageMakerNotebookInstanceLifecycleConfigResources retrieves all AWSSageMakerNotebookInstanceLifecycleConfig items from an AWS CloudFormation template
func (t *Template) GetAllAWSSageMakerNotebookInstanceLifecycleConfigResources() map[string]*resources.AWSSageMakerNotebookInstanceLifecycleConfig {
	results := map[string]*resources.AWSSageMakerNotebookInstanceLifecycleConfig{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSSageMakerNotebookInstanceLifecycleConfig:
			results[name] = resource
		}
	}
	return results
}

// GetAWSSageMakerNotebookInstanceLifecycleConfigWithName retrieves all AWSSageMakerNotebookInstanceLifecycleConfig items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSageMakerNotebookInstanceLifecycleConfigWithName(name string) (*resources.AWSSageMakerNotebookInstanceLifecycleConfig, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSSageMakerNotebookInstanceLifecycleConfig:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSSageMakerNotebookInstanceLifecycleConfig not found", name)
}

// GetAllAWSSecretsManagerResourcePolicyResources retrieves all AWSSecretsManagerResourcePolicy items from an AWS CloudFormation template
func (t *Template) GetAllAWSSecretsManagerResourcePolicyResources() map[string]*resources.AWSSecretsManagerResourcePolicy {
	results := map[string]*resources.AWSSecretsManagerResourcePolicy{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSSecretsManagerResourcePolicy:
			results[name] = resource
		}
	}
	return results
}

// GetAWSSecretsManagerResourcePolicyWithName retrieves all AWSSecretsManagerResourcePolicy items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSecretsManagerResourcePolicyWithName(name string) (*resources.AWSSecretsManagerResourcePolicy, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSSecretsManagerResourcePolicy:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSSecretsManagerResourcePolicy not found", name)
}

// GetAllAWSSecretsManagerRotationScheduleResources retrieves all AWSSecretsManagerRotationSchedule items from an AWS CloudFormation template
func (t *Template) GetAllAWSSecretsManagerRotationScheduleResources() map[string]*resources.AWSSecretsManagerRotationSchedule {
	results := map[string]*resources.AWSSecretsManagerRotationSchedule{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSSecretsManagerRotationSchedule:
			results[name] = resource
		}
	}
	return results
}

// GetAWSSecretsManagerRotationScheduleWithName retrieves all AWSSecretsManagerRotationSchedule items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSecretsManagerRotationScheduleWithName(name string) (*resources.AWSSecretsManagerRotationSchedule, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSSecretsManagerRotationSchedule:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSSecretsManagerRotationSchedule not found", name)
}

// GetAllAWSSecretsManagerSecretResources retrieves all AWSSecretsManagerSecret items from an AWS CloudFormation template
func (t *Template) GetAllAWSSecretsManagerSecretResources() map[string]*resources.AWSSecretsManagerSecret {
	results := map[string]*resources.AWSSecretsManagerSecret{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSSecretsManagerSecret:
			results[name] = resource
		}
	}
	return results
}

// GetAWSSecretsManagerSecretWithName retrieves all AWSSecretsManagerSecret items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSecretsManagerSecretWithName(name string) (*resources.AWSSecretsManagerSecret, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSSecretsManagerSecret:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSSecretsManagerSecret not found", name)
}

// GetAllAWSSecretsManagerSecretTargetAttachmentResources retrieves all AWSSecretsManagerSecretTargetAttachment items from an AWS CloudFormation template
func (t *Template) GetAllAWSSecretsManagerSecretTargetAttachmentResources() map[string]*resources.AWSSecretsManagerSecretTargetAttachment {
	results := map[string]*resources.AWSSecretsManagerSecretTargetAttachment{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSSecretsManagerSecretTargetAttachment:
			results[name] = resource
		}
	}
	return results
}

// GetAWSSecretsManagerSecretTargetAttachmentWithName retrieves all AWSSecretsManagerSecretTargetAttachment items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSecretsManagerSecretTargetAttachmentWithName(name string) (*resources.AWSSecretsManagerSecretTargetAttachment, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSSecretsManagerSecretTargetAttachment:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSSecretsManagerSecretTargetAttachment not found", name)
}

// GetAllAWSServerlessApiResources retrieves all AWSServerlessApi items from an AWS CloudFormation template
func (t *Template) GetAllAWSServerlessApiResources() map[string]*resources.AWSServerlessApi {
	results := map[string]*resources.AWSServerlessApi{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSServerlessApi:
			results[name] = resource
		}
	}
	return results
}

// GetAWSServerlessApiWithName retrieves all AWSServerlessApi items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSServerlessApiWithName(name string) (*resources.AWSServerlessApi, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSServerlessApi:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSServerlessApi not found", name)
}

// GetAllAWSServerlessApplicationResources retrieves all AWSServerlessApplication items from an AWS CloudFormation template
func (t *Template) GetAllAWSServerlessApplicationResources() map[string]*resources.AWSServerlessApplication {
	results := map[string]*resources.AWSServerlessApplication{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSServerlessApplication:
			results[name] = resource
		}
	}
	return results
}

// GetAWSServerlessApplicationWithName retrieves all AWSServerlessApplication items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSServerlessApplicationWithName(name string) (*resources.AWSServerlessApplication, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSServerlessApplication:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSServerlessApplication not found", name)
}

// GetAllAWSServerlessFunctionResources retrieves all AWSServerlessFunction items from an AWS CloudFormation template
func (t *Template) GetAllAWSServerlessFunctionResources() map[string]*resources.AWSServerlessFunction {
	results := map[string]*resources.AWSServerlessFunction{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSServerlessFunction:
			results[name] = resource
		}
	}
	return results
}

// GetAWSServerlessFunctionWithName retrieves all AWSServerlessFunction items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSServerlessFunctionWithName(name string) (*resources.AWSServerlessFunction, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSServerlessFunction:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSServerlessFunction not found", name)
}

// GetAllAWSServerlessLayerVersionResources retrieves all AWSServerlessLayerVersion items from an AWS CloudFormation template
func (t *Template) GetAllAWSServerlessLayerVersionResources() map[string]*resources.AWSServerlessLayerVersion {
	results := map[string]*resources.AWSServerlessLayerVersion{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSServerlessLayerVersion:
			results[name] = resource
		}
	}
	return results
}

// GetAWSServerlessLayerVersionWithName retrieves all AWSServerlessLayerVersion items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSServerlessLayerVersionWithName(name string) (*resources.AWSServerlessLayerVersion, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSServerlessLayerVersion:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSServerlessLayerVersion not found", name)
}

// GetAllAWSServerlessSimpleTableResources retrieves all AWSServerlessSimpleTable items from an AWS CloudFormation template
func (t *Template) GetAllAWSServerlessSimpleTableResources() map[string]*resources.AWSServerlessSimpleTable {
	results := map[string]*resources.AWSServerlessSimpleTable{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSServerlessSimpleTable:
			results[name] = resource
		}
	}
	return results
}

// GetAWSServerlessSimpleTableWithName retrieves all AWSServerlessSimpleTable items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSServerlessSimpleTableWithName(name string) (*resources.AWSServerlessSimpleTable, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSServerlessSimpleTable:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSServerlessSimpleTable not found", name)
}

// GetAllAWSServiceCatalogAcceptedPortfolioShareResources retrieves all AWSServiceCatalogAcceptedPortfolioShare items from an AWS CloudFormation template
func (t *Template) GetAllAWSServiceCatalogAcceptedPortfolioShareResources() map[string]*resources.AWSServiceCatalogAcceptedPortfolioShare {
	results := map[string]*resources.AWSServiceCatalogAcceptedPortfolioShare{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSServiceCatalogAcceptedPortfolioShare:
			results[name] = resource
		}
	}
	return results
}

// GetAWSServiceCatalogAcceptedPortfolioShareWithName retrieves all AWSServiceCatalogAcceptedPortfolioShare items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSServiceCatalogAcceptedPortfolioShareWithName(name string) (*resources.AWSServiceCatalogAcceptedPortfolioShare, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSServiceCatalogAcceptedPortfolioShare:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSServiceCatalogAcceptedPortfolioShare not found", name)
}

// GetAllAWSServiceCatalogCloudFormationProductResources retrieves all AWSServiceCatalogCloudFormationProduct items from an AWS CloudFormation template
func (t *Template) GetAllAWSServiceCatalogCloudFormationProductResources() map[string]*resources.AWSServiceCatalogCloudFormationProduct {
	results := map[string]*resources.AWSServiceCatalogCloudFormationProduct{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSServiceCatalogCloudFormationProduct:
			results[name] = resource
		}
	}
	return results
}

// GetAWSServiceCatalogCloudFormationProductWithName retrieves all AWSServiceCatalogCloudFormationProduct items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSServiceCatalogCloudFormationProductWithName(name string) (*resources.AWSServiceCatalogCloudFormationProduct, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSServiceCatalogCloudFormationProduct:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSServiceCatalogCloudFormationProduct not found", name)
}

// GetAllAWSServiceCatalogCloudFormationProvisionedProductResources retrieves all AWSServiceCatalogCloudFormationProvisionedProduct items from an AWS CloudFormation template
func (t *Template) GetAllAWSServiceCatalogCloudFormationProvisionedProductResources() map[string]*resources.AWSServiceCatalogCloudFormationProvisionedProduct {
	results := map[string]*resources.AWSServiceCatalogCloudFormationProvisionedProduct{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSServiceCatalogCloudFormationProvisionedProduct:
			results[name] = resource
		}
	}
	return results
}

// GetAWSServiceCatalogCloudFormationProvisionedProductWithName retrieves all AWSServiceCatalogCloudFormationProvisionedProduct items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSServiceCatalogCloudFormationProvisionedProductWithName(name string) (*resources.AWSServiceCatalogCloudFormationProvisionedProduct, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSServiceCatalogCloudFormationProvisionedProduct:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSServiceCatalogCloudFormationProvisionedProduct not found", name)
}

// GetAllAWSServiceCatalogLaunchNotificationConstraintResources retrieves all AWSServiceCatalogLaunchNotificationConstraint items from an AWS CloudFormation template
func (t *Template) GetAllAWSServiceCatalogLaunchNotificationConstraintResources() map[string]*resources.AWSServiceCatalogLaunchNotificationConstraint {
	results := map[string]*resources.AWSServiceCatalogLaunchNotificationConstraint{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSServiceCatalogLaunchNotificationConstraint:
			results[name] = resource
		}
	}
	return results
}

// GetAWSServiceCatalogLaunchNotificationConstraintWithName retrieves all AWSServiceCatalogLaunchNotificationConstraint items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSServiceCatalogLaunchNotificationConstraintWithName(name string) (*resources.AWSServiceCatalogLaunchNotificationConstraint, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSServiceCatalogLaunchNotificationConstraint:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSServiceCatalogLaunchNotificationConstraint not found", name)
}

// GetAllAWSServiceCatalogLaunchRoleConstraintResources retrieves all AWSServiceCatalogLaunchRoleConstraint items from an AWS CloudFormation template
func (t *Template) GetAllAWSServiceCatalogLaunchRoleConstraintResources() map[string]*resources.AWSServiceCatalogLaunchRoleConstraint {
	results := map[string]*resources.AWSServiceCatalogLaunchRoleConstraint{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSServiceCatalogLaunchRoleConstraint:
			results[name] = resource
		}
	}
	return results
}

// GetAWSServiceCatalogLaunchRoleConstraintWithName retrieves all AWSServiceCatalogLaunchRoleConstraint items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSServiceCatalogLaunchRoleConstraintWithName(name string) (*resources.AWSServiceCatalogLaunchRoleConstraint, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSServiceCatalogLaunchRoleConstraint:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSServiceCatalogLaunchRoleConstraint not found", name)
}

// GetAllAWSServiceCatalogLaunchTemplateConstraintResources retrieves all AWSServiceCatalogLaunchTemplateConstraint items from an AWS CloudFormation template
func (t *Template) GetAllAWSServiceCatalogLaunchTemplateConstraintResources() map[string]*resources.AWSServiceCatalogLaunchTemplateConstraint {
	results := map[string]*resources.AWSServiceCatalogLaunchTemplateConstraint{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSServiceCatalogLaunchTemplateConstraint:
			results[name] = resource
		}
	}
	return results
}

// GetAWSServiceCatalogLaunchTemplateConstraintWithName retrieves all AWSServiceCatalogLaunchTemplateConstraint items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSServiceCatalogLaunchTemplateConstraintWithName(name string) (*resources.AWSServiceCatalogLaunchTemplateConstraint, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSServiceCatalogLaunchTemplateConstraint:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSServiceCatalogLaunchTemplateConstraint not found", name)
}

// GetAllAWSServiceCatalogPortfolioResources retrieves all AWSServiceCatalogPortfolio items from an AWS CloudFormation template
func (t *Template) GetAllAWSServiceCatalogPortfolioResources() map[string]*resources.AWSServiceCatalogPortfolio {
	results := map[string]*resources.AWSServiceCatalogPortfolio{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSServiceCatalogPortfolio:
			results[name] = resource
		}
	}
	return results
}

// GetAWSServiceCatalogPortfolioWithName retrieves all AWSServiceCatalogPortfolio items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSServiceCatalogPortfolioWithName(name string) (*resources.AWSServiceCatalogPortfolio, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSServiceCatalogPortfolio:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSServiceCatalogPortfolio not found", name)
}

// GetAllAWSServiceCatalogPortfolioPrincipalAssociationResources retrieves all AWSServiceCatalogPortfolioPrincipalAssociation items from an AWS CloudFormation template
func (t *Template) GetAllAWSServiceCatalogPortfolioPrincipalAssociationResources() map[string]*resources.AWSServiceCatalogPortfolioPrincipalAssociation {
	results := map[string]*resources.AWSServiceCatalogPortfolioPrincipalAssociation{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSServiceCatalogPortfolioPrincipalAssociation:
			results[name] = resource
		}
	}
	return results
}

// GetAWSServiceCatalogPortfolioPrincipalAssociationWithName retrieves all AWSServiceCatalogPortfolioPrincipalAssociation items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSServiceCatalogPortfolioPrincipalAssociationWithName(name string) (*resources.AWSServiceCatalogPortfolioPrincipalAssociation, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSServiceCatalogPortfolioPrincipalAssociation:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSServiceCatalogPortfolioPrincipalAssociation not found", name)
}

// GetAllAWSServiceCatalogPortfolioProductAssociationResources retrieves all AWSServiceCatalogPortfolioProductAssociation items from an AWS CloudFormation template
func (t *Template) GetAllAWSServiceCatalogPortfolioProductAssociationResources() map[string]*resources.AWSServiceCatalogPortfolioProductAssociation {
	results := map[string]*resources.AWSServiceCatalogPortfolioProductAssociation{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSServiceCatalogPortfolioProductAssociation:
			results[name] = resource
		}
	}
	return results
}

// GetAWSServiceCatalogPortfolioProductAssociationWithName retrieves all AWSServiceCatalogPortfolioProductAssociation items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSServiceCatalogPortfolioProductAssociationWithName(name string) (*resources.AWSServiceCatalogPortfolioProductAssociation, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSServiceCatalogPortfolioProductAssociation:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSServiceCatalogPortfolioProductAssociation not found", name)
}

// GetAllAWSServiceCatalogPortfolioShareResources retrieves all AWSServiceCatalogPortfolioShare items from an AWS CloudFormation template
func (t *Template) GetAllAWSServiceCatalogPortfolioShareResources() map[string]*resources.AWSServiceCatalogPortfolioShare {
	results := map[string]*resources.AWSServiceCatalogPortfolioShare{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSServiceCatalogPortfolioShare:
			results[name] = resource
		}
	}
	return results
}

// GetAWSServiceCatalogPortfolioShareWithName retrieves all AWSServiceCatalogPortfolioShare items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSServiceCatalogPortfolioShareWithName(name string) (*resources.AWSServiceCatalogPortfolioShare, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSServiceCatalogPortfolioShare:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSServiceCatalogPortfolioShare not found", name)
}

// GetAllAWSServiceCatalogTagOptionResources retrieves all AWSServiceCatalogTagOption items from an AWS CloudFormation template
func (t *Template) GetAllAWSServiceCatalogTagOptionResources() map[string]*resources.AWSServiceCatalogTagOption {
	results := map[string]*resources.AWSServiceCatalogTagOption{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSServiceCatalogTagOption:
			results[name] = resource
		}
	}
	return results
}

// GetAWSServiceCatalogTagOptionWithName retrieves all AWSServiceCatalogTagOption items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSServiceCatalogTagOptionWithName(name string) (*resources.AWSServiceCatalogTagOption, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSServiceCatalogTagOption:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSServiceCatalogTagOption not found", name)
}

// GetAllAWSServiceCatalogTagOptionAssociationResources retrieves all AWSServiceCatalogTagOptionAssociation items from an AWS CloudFormation template
func (t *Template) GetAllAWSServiceCatalogTagOptionAssociationResources() map[string]*resources.AWSServiceCatalogTagOptionAssociation {
	results := map[string]*resources.AWSServiceCatalogTagOptionAssociation{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSServiceCatalogTagOptionAssociation:
			results[name] = resource
		}
	}
	return results
}

// GetAWSServiceCatalogTagOptionAssociationWithName retrieves all AWSServiceCatalogTagOptionAssociation items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSServiceCatalogTagOptionAssociationWithName(name string) (*resources.AWSServiceCatalogTagOptionAssociation, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSServiceCatalogTagOptionAssociation:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSServiceCatalogTagOptionAssociation not found", name)
}

// GetAllAWSServiceDiscoveryHttpNamespaceResources retrieves all AWSServiceDiscoveryHttpNamespace items from an AWS CloudFormation template
func (t *Template) GetAllAWSServiceDiscoveryHttpNamespaceResources() map[string]*resources.AWSServiceDiscoveryHttpNamespace {
	results := map[string]*resources.AWSServiceDiscoveryHttpNamespace{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSServiceDiscoveryHttpNamespace:
			results[name] = resource
		}
	}
	return results
}

// GetAWSServiceDiscoveryHttpNamespaceWithName retrieves all AWSServiceDiscoveryHttpNamespace items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSServiceDiscoveryHttpNamespaceWithName(name string) (*resources.AWSServiceDiscoveryHttpNamespace, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSServiceDiscoveryHttpNamespace:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSServiceDiscoveryHttpNamespace not found", name)
}

// GetAllAWSServiceDiscoveryInstanceResources retrieves all AWSServiceDiscoveryInstance items from an AWS CloudFormation template
func (t *Template) GetAllAWSServiceDiscoveryInstanceResources() map[string]*resources.AWSServiceDiscoveryInstance {
	results := map[string]*resources.AWSServiceDiscoveryInstance{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSServiceDiscoveryInstance:
			results[name] = resource
		}
	}
	return results
}

// GetAWSServiceDiscoveryInstanceWithName retrieves all AWSServiceDiscoveryInstance items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSServiceDiscoveryInstanceWithName(name string) (*resources.AWSServiceDiscoveryInstance, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSServiceDiscoveryInstance:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSServiceDiscoveryInstance not found", name)
}

// GetAllAWSServiceDiscoveryPrivateDnsNamespaceResources retrieves all AWSServiceDiscoveryPrivateDnsNamespace items from an AWS CloudFormation template
func (t *Template) GetAllAWSServiceDiscoveryPrivateDnsNamespaceResources() map[string]*resources.AWSServiceDiscoveryPrivateDnsNamespace {
	results := map[string]*resources.AWSServiceDiscoveryPrivateDnsNamespace{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSServiceDiscoveryPrivateDnsNamespace:
			results[name] = resource
		}
	}
	return results
}

// GetAWSServiceDiscoveryPrivateDnsNamespaceWithName retrieves all AWSServiceDiscoveryPrivateDnsNamespace items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSServiceDiscoveryPrivateDnsNamespaceWithName(name string) (*resources.AWSServiceDiscoveryPrivateDnsNamespace, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSServiceDiscoveryPrivateDnsNamespace:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSServiceDiscoveryPrivateDnsNamespace not found", name)
}

// GetAllAWSServiceDiscoveryPublicDnsNamespaceResources retrieves all AWSServiceDiscoveryPublicDnsNamespace items from an AWS CloudFormation template
func (t *Template) GetAllAWSServiceDiscoveryPublicDnsNamespaceResources() map[string]*resources.AWSServiceDiscoveryPublicDnsNamespace {
	results := map[string]*resources.AWSServiceDiscoveryPublicDnsNamespace{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSServiceDiscoveryPublicDnsNamespace:
			results[name] = resource
		}
	}
	return results
}

// GetAWSServiceDiscoveryPublicDnsNamespaceWithName retrieves all AWSServiceDiscoveryPublicDnsNamespace items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSServiceDiscoveryPublicDnsNamespaceWithName(name string) (*resources.AWSServiceDiscoveryPublicDnsNamespace, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSServiceDiscoveryPublicDnsNamespace:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSServiceDiscoveryPublicDnsNamespace not found", name)
}

// GetAllAWSServiceDiscoveryServiceResources retrieves all AWSServiceDiscoveryService items from an AWS CloudFormation template
func (t *Template) GetAllAWSServiceDiscoveryServiceResources() map[string]*resources.AWSServiceDiscoveryService {
	results := map[string]*resources.AWSServiceDiscoveryService{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSServiceDiscoveryService:
			results[name] = resource
		}
	}
	return results
}

// GetAWSServiceDiscoveryServiceWithName retrieves all AWSServiceDiscoveryService items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSServiceDiscoveryServiceWithName(name string) (*resources.AWSServiceDiscoveryService, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSServiceDiscoveryService:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSServiceDiscoveryService not found", name)
}

// GetAllAWSStepFunctionsActivityResources retrieves all AWSStepFunctionsActivity items from an AWS CloudFormation template
func (t *Template) GetAllAWSStepFunctionsActivityResources() map[string]*resources.AWSStepFunctionsActivity {
	results := map[string]*resources.AWSStepFunctionsActivity{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSStepFunctionsActivity:
			results[name] = resource
		}
	}
	return results
}

// GetAWSStepFunctionsActivityWithName retrieves all AWSStepFunctionsActivity items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSStepFunctionsActivityWithName(name string) (*resources.AWSStepFunctionsActivity, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSStepFunctionsActivity:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSStepFunctionsActivity not found", name)
}

// GetAllAWSStepFunctionsStateMachineResources retrieves all AWSStepFunctionsStateMachine items from an AWS CloudFormation template
func (t *Template) GetAllAWSStepFunctionsStateMachineResources() map[string]*resources.AWSStepFunctionsStateMachine {
	results := map[string]*resources.AWSStepFunctionsStateMachine{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSStepFunctionsStateMachine:
			results[name] = resource
		}
	}
	return results
}

// GetAWSStepFunctionsStateMachineWithName retrieves all AWSStepFunctionsStateMachine items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSStepFunctionsStateMachineWithName(name string) (*resources.AWSStepFunctionsStateMachine, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSStepFunctionsStateMachine:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSStepFunctionsStateMachine not found", name)
}

// GetAllAWSWAFByteMatchSetResources retrieves all AWSWAFByteMatchSet items from an AWS CloudFormation template
func (t *Template) GetAllAWSWAFByteMatchSetResources() map[string]*resources.AWSWAFByteMatchSet {
	results := map[string]*resources.AWSWAFByteMatchSet{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSWAFByteMatchSet:
			results[name] = resource
		}
	}
	return results
}

// GetAWSWAFByteMatchSetWithName retrieves all AWSWAFByteMatchSet items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSWAFByteMatchSetWithName(name string) (*resources.AWSWAFByteMatchSet, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSWAFByteMatchSet:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSWAFByteMatchSet not found", name)
}

// GetAllAWSWAFIPSetResources retrieves all AWSWAFIPSet items from an AWS CloudFormation template
func (t *Template) GetAllAWSWAFIPSetResources() map[string]*resources.AWSWAFIPSet {
	results := map[string]*resources.AWSWAFIPSet{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSWAFIPSet:
			results[name] = resource
		}
	}
	return results
}

// GetAWSWAFIPSetWithName retrieves all AWSWAFIPSet items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSWAFIPSetWithName(name string) (*resources.AWSWAFIPSet, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSWAFIPSet:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSWAFIPSet not found", name)
}

// GetAllAWSWAFRuleResources retrieves all AWSWAFRule items from an AWS CloudFormation template
func (t *Template) GetAllAWSWAFRuleResources() map[string]*resources.AWSWAFRule {
	results := map[string]*resources.AWSWAFRule{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSWAFRule:
			results[name] = resource
		}
	}
	return results
}

// GetAWSWAFRuleWithName retrieves all AWSWAFRule items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSWAFRuleWithName(name string) (*resources.AWSWAFRule, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSWAFRule:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSWAFRule not found", name)
}

// GetAllAWSWAFSizeConstraintSetResources retrieves all AWSWAFSizeConstraintSet items from an AWS CloudFormation template
func (t *Template) GetAllAWSWAFSizeConstraintSetResources() map[string]*resources.AWSWAFSizeConstraintSet {
	results := map[string]*resources.AWSWAFSizeConstraintSet{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSWAFSizeConstraintSet:
			results[name] = resource
		}
	}
	return results
}

// GetAWSWAFSizeConstraintSetWithName retrieves all AWSWAFSizeConstraintSet items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSWAFSizeConstraintSetWithName(name string) (*resources.AWSWAFSizeConstraintSet, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSWAFSizeConstraintSet:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSWAFSizeConstraintSet not found", name)
}

// GetAllAWSWAFSqlInjectionMatchSetResources retrieves all AWSWAFSqlInjectionMatchSet items from an AWS CloudFormation template
func (t *Template) GetAllAWSWAFSqlInjectionMatchSetResources() map[string]*resources.AWSWAFSqlInjectionMatchSet {
	results := map[string]*resources.AWSWAFSqlInjectionMatchSet{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSWAFSqlInjectionMatchSet:
			results[name] = resource
		}
	}
	return results
}

// GetAWSWAFSqlInjectionMatchSetWithName retrieves all AWSWAFSqlInjectionMatchSet items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSWAFSqlInjectionMatchSetWithName(name string) (*resources.AWSWAFSqlInjectionMatchSet, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSWAFSqlInjectionMatchSet:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSWAFSqlInjectionMatchSet not found", name)
}

// GetAllAWSWAFWebACLResources retrieves all AWSWAFWebACL items from an AWS CloudFormation template
func (t *Template) GetAllAWSWAFWebACLResources() map[string]*resources.AWSWAFWebACL {
	results := map[string]*resources.AWSWAFWebACL{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSWAFWebACL:
			results[name] = resource
		}
	}
	return results
}

// GetAWSWAFWebACLWithName retrieves all AWSWAFWebACL items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSWAFWebACLWithName(name string) (*resources.AWSWAFWebACL, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSWAFWebACL:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSWAFWebACL not found", name)
}

// GetAllAWSWAFXssMatchSetResources retrieves all AWSWAFXssMatchSet items from an AWS CloudFormation template
func (t *Template) GetAllAWSWAFXssMatchSetResources() map[string]*resources.AWSWAFXssMatchSet {
	results := map[string]*resources.AWSWAFXssMatchSet{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSWAFXssMatchSet:
			results[name] = resource
		}
	}
	return results
}

// GetAWSWAFXssMatchSetWithName retrieves all AWSWAFXssMatchSet items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSWAFXssMatchSetWithName(name string) (*resources.AWSWAFXssMatchSet, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSWAFXssMatchSet:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSWAFXssMatchSet not found", name)
}

// GetAllAWSWAFRegionalByteMatchSetResources retrieves all AWSWAFRegionalByteMatchSet items from an AWS CloudFormation template
func (t *Template) GetAllAWSWAFRegionalByteMatchSetResources() map[string]*resources.AWSWAFRegionalByteMatchSet {
	results := map[string]*resources.AWSWAFRegionalByteMatchSet{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSWAFRegionalByteMatchSet:
			results[name] = resource
		}
	}
	return results
}

// GetAWSWAFRegionalByteMatchSetWithName retrieves all AWSWAFRegionalByteMatchSet items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSWAFRegionalByteMatchSetWithName(name string) (*resources.AWSWAFRegionalByteMatchSet, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSWAFRegionalByteMatchSet:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSWAFRegionalByteMatchSet not found", name)
}

// GetAllAWSWAFRegionalIPSetResources retrieves all AWSWAFRegionalIPSet items from an AWS CloudFormation template
func (t *Template) GetAllAWSWAFRegionalIPSetResources() map[string]*resources.AWSWAFRegionalIPSet {
	results := map[string]*resources.AWSWAFRegionalIPSet{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSWAFRegionalIPSet:
			results[name] = resource
		}
	}
	return results
}

// GetAWSWAFRegionalIPSetWithName retrieves all AWSWAFRegionalIPSet items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSWAFRegionalIPSetWithName(name string) (*resources.AWSWAFRegionalIPSet, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSWAFRegionalIPSet:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSWAFRegionalIPSet not found", name)
}

// GetAllAWSWAFRegionalRuleResources retrieves all AWSWAFRegionalRule items from an AWS CloudFormation template
func (t *Template) GetAllAWSWAFRegionalRuleResources() map[string]*resources.AWSWAFRegionalRule {
	results := map[string]*resources.AWSWAFRegionalRule{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSWAFRegionalRule:
			results[name] = resource
		}
	}
	return results
}

// GetAWSWAFRegionalRuleWithName retrieves all AWSWAFRegionalRule items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSWAFRegionalRuleWithName(name string) (*resources.AWSWAFRegionalRule, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSWAFRegionalRule:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSWAFRegionalRule not found", name)
}

// GetAllAWSWAFRegionalSizeConstraintSetResources retrieves all AWSWAFRegionalSizeConstraintSet items from an AWS CloudFormation template
func (t *Template) GetAllAWSWAFRegionalSizeConstraintSetResources() map[string]*resources.AWSWAFRegionalSizeConstraintSet {
	results := map[string]*resources.AWSWAFRegionalSizeConstraintSet{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSWAFRegionalSizeConstraintSet:
			results[name] = resource
		}
	}
	return results
}

// GetAWSWAFRegionalSizeConstraintSetWithName retrieves all AWSWAFRegionalSizeConstraintSet items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSWAFRegionalSizeConstraintSetWithName(name string) (*resources.AWSWAFRegionalSizeConstraintSet, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSWAFRegionalSizeConstraintSet:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSWAFRegionalSizeConstraintSet not found", name)
}

// GetAllAWSWAFRegionalSqlInjectionMatchSetResources retrieves all AWSWAFRegionalSqlInjectionMatchSet items from an AWS CloudFormation template
func (t *Template) GetAllAWSWAFRegionalSqlInjectionMatchSetResources() map[string]*resources.AWSWAFRegionalSqlInjectionMatchSet {
	results := map[string]*resources.AWSWAFRegionalSqlInjectionMatchSet{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSWAFRegionalSqlInjectionMatchSet:
			results[name] = resource
		}
	}
	return results
}

// GetAWSWAFRegionalSqlInjectionMatchSetWithName retrieves all AWSWAFRegionalSqlInjectionMatchSet items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSWAFRegionalSqlInjectionMatchSetWithName(name string) (*resources.AWSWAFRegionalSqlInjectionMatchSet, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSWAFRegionalSqlInjectionMatchSet:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSWAFRegionalSqlInjectionMatchSet not found", name)
}

// GetAllAWSWAFRegionalWebACLResources retrieves all AWSWAFRegionalWebACL items from an AWS CloudFormation template
func (t *Template) GetAllAWSWAFRegionalWebACLResources() map[string]*resources.AWSWAFRegionalWebACL {
	results := map[string]*resources.AWSWAFRegionalWebACL{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSWAFRegionalWebACL:
			results[name] = resource
		}
	}
	return results
}

// GetAWSWAFRegionalWebACLWithName retrieves all AWSWAFRegionalWebACL items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSWAFRegionalWebACLWithName(name string) (*resources.AWSWAFRegionalWebACL, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSWAFRegionalWebACL:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSWAFRegionalWebACL not found", name)
}

// GetAllAWSWAFRegionalWebACLAssociationResources retrieves all AWSWAFRegionalWebACLAssociation items from an AWS CloudFormation template
func (t *Template) GetAllAWSWAFRegionalWebACLAssociationResources() map[string]*resources.AWSWAFRegionalWebACLAssociation {
	results := map[string]*resources.AWSWAFRegionalWebACLAssociation{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSWAFRegionalWebACLAssociation:
			results[name] = resource
		}
	}
	return results
}

// GetAWSWAFRegionalWebACLAssociationWithName retrieves all AWSWAFRegionalWebACLAssociation items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSWAFRegionalWebACLAssociationWithName(name string) (*resources.AWSWAFRegionalWebACLAssociation, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSWAFRegionalWebACLAssociation:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSWAFRegionalWebACLAssociation not found", name)
}

// GetAllAWSWAFRegionalXssMatchSetResources retrieves all AWSWAFRegionalXssMatchSet items from an AWS CloudFormation template
func (t *Template) GetAllAWSWAFRegionalXssMatchSetResources() map[string]*resources.AWSWAFRegionalXssMatchSet {
	results := map[string]*resources.AWSWAFRegionalXssMatchSet{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSWAFRegionalXssMatchSet:
			results[name] = resource
		}
	}
	return results
}

// GetAWSWAFRegionalXssMatchSetWithName retrieves all AWSWAFRegionalXssMatchSet items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSWAFRegionalXssMatchSetWithName(name string) (*resources.AWSWAFRegionalXssMatchSet, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSWAFRegionalXssMatchSet:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSWAFRegionalXssMatchSet not found", name)
}

// GetAllAWSWorkSpacesWorkspaceResources retrieves all AWSWorkSpacesWorkspace items from an AWS CloudFormation template
func (t *Template) GetAllAWSWorkSpacesWorkspaceResources() map[string]*resources.AWSWorkSpacesWorkspace {
	results := map[string]*resources.AWSWorkSpacesWorkspace{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AWSWorkSpacesWorkspace:
			results[name] = resource
		}
	}
	return results
}

// GetAWSWorkSpacesWorkspaceWithName retrieves all AWSWorkSpacesWorkspace items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSWorkSpacesWorkspaceWithName(name string) (*resources.AWSWorkSpacesWorkspace, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AWSWorkSpacesWorkspace:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AWSWorkSpacesWorkspace not found", name)
}

// GetAllAlexaASKSkillResources retrieves all AlexaASKSkill items from an AWS CloudFormation template
func (t *Template) GetAllAlexaASKSkillResources() map[string]*resources.AlexaASKSkill {
	results := map[string]*resources.AlexaASKSkill{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case *resources.AlexaASKSkill:
			results[name] = resource
		}
	}
	return results
}

// GetAlexaASKSkillWithName retrieves all AlexaASKSkill items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAlexaASKSkillWithName(name string) (*resources.AlexaASKSkill, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case *resources.AlexaASKSkill:
			return resource, nil
		}
	}
	return nil, fmt.Errorf("resource %q of type AlexaASKSkill not found", name)
}
