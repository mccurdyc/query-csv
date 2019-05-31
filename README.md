# query-csv

**USE AT YOUR OWN RISK**

query-csv is a simple tool that uses a CSV to populate SQL query parameters.

## Getting Started

### Requirements
  + `Go`

### Installing

```bash
go get github.com/mccurdyc/query-csv
```

## Usage

Make sure that your connection string and query follow the requirements for the
chosen database driver. For example, with the PostgreSQL driver, the connection
string and query parameters should follow that outlined in the [Supported Database Drivers section](#postgres).

**BE CAREFUL**

### Example

```bash
make build
./bin/query \
  --file=<path/to/csv> \
  --conn_str="user=<username> password=<password> dbname=<dbname> host=<db-host> port=<db-port> sslmode=disable search_path=<search-path>" \
  --db_driver="postgres" \
  --query="UPDATE table_a SET b_id = b.id FROM (SELECT id FROM table_b WHERE field = \$2) AS b WHERE a.some_field = \$1"
```

## Supported Database Drivers

### postgres
  + library - `github.com/lib/pq`
  + [connection string document](https://godoc.org/github.com/lib/pq#hdr-Connection_String_Parameters)
  + [SQL query parameter markers](https://godoc.org/github.com/lib/pq#hdr-Queries)

### sqlite
  + library - `github.com/mattn/go-sqlite3`
  + [connection string document](https://github.com/mattn/go-sqlite3#connection-string)
  + `?`

## TODOs

+ help menu
+ handle CSV header line nicely.
+ make it clear somewhere that values in CSV must line up to query parameters.

## License
+ [GNU General Public License Version 3](./LICENSE)
