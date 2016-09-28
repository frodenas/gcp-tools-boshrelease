# Google Cloud Platform Tools BOSH Release

This is a [BOSH](http://bosh.io/) release for [Google Cloud Platform](https://cloud.google.com/) Tools:

* A job acting as a Syslog endpoint to send platform logs to [StackDriver Logging](https://cloud.google.com/logging/)
* A job that forwards [Cloud Foundry Firehose](https://docs.cloudfoundry.org/loggregator/architecture.html#firehose) event data (including application logs) to [StackDriver Logging](https://cloud.google.com/logging/)
* A job (to be collocated with other jobs) to send host metrics to [StackDriver Monitoring](https://cloud.google.com/monitoring/)

## Disclaimer

** This BOSH release can only be deployed to [Google Cloud Platform](https://cloud.google.com/)**

This is NOT presently a production ready BOSH release. This is just a Proof of Concept. It is suitable for experimentation and may not become supported in the future.

## Acknowledgements

Many thanks to [Etourneau Gwenn](https://github.com/shinji62) for his fantastic [firehose-to-fluentd](https://github.com/shinji62/firehose-to-fluentd) utility.

## Usage

### Upload the BOSH release

To use this BOSH release, first upload it to your BOSH:

```
bosh target BOSH_HOST
bosh upload https://storage.googleapis.com/bosh-releases/gcp-tools-1.tgz
```

### Deploying StackDriver Logging

Create a deployment file (use the [gcp-tools.yml](https://github.com/frodenas/gcp-tools-release/blob/master/manifests/gcp-tools.yml) example manifest file as a starting point).

Using the previous created deployment manifest, now deploy it:

```
bosh deployment path/to/deployment.yml
bosh -n deploy
```

Once deployed:
* the `google_fluentd` will act as a Syslog endpoint and will forward logs to [StackDriver Logging](https://cloud.google.com/logging/)
* the `firehose_to_fluentd` will forward all [Cloud Foundry Firehose](https://docs.cloudfoundry.org/loggregator/architecture.html#firehose) event data (including application logs) to [StackDriver Logging](https://cloud.google.com/logging/)

If you want to send all your Cloud Foundry component's logs to [StackDriver Logging](https://cloud.google.com/logging/), configure your Cloud Foundry manifest adding (or updating):

```
properties:
  ...
  syslog_daemon_config:
    address: <google-fluentd job instance IP address>
    port: 514
    transport: udp
```

### Deploying StackDriver Monitoring

Add the `gcp-tools` release to the `release` section of your existing deployment manifest:

```
release:
   ...
  - name: gcp-tools
    version: "1"
```

Collocate the `stackdriver_agent` job template in all job instances:

```
jobs:
  - name: nats
    templates:
      - name: nats
        release: cf
      - name: metron_agent
        release: cf
      - name: stackdriver_agent
        release: gcp-tools
```

Once deployed, the `stackdriver_agent` on every instance will send host metrics to [StackDriver Monitoring](https://cloud.google.com/monitoring/).

## Contributing

In the spirit of [free software](http://www.fsf.org/licensing/essays/free-sw.html), **everyone** is encouraged to help improve this project.

Here are some ways *you* can contribute:

* by using alpha, beta, and prerelease versions
* by reporting bugs
* by suggesting new features
* by writing or editing documentation
* by writing specifications
* by writing code (**no patch is too small**: fix typos, add comments, clean up inconsistent whitespace)
* by refactoring code
* by closing [issues](https://github.com/frodenas/gcp-tools-release/issues)
* by reviewing patches

### Submitting an Issue

We use the [GitHub issue tracker](https://github.com/frodenas/gcp-tools-release/issues) to track bugs and features. Before submitting a bug report or feature request, check to make sure it hasn't already been submitted. You can indicate support for an existing issue by voting it up. When submitting a bug report, please include a [Gist](http://gist.github.com/) that includes a stack trace and any details that may be necessary to reproduce the bug,. Ideally, a bug report should include a pull request with failing specs.

### Submitting a Pull Request

1. Fork the project.
2. Create a topic branch.
3. Implement your feature or bug fix.
4. Commit and push your changes.
5. Submit a pull request.

## Copyright

Copyright (c) 2016 Ferran Rodenas. See [LICENSE](https://github.com/frodenas/gcp-tools-release/blob/master/LICENSE) for details.