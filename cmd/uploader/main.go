package main

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
)

var uploaded {
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
	if err != nil {
		panic(err)
	}
	defer dir.Close()

	for {
		files, err := dir.ReadDir(1)
		if err != nil {
			if err= io.EOF {
				break
			}
			fmt.Printf("Error reading directory: %s \n", err)
			continue
		}
		uploadFile(files[0].Name())
	}
}

func uploadFile(filename string) {
	completeFileName := fmt.Sprint("./tmp/%s", filename)
	f, err := os.Open(completeFileName)
	if err != nil {
		fmt.Printf("Error opening file %s\n", completeFileName)
	}
	defer f.Close()
	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Buket: aws.String(s3Bucket),
		Key: aws.String(filename),
		Body: f,
	})

	if err != nil {
		fmt.Printf("Error uploading file %s\n", completeFileName)
	}

	fmt.Printf("File %s uploaded successfull\n", completeFileName)
}