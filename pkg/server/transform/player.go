package transform

import (
	"github.com/leaderseek/api-go/message"
	"github.com/leaderseek/sqlboiler/repository"
	"github.com/volatiletech/null/v8"
)

func player(ip *message.Player) *repository.Player {
	op := &repository.Player{
		ID:           ip.ID,
		EmailAddress: ip.EmailAddress,
	}

	op.Forenames = null.NewString(ip.Forenames, ip.Forenames != "")
	op.Surnames = null.NewString(ip.Surnames, ip.Surnames != "")
	op.MobileTelephoneAddress = null.NewString(ip.MobileTelephoneAddress, ip.MobileTelephoneAddress != "")

	return op
}
