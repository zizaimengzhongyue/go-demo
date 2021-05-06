package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const DB = "test"
const COL = "my-collection"

var client *mongo.Client

func Create() {
	err := client.Database(DB).CreateCollection(context.Background(), COL)
	if err != nil {
		panic(err)
	}
}

func count(col *mongo.Collection) {
	filter := bson.D{}
	count, err := col.CountDocuments(context.Background(), filter)
	if err != nil {
		panic(err)
	}
	fmt.Printf("当前文档数量为: %d\n", count)
}

func Insert() {
	ctx := context.Background()
	col := client.Database(DB).Collection(COL)

	doc := map[string]interface{}{
		"title": "第一条数据", "id": 1, "content": "这是插入的第一条数据",
	}
	singleRes, err := col.InsertOne(ctx, doc)
	if err != nil {
		panic(err)
	}
	fmt.Println(singleRes)
	count(col)

	docs := []interface{}{
		map[string]interface{}{"title": "第二条数据", "id": 2, "content": "这是插入的第二条数据"},
		map[string]interface{}{"title": "第三条数据", "id": 3, "content": "这是插入的第三条数据"},
	}
	multiRes, err := col.InsertMany(ctx, docs)
	if err != nil {
		panic(err)
	}
	fmt.Println(multiRes)
	count(col)
}

func Delete() {
	ctx := context.Background()
	col := client.Database(DB).Collection(COL)

	count(col)

	filter := bson.D{{"id", 1}}
	res, err := col.DeleteOne(ctx, filter)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.DeletedCount)
	count(col)

	filter = bson.D{}
	res, err = col.DeleteMany(ctx, filter)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.DeletedCount)
	count(col)
}

func Update() {
	ctx := context.Background()
	col := client.Database(DB).Collection(COL)

	Find()

	filter := bson.D{{"id", 1}}
	data := bson.D{{"$set", bson.D{{"content", "这是更新后的文章"}}}}
	res, err := col.UpdateOne(ctx, filter, data)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
	Find()

	filter = bson.D{}
	data = bson.D{{"$set", bson.D{{"content", "这是第二次更新后的文章"}}}}
	res, err = col.UpdateMany(ctx, filter, data)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
	Find()
}

func Find() {
	ctx := context.Background()
	col := client.Database(DB).Collection(COL)

	data := map[string]interface{}{}
	filter := bson.D{{"id", 1}}
	err := col.FindOne(ctx, filter).Decode(&data)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)

	datas := []map[string]interface{}{}
	filter = bson.D{}
	cur, err := col.Find(ctx, filter)
	if err != nil {
		panic(err)
	}
	err = cur.All(ctx, &datas)
	if err != nil {
		panic(err)
	}
	fmt.Println(datas)
}

func Drop() {
	ctx := context.Background()
	err := client.Database(DB).Collection(COL).Drop(ctx)
	if err != nil {
		panic(err)
	}
}

func main() {
	Create()
	defer Drop()
	Insert()
	Find()
	Update()
	Delete()
}

func init() {
	var err error
	client, err = mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
}
