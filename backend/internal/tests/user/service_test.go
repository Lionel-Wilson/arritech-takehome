package user

import (
	"github.com/Lionel-Wilson/arritech-takehome/internal/entity"
	"github.com/Lionel-Wilson/arritech-takehome/internal/user"
	"github.com/Lionel-Wilson/arritech-takehome/internal/user/domain"
	"github.com/Lionel-Wilson/arritech-takehome/internal/user/storage"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"testing"
)

func Test_service_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := storage.NewMockUserRepository(ctrl)
	logger := logrus.New()
	userService := user.NewUserService(logger, mockRepository)

	tests := []struct {
		name    string
		user    domain.User
		mock    func()
		wantErr error
	}{
		{
			name: "successfully create user",
			user: domain.User{
				Firstname: "Son",
				Lastname:  "Goku",
				Age:       30,
				Email:     "Kakarot@gmail.com",
			},
			mock: func() {

				mockRepository.EXPECT().InsertUser(gomock.Any(),
					&entity.User{
						Firstname: "Son",
						Lastname:  "Goku",
						Age:       30,
						Email:     "Kakarot@gmail.com",
					},
				).Return(nil)

			},
			wantErr: nil,
		},
		{
			name: "under age error",
			user: domain.User{
				Firstname: "Son",
				Lastname:  "Goku",
				Age:       12,
				Email:     "Kakarot@gmail.com",
			},
			mock: func() {

			},
			wantErr: user.ErrUserMustBeAtLeast18YearsOld,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				tt.mock()

				ctx := t.Context()

				err := userService.CreateUser(ctx, tt.user)
				if tt.wantErr == nil {
					require.NoError(t, err)
				} else {
					require.EqualError(t, err, tt.wantErr.Error())
				}
			},
		)
	}
}
