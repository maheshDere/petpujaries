package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"petpujaris/uploader"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to backend %v", err)
	}

	client := uploader.NewUploadServiceClient(conn)
	gc := NewGRPCClient(client)
	err = gc.UploadFile(context.Background(), "MealUpload.csv") //example file: AusVSIndMatch.csv //Restaurant Meal Upload.xlsx
	if err != nil {
		fmt.Println(err)
	}

}

type GRPCClient struct {
	Client    uploader.UploadServiceClient
	ChunkSize int
}

func NewGRPCClient(cc uploader.UploadServiceClient) *GRPCClient {
	return &GRPCClient{cc, 1024}

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
		Data: &uploader.UploadFileRequest_Info{&uploader.FileInfo{Modulename: "meal"}},
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
