package common

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"

	"github.com/joho/godotenv"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"golang.org/x/net/context"
)

var minioClient *minio.Client
var policy = `{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Action": [
                "s3:GetObject",
                "s3:PutObject"
            ],
            "Effect": "Allow",
            "Principal": "*",
            "Resource": [
                "arn:aws:s3:::%s/*"
            ]
        }
    ]
}`

func init() {
	errEnv := godotenv.Load(".env")
	if errEnv != nil {
		log.Fatalf("Error loading .env file")
	}

	var (
		minioEndpoint  = os.Getenv("MINIO_END_POINT")
		minioAccessKey = os.Getenv("MINIO_ACCESS_KEY")
		minioSecretKey = os.Getenv("MINIO_SECRET_KEY")
		minioBucket    = os.Getenv("MINIO_BUCKET")
	)

	if v, ok := os.LookupEnv("MINIO_END_POINT"); ok {
		fmt.Printf("Database name: %s\n", v)
	} else {
		fmt.Println("Key does not exists")
	}

	// Initialize Client
	ctx := context.Background()
	client, err := minio.New(minioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(minioAccessKey, minioSecretKey, ""),
		Secure: false,
	})
	if err != nil {
		panic(err)
	}

	minioClient = client

	// Create the MinIO bucket if it doesn't exist
	err = minioClient.MakeBucket(ctx, minioBucket, minio.MakeBucketOptions{})
	if err != nil {
		exists, errBucketExists := minioClient.BucketExists(ctx, minioBucket)
		if errBucketExists == nil && exists {
			// handle set policy for bucket
			policyBucket := fmt.Sprintf(policy, minioBucket)
			minioClient.SetBucketPolicy(ctx, minioBucket, policyBucket)
			log.Printf("We already own %s\n", minioBucket)
		} else {
			log.Fatalln(err)
		}
	}
}

func CreateBucket(ctx context.Context, minioBucket string) (bool, error) {
	exists, errBucketExists := minioClient.BucketExists(ctx, minioBucket)
	if errBucketExists == nil && exists {
		return true, nil
	}

	err := minioClient.MakeBucket(ctx, minioBucket, minio.MakeBucketOptions{ObjectLocking: false})

	if err != nil {
		return false, err
	}

	// handle set policy for bucket
	policyBucket := fmt.Sprintf(policy, minioBucket)
	errSetPolicy := minioClient.SetBucketPolicy(ctx, minioBucket, policyBucket)
	if errSetPolicy != nil {
		return false, err
	}

	return true, nil
}

// func UploadImage(file *multipart.FileHeader) (string, error) {
// 	ctx := context.Background()

// 	objectName := file.Filename

// 	_, err := minioClient.FPutObject(ctx, minioBucket, objectName, file.Filename, minio.PutObjectOptions{ContentType: "image/png"})
// 	if err != nil {
// 		return "", err
// 	}

// 	return objectName, nil
// }

func UploadImage(file *multipart.FileHeader, minioBucket string) (string, error) {
	ctx := context.Background()

	// handle check and create bucket
	isCreated, errCreate := CreateBucket(ctx, minioBucket)

	if errCreate != nil && !isCreated {
		return "", errCreate
	}

	objectName := GenerateRandomString(10) + "_" + file.Filename

	// Open the uploaded file
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Get the file size
	fileSize := file.Size

	// Upload the image content to MinIO
	_, err = minioClient.PutObject(ctx, minioBucket, objectName, src, fileSize, minio.PutObjectOptions{ContentType: "image/png"})
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%s", minioBucket, objectName), nil
}

func GetImage(objectName string) ([]byte, error) {
	var (
		minioBucket = os.Getenv("MINIO_BUCKET")
	)
	data, errGetImage := minioClient.GetObject(context.Background(), minioBucket, objectName, minio.GetObjectOptions{})

	if errGetImage != nil {
		return nil, errGetImage
	}

	imageBytes, errRead := io.ReadAll(data)

	if errRead != nil {
		return nil, errRead
	}
	return imageBytes, nil
}

func RemoveImage(imagePath, bucket string) bool {
	err := minioClient.RemoveObject(context.Background(), bucket, imagePath, minio.RemoveObjectOptions{
		ForceDelete:      true,
		GovernanceBypass: false,
		VersionID:        "",
	})

	return err == nil
}

func UploadMultipleImage(files []*multipart.FileHeader, minioBucket string) (*[]string, error) {
	var paths []string
	ctx := context.Background()

	// handle check and create bucket
	isCreated, errCreate := CreateBucket(ctx, minioBucket)

	if errCreate != nil && !isCreated {
		return nil, errCreate
	}

	for _, file := range files {
		path, err := UploadImage(file, minioBucket)
		if err != nil {
			return nil, err
		}
		paths = append(paths, path)
	}

	return &paths, nil
}
