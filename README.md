# doramaba

doramaba とかテレビバって考えてたけど，ShiJimi にしよっかな  
(Shiny and Jiminy)(光り輝く原石を・驚きとともに))略してしじみ．

# To Do List
## ますと
1. バグ探し（だいたいちゃんと動くように，エラー画面が出ないように）
1. コメントのテレビを動画に変更

## その他
1. コメント分析（とりあえず総数）
1. レビュー分析（とりあえず総数，スター，タグ）
1. 
1. ポイントの利用方法考える．（１日一回ログインで1ポイント獲得）
    1. 表示数の上限に利用できる（自身の投稿，テレビのコメント，テレビの検索）
    1. 現状：コメント表示上限200(テレビごと), レビュー：100(テレビごと), tv（おすすめ）:100，コメント:5~180字，レビュー20~400字，コメント過去1000(ユーザごと)，レビュー過去：100(ユーザごと)，watched_tv:1000(ユーザごと)，wtw_tv：1000(ユーザごと)
    1. かわいいバッジがもらえる(寄付ありがとう（有料），ログイン100日目，レビュー20番組，番組作成10個，1000コメント達成，祝5年，)
    1. ポイントの購入方法は要検討（ひとまずなし/paypal）
    1. おすすめ欄のタイトルのみ表示
    1. 自身のコメントの検索機能
    1. 広告の非表示
1. ログアウトページてきとう
1. ログインからの導線
1. カスタマイズ機能
1. ポップアップ系
1. スマホコピペ・アドレスバー(iosのwindow.scrollToがきかない，疑似的にスクロールアクションを起こせれば．)
1. js場所（トップページでツールバーずれることがある）
1. tv edit部分のタイトルかぶり処理，なぜかFlagがNULLになる
1. コメント・レビューの分析表示部分
1. コード関数化してまとめる jsもgoも
1. androidでの見た目がmarital，ちょっと変かも
1. コメントpullhookでtwitterぽく上スクロールしたい（ajaxでAPI叩く→更新部分の要素のみ追加）
1. いろんな説明文
1. 各ユーザへの管理側からのメッセージ
1. カテゴリクリック・キャストクリックで検索
1. 寄付マーククリックアクション
1. お知らせブログ・「機能追加しました！」
1. googleアナリティクス
1. google広告挿入(すぐにしない)
1. トップ画面に今季の映画？
1. 各コメント欄で関連番組
1. indexからドラマのみ検索（映画多すぎ）
1. コメントの通報（記録してAI化）
1. 登録番組の通報（悪質投稿）
1. テレビ登録の承認機能(CountAuthorization)
1. 最初のページ（ふわっと）
1. 現在〇人が参加しています．
1. 自分的ランキング（今季ベスト３，3大西島さんのドラマ）
1. 質の高いユーザ，悪質ユーザ
1. 3大〇〇
1. スワイプで詳細（コメント：単語，活発ユーザ，閲覧履歴，レビュー：星（ある程度集まったら表示），タグ頻度，単語，テレビ index：ジャンルごとトップとか，見た人ランキングとか）
1. あなたにおすすめアルゴリズム
1. ダイレクトメッセージ・フォロー機能はポイント化
1. グループトーク
1. その他サイト作成（CM・本・バラエティ・豆知識・スポーツ観戦・ニュース・不満・3 大〇〇）

# Improve Coding
1. javascript 内で cookie が使えません！
1. 今どきの javascript の書き方
1. shijimi.com/tv/〇〇の tv を別サイトの場合は変えたいのだが，それって別の go を起動してもいける？

# 知識いろいろ
1. vscode が便利
1. prettier で自動コード補正
1. ESlint でコード指南
1. travisいれたい
1. bee run -downdoc=true -gendoc=true でswagger発動
1. javascript → main.go内のgo関数で処理？
1. regexpは遅いらしい（多様部分には使わない）
1. 検索は自動的に「にっぽん」でも「ニッポン」でもヒットする(.ymlの--collation-server=変更でストップ)

# 環境構築
docker, docker-composeが必要
1. git clone ~
1. static/js/common.jsのトップにあるURLを変更（自機の場合:localhost）
1. docker-compose up -d --build　(imageの作成・コンテナの作成・コンテナの起動)
1. docker exec -it go_app /bin/sh　(コンテナに入る))
1. bee run (サーバ起動)
1. localhost:8080番でアプリ，8000番でphpmyadmin(たまにエラー出る)
1. exit (コンテナから脱出)
1. docker-compose down (コンテナの停止・削除)
1. sudo chown 自分 -R . (MySQLフォルダの権限エラーを防ぐ)

# 外部公開
1. CDNのバージョンを固定する
1. 開発モードやめる
1. prod に変更
1. パスワードをprod版にする

## herokuの場合
1. https://qiita.com/pitcher292/items/1ca39c7b0dbd79298c0b
1. 30分でスタンバイ状態：再起動時にきどうしなおしちゃう
1. herokuだとsearch_historyがでてこない

# 公開へのロードマップ
1. konohaVPS900円借りる
1. SSD50GB上に全部乗せる（DB：mysql 500/monthもあるらしい）
1. のっける

# エラーがおきた
1. ローカルPCからリクエストが投げられない．クロスドメインエラー
    1. https://qiita.com/growsic/items/a919a7e2a665557d9cf4
    1. または，common.jsのURLがおかしい
1. herokuにあげたアプリにリクエスト投げられない
    1. httpではだめ．https

# ポイント番号
1. 寄付：1~10
1. 毎日ログイン：10~20
1. 機能系：100~200
    1. コメント閲覧
    1. コメント投稿
    1. コメント検索
    1. レビュー閲覧
    1. レビュー投稿
    1. レビュー検索
    1. テレビ閲覧
    1. テレビ投稿
    1. テレビ検索
1. 
1. 
1. 
