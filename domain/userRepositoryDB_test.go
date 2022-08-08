package domain

import (
	"mpbe/errs"
	"reflect"
	"testing"

	"gorm.io/gorm"
)

func TestNewUserRepositoryDB(t *testing.T) {
	type args struct {
		client *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want UserRepositoryDB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserRepositoryDB(tt.args.client); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserRepositoryDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepositoryDB_RegisterUserInput(t *testing.T) {
	type args struct {
		users Users
	}
	tests := []struct {
		name  string
		u     UserRepositoryDB
		args  args
		want  Users
		want1 *errs.AppErr
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.u.RegisterUserInput(tt.args.users)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserRepositoryDB.RegisterUserInput() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("UserRepositoryDB.RegisterUserInput() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestUserRepositoryDB_LoginUserInput(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name  string
		u     UserRepositoryDB
		args  args
		want  Users
		want1 *errs.AppErr
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.u.LoginUserInput(tt.args.email)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserRepositoryDB.LoginUserInput() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("UserRepositoryDB.LoginUserInput() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
