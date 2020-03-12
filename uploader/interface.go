package uploader

import "context"

type UploaderService interface {
	SaveBulkdata(ctx context.Context, module string, data [][]string) error
}
