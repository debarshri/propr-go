#### Propr

A "key/value" property server written in go.


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


