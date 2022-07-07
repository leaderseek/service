package transform

import (
	sp "github.com/leaderseek/api-go/service/param"
	wp "github.com/leaderseek/definition/workflow/param"
	"github.com/leaderseek/service/pkg/config"
)

func TeamAddMembersRequest(cfg *config.ServerConfig, in *sp.TeamAddMembersRequest) *wp.TeamAddMembersRequest {
	out := new(wp.TeamAddMembersRequest)

	out.DBConnectionString = cfg.DBConnection
	out.TeamID = in.TeamID

	for _, ip := range in.Players {
		out.Players = append(out.Players, player(ip))
	}

	return out
}
