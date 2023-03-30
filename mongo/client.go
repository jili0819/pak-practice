package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type Client struct {
	*mongo.Client
	tableName string
}

func (c *Client) SaveOn(ctx context.Context, collection string, data interface{}) {
}
