<!DOCTYPE html>
<html lang="ja">
  <head>
    {{ template "/common/header.tpl" . }}
  </head>

  <body>
    <ons-page>
      {{ template "/common/toolbar.tpl" . }}
      {{ template "/common/alert.tpl" . }}

      <form id="login-user" action="/tv/user/forget_username" method="post">
        <div class="input-table">
          <p>
            <label for="password" style="margin: 0 30px 0 30px;"
              >＜パスワード＞</label
            >
            <ons-input
              name="password"
              modifier="underbar"
              type="password"
              placeholder="パスワード"
              minlength="8"
              id="password"
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
              placeholder="あなたの小学校の名前は?
          "
              maxlength="100"
              float
              required
            ></ons-input>
          </p>
          <p>
            <label class="left">
              <ons-checkbox input-id="password-check"></ons-checkbox>
            </label>
            <label for="password-check" class="center">
              パスワードを表示
            </label>
          </p>
          <p class="create-top-margin">
            <button class="button button--outline">ユーザー名を検索</button>
          </p>
        </div>
      </form>
      <template id="alert-username-not-found.html">
        <ons-alert-dialog id="alert-username-not-found" modifier="rowfooter">
          <div class="alert-dialog-title">Alert</div>
          <div class="alert-dialog-content">
            ユーザーが見つかりません。
          </div>
          <div class="alert-dialog-footer">
            <ons-alert-dialog-button
              onclick="hideAlertDialog('alert-username-not-found')"
              >OK</ons-alert-dialog-button
            >
          </div>
        </ons-alert-dialog>
      </template>

      <template id="confirm-username-dialog.html">
        <ons-alert-dialog id="confirm-username-dialog" modifier="rowfooter">
          <div class="alert-dialog-title">Alert</div>
          <div class="alert-dialog-content">
            あなたのユーザー名は，<br />「{{.User.Username}}」です
          </div>
          <div class="alert-dialog-footer">
            <ons-alert-dialog-button
              onclick="hideAlertDialog('confirm-username-dialog')"
              >OK</ons-alert-dialog-button
            >
          </div>
        </ons-alert-dialog>
      </template>
    </ons-page>
    <script type="text/javascript" src="/static/js/common.js"></script>
    <script type="text/javascript">
       const name = {{.User.Username}};
       if (name === null) { ; }
       else {
         if (name == ""){
          dialogBoxEveryone('alert-username-not-found');
        } else {
          dialogBoxEveryone('confirm-username-dialog');
        }
      }
    </script>
  </body>
</html>
