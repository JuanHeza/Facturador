package aws

// https://aws.github.io/aws-sdk-go-v2/docs/getting-started/

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

//https://github.dev/awsdocs/aws-doc-sdk-examples/blob/main/gov2/s3/actions/bucket_basics.go#L84

type BucketBasics struct {
	S3Client *s3.Client
}

func newBucketBasics()(BucketBasics){
    // Load the Shared AWS Configuration (~/.aws/config)
    cfg, err := config.LoadDefaultConfig(context.TODO())
    if err != nil {
        log.Fatal(err)
    }

    // Create an Amazon S3 service client
    return BucketBasics{
        S3Client: s3.NewFromConfig(cfg),
    }
}

func pages() {
    // Load the Shared AWS Configuration (~/.aws/config)
    cfg, err := config.LoadDefaultConfig(context.TODO())
    if err != nil {
        log.Fatal(err)
    }

    // Create an Amazon S3 service client
    client := s3.NewFromConfig(cfg)

    // Get the first page of results for ListObjectsV2 for a bucket
    output, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
        Bucket: aws.String("my-bucket"),
    })
    if err != nil {
        log.Fatal(err)
    }

    log.Println("first page results:")
    for _, object := range output.Contents {
        log.Printf("key=%s size=%d", aws.ToString(object.Key), object.Size)
    }
}

func (basics BucketBasics) UploadFile(bucketName string, objectKey string, fileName string) error {
    file, err := os.Open(fileName)
    if err != nil {
        log.Printf("Couldn't open file %v to upload. Here's why: %v\n", fileName, err)
    } else {
        defer file.Close()
        _, err = basics.S3Client.PutObject(context.TODO(), &s3.PutObjectInput{
            Bucket: aws.String(bucketName),
            Key:    aws.String(objectKey),
            Body:   file,
        })
        if err != nil {
            log.Printf("Couldn't upload file %v to %v:%v. Here's why: %v\n",
                fileName, bucketName, objectKey, err)
        }
    }
    return err
}

func (basics BucketBasics) DownloadFile(bucketName string, objectKey string, fileName string) error {
    result, err := basics.S3Client.GetObject(context.TODO(), &s3.GetObjectInput{
        Bucket: aws.String(bucketName),
        Key:    aws.String(objectKey),
    })
    if err != nil {
        log.Printf("Couldn't get object %v:%v. Here's why: %v\n", bucketName, objectKey, err)
        return err
    }
    defer result.Body.Close()
    file, err := os.Create(fileName)
    if err != nil {
        log.Printf("Couldn't create file %v. Here's why: %v\n", fileName, err)
        return err
    }
    defer file.Close()
    body, err := io.ReadAll(result.Body)
    if err != nil {
        log.Printf("Couldn't read object body from %v. Here's why: %v\n", objectKey, err)
    }
    _, err = file.Write(body)
    return err
}

func (basics BucketBasics) ListObjects(bucketName string) ([]types.Object, error) {
    result, err := basics.S3Client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
        Bucket: aws.String(bucketName),
    })
    var contents []types.Object
    if err != nil {
        log.Printf("Couldn't list objects in bucket %v. Here's why: %v\n", bucketName, err)
    } else {
        contents = result.Contents
    }
    return contents, err
}


func (basics BucketBasics) DeleteObjects(bucketName string, objectKeys []string) error {
    var objectIds []types.ObjectIdentifier
    for _, key := range objectKeys {
        objectIds = append(objectIds, types.ObjectIdentifier{Key: aws.String(key)})
    }
    _, err := basics.S3Client.DeleteObjects(context.TODO(), &s3.DeleteObjectsInput{
        Bucket: aws.String(bucketName),
        Delete: &types.Delete{Objects: objectIds},
    })
    if err != nil {
        log.Printf("Couldn't delete objects from bucket %v. Here's why: %v\n", bucketName, err)
    }
    return err
}