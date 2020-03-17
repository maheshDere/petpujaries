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

	"github.com/urfave/cli"

	"google.golang.org/grpc"
)

type GRPCDownloaderClient struct {
	Client downloader.DownloadServiceClient
}

type GRPCClient struct {
	Client    uploader.UploadServiceClient
	ChunkSize int
}

var gc *GRPCClient
var gdc *GRPCDownloaderClient

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to backend %v", err)
	}

	client := uploader.NewUploadServiceClient(conn)
	gc = NewGRPCClient(client)
	downloaderClient := downloader.NewDownloadServiceClient(conn)
	gdc = NewGRPCDownloaderClient(downloaderClient)

	cliApp := cli.NewApp()
	cliApp.Name = "data uploadation service"
	cliApp.Version = "1.0.0"
	cliApp.Commands = []cli.Command{
		{
			Name:  "employee_data_upload",
			Usage: "upload users data from csv file to postgres database",
			Action: func(c *cli.Context) error {
				fileName := c.Args().Get(0)
				err := uploadUsersData(fileName)
				return err
			},
		},
		{
			Name:  "meal_data_upload",
			Usage: "upload restorent data from csv to postgres database",
			Action: func(c *cli.Context) error {
				fileName := c.Args().Get(0)
				err := uploadMealData(fileName)
				return err
			},
		},
		{
			Name:  "meal_scheduler_data_upload",
			Usage: "upload meal scheduler data from csv to postgres database",
			Action: func(c *cli.Context) error {
				fileName := c.Args().Get(0)
				err := uploadMealSchedulerData(fileName)
				return err
			},
		},
		{
			Name:  "download_user__upload_data_template",
			Usage: "get primary users data to generate user uploadation csv file",
			Action: func(c *cli.Context) error {
				err := DownloadUserTemplateData()
				return err
			},
		},
	}
	if err := cliApp.Run(os.Args); err != nil {
		panic(err)
	}

}

func uploadUsersData(fileName string) error {
	module := "employee"
	err := gc.UploadFile(context.Background(), fileName, module) //example file: AusVSIndMatch.csv //Restaurant Meal Upload.xlsx
	if err != nil {
		fmt.Println(err)
	}

	return err
}

func uploadMealData(fileName string) error {
	moduleName := "meal"
	err := gc.UploadFile(context.Background(), fileName, moduleName) //example file: AusVSIndMatch.csv //Restaurant Meal Upload.xlsx
	if err != nil {
		fmt.Println(err)
	}

	return err
}

func uploadMealSchedulerData(fileName string) error {
	return nil
}

func DownloadUserTemplateData() error {
	gdc.DownloadUserPrimarydata(context.Background())
	return nil
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

func (gc *GRPCClient) UploadFile(ctx context.Context, f, module string) error {
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
		Data: &uploader.UploadFileRequest_Info{&uploader.FileInfo{Modulename: module, Userid: 1}},
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

func NewGRPCDownloaderClient(cc downloader.DownloadServiceClient) *GRPCDownloaderClient {
	return &GRPCDownloaderClient{cc}
}

func NewGRPCClient(cc uploader.UploadServiceClient) *GRPCClient {
	return &GRPCClient{cc, 1024}
}
