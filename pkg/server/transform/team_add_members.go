package transform

import (
	sp "github.com/leaderseek/api-go/service/param"
	wp "github.com/leaderseek/definition/workflow/param"
	"github.com/leaderseek/service/pkg/server/config"
)

func TeamAddMembersRequest(cfg *config.Config, in *sp.TeamAddMembersRequest) *wp.TeamAddMembersRequest {
	out := new(wp.TeamAddMembersRequest)

	out.DBConnectionString = cfg.DBConnectionString
	out.TeamID = in.TeamID

	for _, ip := range in.Players {
		out.Players = append(out.Players, player(ip))
	}

	return out
}
