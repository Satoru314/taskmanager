# ベースイメージを指定 (バージョンはお好みで)
FROM postgres:16

# デフォルトのPostgreSQLユーザー名を設定
ENV POSTGRES_USER=docker