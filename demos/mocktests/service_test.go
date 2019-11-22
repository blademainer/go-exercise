package mock

import (
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func Test_aImpl_Hello(t *testing.T) {
	type fields struct {
		UserRepository UserRepository
	}
	user1 := &User{1, "zhangsan", 20}
	user2 := &User{2, "zhanger", 40}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock := NewMockUserRepository(ctrl)
	mock.EXPECT().Get(user1.id).Return(user1, nil).AnyTimes()
	mock.EXPECT().Get(user2.id).Return(user2, nil).AnyTimes()

	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *User
		wantErr bool
	}{
		{
			name:    "test-mock",
			fields:  fields{mock},
			args:    args{user1.id},
			want:    user1,
			wantErr: false,
		},
		{
			name:    "test-mock2",
			fields:  fields{mock},
			args:    args{user2.id},
			want:    user2,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &aImpl{
				UserRepository: tt.fields.UserRepository,
			}
			got, err := a.Hello(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Hello() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Hello() got = %v, want %v", got, tt.want)
			}
		})
	}
}
