## POSTの実行方法
```bash

$body = @{
    sample = "sample"
    sample = 3
    hint = @("黄色")
} | ConvertTo-Json
Invoke-RestMethod -Method POST -Uri http://localhost:1323/♯♯♯ -Body $body -ContentType "application/json;charset=UTF-8"#URL部分に該当すURLを入れる
# or
curl --location --request POST 'http://localhost:1323/###' --header 'Content-Type: application/json' --data-raw '{key: value}'

```
## GETの実行方法
```bash
'http://localhost:1323/###?###=###&$$$=$$$'

```

