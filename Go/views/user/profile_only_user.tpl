<ons-pull-hook id="pull-hook"></ons-pull-hook>
<ons-list id="my-profile">
  <ons-list-header>
    <div class="area-left">作成日：{{.User.Created|dateformatJst}}</div>
  </ons-list-header>
  <ons-list-item>
    <ons-row>
      <ons-col>
        <div class="title">
          {{.User.Username}}
        </div>
        <div class="content">
          <p>年齢　：{{.User.Age|birthday2Age}}</p>
          <p>居住地：{{.User.Address}}</p>
          <p>ポイント：{{.User.MoneyPoint}}</p>
          <a href="/tv/user/edit">
            <button
              class="button button--light"
              style="line-height: 12px; margin-bottom: 12px;"
            >
              編集
            </button>
          </a>
          <div class="area-right-float" style="margin-right: 5px;">
            <ons-button
              modifier="quiet"
              onclick="dialogBoxEveryone('hint-dialog')"
              ><i class="far fa-question-circle hint-icon-right"></i
            ></ons-button>
          </div>
        </div>
      </ons-col>
      <ons-col width="50%">
        <div class="profile-image">
          <img src="{{.User.IconUrl}}" alt="{{.User.Username}}" width="100%" />
        </div>
      </ons-col>
    </ons-row>
    <ons-row>
      <span id="badges"></span>
    </ons-row>
  </ons-list-item>
</ons-list>
<template id="hint-dialog.html">
  <ons-dialog id="hint-dialog" modifier="large" cancelable fullscreen>
    <ons-page>
      <ons-toolbar id="hint-hide-swipe">
        <div class="left">
          <ons-button
            id="cancel-button"
            onclick="hideAlertDialog('hint-dialog')"
          >
            <i class="fas fa-window-close"></i>
          </ons-button>
        </div>
        <div class="center">
          ヒント <i class="far fa-question-circle hint-font-size"></i>
        </div>
      </ons-toolbar>
      <div class="scroller list-margin">
        <ol>
          <li>
            ログインごとにポイントが貯まります。<br />ポイントは機能のカスタマイズに利用できるようにします。
          </li>
          <li>バッジ：ときどき付与されます。</li>
          <li>TM：1000件が表示されます。(初期状態)</li>
          <li>レビュー：100件が表示されます。(初期状態)</li>
          <li>見た：1000件が表示されます。(初期状態)</li>
          <li>見たい：1000件が表示されます。(初期状態)</li>
          <li>見たログが5つ以上たまると、番組のおすすめ機能が発動します。</li>
          <li>過去のコメントの編集はできません。</li>
        </ol>
      </div>
    </ons-page>
    <script>
      hideSwipeToolbar('hint-hide-swipe', 'hint-dialog');
    </script>
  </ons-dialog>
</template>
