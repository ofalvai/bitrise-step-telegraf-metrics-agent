# telegraf-metrics-agent

A Bitrise step that runs the [Telegraf](https://www.influxdata.com/time-series-platform/telegraf/) metrics collector agent in the background.

Metrics are written to `$BITRISE_DEPLOY_DIR/metrics.out` (by default) in InfluxDB line protocol format. Place this step at the beginning of a workflow, then a Deploy to Bitrise.io step at the end, so you can download `metrics.out` as an artifact.

By default, the following metrics are collected:

- [CPU](https://github.com/influxdata/telegraf/tree/master/plugins/inputs/cpu)
- [Memory](https://github.com/influxdata/telegraf/tree/master/plugins/inputs/mem)
- [Process stats](https://github.com/influxdata/telegraf/tree/master/plugins/inputs/procstat)
- [Disk IO](https://github.com/influxdata/telegraf/tree/master/plugins/inputs/diskio)
- [Network](https://github.com/influxdata/telegraf/blob/master/plugins/inputs/net/NET_README.md)

## How to use

Include as a custom step in your `bitrise.yml`:

``` yml
steps:
  - git::https://github.com/ofalvai/bitrise-step-telegraf-metrics-agent.git@master:
```

You need to add this manually to the `bitrise.yml` file, but after saving the file, you can use the workflow editor to tweak the step inputs.

## Inputs

- `telegraf_conf`: Contents of the `telegraf.conf` file ([docs](https://docs.influxdata.com/telegraf/v1.18/administration/configuration/)). See `step.yml` or the Bitrise workflow editor for the default config
