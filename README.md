# GoSocialNetwork
Project built based on Udemy course "Backend Engineering with Go".

## Benchmarking the usage of redis

For example, for the users endpoint, just run the following command, enabling and then disabling redis through `REDIS_ENABLED` environment variable.

```
npx autocannon http://192.168.0.5:8080/v1/users/107 --connections 50 --duration 5 -H "Authorization: Bearer <token>"
```