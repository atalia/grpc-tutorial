#coding:utf-8

import grpc
import route_guide_pb2_grpc as route_guide_pb2_grpc
import route_guide_pb2 as route_guide_pb2

def run():
    with grpc.insecure_channel("localhost:8080") as channel:
        stub = route_guide_pb2_grpc.RouteGuideStub(channel)
        feature = stub.GetFeature(route_guide_pb2.Point(latitude=409146138, longitude=-746188906))
        if not feature.location:
            print "Server returned incomplete feature"
            return

        if feature.name:
            print "Feature called %s at %s" % (feature.name, feature.location)
        else:
            print "Found no feature at %s" % (feature.location)

run()