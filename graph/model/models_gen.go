// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type DeleteJobResponse struct {
	DeletedJobID string `json:"deletedJobId"`
}

type JobListing struct {
	ID          string `json:"_id"`
	Title       string `json:"title"`
	Company     string `json:"company"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

type Mutation struct {
}

type Query struct {
}

type CreateJoblistingInput struct {
	Title       string `json:"title"`
	Company     string `json:"company"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

type UpdateJoblistingInput struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	URL         *string `json:"url,omitempty"`
}
