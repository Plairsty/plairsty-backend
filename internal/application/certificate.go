package model

import (
	certificatePb "awesomeProject/internal/proto/certificates"
	"bytes"
	"context"
	"io"
	"log"
)

func (app *Application) AddCertificate(
	server certificatePb.CertificateService_AddCertificateServer,
) error {
	data := bytes.Buffer{}
	userId := int64(0)
	certFields := &certificatePb.CertificateFields{}
	for {
		log.Println("Waiting for data")
		req, err := server.Recv()
		if err == io.EOF {
			app.logger.Println("EOF")
			break
		}
		if err != nil {
			return err
		}
		userId = req.GetUserId()
		certFields = req.GetCertificate()
		chunk := req.Data.GetData()
		_, err = data.Write(chunk)
		if err != nil {
			return err
		}
	}
	cert := certificatePb.CertificateData{
		Data: data.Bytes(),
	}
	_, err := app.persistence.Certificate.Insert(userId, &cert, certFields)
	if err != nil {
		return err
	}
	return server.SendAndClose(&certificatePb.AddCertificateResponse{
		Status: certificatePb.STATUS_APPROVED,
	})
}

func (app *Application) GetAllCertificate(
	_ context.Context,
	req *certificatePb.GetAllCertificateRequest,
) (*certificatePb.GetAllCertificateResponse, error) {
	field, err := app.persistence.Certificate.GetAll(req.GetUserId())
	if err != nil {
		return nil, err
	}
	return &certificatePb.GetAllCertificateResponse{
		Certificate: field,
	}, nil
}

func (app *Application) GetCertificate(
	_ context.Context,
	req *certificatePb.GetCertificateRequest,
) (*certificatePb.GetCertificateResponse, error) {
	field, err := app.persistence.Certificate.Get(req.GetCertificateId())
	if err != nil {
		return nil, err
	}
	return &certificatePb.GetCertificateResponse{
		Certificate: field,
	}, nil
}

func (app *Application) UpdateCertificate(
	_ context.Context,
	req *certificatePb.UpdateCertificateRequest,
) (*certificatePb.UpdateCertificateResponse, error) {
	err := app.persistence.Certificate.Update(req.GetUserId(), req.GetCertificateId(), req.GetCertificate())
	if err != nil {
		return nil, err
	}
	return &certificatePb.UpdateCertificateResponse{
		Status: certificatePb.STATUS_APPROVED,
	}, nil
}

func (app *Application) DeleteCertificate(
	_ context.Context,
	req *certificatePb.DeleteCertificateRequest,
) (*certificatePb.DeleteCertificateResponse, error) {
	err := app.persistence.Certificate.Delete(req.GetUserId(), req.GetCertificateId())
	if err != nil {
		return nil, err
	}
	return &certificatePb.DeleteCertificateResponse{
		Status: certificatePb.STATUS_APPROVED,
	}, nil
}

func (app *Application) ChangeCertificateStatus(
	context.Context,
	*certificatePb.ChangeCertificateStatusRequest,
) (*certificatePb.ChangeCertificateStatusResponse, error) {
	return nil, nil
}
