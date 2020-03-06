package uploader

type UploaderService interface {
	SaveBulkdata(data [][]string) error
}
