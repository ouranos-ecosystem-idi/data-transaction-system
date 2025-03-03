# 部品登録
A社が部品登録する例を示します。  
識別子およびアクセストークンは適宜置き換えが必要ですのでご注意下さい。

## 1. 事業所の登録
Action (A社): 下記の```curl```コマンドを実行し、A社に事業所(A工場)を作成します。

事業者Aには、自社の製品を製造している事業所の登録がないため、新規に作成する必要があります。

```
curl --location --request PUT 'http://localhost:8081/api/v1/authInfo?dataTarget=plant' \
--header 'apiKey: Sample-APIKey1' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJub25lIiwidHlwIjoiSldUIn0.eyJvcGVyYXRvcl9pZCI6ImIzOWU2MjQ4LWM4ODgtNTZjYS1kOWQwLTg5ZGUxYjFhZGM4ZSIsImVtYWlsIjoib2VtX2FAZXhhbXBsZS5jb20iLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsImF1dGhfdGltZSI6MTcxMzc3MDc5MiwidXNlcl9pZCI6IjUxZjk3MGE3LTFiYmUtNDAxNS1hNzI0LWIzZjUwYzljMWRjZCIsImZpcmViYXNlIjp7ImlkZW50aXRpZXMiOnsiZW1haWwiOlsib2VtX2FAZXhhbXBsZS5jb20iXX0sInNpZ25faW5fcHJvdmlkZXIiOiJwYXNzd29yZCJ9LCJpYXQiOjE3MTM3NzA3OTIsImV4cCI6MTcxMzc3NDM5MiwiYXVkIjoibG9jYWwiLCJpc3MiOiJodHRwczovL3NlY3VyZXRva2VuLmdvb2dsZS5jb20vbG9jYWwiLCJzdWIiOiI1MWY5NzBhNy0xYmJlLTQwMTUtYTcyNC1iM2Y1MGM5YzFkY2QifQ.' \
--data '{
  "openPlantId": "1234567890123012345",
  "operatorId": "b39e6248-c888-56ca-d9d0-89de1b1adc8e",
  "plantAddress": "xx県xx市xxxx町1-1-1234",
  "plantId": null,
  "plantName": "A工場",
  "plantAttribute": {}
}'
```

```json
{
    "plantId": "0cc8b4be-c727-4411-b478-2c874fbc6c25",
    "operatorId": "b39e6248-c888-56ca-d9d0-89de1b1adc8e",
    "plantName": "A工場",
    "plantAddress": "xx県xx市xxxx町1-1-1234",
    "openPlantId": "1234567890123012345",
    "plantAttribute": {
        "globalPlantId": null
    }
}
```

## 2. 親部品情報の作成
Action (A社): 下記の```curl```コマンドを実行し、A社の部品Aを登録します。  
登録する製品の情報は下記です。  
- 部品項目:部品A
- 補助項目:modelA
- 事業所識別子（内部）:{事業所Aで登録した識別子}
- 事業者識別子（内部）:b39e6248-c888-56ca-d9d0-89de1b1adc8e


```
curl --location --request PUT 'http://localhost:8080/api/v1/datatransport?dataTarget=parts' \
--header 'apiKey: Sample-APIKey1' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJub25lIiwidHlwIjoiSldUIn0.eyJvcGVyYXRvcl9pZCI6ImIzOWU2MjQ4LWM4ODgtNTZjYS1kOWQwLTg5ZGUxYjFhZGM4ZSIsImVtYWlsIjoib2VtX2FAZXhhbXBsZS5jb20iLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsImF1dGhfdGltZSI6MTcxMzc3MjYzMiwidXNlcl9pZCI6IjUxZjk3MGE3LTFiYmUtNDAxNS1hNzI0LWIzZjUwYzljMWRjZCIsImZpcmViYXNlIjp7ImlkZW50aXRpZXMiOnsiZW1haWwiOlsib2VtX2FAZXhhbXBsZS5jb20iXX0sInNpZ25faW5fcHJvdmlkZXIiOiJwYXNzd29yZCJ9LCJpYXQiOjE3MTM3NzI2MzIsImV4cCI6MTcxMzc3NjIzMiwiYXVkIjoibG9jYWwiLCJpc3MiOiJodHRwczovL3NlY3VyZXRva2VuLmdvb2dsZS5jb20vbG9jYWwiLCJzdWIiOiI1MWY5NzBhNy0xYmJlLTQwMTUtYTcyNC1iM2Y1MGM5YzFkY2QifQ.' \
--data '{
  "amountRequired": null,
  "amountRequiredUnit": "kilogram",
  "operatorId": "b39e6248-c888-56ca-d9d0-89de1b1adc8e",
  "partsName": "部品A",
  "plantId": "0cc8b4be-c727-4411-b478-2c874fbc6c25",
  "supportPartsName": "modelA",
  "terminatedFlag": false,
  "traceId": null,
  "partsLabelName": "PartsA",
  "partsAddInfo1": "Ver3.0",
  "partsAddInfo2": "2024-12-01-2024-12-31",
  "partsAddInfo3": "任意の情報が入ります"
}'
```

```json
{
    "traceId": "8fc6aa29-5f4f-476e-85e3-2d1b54715891",
    "operatorId": "b39e6248-c888-56ca-d9d0-89de1b1adc8e",
    "plantId": "0cc8b4be-c727-4411-b478-2c874fbc6c25",
    "partsName": "部品A",
    "supportPartsName": "modelA",
    "terminatedFlag": false,
    "amountRequired": null,
    "amountRequiredUnit": "kilogram",
    "partsLabelName": "PartsA",
    "partsAddInfo1": "Ver3.0",
    "partsAddInfo2": "2024-12-01-2024-12-31",
    "partsAddInfo3": "任意の情報が入ります"
}
```

製品登録が成功するとトレース識別子が付番されます。

## 3. 部品構成情報の登録
Action (A社): 下記のcurlコマンドを実行し、作成した自社の部品に対して子部品を登録します。

登録する子部品の製品の情報は下記です。  
- 部品項目:部品A1
- 補助項目:modelA-1
- 事業所識別子（内部）:{事業所Aで登録した識別子}
- 事業者識別子（内部）:b39e6248-c888-56ca-d9d0-89de1b1adc8e
```
curl --location --request PUT 'http://localhost:8080/api/v1/datatransport?dataTarget=partsStructure' \
--header 'apiKey: Sample-APIKey1' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJub25lIiwidHlwIjoiSldUIn0.eyJvcGVyYXRvcl9pZCI6ImIzOWU2MjQ4LWM4ODgtNTZjYS1kOWQwLTg5ZGUxYjFhZGM4ZSIsImVtYWlsIjoib2VtX2FAZXhhbXBsZS5jb20iLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsImF1dGhfdGltZSI6MTcxMzc3MjYzMiwidXNlcl9pZCI6IjUxZjk3MGE3LTFiYmUtNDAxNS1hNzI0LWIzZjUwYzljMWRjZCIsImZpcmViYXNlIjp7ImlkZW50aXRpZXMiOnsiZW1haWwiOlsib2VtX2FAZXhhbXBsZS5jb20iXX0sInNpZ25faW5fcHJvdmlkZXIiOiJwYXNzd29yZCJ9LCJpYXQiOjE3MTM3NzI2MzIsImV4cCI6MTcxMzc3NjIzMiwiYXVkIjoibG9jYWwiLCJpc3MiOiJodHRwczovL3NlY3VyZXRva2VuLmdvb2dsZS5jb20vbG9jYWwiLCJzdWIiOiI1MWY5NzBhNy0xYmJlLTQwMTUtYTcyNC1iM2Y1MGM5YzFkY2QifQ.' \
--data '{
  "parentPartsModel": {
    "amountRequired": null,
    "amountRequiredUnit": "kilogram",
    "operatorId": "b39e6248-c888-56ca-d9d0-89de1b1adc8e",
    "partsName": "部品A",
    "plantId": "0cc8b4be-c727-4411-b478-2c874fbc6c25",
    "supportPartsName": "modelA",
    "terminatedFlag": false,
    "traceId": "8fc6aa29-5f4f-476e-85e3-2d1b54715891",
    "partsLabelName": "PartsA",
    "partsAddInfo1": "Ver3.0",
    "partsAddInfo2": "2024-12-01-2024-12-31",
    "partsAddInfo3": "任意の情報が入ります"
  },
  "childrenPartsModel": [
    {
      "amountRequired": 5,
      "amountRequiredUnit": "kilogram",
      "operatorId": "b39e6248-c888-56ca-d9d0-89de1b1adc8e",
      "partsName": "部品A1",
      "plantId": "0cc8b4be-c727-4411-b478-2c874fbc6c25",
      "supportPartsName": "modelA-1",
      "terminatedFlag": false,
      "traceId": null,
      "partsLabelName": "PartsA1",
      "partsAddInfo1": "Ver3.0",
      "partsAddInfo2": "2024-12-01-2024-12-31",
      "partsAddInfo3": "任意の情報が入ります"
    }
  ]
}'
```

```json
{
    "parentPartsModel": {
        "traceId": "8fc6aa29-5f4f-476e-85e3-2d1b54715891",
        "operatorId": "b39e6248-c888-56ca-d9d0-89de1b1adc8e",
        "plantId": "0cc8b4be-c727-4411-b478-2c874fbc6c25",
        "partsName": "部品A",
        "supportPartsName": "modelA",
        "terminatedFlag": false,
        "amountRequired": null,
        "amountRequiredUnit": "kilogram",
        "partsLabelName": "PartsA",
        "partsAddInfo1": "Ver3.0",
        "partsAddInfo2": "2024-12-01-2024-12-31",
        "partsAddInfo3": "任意の情報が入ります"
    },
    "childrenPartsModel": [
        {
            "traceId": "40b77952-2c89-49be-8ce9-7c64a15e0ae7",
            "operatorId": "b39e6248-c888-56ca-d9d0-89de1b1adc8e",
            "plantId": "0cc8b4be-c727-4411-b478-2c874fbc6c25",
            "partsName": "部品A1",
            "supportPartsName": "modelA-1",
            "terminatedFlag": false,
            "amountRequired": 5,
            "amountRequiredUnit": "kilogram",
            "partsLabelName": "PartsA1",
            "partsAddInfo1": "Ver3.0",
            "partsAddInfo2": "2024-12-01-2024-12-31",
            "partsAddInfo3": "任意の情報が入ります"
        }
    ]
}
```

登録された部品構成情報が返却されます。  
製品登録が成功すると子部品に対してトレース識別子が付番されます。

# CFP結果提出の依頼
A社がB社にCFP結果提出の依頼を作成する例を示します。

## 1. B社の事業者識別子（内部）の検索
Action (A社): 下記の```curl```コマンドを実行し、B社の事業者識別子（内部）を検索します。

A社はあらかじめB社の事業者識別子（公開）を知っている必要があります。

```
curl --location --request GET 'http://localhost:8081/api/v1/authInfo?dataTarget=operator&openOperatorId=1234567890124' \
--header 'apiKey: Sample-APIKey1' \
--header 'Authorization: Bearer eyJhbGciOiJub25lIiwidHlwIjoiSldUIn0.eyJvcGVyYXRvcl9pZCI6ImIzOWU2MjQ4LWM4ODgtNTZjYS1kOWQwLTg5ZGUxYjFhZGM4ZSIsImVtYWlsIjoib2VtX2FAZXhhbXBsZS5jb20iLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsImF1dGhfdGltZSI6MTcxMzc3MjYzMiwidXNlcl9pZCI6IjUxZjk3MGE3LTFiYmUtNDAxNS1hNzI0LWIzZjUwYzljMWRjZCIsImZpcmViYXNlIjp7ImlkZW50aXRpZXMiOnsiZW1haWwiOlsib2VtX2FAZXhhbXBsZS5jb20iXX0sInNpZ25faW5fcHJvdmlkZXIiOiJwYXNzd29yZCJ9LCJpYXQiOjE3MTM3NzI2MzIsImV4cCI6MTcxMzc3NjIzMiwiYXVkIjoibG9jYWwiLCJpc3MiOiJodHRwczovL3NlY3VyZXRva2VuLmdvb2dsZS5jb20vbG9jYWwiLCJzdWIiOiI1MWY5NzBhNy0xYmJlLTQwMTUtYTcyNC1iM2Y1MGM5YzFkY2QifQ.'
```

```json
{
    "operatorId": "15572d1c-ec13-0d78-7f92-dd4278871373",
    "operatorName": "B社",
    "operatorAddress": "東京都渋谷区xx",
    "openOperatorId": "1234567890124",
    "operatorAttribute": {
        "globalOperatorId": "1234ABCD5678EFGH0124"
    }
}
```

## 2. A社からB社への取引関係の作成
Action (A社): 下記の```curl```コマンドを実行し、部品A1の取引関係をB社に対して作成し、紐付け回答を依頼します。

```
curl --location --request PUT 'http://localhost:8080/api/v1/datatransport?dataTarget=tradeRequest' \
--header 'apiKey: Sample-APIKey1' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJub25lIiwidHlwIjoiSldUIn0.eyJvcGVyYXRvcl9pZCI6ImIzOWU2MjQ4LWM4ODgtNTZjYS1kOWQwLTg5ZGUxYjFhZGM4ZSIsImVtYWlsIjoib2VtX2FAZXhhbXBsZS5jb20iLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsImF1dGhfdGltZSI6MTcxMzc3MjYzMiwidXNlcl9pZCI6IjUxZjk3MGE3LTFiYmUtNDAxNS1hNzI0LWIzZjUwYzljMWRjZCIsImZpcmViYXNlIjp7ImlkZW50aXRpZXMiOnsiZW1haWwiOlsib2VtX2FAZXhhbXBsZS5jb20iXX0sInNpZ25faW5fcHJvdmlkZXIiOiJwYXNzd29yZCJ9LCJpYXQiOjE3MTM3NzI2MzIsImV4cCI6MTcxMzc3NjIzMiwiYXVkIjoibG9jYWwiLCJpc3MiOiJodHRwczovL3NlY3VyZXRva2VuLmdvb2dsZS5jb20vbG9jYWwiLCJzdWIiOiI1MWY5NzBhNy0xYmJlLTQwMTUtYTcyNC1iM2Y1MGM5YzFkY2QifQ.' \
--data '{
  "statusModel": {
    "message": "来月中にご回答をお願いします。",
    "replyMessage": null,
    "requestStatus": {},
    "requestType": "CFP",
    "statusId": null,
    "tradeId": null,
    "responseDueDate": "2024-12-31"
  },
  "tradeModel": {
    "downstreamOperatorId": "b39e6248-c888-56ca-d9d0-89de1b1adc8e",
    "downstreamTraceId": "40b77952-2c89-49be-8ce9-7c64a15e0ae7",
    "tradeId": null,
    "upstreamOperatorId": "15572d1c-ec13-0d78-7f92-dd4278871373",
    "upstreamTraceId": null
  }
}'
```

```json
{
    "tradeModel": {
        "tradeId": "f475cb75-b3b8-4427-9e8d-376377f1c795",
        "downstreamOperatorId": "b39e6248-c888-56ca-d9d0-89de1b1adc8e",
        "upstreamOperatorId": "15572d1c-ec13-0d78-7f92-dd4278871373",
        "downstreamTraceId": "40b77952-2c89-49be-8ce9-7c64a15e0ae7",
        "upstreamTraceId": null
    },
    "statusModel": {
        "statusId": "92e42cc2-7768-4bc8-b956-356a40a7a506",
        "tradeId": "f475cb75-b3b8-4427-9e8d-376377f1c795",
        "requestStatus": {
            "completedCount": null,
            "completedCountModifiedAt": null,
            "tradesCount": null,
            "tradesCountModifiedAt": null
        },
        "message": "来月中にご回答をお願いします。",
        "replyMessage": null,
        "requestType": "CFP",
        "responseDueDate":"2024-12-31"
    }
}
```

製品登録が成功すると取引関係情報識別子が付番されます。  
※ これらはダミー実装のため、本番の返却値とは一部値が異なる可能性がございます。