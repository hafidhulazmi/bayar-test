package main

import "testing"

func TestHitungHargaTotal(t *testing.T) {
	type args struct {
		hargaItem float64
		ongkir    float64
		qty       int
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "failed in harga item",
			args: args{
				hargaItem: 0,
				ongkir:    1,
				qty:       1,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "failed in ongkir",
			args: args{
				hargaItem: 1,
				ongkir:    0,
				qty:       1,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "failed in qty",
			args: args{
				hargaItem: 1,
				ongkir:    1,
				qty:       0,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "success",
			args: args{
				hargaItem: 15000,
				ongkir:    10000,
				qty:       2,
			},
			want:    45000,
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HitungHargaTotal(tt.args.hargaItem, tt.args.ongkir, tt.args.qty)
			if (err != nil) != tt.wantErr {
				t.Errorf("HitungHargaTotal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HitungHargaTotal() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPembayaranBarang(t *testing.T) {
	type args struct {
		hargaTotal       float64
		metodePembayaran string
		dicicil          bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "failed in harga total",
			args: args{
				hargaTotal: 0,
			},
			wantErr: true,
		},
		{
			name: "failed in metode pembayaran",
			args: args{
				metodePembayaran: "",
			},
			wantErr: true,
		},
		{
			name: "success",
			args: args{
				hargaTotal:       500000,
				metodePembayaran: "kredit",
				dicicil:          true,
			},
			wantErr: false,
		},
		{
			name: "success 1",
			args: args{
				hargaTotal:       500000,
				metodePembayaran: "cod",
				dicicil:          false,
			},
			wantErr: false,
		},
		{
			name: "success 2",
			args: args{
				hargaTotal:       500000,
				metodePembayaran: "transfer",
				dicicil:          false,
			},
			wantErr: false,
		},
		{
			name: "success 3",
			args: args{
				hargaTotal:       500000,
				metodePembayaran: "debit",
				dicicil:          false,
			},
			wantErr: false,
		},
		{
			name: "success 4",
			args: args{
				hargaTotal:       500000,
				metodePembayaran: "gerai",
				dicicil:          false,
			},
			wantErr: false,
		},
		{
			name: "failed in harga total",
			args: args{
				hargaTotal:       499999,
				metodePembayaran: "kredit",
				dicicil:          true,
			},
			wantErr: true,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := PembayaranBarang(tt.args.hargaTotal, tt.args.metodePembayaran, tt.args.dicicil)
			if (err != nil) != tt.wantErr {
				t.Errorf("PembayaranBarang() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
