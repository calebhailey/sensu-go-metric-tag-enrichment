# Sensu Go Metric Tag Enrichment

- [Overview](#overview)
  - [Background](#background)
- [Example Usage](#example-usage)
- [Configuration](#configuration)
  - [Asset Definition](#asset-definition)
  - [Mutator Definition](#mutator-definition)
- [About](#about)
- [Roadmap](#roadmap)

## Overview

Sensu Metric Tag Enrichment is a Sensu Go backend integration for enriching
metric points with [Entity metadata][entity-metadata] (e.g. entity name,
labels), and the originating Check name by adding them as
[Metric Tags][metric-tags]. Sensu Metric Tag Enrichment is useful for adding
context to generic metrics collected from plugins that generate metrics in
formats that don't support tags (e.g. Nagios Perfdata).

This integration provides the following features:

- Event Mutator

[entity-metadata]: https://docs.sensu.io/sensu-go/latest/reference/entities/#metadata-attributes
[metric-tags]: https://docs.sensu.io/sensu-go/latest/reference/events/#metrics

### Background

Sensu Go has a built-in feature called ["metric extraction"][metric-extraction]
that parses check output for metric data, and extract it into a generic format
for processing by a Sensu Handler. The benefit of this approach is that Sensu
Handlers can process metrics collected in any format without the complexity of
having to directly support them &ndash; Handlers can simply consume metrics in
the generic Sensu format and convert them to the desired/required format for the
intended backend storage system. There are also certain performance and
scalability benefits derived from performing metric extraction at the "edge"
(i.e. by the Sensu Agents) as opposed to performing this conversion on the
backend nodes.  

Metric extraction is enabled by adding the `output_metric_format` attribute to a
[Sensu Go CheckConfig][checkconfig] spec; supported formats include
`nagios_perfdata`, `graphite_plaintext`, `influxdb_line`, and `opentsdb_line`.  
However, not all metric sources are created equally. For example, Nagios plugins
that emit Nagios Performance Data ("PerfData") do not support tags.
Historically, it was popular to adopt a dot-notation based metric naming
convention for generic metrics (e.g. a generically named `response_time` metric
could become `server01.http_service.response_time`). Now that most modern time
series databases support custom tags, it is more practical to provide context
about collected metrics as tags (e.g. `host: server01`, and
`check: http_service`).

This integration provides some basic support for enriching extracted metrics
with basic Entity metadata and the originating Sensu Check name.

[metric-extraction]: https://docs.sensu.io/sensu-go/5.1/guides/extract-metrics-with-checks/

## Configuration

1. Create a Sensu Asset named `metric-tag-enrichment` (see example
  [Asset Definition](#asset-definition)).

2. Create a Sensu Mutator named `metric-tag-enrichment` (see example
  [Mutator Definition](#mutator-definition)).

3. Configure a Sensu Handler to use the Metric Tag Enrichment mutator.

   Add the `spec.mutator` attribute to an existing Sensu Handler resource using
   `sensuctl`s interactive update mode:

   ```
   $ sensuctl handler update influxdb
   
   ? Environment variables:
   ? Filters:
   ? Mutator: metric-tag-enrichment
   ? Timeout: 5
   ? Type: pipe
   ? Runtime Assets: sensu-influxdb-handler
   ? Command: sensu-influxdb-handler -a http://influxdb:8086 -d 'sensu' -u 'admin' -p 'admin'
   ```

### Asset Definition

```yaml
---
type: Asset
api_version: core/v2
metadata:
  namespace: default
  name: metric-tag-enrichment
  labels:
    sensu.io/integration/type: backend
    sensu.io/integration/providers: mutator
spec:
  url: https://github.com/calebhailey/sensu-go-metric-tag-enrichment
  sha512: #
```

### Mutator Definition

```yaml
---
type: Mutator
api_version: core/v2
metadata:
  namespace: default
  name: metric-tag-enrichment
spec:
  command: metric-tag-enrichment
  timeout: 5
  runtime_assets:
  - metric-tag-enrichment
```

## About

This project was an experiment for me to learn about Golang, so I'm approaching
things very naively. If you see room for improvement or any obvious mistakes,
please open an issue to help me learn! :) Thanks in advance for your feedback!

The motivation for this project was to attempt implementing a solution for an
outstanding [Sensu Go feature request][2160] to provide support for enriching
extracted metrics with tags. Off to a good start!

[2160]: https://github.com/sensu/sensu-go/issues/2160

## Roadmap

- Add support for toggling the behavior of the mutator; e.g.:
  - `--entity-name` toggle addition of an `"entity"` tag (source:
    `event.Entity.Name`).
  - `--entity-labels` toggle addition of Entity Labels as tags (source:
    `event.Entity.Labels`).
  - `--check-name` toggle addition of a `"check"` tag (source:
    `event.Check.Name`).
  - `--check-labels` toggle addition of Check Labels as tags (source:
    `event.Check.Labels`).
- Add support for customizing the behavior of the Entity Label and Check Label
  toggles; e.g.:
  - `--entity-labels region,datacenter` filter which Entity Labels are added as
    tags.
  - `--check-labels application_id,team` filter which Check Labels are added as
    tags.
