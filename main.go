package main

import (
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// CreatePreSignedListObjectsV2 returns a pre-signed GET URL that
// lists objects in a bucket
func main() {

	// Initialize a session in us-east-1 that the SDK will use to load
	// credentials from ~/.aws/credentials
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewSharedCredentials("", "default"),
	})
	// Create S3 service client
	svc := s3.New(sess)

	req, _ := svc.ListObjectsV2Request(&s3.ListObjectsV2Input{
		Bucket: aws.String("bucket"), // Change this to the name of your bucket
		Prefix: aws.String(""),       // This allows for a subset of the bucket
	})

	Duration := 10000 // Time in seconds the pre-sign will be valid for

	urlStr, err := req.Presign(time.Second * time.Duration(Duration))

	if err != nil {
		log.Println("Failed to sign request\n", err)
	}

	fmt.Println(urlStr)
}
