package configs

import (
	"fmt"
	"net/http"
	"time"
	"web-service-gin/utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateSession() *session.Session {
	utils.LoadConfig()
	sess := session.Must(session.NewSession(
		&aws.Config{
			Region: aws.String(utils.Config.AwsRegion),
			Credentials: credentials.NewStaticCredentials(
				utils.Config.AwsAccessKey,
				utils.Config.AwsSecretKey,
				"",
			),
		},
	))
	return sess
}

func CreateS3Session(sess *session.Session) *s3.S3 {
	s3session := s3.New(sess)
	return s3session
}

func UploadObject() gin.HandlerFunc {
	return func(c *gin.Context) {
		utils.LoadConfig()
		fileName := uuid.New().String() + ".jpg"
		fmt.Println(fileName)
		sess := CreateSession()
		s3session := s3.New(sess)
		r, _ := s3session.PutObjectRequest(
			&s3.PutObjectInput{
				Bucket: &utils.Config.AWSS3Bucket,
				Key:    &fileName,
				//ACL:    aws.String("public-read-write"),
			})
		// Upload to s3

		//r.HTTPRequest.Header.Add("X-Amz-Acl", "public-read")
		// q := r.HTTPRequest.URL.Query()
		// q.Add("acl", "public-read")
		// q.Add("Content-Type", c.ContentType())
		presigned, err := r.Presign(5 * time.Minute)
		if err != nil {
			fmt.Println(err)
		}
		utils.SendResponse(c, http.StatusOK, "Success", gin.H{"presigned_url": presigned, "key": fileName})
	}
}

func GetFileLink() gin.HandlerFunc {
	return func(c *gin.Context) {
		utils.LoadConfig()
		key := c.Query("key")
		url := "https://" + utils.Config.AWSS3Bucket + ".s3." + utils.Config.AwsRegion + ".amazonaws.com/" + key
		utils.SendResponse(c, http.StatusOK, "Success", url)
	}
}
