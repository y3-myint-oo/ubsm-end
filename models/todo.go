package models

import "gopkg.in/mgo.v2/bson"

// Todo - Sample
type Todo struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Title     string        `json:"title"`
	Completed bool          `json:"completed"`
}

// MockTodos - Mock data set for Todo
var MockTodos = []Todo{
	Todo{"507f1f77bcf86cd799439011", "I just watched a program about beavers. It was the best dam program I've ever seen.", true},
	Todo{"507f1f77bcf86cd799439012", "Why did the coffee file a police report? It got mugged.", false},
	Todo{"507f1f77bcf86cd799439013", "How does a penguin build it's house? Igloos it together.", true},
}
