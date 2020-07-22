# Generating the client

```
java \
  -jar openapi-generator-cli.jar \
  generate \
  -p enumClassPrefix=true \
  -i https://raw.githubusercontent.com/XeroAPI/Xero-OpenAPI/master/accounting-yaml/xero_accounting.yaml \
  -g go \
  -o .
```
