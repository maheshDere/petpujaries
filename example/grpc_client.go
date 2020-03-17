package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"petpujaris/downloader"
	"petpujaris/uploader"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to backend %v", err)
	}

	// client := uploader.NewUploadServiceClient(conn)
	// gc := NewGRPCClient(client)
	// err = gc.UploadFile(context.Background(), "MealUpload.csv") //example file: AusVSIndMatch.csv //Restaurant Meal Upload.xlsx
	// if err != nil {
	// 	fmt.Println(err)
	// }

	downloaderClient := downloader.NewDownloadServiceClient(conn)
	gdc := NewGRPCDownloaderClient(downloaderClient)
	gdc.DownloadUserPrimarydata(context.Background())
}

type GRPCDownloaderClient struct {
	Client downloader.DownloadServiceClient
}

type GRPCClient struct {
	Client    uploader.UploadServiceClient
	ChunkSize int
}

func NewGRPCDownloaderClient(cc downloader.DownloadServiceClient) *GRPCDownloaderClient {
	return &GRPCDownloaderClient{cc}
}

func NewGRPCClient(cc uploader.UploadServiceClient) *GRPCClient {
	return &GRPCClient{cc, 1024}
}

func (gdc *GRPCDownloaderClient) DownloadUserPrimarydata(ctx context.Context) {
	var req downloader.EmployeeFileDownloadRequest
	req.AdminID = uint64(7)
	req.TotalEmployeeCount = uint64(4)
	response, err := gdc.Client.DownloadEmployeeFileData(ctx, &req)
	if err != nil {
		fmt.Println("error :", err)
	}
	for _, userData := range response.EmployeeDetails {
		for _, userinfo := range userData.EmployeeData {
			fmt.Printf("%v  ", userinfo)
		}
		fmt.Printf("\n")
	}
}

func (gc *GRPCClient) UploadFile(ctx context.Context, f string) error {

	file, err := os.Open(f)
	if err != nil {
		return fmt.Errorf("File not open %v", err)
	}
	defer file.Close()

	stream, err := gc.Client.UploadFile(ctx)
	if err != nil {
		return fmt.Errorf("Could not create client connection %v", err)
	}

	err = stream.Send(&uploader.UploadFileRequest{
		Data: &uploader.UploadFileRequest_Info{&uploader.FileInfo{Modulename: "employee"}},
	})

	if err != nil {
		return fmt.Errorf("data could not send to server %v", err)
	}

	reader := bufio.NewReader(file)
	buf := make([]byte, gc.ChunkSize)
	for {
		n, err := reader.Read(buf)
		if err == io.EOF {
			fmt.Println("no more data send")
			break
		}

		if err != nil {
			return fmt.Errorf("file could not read %v", err)
		}

		err = stream.Send(&uploader.UploadFileRequest{
			Data: &uploader.UploadFileRequest_Chuckdata{
				Chuckdata: buf[:n],
			},
		})

		if err != nil {
			return fmt.Errorf("file could not send %v", err)
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		return fmt.Errorf("could not receive the response %v", err)
	}

	fmt.Println("File uploaded successfully", res.GetStatus(), res.GetMessage(), res.GetSize())
	return nil
}
