package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type Trainer struct {
	Name string
	Age  int
	city string
}

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

type db_test_st struct {
}

func Test() {
	db := DB_base{}
	//db.set_userpwd()
	db.init_db("10.2.2.11:27017", "test1")

	ash := Trainer{"Ash", 100, "city"}
	//a := map[string]int32{"a": 1}
	//xx, err := bson.MarshalJSON(ash) //
	//if err != nil {
	//	fmt.Printf("MarshalJSON err:%s", err.Error())
	//	return
	//}
	//fmt.Print(string(xx))
	//var bdoc interface{}
	////err = bson.UnmarshalJSON([]byte(`{"id": 1,"name": "A green door","price": 12.50,"tags": ["home", "green"]}`), &bdoc)
	//err = bson.UnmarshalJSON([]byte(string(xx)), &bdoc)

	//c := db.get_conn()
	//c.Collection("a").InsertOne(context.Background(), bdoc)
	db.Insert("a", bson.M{"age": ash.Age, "city": ash.city, "name": ash.Name})
}

func Test1() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second) //必须到达指定的时间，才会调用 cancel
	defer cancel()

	var err error
	var client *mongo.Client
	client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://10.2.2.11:27017"))
	if err != nil {

		fmt.Println(err)
		return
	}
	var dbase *mongo.Database = client.Database("test")
	var collection *mongo.Collection = dbase.Collection("bbb")

	ash := Trainer{"Ash", 100, "city"}
	var insertResult *mongo.InsertOneResult
	//insertResult, err = collection.InsertOne(context.TODO(), ash)
	insertResult, err = collection.InsertOne(context.TODO(), bson.M{"name": ash.Name, "age": ash.Age, "city": ash.city})
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	if false {

		for i := 0; i < 10; i++ {

			ash := Trainer{"Ash", i, "Pallet Town"}
			var insertResult *mongo.InsertOneResult
			insertResult, err = collection.InsertOne(context.TODO(), ash)
			fmt.Println("Inserted a single document: ", insertResult.InsertedID)
		}
	}

	if err != nil {
		fmt.Println("---------err:")
		fmt.Println(err)
	}

	//bson.M 对应 type M map[string]interface{}
	//bson.D 对应  type D []DocElem

	{
		//单个
		//*mongo.SingleResult
		var result *mongo.SingleResult = collection.FindOne(ctx, bson.M{"age": 3})
		//var result *mongo.SingleResult = collection.FindOne(ctx, bson.M{})
		var bsonDoc Trainer
		//转成数据实例
		err = result.Decode(&bsonDoc)
		if err != nil {
			fmt.Println(err)
		}
	}

	if false {
		//多个
		findOptions := options.Find()
		//findOptions.SetSort(bson.D{{"_id", 1}})
		filter := bson.M{}

		//error
		cur, err := collection.Find(context.Background(), filter, findOptions)
		var jsonMaps []map[string]interface{}
		jsonMaps = make([]map[string]interface{}, 0)
		for cur.Next(context.Background()) {
			var bsonDoc Trainer
			var jsonMap map[string]interface{}
			jsonMap = make(map[string]interface{})
			err = cur.Decode(&bsonDoc)
			if err != nil {
				fmt.Println(err)
				continue
			}
			jsonMap[bsonDoc.Name] = bsonDoc
			jsonMaps = append(jsonMaps, jsonMap)
		}
	}

	filter := bson.M{"age": 100}

	//a := map[string]int32{"a": 1}
	//xx, err := bson.MarshalJSON(a) //
	//fmt.Print(string(xx))
	//var bdoc interface{}
	////err = bson.UnmarshalJSON([]byte(`{"id": 1,"name": "A green door","price": 12.50,"tags": ["home", "green"]}`), &bdoc)
	//err = bson.UnmarshalJSON([]byte(string(xx)), &bdoc)

	upinfo := bson.M{"addkey": 666}
	update := bson.M{"$set": upinfo}
	var opt options.UpdateOptions
	bb := false
	opt.Upsert = &bb
	//res, err := collection.UpdateOne(context.Background(), filter, update, &opt)
	res, err := collection.UpdateMany(context.Background(), filter, update, &opt)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("update a single document MatchedCount: ", res.MatchedCount)

	//collection.Indexes().CreateMany()
	//collection.Indexes().CreateOne()
	//collection.UpdateOne()
	//collection.UpdateMany()
	//collection.DeleteMany()
	//collection.DeleteOne()
	//collection.Drop()
	//collection.FindOneAndUpdate()
	//collection.Aggregate()

	//fmt.Println("find a single document: ", insertResult.InsertedID)
}
func NewDB(user string, pwd string, crypto bool, _db_addr string, _db_name string) *DB_base {
	db := &DB_base{}
	db.db_user = user
	db.db_pwd = pwd
	if !db.init_db(_db_addr, _db_name) {
		return nil
	}
	return db
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
	//insertResult, err := coll.InsertMany(context.Background(), docs)
	_, err := coll.InsertMany(ctx, docs)
	if err != nil {
		db.db_err = err.Error()
	}
	//fmt.Sprintf("Inserted a single documents:%v ", insertResult.InsertedIDs)
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
