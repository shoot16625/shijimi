<!DOCTYPE html>
<html lang="ja">
  <head>
    {{ template "/common/header.tpl" . }}
    {{ template "/common/alert.tpl" . }}
  </head>

  <body>
    <ons-page>
      {{ template "/common/toolbar.tpl" . }}
      <form id="login-user" action="/tv/user/login_admin" method="post">
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
              float
              required
            ></ons-input>
          </p>
          <p>
            <ons-input
              name="key"
              modifier="underbar"
              type="password"
              placeholder="キー"
              minlength="8"
              maxlength="50"
              float
              required
            ></ons-input>
          </p>
          <p style="margin: 30px;">
            <button class="button button--outline">login</button>
          </p>
        </div>
      </form>
    </ons-page>

    {{ template "/common/js.tpl" . }}
  </body>
</html>
