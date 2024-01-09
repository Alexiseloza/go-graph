package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"go-graph/graph/model"
)

type DB struct {
	Client *mongo.Client
}

var ApplyURI = "YOUR MONGO URI HERE"

func Connect() *DB {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(ApplyURI))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	return &DB{
		Client: client,
	}
}

// GET a simple Job or user
func (db *DB) GetJob(id string) *model.JobListing {
	jobCollection := db.Client.Database("graphql-job-board").Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	var JobListing model.JobListing
	err := jobCollection.FindOne(ctx, filter).Decode(&JobListing)
	if err != nil {
		fmt.Println(err)

	}
	return &JobListing

}

// Get all Jobs listing
func (db *DB) GetJobs() []*model.JobListing {
	jobCollection := db.Client.Database("graphql-job-board").Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var JobListing []*model.JobListing
	cursor, err := jobCollection.Find(ctx, bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	if err = cursor.All(context.TODO(), &JobListing); err != nil {
		panic(err)
	}
	return JobListing

}

// Create a new Job
func (db *DB) CreateJoblisting(jobInfo model.CreateJoblistingInput) *model.JobListing {
	jobCollection := db.Client.Database("graphql-job-board").Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	insert, err := jobCollection.InsertOne(ctx, bson.M{"title": jobInfo.Title, "description": jobInfo.Description, "url": jobInfo.URL, "company": jobInfo.Company})
	if err != nil {
		log.Fatal("Could not connect to MongoDB: ", err)
	}
	insertID := insert.InsertedID.(primitive.ObjectID).Hex()
	returnJobListing := model.JobListing{ID: insertID, Title: jobInfo.Title, Company: jobInfo.Company, Description: jobInfo.Description, URL: jobInfo.URL}
	return &returnJobListing
}

// Update Job
func (db *DB) UpdateJoblisting(jobId string, jobInfo model.UpdateJoblistingInput) *model.JobListing {
	jobCollection := db.Client.Database("graphql-job-board").Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	updateJobInfo := bson.M{}
	if jobInfo.Title != nil {
		updateJobInfo["title"] = jobInfo.Title
	}
	if jobInfo.Description != nil {
		updateJobInfo["description"] = jobInfo.Description
	}
	if jobInfo.URL != nil {
		updateJobInfo["url"] = jobInfo.URL
	}
	_id, _ := primitive.ObjectIDFromHex(jobId)
	filter := bson.M{"_id": _id}
	update := bson.M{"$set": updateJobInfo}

	results := jobCollection.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var JobListing model.JobListing
	if err := results.Decode(&JobListing); err != nil {
		log.Fatal("Error decoding the response from mongoDB: ", err)
	}
	return &JobListing

}

// delete Job
func (db *DB) DeleteJoblisting(jobId string) *model.DeleteJobResponse {
	jobCollection := db.Client.Database("graphql-job-board").Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_id, _ := primitive.ObjectIDFromHex(jobId)
	filter := bson.M{"_id": _id}
	_, err := jobCollection.DeleteOne(ctx, filter)
	if err != nil {
		fmt.Println("Could not delete the document: ", err)
		return nil
	}
	return &model.DeleteJobResponse{DeletedJobID: jobId}
}
