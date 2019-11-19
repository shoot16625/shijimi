<!DOCTYPE html>
<html lang="ja">
  <head>
    {{ template "/common/header.tpl" . }}
  </head>

  <body>
    <ons-page>
      {{ template "/common/toolbar.tpl" . }}
      {{ template "/common/alert.tpl" . }}

      <form id="update-user" action="/tv/user/{{.User.Id}}" method="post">
        <div class="input-table">
          <p>
            <label for="password">パスワード</label>
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
            <label class="left">
              <ons-checkbox input-id="password-check"></ons-checkbox>
            </label>
            <label for="password-check" class="center">
              パスワードを表示
            </label>
          </p>
          <p class="create-top-bottom-margin">
            <input type="hidden" name="_method" value="PUT" />
            <button class="button button--outline">パスワード再設定</button>
          </p>
        </div>
      </form>
    </ons-page>
    {{ template "/common/js.tpl" . }}
  </body>
</html>
