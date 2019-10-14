<!DOCTYPE html>
<html lang="ja">
<head>
  {{ template "/common/header.tpl" . }}
</head>

<body>
  <ons-page>
    {{ template "/common/toolbar.tpl" . }}
    {{ template "/common/alert.tpl" . }}

    <form id="update_user" action="/tv/user/{{.User.Id}}" method="post">
      <div style="text-align: center; margin-top: 30px;">
        <p>
          <label for="password">パスワード</label>
          <ons-input name="password" modifier="underbar" type="password" placeholder="パスワード" minlength="8" id="password" float required></ons-input>
        </p>
        <input type="hidden" name="username" value="{{.User.Username}}"></input>
        <input type="hidden" name="age" value="{{.User.Age}}"></input>
        <input type="hidden" name="job" value="{{.User.Job}}"></input>
        <input type="hidden" name="gender" value="{{.User.Gender}}"></input>
        <input type="hidden" name="address" value="{{.User.Address}}"></input>
        <input type="hidden" name="IconUrl" value="{{.User.IconUrl}}"></input>
        <input type="hidden" name="marital" value="{{.User.Marital}}"></input>
        <input type="hidden" name="SecondPassword" value="{{.User.SecondPassword}}"></input>
        <p>
         <label class="left">
          <ons-checkbox input-id="password-check"></ons-checkbox>
        </label>
        <label for="password-check" class="center">
          パスワードを表示
        </label>
        </p>
        <p style="margin-top: 30px;">
          <input type="hidden" name="_method" value="PUT">
          <button class="button button--outline">パスワード再設定</button>
        </p>
      </div>
    </form>
  </ons-page>
  <script type="text/javascript" src="/static/js/common.js"></script>

</body>
</html>