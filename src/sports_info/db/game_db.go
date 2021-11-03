package db

import (
	"context"
	"fmt"
	"sports_info/game/crypto_aux"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB_base struct {
	db_name_base    string
	db_name         string
	db_addr         string
	db_err          string
	db_user         string
	db_pwd          string
	b_auth          bool
	b_init          bool
	b_close         bool
	b_useReplicaSet bool
	db_uri          string        //("mongodb://localhost:27017");
	db_client       *mongo.Client //db_client(db_uri)
	db_database     *mongo.Database
	db_ctx_base     context.Context
	f_cancel        context.CancelFunc
}

func NewDB(user string, pwd string, crypto bool, _db_addr string, _db_name string) (*DB_base, string) {
	db := &DB_base{}
	if crypto {
		key := "94@!#(*13&32!@)("
		{
			dec_64 := crypto_aux.BaseDeEncode(string(user[:]))
			db.db_user = string(crypto_aux.AesDecryptECB([]byte(dec_64), []byte(key)))
		}

		{
			dec_64 := crypto_aux.BaseDeEncode(string(pwd[:]))
			db.db_pwd = string(crypto_aux.AesDecryptECB([]byte(dec_64), []byte(key)))
		}

	} else {
		db.db_user = user
		db.db_pwd = pwd
	}
	if !db.init_db(_db_addr, _db_name) {
		return nil, db.db_err
	}
	return db, ""
}

func (db *DB_base) init_db(_db_addr string, _db_name string) bool {

	if db.b_init {
		db.db_err = "error, init aleary"
		return false
	}

	db.db_addr = _db_addr
	db.db_name_base = _db_name
	db.db_name = _db_name
	db.db_err = ""
	db.b_init = true

	base := db.get_conn()
	return base != nil
}

func (db *DB_base) Insert(table_name string, doc interface{}) string {
	base := db.get_conn()
	if base == nil {
		return db.db_err
	}
	ctx, cancel := context.WithCancel(db.db_ctx_base)
	defer cancel()

	coll := base.Collection(table_name)
	insertResult, err := coll.InsertOne(ctx, doc)
	if err != nil {
		db.db_err = err.Error()
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	return db.db_err
}

func (db *DB_base) Insertmany(table_name string, docs []interface{}) string {
	base := db.get_conn()
	if base == nil {
		return db.db_err
	}
	ctx, cancel := context.WithCancel(db.db_ctx_base)
	defer cancel()
	coll := base.Collection(table_name)
	_, err := coll.InsertMany(ctx, docs)
	if err != nil {
		db.db_err = err.Error()
	}
	return db.db_err
}
func (db *DB_base) Findone(table_name string, filter interface{}, opts *options.FindOneOptions) *mongo.SingleResult {
	base := db.get_conn()
	if base == nil {
		return nil
	}
	ctx, cancel := context.WithCancel(db.db_ctx_base)
	defer cancel()
	coll := base.Collection(table_name)
	return coll.FindOne(ctx, filter, opts)
}

func (db *DB_base) Find(table_name string, filter interface{}, opts *options.FindOptions) *mongo.Cursor {
	base := db.get_conn()
	if base == nil {
		return nil
	}
	ctx, cancel := context.WithCancel(db.db_ctx_base)
	defer cancel()
	coll := base.Collection(table_name)
	res, err := coll.Find(ctx, filter, opts)
	if err != nil {
		db.db_err = err.Error()
		return nil
	}
	return res
}

//func (db *DB_base) aggregate(  table_name string,  mongocxx::pipeline& pipline,  mongocxx::options::aggregate& aggregate= mongocxx::options::aggregate{}){
//}
func (db *DB_base) Updateone(table_name string, filter interface{}, update interface{}, opts *options.UpdateOptions) *mongo.UpdateResult {
	base := db.get_conn()
	if base == nil {
		return nil
	}
	ctx, cancel := context.WithCancel(db.db_ctx_base)
	defer cancel()
	coll := base.Collection(table_name)
	res, err := coll.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		db.db_err = err.Error()
		return nil
	}
	return res
}

func (db *DB_base) Updatemany(table_name string, filter interface{}, update interface{}, opts *options.UpdateOptions) *mongo.UpdateResult {
	base := db.get_conn()
	if base == nil {
		return nil
	}
	ctx, cancel := context.WithCancel(db.db_ctx_base)
	defer cancel()
	coll := base.Collection(table_name)
	res, err := coll.UpdateMany(ctx, filter, update, opts)
	if err != nil {
		db.db_err = err.Error()
		return nil
	}
	return res
}
func (db *DB_base) Deleteone(table_name string, filter interface{}, options interface{}) {
}
func (db *DB_base) Deletemany(table_name string, filter interface{}, options interface{}) {
}
func (db *DB_base) Create_index(table_name string, keys interface{}, unique bool, opts *options.CreateIndexesOptions) string {
	base := db.get_conn()
	if base == nil {
		return ""
	}
	coll := base.Collection(table_name)
	op := options.Index().SetUnique(unique)
	model := mongo.IndexModel{
		Keys:    keys,
		Options: op,
	}
	ctx, cancel := context.WithCancel(db.db_ctx_base)
	defer cancel()
	res, err := coll.Indexes().CreateOne(ctx, model, opts)
	//res, err := coll.Indexes().CreateMany()
	if err != nil {
		db.db_err = err.Error()
		return ""
	}
	return res
}
func (db *DB_base) Get_count(table_name string, bObj interface{}) {
}
func (db *DB_base) ClearTable(table_name string) {
}
func (db *DB_base) Rename(table_name string, target string) {
}
func (db *DB_base) FindAndModify(table_name string, filter interface{}, update interface{}, options interface{}) {
}
func (db *DB_base) FindAndRemove(table_name string, filter interface{}, update interface{}, options interface{}) {
}
func (db *DB_base) Has_collection(table_name string) {
}
func (db *DB_base) Get_last_error() string {
	return db.db_err
}

//func (db *DB_base) init_index() {
//}
func (db *DB_base) Close() {
	db.b_close = true
	db.f_cancel()
}

func (db *DB_base) get_conn() *mongo.Database {
	if !db.b_init {
		return nil
	}

	if db.b_close {
		return nil
	}

	if db.db_client == nil {
		if (len(db.db_user) > 0) && (len(db.db_pwd) > 0) {
			db.db_uri = "mongodb://" + db.db_user + ":" + db.db_pwd + "@" + db.db_addr + "/?authSource=admin"
		} else {
			db.db_uri = "mongodb://" + db.db_addr + "/"
		}
		ctx, cancel := context.WithCancel(context.Background())

		db.db_ctx_base = ctx
		db.f_cancel = cancel

		var err error

		var client *mongo.Client
		client, err = mongo.Connect(db.db_ctx_base, options.Client().ApplyURI(db.db_uri))
		if err != nil {
			db.db_err = err.Error()
			return nil
		}

		db.db_client = client

		var dbase *mongo.Database = client.Database(db.db_name)
		db.db_database = dbase

		db.db_err = ""
	}
	return db.db_database
}
