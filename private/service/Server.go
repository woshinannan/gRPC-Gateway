
package service

import (
	"context"
	"errors"
	"fmt"
	"gRPC-Gateway/private/pb"
)

type Server struct {}
func (s *Server) Echo(ctx context.Context, in *pb.ReqMsg) (*pb.RspMsg, error) {
	var err error
	var rsp  =new(pb.RspMsg)

	if in.Name == "" {
		err = errors.New("the field of Name is nil")
	} else {
		rsp.Age_Name = fmt.Sprintf("Name:%s, Age:%d", in.Name, in.Age)
	}

	return rsp, err
}
