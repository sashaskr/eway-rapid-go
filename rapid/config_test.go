package rapid

import (
	"reflect"
	"testing"
)

func TestNewConfig(t *testing.T) {
	type args struct {
		testing     bool
		apiKey      string
		apiPassword string
	}

	tests := []struct {
		name string
		args args
		want *Config
	}{
		{
			"config set to testing and with API token and app token",
			args{
				testing:     true,
				apiKey:      APIKeyEnv,
				apiPassword: APIPasswordEnv,
			},
			&Config{
				testing:     true,
				apiKey:      "EWAY_API_KEY",
				apiPassword: "EWAY_API_PASSWORD",
			},
		},
		{
			"config set to testing with develop and only app token",
			args{
				testing:     true,
				apiPassword: APIPasswordEnv,
			},
			&Config{
				testing:     true,
				apiPassword: "EWAY_API_PASSWORD",
			},
		},
		{
			"config set to testing with develop and only API token",
			args{
				testing: true,
				apiKey:  APIKeyEnv,
			},
			&Config{
				testing: true,
				apiKey:  "EWAY_API_KEY",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConfig(tt.args.testing, tt.args.apiKey, tt.args.apiPassword); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
