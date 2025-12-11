package interfaces

import (
	"context"
	"ewallet-ums/external"
)

type IExternal interface {
	CreateWallet(ctx context.Context, userID int) (*external.Wallet, error)
}