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

<template id="alert_username_dialog.html">
  <ons-alert-dialog id="alert_username_dialog" modifier="rowfooter">
    <div class="alert-dialog-title">Alert</div>
    <div class="alert-dialog-content">
      ユーザー名、または、パスワードが誤っています。
    </div>
    <div class="alert-dialog-footer">
      <ons-alert-dialog-button
        onclick="hideAlertDialog('alert_username_dialog')"
        >OK</ons-alert-dialog-button
      >
    </div>
  </ons-alert-dialog>
</template>

<template id="alert_username_duplicate.html">
  <ons-alert-dialog id="alert_username_duplicate" modifier="rowfooter">
    <div class="alert-dialog-title">Alert</div>
    <div class="alert-dialog-content">
      {{.User.Username}} はすでに存在しています。
    </div>
    <div class="alert-dialog-footer">
      <ons-alert-dialog-button
        onclick="hideAlertDialog('alert_username_duplicate')"
        >OK</ons-alert-dialog-button
      >
    </div>
  </ons-alert-dialog>
</template>

<template id="alert_minlength.html">
  <ons-alert-dialog id="alert_minlength" modifier="rowfooter">
    <div class="alert-dialog-title">Alert</div>
    <div class="alert-dialog-content">
      5文字以上入力してください。
    </div>
    <div class="alert-dialog-footer">
      <ons-alert-dialog-button onclick="hideAlertDialog('alert_minlength')"
        >OK</ons-alert-dialog-button
      >
    </div>
  </ons-alert-dialog>
</template>
