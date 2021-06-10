package net

import (
	"net"
	"testing"
)

func TestIP2Number(t *testing.T) {
	type args struct {
		ip net.IP
	}
	tests := []struct {
		name    string
		args    args
		want    uint
		wantErr bool
	}{
		{
			name:    "ip1",
			args:    args{ip: net.ParseIP("192.168.0.10")},
			want:    3232235530,
			wantErr: false,
		},
		{
			name:    "ip2",
			args:    args{ip: net.ParseIP("192.168.0.")},
			want:    3232235535,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IP2Number(tt.args.ip)
			if (err != nil) != tt.wantErr {
				t.Errorf("IP2Number() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IP2Number() = %v, want %v", got, tt.want)
			}
		})
	}
}
