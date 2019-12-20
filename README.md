# doramaba

doramaba とかテレビバって考えてたけど，ShiJimi にしよっかな  
(Shiny and Jiminy)(光り輝く原石を・驚きとともに))略してしじみ．

# To Do List
## ますと
1. バグ探し（だいたいちゃんと動くように，エラー画面が出ないように）
1. ヒントボタンかいてく
1. ポイント利用する
1. コードの短縮・説明かく
1. 季節ごとの自分的ランキング（コミュニティを作る？（キャストの部屋・今季の部屋（プロフィールでは，自分のだけ見れる）））
1. サーバの契約・ドメイン取得・SSL化（年明けかな）
1. metabase

## その他
1. ホーム画面に登録ポップアップ（serviceworker(chrome/firefox/androidのみ動作中)）
1. おすすめページでデータ数が40超えてくると，2重に繰り返される（表示側の問題/chromeのみ）
1. スクロールポジションがコメントリロード時にも固定される（表示側の問題/firefoxのみ）
1. 時間指定検索で9時間のズレ（herokuのみ）
1. たまにcomenntのpostデータが欠損する（herokuのみ）
1. コメントリロードすると検索キャッシュ消える（herokuのみ）
1. 自動更新（検索の場合どうなるか(検索時は無効化しといた)）
1. コード埋め込み用script
1. ポイントの利用方法考える．（１日一回ログインで1ポイント獲得）
    1. 表示数の上限に利用できる（自身の投稿，テレビのコメント，テレビの検索）
    1. 現状：コメント表示上限200(テレビごと), レビュー：100(テレビごと), tv（おすすめ）:100，コメント:5~180字，レビュー20~400字，コメント過去1000(ユーザごと)，レビュー過去：100(ユーザごと)，watched_tv:1000(ユーザごと)，wtw_tv：1000(ユーザごと)
    1. かわいいバッジがもらえる(寄付ありがとう（有料），ログイン100日目，レビュー20番組，番組作成10個，1000コメント達成，祝5年，)
    1. ポイントの購入方法は要検討（ひとまずなし/paypal）
    1. おすすめ欄のタイトルのみ表示
    1. 自身のコメントの検索機能
    1. 広告の非表示
    1. 自動リロード時間を設定（デフォ：30sec）
1. 推薦あるご
    1. 今：喫緊10見たのキャストが出てる他番組
    1. とりまシーズンで一個に
1. ログアウトページてきとう
1. ログインからの導線
1. カスタマイズ機能
1. ポップアップ系
1. スマホコピペ・アドレスバー(iosのwindow.scrollToがきかない，疑似的にスクロールアクションを起こせれば．)
1. js場所（トップページでツールバーずれることがある）
1. tv edit部分のタイトルかぶり処理，なぜかFlagがNULLになる
1. コード関数化してまとめる jsもgoも
1. androidでの見た目がmarital，ちょっと変かも
1. 各ユーザへの管理側からのメッセージ
1. カテゴリクリック・キャストクリックで検索
1. 寄付マーククリックアクション
1. 同一ユーザで複数ログイン時，いいね同時クリックで複数カウントされる．
1. googleアナリティクス
1. google広告挿入(すぐにしない)
1. ランキング（カテゴリ別・お気に入りポイント別・映画別・ドラマ別・）グラフ可視化
1. 各コメント欄で関連番組
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
1. その他サイト作成（CM・本・バラエティ・豆知識・スポーツ観戦・ニュース・不満・3大〇〇）

# Improve Coding
1. javascript 内で cookie が使えません！
1. 今どきの javascript の書き方
1. shijimi.com/tv/〇〇の tv を別サイトの場合は変えたいのだが，それって別の go を起動してもいける？

# 機能構成
1. admin用ログインページ tv/user/login_admin_page


# 知識いろいろ
1. vscode が便利
1. prettier で自動コード補正
1. ESlint でコード指南
1. travisいれたい
1. bee run -downdoc=true -gendoc=true でswagger発動
1. javascript → main.go内のgo関数で処理？
1. regexpは遅いらしい（多用部分には使わない）
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
1. heroku:60秒以内にサーバーを起動しないと落ちるため、多くのデータスクレイピングは不可能

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
1. c.Dataで引き継げない
    1. 小文字はまずい．大文字

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

# herokuへのアップ方法
```
heroku apps:destroy -a shijimi --confirm shijimi
git remote rm heroku


git clone git@github.com:shoot16625/shijimi.git
cd shijimi/Go

(
heroku login
heroku container:login
)
heroku create -a shijimi
heroku git:remote -a shijimi
heroku addons:add cleardb:ignite
heroku config | grep CLEARDB_DATABASE_URL


conf内：sqlconとprod変更
common.jsのURL変更
main.go：sqlconn変更/投入データ変更
models/comment.go heroku 時間で検索部分
comment/showタイムライン時間

heroku container:push web -a shijimi
heroku container:release web -a shijimi
heroku open
heroku logs --tail


Top@1060..cd
```