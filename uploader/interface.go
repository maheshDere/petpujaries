package uploader

import "context"

type UploaderService interface {
	SaveBulkdata(ctx context.Context, module string, userID int64, data [][]string) [][]string
}
