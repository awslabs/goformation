package resources

// AWSServerlessFunction_StringOrS3Location is a helper struct that can hold either String or S3Location values
type AWSServerlessFunction_StringOrS3Location struct {
	String string

	S3Location AWSServerlessFunction_S3Location
}
