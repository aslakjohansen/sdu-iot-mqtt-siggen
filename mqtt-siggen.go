package main

import (
    "fmt"
    "io/ioutil"
    "encoding/json"
    "github.com/eclipse/paho.mqtt.golang"
)

type Sample struct {
    Time  float64 `json:"time"`
    Value float64 `json:"value"`
}

type Signal struct {
    Topic   string   `json:"topic"`
    Samples []Sample `json:"samples"`
}
type Config []Signal

const (
    config_filename string = "config.json"
)

var (
    brokers []string = []string{"tcp://127.0.0.1:1883"}
)

func create_broker () mqtt.Client {
    // configure options
    options := mqtt.NewClientOptions()
    for _, broker := range brokers {
      options.AddBroker(broker)
    }
    
    client := mqtt.NewClient(options)
    if token := client.Connect(); token.Wait() && token.Error() != nil {
        panic(token.Error())
    }
    
    return client
}

func read_config (filename string) Config {
    var config Config
    
    // read config file
    data, err := ioutil.ReadFile(config_filename)
    if err!=nil {
        panic("Unable to load config file: "+err.Error())
    }
    
    // initial partial parsing of config file
    err = json.Unmarshal(data, &config)
    if err!=nil {
        panic("Unable to unmarshal config file: "+err.Error())
    }
    
    return config
}

func main () {
    config := read_config(config_filename)
    client := create_broker()
    
    fmt.Println("Ready")
    fmt.Println(config)
    fmt.Println(client)
}
