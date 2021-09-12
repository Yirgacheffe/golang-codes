package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var verbose bool
var author string
var region string

var rootCmd = &cobra.Command{
	Use:   "unlessctl",
	Long:  `Unless platform configuration command line utility for service operators to debug and diagnose.`,
	Short: "Unless control interface",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("cmd goes here...!")
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Long:  `All software has version.`,
	Short: "Print the version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("unlessctl v0.9 -- HEAD")
	},
}

func newCommand() *cobra.Command {
	return &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Foo")
		},
		Use:   "foo",
		Short: "Command foo",
		Long:  "This is a new Command Foo",
	}
}

func init() {
	// cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is &HOME/unless.yaml)")

	rootCmd.PersistentFlags().StringVar(&author, "author", "midfall", "Author name for copyright attribution")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().Bool("viper", true, "Use Viper for configuration")

	rootCmd.Flags().StringVarP(&region, "region", "r", "", "AWS region (required)")
	// rootCmd.MarkFlagRequired("region")
	rootCmd.AddCommand(versionCmd)

	// rootCmd.SetVersionTemplate("")
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
		return
	}

	home := os.Getenv("HOME") // home, err := homedir.Dir()
	/*
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	*/
	viper.AddConfigPath(home)
	viper.SetConfigName(".cobra") // search config with name ".cobra"

	// viper.SetConfigName("config")
	// viper.SetConfigType("yaml")
	// viper.AddConfigPath("/etc/appname")
	// viper.AddConfigPath("&HOME/.appname")
	// viper.AddConfigPath(".")
	// viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
		// os.Exit(1)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
