package domain

import (
	"mpbe/errs"
	"reflect"
	"testing"

	"gorm.io/gorm"
)

func TestNewEmployeeRepositoryDB(t *testing.T) {
	type args struct {
		client *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want EmployeeRepositoryDB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEmployeeRepositoryDB(tt.args.client); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEmployeeRepositoryDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmployeeRepositoryDB_FindAll(t *testing.T) {
	tests := []struct {
		name  string
		e     *EmployeeRepositoryDB
		want  []Employees
		want1 *errs.AppErr
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.e.FindAll()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EmployeeRepositoryDB.FindAll() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("EmployeeRepositoryDB.FindAll() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestEmployeeRepositoryDB_FindByID(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name  string
		e     EmployeeRepositoryDB
		args  args
		want  Employees
		want1 *errs.AppErr
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.e.FindByID(tt.args.id)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EmployeeRepositoryDB.FindByID() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("EmployeeRepositoryDB.FindByID() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestEmployeeRepositoryDB_CreateEmployeeInput(t *testing.T) {
	type args struct {
		employees Employees
	}
	tests := []struct {
		name  string
		e     EmployeeRepositoryDB
		args  args
		want  Employees
		want1 *errs.AppErr
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.e.CreateEmployeeInput(tt.args.employees)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EmployeeRepositoryDB.CreateEmployeeInput() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("EmployeeRepositoryDB.CreateEmployeeInput() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestEmployeeRepositoryDB_DeleteEmployee(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name  string
		e     EmployeeRepositoryDB
		args  args
		want  Employees
		want1 *errs.AppErr
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.e.DeleteEmployee(tt.args.id)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EmployeeRepositoryDB.DeleteEmployee() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("EmployeeRepositoryDB.DeleteEmployee() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestEmployeeRepositoryDB_UpdateEmployee(t *testing.T) {
	type args struct {
		id        int
		employees Employees
	}
	tests := []struct {
		name  string
		e     EmployeeRepositoryDB
		args  args
		want  Employees
		want1 *errs.AppErr
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.e.UpdateEmployee(tt.args.id, tt.args.employees)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EmployeeRepositoryDB.UpdateEmployee() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("EmployeeRepositoryDB.UpdateEmployee() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
