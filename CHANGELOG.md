# GoFormation Versioning Changelog

# [2.3.0](https://github.com/vrealzhou/goformation/compare/v2.2.2...v2.3.0) (2019-03-20)


### Bug Fixes

* **parser:** Unmarshalling of resources with polymorphic properties (like S3 events) now works ([#188](https://github.com/vrealzhou/goformation/issues/188)) ([8eff90a](https://github.com/vrealzhou/goformation/commit/8eff90a))


### Features

* **sam:** Add support for `AWS::Serverless::Api.TracingEnabled`, `AWS::Serverless::Function.PermissionsBoundary`, `AWS::Serverless::Function.DynamoEvent.Enabled`, `AWS::Serverless::Function.KinesisEvent.Enabled`, and `AWS::Serverless::Function.SQSEvent.Enabled` ([#191](https://github.com/vrealzhou/goformation/issues/191)) ([38f0187](https://github.com/vrealzhou/goformation/commit/38f0187))
* **schema:** AWS CloudFormation Update (2019-03-15) ([#189](https://github.com/vrealzhou/goformation/issues/189)) ([8b332a4](https://github.com/vrealzhou/goformation/commit/8b332a4))

## [2.2.2](https://github.com/vrealzhou/goformation/compare/v2.2.1...v2.2.2) (2019-03-13)


### Bug Fixes

* **parser:** Select the correct AWS CloudFormation resource type based on similarity ([#183](https://github.com/vrealzhou/goformation/issues/183)) ([5749b23](https://github.com/vrealzhou/goformation/commit/5749b23))

## [2.2.1](https://github.com/vrealzhou/goformation/compare/v2.2.0...v2.2.1) (2019-03-10)


### Bug Fixes

* **parser:** fix invalid YAML template error for custom tag marshaler ([#177](https://github.com/vrealzhou/goformation/issues/177)) ([035d438](https://github.com/vrealzhou/goformation/commit/035d438))

# [2.2.0](https://github.com/vrealzhou/goformation/compare/v2.1.5...v2.2.0) (2019-03-10)


### Features

* **schema:** regenerated resources to apply SAM schema fixes from previous PR ([b30c019](https://github.com/vrealzhou/goformation/commit/b30c019))

## [2.1.5](https://github.com/vrealzhou/goformation/compare/v2.1.4...v2.1.5) (2019-03-10)


### Bug Fixes

* **parser:** do not break if a non-intrinsic `Condition` statement is found in a YAML template ([#169](https://github.com/vrealzhou/goformation/issues/169)) ([e4671e3](https://github.com/vrealzhou/goformation/commit/e4671e3))

## [2.1.4](https://github.com/vrealzhou/goformation/compare/v2.1.3...v2.1.4) (2019-03-10)


### Bug Fixes

* **schema:** fixed incorrect field type for AWS::Serverless::Application.Location ([#167](https://github.com/vrealzhou/goformation/issues/167)) ([3f1817b](https://github.com/vrealzhou/goformation/commit/3f1817b))

## [2.1.3](https://github.com/vrealzhou/goformation/compare/v2.1.2...v2.1.3) (2019-03-10)


### Bug Fixes

* **schema:** maps within YAML templates should allow unknown fields/properties ([3b6e359](https://github.com/vrealzhou/goformation/commit/3b6e359))

## [2.1.2](https://github.com/vrealzhou/goformation/compare/v2.1.1...v2.1.2) (2019-03-10)


### Bug Fixes

* **CI:** fix broken GitHub PR integration ([#185](https://github.com/vrealzhou/goformation/issues/185)) ([d42d00a](https://github.com/vrealzhou/goformation/commit/d42d00a))

## [2.1.1](https://github.com/vrealzhou/goformation/compare/v2.1.0...v2.1.1) (2019-03-10)


### Bug Fixes

* **CI:** only run semantic-release on push-to-master (not on pull requests) ([#184](https://github.com/vrealzhou/goformation/issues/184)) ([c83945a](https://github.com/vrealzhou/goformation/commit/c83945a))

# [2.1.0](https://github.com/vrealzhou/goformation/compare/v2.0.0...v2.1.0) (2019-03-10)


### Features

* **CI:** auto-generate AUTHORS.md file ([b37af7b](https://github.com/vrealzhou/goformation/commit/b37af7b))

# Semantic Versioning Changelog

# [2.0.0](https://github.com/vrealzhou/goformation/compare/v1.4.1...v2.0.0) (2019-03-10)


### Code Refactoring

* **generator:** moving resources and policies into their own packages ([#161](https://github.com/vrealzhou/goformation/issues/161)) ([03a0123](https://github.com/vrealzhou/goformation/commit/03a0123))


### BREAKING CHANGES

* **generator:** this PR refactors the auto-generated CloudFormation resources out of the cloudformation package and into a dedicated package (resources). This helps keep the auto generated files separate from others.

E.g. cloudformation.AWSSnsTopic{} becomes resources.AWSSnsTopic{}

## [1.4.1](https://github.com/vrealzhou/goformation/compare/v1.4.0...v1.4.1) (2019-03-10)


### Bug Fixes

* **spec:** corrected AWS::Serverless::Api.Auth.Authorizers to be of type JSON rather than string  ([#164](https://github.com/vrealzhou/goformation/issues/164)) ([4cf1bee](https://github.com/vrealzhou/goformation/commit/4cf1bee))

# [1.4.0](https://github.com/vrealzhou/goformation/compare/v1.3.0...v1.4.0) (2019-03-09)


### Features

* **parser:** Default to parsing as YAML unless the filename ends in .json ([#176](https://github.com/vrealzhou/goformation/issues/176)) ([42e7146](https://github.com/vrealzhou/goformation/commit/42e7146))

# [1.3.0](https://github.com/vrealzhou/goformation/compare/v1.2.1...v1.3.0) (2019-03-09)


### Bug Fixes

* **CI:** speed up PR builds by only downloading the cfn spec and regenerating resources on cron schedule (not on every build) ([7ae2a32](https://github.com/vrealzhou/goformation/commit/7ae2a32))
* **CI:** Update TravisCI configuration based on https://github.com/se… ([#180](https://github.com/vrealzhou/goformation/issues/180)) ([88e1e85](https://github.com/vrealzhou/goformation/commit/88e1e85))
* **CI:** Update TravisCI configuration for semantic-release to use jobs ([f6c2fee](https://github.com/vrealzhou/goformation/commit/f6c2fee))


### Features

* Added semantic-release CI setup ([a9b368a](https://github.com/vrealzhou/goformation/commit/a9b368a))
* Added semantic-release configuration file ([3b25fdb](https://github.com/vrealzhou/goformation/commit/3b25fdb))
