package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	"gopkg.in/mgo.v2/bson"

	"github.com/aerospike/aerospike-client-go"
	"github.com/codegangsta/cli"
	"github.com/msempere/aeropipe/util"
)

func read(client *aerospike.Client, namespace, set string) {
	record, err := client.ScanAll(nil, namespace, set)
	util.PanicOnError("Error scanning database", err)

	keys := make([]string, 0)
	values := make(map[string]string)

	for res := range record.Results() {
		if res.Err != nil {
			util.PanicOnError("Error processing results", err)
		} else {
			for _, value := range res.Record.Bins {
				key := res.Record.Key.Value().String()
				util.PanicOnError("Error processing key", err)

				keys = append(keys, key)
				values[key] = value.(string)
			}
		}
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(values[k])
	}
}

func store(client *aerospike.Client, policy *aerospike.WritePolicy, namespace, list string) {
	s := bufio.NewScanner(os.Stdin)
	var i uint64 = 0

	for s.Scan() {
		binkey := bson.NewObjectId()
		key, err := aerospike.NewKey(namespace, list, binkey)
		util.PanicOnError("Error generating new key", err)

		err = client.PutBins(policy, key, aerospike.NewBin("", s.Text()))
		util.PanicOnError("Error while inserting bins", err)
		i += 1
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "aeropipe"
	app.Usage = "Treat Aerospike Large Lists unix-like pipelines"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "host",
			Value:  "127.0.0.1",
			Usage:  "Aerospike host",
			EnvVar: "AEROSPIKE_HOST",
		},
		cli.IntFlag{
			Name:   "port",
			Value:  3000,
			Usage:  "Aerospike port",
			EnvVar: "AEROSPIKE_PORT",
		},
		cli.StringFlag{
			Name:   "namespace",
			Value:  "aeropipe_ns",
			Usage:  "Aeropipe namespace",
			EnvVar: "AEROSPIKE_AEROPIPE_NAMESPACE",
		},
		cli.StringFlag{
			Name:   "list",
			Value:  "aeropipe_list",
			Usage:  "Aeropipe list",
			EnvVar: "AEROSPIKE_AEROPIPE_LIST",
		},
		cli.StringFlag{
			Name:   "key",
			Value:  "aeropipe_key",
			Usage:  "Aeropipe key",
			EnvVar: "AEROSPIKE_AEROPIPE_KEY",
		},
	}

	app.Action = func(c *cli.Context) {

		client, err := aerospike.NewClient(c.String("host"), c.Int("port"))
		util.PanicOnError("Error generating aerospike client", err)
		if !client.IsConnected() {
			util.PanicOnError("Error, not connected to database", err)
		}
		defer client.Close()

		policy := aerospike.NewWritePolicy(0, 0)
		policy.SendKey = true

		if util.IsTTY() {
			read(client, c.String("namespace"), c.String("set"))
		} else {
			store(client, policy, c.String("namespace"), c.String("list"))
		}
	}

	app.Run(os.Args)
}
