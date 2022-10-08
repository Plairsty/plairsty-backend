package model

import (
	resumePb "awesomeProject/internal/proto/resume"
	"bytes"
	"errors"
	"io"
	"log"
)

const maxImageSize = 2 << 20

func (app *Application) UploadResume(server resumePb.ResumeService_UploadResumeServer) error {
	data := bytes.Buffer{}
	dataSize := 0
	for {
		log.Println("Waiting for data")
		req, err := server.Recv()
		if err == io.EOF {
			log.Println("EOF")
			break
		}
		if err != nil {
			return err
		}

		chunk := req.GetResume().GetData()
		size := len(chunk)
		log.Printf("received a chunk with size: %d", size)
		dataSize += size
		if dataSize > maxImageSize {
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
	err := app.persistence.Resume.Insert(&resume, 355284088) // This id will be passed by interceptor
	if err != nil {
		return err
	}
	return server.SendAndClose(&resumePb.ResumeUploadResponse{
		Status: resumePb.STATUS_STATUS_APPROVED,
	})
}
