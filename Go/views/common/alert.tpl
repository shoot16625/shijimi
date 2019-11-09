<template id="alert-only-user-dialog.html">
  <ons-alert-dialog id="alert-only-user-dialog" modifier="rowfooter">
    <div class="alert-dialog-title">Alert</div>
    <div class="alert-dialog-content">
      <span>この機能はログインユーザーのみ</span
      ><span class="new-line">利用できます。</span>
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
        <form
          id="delete-user"
          action="/tv/user/{{.User.Id}}"
          method="post"
          onSubmit="showLoading();"
        >
          <input type="hidden" name="_method" value="DELETE" />
          <button id="delete-user-button" class="button--quiet" type="submit">
            OK
          </button>
        </form>
      </ons-alert-dialog-button>
    </div>
  </ons-alert-dialog>
</template>
<template id="alert-review-twice.html">
  <ons-alert-dialog id="alert-review-twice" modifier="rowfooter">
    <div class="alert-dialog-title">Alert</div>
    <div class="alert-dialog-content">
      <span>既にレビューが行われています。</span
      ><span class="new-line">変更したい場合は</span
      ><span class="new-line">マイページから削除後、</span
      ><span class="new-line">お願いします。</span>
    </div>
    <div class="alert-dialog-footer">
      <ons-alert-dialog-button onclick="hideAlertDialog('alert-review-twice')"
        >OK</ons-alert-dialog-button
      >
    </div>
  </ons-alert-dialog>
</template>
