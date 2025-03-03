# 部品登録紐付け
B社が部品登録紐付けをする例を示します。  
識別子およびアクセストークンは適宜置き換えが必要ですのでご注意下さい。

## 1. 事業者認証の実行
Action (B社): 下記の```curl```コマンドを実行し、B社の認証情報を取得します。

前提条件として、プラットフォーム認定を受けた事業者に発行されるApiKeyを指定しAPIを実行します。
また運営事業者から各事業者はAccountIdとPasswordを事前に払い出されています。

```
curl --location --request POST 'http://localhost:8081/auth/login' \
--header 'Content-Type: application/json' \
--header 'apiKey: Sample-APIKey1' \
--data-raw '{
  "operatorAccountId": "supplier_b@example.com",
  "accountPassword": "supplierB&user_01"
}'
```

```json
{
    "accessToken": "eyJhbGciOiJub25lIiwidHlwIjoiSldUIn0.eyJvcGVyYXRvcl9pZCI6IjE1NTcyZDFjLWVjMTMtMGQ3OC03ZjkyLWRkNDI3ODg3MTM3MyIsImVtYWlsIjoic3VwcGxpZXJfYkBleGFtcGxlLmNvbSIsImVtYWlsX3ZlcmlmaWVkIjpmYWxzZSwiYXV0aF90aW1lIjoxNzEzNzc2ODUxLCJ1c2VyX2lkIjoiMzBhNTE1OGItN2E0NS00ODZjLTgxMmMtNTllZDBmM2RjYjA1IiwiZmlyZWJhc2UiOnsiaWRlbnRpdGllcyI6eyJlbWFpbCI6WyJzdXBwbGllcl9iQGV4YW1wbGUuY29tIl19LCJzaWduX2luX3Byb3ZpZGVyIjoicGFzc3dvcmQifSwiaWF0IjoxNzEzNzc2ODUxLCJleHAiOjE3MTM3ODA0NTEsImF1ZCI6ImxvY2FsIiwiaXNzIjoiaHR0cHM6Ly9zZWN1cmV0b2tlbi5nb29nbGUuY29tL2xvY2FsIiwic3ViIjoiMzBhNTE1OGItN2E0NS00ODZjLTgxMmMtNTllZDBmM2RjYjA1In0.",
    "refreshToken": "eyJfQXV0aEVtdWxhdG9yUmVmcmVzaFRva2VuIjoiRE8gTk9UIE1PRElGWSIsImxvY2FsSWQiOiIzMGE1MTU4Yi03YTQ1LTQ4NmMtODEyYy01OWVkMGYzZGNiMDUiLCJwcm92aWRlciI6InBhc3N3b3JkIiwiZXh0cmFDbGFpbXMiOnt9LCJwcm9qZWN0SWQiOiJsb2NhbCJ9"
}
```

## 2. 事業所の登録
Action (B社): 下記の```curl```コマンドを実行し、B社に事業所(B工場)を作成します。

事業者Bには、自社の製品を製造している事業所の登録がないため、新規に作成する必要があります。

```
curl --location --request PUT 'http://localhost:8081/api/v1/authInfo?dataTarget=plant' \
--header 'apiKey: Sample-APIKey1' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJub25lIiwidHlwIjoiSldUIn0.eyJvcGVyYXRvcl9pZCI6IjE1NTcyZDFjLWVjMTMtMGQ3OC03ZjkyLWRkNDI3ODg3MTM3MyIsImVtYWlsIjoic3VwcGxpZXJfYkBleGFtcGxlLmNvbSIsImVtYWlsX3ZlcmlmaWVkIjpmYWxzZSwiYXV0aF90aW1lIjoxNzEzNzc2ODUxLCJ1c2VyX2lkIjoiMzBhNTE1OGItN2E0NS00ODZjLTgxMmMtNTllZDBmM2RjYjA1IiwiZmlyZWJhc2UiOnsiaWRlbnRpdGllcyI6eyJlbWFpbCI6WyJzdXBwbGllcl9iQGV4YW1wbGUuY29tIl19LCJzaWduX2luX3Byb3ZpZGVyIjoicGFzc3dvcmQifSwiaWF0IjoxNzEzNzc2ODUxLCJleHAiOjE3MTM3ODA0NTEsImF1ZCI6ImxvY2FsIiwiaXNzIjoiaHR0cHM6Ly9zZWN1cmV0b2tlbi5nb29nbGUuY29tL2xvY2FsIiwic3ViIjoiMzBhNTE1OGItN2E0NS00ODZjLTgxMmMtNTllZDBmM2RjYjA1In0.' \
--data '{
  "openPlantId": "1234567890124012345",
  "operatorId": "15572d1c-ec13-0d78-7f92-dd4278871373",
  "plantAddress": "xx県xx市xxxx町2-1-1234",
  "plantId": null,
  "plantName": "B工場",
  "plantAttribute": {}
}'
```

```json
{
    "plantId": "544a5a35-dab3-469f-8ff5-116a4fe483e8",
    "operatorId": "15572d1c-ec13-0d78-7f92-dd4278871373",
    "plantName": "B工場",
    "plantAddress": "xx県xx市xxxx町2-1-1234",
    "openPlantId": "1234567890124012345",
    "plantAttribute": {
        "globalPlantId": null
    }
}
```

## 3. 親部品情報の作成
Action (B社): 下記の```curl```コマンドを実行し、B社の部品Bを登録します。  
登録する製品の情報は下記です。  
- 部品項目:部品B
- 補助項目:modelB
- 事業所識別子（内部）:{事業所(工場B)で登録した識別子}
- 事業者識別子（内部）:15572d1c-ec13-0d78-7f92-dd4278871373


```
curl --location --request PUT 'http://localhost:8080/api/v1/datatransport?dataTarget=parts' \
--header 'apiKey: Sample-APIKey1' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJub25lIiwidHlwIjoiSldUIn0.eyJvcGVyYXRvcl9pZCI6IjE1NTcyZDFjLWVjMTMtMGQ3OC03ZjkyLWRkNDI3ODg3MTM3MyIsImVtYWlsIjoic3VwcGxpZXJfYkBleGFtcGxlLmNvbSIsImVtYWlsX3ZlcmlmaWVkIjpmYWxzZSwiYXV0aF90aW1lIjoxNzEzNzc2ODUxLCJ1c2VyX2lkIjoiMzBhNTE1OGItN2E0NS00ODZjLTgxMmMtNTllZDBmM2RjYjA1IiwiZmlyZWJhc2UiOnsiaWRlbnRpdGllcyI6eyJlbWFpbCI6WyJzdXBwbGllcl9iQGV4YW1wbGUuY29tIl19LCJzaWduX2luX3Byb3ZpZGVyIjoicGFzc3dvcmQifSwiaWF0IjoxNzEzNzc2ODUxLCJleHAiOjE3MTM3ODA0NTEsImF1ZCI6ImxvY2FsIiwiaXNzIjoiaHR0cHM6Ly9zZWN1cmV0b2tlbi5nb29nbGUuY29tL2xvY2FsIiwic3ViIjoiMzBhNTE1OGItN2E0NS00ODZjLTgxMmMtNTllZDBmM2RjYjA1In0.' \
--data '{
  "amountRequired": null,
  "amountRequiredUnit": "kilogram",
  "operatorId": "15572d1c-ec13-0d78-7f92-dd4278871373",
  "partsName": "部品B",
  "plantId": "544a5a35-dab3-469f-8ff5-116a4fe483e8",
  "supportPartsName": "modelB",
  "terminatedFlag": true,
  "traceId": null,
  "partsLabelName": "PartsB",
  "partsAddInfo1": "Ver3.0",
  "partsAddInfo2": "2024-12-01-2024-12-31",
  "partsAddInfo3": "任意の情報が入ります"
}'
```

```json
{
    "traceId": "2fb97052-250b-44de-acbb-1ba63e28af71",
    "operatorId": "15572d1c-ec13-0d78-7f92-dd4278871373",
    "plantId": "544a5a35-dab3-469f-8ff5-116a4fe483e8",
    "partsName": "部品B",
    "supportPartsName": "modelB",
    "terminatedFlag": true,
    "amountRequired": null,
    "amountRequiredUnit": "kilogram",
    "partsLabelName": "PartsB",
    "partsAddInfo1": "Ver3.0",
    "partsAddInfo2": "2024-12-01-2024-12-31",
    "partsAddInfo3": "任意の情報が入ります"
}
```

製品登録が成功するとトレース識別子が付番されます。

## 4. 部品登録紐付けの依頼確認
Action (B社): 下記のcurlコマンドを実行し、自社に対して紐付け依頼を受領している一覧を検索する

```
curl --location --request GET 'http://localhost:8080/api/v1/datatransport?dataTarget=tradeResponse' \
--header 'apiKey: Sample-APIKey1' \
--header 'Authorization: Bearer eyJhbGciOiJub25lIiwidHlwIjoiSldUIn0.eyJvcGVyYXRvcl9pZCI6IjE1NTcyZDFjLWVjMTMtMGQ3OC03ZjkyLWRkNDI3ODg3MTM3MyIsImVtYWlsIjoic3VwcGxpZXJfYkBleGFtcGxlLmNvbSIsImVtYWlsX3ZlcmlmaWVkIjpmYWxzZSwiYXV0aF90aW1lIjoxNzEzNzc2ODUxLCJ1c2VyX2lkIjoiMzBhNTE1OGItN2E0NS00ODZjLTgxMmMtNTllZDBmM2RjYjA1IiwiZmlyZWJhc2UiOnsiaWRlbnRpdGllcyI6eyJlbWFpbCI6WyJzdXBwbGllcl9iQGV4YW1wbGUuY29tIl19LCJzaWduX2luX3Byb3ZpZGVyIjoicGFzc3dvcmQifSwiaWF0IjoxNzEzNzc2ODUxLCJleHAiOjE3MTM3ODA0NTEsImF1ZCI6ImxvY2FsIiwiaXNzIjoiaHR0cHM6Ly9zZWN1cmV0b2tlbi5nb29nbGUuY29tL2xvY2FsIiwic3ViIjoiMzBhNTE1OGItN2E0NS00ODZjLTgxMmMtNTllZDBmM2RjYjA1In0.'
```

```json
[
    {
        "statusModel": {
            "statusId": "92e42cc2-7768-4bc8-b956-356a40a7a506",
            "tradeId": "f475cb75-b3b8-4427-9e8d-376377f1c795",
            "requestStatus": {
                "cfpResponseStatus": "NOT_COMPLETED",
                "tradeTreeStatus": "UNTERMINATED",
                "completedCount": 0,
                "completedCountModifiedAt": "2024-11-15T04:49:44Z",
                "tradesCount": 1,
                "tradesCountModifiedAt": "2024-11-15T04:49:44Z"
            },
            "message": "来月中にご回答をお願いします。",
            "replyMessage": null,
            "requestType": "CFP",
            "responseDueDate": "2024-12-31"
        },
        "tradeModel": {
            "tradeId": "f475cb75-b3b8-4427-9e8d-376377f1c795",
            "downstreamOperatorId": "b39e6248-c888-56ca-d9d0-89de1b1adc8e",
            "upstreamOperatorId": "15572d1c-ec13-0d78-7f92-dd4278871373",
            "downstreamTraceId": "40b77952-2c89-49be-8ce9-7c64a15e0ae7",
            "upstreamTraceId": null
        },
        "partsModel": {
            "traceId": "40b77952-2c89-49be-8ce9-7c64a15e0ae7",
            "operatorId": "b39e6248-c888-56ca-d9d0-89de1b1adc8e",
            "plantId": "0cc8b4be-c727-4411-b478-2c874fbc6c25",
            "partsName": "部品A1",
            "supportPartsName": "modelA-1",
            "terminatedFlag": false,
            "amountRequired": null,
            "amountRequiredUnit": "kilogram",
            "partsLabelName": "PartsA1",
            "partsAddInfo1": "Ver3.0",
            "partsAddInfo2": "2024-12-01-2024-12-31",
            "partsAddInfo3": "任意の情報が入ります"
        }
    }
]
```

回答依頼を受けている製品情報と依頼メッセージを確認することができます。

## 5. 部品登録紐付けの登録
Action (B社): 下記のcurlコマンドを実行し、受領している依頼に対して自社の製品を紐付け登録する

登録する情報は下記です。
- tradeId:紐付け回答する対象の取引関係情報識別子
- traceId:紐付けする自社製品のトレース識別子

```
curl --location --request PUT 'http://localhost:8080/api/v1/datatransport?dataTarget=tradeResponse&tradeId=f475cb75-b3b8-4427-9e8d-376377f1c795&traceId=2fb97052-250b-44de-acbb-1ba63e28af71' \
--header 'apiKey: Sample-APIKey1' \
--header 'Authorization: Bearer eyJhbGciOiJub25lIiwidHlwIjoiSldUIn0.eyJvcGVyYXRvcl9pZCI6IjE1NTcyZDFjLWVjMTMtMGQ3OC03ZjkyLWRkNDI3ODg3MTM3MyIsImVtYWlsIjoic3VwcGxpZXJfYkBleGFtcGxlLmNvbSIsImVtYWlsX3ZlcmlmaWVkIjpmYWxzZSwiYXV0aF90aW1lIjoxNzEzNzc2ODUxLCJ1c2VyX2lkIjoiMzBhNTE1OGItN2E0NS00ODZjLTgxMmMtNTllZDBmM2RjYjA1IiwiZmlyZWJhc2UiOnsiaWRlbnRpdGllcyI6eyJlbWFpbCI6WyJzdXBwbGllcl9iQGV4YW1wbGUuY29tIl19LCJzaWduX2luX3Byb3ZpZGVyIjoicGFzc3dvcmQifSwiaWF0IjoxNzEzNzc2ODUxLCJleHAiOjE3MTM3ODA0NTEsImF1ZCI6ImxvY2FsIiwiaXNzIjoiaHR0cHM6Ly9zZWN1cmV0b2tlbi5nb29nbGUuY29tL2xvY2FsIiwic3ViIjoiMzBhNTE1OGItN2E0NS00ODZjLTgxMmMtNTllZDBmM2RjYjA1In0.'
```

```json
{
    "tradeId": "f475cb75-b3b8-4427-9e8d-376377f1c795",
    "downstreamOperatorId": "b39e6248-c888-56ca-d9d0-89de1b1adc8e",
    "upstreamOperatorId": "15572d1c-ec13-0d78-7f92-dd4278871373",
    "downstreamTraceId": "40b77952-2c89-49be-8ce9-7c64a15e0ae7",
    "upstreamTraceId": "2fb97052-250b-44de-acbb-1ba63e28af71"
}
```

※ これらはダミー実装のため、本番の返却値とは一部値が異なる可能性がございます。