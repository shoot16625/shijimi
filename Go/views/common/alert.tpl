<template id="alert-only-user-dialog.html">
  <ons-alert-dialog id="alert-only-user-dialog" modifier="rowfooter">
    <div class="alert-dialog-title">Alert</div>
    <div class="alert-dialog-content">
      この機能はログインユーザーのみ<br />利用できます。
    </div>
    <div class="alert-dialog-footer">
      <ons-alert-dialog-button
        onclick="hideAlertDialog('alert-only-user-dialog')"
        >OK</ons-alert-dialog-button
      >
    </div>
  </ons-alert-dialog>
</template>

<template id="alert-username-dialog.html">
  <ons-alert-dialog id="alert-username-dialog" modifier="rowfooter">
    <div class="alert-dialog-title">Alert</div>
    <div class="alert-dialog-content">
      ユーザー名、または、パスワードが誤っています。
    </div>
    <div class="alert-dialog-footer">
      <ons-alert-dialog-button
        onclick="hideAlertDialog('alert-username-dialog')"
        >OK</ons-alert-dialog-button
      >
    </div>
  </ons-alert-dialog>
</template>

<template id="alert-username-duplicate.html">
  <ons-alert-dialog id="alert-username-duplicate" modifier="rowfooter">
    <div class="alert-dialog-title">Alert</div>
    <div class="alert-dialog-content">
      {{.User.Username}} はすでに存在しています。
    </div>
    <div class="alert-dialog-footer">
      <ons-alert-dialog-button
        onclick="hideAlertDialog('alert-username-duplicate')"
        >OK</ons-alert-dialog-button
      >
    </div>
  </ons-alert-dialog>
</template>

<template id="alert-min-length.html">
  <ons-alert-dialog id="alert-min-length" modifier="rowfooter">
    <div class="alert-dialog-title">Alert</div>
    <div class="alert-dialog-content">
      5文字以上入力してください。
    </div>
    <div class="alert-dialog-footer">
      <ons-alert-dialog-button onclick="hideAlertDialog('alert-min-length')"
        >OK</ons-alert-dialog-button
      >
    </div>
  </ons-alert-dialog>
</template>

<template id="unsubscribe-dialog.html">
  <ons-alert-dialog id="unsubscribe-dialog" modifier="rowfooter">
    <div class="alert-dialog-title">Alert</div>
    <div class="alert-dialog-content">
      本当に退会しますか？<br />あなたの全ての投稿データが削除されます。
    </div>
    <div class="alert-dialog-footer">
      <ons-alert-dialog-button onclick="hideAlertDialog('unsubscribe-dialog')"
        >Cancel</ons-alert-dialog-button
      >
      <ons-alert-dialog-button>
        <form id="delete-user" action="/tv/user/{{.User.Id}}" method="post">
          <input type="hidden" name="_method" value="DELETE" />
          <button class="button--quiet" type="submit">
            OK
          </button>
        </form>
      </ons-alert-dialog-button>
    </div>
  </ons-alert-dialog>
</template>
