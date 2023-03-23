/*
Copyright © 2022 The Infratographer Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	dbm "go.infratographer.com/tenant-api/db"
	"go.infratographer.com/tenant-api/internal/config"
	"go.infratographer.com/x/crdbx"
	"go.infratographer.com/x/goosex"
	"go.infratographer.com/x/loggingx"
	"go.infratographer.com/x/otelx"
	"go.infratographer.com/x/versionx"
	"go.infratographer.com/x/viperx"
	"go.uber.org/zap"
)

const appName = "tenant-api"

var (
	cfgFile string
	logger  *zap.Logger
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "infra-tenant-api",
	Short: "Infratographer Tenant API Service handles tenant hierarchy",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is /etc/infratographer/tenant-api.yaml)")
	viperx.MustBindFlag(viper.GetViper(), "config", rootCmd.PersistentFlags().Lookup("config"))

	rootCmd.PersistentFlags().String("nats-url", "nats://nats:4222", "NATS server connection url")
	viperx.MustBindFlag(viper.GetViper(), "nats.url", rootCmd.PersistentFlags().Lookup("nats-url"))

	rootCmd.PersistentFlags().String("nats-creds-file", "", "Path to the file containing the NATS nkey keypair")
	viperx.MustBindFlag(viper.GetViper(), "nats.creds-file", rootCmd.PersistentFlags().Lookup("nats-creds-file"))

	rootCmd.PersistentFlags().String("nats-subject-prefix", "com.infratographer.events", "prefix for NATS subjects")
	viperx.MustBindFlag(viper.GetViper(), "nats.subject-prefix", rootCmd.PersistentFlags().Lookup("nats-subject-prefix"))

	rootCmd.PersistentFlags().String("nats-stream-name", "tenant-api", "nats stream name")
	viperx.MustBindFlag(viper.GetViper(), "nats.stream-name", rootCmd.PersistentFlags().Lookup("nats-stream-name"))

	// Logging flags
	loggingx.MustViperFlags(viper.GetViper(), rootCmd.PersistentFlags())

	// Register version command
	versionx.RegisterCobraCommand(rootCmd, func() { versionx.PrintVersion(logger.Sugar()) })

	// OTEL Flags
	otelx.MustViperFlags(viper.GetViper(), rootCmd.Flags())

	// Database Flags
	crdbx.MustViperFlags(viper.GetViper(), rootCmd.Flags())

	// Add migrate command
	goosex.RegisterCobraCommand(rootCmd, func() {
		goosex.SetBaseFS(dbm.Migrations)
		goosex.SetLogger(logger.Sugar())
		goosex.SetDBURI(config.AppConfig.CRDB.GetURI())
	})
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath("/etc/infratographer/")
		viper.SetConfigType("yaml")
		viper.SetConfigName("tenant-api")
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.SetEnvPrefix("tenantapi")

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	err := viper.ReadInConfig()

	logger = loggingx.InitLogger(appName, config.AppConfig.Logging).Desugar()

	if err == nil {
		logger.Info("using config file",
			zap.String("file", viper.ConfigFileUsed()),
		)
	}

	err = viper.Unmarshal(&config.AppConfig)
	if err != nil {
		logger.Fatal("unable to decode app config", zap.Error(err))
	}
}
