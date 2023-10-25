package service

import (
	"context"
	"errors"
	"go.uber.org/mock/gomock"
	"reflect"
	"testing"
	"wb-tech/internal/model"
)

func TestService_GetFromCache(t *testing.T) {
	existsUID := "testUid123"
	missingUID := "missingUID12"
	tests := []struct {
		name              string
		uid               string
		expectedOrderData model.OrderData
		wantErr           bool
	}{
		{
			name: "Success",
			uid:  existsUID,
			expectedOrderData: model.OrderData{
				OrderUid: existsUID,
			},
			wantErr: false,
		},
		{
			name:    "Order data not found",
			uid:     missingUID,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := NewMockrepository(ctrl)

			s := &Service{
				sc:   nil,
				repo: mockRepo,
				Cache: map[string]model.OrderData{
					existsUID: model.OrderData{
						OrderUid: existsUID,
					},
				},
			}

			got, err := s.GetFromCache(tt.uid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFromCache() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.expectedOrderData) {
				t.Errorf("GetFromCache() got = %v, want %v", got, tt.expectedOrderData)
			}
		})
	}
}

func TestService_LoadCache(t *testing.T) {
	tests := []struct {
		name       string
		repoBehave func(repository *Mockrepository)
		wantErr    bool
	}{
		{
			name: "success",
			repoBehave: func(repository *Mockrepository) {
				repository.EXPECT().LoadCache(gomock.Any()).Return(map[string]model.OrderData{}, nil)
			},
			wantErr: false,
		},
		{
			name: "success",
			repoBehave: func(repository *Mockrepository) {
				repository.EXPECT().LoadCache(gomock.Any()).Return(map[string]model.OrderData{}, errors.New("same error"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := NewMockrepository(ctrl)

			if tt.repoBehave != nil {
				tt.repoBehave(mockRepo)
			}

			s := Service{
				sc:    nil,
				repo:  mockRepo,
				Cache: map[string]model.OrderData{},
			}
			if err := s.LoadCache(context.Background()); (err != nil) != tt.wantErr {
				t.Errorf("LoadCache() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
