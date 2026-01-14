package config

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNew_OK(t *testing.T) {
	t.Setenv("DATABASE_PATH", "./data/service_database.db")
	
	cfg, err := New("test_data/config.json")

	require.NoError(t, err)
	require.Equal(t, "./data/service_database.db",cfg.DBPath)
	require.Equal(t, ":8080", cfg.Port)
	require.Equal(t, 10*time.Second, time.Duration(cfg.ReadTimeout))
	require.Equal(t, 10*time.Second, time.Duration(cfg.WriteTimeout))
}

func TestNew_Errors(t *testing.T) {
	tests := []struct {
		name       string
		envValue   string
		configPath string 
		wantErr    error
	}{ 
		{
			name       : "config file not found",
			envValue   : "./data/service_database.db",
			configPath : "test_data/missing.json",
			wantErr    : ErrConfigNotFound,
		},
		{
			name       : "DATABASE_PATH is not set",
			envValue   : "",
			configPath : "test_data/config.json",
			wantErr    : ErrDBPathNotSet,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("DATABASE_PATH", tt.envValue)

			cfg, err := New(tt.configPath)
			require.Nil(t, cfg)
			require.ErrorIs(t, err, tt.wantErr)
		})
	}
}

func TestMustNewCfg_PanicsOnError(t *testing.T) {
	t.Setenv("DATABASE_PATH", "")

	require.Panics(t, func() {
		MustNewCfg("test_data/config.json")
	})
}

func TestDuration_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    time.Duration
		wantErr bool
	}{
		{
			name  : "valid duration",
			input : `"15s"`,
			want  : 15 * time.Second,
		},
		{
			name    : "invalid format",
			input   : `"abc"`,
			wantErr : true,
		},
		{
			name : "not a string",
			input : `10`,
			wantErr: true,
		},
		{
			name: "zero duration",
			input: `"0s"`,
			wantErr: true,
		},
		{
			name: "negative duration",
			input: `"-10s"`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var d Duration
			err := json.Unmarshal([]byte(tt.input), &d)

			if tt.wantErr {
				require.Error(t, err)
				require.ErrorIs(t, err, ErrInvalidDuration)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.want, time.Duration(d))
		})
	}
}
