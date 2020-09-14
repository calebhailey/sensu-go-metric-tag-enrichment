package main

import (
	"strings"

	"github.com/sensu-community/sensu-plugin-sdk/sensu"
	"github.com/sensu/sensu-go/types"
)

// Config represents the mutator plugin config.
type Config struct {
	sensu.PluginConfig
	EntityName      bool
	EntityNamespace bool
	EntityLabels    bool
	CheckName       bool
	CheckLabels     bool
	EventLabels     bool
	Filter          string
	Tags            string
	Required        string
}

var (
	config = Config{
		PluginConfig: sensu.PluginConfig{
			Name:     "sensu-metric-tag-mutator",
			Short:    "Enrich Sensu Metrics with Event, Entity, and Check labels as metric tags.",
			Keyspace: "sensu.io/plugins/sensu-metric-tag-mutator/config",
		},
	}

	options = []*sensu.PluginConfigOption{
		&sensu.PluginConfigOption{
			Path:      "entity",
			Env:       "",
			Argument:  "entity",
			Shorthand: "e",
			Default:   false,
			Usage:     "Add an \"entity\" tag containing the entity name to every metric.",
			Value:     &config.EntityName,
		},
		&sensu.PluginConfigOption{
			Path:      "namespace",
			Env:       "",
			Argument:  "namespace",
			Shorthand: "n",
			Default:   false,
			Usage:     "Add a \"namespace\" tag containing the entity namespace to every metric.",
			Value:     &config.EntityNamespace,
		},
		&sensu.PluginConfigOption{
			Path:      "entity-labels",
			Env:       "",
			Argument:  "entity-labels",
			Shorthand: "",
			Default:   false,
			Usage:     "Add a tag for every entity label to every metric.",
			Value:     &config.EntityLabels,
		},
		&sensu.PluginConfigOption{
			Path:      "check",
			Env:       "",
			Argument:  "check",
			Shorthand: "c",
			Default:   false,
			Usage:     "Add a \"check\" tag containing the check name to every metric.",
			Value:     &config.CheckName,
		},
		&sensu.PluginConfigOption{
			Path:      "check-labels",
			Env:       "",
			Argument:  "check-labels",
			Shorthand: "",
			Default:   false,
			Usage:     "Add a tag for every check label to every metric.",
			Value:     &config.CheckLabels,
		},
		&sensu.PluginConfigOption{
			Path:      "event-labels",
			Env:       "",
			Argument:  "event-labels",
			Shorthand: "",
			Default:   false,
			Usage:     "Add a tag for every event label to every metric.",
			Value:     &config.EventLabels,
		},
		&sensu.PluginConfigOption{
			Path:      "require",
			Env:       "",
			Argument:  "require",
			Shorthand: "r",
			Default:   "",
			Usage:     "Comma separated list of tags to include on all metrics, even if no matching labels are found",
			Value:     &config.Required,
		},
		&sensu.PluginConfigOption{
			Path:      "filter",
			Env:       "",
			Argument:  "filter",
			Shorthand: "f",
			Default:   "sensu.io/managed_by",
			Usage:     "Comma separated list of tags to exclude.",
			Value:     &config.Filter,
		},
	}
)

func main() {
	mutator := sensu.NewGoMutator(&config.PluginConfig, options, checkArgs, executeMutator)
	mutator.Execute()
}

func indexOf(items []string, item string) int {
	for k, v := range items {
		if v == item {
			return k
		}
	}
	return -1
}

func checkArgs(_ *types.Event) error {
	// if len(mutatorConfig.Example) == 0 {
	// 	return fmt.Errorf("--example or MUTATOR_EXAMPLE environment variable is required")
	// }
	return nil
}

func addTag(point *types.MetricPoint, key string, value string) {
	if len(config.Filter) == 0 || indexOf(strings.Split(config.Filter, ","), key) < 0 {
		point.Tags = append(
			point.Tags,
			&types.MetricTag{
				Name:  key,
				Value: value,
			},
		)
	}
}

func extractTags(event *types.Event) {
	if event.Metrics != nil {
		for _, point := range event.Metrics.Points {
			if config.EntityNamespace {
				addTag(point, "namespace", event.Namespace)
			}
			if config.EntityName && event.Entity != nil {
				addTag(point, "entity", event.Entity.Name)
			}
			if config.CheckName && event.Check != nil {
				addTag(point, "check", event.Check.Name)
			}
			if config.EventLabels && event.Labels != nil {
				for name, value := range event.Labels {
					addTag(point, name, value)
				}
			}
			if config.EntityLabels && event.Entity != nil && event.Entity.Labels != nil {
				for name, value := range event.Entity.Labels {
					addTag(point, name, value)
				}
			}
			if config.CheckLabels && event.Check != nil && event.Check.Labels != nil {
				for name, value := range event.Check.Labels {
					addTag(point, name, value)
				}
			}
		}
	}
}

func executeMutator(event *types.Event) (*types.Event, error) {
	extractTags(event)
	return event, nil
}
