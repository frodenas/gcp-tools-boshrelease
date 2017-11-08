/*
 * Copyright 2017 Google Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"os"
	"time"

	"github.com/cloudfoundry-community/stackdriver-tools/src/stackdriver-nozzle/cloudfoundry"
	"github.com/cloudfoundry-community/stackdriver-tools/src/stackdriver-nozzle/heartbeat"

	"github.com/cloudfoundry-community/go-cfclient"
	"github.com/cloudfoundry/lager"
)

func main() {
	logger := lager.NewLogger("firehose-rate-script")
	logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))

	apiEndpoint := os.Getenv("FIREHOSE_ENDPOINT")
	username := os.Getenv("FIREHOSE_USERNAME")
	password := os.Getenv("FIREHOSE_PASSWORD")
	_, skipSSLValidation := os.LookupEnv("FIREHOSE_SKIP_SSL")

	cfConfig := &cfclient.Config{
		ApiAddress:        apiEndpoint,
		Username:          username,
		Password:          password,
		SkipSslValidation: skipSSLValidation}

	cfClient, err := cfclient.NewClient(cfConfig)
	if err != nil {
		logger.Fatal("NewClient", err)
	}
	client := cloudfoundry.NewFirehose(cfConfig, cfClient, "firehose-rate-script")

	messages, _ := client.Connect()

	period := time.Duration(1 * time.Second)
	heartbeater := heartbeat.NewTelemetry(logger, period, nil)
	heartbeater.Start()
	defer heartbeater.Stop()
	for _ = range messages {
		heartbeater.Increment("count")
	}
}
