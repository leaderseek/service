package server

import (
	"context"

	"github.com/leaderseek/api-go/service"
	sp "github.com/leaderseek/api-go/service/param"
	"github.com/leaderseek/definition/workflow"
	wp "github.com/leaderseek/definition/workflow/param"
	"github.com/leaderseek/service/pkg/config"
	"github.com/leaderseek/service/pkg/options"
	"github.com/leaderseek/service/pkg/server/transform"
	"go.temporal.io/sdk/client"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	service.UnimplementedLeaderseekServer
	logger *zap.Logger
	client client.Client
	cfg    *config.ServerConfig
}

func NewServer(logger *zap.Logger, client client.Client, cfg *config.ServerConfig) *Server {
	return &Server{
		logger: logger,
		client: client,
		cfg:    cfg,
	}
}

func logError(logger *zap.Logger, err error, code codes.Code, msg string) error {
	logger.Error(msg, zap.Error(err))
	return status.Error(code, msg)
}

func (s *Server) TeamCreate(ctx context.Context, req *sp.TeamCreateRequest) (*sp.TeamCreateResponse, error) {
	wReq := transform.TeamCreateRequest(s.cfg, req)

	wRun, err := s.client.ExecuteWorkflow(ctx, options.StartWorkflowOptionsDefault(), workflow.TeamCreate, wReq)
	if err != nil {
		return nil, logError(s.logger, err, codes.Internal, "failed to execute workflow")
	}

	wRes := new(wp.TeamCreateResponse)
	if err := wRun.Get(ctx, wRes); err != nil {
		return nil, logError(s.logger, err, codes.Internal, "failed to complete workflow")
	}

	res := transform.TeamCreateResponse(wRes)

	return res, nil
}

func (s *Server) TeamAddMembers(ctx context.Context, req *sp.TeamAddMembersRequest) (*emptypb.Empty, error) {
	wReq := transform.TeamAddMembersRequest(s.cfg, req)

	wRun, err := s.client.ExecuteWorkflow(ctx, options.StartWorkflowOptionsDefault(), workflow.TeamAddMembers, wReq)
	if err != nil {
		return nil, logError(s.logger, err, codes.Internal, "failed to execute workflow")
	}

	if err := wRun.Get(ctx, nil); err != nil {
		return nil, logError(s.logger, err, codes.Internal, "failed to complete workflow")
	}

	return &emptypb.Empty{}, nil
}
