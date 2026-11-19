package model

import (
	"context"
	"time"

	"github.com/mrmissx/stashbin/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Slug is the _id, so uniqueness is enforced by Mongo itself and the old
// Postgres "duplicate slug" retry keeps working via IsDuplicateKeyError.
type Document struct {
	Slug      string    `json:"key" form:"slug" bson:"_id"`
	Content   string    `json:"content" form:"content" bson:"content"`
	CreatedAt time.Time `json:"date" bson:"created_at"`
	Views     int64     `json:"views" bson:"views"`
}

const maxSlugRetries = 3

func documents(db *mongo.Database) *mongo.Collection {
	return db.Collection("documents")
}

func insertDocument(db *mongo.Database, doc *Document, retry int) error {
	if retry > maxSlugRetries {
		utils.Logger.Warning("Failed to insert document after 3 retries")
		return nil
	}

	doc.Slug = utils.CreateSlug()
	doc.CreatedAt = time.Now()
	doc.Views = 0

	_, err := documents(db).InsertOne(context.Background(), doc)
	if err != nil && mongo.IsDuplicateKeyError(err) {
		utils.Logger.Warning("Slug %s already exists, retrying...", doc.Slug)
		return insertDocument(db, doc, retry+1)
	}
	return err
}

func (doc *Document) Create(db *mongo.Database) error {
	return insertDocument(db, doc, 0)
}

func (doc *Document) incrementViews(db *mongo.Database) {
	_, err := documents(db).UpdateOne(
		context.Background(),
		bson.M{"_id": doc.Slug},
		bson.M{"$inc": bson.M{"views": 1}},
	)
	if err != nil {
		utils.Logger.Error("Failed to increment views: %v", err)
	}
}

func (doc *Document) GetBySlug(db *mongo.Database, slug string) error {
	err := documents(db).FindOne(context.Background(), bson.M{"_id": slug}).Decode(doc)
	if err != nil {
		return err
	}

	go doc.incrementViews(db)

	return nil
}
