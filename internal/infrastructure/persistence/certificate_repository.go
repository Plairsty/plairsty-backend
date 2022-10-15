package persistence

import (
	certificatePb "awesomeProject/internal/proto/certificates"
	"awesomeProject/utils"
	"bytes"
	"database/sql"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"log"
)

type CertificateRepository struct {
	DB *sql.DB
	S3 *S3
}

func (r CertificateRepository) Insert(
	userId int64,
	certificateData *certificatePb.CertificateData,
	certificate *certificatePb.CertificateFields,
) (int64, error) {
	file := certificateData.GetData()
	f := bytes.NewReader(file)
	uploader := utils.S3Uploader(r.S3.ApiId, r.S3.ApiToken, r.S3.Region)
	key := fmt.Sprintf("%s%d.pdf", r.S3.Key, userId)
	res, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(r.S3.Bucket),
		Key:    aws.String(key),
		Body:   f,
		ACL:    aws.String("public-read"),
	})
	log.Println("[SENSITIVE] File Uploaded Successfully, URL : ", res.Location)
	if err != nil {
		return 0, err
	}

	query := `
		INSERT INTO certifications (user_id, title, description, url, cert_id, created_at, updated_at, allotted_date, expiry, type, level, category, issuer, domain)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
		RETURNING id`
	args := []interface{}{
		userId,
		certificate.GetTitle(),
		certificate.GetDescription(),
		res.Location,
		certificate.GetCertificateId(),
		certificate.GetAllocatedDate(),
		certificate.GetExpiryDate(),
		certificate.GetCertificateType(),
		certificate.GetCertificateLevel(),
		certificate.GetCertificateCategory(),
		certificate.GetCertificateIssuer(),
		certificate.GetDomain(),
	}
	var id int64
	err = r.DB.QueryRow(query, args...).Scan(&id)
	if err != nil {
		log.Println("[ERROR] Error while inserting certificate data", err)
		return 0, err
	}
	return id, nil
}

func (r CertificateRepository) Get(id int64) (*certificatePb.CertificateFields, error) {
	// Get the certificate from the database by certificate id
	query := `
		SELECT id, title, description, url, cert_id, allotted_date, expiry, type, level, category, issuer, domain
		FROM certifications
		WHERE cert_id = $1`
	var certificate certificatePb.CertificateFields
	err := r.DB.QueryRow(query, id).Scan(
		&certificate.Id,
		&certificate.Title,
		&certificate.Description,
		&certificate.CertificateFileUrl,
		&certificate.CertificateId,
		&certificate.AllocatedDate,
		&certificate.ExpiryDate,
		&certificate.CertificateType,
		&certificate.CertificateLevel,
		&certificate.CertificateCategory,
		&certificate.CertificateIssuer,
		&certificate.Domain,
	)
	if err != nil {
		log.Println("[ERROR] Error while getting certificate data", err)
		return nil, err
	}
	return &certificate, nil
}

func (r CertificateRepository) GetAll(id int64) ([]*certificatePb.CertificateFields, error) {
	// Get all the certificates from the database by student id
	query := `
		SELECT id, title, description, url, cert_id, allotted_date, expiry, type, level, category, issuer, domain
		FROM certifications
		WHERE user_id = $1`
	rows, err := r.DB.Query(query, id)
	if err != nil {
		log.Println("[ERROR] Error while getting certificate data", err)
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println("[ERROR] Error while closing rows", err)
		}
	}(rows)
	var certificates []*certificatePb.CertificateFields
	for rows.Next() {
		var certificate certificatePb.CertificateFields
		err := rows.Scan(
			&certificate.Id,
			&certificate.Title,
			&certificate.Description,
			&certificate.CertificateFileUrl,
			&certificate.CertificateId,
			&certificate.AllocatedDate,
			&certificate.ExpiryDate,
			&certificate.CertificateType,
			&certificate.CertificateLevel,
			&certificate.CertificateCategory,
			&certificate.CertificateIssuer,
			&certificate.Domain,
		)
		if err != nil {
			log.Println("[ERROR] Error while getting certificate data", err)
			return nil, err
		}
		certificates = append(certificates, &certificate)
	}
	return certificates, nil
}

func (r CertificateRepository) Update(userId, certId int64, certificate *certificatePb.CertificateFields) error {
	// get the user id and certificate id and then update the certificate
	query := `
		UPDATE certifications
		SET title = $1, description = $2, allotted_date = $3, expiry = $4, type = $5, level = $6, category = $7, issuer = $8, domain = $9
		WHERE user_id = $10 AND cert_id = $11`
	_, err := r.DB.Exec(query,
		certificate.GetTitle(),
		certificate.GetDescription(),
		certificate.GetAllocatedDate(),
		certificate.GetExpiryDate(),
		certificate.GetCertificateType(),
		certificate.GetCertificateLevel(),
		certificate.GetCertificateCategory(),
		certificate.GetCertificateIssuer(),
		certificate.GetDomain(),
		userId,
		certId,
	)
	if err != nil {
		log.Println("[ERROR] Error while updating certificate data", err)
		return err
	}
	return nil
}

func (r CertificateRepository) Delete(userId, certId int64) error {
	// get the user id and certificate id and then delete the certificate
	query := `
		DELETE FROM certifications
		WHERE user_id = $1 AND cert_id = $2`
	_, err := r.DB.Exec(query, userId, certId)
	if err != nil {
		log.Println("[ERROR] Error while deleting certificate data", err)
		return err
	}
	return nil
}

func (r CertificateRepository) ChangeStatus(id int64, status certificatePb.STATUS) error {
	// Change the status of the certificate by certificate id
	panic("implement me")
}
