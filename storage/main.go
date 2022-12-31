package storage

type ErrorsMap struct {
	Error error
	Items []interface{}
}

type CollectionEnum string

const (
	AnimesCollection CollectionEnum = "Animes"
	GenresCollection CollectionEnum = "Genres"
)

func (dcn CollectionEnum) String() string {
	return string(dcn)
}

type Collection interface {
	Upload(any) error
	UploadManySync(map[string]interface{}) []ErrorsMap
	UploadManyAsync(map[string]interface{}) []ErrorsMap
	List() []interface{}
}

type StorageManager interface {
	Col(CollectionEnum) Collection
	Animes() Collection
	Genres() Collection
}
