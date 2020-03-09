package uploader

import (
	"bufio"
	context "context"
	fmt "fmt"
	"io"
	"net"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

func TestUploaderHandler_UploadFile(t *testing.T) {
	ctx := context.TODO()
	mockUploaderService, mockFileOperation, uploaderHandler := setupHandler()
	srv, listener := startGRPCServer(uploaderHandler, t)
	defer srv.Stop()

	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(getBufDialer(listener)), grpc.WithInsecure())
	assert.NoError(t, err)
	defer conn.Close()

	client := NewUploadServiceClient(conn)
	gc := NewGRPCClient(client)
	t.Run("stream file data", func(t *testing.T) {
		mockFileOperation.On("Reader", mock.Anything).Return([][]string{{"mockrecord"}}, nil)
		mockUploaderService.On("SaveBulkdata", mock.Anything).Return(nil)
		res, err := gc.UploadFile(ctx, "../test/test.csv", t)
		assert.NoError(t, err)
		assert.Equal(t, uint32(200), res.GetStatus())
	})
}

type GRPCClient struct {
	Client    UploadServiceClient
	ChunkSize int
}

func NewGRPCClient(cc UploadServiceClient) *GRPCClient {
	return &GRPCClient{cc, 1024}

}

func (gc *GRPCClient) UploadFile(ctx context.Context, f string, t *testing.T) (*UploadFileResponse, error) {
	file, err := os.Open(f)
	assert.NoError(t, err)
	defer file.Close()

	stream, err := gc.Client.UploadFile(ctx)
	assert.NoError(t, err)

	err = stream.Send(&UploadFileRequest{
		Data: &UploadFileRequest_Info{&FileInfo{Modulename: "employee"}},
	})
	assert.NoError(t, err)

	reader := bufio.NewReader(file)
	buf := make([]byte, gc.ChunkSize)
	for {
		n, err := reader.Read(buf)
		if err == io.EOF {
			fmt.Println("no more data send")
			break
		}
		assert.NoError(t, err)

		err = stream.Send(&UploadFileRequest{
			Data: &UploadFileRequest_Chuckdata{
				Chuckdata: buf[:n],
			},
		})
		assert.NoError(t, err)
	}

	res, err := stream.CloseAndRecv()
	assert.NoError(t, err)

	return res, nil
}

func setupHandler() (*MockUploaderService, *MockXLSXFileService, *UploaderHandler) {
	mockUploaderService := new(MockUploaderService)
	mockFileOperation := new(MockXLSXFileService)
	return mockUploaderService, mockFileOperation, NewUploaderHandler(mockUploaderService, mockFileOperation)
}

func getBufDialer(listener *bufconn.Listener) func(context.Context, string) (net.Conn, error) {
	return func(ctx context.Context, url string) (net.Conn, error) {
		return listener.Dial()
	}
}

func startGRPCServer(uploaderHandler *UploaderHandler, t *testing.T) (*grpc.Server, *bufconn.Listener) {
	bufferSize := 1024 * 1024
	listener := bufconn.Listen(bufferSize)

	srv := grpc.NewServer()
	RegisterUploadServiceServer(srv, uploaderHandler)
	go func() {
		err := srv.Serve(listener)
		assert.NoError(t, err)

	}()
	return srv, listener
}
