// Package repository contains methods to communicate with postgres and gRPC servers
package repository

import (
	"context"
	"fmt"

	proto "github.com/eugenshima/price-service/proto"
	"github.com/eugenshima/trading-service/internal/model"
)

// PriceServiceClient struct ....
type PriceServiceClient struct {
	client proto.PriceServiceClient
}

// NewPriceServiceClient creates a new PriceServiceClient
func NewPriceServiceClient(client proto.PriceServiceClient) *PriceServiceClient {
	return &PriceServiceClient{client: client}
}

// AddSubscriber method adds a subscriber to the list of subscribers
func (c *PriceServiceClient) AddSubscriber(ctx context.Context, selectedShares []string) (*model.Share, error) {
	stream, err := c.client.Subscribe(ctx, &proto.SubscribeRequest{ShareName: selectedShares})
	if err != nil {
		return nil, fmt.Errorf("subscribe: %w", err)
	}
	shares, err := stream.Recv()
	if err != nil {
		return nil, fmt.Errorf("recv: %w", err)
	}

	share := &model.Share{
		ShareName:  shares.Shares[0].ShareName,
		SharePrice: shares.Shares[0].SharePrice,
	}

	return share, nil
}
