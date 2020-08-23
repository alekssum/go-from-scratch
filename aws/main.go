package main

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func main() {

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("eu-north-1"),
		Credentials: credentials.NewStaticCredentials("AKIAITRITV6KPONAN2KA", "dNu+SfQ+3Jr05oGJ4+t/uWrQ97nRCVZgSVhMWcGu", ""),
	})
	if err != nil {
		log.Fatal(err)
	}

	uploader := s3manager.NewUploader(sess)

	fileName := "test.txt"
	file, err := os.OpenFile(fileName, os.O_RDWR, 0755)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("alekssum-eshop-test"),
		Key:    aws.String(fileName),
		Body:   file,
	})

	if err != nil {
		log.Fatal(err)

	}

	log.Print("success")

}
