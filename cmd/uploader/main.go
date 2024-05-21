package main

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
	"sync"
)

var {
	s3Client *s3.S3
	s3Bucket string
}

func init() {
	sess, err := session.NewSession{
		Region: aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials{
			"ASDFASDFASDFFKFKIEMMAKNTIME",
			"asdfasdfasdfdaf/asdkfjçksjfefadfkey",
			"",
		}
	}

	if err != nil {
		panic(err)
	}

	S3.s3Client = s3.New(sess)
	S3.s3Bucket = "goexpert-bucket-exemplo"
}

func main() {
	dir, err := os.Open("./tmp")
	var wg sync.WaitGroup
	if err != nil {
		panic(err)
	}
	defer dir.Close()
	uploadControl := make(chan struct{}, 200)
	errorFileUpload := make(chan string, 10)

	go func() {
		for {
			select {
			case filename <- errorFileUpload:
				uploadControl <- struct{}{}
				wg.Add(1)
				go uploadFile(filename, uploadControl, errorFileUpload)
			}
		}
	}()

	for {
		files, err := dir.ReadDir(1)
		if err != nil {
			if err= io.EOF {
				break
			}
			fmt.Printf("Error reading directory: %s \n", err)
			continue
		}
		wg.Add(1)
		uploadControl <- struct{}{}
		go uploadFile(files[0].Name(), uploadControl, errorFileUpload)
	}
	wg.Wait()
}

func uploadFile(filename string, uploadControl <-chan struct{}, errorFileUpload chan<- string) {
	defer wg.Done();
	completeFileName := fmt.Sprint("./tmp/%s", filename)
	f, err := os.Open(completeFileName)
	if err != nil {
		fmt.Printf("Error opening file %s\n", completeFileName)
		<-uploadControl
		errorFileUpload <- completeFileName
	}
	defer f.Close()
	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Buket: aws.String(s3Bucket),
		Key: aws.String(filename),
		Body: f,
	})

	if err != nil {
		fmt.Printf("Error uploading file %s\n", completeFileName)
		<-uploadControl
		errorFileUpload <- completeFileName
	}

	fmt.Printf("File %s uploaded successfull\n", completeFileName)
	<-uploadControl
}