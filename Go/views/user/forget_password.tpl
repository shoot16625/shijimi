<!DOCTYPE html>
<html lang="ja">
<head>
  {{ template "/common/header.tpl" . }}
</head>

<body>
  <ons-page>
    {{ template "/common/toolbar.tpl" . }}
    {{ template "/common/alert.tpl" . }}

    <form id="login_user" action="/tv/user/forget_password" method="post">
      <div style="text-align: center; margin-top: 30px;">
        <p>
          <ons-input name="username" modifier="underbar" placeholder="ユーザー名" float required></ons-input>
        </p>
        <p>
          <label for="SecondPassword">第2パスワード</label>
          <ons-input id="SecondPassword" name="SecondPassword"modifier="underbar" placeholder="あなたの小学校の名前は?
          " maxlength="100" float required></ons-input>
        </p>
        <p style="margin-top: 30px;">
          <button class="button button--outline">パスワードを検索</button>
        </p>
      </div>
    </form>
    <template id="alert_username_notfound.html">
      <ons-alert-dialog id="alert_username_notfound" modifier="rowfooter">
        <div class="alert-dialog-title">Alert</div>
        <div class="alert-dialog-content">
          ユーザー情報が誤っています。
        </div>
        <div class="alert-dialog-footer">
          <ons-alert-dialog-button onclick="hideAlertDialog('alert_username_notfound')">OK</ons-alert-dialog-button>
        </div>
      </ons-alert-dialog>
    </template>

  </ons-page>
  <script type="text/javascript" src="/static/js/common.js"></script>
  <script type="text/javascript">
    const name = {{.User}};
    console.log(name);
    if (name != "null"){
      DialogBoxEveryone('alert_username_notfound')
    }
  </script>
</body>
</html>