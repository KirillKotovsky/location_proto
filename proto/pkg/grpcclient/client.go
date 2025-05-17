package grpcclient

import (
	"context"

	locationpb "github.com/KirillKotovsky/location_proto/proto"
	"google.golang.org/grpc"
)

type Client interface {
	StoreLocation(ctx context.Context, loc *locationpb.Location) error
}

type client struct {
	conn   *grpc.ClientConn
	client locationpb.LocationTrackerClient
}

func New(conn *grpc.ClientConn) Client {
	return &client{
		conn:   conn,
		client: locationpb.NewLocationTrackerClient(conn),
	}
}

func (c *client) StoreLocation(ctx context.Context, loc *locationpb.Location) error {
	_, err := c.client.StoreLocation(ctx, loc)
	return err
}
