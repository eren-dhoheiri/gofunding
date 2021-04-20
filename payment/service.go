package payment

import (
	"backend_funding/user"
	"os"
	"strconv"

	midtrans "github.com/veritrans/go-midtrans"
)

type service struct {
}

type Service interface {
	GetPaymentResponse(transaction Transaction, user user.User) (midtrans.SnapResponse, error)
}

func NewService() *service {
	return &service{}
}

func (s *service) GetPaymentResponse(transaction Transaction, user user.User) (midtrans.SnapResponse, error) {

	MIDTRANS_SERVER_KEY := os.Getenv("MIDTRANS_SERVER_KEY")
	MIDTRANS_CLIENT_KEY := os.Getenv("MIDTRANS_CLIENT_KEY")

	midclient := midtrans.NewClient()
	midclient.ServerKey = MIDTRANS_SERVER_KEY
	midclient.ClientKey = MIDTRANS_CLIENT_KEY
	midclient.APIEnvType = midtrans.Sandbox

	snapGateway := midtrans.SnapGateway{
		Client: midclient,
	}

	snapReq := &midtrans.SnapReq{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(transaction.ID),
			GrossAmt: int64(transaction.Amount),
		},
		CustomerDetail: &midtrans.CustDetail{
			Email: user.Email,
			FName: user.Name,
		},
	}

	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return snapTokenResp, err
	}

	return snapTokenResp, nil
}