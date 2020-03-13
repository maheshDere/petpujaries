package uploader

import (
	"bytes"
	"io"
	"petpujaris/filemanager"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type UploaderHandler struct {
	Service     UploaderService
	FileService filemanager.FileOperation
}

func NewUploaderHandler(service UploaderService, fileService filemanager.FileOperation) *UploaderHandler {
	return &UploaderHandler{Service: service, FileService: fileService}
}

const maxFileSize = 1 << 20

func (uh *UploaderHandler) UploadFile(stream UploadService_UploadFileServer) error {
	req, err := stream.Recv()
	if err != nil {
		return status.Errorf(codes.Unknown, "can not recevice file")
	}

	moduleName := req.GetInfo().GetModulename()

	fileData := bytes.Buffer{}
	fileSize := 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return status.Errorf(codes.Unknown, "can not recevice file")
		}

		chunk := req.GetChuckdata()
		size := len(chunk)
		fileSize += size

		if fileSize >= maxFileSize {
			return status.Errorf(codes.InvalidArgument, "file size too large")
		}

		_, err = fileData.Write(chunk)
		if err != nil {
			return status.Errorf(codes.Unknown, "can not recevice file")
		}
	}

	err = stream.SendAndClose(&UploadFileResponse{
		Message: "Upload received with success",
		Status:  200,
		Size:    uint32(fileSize),
	})
	if err != nil {
		return status.Errorf(codes.Unknown, "response not send")
	}

	data, err := uh.FileService.Reader(&fileData)
	if err != nil {
		return status.Errorf(codes.Unknown, "error in read file")
	}

	uh.Service.SaveBulkdata(stream.Context(), moduleName, 0, data)
	return nil
}
