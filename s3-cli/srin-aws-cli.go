package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

const (
	create_bucket string = "create-bucket"
	list_buckets  string = "list-buckets"
	delete_bucket string = "delete-bucket"
	upload_object string = "upload-object"
	list_objects  string = "list-objects"
	delete_object string = "delete-object"
)

func main() {
	var arguments = os.Args[1:]

	help := flag.Bool("help", false, "Show help")

	flag.Parse()

	if *help {
		fmt.Println("srin-aws-cli")
		fmt.Println("")
		fmt.Println("Usage:")
		fmt.Println("srin-aws-cli <command> <options>")
		fmt.Println("")
		fmt.Println("Commands:")
		fmt.Println("  create-bucket <bucket-name>")
		fmt.Println("  list-buckets")
		fmt.Println("  delete-bucket <bucket-name>")
		fmt.Println("  upload-object <bucket-name>  --key=<bucket-file-path> --file=<current-file-path>")
		fmt.Println("  list-objects <bucket-name>")
		fmt.Println("  delete-object <bucket-name> --key=<bucket-file-path>")

		os.Exit(0)
	}

	if len(arguments) == 0 {
		fmt.Println("Usage: srin-aws-cli <command>")

		os.Exit(0)
	}

	// Initialize a session in ap-southeast-1 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("ap-southeast-1"),
	})

	// Create S3 service client
	svc := s3.New(sess)

	switch arguments[0] {
	case create_bucket:
		if len(arguments) != 2 {
			fmt.Println("Usage: srin-aws-cli create-bucket <bucket-name>")
			os.Exit(0)
		}

		bucketName := arguments[1]

		createBucket(svc, &bucketName)
		break
	case delete_bucket:
		if len(arguments) != 2 {
			fmt.Println("Usage: srin-aws-cli delete-bucket <bucket-name>")
			os.Exit(0)
		}

		bucketName := arguments[1]

		deleteBucket(svc, &bucketName)
		break
	case list_buckets:
		listBuckets(svc)
		break
	case list_objects:
		if len(arguments) != 2 {
			fmt.Println("Usage: srin-aws-cli list-objects <bucket-name>")
			os.Exit(0)
		}

		bucketName := arguments[1]

		listObjects(svc, &bucketName)
		break
	case upload_object:
		if len(arguments) != 2 {
			fmt.Println("Usage: srin-aws-cli upload-object <bucket-name> --key=<bucket-file-path> --file=<current-file-path>")
			os.Exit(0)
		}

		bucketName := arguments[1]

		fileName := flag.String("key", "", "The key of the object to upload")
		filePath := flag.String("file", "", "The path of the file to upload")

		flag.Parse()

		if *fileName == "" || *filePath == "" {
			flag.Usage()
			os.Exit(0)
		}

		uploadObject(sess, filePath, &bucketName, fileName)
		break
	case delete_object:
		if len(arguments) != 2 {
			fmt.Println("Usage: srin-aws-cli delete-object <bucket-name> --key=<bucket-file-path>")
			os.Exit(0)
		}

		bucketName := arguments[1]

		fileName := flag.String("key", "", "The key of the object to upload")

		flag.Parse()

		if *fileName == "" {
			flag.Usage()
			os.Exit(0)
		}

		deleteObject(svc, &bucketName, fileName)
		break
	default:
		flag.Usage()
		os.Exit(0)

	}
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

func listBuckets(svc *s3.S3) {

	result, err := svc.ListBuckets(nil)

	if err != nil {
		exitErrorf("Unable to list buckets, %v", err)
	}

	fmt.Println("Buckets:")

	for _, b := range result.Buckets {
		fmt.Printf("* %s created on %s\n",
			aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	}

}

func createBucket(svc *s3.S3, bucket *string) error {

	response, err := svc.CreateBucket(&s3.CreateBucketInput{
		Bucket: bucket,
	})

	if err != nil {
		exitErrorf("Unable to create bucket, %v \n", err)
	}

	fmt.Printf("Bucket created, %v \n", *response.Location)

	err = svc.WaitUntilBucketExists(&s3.HeadBucketInput{
		Bucket: bucket,
	})

	if err != nil {
		exitErrorf("Error in create bucket, %v \n", err)
	}

	return nil
}

func deleteBucket(svc *s3.S3, bucket *string) error {

	_, err := svc.DeleteBucket(&s3.DeleteBucketInput{
		Bucket: bucket,
	})

	if err != nil {
		exitErrorf("Unable to delete bucket, %v \n", err)
	}

	err = svc.WaitUntilBucketNotExists(&s3.HeadBucketInput{
		Bucket: bucket,
	})

	if err != nil {
		exitErrorf("Error in delete bucket, %v \n", err)
	}

	fmt.Printf("Bucket deleted, %v \n\n", *bucket)

	return nil
}

func listObjects(svc *s3.S3, bucket *string) {

	result, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: bucket,
	})

	if err != nil {
		exitErrorf("Unable to list objects, %v", err)
	}

	fmt.Println("Objects in " + *bucket + ":")

	for _, item := range result.Contents {
		fmt.Println("Name:          ", *item.Key)
		fmt.Println("Last modified: ", *item.LastModified)
		fmt.Println("Size:          ", *item.Size)
		fmt.Println("Storage class: ", *item.StorageClass)
		fmt.Println("")
	}

}

func uploadObject(sess *session.Session, filePath *string, bucketName *string, fileName *string) error {
	uploader := s3manager.NewUploader(sess)

	file, err := os.Open(*filePath)

	if err != nil {
		exitErrorf("error in opening file, %v", err)
	}

	defer file.Close()

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: bucketName,
		Key:    fileName,
		Body:   file,
	})

	if err != nil {
		exitErrorf("Upload file error, %v", err)
	}

	fmt.Println("Successfully uploaded file!")

	return err
}

func deleteObject(svc *s3.S3, bucket *string, item *string) error {
	_, err := svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: bucket,
		Key:    item,
	})
	if err != nil {
		exitErrorf("Unable to delete file, %v", err)
	}

	err = svc.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: bucket,
		Key:    item,
	})

	if err != nil {
		exitErrorf("Delete file error, %v", err)
	}

	fmt.Println("Successfully delete file!")

	return nil
}
