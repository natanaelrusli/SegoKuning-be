package model

// NOTE-CLEAN1: model (or domain model layer), represent the core data structure and business logic, this entity encapsulate the essential data and behaviour
// in BE, typically the model will be returned to the user via JSON
// Don't think that this model is the same with what we store in the database, because in clean architecture we need to have separation of concern for the model to not depend to anything

type Image struct {
	ID  int    `json:"id"`
	URL string `json:"url"`
}
