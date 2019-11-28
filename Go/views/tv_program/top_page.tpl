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
      <!-- <div class="background" style="background-color: white;"></div> -->
      <!-- <ons-toolbar class="toolbar" id="my-toolbar"></ons-toolbar> -->
      {{ template "/common/toolbar.tpl" . }}
      <ons-pull-hook id="pull-hook"></ons-pull-hook>
      {{ template "/common/alert.tpl" . }}
      <div class="list-margin">
        <ons-card
          style="text-align: center;background-color:linen;margin:10px;"
        >
          「ShiJimi」<br />
          ドラマ・映画の情報共有SNS
        </ons-card>
        <ons-row class="create-top-margin-20">
          <ons-col width="15%"></ons-col>
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
              onclick="dialogBoxEveryone('hint-dialog')"
              ><i class="far fa-question-circle hint-icon-right"></i
            ></ons-button>
          </ons-col>
        </ons-row>
        <ons-row>
          <ons-col>
            <p>
              <ons-button
                modifier="quiet"
                onclick="location.href='tv/tv_program/index'"
                >あなたにおすすめ</ons-button
              >
            </p>
            <p>
              <ons-button
                modifier="quiet"
                onclick="goOtherPage({{.UserId}}, 0, 'tv/tv_program/create_page')"
                >ドラマ・映画をつくる</ons-button
              >
            </p>
            <p>
              <ons-button
                modifier="quiet"
                onclick="location.href='tv/tv_program/comment/1'"
                >お問い合わせ</ons-button
              >
            </p>
          </ons-col>
          <ons-col>
            <p>
              <ons-button
                modifier="quiet"
                onclick="location.href='tv/user/create'"
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
                onclick="location.href='tv/user/logout'"
                >ログアウト</ons-button
              >
            </p>
          </ons-col>
        </ons-row>

        <ons-toast id="loginErrorToast" animation="ascend"
          >ログインに失敗しました <i class="far fa-sad-tear"></i
          ><button onclick="loginErrorToast.hide()">
            ok
          </button></ons-toast
        >
        <div class="on-air-drama">
          <h2>
            <i class="fas fa-tv" style="color: skyblue;"></i> 現在放送中のドラマ
          </h2>

          <p class="drama-on-air-carousel">
            <i class="far fa-moon" style="color:rgb(235, 200, 3);"></i> 月
          </p>
          <ons-carousel
            id="carousel01"
            auto-refresh
            auto-scroll
            auto-scroll-ratio="0.15"
            swipeable
            overscrollable
            item-width="200px"
            class="dramas-on-air"
          >
            {{ range.TvProgramMon }}
            <ons-carousel-item
              modifier="nodivider"
              id="{{.Id}}"
              name="{{.Title}}"
            >
              <div class="area-center drama-on-air-carousel-padding">
                <div class="thumbnail">
                  <img
                    src="{{.ImageUrl}}"
                    alt="{{.Title}}"
                    class="image-carousel"
                    onerror="this.src='/static/img/tv_img/hanko_02.png'"
                  />
                  <a href="/tv/tv_program/comment/{{.Id}}"></a>
                </div>
                {{if .ImageUrlReference}}
                <div class="reference">From:{{.ImageUrlReference}}</div>
                {{ end }}
                <div>{{.Title}}</div>
              </div>
            </ons-carousel-item>
            {{ end }}
          </ons-carousel>
          <p class="drama-on-air-carousel">
            <i class="fas fa-fire" style="color:rgb(235, 30, 30);"></i> 火
          </p>
          <ons-carousel
            id="carousel02"
            auto-refresh
            auto-scroll
            auto-scroll-ratio="0.15"
            swipeable
            overscrollable
            item-width="200px"
            class="dramas-on-air"
          >
            {{ range.TvProgramTue }}
            <ons-carousel-item
              modifier="nodivider"
              id="{{.Id}}"
              name="{{.Title}}"
            >
              <div class="area-center drama-on-air-carousel-padding">
                <div class="thumbnail">
                  <img
                    src="{{.ImageUrl}}"
                    alt="{{.Title}}"
                    class="image-carousel"
                    onerror="this.src='/static/img/tv_img/hanko_02.png'"
                  />
                  <a href="/tv/tv_program/comment/{{.Id}}"></a>
                </div>
                {{if .ImageUrlReference}}
                <div class="reference">From:{{.ImageUrlReference}}</div>
                {{ end }}
                <div>{{.Title}}</div>
              </div>
            </ons-carousel-item>
            {{ end }}
          </ons-carousel>
          <p class="drama-on-air-carousel">
            <i class="fas fa-tint" style="color:rgb(95, 149, 231);"></i> 水
          </p>
          <ons-carousel
            id="carousel03"
            auto-refresh
            auto-scroll
            auto-scroll-ratio="0.15"
            swipeable
            overscrollable
            item-width="200px"
            class="dramas-on-air"
          >
            {{ range.TvProgramWed }}
            <ons-carousel-item
              modifier="nodivider"
              id="{{.Id}}"
              name="{{.Title}}"
            >
              <div class="area-center drama-on-air-carousel-padding">
                <div class="thumbnail">
                  <img
                    src="{{.ImageUrl}}"
                    alt="{{.Title}}"
                    class="image-carousel"
                    onerror="this.src='/static/img/tv_img/hanko_02.png'"
                  />
                  <a href="/tv/tv_program/comment/{{.Id}}"></a>
                </div>
                {{if .ImageUrlReference}}
                <div class="reference">From:{{.ImageUrlReference}}</div>
                {{ end }}
                <div>{{.Title}}</div>
              </div>
            </ons-carousel-item>
            {{ end }}
          </ons-carousel>
          <p class="drama-on-air-carousel">
            <i class="fas fa-tree" style="color:green;"></i> 木
          </p>
          <ons-carousel
            id="carousel04"
            auto-refresh
            auto-scroll
            auto-scroll-ratio="0.15"
            swipeable
            overscrollable
            item-width="200px"
            class="dramas-on-air"
          >
            {{ range.TvProgramThu }}
            <ons-carousel-item
              modifier="nodivider"
              id="{{.Id}}"
              name="{{.Title}}"
            >
              <div class="area-center drama-on-air-carousel-padding">
                <div class="thumbnail">
                  <img
                    src="{{.ImageUrl}}"
                    alt="{{.Title}}"
                    class="image-carousel"
                    onerror="this.src='/static/img/tv_img/hanko_02.png'"
                  />
                  <a href="/tv/tv_program/comment/{{.Id}}"></a>
                </div>
                {{if .ImageUrlReference}}
                <div class="reference">From:{{.ImageUrlReference}}</div>
                {{ end }}
                <div>{{.Title}}</div>
              </div>
            </ons-carousel-item>
            {{ end }}
          </ons-carousel>
          <p class="drama-on-air-carousel">
            <i class="fas fa-coins" style="color:rgb(187, 162, 24);"></i> 金
          </p>
          <ons-carousel
            id="carousel05"
            auto-refresh
            auto-scroll
            auto-scroll-ratio="0.15"
            swipeable
            overscrollable
            item-width="200px"
            class="dramas-on-air"
          >
            {{ range.TvProgramFri }}
            <ons-carousel-item
              modifier="nodivider"
              id="{{.Id}}"
              name="{{.Title}}"
            >
              <div class="area-center drama-on-air-carousel-padding">
                <div class="thumbnail">
                  <img
                    src="{{.ImageUrl}}"
                    alt="{{.Title}}"
                    class="image-carousel"
                    onerror="this.src='/static/img/tv_img/hanko_02.png'"
                  />
                  <a href="/tv/tv_program/comment/{{.Id}}"></a>
                </div>
                {{if .ImageUrlReference}}
                <div class="reference">From:{{.ImageUrlReference}}</div>
                {{ end }}
                <div>{{.Title}}</div>
              </div>
            </ons-carousel-item>
            {{ end }}
          </ons-carousel>
          <p class="drama-on-air-carousel">
            <i class="fas fa-globe" style="color:rgb(138, 193, 219);"></i> 土
          </p>
          <ons-carousel
            id="carousel06"
            auto-refresh
            auto-scroll
            auto-scroll-ratio="0.15"
            swipeable
            overscrollable
            item-width="200px"
            class="dramas-on-air"
          >
            {{ range.TvProgramSat }}
            <ons-carousel-item
              modifier="nodivider"
              id="{{.Id}}"
              name="{{.Title}}"
            >
              <div class="area-center drama-on-air-carousel-padding">
                <div class="thumbnail">
                  <img
                    src="{{.ImageUrl}}"
                    alt="{{.Title}}"
                    class="image-carousel"
                    onerror="this.src='/static/img/tv_img/hanko_02.png'"
                  />
                  <a href="/tv/tv_program/comment/{{.Id}}"></a>
                </div>
                {{if .ImageUrlReference}}
                <div class="reference">From:{{.ImageUrlReference}}</div>
                {{ end }}
                <div>{{.Title}}</div>
              </div>
            </ons-carousel-item>
            {{ end }}
          </ons-carousel>
          <p class="drama-on-air-carousel">
            <i class="fas fa-sun" style="color:rgb(255, 166, 0);"></i> 日
          </p>
          <ons-carousel
            id="carousel07"
            auto-refresh
            auto-scroll
            auto-scroll-ratio="0.15"
            swipeable
            overscrollable
            item-width="200px"
            class="dramas-on-air"
          >
            {{ range.TvProgramSun }}
            <ons-carousel-item
              modifier="nodivider"
              id="{{.Id}}"
              name="{{.Title}}"
            >
              <div class="area-center drama-on-air-carousel-padding">
                <div class="thumbnail">
                  <img
                    src="{{.ImageUrl}}"
                    alt="{{.Title}}"
                    class="image-carousel"
                    onerror="this.src='/static/img/tv_img/hanko_02.png'"
                  />
                  <a href="/tv/tv_program/comment/{{.Id}}"></a>
                </div>
                {{if .ImageUrlReference}}
                <div class="reference">From:{{.ImageUrlReference}}</div>
                {{ end }}
                <div>{{.Title}}</div>
              </div>
            </ons-carousel-item>
            {{ end }}
          </ons-carousel>
        </div>
        <div class="on-air-movie">
          <h2>
            <i class="fas fa-film" style="color:chocolate;"></i> 最近の映画
          </h2>
          <p class="drama-on-air-carousel"></p>
          <ons-carousel
            id="carousel10"
            auto-refresh
            auto-scroll
            auto-scroll-ratio="0.15"
            swipeable
            overscrollable
            item-width="200px"
            class="dramas-on-air"
          >
            {{ range.TvProgramMovie }}
            <ons-carousel-item
              modifier="nodivider"
              id="{{.Id}}"
              name="{{.Title}}"
            >
              <div class="area-center drama-on-air-carousel-padding">
                <div class="thumbnail">
                  <img
                    src="{{.ImageUrl}}"
                    alt="{{.Title}}"
                    class="image-carousel"
                    onerror="this.src='/static/img/tv_img/hanko_02.png'"
                  />
                  <a href="/tv/tv_program/comment/{{.Id}}"></a>
                </div>
                {{if .ImageUrlReference}}
                <div class="reference">From:{{.ImageUrlReference}}</div>
                {{ end }}
                <div>{{.Title}}</div>
              </div>
            </ons-carousel-item>
            {{ end }}
          </ons-carousel>
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
            <ol>
              <li>
                ツールバーの<b>しじみ</b>をクリックすると、<br />トップページへ移動できます。
              </li>
              <li>
                <b>しじみ</b>の周辺をクリックすると、<br />上部へ移動できます。
              </li>
              <li>ユーザ登録をすれば、<br />誰でも番組の作成・編集が可能。</li>
              <li>
                検索：複数のキーワードで指定したいときは、スペースで区切ってください。
              </li>
              <li>機能は今後も追加していく予定です。</li>
              <li>バグ：→ お問い合わせへポスト</li>
              <li>アドレスバーの消し方(iphone)：アドレスバーを上スワイプ</li>
              <li>ツールバーを横スワイプ：ダイアログが隠れます。</li>
            </ol>
          </div>
        </ons-page>
        <script>
          hideSwipeToolbar('hint-hide-swipe', 'hint-dialog');
        </script>
      </ons-dialog>
    </template>
    {{ template "/common/js.tpl" . }}
    <script type="text/javascript">
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

    <!-- 保留 -->
    <!-- <script>
      ons.ready(function(){
        document.getElementById("my-toolbar").innerHTML= '<div class="left" id="mypage-toolbar"><ons-toolbar-button icon="md-face" style="font-size:24px;" onclick="location.href=\'/tv/user/show\'"></ons-toolbar-button></div><div class="center" id="image-toolbar"><div class="area-center"><img src="/static/img/shijimi-transparence.png" alt="shijimi" height="42px;" onclick="location.href=\'/\'"/></div></div><div class="right"><ons-toolbar-button icon="fa-search" onclick="dialogBoxEveryone(\'search-toolbar\')"></ons-toolbar-button></div>';
      })
    </script> -->
    <script>
        // ログイン失敗ポップアップ
        if ({{.LoginError}}) {
          document.querySelector('ons-toast').show();
          setTimeout(function() {
            document.querySelector('ons-toast').hide();
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
