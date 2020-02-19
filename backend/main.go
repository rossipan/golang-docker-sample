package main

import (
    "fmt"
    "net/http"
    "github.com/spf13/viper"
)

type Config struct {
    UserName     string
}

func SetupDefault() {
    viper.SetDefault("env_test_user_name", "xrex")
}

func SetupGlobal(config *Config) {
    config.UserName = viper.GetString("env_test_user_name")
}

var configs = &Config{}

func init() {
    viper.AutomaticEnv()
    SetupDefault()
    SetupGlobal(configs)

}

func main() {
    username := configs.UserName
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello %s, you've requested: %s\n", username, r.URL.Path)
    })

    http.ListenAndServe(":8080", nil)
}