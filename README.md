[![Sensu Bonsai Asset](https://img.shields.io/badge/Bonsai-Download%20Me-brightgreen.svg?colorB=89C967&logo=sensu)](https://bonsai.sensu.io/assets/calebhailey/sensu-metric-tag-mutator)
![Go Test](https://github.com/calebhailey/sensu-metric-tag-mutator/workflows/Go%20Test/badge.svg)
![goreleaser](https://github.com/calebhailey/sensu-metric-tag-mutator/workflows/goreleaser/badge.svg)

# Sensu Metric Tag Mutator

## Table of Contents
- [Overview](#overview)
- [Usage examples](#usage)
- [Configuration](#configuration)
  - [Asset registration](#asset-registration)
  - [Mutator definition](#mutator-definition)
- [Installation from source](#installation-from-source)
- [Additional notes](#additional-notes)
- [Contributing](#contributing)

## Overview

The Sensu Metric Tag Mutator is a Sensu Go plugin for enriching metric points 
with [Entity metadata][entity-metadata] (e.g. entity name, labels) and 
[Check metadata][check-metadata] (e.g. check name, labels), adding them as 
[Metric Tags][metric-tags]. The Sensu Metric Tag Mutator is useful for adding 
context to generic metrics collected from plugins that generate metrics in 
formats that don't support tags (e.g. Nagios Perfdata).

[entity-metadata]: https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-entities/entities/#metadata-attributes
[check-metadata]: https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-schedule/checks/#metadata-attributes
[metric-tags]: https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-events/events/#points-attributes

## Usage

```
Enrich Sensu Metrics with Event, Entity, and Check labels as metric tags.

Usage:
  sensu-metric-tag-mutator [flags]
  sensu-metric-tag-mutator [command]

Available Commands:
  help        Help about any command
  version     Print the version number of this plugin

Flags:
  -c, --check            Add a "check" tag containing the check name to every metric.
      --check-labels     Add a tag for every check label to every metric.
  -e, --entity           Add an "entity" tag containing the entity name to every metric.
      --entity-labels    Add a tag for every entity label to every metric.
      --event-labels     Add a tag for every event label to every metric.
  -f, --filter string    Comma separated list of tags to exclude. (default "sensu.io/managed_by")
  -h, --help             help for sensu-metric-tag-mutator
  -n, --namespace        Add a "namespace" tag containing the entity namespace to every metric.
  -r, --require string   Comma separated list of tags to include on all metrics, even if no matching  labels are found

Use "sensu-metric-tag-mutator [command] --help" for more information about a command.
```

## Configuration

### Asset registration

[Sensu Assets][10] are the best way to make use of this plugin. If you're not using an asset, please
consider doing so! If you're using sensuctl 5.13 with Sensu Backend 5.13 or later, you can use the
following command to add the asset:

```
$ sensuctl asset add calebhailey/sensu-metric-tag-mutator
```

If you're using an earlier version of sensuctl, you can find the asset on the [Bonsai Asset Index](https://bonsai.sensu.io/assets/calebhailey/sensu-metric-tag-mutator).

### Mutator definition

```yml
---
type: Mutator
api_version: core/v2
metadata:
  name: sensu-metric-tag-mutator
spec:
  command: >- 
    sensu-metric-tag-mutator
    --namespace 
    --entity 
    --check
    --event-labels
    --entity-labels
    --check-labels
    --filter sensu.io/managed_by,foo,bar 
  runtime_assets:
  - calebhailey/sensu-metric-tag-mutator
```

## Installation from source

The preferred way of installing and deploying this plugin is to use it as an Asset. If you would
like to compile and install the plugin from source or contribute to it, download the latest version
or create an executable script from this source.

From the local path of the sensu-metric-tag-mutator repository:

```
go build
```

## Additional notes

## Contributing

For more information about contributing to this plugin, see [Contributing][1].

[1]: https://github.com/sensu/sensu-go/blob/master/CONTRIBUTING.md
[2]: https://github.com/sensu-community/sensu-plugin-sdk
[3]: https://github.com/sensu-plugins/community/blob/master/PLUGIN_STYLEGUIDE.md
[4]: https://github.com/sensu-community/mutator-plugin-template/blob/master/.github/workflows/release.yml
[5]: https://github.com/sensu-community/mutator-plugin-template/actions
[6]: https://docs.sensu.io/sensu-go/latest/reference/mutators/
[7]: https://github.com/sensu-community/mutator-plugin-template/blob/master/main.go
[8]: https://bonsai.sensu.io/
[9]: https://github.com/sensu-community/sensu-plugin-tool
[10]: https://docs.sensu.io/sensu-go/latest/reference/assets/
