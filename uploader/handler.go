package uploader

import (
	"bytes"
	fmt "fmt"
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
	fmt.Println("moduleName: ", moduleName)
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

	result, err := uh.FileService.Reader(&fileData)
	for _, v := range result {
		for k1, v1 := range v {
			fmt.Printf("k:%d and v:  %s \n", k1, v1)
		}
		fmt.Println("**************")
	}

	uh.Service.SaveBulkdata(stream.Context(), result)

	return nil
}
