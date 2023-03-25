package config

import "github.com/spf13/viper"

func Load[T any](cfg *T, path, name string) error {
	viper := viper.New()

	viper.SetConfigType("yaml")
	viper.SetConfigName(name)
	viper.AddConfigPath(path)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return err
	}

	return nil
}
