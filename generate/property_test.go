package main_test

import (
	"encoding/json"

	"github.com/awslabs/goformation/cloudformation"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Goformation Code Generator", func() {

	Context("with a single generated custom property", func() {

		Context("specified as a Go struct", func() {

			property := &cloudformation.AWSServerlessFunction_S3Location{
				Bucket:  cloudformation.NewString("test-bucket"),
				Key:     cloudformation.NewString("test-key"),
				Version: cloudformation.NewInteger(123),
			}
			expected := []byte(`{"Bucket":"test-bucket","Key":"test-key","Version":123}`)

			result, err := json.Marshal(property)
			It("should marshal to JSON successfully", func() {
				Expect(err).ToNot(HaveOccurred())
				Expect(result).To(Equal(expected))
			})

		})
		Context("specified as JSON", func() {

			property := []byte(`{"Bucket":"test-bucket","Key":"test-key","Version":123}`)
			expected := cloudformation.NewValue(
				&cloudformation.AWSServerlessFunction_S3Location{
					Bucket:  cloudformation.NewString("test-bucket"),
					Key:     cloudformation.NewString("test-key"),
					Version: cloudformation.NewInteger(123),
				},
			)

			result := cloudformation.NewValue(&cloudformation.AWSServerlessFunction_S3Location{})
			err := json.Unmarshal(property, result)
			It("should unmarshal to a Go struct successfully", func() {
				Expect(err).ToNot(HaveOccurred())
				expectedData, err := expected.MarshalJSON()
				Expect(err).ToNot(HaveOccurred())
				err = expected.UnmarshalJSON(expectedData)
				Expect(err).ToNot(HaveOccurred())
				Expect(result.Raw()).To(Equal(expected.Raw()))
			})

		})

	})

	Context("with a polymorphic property", func() {

		Context("with a primitive value", func() {

			Context("specified as a Go struct", func() {

				property := cloudformation.NewString("test-primitive-value")

				expected := []byte(`"test-primitive-value"`)
				result, err := json.Marshal(property)
				It("should marshal to JSON successfully", func() {
					Expect(err).ToNot(HaveOccurred())
					Expect(result).To(Equal(expected))
				})

			})

			Context("specified as JSON", func() {

				property := []byte(`"test-primitive-value"`)
				expected := cloudformation.NewString("test-primitive-value")

				var result cloudformation.Value
				err := json.Unmarshal(property, &result)
				It("should unmarshal to a Go struct successfully", func() {
					Expect(err).ToNot(HaveOccurred())
					Expect(result).To(Equal(*expected))
				})

			})

		})

		Context("with a custom type value", func() {

			Context("specified as a Go struct", func() {

				property := cloudformation.AWSServerlessFunction_S3Location{
					Bucket:  cloudformation.NewString("test-bucket"),
					Key:     cloudformation.NewString("test-key"),
					Version: cloudformation.NewInteger(123),
				}

				expected := []byte(`{"Bucket":"test-bucket","Key":"test-key","Version":123}`)

				result, err := json.Marshal(property)
				It("should marshal to JSON successfully", func() {
					Expect(err).ToNot(HaveOccurred())
					Expect(result).To(Equal(expected))
				})

			})

			Context("specified as JSON", func() {

				property := []byte(`{"Bucket":"test-bucket","Key":"test-key","Version":123}`)

				expected := cloudformation.NewValue(
					&cloudformation.AWSServerlessFunction_S3Location{
						Bucket:  cloudformation.NewString("test-bucket"),
						Key:     cloudformation.NewString("test-key"),
						Version: cloudformation.NewInteger(123),
					},
				)

				result := &cloudformation.Value{}
				err := result.UnmarshalJSON(property)
				It("should unmarshal to a Go struct successfully", func() {
					Expect(err).ToNot(HaveOccurred())
					expectedData, err := expected.MarshalJSON()
					Expect(err).ToNot(HaveOccurred())
					err = expected.UnmarshalJSON(expectedData)
					Expect(err).ToNot(HaveOccurred())
					Expect(result.Raw()).To(Equal(expected.Raw()))
				})

			})

		})

	})

})
