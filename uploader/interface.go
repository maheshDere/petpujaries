package uploader

import "context"

type UploaderService interface {
	SaveBulkdata(ctx context.Context, data [][]string) error
}
