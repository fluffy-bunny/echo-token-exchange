# Infrastructure

## Traefik

[echo-infrastructure](https://github.com/fluffy-bunny/echo-infrastructure)  

```bash
docker-compose -f .\docker\inmemory\docker-compose-token-exchange.yml up
```

### Request

```curl
curl --location --request POST 'http://echo-tx.docker.localhost/token' \
--header 'Authorization: Basic YjJiLWNsaWVudDpzZWNyZXQ=' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'grant_type=client_credentials' \
--data-urlencode 'scope=a b c users.read invoices'
```

### Response

```json
{
    "access_token": "eyJhbGciOiJFUzI1NiIsImtpZCI6IjBiMmNkMmU1NGM5MjRjZTg5ZjAxMGYyNDI4NjIzNjdkIiwidHlwIjoiSldUIn0.eyJhdWQiOlsiYjJiLWNsaWVudCIsInVzZXJzIiwiaW52b2ljZXMiXSwiY2xpZW50X2lkIjoiYjJiLWNsaWVudCIsImV4cCI6MTY1MDkyNTQ3MywiaWF0IjoxNjUwOTIxODczLCJpc3MiOiJodHRwOi8vZWNoby10eC5kb2NrZXIubG9jYWxob3N0LyIsImp0aSI6ImM5amgzNGIyZmVqOGs4dDU0cDlnIiwic2NvcGUiOlsiYSIsImIiLCJjIiwidXNlcnMucmVhZCIsImludm9pY2VzIl19.V2AL27xx2jmvRCv6wUujyocv1WX1gE8ABMiqF96hcjBM2VHJVbuYZsyDJa2EqFXoU94K3tCTSGxvRocDHivq8Q",
    "expires_in": 3600,
    "scope": "a b c users.read invoices",
    "token_type": "Bearer"
}
```
### JWT.IO

```json
{
  "aud": [
    "b2b-client",
    "users",
    "invoices"
  ],
  "client_id": "b2b-client",
  "exp": 1650925473,
  "iat": 1650921873,
  "iss": "http://echo-tx.docker.localhost/",
  "jti": "c9jh34b2fej8k8t54p9g",
  "scope": [
    "a",
    "b",
    "c",
    "users.read",
    "invoices"
  ]
}
```

