package main_test

import (
	"encoding/json"

	"github.com/vrealzhou/goformation/cloudformation/resources"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Resource", func() {

	Context("with a resource that contains a mix of primitive, custom and polymorphic properties", func() {

		Context("described as Go structs", func() {

			Context("with a simple primitive used for a polymorphic property", func() {

				codeuri := "s3://bucket/key"
				resource := &resources.AWSServerlessFunction{
					Runtime: "nodejs6.10",
					CodeUri: &resources.AWSServerlessFunction_CodeUri{
						String: &codeuri,
					},
				}

				expected := []byte(`{"Type":"AWS::Serverless::Function","Properties":{"CodeUri":"s3://bucket/key","Runtime":"nodejs6.10"}}`)

				result, err := json.Marshal(resource)
				It("should marshal to JSON successfully", func() {
					Expect(result).To(Equal(expected))
					Expect(err).To(BeNil())
				})

			})

			Context("with a custom type used for a polymorphic property", func() {

				resource := &resources.AWSServerlessFunction{
					Runtime: "nodejs6.10",
					CodeUri: &resources.AWSServerlessFunction_CodeUri{
						S3Location: &resources.AWSServerlessFunction_S3Location{
							Bucket:  "test-bucket",
							Key:     "test-key",
							Version: 123,
						},
					},
				}

				expected := []byte(`{"Type":"AWS::Serverless::Function","Properties":{"CodeUri":{"Bucket":"test-bucket","Key":"test-key","Version":123},"Runtime":"nodejs6.10"}}`)

				result, err := json.Marshal(resource)
				It("should marshal to JSON successfully", func() {
					Expect(result).To(Equal(expected))
					Expect(err).To(BeNil())
				})

			})

		})

	})

	Context("with a resource that has some resource attributes defined", func() {

		Context("described as Go structs", func() {

			Context("with a dependency on another resource", func() {

				resource := &resources.AWSEC2Instance{
					ImageId: "ami-0123456789",
				}
				resource.SetDependsOn([]string{"MyDependency"})

				expected := []byte(`{"Type":"AWS::EC2::Instance","Properties":{"ImageId":"ami-0123456789"},"DependsOn":["MyDependency"]}`)

				result, err := json.Marshal(resource)
				It("should marshal to JSON successfully", func() {
					Expect(result).To(Equal(expected))
					Expect(err).To(BeNil())
				})

			})

			Context("with a metadata attribute", func() {

				resource := &resources.AWSS3Bucket{
					BucketName: "MyBucket",
				}
				resource.SetMetadata(map[string]interface{}{"Object1": "Location1", "Object2": "Location2"})

				expected := []byte(`{"Type":"AWS::S3::Bucket","Properties":{"BucketName":"MyBucket"},"Metadata":{"Object1":"Location1","Object2":"Location2"}}`)

				result, err := json.Marshal(resource)
				It("should marshal to JSON successfully", func() {
					Expect(result).To(Equal(expected))
					Expect(err).To(BeNil())
				})

			})

		})

		Context("specified as JSON", func() {

			Context("with a dependency on another resource", func() {

				property := []byte(`{"Type":"AWS::EC2::Instance","Properties":{"ImageId":"ami-0123456789"},"DependsOn":["MyDependency"]}`)
				expected := &resources.AWSEC2Instance{
					ImageId: "ami-0123456789",
				}
				expected.SetDependsOn([]string{"MyDependency"})

				result := &resources.AWSEC2Instance{}
				err := json.Unmarshal(property, result)
				It("should unmarshal to a Go struct successfully", func() {
					Expect(result).To(Equal(expected))
					Expect(err).To(BeNil())
				})

			})

			Context("with a metadata attribute", func() {

				property := []byte(`{"Type":"AWS::S3::Bucket","Properties":{"BucketName":"MyBucket"},"Metadata":{"Object1":"Location1","Object2":"Location2"}}`)
				expected := &resources.AWSS3Bucket{
					BucketName: "MyBucket",
				}
				expected.SetMetadata(map[string]interface{}{"Object1": "Location1", "Object2": "Location2"})

				result := &resources.AWSS3Bucket{}
				err := json.Unmarshal(property, result)
				It("should unmarshal to a Go struct successfully", func() {
					Expect(result).To(Equal(expected))
					Expect(err).To(BeNil())
				})

			})

		})

	})

	Context("with a custom property type resource", func() {

		Context("described as Go structs", func() {

			Context("with a list type", func() {

				subproperty := &resources.AWSServerlessFunction_S3Event{
					Bucket: "my-bucket",
					Events: &resources.AWSServerlessFunction_Events{
						StringArray: &[]string{"s3:ObjectCreated:*", "s3:ObjectRemoved:*"},
					},
				}

				expected := []byte(`{"Bucket":"my-bucket","Events":["s3:ObjectCreated:*","s3:ObjectRemoved:*"]}`)

				result, err := json.Marshal(subproperty)
				It("should marshal to JSON successfully", func() {
					Expect(result).To(Equal(expected))
					Expect(err).To(BeNil())
				})

			})

			Context("with a primitive type", func() {

				event := "s3:ObjectCreated:*"
				subproperty := &resources.AWSServerlessFunction_S3Event{
					Bucket: "my-bucket",
					Events: &resources.AWSServerlessFunction_Events{
						String: &event,
					},
				}

				expected := []byte(`{"Bucket":"my-bucket","Events":"s3:ObjectCreated:*"}`)

				result, err := json.Marshal(subproperty)
				It("should marshal to JSON successfully", func() {
					Expect(result).To(Equal(expected))
					Expect(err).To(BeNil())
				})

			})

		})
	})
})
