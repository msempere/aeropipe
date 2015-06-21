# aeropipe
Treat Aerospike Large Lists unix-like pipelines

## Installation:
```
make buid
make install
```

## Setup
Aeropipe needs the following environment vars in order to connect to Aerospike:
 - 			AEROSPIKE_HOST (Database host, default: 127.0.0.1)
 - 			AEROSPIKE_PORT (Database port, default: 3000)
 - 			AEROSPIKE_AEROPIPE_NAMESPACE (Namespace configuration for Aerospike. [More info] (https://www.aerospike.com/docs/operations/configure/namespace/), default: aeropipe_ns)
 - 			AEROSPIKE_AEROPIPE_LIST (List where the data will be placed/located, default: aeropipe_list)

Example: (~/.basrc)
```
export AEROSPIKE_HOST="127.0.0.1"
export AEROSPIKE_PORT=3000
export AEROSPIKE_AEROPIPE_NAMESPACE="ssd_namespace"
export AEROSPIKE_AEROPIPE_LIST="my_awesome_list"
```

## Usage examples:

Pipe Apache access logs to Aerospike:
```
cat /var/log/apache2/access.log | aeropipe
```

Pipe Aerospike content to file:
```
aeropipe > /tmp/apache2/access.log
```

## License
Distributed under MIT license. See `LICENSE` for more information.
