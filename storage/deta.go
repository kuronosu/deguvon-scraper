package storage

import (
	"fmt"
	"os"

	"github.com/deta/deta-go/deta"
	"github.com/deta/deta-go/service/base"
	"github.com/kuronosu/deguvon-scraper/utils"
)

// Deta collection implementation

type DetaCollection struct {
	name    CollectionEnum
	db      *base.Base
	Verbose bool
}

func (col *DetaCollection) uploadMany(data map[string]interface{}, verbose bool, progressbarFunc utils.ProgressbarFunc[[]any]) []ErrorsMap {
	errors := make([]ErrorsMap, 0)

	out := os.Stdout
	if !verbose {
		out = nil
	}
	fmt.Println("Uploading " + col.name.String())

	progressbarFunc(utils.ChunkMap(data, 25), "Uploading "+col.name.String(), 50, out, func(dataChunk []any) {
		_, err := col.db.PutMany(dataChunk)
		if err != nil {
			errors = append(errors, ErrorsMap{err, dataChunk})
		}
	})
	return errors
}

func (col DetaCollection) Upload(data any) error {
	_, err := col.db.Put(data)
	return err
}

func (col DetaCollection) UploadManySync(data map[string]interface{}) []ErrorsMap {
	return col.uploadMany(data, col.Verbose, utils.Progressbar[[]any])
}

func (col DetaCollection) UploadManyAsync(data map[string]interface{}) []ErrorsMap {
	return col.uploadMany(data, col.Verbose, utils.ProgressbarAsync[[]any])
}

func (col DetaCollection) List() []interface{} {
	res := make([]interface{}, 0)
	col.db.Fetch(&base.FetchInput{
		Q:    nil,
		Dest: &res,
	})
	return res
}

func (dc *DetaCollection) Name() string {
	return dc.name.String()
}

// Deta storage manager implementation

type DetaStorageManager struct {
	deta *deta.Deta
	dbs  map[CollectionEnum]DetaCollection
}

var collections = []CollectionEnum{
	AnimesCollection,
	GenresCollection,
}

func NewDetaStorageManager(verbose bool) (StorageManager, error) {
	d, err := deta.New()
	if err != nil {
		fmt.Println("failed to init new Deta instance:", err)
		return nil, err
	}
	dsm := &DetaStorageManager{
		deta: d,
		dbs:  make(map[CollectionEnum]DetaCollection),
	}

	for _, collection := range collections {
		err := dsm.initCol(collection, verbose)
		if err != nil {
			return nil, err
		}
	}

	for _, db := range dsm.dbs {
		db.Verbose = verbose
	}

	return dsm, nil
}

func (dsm *DetaStorageManager) initCol(name CollectionEnum, verbose bool) error {
	if _, ok := dsm.dbs[name]; !ok {
		db, err := base.New(dsm.deta, name.String())
		if err != nil {
			return err
		}
		dsm.dbs[name] = DetaCollection{name, db, verbose}
	}
	return nil
}

func (dsm *DetaStorageManager) Col(name CollectionEnum) Collection {
	if _, ok := dsm.dbs[name]; !ok {
		err := dsm.initCol(name, false)
		if err != nil {
			return nil
		}
	}
	return dsm.dbs[name]
}

func (dsm *DetaStorageManager) Animes() Collection {
	return dsm.Col(AnimesCollection)
}

func (dsm *DetaStorageManager) Genres() Collection {
	return dsm.Col(GenresCollection)
}
