package main_test

import (
	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/paulmaddox/goformation/resources"
)

var _ = Describe("Goformation Code Generator", func() {

	Context("with a single generated custom property", func() {

		Context("specified as a Go struct", func() {

			property := &resources.AWSServerlessFunction_S3Location{
				Bucket:  "test-bucket",
				Key:     "test-key",
				Version: 123,
			}
			expected := []byte(`{"Bucket":"test-bucket","Key":"test-key","Version":123}`)

			result, err := json.Marshal(property)
			It("should marshal to JSON successfully", func() {
				Expect(result).To(Equal(expected))
				Expect(err).To(BeNil())
			})

		})
		Context("specified as JSON", func() {

			property := []byte(`{"Bucket":"test-bucket","Key":"test-key","Version":123}`)
			expected := &resources.AWSServerlessFunction_S3Location{
				Bucket:  "test-bucket",
				Key:     "test-key",
				Version: 123,
			}

			result := &resources.AWSServerlessFunction_S3Location{}
			err := json.Unmarshal(property, result)
			It("should unmarshal to a Go struct successfully", func() {
				Expect(result).To(Equal(expected))
				Expect(err).To(BeNil())
			})

		})

	})

	Context("with a polymorphic property", func() {

		Context("with a primitive value", func() {

			Context("specified as a Go struct", func() {

				value := "test-primitive-value"
				property := &resources.AWSServerlessFunction_StringOrS3Location{
					String: &value,
				}

				expected := []byte(`"test-primitive-value"`)
				result, err := json.Marshal(property)
				It("should marshal to JSON successfully", func() {
					Expect(result).To(Equal(expected))
					Expect(err).To(BeNil())
				})

			})

			Context("specified as JSON", func() {

				property := []byte(`"test-primitive-value"`)
				value := "test-primitive-value"
				expected := &resources.AWSServerlessFunction_StringOrS3Location{
					String: &value,
				}

				result := &resources.AWSServerlessFunction_StringOrS3Location{}
				err := json.Unmarshal(property, result)
				It("should unmarshal to a Go struct successfully", func() {
					Expect(result).To(Equal(expected))
					Expect(err).To(BeNil())
				})

			})

		})

		Context("with a custom type value", func() {

			Context("specified as a Go struct", func() {

				property := &resources.AWSServerlessFunction_StringOrS3Location{
					S3Location: &resources.AWSServerlessFunction_S3Location{
						Bucket:  "test-bucket",
						Key:     "test-key",
						Version: 123,
					},
				}

				expected := []byte(`{"Bucket":"test-bucket","Key":"test-key","Version":123}`)

				result, err := json.Marshal(property)
				It("should marshal to JSON successfully", func() {
					Expect(result).To(Equal(expected))
					Expect(err).To(BeNil())
				})

			})

			Context("specified as JSON", func() {

				property := []byte(`{"Bucket":"test-bucket","Key":"test-key","Version":123}`)

				expected := &resources.AWSServerlessFunction_StringOrS3Location{
					S3Location: &resources.AWSServerlessFunction_S3Location{
						Bucket:  "test-bucket",
						Key:     "test-key",
						Version: 123,
					},
				}

				result := &resources.AWSServerlessFunction_StringOrS3Location{}
				err := json.Unmarshal(property, result)
				It("should unmarshal to a Go struct successfully", func() {
					Expect(result).To(Equal(expected))
					Expect(err).To(BeNil())
				})

			})

		})

	})

})
