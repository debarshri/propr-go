#### Propr

A "key/value" property server written in go.
Uses bolt db

##### To build and run
./build.sh
./server

##### Usage

`curl -X POST "http://localhost:1323/create/property1"`

`curl -X POST "http://localhost:1323/add/property1/key1/value1"`

`curl "http://localhost:1323/get/property1/key1"`

would give

`value`

`curl "http://localhost:1323/get/property1"`

would give

`key1=value1`

It also has a UI in localhost:1323


However,

### TODO
Create a cluster of boltdb databases with common rpc functions for get put delete create based on Bloomfiltering on keys.
Working on it, bitches.


