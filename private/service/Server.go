
package service

import (
	"context"
	"errors"
	"fmt"
	"gRPC-Gateway/private/pb"
	"github.com/jeanphorn/log4go"
)

type Server struct {}
func (s *Server) Echo(ctx context.Context, in *pb.ReqMsg) (*pb.RspMsg, error) {
	var err error
	var rsp  =new(pb.RspMsg)

	log4go.LOGGER("Test").Info("test begin")
	log4go.LOGGER("Test").Info("input name=%s, age=%d", in.Name, in.Age)
	log4go.LOGGER("Test").Info("input=%+v", in)
	if in.Name == "" {
		err = errors.New("the field of Name is nil")
	} else {
		rsp.Age_Name = fmt.Sprintf("Name:%s, Age:%d", in.Name, in.Age)
	}

	return rsp, err
}
