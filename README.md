# AWS GoFormation

TODO Definition

## Installation

The easiest way to get GoFormation is through `go get`:

```
go get github.com/awslabs/aws-goformation
```

This will get you a fresh copy of AWS GoFormation directly from the repository, into your `$GOPATH`.

## Usage

For using GoFormation you just need to reference in your Go app, whenever you want to use it:

### Opening the template file

If you want GoFormation to manage the opening of the file before the parsing, the way to proceed is `goformation.Open("template-file.yaml")`:

```
// my_file.go
package main

import "github.com/awslabs/goformation"

func main() {
	template, errors, logs := goformation.Open("my-template.yaml")
	// Do something with your template parsed.
}
```

### Parsing template's contents

If you rather use directly the contents of a template, then you should use `goformation.Parse("template contents...")`:

```
// my_file.go
package main

import "github.com/awslabs/goformation"

func main() {
	var textTemplate []byte = ... // Get your template's contents somewhere
	template, errors, logs := goformation.Open(textTemplate)
	// Do something with your template parsed.
}
```

### Using the parsed template

Once your template is parsed, you can easily get the resources parsed, to do any action with them - _NOTE: Currently, AWS GoFormation only supports `AWS::Serverless::Function` resources._ -:

```
// my_file.go
package main

import (
	"github.com/awslabs/goformation",
	. "github.com/awslabs/goformation/resources"
)

func main() {
	template, errors, logs := goformation.Open("my-template.yaml")
	// Verify there's no errors on parsing

	resources := template.Resources() // Get All resources
	functions := template.GetResourcesByType("AWS::Serverless::Function") // Get only Serverless Functions

	// Iterate over the parsed functions
	for fnName, fnData := range functions {	
		fnParsed := fnData.(AWSServerlessFunction)

		// Get function data
		fnRuntime := fnParsed.Runtime()
		log.Printf("Runtime: %s", fnRuntime) // Outputs the function's runtime
	}
}
```

## Documentation

TODO

### Parsing mechanism

TODO

#### Unmarshalling

TODO

##### Line detection

TODO

#### Scaffold

TODO

#### Post process

TODO

## Project Status

TODO

## Contributing

TODO