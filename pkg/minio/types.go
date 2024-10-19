package minio

import "time"

// UserIdentity represents the identity of the user involved in the event.
// It contains the PrincipalID, which uniquely identifies the user within the system.
type UserIdentity struct {
	PrincipalID string `json:"principalId"`
}

// RequestParameters holds information about the request that triggered the event.
// It includes details such as the source IP address from which the request originated.
type RequestParameters struct {
	SourceIPAddress string `json:"sourceIPAddress"`
}

// ResponseElements captures metadata returned in the response to the request.
// This includes specific identifiers like the Amazon S3 request ID and an additional ID used internally by Amazon.
type ResponseElements struct {
	XAmzRequestID string `json:"x-amz-request-id"`
	XAmzID2       string `json:"x-amz-id-2"`
}

// BucketOwnerIdentity represents the identity of the bucket's owner.
// It contains the PrincipalID of the owner.
type BucketOwnerIdentity struct {
	PrincipalID string `json:"principalId"`
}

// Bucket contains information about an Amazon S3 bucket involved in the event.
// This includes the bucket's name, owner identity, and the Amazon Resource Name (ARN).
type Bucket struct {
	Name          string              `json:"name"`
	OwnerIdentity BucketOwnerIdentity `json:"ownerIdentity"`
	Arn           string              `json:"arn"`
}

// S3Object represents an object stored in an S3 bucket.
// It includes details such as the object's key (its unique identifier in the bucket), size in bytes, the ETag for the object, and versioning information.
type S3Object struct {
	Key       string `json:"key"`
	Size      int    `json:"size"`
	ETag      string `json:"eTag"`
	VersionID string `json:"versionId"`
	Sequencer string `json:"sequencer"`
}

// S3 contains detailed information about an Amazon S3 event.
// It includes metadata about the S3 schema version, the specific configuration that triggered the event, the involved bucket, and the object within that bucket.
type S3 struct {
	S3SchemaVersion string   `json:"s3SchemaVersion"`
	ConfigurationID string   `json:"configurationId"`
	Bucket          Bucket   `json:"bucket"`
	Object          S3Object `json:"object"`
}

// GlacierRestoreEventData holds details about a restoration event from Amazon Glacier storage.
// This includes the expiration time of the restored object's lifecycle and the storage class used for restoration.
type GlacierRestoreEventData struct {
	LifecycleRestorationExpiryTime time.Time `json:"lifecycleRestorationExpiryTime"`
	LifecycleRestoreStorageClass   string    `json:"lifecycleRestoreStorageClass"`
}

// GlacierEventData represents the data associated with a Glacier event.
// It contains the restore event data specific to the lifecycle and storage class.
type GlacierEventData struct {
	RestoreEventData GlacierRestoreEventData `json:"restoreEventData"`
}

// Record represents a single event record in Amazon S3 or Glacier.
// It includes information such as the event version, source, region, time, event name, user identity, request and response details, and the associated S3 or Glacier event data.
type Record struct {
	EventVersion      string            `json:"eventVersion"`
	EventSource       string            `json:"eventSource"`
	AwsRegion         string            `json:"awsRegion"`
	EventTime         time.Time         `json:"eventTime"`
	EventName         string            `json:"eventName"`
	UserIdentity      UserIdentity      `json:"userIdentity"`
	RequestParameters RequestParameters `json:"requestParameters"`
	ResponseElements  ResponseElements  `json:"responseElements"`
	S3                S3                `json:"s3"`
	GlacierEventData  GlacierEventData  `json:"glacierEventData"`
}

// Event represents a collection of event records triggered in Amazon S3 or Glacier.
// It is structured as an array of Record structs, each containing detailed information about a specific event.
type Event struct {
	Records []Record `json:"Records"`
}
