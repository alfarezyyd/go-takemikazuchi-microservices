package configs

import (
	"github.com/spf13/viper"
	"googlemaps.github.io/maps"
)

type GoogleMaps struct {
	mapsClient  *maps.Client
	viperConfig *viper.Viper
}

func NewGoogleMaps(viperConfig *viper.Viper) *GoogleMaps {
	return &GoogleMaps{
		viperConfig: viperConfig,
	}
}

func (googleMapsClient *GoogleMaps) InitializeGoogleMaps() *maps.Client {
	if googleMapsClient.mapsClient == nil {
		client, err := maps.NewClient(maps.WithAPIKey(googleMapsClient.viperConfig.GetString("GOOGLE_MAPS_API_KEY")))
		if err != nil {
			return nil
		}
		googleMapsClient.mapsClient = client
	}
	return googleMapsClient.mapsClient
}
