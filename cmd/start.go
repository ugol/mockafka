/*
Copyright © 2022 Ugo Landini <ugo.landini@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start mock Kafka server",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		n, err := cmd.Flags().GetInt("brokers")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Starting Mock Kafka with %d brokers\n", n)
			startMock(n)
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	startCmd.Flags().Int("brokers", 3, "number of brokers")
}

func startMock(n int) {
	// trying to get rid of 'no bootstrap.servers configured' librdkafka message
	//conf := make(map[string]kafka.ConfigValue)
	//conf["bootstrap.servers"] = "127.0.0.1:61832"
	//_, err = kafka.NewConsumer((*kafka.ConfigMap)(&conf))

	mock, err := kafka.NewMockCluster(n)

	if err != nil {
		fmt.Printf("Error creating mock cluster: %s", err)
		os.Exit(1)
	}
	fmt.Printf("bootstrap.servers=%s\n", mock.BootstrapServers())

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// Process messages
	run := true
	for run == true {
		select {
		case sig := <-sigchan:
			fmt.Printf("Kafka Mock server received signal: %v\n", sig)
			run = false
		default:
			// do nothing
		}
	}

	mock.Close()
}
