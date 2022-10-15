package model

import (
	certificatePb "awesomeProject/internal/proto/certificates"
	"context"
)

func (app *Application) AddCertificate(
	certificatePb.CertificateService_AddCertificateServer,
) error {
	return nil
}

func (app *Application) GetAllCertificate(
	context.Context,
	*certificatePb.GetAllCertificateRequest,
) (*certificatePb.GetAllCertificateResponse, error) {
	return nil, nil
}

func (app *Application) GetCertificate(
	context.Context,
	*certificatePb.GetCertificateRequest,
) (*certificatePb.GetCertificateResponse, error) {
	return nil, nil
}

func (app *Application) UpdateCertificate(
	context.Context,
	*certificatePb.UpdateCertificateRequest,
) (*certificatePb.UpdateCertificateResponse, error) {
	return nil, nil
}

func (app *Application) DeleteCertificate(
	context.Context,
	*certificatePb.DeleteCertificateRequest,
) (*certificatePb.DeleteCertificateResponse, error) {
	return nil, nil
}

func (app *Application) ChangeCertificateStatus(
	context.Context,
	*certificatePb.ChangeCertificateStatusRequest,
) (*certificatePb.ChangeCertificateStatusResponse, error) {
	return nil, nil
}
