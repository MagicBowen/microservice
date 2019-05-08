## elasticsearch

### cmd

```sh
curl -u elastic:changeme -XGET 'http://localhost:9200/_cluster/stats?pretty'  
curl 'localhost:9200/_cat/nodes?v' 
curl 'localhost:9200/_cat/health?v'

curl 'localhost:9200/_cat/indices?v'
curl -XPUT 'localhost:9200/customer?pretty'

curl -XPUT 'localhost:9200/customer/external/1?pretty' -d '
{
　  "name": "John Doe"
}'

curl -XGET 'localhost:9200/customer/external/1?pretty'

curl -XDELETE 'localhost:9200/customer?pretty'

curl -XPOST 'localhost:9200/customer/external/1/_update?pretty' -d '
{
 "doc": { "name": "Jane Doe", "age": 20 }
}'

curl -XPOST 'localhost:9200/bank/_search?pretty' -d '
{
 "query": { "match_all": {} }
}'

```