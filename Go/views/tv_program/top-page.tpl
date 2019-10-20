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
      <div style="margin-left: 5px; margin-right: 5px;">
        <ons-card
          style="text-align: center;background-color:linen;margin:10px;"
        >
          「ShiJimi」<br />
          SNSとレビューサイトを足して2で割ったような．
        </ons-card>
        <form id="search_form" action="/tv/tv_program/search" method="post">
          <p style="margin-top: 20px;">
            <ons-search-input
              name="search_word"
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
                onclick="GoOtherPage({{.UserId}}, 'tv/tv_program/create_page')"
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
                onclick="DialogBoxEveryone('login_dialog')"
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
            <i class="far fa-moon" style="color:gold;"></i> 月
          </p>
          <ons-carousel
            id="carousel01"
            auto-refresh
            auto-scroll
            auto-scroll-ratio="0.15"
            swipeable
            overscrollable
            item-width="200px"
            class="doramas_on_air"
            style="margin-bottom: 10px;"
          >
            {{ range.TvProgram_mon }}
            <ons-carousel-item
              modifier="nodivider"
              class="tv_program"
              name="{{.Title}}"
            >
              <div style="padding: 3px; text-align: center;">
                <div class="thumbnail">
                  <img
                    src="{{.ImageUrl}}"
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
            <i class="fas fa-fire" style="color:red;"></i> 火
          </p>
          <ons-carousel
            id="carousel02"
            auto-refresh
            auto-scroll
            auto-scroll-ratio="0.15"
            swipeable
            overscrollable
            item-width="200px"
            class="doramas_on_air"
            style="margin-bottom: 10px;"
          >
            {{ range.TvProgram_tue }}
            <ons-carousel-item
              modifier="nodivider"
              class="tv_program"
              name="{{.Title}}"
            >
              <div style="padding: 3px; text-align: center;">
                <div class="thumbnail">
                  <img
                    src="{{.ImageUrl}}"
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
            <i class="fas fa-tint" style="color:blue;"></i> 水
          </p>
          <ons-carousel
            id="carousel03"
            auto-refresh
            auto-scroll
            auto-scroll-ratio="0.15"
            swipeable
            overscrollable
            item-width="200px"
            class="doramas_on_air"
            style="margin-bottom: 10px;"
          >
            {{ range.TvProgram_wed }}
            <ons-carousel-item
              modifier="nodivider"
              class="tv_program"
              name="{{.Title}}"
            >
              <div style="padding: 3px; text-align: center;">
                <div class="thumbnail">
                  <img
                    src="{{.ImageUrl}}"
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
            class="doramas_on_air"
            style="margin-bottom: 10px;"
          >
            {{ range.TvProgram_thu }}
            <ons-carousel-item
              modifier="nodivider"
              class="tv_program"
              name="{{.Title}}"
            >
              <div style="padding: 3px; text-align: center;">
                <div class="thumbnail">
                  <img
                    src="{{.ImageUrl}}"
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
            <i class="fas fa-coins" style="color:gold;"></i> 金
          </p>
          <ons-carousel
            id="carousel05"
            auto-refresh
            auto-scroll
            auto-scroll-ratio="0.15"
            swipeable
            overscrollable
            item-width="200px"
            class="doramas_on_air"
            style="margin-bottom: 10px;"
          >
            {{ range.TvProgram_fri }}
            <ons-carousel-item
              modifier="nodivider"
              class="tv_program"
              name="{{.Title}}"
            >
              <div style="padding: 3px; text-align: center;">
                <div class="thumbnail">
                  <img
                    src="{{.ImageUrl}}"
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
            <i class="fas fa-globe" style="color:rgb(116, 168, 192);"></i> 土
          </p>
          <ons-carousel
            id="carousel06"
            auto-refresh
            auto-scroll
            auto-scroll-ratio="0.15"
            swipeable
            overscrollable
            item-width="200px"
            class="doramas_on_air"
            style="margin-bottom: 10px;"
          >
            {{ range.TvProgram_sat }}
            <ons-carousel-item
              modifier="nodivider"
              class="tv_program"
              name="{{.Title}}"
            >
              <div style="padding: 3px; text-align: center;">
                <div class="thumbnail">
                  <img
                    src="{{.ImageUrl}}"
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
            <i class="fas fa-sun" style="color:gold;"></i> 日
          </p>
          <ons-carousel
            id="carousel07"
            auto-refresh
            auto-scroll
            auto-scroll-ratio="0.15"
            swipeable
            overscrollable
            item-width="200px"
            class="doramas_on_air"
            style="margin-bottom: 10px;"
          >
            {{ range.TvProgram_sun }}
            <ons-carousel-item
              modifier="nodivider"
              class="tv_program"
              name="{{.Title}}"
            >
              <div style="padding: 3px; text-align: center;">
                <div class="thumbnail">
                  <img
                    src="{{.ImageUrl}}"
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

    <template id="login_dialog.html">
      <ons-dialog id="login_dialog" modifier="large" cancelable fullscreen>
        <ons-page>
          <ons-toolbar>
            <div class="left">
              <ons-button
                id="cancel_button"
                onclick="hideAlertDialog('login_dialog')"
              >
                <i class="fas fa-window-close"></i>
              </ons-button>
            </div>
            <div class="center">ログイン</div>
          </ons-toolbar>
          <form id="login_user" action="/tv/user/login" method="post">
            <div style="text-align: center; margin-top: 30px;">
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
          <p style="text-align: right;">
            <a href="tv/user/forget_username_page">ユーザー名を忘れたら...</a>
          </p>
          <p style="text-align: right;">
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
        AutoScroll(carousel01, {{.TvProgram_mon}}.length);
      }
      if ({{.TvProgram_tue}}) {
        AutoScroll(carousel02, {{.TvProgram_tue}}.length);
      }
      if ({{.TvProgram_wed}}) {
        AutoScroll(carousel03, {{.TvProgram_wed}}.length);
      }
      if({{.TvProgram_thu}}){
        AutoScroll(carousel04, {{.TvProgram_thu}}.length);
      }
      if ({{.TvProgram_fri}}) {
        AutoScroll(carousel05, {{.TvProgram_fri}}.length);
      }
      if ({{.TvProgram_sat}}) {
        AutoScroll(carousel06, {{.TvProgram_sat}}.length);
      }
      if ({{.TvProgram_sun}}) {
        AutoScroll(carousel07, {{.TvProgram_sun}}.length);
      }
    </script>
  </body>
</html>
