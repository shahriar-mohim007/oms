package main

import pb "github.com/shahriar-mohim007/commons/api"

type CreateOrderRequest struct {
	Order         *pb.Order `"json": order`
	RedirectToURL string    `"json": redirectToURL`
}
