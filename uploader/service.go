package uploader

import "petpujaris/workers"

type uploaderService struct{}

func (rs uploaderService) SaveBulkdata(data [][]string) error {
	p := workers.NewPool(10, len(data), data)
	p.Run()
	return nil
}

func NewUploaderService() UploaderService {
	return uploaderService{}
}
