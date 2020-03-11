package s3connection

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// DownloadFromS3Bucket downloads xml file from s3 bucket
func DownloadFromS3Bucket() {

	bucket := "golang-xmlfilereader/TestResults"
	item := "a705804a-a676-4020-b36b-2482ca7bd540-testsuite.xml"

	file, err := os.Create(item)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String("ap-southeast-2"),
		Credentials: credentials.NewEnvCredentials(),
		// Credentials: (
		//     AWS_ACCESS_KEY: env.AWS_ACCESS_KEY,
		//     AWS_ACCESS_KEY_ID: env.AWS_SECRET_KEY
		// )
	})

	downloader := s3manager.NewDownloader(sess)

	numBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(item),
		})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Downloaded", file.Name(), numBytes, "bytes")
}
