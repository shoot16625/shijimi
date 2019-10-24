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
      {{ template "/common/alert.tpl" . }}
      <div class="list-margin">
        <ons-card
          style="text-align: center;background-color:linen;margin:10px;"
        >
          「ShiJimi」<br />
          SNSとレビューサイトを足して<br>2で割ったような．
        </ons-card>
        <form id="search_form" action="/tv/tv_program/search" method="post">
          <p style="margin-top: 20px;">
            <ons-search-input
              name="search-word"
              placeholder="ドラマ・映画を検索"
            ></ons-search-input>
          </p>
        </form>
        <ons-row>
          <ons-col>
            <p>
              <ons-button
                modifier="quiet"
                onclick="location.href='tv/tv_program/index'"
                >すべてのドラマ</ons-button
              >
            </p>
            <p>
              <ons-button
                modifier="quiet"
                onclick="goOtherPage({{.UserId}}, 'tv/tv_program/create_page')"
                >ドラマ・映画をつくる</ons-button
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

        <div class="on_airs">
          <h2>
            <i class="fas fa-tv" style="color: skyblue;"></i> 現在放送中のドラマ
          </h2>

          <p
            style="margin-top: 25px;border-bottom: solid thin lightgray;text-align: left;"
          >
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
            class="doramas-on-air"
          >
            {{ range.TvProgram_mon }}
            <ons-carousel-item
              modifier="nodivider"
              id="{{.Id}}"
              name="{{.Title}}"
            >
              <div class="area-center" style="padding: 3px;">
                <div class="thumbnail">
                  <img
                    src="{{.ImageURL}}"
                    alt="{{.Title}}"
                    style="width: 80%"
                  />
                  <a href="/tv/tv_program/comment/{{.Id}}"></a>
                </div>
                <div>{{.Title}}</div>
              </div>
            </ons-carousel-item>
            {{ end }}
          </ons-carousel>
          <p
            style="margin-top: 25px;border-bottom: solid thin lightgray;text-align: left;"
          >
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
            class="doramas-on-air"
          >
            {{ range.TvProgram_tue }}
            <ons-carousel-item
              modifier="nodivider"
              id="{{.Id}}"
              name="{{.Title}}"
            >
              <div class="area-center" style="padding: 3px;">
                <div class="thumbnail">
                  <img
                    src="{{.ImageURL}}"
                    alt="{{.Title}}"
                    style="width: 80%"
                  />
                  <a href="/tv/tv_program/comment/{{.Id}}"></a>
                </div>
                <div>{{.Title}}</div>
              </div>
            </ons-carousel-item>
            {{ end }}
          </ons-carousel>
          <p
            style="margin-top: 25px;border-bottom: solid thin lightgray;text-align: left;"
          >
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
            class="doramas-on-air"
          >
            {{ range.TvProgram_wed }}
            <ons-carousel-item
              modifier="nodivider"
              id="{{.Id}}"
              name="{{.Title}}"
            >
              <div class="area-center" style="padding: 3px;">
                <div class="thumbnail">
                  <img
                    src="{{.ImageURL}}"
                    alt="{{.Title}}"
                    style="width: 80%"
                  />
                  <a href="/tv/tv_program/comment/{{.Id}}"></a>
                </div>
                <div>{{.Title}}</div>
              </div>
            </ons-carousel-item>
            {{ end }}
          </ons-carousel>
          <p
            style="margin-top: 25px;border-bottom: solid thin lightgray;text-align: left;"
          >
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
            class="doramas-on-air"
          >
            {{ range.TvProgram_thu }}
            <ons-carousel-item
              modifier="nodivider"
              id="{{.Id}}"
              name="{{.Title}}"
            >
              <div class="area-center" style="padding: 3px;">
                <div class="thumbnail">
                  <img
                    src="{{.ImageURL}}"
                    alt="{{.Title}}"
                    style="width: 80%"
                  />
                  <a href="/tv/tv_program/comment/{{.Id}}"></a>
                </div>
                <div>{{.Title}}</div>
              </div>
            </ons-carousel-item>
            {{ end }}
          </ons-carousel>
          <p
            style="margin-top: 25px;border-bottom: solid thin lightgray;text-align: left;"
          >
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
            class="doramas-on-air"
          >
            {{ range.TvProgram_fri }}
            <ons-carousel-item
              modifier="nodivider"
              id="{{.Id}}"
              name="{{.Title}}"
            >
              <div class="area-center" style="padding: 3px;">
                <div class="thumbnail">
                  <img
                    src="{{.ImageURL}}"
                    alt="{{.Title}}"
                    style="width: 80%"
                  />
                  <a href="/tv/tv_program/comment/{{.Id}}"></a>
                </div>
                <div>{{.Title}}</div>
              </div>
            </ons-carousel-item>
            {{ end }}
          </ons-carousel>
          <p
            style="margin-top: 25px;border-bottom: solid thin lightgray;text-align: left;"
          >
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
            class="doramas-on-air"
          >
            {{ range.TvProgram_sat }}
            <ons-carousel-item
              modifier="nodivider"
              id="{{.Id}}"
              name="{{.Title}}"
            >
              <div class="area-center" style="padding: 3px;">
                <div class="thumbnail">
                  <img
                    src="{{.ImageURL}}"
                    alt="{{.Title}}"
                    style="width: 80%"
                  />
                  <a href="/tv/tv_program/comment/{{.Id}}"></a>
                </div>
                <div>{{.Title}}</div>
              </div>
            </ons-carousel-item>
            {{ end }}
          </ons-carousel>
          <p
            style="margin-top: 25px;border-bottom: solid thin lightgray;text-align: left;"
          >
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
            class="doramas-on-air"
          >
            {{ range.TvProgram_sun }}
            <ons-carousel-item
              modifier="nodivider"
              id="{{.Id}}"
              name="{{.Title}}"
            >
              <div class="area-center" style="padding: 3px;">
                <div class="thumbnail">
                  <img
                    src="{{.ImageURL}}"
                    alt="{{.Title}}"
                    style="width: 80%"
                  />
                  <a href="/tv/tv_program/comment/{{.Id}}"></a>
                </div>
                <div>{{.Title}}</div>
              </div>
            </ons-carousel-item>
            {{ end }}
          </ons-carousel>
        </div>
      </div>
    </ons-page>

    <template id="login-dialog.html">
      <ons-dialog id="login-dialog" modifier="large" cancelable fullscreen>
        <ons-page>
          <ons-toolbar>
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
              <p>
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
        </script>
      </ons-dialog>
    </template>

    <script type="text/javascript" src="/static/js/common.js"></script>
    <script type="text/javascript">
      if({{.TvProgram_mon}}){
        autoScroll(carousel01, {{.TvProgram_mon}}.length);
      }
      if ({{.TvProgram_tue}}) {
        autoScroll(carousel02, {{.TvProgram_tue}}.length);
      }
      if ({{.TvProgram_wed}}) {
        autoScroll(carousel03, {{.TvProgram_wed}}.length);
      }
      if({{.TvProgram_thu}}){
        autoScroll(carousel04, {{.TvProgram_thu}}.length);
      }
      if ({{.TvProgram_fri}}) {
        autoScroll(carousel05, {{.TvProgram_fri}}.length);
      }
      if ({{.TvProgram_sat}}) {
        autoScroll(carousel06, {{.TvProgram_sat}}.length);
      }
      if ({{.TvProgram_sun}}) {
        autoScroll(carousel07, {{.TvProgram_sun}}.length);
      }
    </script>
  </body>
</html>
