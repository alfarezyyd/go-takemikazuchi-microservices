package configs

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"github.com/spf13/viper"
)

type MidtransService struct {
	snapClient  snap.Client
	viperConfig *viper.Viper
}

func NewMidtransService(viperConfig *viper.Viper) *MidtransService {
	return &MidtransService{
		viperConfig: viperConfig,
	}
}

func (midtransService *MidtransService) InitializeMidtransConfiguration() *snap.Client {
	midtransService.snapClient.New(midtransService.viperConfig.GetString("MIDTRANS_SERVER_KEY"), midtrans.Sandbox)
	return &midtransService.snapClient
}
