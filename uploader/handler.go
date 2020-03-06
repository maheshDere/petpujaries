package uploader

import (
	"bytes"
	fmt "fmt"
	"io"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type UploaderHandler struct {
}

func NewUploaderHandler() *UploaderHandler {
	return &UploaderHandler{}

}

const maxFileSize = 1 << 20

func (s *UploaderHandler) UploadFile(stream UploadService_UploadFileServer) error {
	req, err := stream.Recv()
	if err != nil {
		fmt.Println("stream recv error", err)
		return status.Errorf(codes.Unknown, "can not recevice file")
	}

	moduleName := req.GetInfo().GetModulename()
	fmt.Println("moduleName: ", moduleName)
	fileData := bytes.Buffer{}
	fileSize := 0

	for {
		fmt.Println("waititng for chunk data")
		req, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("no more data send")
			break
		}
		if err != nil {
			fmt.Println("stream recv error", err)
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
			fmt.Println("Buffer error", err)
			return status.Errorf(codes.Unknown, "can not recevice file")
		}
	}

	fmt.Println("File successfully get")
	fmt.Println(fileData.String())
	err = stream.SendAndClose(&UploadFileResponse{
		Message: "Upload received with success",
		Status:  200,
		Size:    uint32(fileSize),
	})

	if err != nil {
		fmt.Println("Not send success response", err)
		return status.Errorf(codes.Unknown, "response not send")
	}
	return nil
}
