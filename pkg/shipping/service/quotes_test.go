package service

import (
	"reflect"
	"testing"

	"frete-rapido-api/pkg/shipping"
	"frete-rapido-api/pkg/shipping/repository"
)

func Test_service_Quotes(t *testing.T) {
	type (
		fields struct {
			shippingRepo func() shipping.Repository
		}

		args struct {
			shippingData *shipping.ShippingDataRequest
		}
	)

	shippingBadExample := &shipping.ShippingDataRequest{
		Shipper: shipping.ShipperInfo{
			RegisteredNumber: "123",
			Token:            "123",
			PlatformCode:     "123",
		},
		Recipient: shipping.Address{
			Type:    1,
			Country: "BRA",
			Zipcode: 123,
		},
		Dispatchers: []shipping.DispatchersInfo{
			{
				RegisteredNumber: "1",
				Zipcode:          0,
				Volumes: []shipping.Volume{
					{
						Category:      "1",
						Amount:        1,
						UnitaryWeight: 1,
						UnitaryPrice:  1,
						Price:         1,
						Height:        1,
						Width:         1,
						Length:        1,
					},
				},
			},
		},
		SimulationType: []int{
			1,
		},
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *shipping.ShippingDataResponse
		wantErr bool
	}{
		{
			name: "should return a validation error",
			fields: fields{
				shippingRepo: func() shipping.Repository {
					repositoryMock := &repository.Mock{}

					return repositoryMock
				},
			},
			args: args{
				&shipping.ShippingDataRequest{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should return a fetch error",
			fields: fields{
				shippingRepo: func() shipping.Repository {
					repositoryMock := &repository.Mock{}

					return repositoryMock
				},
			},
			args: args{
				shippingBadExample,
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				shippingRepo: tt.fields.shippingRepo(),
			}

			got, err := s.Quotes(tt.args.shippingData)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.Quotes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.Quotes() = %v, want %v", got, tt.want)
			}
		})
	}
}
