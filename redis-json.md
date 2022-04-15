# REDIS JSON and REDIS SEARCH

```redis
FT.DROPINDEX mmIdx
FT.CREATE mmIdx ON JSON SCHEMA $.meta.type AS type TEXT $.meta.client AS client TEXT $.meta.subject AS subject TEXT
JSON.SET mm . '{"meta":{"type":"refresh","subject":"herb","client":"001"},"data":{"name":"Paul John","email":"paul.john@example.com","age":"42","country":"London"}}'  


FT.SEARCH mmIdx "@subject:herb"
FT.SEARCH mmIdx "@type:refresh"
FT.SEARCH mmIdx "@client:a"
FT.SEARCH mmIdx '@subject:herb, @client:002'
```
