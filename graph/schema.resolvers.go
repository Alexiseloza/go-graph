package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.42

import (
	"context"
	"go-graph/database"
	"go-graph/graph/model"
)

var db = database.Connect()

// CreateJoblisting is the resolver for the createJoblisting field.
func (r *mutationResolver) CreateJoblisting(ctx context.Context, input model.CreateJoblistingInput) (*model.JobListing, error) {
	return db.CreateJoblisting(input), nil
}

// UpdateJoblisting is the resolver for the updateJoblisting field.
func (r *mutationResolver) UpdateJoblisting(ctx context.Context, id string, input model.UpdateJoblistingInput) (*model.JobListing, error) {
	return db.UpdateJoblisting(id, input), nil
}

// DeleteJoblisting is the resolver for the deleteJoblisting field.
func (r *mutationResolver) DeleteJoblisting(ctx context.Context, id string) (*model.DeleteJobResponse, error) {
	return db.DeleteJoblisting(id), nil
}

// Jobs is the resolver for the jobs field.
func (r *queryResolver) Jobs(ctx context.Context) ([]*model.JobListing, error) {
	return db.GetJobs(), nil
}

// Job is the resolver for the job field.
func (r *queryResolver) Job(ctx context.Context, id string) (*model.JobListing, error) {
	return db.GetJob(id), nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }