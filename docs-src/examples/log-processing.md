# Log Processing

## HashiCorp Vault Audit Log

HashiCorp Vault can write JSON-formatted audit logs.

`gq` can be used to process a stream of these logs to CSV (for example). All
extra whitespace in error messages is compressed to a single space and trimmed.

```sh
curl -sL https://github.com/hashicorp/vault-guides/raw/master/monitoring-troubleshooting/vault-audit.log \
| gq -l 'list .time .type .request.path .request.operation (regexReplaceAll "\\s+" (default "" .error | trim) " ") | join ","'
```

Which will output the following (trimmed to the last 5 lines for brevity):

```
2020-04-30T19:12:58.5648483Z,response,sys/mounts,read,1 error occurred: * permission denied
2020-04-30T19:32:00.7744629Z,request,auth/userpass/login/lab-user-4,update,
2020-04-30T19:32:00.9207237Z,response,auth/userpass/login/lab-user-4,update,
2020-04-30T19:35:23.1771431Z,request,auth/userpass/login/lab-user-5,update,
2020-04-30T19:35:23.2895529Z,response,auth/userpass/login/lab-user-5,update,
```
