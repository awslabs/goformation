# GoFormation Versioning Changelog

# [7.2.0](https://github.com/awslabs/goformation/compare/v7.1.3...v7.2.0) (2023-01-13)


### Bug Fixes

* **ci:** bump semantic-release-action ([510f9c6](https://github.com/awslabs/goformation/commit/510f9c6e5ef243f3179ba4568b7365b3366c5c5d))


### Features

* switch go-yaml implementation to most recent version ([#535](https://github.com/awslabs/goformation/issues/535)) ([0ca6ce2](https://github.com/awslabs/goformation/commit/0ca6ce24fa93544b7fa52054bef2fa4e4f82bad2))

## [7.1.3](https://github.com/awslabs/goformation/compare/v7.1.2...v7.1.3) (2023-01-06)


### Bug Fixes

* **schema:** CloudFormation Updates ([#532](https://github.com/awslabs/goformation/issues/532)) ([d94f3f2](https://github.com/awslabs/goformation/commit/d94f3f2bc7ac2b6a28fb5155018664727a934f8c))

## [7.1.2](https://github.com/awslabs/goformation/compare/v7.1.1...v7.1.2) (2022-12-31)


### Bug Fixes

* **schema:** CloudFormation Updates ([#531](https://github.com/awslabs/goformation/issues/531)) ([83b04c9](https://github.com/awslabs/goformation/commit/83b04c96fb481eb7ac4a98397c2f88989ae21821))
* **schema:** CloudFormation Updates ([#531](https://github.com/awslabs/goformation/issues/531)) ([d72e4af](https://github.com/awslabs/goformation/commit/d72e4af5ba72e2c0f9905a4041677264fcbf7def))

## [7.1.1](https://github.com/awslabs/goformation/compare/v7.1.0...v7.1.1) (2022-12-23)


### Bug Fixes

* **schema:** CloudFormation Updates ([#530](https://github.com/awslabs/goformation/issues/530)) ([a65a99f](https://github.com/awslabs/goformation/commit/a65a99fcad4ae808b5c6797cbfe2f84893e5afac))

# [7.1.0](https://github.com/awslabs/goformation/compare/v7.0.7...v7.1.0) (2022-12-20)


### Features

* remove go1.16 and add go1.19 support ([#529](https://github.com/awslabs/goformation/issues/529)) ([583451d](https://github.com/awslabs/goformation/commit/583451d6377ec3e35a59b2eeb6b9685021f7e70c))

## [7.0.7](https://github.com/awslabs/goformation/compare/v7.0.6...v7.0.7) (2022-12-06)


### Bug Fixes

* **schema:** CloudFormation Updates ([#525](https://github.com/awslabs/goformation/issues/525)) ([fa6c239](https://github.com/awslabs/goformation/commit/fa6c23958205723db0e3f2d93edba97d848d6388))

## [7.0.6](https://github.com/awslabs/goformation/compare/v7.0.5...v7.0.6) (2022-12-01)


### Bug Fixes

* generation of AppFlow properties without type ([bfcd40f](https://github.com/awslabs/goformation/commit/bfcd40f16b3f395de2e02c483a15ddb52f550f8f))
* **schema:** CloudFormation Updates ([#524](https://github.com/awslabs/goformation/issues/524)) ([4fbffa5](https://github.com/awslabs/goformation/commit/4fbffa57d7591e5ae9b6ea08551f90727facc92a))

## [7.0.5](https://github.com/awslabs/goformation/compare/v7.0.4...v7.0.5) (2022-10-26)


### Bug Fixes

* **schema:** CloudFormation Updates ([065bf7e](https://github.com/awslabs/goformation/commit/065bf7e1f075d83c47da398ecf31b6d0582165d1))

## [7.0.4](https://github.com/awslabs/goformation/compare/v7.0.3...v7.0.4) (2022-10-22)


### Bug Fixes

* **schema:** CloudFormation Updates ([590b489](https://github.com/awslabs/goformation/commit/590b4893c639534520a3479ca6a00779fd194031))

## [7.0.3](https://github.com/awslabs/goformation/compare/v7.0.2...v7.0.3) (2022-10-15)


### Bug Fixes

* **schema:** CloudFormation Updates ([#510](https://github.com/awslabs/goformation/issues/510)) ([25e2ea4](https://github.com/awslabs/goformation/commit/25e2ea4445d58bf0a44c2b6f738338bfd6f2d0a5))

## [7.0.2](https://github.com/awslabs/goformation/compare/v7.0.1...v7.0.2) (2022-10-08)


### Bug Fixes

* **schema:** CloudFormation Updates ([998c192](https://github.com/awslabs/goformation/commit/998c1928695e32c580e20a2e79df0c99e4d366d9))

## [7.0.1](https://github.com/awslabs/goformation/compare/v7.0.0...v7.0.1) (2022-09-30)


### Bug Fixes

* **schema:** CloudFormation Updates ([cdcc602](https://github.com/awslabs/goformation/commit/cdcc6022ec63ecbe7a7c20c3c12bf68baad97dfe))

# [7.0.0](https://github.com/awslabs/goformation/compare/v6.11.0...v7.0.0) (2022-09-30)


* feat!: bump release to v7 ([a30de92](https://github.com/awslabs/goformation/commit/a30de9219e65b2b2f56df225eb687abd23aaa49c))


### BREAKING CHANGES

* Pointers are not used for Lists, Maps and interface{} members.

# [6.11.0](https://github.com/awslabs/goformation/compare/v6.10.0...v6.11.0) (2022-09-30)


### Features

* Release v7 ([#499](https://github.com/awslabs/goformation/issues/499)) ([28c3768](https://github.com/awslabs/goformation/commit/28c37685fdeabd5ec3745413601eabfb1d480f67))

# [7.0.0-beta.1](https://github.com/awslabs/goformation/compare/v6.10.0...v7.0.0-beta.1) (2022-09-29)


* feat(generate)!: Exclude Map and List from pointer ([aec3913](https://github.com/awslabs/goformation/commit/aec39137ec286627cc72cc106d9e4bc53e1fd16b))


### BREAKING CHANGES

* Goformation no longer generates pointer types for
Lists, Maps and interface{} types.

# [6.10.0](https://github.com/awslabs/goformation/compare/v6.9.0...v6.10.0) (2022-09-02)


### Features

* **sam:** add missing fields to ScheduledEvents and remove required versions from IAM ([33395af](https://github.com/awslabs/goformation/commit/33395afd746f31bfab003f56cf50cd984fc7d3b3))

# [6.9.0](https://github.com/awslabs/goformation/compare/v6.8.0...v6.9.0) (2022-09-02)


### Features

* **schema:** serverless http api cors configuration ([a90bb03](https://github.com/awslabs/goformation/commit/a90bb03ca3dfe7de1442c2cefe80aea4ea2aa211))

# [6.8.0](https://github.com/awslabs/goformation/compare/v6.7.2...v6.8.0) (2022-09-02)


### Features

* **generator:** remove generation of interface pointers ([315dde3](https://github.com/awslabs/goformation/commit/315dde31de45419c105e889a5bf29b6e4b1f42c0))

## [6.7.2](https://github.com/awslabs/goformation/compare/v6.7.1...v6.7.2) (2022-09-02)


### Bug Fixes

* **schema:** CloudFormation Updates ([7e80942](https://github.com/awslabs/goformation/commit/7e809428bd2c81f4ef52b9034bca39fe7e5c2442))

## [6.7.1](https://github.com/awslabs/goformation/compare/v6.7.0...v6.7.1) (2022-08-19)


### Bug Fixes

* **schema:** CloudFormation Updates ([7fedc99](https://github.com/awslabs/goformation/commit/7fedc99204bfa22d0ffc072f620c59a5bc9ac50d))

# [6.7.0](https://github.com/awslabs/goformation/compare/v6.6.6...v6.7.0) (2022-08-17)


### Features

* **if intrinsics:** generalized solution to support more types ([c66e47b](https://github.com/awslabs/goformation/commit/c66e47bcabe0c26f1dc81d9e671473a0c4fe7ef1))

## [6.6.6](https://github.com/awslabs/goformation/compare/v6.6.5...v6.6.6) (2022-08-15)


### Bug Fixes

* **generator:** remove unused import ([cf87ba6](https://github.com/awslabs/goformation/commit/cf87ba6f38737d60d15f7ece7384bee3a61a5b7c))
* **resource.template:** remove print to standard output when JSON unmarshal in a resource fails ([d64f719](https://github.com/awslabs/goformation/commit/d64f7194966f62cf8b7789cd8efaff4aaaaec0c3))
* **resource.template:** remove print to standard output when JSON unmarshal in a resource fails (output of go generate) ([c039ac4](https://github.com/awslabs/goformation/commit/c039ac4e9de2cb106f61252dbfe68164dbff38da))

## [6.6.5](https://github.com/awslabs/goformation/compare/v6.6.4...v6.6.5) (2022-08-15)


### Bug Fixes

* **schema:** CloudFormation Updates ([83f2d49](https://github.com/awslabs/goformation/commit/83f2d49319fef2d331eadab1c42b3e5774cd6613))

## [6.6.4](https://github.com/awslabs/goformation/compare/v6.6.3...v6.6.4) (2022-08-06)


### Bug Fixes

* **schema:** CloudFormation Updates ([bd8a2ac](https://github.com/awslabs/goformation/commit/bd8a2acffb312cedcc5c1a0c2a2fc4bfcd440e4a))

## [6.6.3](https://github.com/awslabs/goformation/compare/v6.6.2...v6.6.3) (2022-07-23)


### Bug Fixes

* **schema:** CloudFormation Updates ([af4f471](https://github.com/awslabs/goformation/commit/af4f471232b86642daa3fb2c8a4588f24e238530))

## [6.6.2](https://github.com/awslabs/goformation/compare/v6.6.1...v6.6.2) (2022-07-16)


### Bug Fixes

* **schema:** CloudFormation Updates ([5d02a2b](https://github.com/awslabs/goformation/commit/5d02a2b7d1eafb7b7d7dec821de9767a5d5f9b53))

## [6.6.1](https://github.com/awslabs/goformation/compare/v6.6.0...v6.6.1) (2022-07-15)


### Bug Fixes

* **schema:** CloudFormation Updates ([35a4b24](https://github.com/awslabs/goformation/commit/35a4b24e21ed39f1005d5c452804cbe2acdd4de1))

# [6.6.0](https://github.com/awslabs/goformation/compare/v6.5.4...v6.6.0) (2022-07-12)


### Bug Fixes

* **intrinsics:** split function ([286dd4c](https://github.com/awslabs/goformation/commit/286dd4c339fe9c64e31eb14d8eb621fad205fa7a))
* **intrinsics:** SplitPtr also as string ([86436f5](https://github.com/awslabs/goformation/commit/86436f5474683881db0e8906e8e62ce079e68208))


### Features

* **intrinsics:** add intrinsics ptr versions ([ffdc5af](https://github.com/awslabs/goformation/commit/ffdc5af445f9a055831ad1130000a343799f51b1))

## [6.5.4](https://github.com/awslabs/goformation/compare/v6.5.3...v6.5.4) (2022-07-08)


### Bug Fixes

* **schema:** CloudFormation Updates ([bc360ab](https://github.com/awslabs/goformation/commit/bc360ab35964268ef7c11110e2a1f6efe59280bd))

## [6.5.3](https://github.com/awslabs/goformation/compare/v6.5.2...v6.5.3) (2022-07-01)


### Bug Fixes

* **schema:** CloudFormation Updates ([0f4ade8](https://github.com/awslabs/goformation/commit/0f4ade82c3a9c3e139929a273c9e9bdc35612e2f))

## [6.5.2](https://github.com/awslabs/goformation/compare/v6.5.1...v6.5.2) (2022-06-24)


### Bug Fixes

* **schema:** CloudFormation Updates ([0de7ca4](https://github.com/awslabs/goformation/commit/0de7ca404158b1c1faa618fec93e032d42eb90d0))

## [6.5.1](https://github.com/awslabs/goformation/compare/v6.5.0...v6.5.1) (2022-06-17)


### Bug Fixes

* **schema:** CloudFormation Updates ([dcee612](https://github.com/awslabs/goformation/commit/dcee61222d00afa083519ecfe5cb60bd07b8b0ae))

# [6.5.0](https://github.com/awslabs/goformation/compare/v6.4.1...v6.5.0) (2022-06-14)


### Bug Fixes

* **schema:** re-generated schema ([eae0a91](https://github.com/awslabs/goformation/commit/eae0a9138e8e5417bd615bcee6ba02301e813ba2))


### Features

* **generator:** add support for new sagemaker properties ([bfd39c4](https://github.com/awslabs/goformation/commit/bfd39c480e6460ef1d1b3cf339ca30d8dd202c92))

## [6.4.1](https://github.com/awslabs/goformation/compare/v6.4.0...v6.4.1) (2022-06-02)


### Bug Fixes

* **schema:** Add AddDefaultAuthorizerToCorsPreflight to Serverless Auth ([637150c](https://github.com/awslabs/goformation/commit/637150c9f24eec6e4c86be18869340f12dd50763))

# [6.4.0](https://github.com/awslabs/goformation/compare/v6.3.0...v6.4.0) (2022-05-03)


### Features

* **schema:** Support condition properties in resources ([b3b7397](https://github.com/awslabs/goformation/commit/b3b739703c2436466cfa9416aeba96840e09423c))

# [6.3.0](https://github.com/awslabs/goformation/compare/v6.2.1...v6.3.0) (2022-05-03)


### Features

* **schema:** Support custom resource types ([1274ccd](https://github.com/awslabs/goformation/commit/1274ccd9317134522e92dd7fae402c7927076182))

## [6.2.1](https://github.com/awslabs/goformation/compare/v6.2.0...v6.2.1) (2022-04-30)


### Bug Fixes

* **schema:** CloudFormation Updates ([7858395](https://github.com/awslabs/goformation/commit/78583955ee7d232715a96a97c9bddf03f3a0ab8d))

# [6.2.0](https://github.com/awslabs/goformation/compare/v6.1.1...v6.2.0) (2022-04-26)


### Bug Fixes

* **schema:** Add RequestModel and RequestParameters for AWS::Serverless::Function.EventSource ([e0c2673](https://github.com/awslabs/goformation/commit/e0c26738fb4f8558cc8538b128a490924592b94c))
* **schema:** Fix JSON Schema generation commas for InclusivePrimitiveItemTypes ([28db940](https://github.com/awslabs/goformation/commit/28db9408fa2a0b4ffc9814dd0509d8026d9de223))


### Features

* **schema:** Add the ability to create items using pattern properties rather than normal references ([7b60160](https://github.com/awslabs/goformation/commit/7b6016076b6653a7c930b15a94262db16aa93b5a))

## [6.1.1](https://github.com/awslabs/goformation/compare/v6.1.0...v6.1.1) (2022-04-25)


### Bug Fixes

* **schema:** Add Version property into IAMPolicyDocument and fix Statement type ([846268a](https://github.com/awslabs/goformation/commit/846268a5028d9ca2cf93dde2b88ef7d3ba01e45d))

# [6.1.0](https://github.com/awslabs/goformation/compare/v6.0.15...v6.1.0) (2022-04-25)


### Bug Fixes

* **generate:** remove duplicated line ([a18d04c](https://github.com/awslabs/goformation/commit/a18d04c1844fbfacfdff709c51d09b6538603714))


### Features

* **schema:** Support generation of array items that should be combined in one anyOf ([d5e468f](https://github.com/awslabs/goformation/commit/d5e468fb4e4f83bcf08345232c8a7745fa0b7ff9))

## [6.0.15](https://github.com/awslabs/goformation/compare/v6.0.14...v6.0.15) (2022-04-22)


### Bug Fixes

* **schema:** string should be a primitivetype ([5fa746c](https://github.com/awslabs/goformation/commit/5fa746cc18b5ab0892f987f48d37989783787b33))

## [6.0.14](https://github.com/awslabs/goformation/compare/v6.0.13...v6.0.14) (2022-04-22)


### Bug Fixes

* **schema:** Add S3WritePolicy to sam policy template ([c9f775e](https://github.com/awslabs/goformation/commit/c9f775ede18530893f446eb7538f5364946fcf6f))

## [6.0.13](https://github.com/awslabs/goformation/compare/v6.0.12...v6.0.13) (2022-04-22)


### Bug Fixes

* **schema:** Add SSMParameterReadPolicy and AWSSecretsManagerGetSecretValuePolicy into AWS::Serverless::Function.SAMPolicyTemplate ([7a85ab9](https://github.com/awslabs/goformation/commit/7a85ab9fa99256ac54e0cf269ad10b1e5207c721))

## [6.0.12](https://github.com/awslabs/goformation/compare/v6.0.11...v6.0.12) (2022-04-22)


### Bug Fixes

* **schema:** Add DynamoDBWritePolicy to sam policy template ([6f08c13](https://github.com/awslabs/goformation/commit/6f08c136b112813bee5f329561ed25122f4a77cf))

## [6.0.11](https://github.com/awslabs/goformation/compare/v6.0.10...v6.0.11) (2022-04-22)


### Bug Fixes

* **schema:** Add Domain in AWS::Serverless::API schema ([dff256a](https://github.com/awslabs/goformation/commit/dff256ac71a956b090bf957765fa3ab32288b13f))

## [6.0.10](https://github.com/awslabs/goformation/compare/v6.0.9...v6.0.10) (2022-04-22)


### Bug Fixes

* **schema:** CloudFormation Updates ([319d00f](https://github.com/awslabs/goformation/commit/319d00fa1d94fdca42e6f9c58f407e1965a07ad7))

## [6.0.9](https://github.com/awslabs/goformation/compare/v6.0.8...v6.0.9) (2022-04-15)


### Bug Fixes

* **schema:** CloudFormation Updates ([8432365](https://github.com/awslabs/goformation/commit/8432365c586ee1f2a69df6afa8c699d90781f527))

## [6.0.8](https://github.com/awslabs/goformation/compare/v6.0.7...v6.0.8) (2022-04-14)


### Performance Improvements

* reduce JSON CloudFormation template size ([f893af7](https://github.com/awslabs/goformation/commit/f893af783a76ef4aff4c27710e2d4202c9fbf91a))

## [6.0.7](https://github.com/awslabs/goformation/compare/v6.0.6...v6.0.7) (2022-04-08)


### Bug Fixes

* **schema:** CloudFormation Updates ([68156bc](https://github.com/awslabs/goformation/commit/68156bc8aecb5ef63e6b9c2c597913d6d156e2cc))

## [6.0.6](https://github.com/awslabs/goformation/compare/v6.0.5...v6.0.6) (2022-04-02)


### Bug Fixes

* **schema:** CloudFormation Updates ([d2d083a](https://github.com/awslabs/goformation/commit/d2d083a8651a7c61f5f7748acb7ba0718e2ed5bd))

## [6.0.5](https://github.com/awslabs/goformation/compare/v6.0.4...v6.0.5) (2022-04-01)


### Bug Fixes

* **schema:** CloudFormation Updates ([9ce0a19](https://github.com/awslabs/goformation/commit/9ce0a19c9aa2c22c487aed1a3b2e24ea6cea731f))

## [6.0.4](https://github.com/awslabs/goformation/compare/v6.0.3...v6.0.4) (2022-03-26)


### Bug Fixes

* **schema:** CloudFormation Updates ([d59706b](https://github.com/awslabs/goformation/commit/d59706b488b4c73b92ef395ceef64160b2ecfb3d))

## [6.0.3](https://github.com/awslabs/goformation/compare/v6.0.2...v6.0.3) (2022-03-18)


### Bug Fixes

* **schema:** CloudFormation Updates ([801c7f8](https://github.com/awslabs/goformation/commit/801c7f886413d3015d69b8c49382631ad9c7528d))

## [6.0.2](https://github.com/awslabs/goformation/compare/v6.0.1...v6.0.2) (2022-03-11)


### Bug Fixes

* **schema:** CloudFormation Updates ([e06f6e2](https://github.com/awslabs/goformation/commit/e06f6e26aeeae9c5aa08181c9394be12ecd61aae))

## [6.0.1](https://github.com/awslabs/goformation/compare/v6.0.0...v6.0.1) (2022-03-04)


### Bug Fixes

* **schema:** CloudFormation Updates ([13095ef](https://github.com/awslabs/goformation/commit/13095ef6facddea686bc92d0b519e3578e5810d9))

# [6.0.0](https://github.com/awslabs/goformation/compare/v5.4.10...v6.0.0) (2022-02-25)


### Bug Fixes

* **generate:** DependsOn should also accept a string ([09908b6](https://github.com/awslabs/goformation/commit/09908b63bdf5cd326032008fa346e88c3765e1ae)), closes [#407](https://github.com/awslabs/goformation/issues/407)
* **policies:** re-create deleted files ([bdd5860](https://github.com/awslabs/goformation/commit/bdd58609a976e73fbd037866eadcb0db1f5a58a0))
* **schema:** generated schema acording to new rules ([d9dc863](https://github.com/awslabs/goformation/commit/d9dc86308402cf56b966a55a23fbcd2ca2c92291))
* **schema:** re-generate schema ([58dc56b](https://github.com/awslabs/goformation/commit/58dc56beb3a58acc391a5b3020f8cf3fabc19d6b))
* **schema:** regenerated with latest code ([33f99bf](https://github.com/awslabs/goformation/commit/33f99bf7d46415148bc42a1906a4b40843fa5c16))


### feature

* **types:** added utils to create pointer types ([4a68a60](https://github.com/awslabs/goformation/commit/4a68a60c13392084a04d71f69663b03959edd76d))


### Features

* **generate:** allow for optional params ([d9bfdff](https://github.com/awslabs/goformation/commit/d9bfdff22704915630f6341efbd11ceac7f3315d))
* **go:** drop support for go 1.13 and 1.14 ([05bb704](https://github.com/awslabs/goformation/commit/05bb704b8824417888bc7880f3478d9f2a380737))
* **go:** drop support for go 1.15 ([2e45a2b](https://github.com/awslabs/goformation/commit/2e45a2bbb85234c489f4c7cbb644c325a793f704))


### BREAKING CHANGES

* **generate:** DependsOn can now parse a single string instead of just a list of strings.
* **types:** use cloudformation.{String,Int,...} as helpers for
creating pointer types.
* **generate:** optional parameters are now marked as a pointer.

# [6.0.0-beta.1](https://github.com/awslabs/goformation/compare/v5.4.9...v6.0.0-beta.1) (2022-02-23)

### Bug Fixes

* **build:** build the project again ([5b8ac7d](https://github.com/awslabs/goformation/commit/5b8ac7dff9a7900d8b8d2dd849fd2d27a344b749))
* change version to v6 ([68254e2](https://github.com/awslabs/goformation/commit/68254e21559c1a5e5c43aa008b0d46cb1509210b))
* **policies:** add policies back ([2e26547](https://github.com/awslabs/goformation/commit/2e265475fa1102011fd18a430da9b51e78eefacb))
* **schema:** generated schema acording to new rules ([4321f42](https://github.com/awslabs/goformation/commit/4321f42ac0164cec7ea5bbebeafcc0792626fa06))
* **schema:** re-generate schema ([b0ccc0a](https://github.com/awslabs/goformation/commit/b0ccc0a1779eaa12cca9b9164f890d3b5fd9f648))
* **tags:** bring tag back. ([a3b2b57](https://github.com/awslabs/goformation/commit/a3b2b57e5939ce5a823e243c1e9470995a28f42d))


* feature(types)!: added utils to create pointer types ([bbf53ec](https://github.com/awslabs/goformation/commit/bbf53ec690533f207846100dc957e9bfea7a6b15))
* feat(generate)!: allow for optional params ([43a094e](https://github.com/awslabs/goformation/commit/43a094e3ca2e30a35645604372f1e4408ef49c6a))

### BREAKING CHANGES

* use cloudformation.{String,Int,...} as helpers for
creating pointer types.
* optional parameters are now marked as a pointer.

## [5.4.10](https://github.com/awslabs/goformation/compare/v5.4.9...v5.4.10) (2022-02-25)


### Bug Fixes

* **schema:** CloudFormation Updates ([c5b4ae3](https://github.com/awslabs/goformation/commit/c5b4ae3f38e61aedf2535a92204c28cac4a048ca))

## [5.4.9](https://github.com/awslabs/goformation/compare/v5.4.8...v5.4.9) (2022-02-21)


### Bug Fixes

* **schema:** CloudFormation Updates ([2f3e802](https://github.com/awslabs/goformation/commit/2f3e8023c760b0869ec947a520ffb60a4ed81c1c))

## [5.4.8](https://github.com/awslabs/goformation/compare/v5.4.7...v5.4.8) (2022-02-11)


### Bug Fixes

* **generator:** updated resources that support update/creation policy ([18c08b9](https://github.com/awslabs/goformation/commit/18c08b988d0ae39d07bfc3e653bddb2621db8276))

## [5.4.7](https://github.com/awslabs/goformation/compare/v5.4.6...v5.4.7) (2022-02-11)


### Bug Fixes

* **schema:** CloudFormation Updates ([bbbbbed](https://github.com/awslabs/goformation/commit/bbbbbeda4560c6d4296324edea11392ed31cc8de))

## [5.4.6](https://github.com/awslabs/goformation/compare/v5.4.5...v5.4.6) (2022-02-08)


### Bug Fixes

* **schema:** CloudFormation Updates ([#422](https://github.com/awslabs/goformation/issues/422)) ([61378b5](https://github.com/awslabs/goformation/commit/61378b5bad094f764289d53c87b2b2c91f118491))

## [5.4.5](https://github.com/awslabs/goformation/compare/v5.4.4...v5.4.5) (2022-01-18)


### Bug Fixes

* **schema:** AWS::CDK::Metadata resource should be automatically generated ([#421](https://github.com/awslabs/goformation/issues/421)) ([65569f7](https://github.com/awslabs/goformation/commit/65569f77c1708e52839ea486d5e10bdabdb24ce0)), closes [#418](https://github.com/awslabs/goformation/issues/418)

## [5.4.4](https://github.com/awslabs/goformation/compare/v5.4.3...v5.4.4) (2022-01-03)


### Bug Fixes

* **sam:** AWS::Serverless::Function Properties Architectures property should have a primitive type specified ([#420](https://github.com/awslabs/goformation/issues/420)) ([3aa91ed](https://github.com/awslabs/goformation/commit/3aa91edcb65cc6cfe8d0b2066944bc773acadc9f))

## [5.4.3](https://github.com/awslabs/goformation/compare/v5.4.2...v5.4.3) (2021-12-30)


### Bug Fixes

* **schema:** Add architectures support for sam functions ([#419](https://github.com/awslabs/goformation/issues/419)) ([b505b69](https://github.com/awslabs/goformation/commit/b505b694b7fc1d663282b67cdeba71013c4ea5ec))

## [5.4.2](https://github.com/awslabs/goformation/compare/v5.4.1...v5.4.2) (2021-12-30)


### Bug Fixes

* **schema:** Add cdkmetada resource ([#418](https://github.com/awslabs/goformation/issues/418)) ([3d1b1f9](https://github.com/awslabs/goformation/commit/3d1b1f91c8e592604becd080a7a61b4064784d03))

## [5.4.1](https://github.com/awslabs/goformation/compare/v5.4.0...v5.4.1) (2021-12-13)


### Bug Fixes

* **schema:** CloudFormation Updates ([#415](https://github.com/awslabs/goformation/issues/415)) ([e560a0f](https://github.com/awslabs/goformation/commit/e560a0fece3347f1aec326f643400423c2c7bc03))

# [5.4.0](https://github.com/awslabs/goformation/compare/v5.3.0...v5.4.0) (2021-11-22)


### Features

* **intrinsics:** Add SubVars to Sub with replacement variables ([#411](https://github.com/awslabs/goformation/issues/411)) ([0940790](https://github.com/awslabs/goformation/commit/094079090ca6b9ec8a98b1eada9fe3ca0e9c163c))

# [5.3.0](https://github.com/awslabs/goformation/compare/v5.2.12...v5.3.0) (2021-11-22)


### Features

* **intrinsics:** add support for base64 encoded string in instrinsic if function ([#414](https://github.com/awslabs/goformation/issues/414)) ([652501b](https://github.com/awslabs/goformation/commit/652501b0c05dd136f7faa5ab7b291f56386f0f3f)), closes [#412](https://github.com/awslabs/goformation/issues/412)

## [5.2.12](https://github.com/awslabs/goformation/compare/v5.2.11...v5.2.12) (2021-11-22)


### Bug Fixes

* **schema:** CloudFormation Updates ([#408](https://github.com/awslabs/goformation/issues/408)) ([2ffeeac](https://github.com/awslabs/goformation/commit/2ffeeac1cd90ec90807af825e49a15b6c8346431))

## [5.2.11](https://github.com/awslabs/goformation/compare/v5.2.10...v5.2.11) (2021-10-06)


### Bug Fixes

* **sam:** DestinationConfig shouldn't contain OnSuccess property ([#406](https://github.com/awslabs/goformation/issues/406)) ([6971966](https://github.com/awslabs/goformation/commit/69719668de26c4d9f54b4db25b2ce42313413375)), closes [#404](https://github.com/awslabs/goformation/issues/404)

## [5.2.10](https://github.com/awslabs/goformation/compare/v5.2.9...v5.2.10) (2021-10-05)


### Bug Fixes

* **schema:** CloudFormation Updates ([#401](https://github.com/awslabs/goformation/issues/401)) ([fa89e23](https://github.com/awslabs/goformation/commit/fa89e2381d69c66fb93496ab35b57a8f52772323))

## [5.2.9](https://github.com/awslabs/goformation/compare/v5.2.8...v5.2.9) (2021-09-03)


### Bug Fixes

* **schema:** CloudFormation Updates ([#400](https://github.com/awslabs/goformation/issues/400)) ([1606bbe](https://github.com/awslabs/goformation/commit/1606bbe6d8a2b0bd206505ec38542e8d7f6512d6))

## [5.2.8](https://github.com/awslabs/goformation/compare/v5.2.7...v5.2.8) (2021-08-27)


### Bug Fixes

* **schema:** CloudFormation Updates ([#398](https://github.com/awslabs/goformation/issues/398)) ([c7ebbd3](https://github.com/awslabs/goformation/commit/c7ebbd3328f69e1eaaa4b315f775b1caadcf4191))

## [5.2.7](https://github.com/awslabs/goformation/compare/v5.2.6...v5.2.7) (2021-08-11)


### Bug Fixes

* **schema:** CloudFormation Updates ([#393](https://github.com/awslabs/goformation/issues/393)) ([b005b8c](https://github.com/awslabs/goformation/commit/b005b8c6e8df3791af021593423f09cb70c316cf))

## [5.2.6](https://github.com/awslabs/goformation/compare/v5.2.5...v5.2.6) (2021-07-16)


### Bug Fixes

* **schema:** Allow any type for Parameter AllowedValues ([#392](https://github.com/awslabs/goformation/issues/392)) ([ccc7fb0](https://github.com/awslabs/goformation/commit/ccc7fb0509b1bdddf208e6a4463dd882ddfa4b69))
* **schema:** CloudFormation Updates ([#390](https://github.com/awslabs/goformation/issues/390)) ([ac83603](https://github.com/awslabs/goformation/commit/ac8360341ffc4e6576102a008a39b1674ea748fa))

## [5.2.5](https://github.com/awslabs/goformation/compare/v5.2.4...v5.2.5) (2021-07-16)


### Bug Fixes

* **schema:** Support Version field in custom resource ([#391](https://github.com/awslabs/goformation/issues/391)) ([eef8f36](https://github.com/awslabs/goformation/commit/eef8f361a61e0a3827b271019672c31a44937664))

## [5.2.4](https://github.com/awslabs/goformation/compare/v5.2.3...v5.2.4) (2021-07-05)


### Bug Fixes

* **schema:** CloudFormation Updates ([#388](https://github.com/awslabs/goformation/issues/388)) ([ae6ed10](https://github.com/awslabs/goformation/commit/ae6ed10530aaf15b34026207d5284e49659ce679))

## [5.2.3](https://github.com/awslabs/goformation/compare/v5.2.2...v5.2.3) (2021-06-29)


### Bug Fixes

* **schema:** CloudFormation Updates ([#387](https://github.com/awslabs/goformation/issues/387)) ([71f83ce](https://github.com/awslabs/goformation/commit/71f83ce7b426e45238f64fc11edc9045c280d992))

## [5.2.2](https://github.com/awslabs/goformation/compare/v5.2.1...v5.2.2) (2021-06-25)


### Bug Fixes

* **schema:** CloudFormation Updates ([#385](https://github.com/awslabs/goformation/issues/385)) ([8b5b816](https://github.com/awslabs/goformation/commit/8b5b8164786d7815c4e9f0b603d099198dfa5d46))

## [5.2.1](https://github.com/awslabs/goformation/compare/v5.2.0...v5.2.1) (2021-06-21)


### Bug Fixes

* **schema:** CloudFormation Updates ([#384](https://github.com/awslabs/goformation/issues/384)) ([823b7a4](https://github.com/awslabs/goformation/commit/823b7a4dba77bb702fdc0eb3e14e6f642afe5af0))

# [5.2.0](https://github.com/awslabs/goformation/compare/v5.1.0...v5.2.0) (2021-06-20)


### Features

* **schema:** Add AWS::Serverless::Function.Auth ([#373](https://github.com/awslabs/goformation/issues/373)) ([fc2877f](https://github.com/awslabs/goformation/commit/fc2877f69d241cb8462f57e865f0bc7b756610f4))

# [5.1.0](https://github.com/awslabs/goformation/compare/v5.0.0...v5.1.0) (2021-06-20)


### Features

* **schema:** CloudFormation Updates ([#383](https://github.com/awslabs/goformation/issues/383)) ([92fa1e3](https://github.com/awslabs/goformation/commit/92fa1e3e8adf95e85b0d39f029a80d51e450437c))

# [5.0.0](https://github.com/awslabs/goformation/compare/v4.19.5...v5.0.0) (2021-06-20)


### Features

* **schema:** Improve SAM Global support ([#376](https://github.com/awslabs/goformation/issues/376)) ([d56929b](https://github.com/awslabs/goformation/commit/d56929b269373078275877578e19f2f554ac6eff)), closes [#199](https://github.com/awslabs/goformation/issues/199) [#305](https://github.com/awslabs/goformation/issues/305)


### BREAKING CHANGES

* **schema:** Improved implementation of Globals (in SAM tempates)

This PR introduces a new implementation for both defining SAM templates with `Global` values, as well as parsing templates containing them.

Note: Globals only apply to SAM templates - if you are using regular CloudFormation templates this breaking change should not impact you. The only impact you might see is if you are creating `cloudFormation.Template` structs manually rather than using the `cloudformation.NewTemplate()` constructor. As part of this change, a new field (`Globals`) was added to the `Template{}` struct. If you are not using the constructor, your compiler will probably complain that this field is missing from the struct instantiation in your code.    

For more information about what `Globals` are in SAM templates, see this link:
https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/sam-specification-template-anatomy-globals.html

In previous versions of Goformation (before v5), goformation was able to parse SAM templates that contained a `Globals` section, however the implementation just overwrote the resource properties in the template, with the global values at parse time.

This meant that:

 - It was not possible to compose a template in Go that had a Globals section.
 - The JSON Schema generated by this repo, had no concept of Globals.

 **This new implementation DOES NOT does not overwrite the values in the template, like the previous implementation did**. It replaces the old implementation completely, and exposes Globals as a series of structs that are marshalled/unmarshalled to/from JSON/YAML.

 This allows you to compose a template:

```go
template := cloudformation.NewTemplate()

template.Globals["Function"] = &global.Function{
    Timeout: 1800,
}
```

As well as parse a JSON/YAML template, and inspect the global properties:

```go
template, err := goformation.Open("../some-template-with-globals.yml")
if err != nil {
    fmt.Printf("failed to open template: %s\n", err)
    os.Exit(1)
} 

fmt.Printf("%v", template.Globals["Function"])

// You can view the global as above, however it's type is downcast to a basic interface.
// If you want to inspect properties specific to the global type (e.g. the timeout for a Lambda function)
// then use the following helper methods to get the various global types from the template:

globalAPI, _:= template.GetServerlessGlobalApi()
globalFunction, _ := template.GetServerlessGlobalFunction()
globalHttpApi, _ := template.GetServerlessGlobalHttpApi()
globalSimpleTable, _ := template.GetServerlessGlobalSimpleTable()

fmt.Printf("Global Function Timeout: %d\n", globalFunction.Timeout)
```

## [4.19.5](https://github.com/awslabs/goformation/compare/v4.19.4...v4.19.5) (2021-06-10)


### Bug Fixes

* **schema:** CloudFormation Updates ([#381](https://github.com/awslabs/goformation/issues/381)) ([0714ecf](https://github.com/awslabs/goformation/commit/0714ecf5f1ae72116208da58fb5d9f1d33e3bb84))

## [4.19.4](https://github.com/awslabs/goformation/compare/v4.19.3...v4.19.4) (2021-05-31)


### Bug Fixes

* **schema:** CloudFormation Updates ([#375](https://github.com/awslabs/goformation/issues/375)) ([5d0d4f2](https://github.com/awslabs/goformation/commit/5d0d4f2f02b9f21d0db73ab6716578046d9abc21))

## [4.19.3](https://github.com/awslabs/goformation/compare/v4.19.2...v4.19.3) (2021-05-21)


### Bug Fixes

* **schema:** CloudFormation Updates ([#372](https://github.com/awslabs/goformation/issues/372)) ([95c8bf5](https://github.com/awslabs/goformation/commit/95c8bf594275cfcd9e3464dadb2e48c6c45357ff))

## [4.19.2](https://github.com/awslabs/goformation/compare/v4.19.1...v4.19.2) (2021-05-18)


### Bug Fixes

* **schema:** CloudFormation Updates ([#371](https://github.com/awslabs/goformation/issues/371)) ([b4a4521](https://github.com/awslabs/goformation/commit/b4a4521545419c5fb910a3a44a1e11e13c52f5e2))

## [4.19.1](https://github.com/awslabs/goformation/compare/v4.19.0...v4.19.1) (2021-04-17)


### Bug Fixes

* **schema:** CloudFormation Updates ([#367](https://github.com/awslabs/goformation/issues/367)) ([1a42e5b](https://github.com/awslabs/goformation/commit/1a42e5b414f371e2fd185d2fe796eca3934f1977))

# [4.19.0](https://github.com/awslabs/goformation/compare/v4.18.3...v4.19.0) (2021-04-17)


### Features

* **schema:** Improve cloudformation If to accept structs ([#368](https://github.com/awslabs/goformation/issues/368)) ([3c1bcd8](https://github.com/awslabs/goformation/commit/3c1bcd83e0e582aedeaf42f22e270b8a9f977863))

## [4.18.3](https://github.com/awslabs/goformation/compare/v4.18.2...v4.18.3) (2021-04-07)


### Bug Fixes

* **schema:** CloudFormation Updates ([#361](https://github.com/awslabs/goformation/issues/361)) ([ece24b8](https://github.com/awslabs/goformation/commit/ece24b8ffa21fac279e13fa20d9ccd2f21fe6e7e))

## [4.18.2](https://github.com/awslabs/goformation/compare/v4.18.1...v4.18.2) (2021-03-23)


### Bug Fixes

* **schema:** CloudFormation Updates ([#357](https://github.com/awslabs/goformation/issues/357)) ([4d320a5](https://github.com/awslabs/goformation/commit/4d320a56aa6fc79cd3a6a951c31a8c6580c68c99))

## [4.18.1](https://github.com/awslabs/goformation/compare/v4.18.0...v4.18.1) (2021-03-08)


### Bug Fixes

* **schema:** CloudFormation Updates ([#353](https://github.com/awslabs/goformation/issues/353)) ([eee9123](https://github.com/awslabs/goformation/commit/eee912305e08843fc112940333b93a6fc6aaf0ed))

# [4.18.0](https://github.com/awslabs/goformation/compare/v4.17.0...v4.18.0) (2021-03-08)


### Features

* **intrinsics:** Add cloudformation.TransformFn() ([#352](https://github.com/awslabs/goformation/issues/352)) ([9a1e331](https://github.com/awslabs/goformation/commit/9a1e3315fd0dcbe2e9a1fcbd0994de7470a2bd9a))

# [4.17.0](https://github.com/awslabs/goformation/compare/v4.16.4...v4.17.0) (2021-02-25)


### Bug Fixes

* **schema:** CloudFormation Updates ([#351](https://github.com/awslabs/goformation/issues/351)) ([53556a3](https://github.com/awslabs/goformation/commit/53556a3ff4a74666966e181d006448655533b2ba))


### Features

* **resources:** Add DependOn, DeletionPolicy and others to CustomResource ([#350](https://github.com/awslabs/goformation/issues/350)) ([6712019](https://github.com/awslabs/goformation/commit/671201953b1b8ba7f476e4f90a2cc5d46f542876))

## [4.16.4](https://github.com/awslabs/goformation/compare/v4.16.3...v4.16.4) (2021-02-20)


### Bug Fixes

* **schema:** CloudFormation Updates ([#349](https://github.com/awslabs/goformation/issues/349)) ([a012fde](https://github.com/awslabs/goformation/commit/a012fde9485a0fa185d9ed05bfc9fcf948fe15eb))

## [4.16.3](https://github.com/awslabs/goformation/compare/v4.16.2...v4.16.3) (2021-02-17)


### Bug Fixes

* **schema:** CloudFormation Updates ([#348](https://github.com/awslabs/goformation/issues/348)) ([5dd2417](https://github.com/awslabs/goformation/commit/5dd2417ddc3a063cbb2bba612938671e60bbb5ca))

## [4.16.2](https://github.com/awslabs/goformation/compare/v4.16.1...v4.16.2) (2021-02-09)


### Bug Fixes

* **schema:** CloudFormation Updates ([#347](https://github.com/awslabs/goformation/issues/347)) ([d49d514](https://github.com/awslabs/goformation/commit/d49d514d3cff735b2e9fffc261e6b6c05f1f8a5a))

## [4.16.1](https://github.com/awslabs/goformation/compare/v4.16.0...v4.16.1) (2021-02-09)


### Bug Fixes

* **template:** field Export on type Output should be pointer ([#299](https://github.com/awslabs/goformation/issues/299)) ([7d5870e](https://github.com/awslabs/goformation/commit/7d5870eff12e2ca9c011fb41b6748d0c3ec5f1d9)), closes [#294](https://github.com/awslabs/goformation/issues/294)

# [4.16.0](https://github.com/awslabs/goformation/compare/v4.15.9...v4.16.0) (2021-02-03)


### Features

* **intrinsics:** Allow for int in Fn::Equals ([#346](https://github.com/awslabs/goformation/issues/346)) ([dd6cd2d](https://github.com/awslabs/goformation/commit/dd6cd2d943eca6b4b0118066665a850ad0f2cc50))

## [4.15.9](https://github.com/awslabs/goformation/compare/v4.15.8...v4.15.9) (2021-01-29)


### Bug Fixes

* **schema:** CloudFormation Updates ([#342](https://github.com/awslabs/goformation/issues/342)) ([f047bed](https://github.com/awslabs/goformation/commit/f047bedef1193c8755bebac7f9edd3d872d52e1a))

## [4.15.8](https://github.com/awslabs/goformation/compare/v4.15.7...v4.15.8) (2021-01-23)


### Bug Fixes

* **schema:** CloudFormation Updates ([#341](https://github.com/awslabs/goformation/issues/341)) ([b65192b](https://github.com/awslabs/goformation/commit/b65192ba5d65da80351698c0eae194c02e5afcf6))

## [4.15.7](https://github.com/awslabs/goformation/compare/v4.15.6...v4.15.7) (2021-01-05)


### Bug Fixes

* **schema:** S3Location or String support for AWS::Serverless::LayerVersion.ContentUri ([#339](https://github.com/awslabs/goformation/issues/339)) ([6e39ebe](https://github.com/awslabs/goformation/commit/6e39ebe6e598e68f366b61b721e1ada2bdf268af)), closes [#337](https://github.com/awslabs/goformation/issues/337)

## [4.15.6](https://github.com/awslabs/goformation/compare/v4.15.5...v4.15.6) (2020-11-22)


### Bug Fixes

* **schema:** CloudFormation Updates ([#333](https://github.com/awslabs/goformation/issues/333)) ([0fec2c4](https://github.com/awslabs/goformation/commit/0fec2c466ad18f0852de492ab837c1cd51897b12))

## [4.15.5](https://github.com/awslabs/goformation/compare/v4.15.4...v4.15.5) (2020-11-06)


### Bug Fixes

* **schema:** CloudFormation Updates ([#331](https://github.com/awslabs/goformation/issues/331)) ([12f9c83](https://github.com/awslabs/goformation/commit/12f9c835cbd9b980e06e30c458f95b14e16a3771))

## [4.15.4](https://github.com/awslabs/goformation/compare/v4.15.3...v4.15.4) (2020-11-01)


### Bug Fixes

* **schema:** CloudFormation Updates ([#330](https://github.com/awslabs/goformation/issues/330)) ([4070319](https://github.com/awslabs/goformation/commit/40703191771425b3519128027478cf740d10f2d9))

## [4.15.3](https://github.com/awslabs/goformation/compare/v4.15.2...v4.15.3) (2020-10-23)


### Bug Fixes

* **schema:** CloudFormation Updates ([#329](https://github.com/awslabs/goformation/issues/329)) ([4c1362b](https://github.com/awslabs/goformation/commit/4c1362bac14a3d6fc2cebd56eddcaf440228a4e9))

## [4.15.2](https://github.com/awslabs/goformation/compare/v4.15.1...v4.15.2) (2020-10-11)


### Bug Fixes

* **schema:** CloudFormation Updates ([#320](https://github.com/awslabs/goformation/issues/320)) ([49879b4](https://github.com/awslabs/goformation/commit/49879b45f79dc9880071facd32d76cc4bf0570eb))

## [4.15.1](https://github.com/awslabs/goformation/compare/v4.15.0...v4.15.1) (2020-10-11)


### Bug Fixes

* **intrinsics:** Join function to allow to use parameters of type `List<>` ([#309](https://github.com/awslabs/goformation/issues/309)) ([6cc1cd3](https://github.com/awslabs/goformation/commit/6cc1cd329047227674caaf1b546066e3043c6616))

# [4.15.0](https://github.com/awslabs/goformation/compare/v4.14.0...v4.15.0) (2020-08-16)


### Features

* **schema:** dummy commit - trigger CI for schema update ([66bc344](https://github.com/awslabs/goformation/commit/66bc344ed6a17a613abbb4d217afcc0b8ea02b48))

# [4.14.0](https://github.com/awslabs/goformation/compare/v4.13.1...v4.14.0) (2020-07-26)


### Features

* **schema:** Add support for Template Outputs ([#291](https://github.com/awslabs/goformation/issues/291)) ([6875c50](https://github.com/awslabs/goformation/commit/6875c50d00d8e1af71d9bad5788446b29ab03513))

## [4.13.1](https://github.com/awslabs/goformation/compare/v4.13.0...v4.13.1) (2020-07-26)


### Bug Fixes

* **schema:** Add Change and Update policies to the Unmarshal method ([#288](https://github.com/awslabs/goformation/issues/288)) ([989b05f](https://github.com/awslabs/goformation/commit/989b05fa78cb9e72f6d59298fb8bb287612f322e))

# [4.13.0](https://github.com/awslabs/goformation/compare/v4.12.0...v4.13.0) (2020-07-26)


### Features

* **schema:** adding AWS::Serverless::StateMachine and FileSystemConfigs to Function ([#284](https://github.com/awslabs/goformation/issues/284)) ([d2d23ca](https://github.com/awslabs/goformation/commit/d2d23cafba606a8ea40649cc666073fa0e2d5ad3))

# [4.12.0](https://github.com/awslabs/goformation/compare/v4.11.0...v4.12.0) (2020-07-21)


### Features

* **schema:** Add new DynamoDBEvent options ([#289](https://github.com/awslabs/goformation/issues/289)) ([741228d](https://github.com/awslabs/goformation/commit/741228d6923ea10f1d22a901bbddf106d5c71cd7))

# [4.11.0](https://github.com/awslabs/goformation/compare/v4.10.1...v4.11.0) (2020-06-28)


### Features

* **schema:** CFN Updates ([#287](https://github.com/awslabs/goformation/issues/287)) ([9778479](https://github.com/awslabs/goformation/commit/97784795e35035b71b946d0ca69ef4d380d3b4a8))

## [4.10.1](https://github.com/awslabs/goformation/compare/v4.10.0...v4.10.1) (2020-06-22)


### Bug Fixes

* **generator:** update the generation making it easier to fix CF schema errors to generate ([#285](https://github.com/awslabs/goformation/issues/285)) ([6751e5b](https://github.com/awslabs/goformation/commit/6751e5b6ecbe1daee45171528cb1300efc6fb300))

# [4.10.0](https://github.com/awslabs/goformation/compare/v4.9.0...v4.10.0) (2020-06-22)


### Features

* **schema:** Serverless eventbridgeruleevent ([#279](https://github.com/awslabs/goformation/issues/279)) ([2a9e572](https://github.com/awslabs/goformation/commit/2a9e572313485023dc4e57cb8facda72a3571307))

# [4.9.0](https://github.com/awslabs/goformation/compare/v4.8.0...v4.9.0) (2020-06-22)


### Features

* **schema:** Add OpenApiVersion field to serverless Api ([#281](https://github.com/awslabs/goformation/issues/281)) ([bccc71b](https://github.com/awslabs/goformation/commit/bccc71b90531fb6bba8465b578fc2accc4dc6e34))

# [4.8.0](https://github.com/awslabs/goformation/compare/v4.7.1...v4.8.0) (2020-04-04)


### Features

* **schema:** Add UpdateReplacePolicy to the templates and the policies so that it is generated for every resource ([#272](https://github.com/awslabs/goformation/issues/272)) ([696c515](https://github.com/awslabs/goformation/commit/696c515bcbb07105683a328ef0e161d62146114b))

## [4.7.1](https://github.com/awslabs/goformation/compare/v4.7.0...v4.7.1) (2020-04-04)


### Bug Fixes

* **intrinsics:** change Fn::Sub to allow AWS pseudo parameters ([#275](https://github.com/awslabs/goformation/issues/275)) ([5a48c27](https://github.com/awslabs/goformation/commit/5a48c27630b945dcdc33133defd0241f898ccc52)), closes [#274](https://github.com/awslabs/goformation/issues/274) [#202](https://github.com/awslabs/goformation/issues/202)

# [4.7.0](https://github.com/awslabs/goformation/compare/v4.6.0...v4.7.0) (2020-02-28)


### Features

* **schema:** Added CloudWatch Logs event for SAM ([#271](https://github.com/awslabs/goformation/issues/271)) ([fedb013](https://github.com/awslabs/goformation/commit/fedb013e3b19ab1242cf8e3ae28a40240103d9b1))

# [4.6.0](https://github.com/awslabs/goformation/compare/v4.5.1...v4.6.0) (2020-02-22)


### Features

* **schema:** CloudFormation Updates (2020-02-22) ([#269](https://github.com/awslabs/goformation/issues/269)) ([ffd88a6](https://github.com/awslabs/goformation/commit/ffd88a6a9b0349853517e811169ee66804d79a2e))

## [4.5.1](https://github.com/awslabs/goformation/compare/v4.5.0...v4.5.1) (2020-02-14)


### Bug Fixes

* **schema, parser:** change Transform json schema to allow multiple macros ([#268](https://github.com/awslabs/goformation/issues/268)) ([072fc74](https://github.com/awslabs/goformation/commit/072fc74628c8ee9a603c2e502ac458af916afc07)), closes [#267](https://github.com/awslabs/goformation/issues/267)

# [4.5.0](https://github.com/awslabs/goformation/compare/v4.4.0...v4.5.0) (2020-02-13)


### Features

* **schema:** CloudFormation Updates (2020-02-13) ([#266](https://github.com/awslabs/goformation/issues/266)) ([bc75922](https://github.com/awslabs/goformation/commit/bc75922eb604d6e43f290912234a644c4d7584b5))

# [4.4.0](https://github.com/awslabs/goformation/compare/v4.3.0...v4.4.0) (2020-01-30)


### Features

* **schema:** CloudFormation Updates (2020-01-30) ([#263](https://github.com/awslabs/goformation/issues/263)) ([fda2d31](https://github.com/awslabs/goformation/commit/fda2d31f384eabbbf432ad1ee77ff8db6d0f2e73))

# [4.3.0](https://github.com/awslabs/goformation/compare/v4.2.0...v4.3.0) (2020-01-30)


### Features

* **schema:** add CloudFormation parameter type ([#259](https://github.com/awslabs/goformation/issues/259)) ([27fe204](https://github.com/awslabs/goformation/commit/27fe204f7addb8cb1bd6e977b0f717c04b09364a))

# [4.2.0](https://github.com/awslabs/goformation/compare/v4.1.0...v4.2.0) (2020-01-29)


### Features

* **parser:** Add support for Conditions ([#260](https://github.com/awslabs/goformation/issues/260)) ([1b00f17](https://github.com/awslabs/goformation/commit/1b00f17a33109023ad8a4471812448dc1d0db776))

# [4.1.0](https://github.com/awslabs/goformation/compare/v4.0.3...v4.1.0) (2019-12-09)


### Features

* **schema:** CloudFormation Updates (2019-12-09) ([#251](https://github.com/awslabs/goformation/issues/251)) ([a23ba41](https://github.com/awslabs/goformation/commit/a23ba416a24649c7296a0bc507c7940d9082ea30))

## [4.0.3](https://github.com/awslabs/goformation/compare/v4.0.2...v4.0.3) (2019-11-30)


### Bug Fixes

* **schema:** AWS::Serverless::Function S3 notification filters ([#249](https://github.com/awslabs/goformation/issues/249)) ([a50ef92](https://github.com/awslabs/goformation/commit/a50ef9291026420ea8a5e74790fc49b8a9c7fd85)), closes [#74](https://github.com/awslabs/goformation/issues/74)

## [4.0.2](https://github.com/awslabs/goformation/compare/v4.0.1...v4.0.2) (2019-11-30)


### Bug Fixes

* **schema:** AWS::Serverless:Api.Cors ([#246](https://github.com/awslabs/goformation/issues/246)) ([62fd56a](https://github.com/awslabs/goformation/commit/62fd56a62586c65722f99dbd4c8308ab42fcfc1d)), closes [#244](https://github.com/awslabs/goformation/issues/244)

## [4.0.1](https://github.com/awslabs/goformation/compare/v4.0.0...v4.0.1) (2019-11-30)


### Bug Fixes

* **schema:** AWS::Serverless::Api.MethodSettings should be a list ([a1f340a](https://github.com/awslabs/goformation/commit/a1f340a07e0ba4f21b8655da2c4d608849278901)), closes [#242](https://github.com/awslabs/goformation/issues/242)

# [4.0.0](https://github.com/awslabs/goformation/compare/v3.1.0...v4.0.0) (2019-11-30)


* Fix method conflicts (#245) ([d0b0a8b](https://github.com/awslabs/goformation/commit/d0b0a8bc322e27f72e840c9847f3c822d4efa933)), closes [#245](https://github.com/awslabs/goformation/issues/245) [#241](https://github.com/awslabs/goformation/issues/241) [#294](https://github.com/awslabs/goformation/issues/294)


### BREAKING CHANGES

* This change refactors the DependsOn, Metadata, CreationPolicy,
UpdatePolicy and DeletionPolicy methods on each resource to a new
name. This is required, as some CloudFormation resources use these
keywords as properties (AWS::AppMesh::Route.GrpcRouteMatch has a
Metadata field for example), which causes a conflict.

`resource.DependsOn()` method is refactored to `resource.AWSCloudFormationDependsOn` field.
`resource.SetDependsOn()` method is refactored to `resource.AWSCloudFormationDependsOn` field.
`resource.Metadata()` method is refactored to `resource.AWSCloudFormationMetadata` field.
`resource.SetMetadata()` method is refactored to `resource.AWSCloudFormationMetadata` field.
`resource.CreationPolicy()` method is refactored to `resource.AWSCloudFormationCreationPolicy` field.
`resource.SetCreationPolicy()` method is refactored to `resource.AWSCloudFormationCreationPolicy` field.
`resource.UpdatePolicy()` method is refactored to `resource.AWSCloudFormationUpdatePolicy` field.
`resource.SetUpdatePolicy()` method is refactored to `resource.AWSCloudFormationUpdatePolicy` field.
`resource.DeletionPolicy()` method is refactored to `resource.AWSCloudFormationDeletionPolicy` field.
`resource.SetDeletionPolicy()` method is refactored to `resource.AWSCloudFormationDeletionPolicy` field.

# [3.1.0](https://github.com/awslabs/goformation/compare/v3.0.1...v3.1.0) (2019-10-29)


### Features

* **schema:** AWS CloudFormation Update (2019-10-29) ([#239](https://github.com/awslabs/goformation/issues/239)) ([7ff8499](https://github.com/awslabs/goformation/commit/7ff84990c89e11815d22e06d377e110ae422cc17))

## [3.0.1](https://github.com/awslabs/goformation/compare/v3.0.0...v3.0.1) (2019-10-29)


### Bug Fixes

* **schema:** Ordered cloudformation/all.go file ([#238](https://github.com/awslabs/goformation/issues/238)) ([91254f3](https://github.com/awslabs/goformation/commit/91254f30925b89db5e79604d812a1ee9279267bd))

# [3.0.0](https://github.com/awslabs/goformation/compare/v2.3.1...v3.0.0) (2019-10-27)


* Group CloudFormation resources by AWS service name (#234) ([d0749e6](https://github.com/awslabs/goformation/commit/d0749e6a8fc5e7b0ddc301aef0170e12c7dc459c)), closes [#234](https://github.com/awslabs/goformation/issues/234)


### BREAKING CHANGES

* this change moves all Cloudformation resources to
packages based on the AWS service name. The main motivation for this is
that building goformation on some platforms (Windows) failed due to too
many files in the old cloudformation/resources package. This new package
style has a nice benefit of slightly nicer to use API, but is a breaking
change and will require refactoring existing codebases to update to v3.

Old usage:

```go
import "github.com/awslabs/goformation/v2/cloudformation/resources"

... snip ...

topic := &resources.AWSSNSTopic{}
```

New usage:

```go
import "github.com/awslabs/goformation/v4/cloudformation/sns"

...snip...

topic := &sns.Topic{}
```

Most tests are still failing at this point and need refactoring.

* fix(schema): Tag handling

Fixed tag handling for new grouped resources style (via new tags.Tag
struct).

* fix(schema): SAM specification

SAM Specification now generates nicely with new grouped resources
format. Also all tests are now passing \o/

# [2.3.0](https://github.com/awslabs/goformation/compare/v2.2.2...v2.3.0) (2019-03-20)


### Bug Fixes

* **parser:** Unmarshalling of resources with polymorphic properties (like S3 events) now works ([#188](https://github.com/awslabs/goformation/issues/188)) ([8eff90a](https://github.com/awslabs/goformation/commit/8eff90a))


### Features

* **sam:** Add support for `AWS::Serverless::Api.TracingEnabled`, `AWS::Serverless::Function.PermissionsBoundary`, `AWS::Serverless::Function.DynamoEvent.Enabled`, `AWS::Serverless::Function.KinesisEvent.Enabled`, and `AWS::Serverless::Function.SQSEvent.Enabled` ([#191](https://github.com/awslabs/goformation/issues/191)) ([38f0187](https://github.com/awslabs/goformation/commit/38f0187))
* **schema:** AWS CloudFormation Update (2019-03-15) ([#189](https://github.com/awslabs/goformation/issues/189)) ([8b332a4](https://github.com/awslabs/goformation/commit/8b332a4))

## [2.2.2](https://github.com/awslabs/goformation/compare/v2.2.1...v2.2.2) (2019-03-13)


### Bug Fixes

* **parser:** Select the correct AWS CloudFormation resource type based on similarity ([#183](https://github.com/awslabs/goformation/issues/183)) ([5749b23](https://github.com/awslabs/goformation/commit/5749b23))

## [2.2.1](https://github.com/awslabs/goformation/compare/v2.2.0...v2.2.1) (2019-03-10)


### Bug Fixes

* **parser:** fix invalid YAML template error for custom tag marshaler ([#177](https://github.com/awslabs/goformation/issues/177)) ([035d438](https://github.com/awslabs/goformation/commit/035d438))

# [2.2.0](https://github.com/awslabs/goformation/compare/v2.1.5...v2.2.0) (2019-03-10)


### Features

* **schema:** regenerated resources to apply SAM schema fixes from previous PR ([b30c019](https://github.com/awslabs/goformation/commit/b30c019))

## [2.1.5](https://github.com/awslabs/goformation/compare/v2.1.4...v2.1.5) (2019-03-10)


### Bug Fixes

* **parser:** do not break if a non-intrinsic `Condition` statement is found in a YAML template ([#169](https://github.com/awslabs/goformation/issues/169)) ([e4671e3](https://github.com/awslabs/goformation/commit/e4671e3))

## [2.1.4](https://github.com/awslabs/goformation/compare/v2.1.3...v2.1.4) (2019-03-10)


### Bug Fixes

* **schema:** fixed incorrect field type for AWS::Serverless::Application.Location ([#167](https://github.com/awslabs/goformation/issues/167)) ([3f1817b](https://github.com/awslabs/goformation/commit/3f1817b))

## [2.1.3](https://github.com/awslabs/goformation/compare/v2.1.2...v2.1.3) (2019-03-10)


### Bug Fixes

* **schema:** maps within YAML templates should allow unknown fields/properties ([3b6e359](https://github.com/awslabs/goformation/commit/3b6e359))

## [2.1.2](https://github.com/awslabs/goformation/compare/v2.1.1...v2.1.2) (2019-03-10)


### Bug Fixes

* **CI:** fix broken GitHub PR integration ([#185](https://github.com/awslabs/goformation/issues/185)) ([d42d00a](https://github.com/awslabs/goformation/commit/d42d00a))

## [2.1.1](https://github.com/awslabs/goformation/compare/v2.1.0...v2.1.1) (2019-03-10)


### Bug Fixes

* **CI:** only run semantic-release on push-to-master (not on pull requests) ([#184](https://github.com/awslabs/goformation/issues/184)) ([c83945a](https://github.com/awslabs/goformation/commit/c83945a))

# [2.1.0](https://github.com/awslabs/goformation/compare/v2.0.0...v2.1.0) (2019-03-10)


### Features

* **CI:** auto-generate AUTHORS.md file ([b37af7b](https://github.com/awslabs/goformation/commit/b37af7b))

# Semantic Versioning Changelog

# [2.0.0](https://github.com/awslabs/goformation/compare/v1.4.1...v2.0.0) (2019-03-10)


### Code Refactoring

* **generator:** moving resources and policies into their own packages ([#161](https://github.com/awslabs/goformation/issues/161)) ([03a0123](https://github.com/awslabs/goformation/commit/03a0123))


### BREAKING CHANGES

* **generator:** this PR refactors the auto-generated CloudFormation resources out of the cloudformation package and into a dedicated package (resources). This helps keep the auto generated files separate from others.

E.g. cloudformation.AWSSnsTopic{} becomes resources.AWSSnsTopic{}

## [1.4.1](https://github.com/awslabs/goformation/compare/v1.4.0...v1.4.1) (2019-03-10)


### Bug Fixes

* **spec:** corrected AWS::Serverless::Api.Auth.Authorizers to be of type JSON rather than string  ([#164](https://github.com/awslabs/goformation/issues/164)) ([4cf1bee](https://github.com/awslabs/goformation/commit/4cf1bee))

# [1.4.0](https://github.com/awslabs/goformation/compare/v1.3.0...v1.4.0) (2019-03-09)


### Features

* **parser:** Default to parsing as YAML unless the filename ends in .json ([#176](https://github.com/awslabs/goformation/issues/176)) ([42e7146](https://github.com/awslabs/goformation/commit/42e7146))

# [1.3.0](https://github.com/awslabs/goformation/compare/v1.2.1...v1.3.0) (2019-03-09)


### Bug Fixes

* **CI:** speed up PR builds by only downloading the cfn spec and regenerating resources on cron schedule (not on every build) ([7ae2a32](https://github.com/awslabs/goformation/commit/7ae2a32))
* **CI:** Update TravisCI configuration based on https://github.com/se… ([#180](https://github.com/awslabs/goformation/issues/180)) ([88e1e85](https://github.com/awslabs/goformation/commit/88e1e85))
* **CI:** Update TravisCI configuration for semantic-release to use jobs ([f6c2fee](https://github.com/awslabs/goformation/commit/f6c2fee))


### Features

* Added semantic-release CI setup ([a9b368a](https://github.com/awslabs/goformation/commit/a9b368a))
* Added semantic-release configuration file ([3b25fdb](https://github.com/awslabs/goformation/commit/3b25fdb))
