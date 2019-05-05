package conf

import "testing"

func TestSetup(t *testing.T) {
	type args struct {
		config map[string]interface{}
	}

	type Mysql struct {
		Name     string `ini:"name"`
		Username string `ini:"username"`
		Password string `ini:"password"`
		Host     string `ini:"host"`
		Port     int    `ini:"port"`
	}

	type Def struct {
		Environment string `ini:"environment"`
		Prefix      string `ini:"prefix"`
		Port        int    `ini:"port"`
		Debug       bool   `ini:"debug"`
		Mysql       Mysql  `ini:"mysql"`
	}

	def := new(Def)

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"t1", args{
			map[string]interface{}{
				"name":      "link_manager",
				"directory": "./demo",
				"data":      def,
			},
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Setup(tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("Setup() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
