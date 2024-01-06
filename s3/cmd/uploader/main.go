package main

import (
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	s3Bucket string
	s3Client *s3.S3
	wg       sync.WaitGroup
)

func main() {
	dir, err := os.Open("./tmp")
	if err != nil {
		panic(err)
	}

	defer dir.Close()

	for {
		uploadControl := make(chan struct{}, 5)
		errorControl := make(chan string)

		go func() {
			for {
				select {
				case filename := <-errorControl:
					fmt.Println("Retrying to upload file", filename)
					uploadControl <- struct{}{}
					wg.Add(1)
					go uploadFile(filename, uploadControl, errorControl)
				}
			}
		}()
		files, err := dir.ReadDir(1)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Failed to read directory", err)
			continue
		}
		wg.Add(1)
		uploadControl <- struct{}{}
		fmt.Println("Upload control", len(uploadControl))
		go uploadFile(files[0].Name(), uploadControl, errorControl)
	}
	wg.Wait()

}

func uploadFile(filename string, uploadControl <-chan struct{}, errorControl chan string) {
	defer wg.Done()
	completeFileName := fmt.Sprintf("./tmp/%s", filename)
	fmt.Println("Uploading file", completeFileName)
	f, err := os.Open(completeFileName)
	if err != nil {
		fmt.Println("Failed to open file", err)
		<-uploadControl
		errorControl <- filename
		return
	}
	defer f.Close()

	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(filename),
		Body:   f,
	})

	if err != nil {
		fmt.Println("Failed to upload file", err)
		<-uploadControl
		errorControl <- filename
		return
	}

	fmt.Println("Successfully uploaded file to", s3Bucket)
	<-uploadControl
}

func init() {
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String("us-east-1"),
			Credentials: credentials.NewStaticCredentials(
				"eita",
				"credo",
				"",
			),
		},
	)

	if err != nil {
		panic(err)
	}

	s3Client = s3.New(sess)
	s3Bucket = "goexpert-felipe-bucket-example"
}
