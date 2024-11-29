# Minimum Ouranos data platform
本リポジトリ（Minimum Ouranos data platform）は企業・業界・国境を跨いだデータ連携・利活用を目指すイニシアティブの
「ウラノス・エコシステム（Ouranos Ecosystem）」におけるデータ流通システムの最小実装を体験するため、
実装の一部をオープンソースとして公開する。

本リポジトリとして体験できる点は以下の観点である。
- データ流通システムおよびユーザ認証システムの理解
- ローカル環境におけるデータ流通システムおよびユーザ認証システムのサンプル実行

本リポジトリ内で実行できるトレーサビリティ管理システムはダミー実装であり、
実際の業務システムと連携する際は挙動が異なる可能性があるため、注意すること。

## 基本概念
データ流通システム内でのデータ連携は、誰もが参加しやすいよう軽量なデータ交換の仕組みとし、他のデータスペースなど基盤外とのデータ連携は、プロトコルをシンプル化・標準化することにより相互運用性を確保する。それらの汎用化されたデータの上で、業界毎に業界固有ルールのシステム化を行う。

現在公開されている[システム化ガイドラインβ版](https://www.ipa.go.jp/digital/architecture/Individual-link/m42obm0000008rd4-att/guideline-for-datacooperation-in-BattCFPDD-beta.pdf "サプライチェーン上のデータ連携の仕組みに関するガイドラインβ版")より、データ連携のシステムアーキテクチャを抜粋して示す。

![データ流通システムのシステムアーキテクチャ](/docs/assets/images/DataPlatformArchitecture.png)
*システム化ガイドラインβ版 37頁参照*

- コネクタ：データのPut/Getのような軽量なデータ交換のI/Fを提供する
- トランスフォーム：他システムとの連携に必要なデータ変換を行う
- アダプタ：他システムへのアクセスを提供する

## ビジネスアーキテクチャ
蓄電池トレーサビリティプロジェクトにおけるデータ流通の利活用にて前提となっているビジネスアーキテクチャを示す。

![ビジネスアーキテクチャ](/docs/assets/images/BusinessArchitecture.png)
*システム化ガイドラインβ版 18頁参照*

蓄電池トレーサビリティ管理システムは、トレース識別子（製品・部品に対してトレースを取るために割り当てる、
各者内でユニークに特定可能な識別子）をインデックスとして、トレース識別子同士を紐付けることで、「製品と調達部品の構成関係」及び
「事業者間の取引関係」を記録して、サプライチェーンの追跡を可能にする。

## 業務フロー
本ユースケースにおいて他社間のデータを連携するまでに前提とする業務フローを示す。

1. 基本契約:製品発注時等に規制対応に関する契約を行う
![業務フロー1](/docs/assets/images/BasicBusinessFlow1.png)

2. 依頼業務:川上企業に部品情報とCFPとDDの登録を依頼する
![業務フロー2](/docs/assets/images/BasicBusinessFlow2.png)

3. CFPの計算から最終製品のCFP提出までの業務:出荷する製品のCFPを当局に提出する
![業務フロー3](/docs/assets/images/BasicBusinessFlow3.png)

3. 部品選定またはCFP変更要求に関する業務:CFPが高い部品を特定すると共に、必要に応じて仕入れ先にCFPの提言を依頼するか、代替部品を選定する
![業務フロー4](/docs/assets/images/BasicBusinessFlow4.png)

# Getting Started
本リポジトリ内のユースケースを体験するための環境構築の手順を示す。

## Prerequisites
本リポジトリのシステムが動作するための動作環境の一式を事前にインストールする。
- go関係のソフトウェアをインストール
  - golang
  - golangci-lint
  - mockery
  - goreturns
- マイグレーションツール（golang-migrate）のインストール
- makeをインストール
- docker, docker-composeをインストール

### 動作確認済み実行環境
|Name                                        |Version |Notes|
|:-------------------------------------------|:-------|:----|
|golang                                      |1.22||
|[golangci-lint](https://golangci-lint.run/usage/install/)|1.56.2||
|[mockery](https://vektra.github.io/mockery/)|v2.42.3||
|[goreturns](https://github.com/sqs/goreturns)|-|go install|
|[golang-migrate](https://github.com/golang-migrate/migrate)|-||
|make                                        |GNU make 3.81||
|docker                                      |-||
|docker-compose                              |-||

## Basic Setup
本リポジトリでは、A社およびB社の2社の事業者を作成しデータ交換を実施する。本モデルケースではトレーサビリティ管理システムはダミー実装にて実現しているため、扱うデータはPostgresのデータベース上にて保存・更新をするが、本番環境においては構成が異なる。

以下、`path/to/data-spaces-backend` はローカル環境に `git clone` された本リポジトリへのパスを、
`path/to/authenticator-backend` はローカル環境に `git clone` されたユーザ認証システム（別リポジトリで提供）へのパスを表す。

### 1.外部依存システムの起動

Postgresの起動およびIdPのエミュレーターの起動
```shell
cd <path/to/authenticator-backend>
$ docker-compose up -d
```

### 2. 認証データのセットアップ

1. データベースのスキーマ定義を反映するため、migrationを実行
```shell
cd <path/to/authenticator-backend>
export POSTGRESQL_URL='postgres://dhuser:passw0rd@localhost:5432/dhlocal?sslmode=disable'
migrate -path setup/migrations -database ${POSTGRESQL_URL} up
```

2. A社とB社のデータをDBに必要な情報を作成
```shell
cd <path/to/authenticator-backend>
# ローカルのDBにseedを追加
./setup/setup_seeds.sh
```

[確認]
- DBに接続し、operatorsテーブルのカラム数を確認

3. A社とB社の認証情報をIdPに作成
※ 事業者情報を初期値から変更する場合は```/cmd/add_local_user/data/seed.csv```を編集する。
```shell
cd <path/to/authenticator-backend>
# ローカルのIdPエミュレーターに事業者情報を追加
make idp-add-local
```

[確認]
- http://localhost:4000 にアクセスし、Firebaseのブラウザのauthenticationのタブの中身を確認

### 3. データ流通システム
1. ビルド手順

```shell
cd <path/to/data-spaces-backend>
go build main.go
docker build -t data-spaces-backend .
```

2. 起動手順

```shell
docker run -v $(pwd)/config/:/app/config/ -td -i --network docker.internal --env-file config/local.env -p 8080:8080 --name data-spaces-backend data-spaces-backend
```

### 4. ユーザ認証システム

1. ビルド手順

```shell
cd <path/to/authenticator-backend>
go build main.go
docker build -t authenticator-backend .
```

2. 起動手順

```shell
docker run -v $(pwd)/config/:/app/config/ -td -i --network docker.internal --env-file config/local.env -p 8081:8081 --name authenticator-backend authenticator-backend
```

# Tutorials
環境準備が整ったら、いくつかの業務フローのチュートリアルを実行できる。

- [事業者認証](/docs/tutorials/OperatorIdentification.md)
- [部品登録およびA社からB社へのCFP結果提出の依頼をする(基本フロー2 #4)](/docs/tutorials/MakePartsStructureAndMakeTradeRequest.md)
- [B社からA社へ部品登録紐付けをする(基本フロー2 #31)](/docs/tutorials/MakeTradeResponse.md)
- [B社からA社へCFP情報の伝達をする(基本フロー3 #5)](/docs/tutorials/ResponseCFP.md)
- [B社の回答情報の取得およびA社の完成品のCFPを算出(基本フロー3 #6, #2)](/docs/tutorials/GetResponseAndMakeCFPReport.md)

また、一部機能に制限はあるが、別リポジトリで提供している [CFPアプリ参考実装例](https://github.com/ouranos-ecosystem-idi/sample-application-cfp-frontend) と連携することで、
Web UIから各種操作を行うことも可能である。
手順は [こちら](/docs/tutorials/IntegrateCFPApp.md) を参照のこと。

# 開発者向けドキュメント
開発者向けのドキュメントは[こちら](/docs/architecture/architecture.md)を参照

# 問合せ及び要望に関して
- 本リポジトリは現状は主に配布目的の運用となるため、IssueやPull Requestに関しては受け付けておりません。

# License
- 本リポジトリはMITライセンスで提供されています。
- ソースコードおよび関連ドキュメントの著作権は株式会社NTTデータグループに帰属します。

# 免責事項
- 本リポジトリの内容は予告なく変更・削除する可能性があります。
- 本リポジトリの利用により生じた損失及び損害等について、いかなる責任も負わないものとします。

