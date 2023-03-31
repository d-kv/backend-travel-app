package config

import "github.com/spf13/viper"

func Load[T any](cfg *T, path, name string) error {
	viper := viper.New()

	viper.SetConfigType("yaml")
	viper.SetConfigName(name)
	viper.AddConfigPath(path)

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return err
	}

	return nil
}
