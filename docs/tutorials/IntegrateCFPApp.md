# データ流通システム・ユーザ認証システム・CFPアプリ参考実装例 連携手順

本ドキュメントでは、以下の3ソフトウェアを連携して動作させる手順について説明する。

* [データ流通システム](https://github.com/ouranos-ecosystem-interop-data-infra/ouranos-ecosystem-data-transaction-system)
* [ユーザ認証システム](https://github.com/ouranos-ecosystem-interop-data-infra/ouranos-ecosystem-user-authentication-system)
* CFPアプリ参考実装例

  * [APサーバ](https://github.com/ouranos-ecosystem-interop-data-infra/ouranos-ecosystem-sample-application-cfp-backend)
  * [Web UI](https://github.com/ouranos-ecosystem-interop-data-infra/ouranos-ecosystem-sample-application-cfp-frontend)
  * [リバースプロキシ](https://github.com/ouranos-ecosystem-interop-data-infra/ouranos-ecosystem-sample-application-cfp-proxy)

## 前提

この例では2台のサーバに、以下に示すソフトウェアを配備する。

* データ流通サーバ (Ubuntu 22.04 LTS): データ流通システム、ユーザ認証システム
* CFPアプリサーバ (Windows 10): CFPアプリ参考実装例

なお本構成では、蓄電池トレーサビリティシステム管理システムには、データ流通システムが提供するダミー実装を使用する。
そのため、通知機能などCFPアプリ参考実装例が提供する機能の一部は利用できない。

## 手順

1. [データ流通システムのドキュメント](https://github.com/ouranos-ecosystem-interop-data-infra/ouranos-ecosystem-data-transaction-system#getting-started) に従い、データ流通システムとユーザ認証システムを、データ流通サーバ上に構築する。

   構築後、外部からのリクエストを URL 中のパスに応じて振り分けるため、以下の手順に従い nginx のインストールと設定を行う。

   ```
   # nginx パッケージをインストールする
   $ sudo apt-get update
   $ sudo apt-get install -y nginx
   
   # 以下の設定を記述した conf ファイルを, /etc/nginx/conf.d 以下に作成する
   $ sudo vi /etc/nginx/conf.d/server.conf 
   $ cat /etc/nginx/conf.d/server.conf 
   server {
       listen       80;
       listen  [::]:80;
   
       location ~* auth {
         client_max_body_size 10M;
         proxy_buffering off;
         proxy_pass http://(データ流通サーバのIPアドレス):8081$request_uri;
         proxy_read_timeout 300;
       }
   
       location /api/v1/datatransport {
         client_max_body_size 10M;
         proxy_buffering off;
         proxy_pass http://(データ流通サーバのIPアドレス):8080$request_uri;
         proxy_read_timeout 300;
       }
   }
   
   # 作成した設定が上書きされないよう、あらかじめ存在する設定ファイルを退避する
   $ sudo mv /etc/nginx/sites-enabled/default /tmp
   
   # 作成した設定を読み込ませるため、nginx サービスを再起動する
   $ sudo systemctl restart nginx.service
   ```

2. 以下のドキュメントに従い, CFPアプリ参考実装例の各コンポーネントをCFPアプリサーバ上に構築する。注意点を以下に示す。

   a. [APサーバ](https://github.com/ouranos-ecosystem-interop-data-infra/ouranos-ecosystem-sample-application-cfp-backend#ビルド起動手順)

      * アプリケーションのビルド前に、src/main/resources/META-INF/spring/common-backend-infra.properties に以下の情報を設定すること。

        | プロパティ | 設定値 |
        | ---- | ---- |
        | URL_DATA_TRANSPORT_SYSTEM      | http://(データ流通サーバのFQDNもしくはIPアドレス)/ |
        | API_KEY_DATA_TRANSPORT_SYSTEM  | Sample-APIKey1 |
        | URL_INTROSPECTION_ENDPOINT     | http://(データ流通サーバのFQDNもしくはIPアドレス)/api/v1/systemAuth/token |
        | API_KEY_INTROSPECTION_ENDPOINT | Sample-APIKey2 |

   b. [Web UI](https://github.com/ouranos-ecosystem-interop-data-infra/ouranos-ecosystem-sample-application-cfp-frontend#起動手順)

      * アプリケーションのビルド前に、プロジェクトのルートディレクトリに以下のファイルを配置しておくこと。

        ```
        $ cat .env.local
        NEXT_PUBLIC_DATA_TRANSPORT_API_BASE_URL=http://localhost
        ```

   c. [リバースプロキシ](https://github.com/ouranos-ecosystem-interop-data-infra/ouranos-ecosystem-sample-application-cfp-proxy#ビルド起動手順)

3. CFPアプリサーバ上でブラウザを開き, http://localhost/ にアクセスすると、以下のページが表示される。

   ![](/docs/assets/images/cfpapp_login.png)

   ユーザ認証システムの構築時に cmd/add_local_user/data/seed.csv で指定した事業者アカウント・パスワードでログインすると、
   CFPアプリ参考実装例を使用することができる。

   ![](/docs/assets/images/cfpapp_parts.png)

   ![](/docs/assets/images/cfpapp_plants.png)
