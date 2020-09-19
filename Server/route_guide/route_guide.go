package route_guide

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"context"
	pb "../proto/routeguide"
	"sync"
)

type routeGuideServer struct {
	saveFeatures []*pb.Feature
	mu sync.Mutex
	routeNotes map[string][]*pb.RouteNote
}

var _ pb.RouteGuideServer = &routeGuideServer{}

func (s *routeGuideServer) GetFeature(ctx context.Context,point *pb.Point) (*pb.Feature, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeature not implemented")
}

func (s *routeGuideServer) ListFeatures(rectangle *pb.Rectangle, steam pb.RouteGuide_ListFeaturesServer) error {
	return status.Errorf(codes.Unimplemented, "method ListFeatures not implemented")
}

func (s *routeGuideServer) RecordRoute(steam pb.RouteGuide_RecordRouteServer) error {
	return status.Errorf(codes.Unimplemented, "method RecordRoute not implemented")
}

func (s *routeGuideServer) RouteChat(steam pb.RouteGuide_RouteChatServer) error {
	return status.Errorf(codes.Unimplemented, "method RouteChat not implemented")
}

