# 回答情報の取得
A社がB社の回答情報を取得する例を示します。  
識別子およびアクセストークンは適宜置き換えが必要ですのでご注意下さい。

## 1. 回答依頼情報の取得
Action (A社): 下記の```curl```コマンドを実行し、A社が回答依頼している依頼一覧を取得する

```
curl --location --request GET 'http://localhost:8080/api/v1/datatransport?dataTarget=tradeRequest' \
--header 'apiKey: Sample-APIKey1' \
--header 'Authorization: Bearer eyJhbGciOiJub25lIiwidHlwIjoiSldUIn0.eyJvcGVyYXRvcl9pZCI6ImIzOWU2MjQ4LWM4ODgtNTZjYS1kOWQwLTg5ZGUxYjFhZGM4ZSIsImVtYWlsIjoib2VtX2FAZXhhbXBsZS5jb20iLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsImF1dGhfdGltZSI6MTcxMzc3OTY3MCwidXNlcl9pZCI6IjUxZjk3MGE3LTFiYmUtNDAxNS1hNzI0LWIzZjUwYzljMWRjZCIsImZpcmViYXNlIjp7ImlkZW50aXRpZXMiOnsiZW1haWwiOlsib2VtX2FAZXhhbXBsZS5jb20iXX0sInNpZ25faW5fcHJvdmlkZXIiOiJwYXNzd29yZCJ9LCJpYXQiOjE3MTM3Nzk2NzAsImV4cCI6MTcxMzc4MzI3MCwiYXVkIjoibG9jYWwiLCJpc3MiOiJodHRwczovL3NlY3VyZXRva2VuLmdvb2dsZS5jb20vbG9jYWwiLCJzdWIiOiI1MWY5NzBhNy0xYmJlLTQwMTUtYTcyNC1iM2Y1MGM5YzFkY2QifQ.'
```

```json
[
    {
        "tradeId": "f475cb75-b3b8-4427-9e8d-376377f1c795",
        "downstreamOperatorId": "b39e6248-c888-56ca-d9d0-89de1b1adc8e",
        "upstreamOperatorId": "15572d1c-ec13-0d78-7f92-dd4278871373",
        "downstreamTraceId": "40b77952-2c89-49be-8ce9-7c64a15e0ae7",
        "upstreamTraceId": "2fb97052-250b-44de-acbb-1ba63e28af71"
    }
]
```

## 2. 依頼情報のステータスを確認
Action (A社): 下記の```curl```コマンドを実行し、A社が回答依頼している依頼一覧を取得する

- statusTarget:自社が依頼元のステータスのみを取得するQueryパラメータを指定します
- traceId:依頼している自社のtraceIdを指定します

```
curl --location --request GET 'http://localhost:8080/api/v1/datatransport?dataTarget=status&statusTarget=REQUEST&traceId=40b77952-2c89-49be-8ce9-7c64a15e0ae7' \
--header 'apiKey: Sample-APIKey1' \
--header 'Authorization: Bearer eyJhbGciOiJub25lIiwidHlwIjoiSldUIn0.eyJvcGVyYXRvcl9pZCI6ImIzOWU2MjQ4LWM4ODgtNTZjYS1kOWQwLTg5ZGUxYjFhZGM4ZSIsImVtYWlsIjoib2VtX2FAZXhhbXBsZS5jb20iLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsImF1dGhfdGltZSI6MTcxMzc3OTY3MCwidXNlcl9pZCI6IjUxZjk3MGE3LTFiYmUtNDAxNS1hNzI0LWIzZjUwYzljMWRjZCIsImZpcmViYXNlIjp7ImlkZW50aXRpZXMiOnsiZW1haWwiOlsib2VtX2FAZXhhbXBsZS5jb20iXX0sInNpZ25faW5fcHJvdmlkZXIiOiJwYXNzd29yZCJ9LCJpYXQiOjE3MTM3Nzk2NzAsImV4cCI6MTcxMzc4MzI3MCwiYXVkIjoibG9jYWwiLCJpc3MiOiJodHRwczovL3NlY3VyZXRva2VuLmdvb2dsZS5jb20vbG9jYWwiLCJzdWIiOiI1MWY5NzBhNy0xYmJlLTQwMTUtYTcyNC1iM2Y1MGM5YzFkY2QifQ.'
```

```json
[
    {
        "statusId": "92e42cc2-7768-4bc8-b956-356a40a7a506",
        "tradeId": "f475cb75-b3b8-4427-9e8d-376377f1c795",
        "requestStatus": {
            "cfpResponseStatus": "COMPLETED",
            "tradeTreeStatus": "TERMINATED",
            "completedCount": 1,
            "completedCountModifiedAt": "2024-11-15T06:33:10Z",
            "tradesCount": 1,
            "tradesCountModifiedAt": "2024-11-15T04:49:44Z"
        },
        "message": "来月中にご回答をお願いします。",
        "replyMessage": null,
        "requestType": "CFP",
        "responseDueDate": "2024-12-31"
    }
]
```
※ これらはダミー実装のため、本番の返却値とは一部値が異なる可能性がございます。

# 完成品のCFP情報を算出する

## 1. 製品にCFP情報を登録

Action (A社): 下記の```curl```コマンドを実行し、A社の製品にCFP情報を登録する

前提条件として、完成品のCFP情報の算出はアプリケーションで実施されます。  
登録する完成品のCFP情報は下記です。
- 前処理自社由来排出量：3.0(kgCO2e/kilogram)
- 製造自社由来排出量：20.0(kgCO2e/kilogram)
- 前処理部品由来排出量：0(kgCO2e/kilogram)
- 製造部品由来排出量：0(kgCO2e/kilogram)
- 前処理工程DQR値
  - TeR:1
  - GeR:2
  - TiR:3
- 製造工程DQR値
  - TeR:2
  - GeR:3
  - TiR:4


```
curl --location --request PUT 'http://localhost:8080/api/v1/datatransport?dataTarget=cfp' \
--header 'apiKey: Sample-APIKey1' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJub25lIiwidHlwIjoiSldUIn0.eyJvcGVyYXRvcl9pZCI6ImIzOWU2MjQ4LWM4ODgtNTZjYS1kOWQwLTg5ZGUxYjFhZGM4ZSIsImVtYWlsIjoib2VtX2FAZXhhbXBsZS5jb20iLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsImF1dGhfdGltZSI6MTcxMzc4MDk3MywidXNlcl9pZCI6IjUxZjk3MGE3LTFiYmUtNDAxNS1hNzI0LWIzZjUwYzljMWRjZCIsImZpcmViYXNlIjp7ImlkZW50aXRpZXMiOnsiZW1haWwiOlsib2VtX2FAZXhhbXBsZS5jb20iXX0sInNpZ25faW5fcHJvdmlkZXIiOiJwYXNzd29yZCJ9LCJpYXQiOjE3MTM3ODA5NzMsImV4cCI6MTcxMzc4NDU3MywiYXVkIjoibG9jYWwiLCJpc3MiOiJodHRwczovL3NlY3VyZXRva2VuLmdvb2dsZS5jb20vbG9jYWwiLCJzdWIiOiI1MWY5NzBhNy0xYmJlLTQwMTUtYTcyNC1iM2Y1MGM5YzFkY2QifQ.' \
--data '[
    {
        "cfpId": null,
        "traceId": "40b77952-2c89-49be-8ce9-7c64a15e0ae7",
        "ghgEmission": 3.0,
        "ghgDeclaredUnit": "kgCO2e/kilogram",
        "cfpType": "preProduction",
        "dqrType": "preProcessing",
        "dqrValue": {
            "TeR": 1,
            "GeR": 2,
            "TiR": 3
        }
    },
    {
        "cfpId": null,
        "traceId": "40b77952-2c89-49be-8ce9-7c64a15e0ae7",
        "ghgEmission": 20.0,
        "ghgDeclaredUnit": "kgCO2e/kilogram",
        "cfpType": "mainProduction",
        "dqrType": "mainProcessing",
        "dqrValue": {
            "TeR": 2,
            "GeR": 3,
            "TiR": 4
        }
    },
    {
        "cfpId": null,
        "traceId": "40b77952-2c89-49be-8ce9-7c64a15e0ae7",
        "ghgEmission": 0,
        "ghgDeclaredUnit": "kgCO2e/kilogram",
        "cfpType": "preComponent",
        "dqrType": "preProcessing",
        "dqrValue": {
            "TeR": 1,
            "GeR": 2,
            "TiR": 3
        }
    },
    {
        "cfpId": null,
        "traceId": "40b77952-2c89-49be-8ce9-7c64a15e0ae7",
        "ghgEmission": 0,
        "ghgDeclaredUnit": "kgCO2e/kilogram",
        "cfpType": "mainComponent",
        "dqrType": "mainProcessing",
        "dqrValue": {
            "TeR": 2,
            "GeR": 3,
            "TiR": 4
        }
    }
]'
```

```json
[
    {
        "cfpId": "ec52116d-78bb-427b-82ee-033bb7a89b5d",
        "traceId": "40b77952-2c89-49be-8ce9-7c64a15e0ae7",
        "ghgEmission": 3,
        "ghgDeclaredUnit": "kgCO2e/kilogram",
        "cfpType": "preProduction",
        "dqrType": "preProcessing",
        "dqrValue": {
            "TeR": 1,
            "GeR": 2,
            "TiR": 3
        }
    },
    {
        "cfpId": "ec52116d-78bb-427b-82ee-033bb7a89b5d",
        "traceId": "40b77952-2c89-49be-8ce9-7c64a15e0ae7",
        "ghgEmission": 20,
        "ghgDeclaredUnit": "kgCO2e/kilogram",
        "cfpType": "mainProduction",
        "dqrType": "mainProcessing",
        "dqrValue": {
            "TeR": 2,
            "GeR": 3,
            "TiR": 4
        }
    },
    {
        "cfpId": "ec52116d-78bb-427b-82ee-033bb7a89b5d",
        "traceId": "40b77952-2c89-49be-8ce9-7c64a15e0ae7",
        "ghgEmission": 0,
        "ghgDeclaredUnit": "kgCO2e/kilogram",
        "cfpType": "preComponent",
        "dqrType": "preProcessing",
        "dqrValue": {
            "TeR": 1,
            "GeR": 2,
            "TiR": 3
        }
    },
    {
        "cfpId": "ec52116d-78bb-427b-82ee-033bb7a89b5d",
        "traceId": "40b77952-2c89-49be-8ce9-7c64a15e0ae7",
        "ghgEmission": 0,
        "ghgDeclaredUnit": "kgCO2e/kilogram",
        "cfpType": "mainComponent",
        "dqrType": "mainProcessing",
        "dqrValue": {
            "TeR": 2,
            "GeR": 3,
            "TiR": 4
        }
    }
]
```

## 2. 登録したCFPの値を取得

Action (A社): 下記の```curl```コマンドを実行し、製品に登録したCFP値を取得する

```
curl --location --request GET 'http://localhost:8080/api/v1/datatransport?dataTarget=cfp&traceIds=40b77952-2c89-49be-8ce9-7c64a15e0ae7' \
--header 'apiKey: Sample-APIKey1' \
--header 'Authorization: Bearer eyJhbGciOiJub25lIiwidHlwIjoiSldUIn0.eyJvcGVyYXRvcl9pZCI6ImIzOWU2MjQ4LWM4ODgtNTZjYS1kOWQwLTg5ZGUxYjFhZGM4ZSIsImVtYWlsIjoib2VtX2FAZXhhbXBsZS5jb20iLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsImF1dGhfdGltZSI6MTcxNDAyMDg0MSwidXNlcl9pZCI6IjA4ZTgyMWRmLTBiZGUtNDg0MC1iZmE0LTczNTc1YmVlOGU1NSIsImZpcmViYXNlIjp7ImlkZW50aXRpZXMiOnsiZW1haWwiOlsib2VtX2FAZXhhbXBsZS5jb20iXX0sInNpZ25faW5fcHJvdmlkZXIiOiJwYXNzd29yZCJ9LCJpYXQiOjE3MTQwMjA4NDEsImV4cCI6MTcxNDAyNDQ0MSwiYXVkIjoibG9jYWwiLCJpc3MiOiJodHRwczovL3NlY3VyZXRva2VuLmdvb2dsZS5jb20vbG9jYWwiLCJzdWIiOiIwOGU4MjFkZi0wYmRlLTQ4NDAtYmZhNC03MzU3NWJlZThlNTUifQ.'
```

```json
[
    {
        "cfpId": null,
        "traceId": "40b77952-2c89-49be-8ce9-7c64a15e0ae7",
        "ghgEmission": 1.5,
        "ghgDeclaredUnit": "kgCO2e/kilogram",
        "cfpType": "preProductionResponse",
        "dqrType": "preProcessingResponse",
        "dqrValue": {
            "TeR": 2.1,
            "GeR": 0,
            "TiR": null
        }
    },
    {
        "cfpId": null,
        "traceId": "40b77952-2c89-49be-8ce9-7c64a15e0ae7",
        "ghgEmission": 10,
        "ghgDeclaredUnit": "kgCO2e/kilogram",
        "cfpType": "mainProductionResponse",
        "dqrType": "mainProcessingResponse",
        "dqrValue": {
            "TeR": 2.1,
            "GeR": 0,
            "TiR": null
        }
    }
]
```

※ これらはダミー実装のため、本番の返却値とは一部値が異なる可能性がございます。