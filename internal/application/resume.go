package model

import (
	resumePb "awesomeProject/internal/proto/resume"
	"bytes"
	"context"
	"errors"
	"io"
	"log"
)

const maxResumeSize = 2 << 20

func (app *Application) UploadResume(server resumePb.ResumeService_UploadResumeServer) error {
	data := bytes.Buffer{}
	dataSize := 0
	userId := 0

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
		userId = int(req.GetId())
		chunk := req.Resume.GetData()
		size := len(chunk)
		dataSize += size

		app.logger.Printf("received a chunk with size: %d", size)
		if dataSize > maxResumeSize {
			return errors.New("file size limit exceeded")
		}
		_, err = data.Write(chunk)
		if err != nil {
			return err
		}
	}
	// Save the file to the S3 bucket
	resume := resumePb.Resume{
		Data: data.Bytes(),
	}
	err := app.persistence.Resume.Insert(&resume, userId) // This id will be passed by interceptor
	if err != nil {
		return err
	}
	return server.SendAndClose(&resumePb.ResumeUploadResponse{
		Status: resumePb.STATUS_STATUS_APPROVED,
	})
}

func (app *Application) UploadResumeUrl(_ context.Context, request *resumePb.UploadResumeUrlRequest) (*resumePb.UploadResumeUrlResponse, error) {
	err := app.persistence.Resume.InsertUrl(request.Url, int(request.Id))
	if err != nil {
		return nil, err
	}
	return &resumePb.UploadResumeUrlResponse{
		Success: true,
	}, nil
}

func (app *Application) GetResume(_ context.Context, req *resumePb.GetResumeRequest) (*resumePb.GetResumeResponse, error) {
	resume, err := app.persistence.Resume.Get(req.GetId())
	if err != nil {
		return nil, err
	}
	return &resumePb.GetResumeResponse{
		Url: resume,
	}, nil
}

func (app *Application) DeleteResume(_ context.Context, req *resumePb.DeleteResumeRequest) (*resumePb.DeleteResumeResponse, error) {
	err := app.persistence.Resume.Delete(req.GetId())
	if err != nil {
		return &resumePb.DeleteResumeResponse{
			Status: resumePb.STATUS_STATUS_REJECTED,
		}, err
	}
	return &resumePb.DeleteResumeResponse{
		Status: resumePb.STATUS_STATUS_APPROVED,
	}, nil
}
