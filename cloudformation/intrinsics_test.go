package cloudformation_test

import (
	"strings"

	"github.com/awslabs/goformation"
	"github.com/awslabs/goformation/cloudformation"
	"github.com/awslabs/goformation/intrinsics"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Goformation", func() {
	Context("with a YAML template that contains an intrinsic", func() {

		tests := []struct {
			Name     string
			Input    string
			Expected string
		}{
			{
				Name:     "Ref",
				Input:    `Description: !Ref tvalue`,
				Expected: `{"Description":{"Ref":"tvalue"}}`,
			},
			{
				Name:     "ImportValue",
				Input:    `Description: !ImportValue tvalue`,
				Expected: `{"Description":{"Fn::ImportValue":"tvalue"}}`,
			},
			{
				Name:     "Base64",
				Input:    `Description: !Base64 tvalue`,
				Expected: `{"Description":{"Fn::Base64":"tvalue"}}`,
			},
			{
				Name:     "GetAZs",
				Input:    `Description: !GetAZs tvalue`,
				Expected: `{"Description":{"Fn::GetAZs":"tvalue"}}`,
			},
			{
				Name:     "Sub",
				Input:    `Description: !Sub tvalue`,
				Expected: `{"Description":{"Fn::Sub":"tvalue"}}`,
			},
			{
				Name:     "GetAtt",
				Input:    `Description: !GetAtt [object, property]`,
				Expected: `{"Description":{"Fn::GetAtt":["object","property"]}}`,
			},
			{
				Name:     "Split",
				Input:    `Description: !Split [d, sss]`,
				Expected: `{"Description":{"Fn::Split":["d","sss"]}}`,
			},
			{
				Name:     "Equals",
				Input:    `Description: !Equals [a, b]`,
				Expected: `{"Description":{"Fn::Equals":["a","b"]}}`,
			},
			{
				Name:     "CIDR",
				Input:    `Description: !Cidr [a, b, c]`,
				Expected: `{"Description":{"Fn::Cidr":["a","b","c"]}}`,
			},
			{
				Name:     "FindInMap",
				Input:    `Description: !FindInMap [a, b, c]`,
				Expected: `{"Description":{"Fn::FindInMap":["a","b","c"]}}`,
			},
			{
				Name:     "If",
				Input:    `Description: !If [a, b, c]`,
				Expected: `{"Description":{"Fn::If":["a","b","c"]}}`,
			},
			{
				Name:     "Join",
				Input:    `Description: !Join [a, [b, c]]`,
				Expected: `{"Description":{"Fn::Join":["a",["b","c"]]}}`,
			},
			{
				Name:     "Select",
				Input:    `Description: !Select [a, [b, c]]`,
				Expected: `{"Description":{"Fn::Select":["a",["b","c"]]}}`,
			},
			{
				Name:     "And",
				Input:    `Description: !And [a, b, c]`,
				Expected: `{"Description":{"Fn::And":["a","b","c"]}}`,
			},
			{
				Name:     "Not",
				Input:    `Description: !Not [a, b, c]`,
				Expected: `{"Description":{"Fn::Not":["a","b","c"]}}`,
			},
			{
				Name:     "Or",
				Input:    `Description: !Or [a, b, c]`,
				Expected: `{"Description":{"Fn::Or":["a","b","c"]}}`,
			},
		}

		for _, test := range tests {
			test := test

			It("should replace "+test.Name+" with the JSON expanded value", func() {
				template, err := goformation.ParseYAMLWithOptions([]byte(test.Input), &intrinsics.ProcessorOptions{
					IntrinsicHandlerOverrides: cloudformation.EncoderIntrinsics,
				})

				Expect(err).To(BeNil())

				raw, err := template.JSON()
				output := string(raw)
				Expect(err).To(BeNil())
				output = strings.Replace(output, " ", "", -1)
				output = strings.Replace(output, "\n", "", -1)

				Expect(test.Expected).To(BeEquivalentTo(output))
			})
		}
	})
})
