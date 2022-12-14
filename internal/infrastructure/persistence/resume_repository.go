package persistence

import (
	resumePb "awesomeProject/internal/proto/resume"
	"awesomeProject/utils"
	"bytes"
	"database/sql"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"log"
	"strconv"
)

type ResumeRepository struct {
	DB *sql.DB
	S3 *S3
}

// Insert Upload the resume to s3
// Get the url
// Insert the resume url into db
func (r ResumeRepository) Insert(resume *resumePb.Resume, id int) error {
	file := resume.Data
	f := bytes.NewReader(file)
	uploader := utils.S3Uploader(r.S3.ApiId, r.S3.ApiToken, r.S3.Region)
	key := fmt.Sprintf("%s%d.pdf", r.S3.Key, id)
	res, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(r.S3.Bucket),
		Key:    aws.String(key),
		Body:   f,
		ACL:    aws.String("public-read"),
	})
	log.Println("[SENSITIVE] File Uploaded Successfully, URL : ", res.Location)
	if err != nil {
		return err
	}
	// Insert the resume into db, overwrite if already exists
	query := `
			INSERT INTO resume (student_id, file_key)
			VALUES ($1, $2)
			ON CONFLICT (student_id)
			DO UPDATE SET file_key = $2`
	_, err = r.DB.Exec(query, id, res.Location)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// Get the resume url from db
// here id is student id
func (r ResumeRepository) Get(id int64) (string, error) {
	var resumeUrl string
	query := `SELECT file_key FROM resume WHERE student_id=$1`
	err := r.DB.QueryRow(query, id).Scan(&resumeUrl)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return resumeUrl, nil
}

func (r ResumeRepository) InsertUrl(url string, id int) error {
	query := `
			INSERT INTO resume (student_id, file_key) 
			VALUES ($1, $2) 
			ON CONFLICT (student_id) 
    		DO UPDATE SET file_key = $2`
	_, err := r.DB.Exec(query, id, url)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil

}

func (r ResumeRepository) Delete(id int64) error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(r.S3.Region),
	})
	if err != nil {
		return err
	}
	// Delete the resume from s3
	svc := s3.New(sess)
	key := r.S3.Key + strconv.FormatInt(id, 10) + ".pdf"
	res, err := svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(r.S3.Bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return err
	}
	log.Println("[SENSITIVE] Resume Deleted Successfully, URL : ", res)
	// Delete the resume from db
	query := `DELETE FROM resume WHERE student_id=$1`
	_, err = r.DB.Exec(query, id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
