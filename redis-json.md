# REDIS JSON and REDIS SEARCH

```redis
FT.DROPINDEX echoTokenStoreIdx
FT.CREATE echoTokenStoreIdx ON JSON SCHEMA $.metadata.type AS type TEXT $.metadata.client_id AS client_id TEXT $.metadata.subject AS subject TEXT
FT.CREATE test              ON JSON SCHEMA $.metadata.type AS type TEXT $.metadata.client_id AS client_id TEXT $.metadata.subject AS subject TEXT

JSON.SET mykey . '{"metadata":{"type":"reference_token","client_id":"001","subject":"1111","expiration":"2022-04-18T09:55:18.4616032-07:00","issued_at":"0001-01-01T00:00:00Z"},"data":{"scope":["offline_access","a","b","c","users.read","invoices"],"aud":["b2b-client","users","invoices"],"client_id":"b2b-client","exp":1650123272,"iat":1650119672,"iss":"http://localhost:1323","jti":"c9dd7u1ld5lnsc78u3f0"}}'  


FT.SEARCH echoTokenStoreIdx "@subject:1111"
FT.SEARCH echoTokenStoreIdx "@type:reference_token"
FT.SEARCH echoTokenStoreIdx "@client_id:001"
FT.SEARCH echoTokenStoreIdx '@subject:1111, @client_id:001'
```
