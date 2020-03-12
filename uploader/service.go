package uploader

import (
	"context"
	"petpujaris/workers"
)

type uploaderService struct {
	WorkerPool workers.Pool
}

func (rs uploaderService) SaveBulkdata(ctx context.Context, module string, data [][]string) error {
	rs.WorkerPool.Run(ctx, module, data)
	return nil
}

func NewUploaderService(w workers.Pool) UploaderService {
	return uploaderService{WorkerPool: w}
}
