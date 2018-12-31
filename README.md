# Sensu Go Metric Tag Enrichment

- [Overview](#overview)
- [Configuration](#configuration)
  - [Asset Definition](#asset-definition)
  - [Mutator Definition](#mutator-definition)
- [About](#about)
- [Roadmap](#roadmap)

## Overview

Sensu Metric Tag Enrichment is a Sensu Go backend integration for enriching
metric points with [Entity metadata][entity-metadata] (e.g. entity name, labels,
etc), by adding them as [Metric Tags][metric-tags]. Sensu Metric Tag Enrichment
is useful for adding context to metrics collected from plugins that generate
generic metrics in formats that don't support tags (e.g. Nagios Perfdata).

This integration provides the following features:

- [Event Mutator](#metrics)

[entity-metadata]: https://docs.sensu.io/sensu-go/latest/reference/entities/#metadata-attributes
[metric-tags]: https://docs.sensu.io/sensu-go/latest/reference/events/#metrics

## Configuration



### Asset Definition

```yaml
---
type: Asset
api_version: core/v2
metadata:
  namespace: default
  name: metric-tag-enrichment
  labels:
    sensu.io.integration.type: backend
    sensu.io.integration.class: community
    sensu.io.integration.providers: mutator
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

## Roamdap

- Add support for configuring the behavior of the mutator (e.g. only add an
  `"entity"` tag, or only add some subset of entity Labels as tags, or
  optionally add Check Labels as tags; etc)
