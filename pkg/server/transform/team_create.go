package transform

import (
	sp "github.com/leaderseek/api-go/service/param"
	wp "github.com/leaderseek/definition/workflow/param"
	"github.com/leaderseek/service/pkg/config"
	"github.com/leaderseek/sqlboiler/repository"
)

func TeamCreateRequest(cfg *config.ServerConfig, in *sp.TeamCreateRequest) *wp.TeamCreateRequest {
	out := new(wp.TeamCreateRequest)

	out.DBConnectionString = cfg.DBConnection

	out.Team = &repository.Team{
		ID:          in.Team.ID,
		DisplayName: in.Team.DisplayName,
	}

	for _, ip := range in.Players {
		out.Players = append(out.Players, player(ip))
	}

	return out
}

func TeamCreateResponse(in *wp.TeamCreateResponse) *sp.TeamCreateResponse {
	return &sp.TeamCreateResponse{
		TeamID: in.TeamID,
	}
}
