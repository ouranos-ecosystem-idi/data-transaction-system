# CFP情報の伝達
B社がA社にCFP情報の伝達をする例を示します。  
識別子およびアクセストークンは適宜置き換えが必要ですのでご注意下さい。

## 1. 製品にCFP情報を登録
Action (B社): 下記の```curl```コマンドを実行し、B社の製品にCFP情報を登録する

前提条件として、完成品のCFP情報の算出はアプリケーションで実施されます。  
登録する完成品のCFP情報は下記です。
- 前処理自社由来排出量：1.5(kgCO2e/kilogram)
- 製造自社由来排出量：10.0(kgCO2e/kilogram)
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
--header 'Authorization: Bearer eyJhbGciOiJub25lIiwidHlwIjoiSldUIn0.eyJvcGVyYXRvcl9pZCI6IjE1NTcyZDFjLWVjMTMtMGQ3OC03ZjkyLWRkNDI3ODg3MTM3MyIsImVtYWlsIjoic3VwcGxpZXJfYkBleGFtcGxlLmNvbSIsImVtYWlsX3ZlcmlmaWVkIjpmYWxzZSwiYXV0aF90aW1lIjoxNzEzNzc2ODUxLCJ1c2VyX2lkIjoiMzBhNTE1OGItN2E0NS00ODZjLTgxMmMtNTllZDBmM2RjYjA1IiwiZmlyZWJhc2UiOnsiaWRlbnRpdGllcyI6eyJlbWFpbCI6WyJzdXBwbGllcl9iQGV4YW1wbGUuY29tIl19LCJzaWduX2luX3Byb3ZpZGVyIjoicGFzc3dvcmQifSwiaWF0IjoxNzEzNzc2ODUxLCJleHAiOjE3MTM3ODA0NTEsImF1ZCI6ImxvY2FsIiwiaXNzIjoiaHR0cHM6Ly9zZWN1cmV0b2tlbi5nb29nbGUuY29tL2xvY2FsIiwic3ViIjoiMzBhNTE1OGItN2E0NS00ODZjLTgxMmMtNTllZDBmM2RjYjA1In0.' \
--data '[
  {
    "cfpId": null,
    "traceId": "2fb97052-250b-44de-acbb-1ba63e28af71",
    "ghgEmission": 1.5,
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
    "traceId": "2fb97052-250b-44de-acbb-1ba63e28af71",
    "ghgEmission": 10.0,
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
    "traceId": "2fb97052-250b-44de-acbb-1ba63e28af71",
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
    "traceId": "2fb97052-250b-44de-acbb-1ba63e28af71",
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
        "cfpId": "6191188d-6e67-4449-8bd3-e121b8f23a65",
        "traceId": "2fb97052-250b-44de-acbb-1ba63e28af71",
        "ghgEmission": 1.5,
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
        "cfpId": "6191188d-6e67-4449-8bd3-e121b8f23a65",
        "traceId": "2fb97052-250b-44de-acbb-1ba63e28af71",
        "ghgEmission": 10,
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
        "cfpId": "6191188d-6e67-4449-8bd3-e121b8f23a65",
        "traceId": "2fb97052-250b-44de-acbb-1ba63e28af71",
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
        "cfpId": "6191188d-6e67-4449-8bd3-e121b8f23a65",
        "traceId": "2fb97052-250b-44de-acbb-1ba63e28af71",
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

## 2. A社にCFP情報およびDQR情報の開示設定
Action (B社): 製品に登録したCFPおよびDQR情報をA社に開示する

本作業はダミー実装では自動的に許可されるため、実施は不要です。（本番環境では個別に開示の設定が必要になります）

※ これらはダミー実装のため、本番の返却値とは一部値が異なる可能性がございます。