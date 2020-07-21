package transaction

import (
	"reflect"
	"testing"
	"time"
)

func TestMapRowToTransaction(t *testing.T) {
	type args struct {
		row []string
	}
	row := []string{"0", "debit", "73555", "processing", "5921", "Thu, 02 Jan 2020 11:15:10 UTC"}
	tr := &Transaction{
		Id:     0,
		Type:   "debit",
		Sum:    735_55,
		Status: "processing",
		MCC:    "5921",
		Date:   time.Date(2020, time.January, 2, 11, 15, 10, 0, time.UTC),
	}
	tests := []struct {
		name    string
		args    args
		want    *Transaction
		wantErr bool
	}{
		{
			name: "correct data",
			args: args{
				row: row,
			},
			want:    tr,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MapRowToTransaction(tt.args.row)
			if (err != nil) != tt.wantErr {
				t.Errorf("MapRowToTransaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapRowToTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}
