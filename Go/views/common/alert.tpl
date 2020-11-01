<template id="alert-only-user-dialog.html">
  <ons-alert-dialog id="alert-only-user-dialog" modifier="rowfooter">
    <div class="alert-dialog-title">Alert</div>
    <div class="alert-dialog-content">
      <span>この機能はログインユーザーのみ</span
      ><span class="new-line">利用できます。</span
      ><span class="new-line"><a href="/">新規登録はこちら</a></span>
    </div>
    <div class="alert-dialog-footer">
      <ons-alert-dialog-button
        onclick="hideAlertDialog('alert-only-user-dialog')"
        >OK</ons-alert-dialog-button
      >
    </div>
  </ons-alert-dialog>
</template>

<template id="alert-only-not-user-dialog.html">
  <ons-alert-dialog id="alert-only-not-user-dialog" modifier="rowfooter">
    <div class="alert-dialog-title">Alert</div>
    <div class="alert-dialog-content">
      <span>非ログインユーザーのみ</span
      ><span class="new-line">利用できます。</span>
    </div>
    <div class="alert-dialog-footer">
      <ons-alert-dialog-button
        onclick="hideAlertDialog('alert-only-not-user-dialog')"
        >OK</ons-alert-dialog-button
      >
    </div>
  </ons-alert-dialog>
</template>

<template id="alert-only-admin-dialog.html">
  <ons-alert-dialog id="alert-only-admin-dialog" modifier="rowfooter">
    <div class="alert-dialog-title">Alert</div>
    <div class="alert-dialog-content">
      <span>この機能は管理者のみ</span
      ><span class="new-line">利用できます。</span>
    </div>
    <div class="alert-dialog-footer">
      <ons-alert-dialog-button
        onclick="hideAlertDialog('alert-only-admin-dialog')"
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

<!-- アラートではない -->
<template id="search-toolbar.html">
  <ons-dialog id="search-toolbar" modifier="large" cancelable fullscreen>
    <ons-page>
      <ons-toolbar id="search-hide-swipe">
        <div class="left">
          <ons-button
            id="cancel-button"
            onclick="hideAlertDialog('search-toolbar')"
            ><i class="fas fa-window-close"></i
          ></ons-button>
        </div>
        <div class="center">
          ドラマ・映画検索
        </div>
      </ons-toolbar>
      <form id="search_tv_program" action="/tv/tv_program/search" method="post">
        <div class="area-center create-top-bottom-margin">
          <p>
            <ons-search-input
              name="search-word"
              placeholder="Search"
              id="search-word"
            ></ons-search-input>
          </p>
          <p class="create-top-bottom-margin">
            <button class="button button--outline">search</button>
          </p>
          <p style="margin-top: 40px;">
            タイトル・出演者・主題歌・スタッフ<br />季節・年・曜日・ジャンルなど
          </p>
        </div>
      </form>
    </ons-page>
    <script>
      hideSwipeToolbar('search-hide-swipe', 'search-toolbar');
    </script>
  </ons-dialog>
</template>
