#!/bin/sh

# /health が 403 を返すようファイル作成を行う
# これにより Pod の ReadinessProbe / LivenessProbe に失敗するようになる
echo > $HOME/.shutdown

# `Pod.spec.terminationGracePeriodSeconds: 30` → 28秒待つ
# この間に以下を行う
# * Request に対し Response が返せてないものを (SIGTERM せずに) 28 秒間待つ
# * Service の Endpoint list から該当 Pod が削除されるため、新規の通信は受け付けない
sleep 28

# `Pod.spec.terminationGracePeriodSeconds: 30` → 残り2秒
# 終了処理を以下に記述
# 例. Nginxプロセスの停止: `nginx -s quit`


