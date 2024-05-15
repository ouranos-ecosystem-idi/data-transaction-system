# 事業者認証
A社がユーザ認証システムで事業者認証をした後に、自身の事業者情報を取得する例を示します。

## 1. 事業者認証の実行
Action (A社): 下記の```curl```コマンドを実行し、認証情報を取得します。

前提条件として、プラットフォーム認定を受けた事業者に発行されるApiKeyを指定しAPIを実行します。
また運営事業者から各事業者はAccountIdとPasswordを事前に払い出されています。

```
curl --location --request POST 'http://localhost:8081/auth/login' \
--header 'Content-Type: application/json' \
--header 'apiKey: Sample-APIKey1' \
--data-raw '{
  "operatorAccountId": "oem_a@example.com",
  "accountPassword": "oemA&user_01"
}'
```

```json
{
    "accessToken": "eyJhbGciOiJub25lIiwidHlwIjoiSldUIn0.eyJvcGVyYXRvcl9pZCI6ImIzOWU2MjQ4LWM4ODgtNTZjYS1kOWQwLTg5ZGUxYjFhZGM4ZSIsImVtYWlsIjoib2VtX2FAZXhhbXBsZS5jb20iLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsImF1dGhfdGltZSI6MTcxMTY5MTg5NCwidXNlcl9pZCI6IjcxN2I0NGM3LWQ0MWEtNGU0Mi1iYTEzLWZjMTMzYzM5M2Y5OSIsImZpcmViYXNlIjp7ImlkZW50aXRpZXMiOnsiZW1haWwiOlsib2VtX2FAZXhhbXBsZS5jb20iXX0sInNpZ25faW5fcHJvdmlkZXIiOiJwYXNzd29yZCJ9LCJpYXQiOjE3MTE2OTE4OTQsImV4cCI6MTcxMTY5NTQ5NCwiYXVkIjoibG9jYWwiLCJpc3MiOiJodHRwczovL3NlY3VyZXRva2VuLmdvb2dsZS5jb20vbG9jYWwiLCJzdWIiOiI3MTdiNDRjNy1kNDFhLTRlNDItYmExMy1mYzEzM2MzOTNmOTkifQ.",
    "refreshToken": "eyJfQXV0aEVtdWxhdG9yUmVmcmVzaFRva2VuIjoiRE8gTk9UIE1PRElGWSIsImxvY2FsSWQiOiI3MTdiNDRjNy1kNDFhLTRlNDItYmExMy1mYzEzM2MzOTNmOTkiLCJwcm92aWRlciI6InBhc3N3b3JkIiwiZXh0cmFDbGFpbXMiOnt9LCJwcm9qZWN0SWQiOiJsb2NhbCJ9"
}
```

認証が成功すると返却値としてjson web tokenが払い出されます。
ユーザ認証システムにはheaderにApiKeyおよびTokenを必ず指定してAPIを実行します。

## 2. 事業者情報の取得
Action (A社): 下記のcurlコマンドを実行し、自社の事業者情報を取得します。
```
curl --location --request GET 'http://localhost:8081/api/v1/authInfo?dataTarget=operator' \
--header 'apiKey: Sample-APIKey1' \
--header 'Authorization: Bearer eyJhbGciOiJub25lIiwidHlwIjoiSldUIn0.eyJvcGVyYXRvcl9pZCI6ImIzOWU2MjQ4LWM4ODgtNTZjYS1kOWQwLTg5ZGUxYjFhZGM4ZSIsImVtYWlsIjoib2VtX2FAZXhhbXBsZS5jb20iLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsImF1dGhfdGltZSI6MTcxMTY5MTg5NCwidXNlcl9pZCI6IjcxN2I0NGM3LWQ0MWEtNGU0Mi1iYTEzLWZjMTMzYzM5M2Y5OSIsImZpcmViYXNlIjp7ImlkZW50aXRpZXMiOnsiZW1haWwiOlsib2VtX2FAZXhhbXBsZS5jb20iXX0sInNpZ25faW5fcHJvdmlkZXIiOiJwYXNzd29yZCJ9LCJpYXQiOjE3MTE2OTE4OTQsImV4cCI6MTcxMTY5NTQ5NCwiYXVkIjoibG9jYWwiLCJpc3MiOiJodHRwczovL3NlY3VyZXRva2VuLmdvb2dsZS5jb20vbG9jYWwiLCJzdWIiOiI3MTdiNDRjNy1kNDFhLTRlNDItYmExMy1mYzEzM2MzOTNmOTkifQ.'
```

A社の事業者情報が返却されます。
```json
{
    "operatorId": "b39e6248-c888-56ca-d9d0-89de1b1adc8e",
    "operatorName": "A社",
    "operatorAddress": "東京都渋谷区xx",
    "openOperatorId": "1234567890123",
    "operatorAttribute": {
        "globalOperatorId": "1234ABCD5678EFGH0123"
    }
}
```