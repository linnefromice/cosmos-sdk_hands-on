# cosmos-sdk_hands-on

```bash
PATH=$PATH:$(go env GOPATH)/bin ignite chain serve
MODULE_BIN=$(go env GOPATH)/bin/handsond
$MODULE_BIN q bank balances $($MODULE_BIN keys show alice -a) --output json | jq
$MODULE_BIN q bank balances $($MODULE_BIN keys show bob -a) --output json | jq
```
