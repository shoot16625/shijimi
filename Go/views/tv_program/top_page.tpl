<!DOCTYPE html>
<html lang="ja">
  <head>
    {{ template "/common/header.tpl" . }}

    <style type="text/css">
      p {
        text-align: center;
        margin: 5px;
      }
    </style>
  </head>
  <body>
    <ons-page id="top-page">
      {{ template "/common/toolbar.tpl" . }}
      <ons-pull-hook id="pull-hook"></ons-pull-hook>
      {{ template "/common/alert.tpl" . }}
      <div class="list-margin">
        <ons-card style="background-color:linen;padding: 10px;">
          <p>「ShiJimi」</p>
          <p>ドラマ・映画の情報共有SNS</p>
          <div class="area-right-float hint-position-relative">
            <ons-button
              modifier="quiet"
              onclick="dialogBoxEveryone('hint-dialog')"
              ><i class="far fa-question-circle hint-icon-right"></i
            ></ons-button>
          </div>
        </ons-card>
        <ons-row>
          <ons-col width="15%" class="area-center"> </ons-col>
          <ons-col width="70%" style="text-align: center;">
            <form id="search_form" action="/tv/tv_program/search" method="post">
              <ons-search-input
                name="search-word"
                placeholder="ドラマ・映画を検索"
              ></ons-search-input>
            </form>
          </ons-col>
          <ons-col>
            <ons-button
              modifier="quiet"
              onclick="location.href='/tv/tv_program/comment/2'"
              ><i class="far fa-envelope buruburu" style="color: darkgray;"></i
            ></ons-button>
          </ons-col>
        </ons-row>
        <ons-row>
          <ons-col>
            <p>
              <ons-button
                modifier="quiet"
                onclick="location.href='/tv/tv_program/index'"
                >あなたにおすすめ</ons-button
              >
            </p>
            <p>
              <ons-button
                modifier="quiet"
                onclick="goOtherPage({{.User.Id}}, 0, '/tv/tv_program/create_page')"
                >ドラマ・映画をつくる</ons-button
              >
            </p>
            <p>
              <ons-button
                modifier="quiet"
                onclick="location.href='/tv/tv_program/comment/1'"
                >お問い合わせ</ons-button
              >
            </p>
          </ons-col>
          <ons-col>
            <p>
              <ons-button
                modifier="quiet"
                onclick="goOtherPage({{.User.Id}}, 0,'/tv/user/create')"
                >新規登録</ons-button
              >
            </p>
            <p>
              <ons-button
                modifier="quiet"
                onclick="dialogBoxEveryone('login-dialog')"
                >ログイン</ons-button
              >
            </p>
            <p>
              <ons-button
                modifier="quiet"
                onclick="location.href='/tv/user/logout'"
                >ログアウト</ons-button
              >
            </p>
          </ons-col>
        </ons-row>

        <div class="on-air-drama" id="on-air-drama"></div>
        <div class="on-air-movie" id="on-air-movie"></div>

        <div class="floating-bottom">
          <div class="toast">
            <div class="toast__message">
              {{ .LoginErrorStatus }} <i class="far fa-sad-tear"></i>
              <button
                class="toast-hide-button"
                onclick="hideToast('.floating-bottom')"
              >
                ok
              </button>
            </div>
          </div>
        </div>

        <ons-button
          modifier="quiet"
          class="add-button comment-list-content-font"
          >ホーム画面にインストールする</ons-button
        >
        <ons-row>
          <ons-col class="area-right">
            <a
              href="https://twitter.com/share"
              class="twitter-share-button"
              data-url="https://shijimi.herokuapp.com/"
              data-text="ドラマ・映画の情報共有SNS"
              data-hashtags="shijimi"
              >Tweet</a
            >
          </ons-col>
          <ons-col>
            <div
              class="fb-share-button"
              data-href="https://shijimi.herokuapp.com/"
              data-layout="button_count"
              data-size="small"
            >
              <a
                target="_blank"
                href="https://www.facebook.com/sharer/sharer.php?u=https%3A%2F%2Fdevelopers.facebook.com%2Fdocs%2Fplugins%2F&amp;src=sdkpreparse"
                class="fb-xfbml-parse-ignore"
                >シェア</a
              >
            </div>
          </ons-col>
        </ons-row>

        <div class="area-right">
          <ons-button
            modifier="quiet"
            id="consent-button"
            onclick="dialogBoxEveryone('terms-of-service')"
            >利用規約</ons-button
          >
        </div>
        <div class="area-right">
          <ons-button
            modifier="quiet"
            id="privacy-button"
            onclick="dialogBoxEveryone('privacy-policy')"
            >プライバシーポリシー</ons-button
          >
        </div>
        <div
          class="area-right"
          style="font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;margin: 10px;"
        >
          &copy; 2020 ShiJimi
        </div>
      </div>
    </ons-page>
    {{ template "/common/js.tpl" . }}
    <template id="terms-of-service.html">
      <ons-dialog id="terms-of-service" modifier="large" cancelable fullscreen>
        <ons-page>
          <ons-toolbar id="terms-of-service-hide-swipe">
            <div class="left">
              <ons-button
                id="cancel-button"
                onclick="hideAlertDialog('terms-of-service')"
              >
                <i class="fas fa-window-close"></i>
              </ons-button>
            </div>
            <div class="center">利用規約</div>
          </ons-toolbar>
          <div class="scroller list-margin">
            {{ template "/common/terms_of_service.tpl" . }}
          </div>
        </ons-page>
        <script>
          hideSwipeToolbar('terms-of-service-hide-swipe', 'terms-of-service');
        </script>
      </ons-dialog>
    </template>
    <template id="privacy-policy.html">
      <ons-dialog id="privacy-policy" modifier="large" cancelable fullscreen>
        <ons-page>
          <ons-toolbar id="privacy-policy-hide-swipe">
            <div class="left">
              <ons-button
                id="cancel-button"
                onclick="hideAlertDialog('privacy-policy')"
              >
                <i class="fas fa-window-close"></i>
              </ons-button>
            </div>
            <div class="center">プライバシーポリシー</div>
          </ons-toolbar>
          <div class="scroller list-margin">
            {{ template "/common/privacy_policy.tpl" . }}
          </div>
        </ons-page>
        <script>
          hideSwipeToolbar('privacy-policy-hide-swipe', 'privacy-policy');
        </script>
      </ons-dialog>
    </template>
    <template id="login-dialog.html">
      <ons-dialog id="login-dialog" modifier="large" cancelable fullscreen>
        <ons-page>
          <ons-toolbar id="login-hide-swipe">
            <div class="left">
              <ons-button
                id="cancel-button"
                onclick="hideAlertDialog('login-dialog')"
              >
                <i class="fas fa-window-close"></i>
              </ons-button>
            </div>
            <div class="center">ログイン</div>
          </ons-toolbar>
          <form id="login-user" action="/tv/user/login" method="post">
            <div class="input-table">
              <p class="margin-bottom-20">
                <ons-input
                  name="username"
                  minlength="2"
                  maxlength="20"
                  modifier="underbar"
                  placeholder="ユーザー名"
                  float
                  required
                ></ons-input>
              </p>
              <p>
                <ons-input
                  name="password"
                  modifier="underbar"
                  type="password"
                  placeholder="パスワード"
                  minlength="8"
                  maxlength="50"
                  id="password"
                  float
                  required
                ></ons-input>
              </p>
              <p style="margin-top:20px;">
                <label class="left">
                  <ons-checkbox input-id="password-check"></ons-checkbox>
                </label>
                <label for="password-check" class="center">
                  パスワードを表示
                </label>
              </p>
              <p style="margin: 30px;">
                <button class="button button--outline">login</button>
              </p>
            </div>
          </form>
          <p class="area-right">
            <a href="tv/user/forget_username_page">ユーザー名を忘れたら...</a>
          </p>
          <p class="area-right">
            <a href="tv/user/forget_password_page">パスワードを忘れたら...</a>
          </p>
        </ons-page>
        <script type="text/javascript">
          $(function() {
            $('#password-check').change(function() {
              if ($(this).prop('checked')) {
                $('#password').attr('type', 'text');
              } else {
                $('#password').attr('type', 'password');
              }
            });
          });

          hideSwipeToolbar('login-hide-swipe', 'login-dialog');
        </script>
      </ons-dialog>
    </template>
    <!-- ヒントページ -->
    <template id="hint-dialog.html">
      <ons-dialog id="hint-dialog" modifier="large" cancelable fullscreen>
        <ons-page>
          <ons-toolbar id="hint-hide-swipe">
            <div class="left">
              <ons-button
                id="cancel-button"
                onclick="hideAlertDialog('hint-dialog')"
              >
                <i class="fas fa-window-close"></i>
              </ons-button>
            </div>
            <div class="center">
              ヒント <i class="far fa-question-circle hint-font-size"></i>
            </div>
          </ons-toolbar>
          <div class="scroller list-margin">
            <ul class="list">
              <li class="list-header">
                ShiJimiの使い方
              </li>
              <li class="list-item hint-list-dialog">
                <div class="list-item__center">
                  ユーザー登録をすることで、作品登録・ブックマーク・ツイートなどの機能が利用できるようになります。
                </div>
              </li>
              <li class="list-item hint-list-dialog">
                <div class="list-item__center">
                  日記感覚で利用することをおすすめいたします。
                </div>
              </li>
              <li class="list-item">
                <div class="list-item__center hint-list-dialog">
                  順次、機能を追加していきますので、気長にお待ち下さい。
                </div>
              </li>
            </ul>
            <ul class="list">
              <li class="list-header">
                ツールバー
              </li>
              <li class="list-item hint-list-dialog">
                <div class="list-item__center">
                  しじみをクリック → 最初のページへ
                </div>
              </li>
              <li class="list-item">
                <div class="list-item__center hint-list-dialog">
                  しじみの周辺をクリック → 上部へスクロール移動
                </div>
              </li>
              <li class="list-item">
                <div class="list-item__center hint-list-dialog">
                  横スワイプ → ダイアログを閉じる
                </div>
              </li>
            </ul>
            <ul class="list">
              <li class="list-header">
                機能
              </li>
              <li class="list-item">
                <div class="list-item__center hint-list-dialog">
                  検索：複数のキーワード指定 → スペース区切り
                </div>
              </li>
              <li class="list-item">
                <div class="list-item__center hint-list-dialog">
                  バグ：<a href="tv/tv_program/comment/1"> お問い合わせ</a
                  >へポストお願いします。
                </div>
              </li>
              <li class="list-item">
                <div class="list-item__center hint-list-dialog">
                  アドレスバーの消し方(iphone)：アドレスバーを上スワイプ
                </div>
              </li>
            </ul>
          </div>
        </ons-page>
        <script>
          hideSwipeToolbar('hint-hide-swipe', 'hint-dialog');
        </script>
      </ons-dialog>
    </template>

    <script>
      ons.ready(function() {
        let elem = document.getElementById('on-air-drama');
        let text =
          '<h2><i class="fas fa-tv" style="color: skyblue;"></i> 現在放送中のドラマ</h2>';

        let subHeader = ['<i class="far fa-moon" style="color:rgb(235, 200, 3);"></i> 月','<i class="fas fa-fire" style="color:rgb(235, 30, 30);"></i> 火','<i class="fas fa-tint" style="color:rgb(95, 149, 231);"></i> 水','<i class="fas fa-tree" style="color:green;"></i> 木','<i class="fas fa-coins" style="color:rgb(187, 162, 24);"></i> 金','<i class="fas fa-globe" style="color:rgb(138, 193, 219);"></i> 土','<i class="fas fa-sun" style="color:rgb(255, 166, 0);"></i> 日'];

        let tvPrograms = [{{ .TvProgramMon }}, {{ .TvProgramTue }}, {{ .TvProgramWed }}, {{ .TvProgramThu }}, {{ .TvProgramFri }}, {{ .TvProgramSat }}, {{ .TvProgramSun }}];

        for (let index = 0; index < tvPrograms.length; index++) {
          text += '<p class="drama-on-air-carousel">'+subHeader[index]+'</p>';
          if (tvPrograms[index]) {
            text += '<ons-carousel id="carousel0' + String(index+1) + '" auto-refresh auto-scroll auto-scroll-ratio="0.15" swipeable overscrollable item-width="200px" class="dramas-on-air">';
            for (let i = 0; i < tvPrograms[index].length; i++) {
              let imageUrlReference = "";
              if (tvPrograms[index][i].ImageUrlReference) {
                imageUrlReference = '<div class="reference">From:' + tvPrograms[index][i].ImageUrlReference + '</div>';
              }

              text += '<ons-carousel-item modifier="nodivider" id="' + tvPrograms[index][i].Id + '" name="' + tvPrograms[index][i].Title + '"><div class="area-center drama-on-air-carousel-padding"><div class="thumbnail"><img data-src="' + tvPrograms[index][i].ImageUrl + '" alt="' + tvPrograms[index][i].Title + '" class="image-carousel lazyload" onerror="this.src=\'/static/img/tv_img/hanko_02.png\'"/><a href="/tv/tv_program/comment/' + tvPrograms[index][i].Id + '"></a></div>' + imageUrlReference + '<div>' + tvPrograms[index][i].Title + '</div></div></ons-carousel-item>';
            }
            text += '</ons-carousel>';
          }
        }
        elem.innerHTML = text;

        elem = document.getElementById('on-air-movie');
        text = '<h2><i class="fas fa-film" style="color:chocolate;"></i> 最近の映画</h2>';
        let tvProgram = {{ .TvProgramMovie }};
        if (tvProgram) {
          text += '<ons-carousel id="carousel10" auto-refresh auto-scroll auto-scroll-ratio="0.15" swipeable overscrollable item-width="200px" class="dramas-on-air">';
          for (let i = 0; i < tvProgram.length; i++) {
            let imageUrlReference = "";
            if (tvProgram[i].ImageUrlReference) {
              imageUrlReference = '<div class="reference">From:' + tvProgram[i].ImageUrlReference + '</div>';
            }
            let imageURL = tvProgram[i].ImageUrl;
            imageURL = imageURL.replace("w=300", "w=200");
            text += '<ons-carousel-item modifier="nodivider" id="' + tvProgram[i].Id + '" name="' + tvProgram[i].Title + '"><div class="area-center drama-on-air-carousel-padding"><div class="thumbnail"><img data-src="' + imageURL + '" alt="' + tvProgram[i].Title + '" class="image-carousel lazyload" onerror="this.src=\'/static/img/tv_img/hanko_02.png\'"/><a href="/tv/tv_program/comment/' + tvProgram[i].Id + '"></a></div>' + imageUrlReference + '<div>' + tvProgram[i].Title + '</div></div></ons-carousel-item>';
          }
          text += '</ons-carousel>';
        }
        elem.innerHTML = text;
      });

      ons.ready(function() {
        if({{.TvProgramMon}}){
          autoScroll(carousel01, {{.TvProgramMon}}.length);
        }
        if ({{.TvProgramTue}}) {
          autoScroll(carousel02, {{.TvProgramTue}}.length);
        }
        if ({{.TvProgramWed}}) {
          autoScroll(carousel03, {{.TvProgramWed}}.length);
        }
        if({{.TvProgramThu}}){
          autoScroll(carousel04, {{.TvProgramThu}}.length);
        }
        if ({{.TvProgramFri}}) {
          autoScroll(carousel05, {{.TvProgramFri}}.length);
        }
        if ({{.TvProgramSat}}) {
          autoScroll(carousel06, {{.TvProgramSat}}.length);
        }
        if ({{.TvProgramSun}}) {
          autoScroll(carousel07, {{.TvProgramSun}}.length);
        }
        if ({{.TvProgramMovie}}) {
          autoScroll(carousel10, {{.TvProgramMovie}}.length);
        }
      });

        // ツイッターガジェット
        !(function(d, s, id) {
          var js,
            fjs = d.getElementsByTagName(s)[0],
            p = /^http:/.test(d.location) ? 'http' : 'https';
          if (!d.getElementById(id)) {
            js = d.createElement(s);
            js.id = id;
            js.src = p + '://platform.twitter.com/widgets.js';
            fjs.parentNode.insertBefore(js, fjs);
          }
        })(document, 'script', 'twitter-wjs');
    </script>
    <div id="fb-root"></div>
    <script
      async
      defer
      crossorigin="anonymous"
      src="https://connect.facebook.net/ja_JP/sdk.js#xfbml=1&version=v4.0"
    ></script>

    <script>
        // ログイン失敗ポップアップ
        if ({{.LoginError}}) {
          $(".floating-bottom").fadeIn();
          setTimeout(function() {
            $(".floating-bottom").fadeOut();
          }, 4000);
        }

      // serviceEorkerの設置
      window.addEventListener('load', function() {
        if ('serviceWorker' in navigator) {
          navigator.serviceWorker
            .register('/serviceWorker.js')
            .then(function(registration) {
              console.log('serviceWorker registed.');
            })
            .catch(function(error) {
              console.warn('serviceWorker error.', error);
            });
        }
      });

      // インストールボタンの機能(たぶんchromeのみ，localhostかhttpsでないと表示されない)
      let deferredPrompt;
      const addBtn = document.querySelector('.add-button');
      addBtn.style.display = 'none';
      window.addEventListener('beforeinstallprompt', e => {
        // Prevent Chrome 67 and earlier from automatically showing the prompt
        e.preventDefault();
        // Stash the event so it can be triggered later.
        deferredPrompt = e;
        // Update UI to notify the user they can add to home screen
        addBtn.style.display = 'block';

        addBtn.addEventListener('click', e => {
          // hide our user interface that shows our A2HS button
          addBtn.style.display = 'none';
          // Show the prompt
          deferredPrompt.prompt();
          // Wait for the user to respond to the prompt
          deferredPrompt.userChoice.then(choiceResult => {
            if (choiceResult.outcome === 'accepted') {
              console.log('User accepted the A2HS prompt');
            } else {
              console.log('User dismissed the A2HS prompt');
            }
            deferredPrompt = null;
          });
        });
      });
    </script>
  </body>
</html>
