package service

import (
	"fmt"
	"reflect"
	"testing"

	"frete-rapido-api/pkg/shipping"
	"frete-rapido-api/pkg/shipping/repository"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

func Test_service_Metrics(t *testing.T) {
	type (
		fields struct {
			shippingRepo func() shipping.Repository
		}
	)

	carriers := map[string]shipping.CarrierMetrics{}
	carriers["frete rapido"] = shipping.CarrierMetrics{
		Total:        50,
		AveragePrice: 50,
		Results:      1,
	}

	carriers["frete rapido 2"] = shipping.CarrierMetrics{
		Total:        10,
		AveragePrice: 5,
		Results:      2,
	}

	metricsResponse := &shipping.MetricsResponse{
		Carriers:    carriers,
		LowestPrice: 10,
		HigherPrice: 50,
	}

	tests := []struct {
		name    string
		fields  fields
		want    *shipping.MetricsResponse
		wantErr bool
	}{
		{
			name: "should return a search error",
			fields: fields{
				func() shipping.Repository {
					m := &repository.Mock{}

					m.On("Search", mock.Anything).Return(nil, fmt.Errorf("error"))
					return m
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should return a not found error",
			fields: fields{
				func() shipping.Repository {
					m := &repository.Mock{}
					m.On("Search", mock.Anything).Return([]*shipping.CarrierDBModel{}, nil)
					return m
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should return the carriers metrics",
			fields: fields{
				func() shipping.Repository {
					m := &repository.Mock{}
					m.On("Search", mock.Anything).Return([]*shipping.CarrierDBModel{
						{
							ID:       uuid.NewString(),
							Name:     "frete rapido",
							Service:  "Normal",
							Deadline: 1,
							Price:    50,
						},
						{
							ID:       uuid.NewString(),
							Name:     "frete rapido 2",
							Service:  "Normal",
							Deadline: 1,
							Price:    5,
						},
						{
							ID:       uuid.NewString(),
							Name:     "frete rapido 2",
							Service:  "Rodoviario",
							Deadline: 1,
							Price:    5,
						},
					}, nil)
					return m
				},
			},
			want:    metricsResponse,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				shippingRepo: tt.fields.shippingRepo(),
			}

			got, err := s.Metrics(&shipping.Filter{})
			if (err != nil) != tt.wantErr {
				t.Errorf("service.Metrics() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.Metrics() = %v, want %v", got, tt.want)
			}
		})
	}
}
