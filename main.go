package main

import (
  "fmt"
  "os"
  "log"
  "encoding/json"
  "github.com/sensu/sensu-go/types"
)

// type ObjectMeta struct{
//   Name string `json:"name"`
//   Namespace string `json:"namespace"`
// }
//
// type Payload struct{
//   ApiVersion string `json:"api_version"`
//   ResourceType string `json:"type"`
//   Metadata ObjectMeta `json:"metadata"`
// }

func main() {
  var stdin *os.File
  var event types.Event

  stdin = os.Stdin

  err := json.NewDecoder(stdin).Decode(&event)
  if err != nil {
      log.Fatal(err)
  }

  if event.Metrics != nil {
    for _, point := range(event.Metrics.Points) {
      point.Tags = append(point.Tags, &types.MetricTag{ Name: "entity", Value: event.Entity.Name } )
      if event.Entity.Labels != nil {
        for name, value := range(event.Entity.Labels) {
          point.Tags = append( point.Tags, &types.MetricTag{ Name: name, Value: value } )
        }
      }
    }
  }

  output, err := json.MarshalIndent(event, "", "  ")
  if err != nil {
      log.Fatal(err)
  }

  os.Stdout.Write(output)
  fmt.Println()
}
