<!DOCTYPE html>
<html lang="ja">
  <head>
    {{ template "/common/header.tpl" . }}
  </head>

  <body>
    <ons-page>
      {{ template "/common/toolbar.tpl" . }}
      {{ template "/common/alert.tpl" . }}

      <form id="login-user" action="/tv/user/forget_password" method="post">
        <div class="input-table">
          <p>
            <ons-input
              name="username"
              modifier="underbar"
              placeholder="ユーザー名"
              float
              required
            ></ons-input>
          </p>
          <p>
            <label for="SecondPassword" style="margin: 0 30px 0 30px;"
              >＜第2パスワード＞</label
            >
            <ons-input
              id="SecondPassword"
              name="SecondPassword"
              modifier="underbar"
              placeholder="あなたの小学校の名前は?"
              maxlength="100"
              float
              required
            ></ons-input>
          </p>
          <p class="create-top-margin">
            <button class="button button--outline">パスワードを再設定</button>
          </p>
        </div>
      </form>
      <template id="alert-username-not-found.html">
        <ons-alert-dialog id="alert-username-not-found" modifier="rowfooter">
          <div class="alert-dialog-title">Alert</div>
          <div class="alert-dialog-content">
            ユーザー情報が誤っています。
          </div>
          <div class="alert-dialog-footer">
            <ons-alert-dialog-button
              onclick="hideAlertDialog('alert-username-not-found')"
              >OK</ons-alert-dialog-button
            >
          </div>
        </ons-alert-dialog>
      </template>
    </ons-page>
    <script type="text/javascript" src="/static/js/common.js"></script>
    <script type="text/javascript">
      const name = {{.User}};
      if (name != null){
        dialogBoxEveryone('alert-username-not-found');
      }
    </script>
  </body>
</html>
