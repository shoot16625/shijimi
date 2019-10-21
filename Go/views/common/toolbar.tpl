<div class="background" style="background-color: white;"></div>

<ons-toolbar class="toolbar">
  <div class="left" id="mypage_toolbar">
    <ons-toolbar-button
      icon="md-face"
      style="font-size:24px;"
      onclick="location.href='/tv/user/show'"
    ></ons-toolbar-button>
  </div>
  <div class="center" id="image_toolbar">
    <img
      src="/static/img/shijimi-touka02.png"
      alt="shijimi"
      height="42px;"
      onclick="location.href='/'"
    />
  </div>
  <div class="right">
    <ons-toolbar-button
      icon="fa-search"
      onclick="dialogBoxEveryone('search_toolbar')"
    ></ons-toolbar-button>
  </div>
</ons-toolbar>

<template id="search_toolbar.html">
  <ons-dialog id="search_toolbar" modifier="large" cancelable fullscreen>
    <ons-page>
      <ons-toolbar>
        <div class="left">
          <ons-button
            id="cancel-button"
            onclick="hideAlertDialog('search_toolbar')"
            style="background:left;color: grey;"
            ><i class="fas fa-window-close"></i
          ></ons-button>
        </div>
        <div class="center">
          ドラマ・映画検索
        </div>
      </ons-toolbar>
      <form id="search_tv_program" action="/tv/tv_program/search" method="post">
        <div style="text-align: center; margin-top: 30px;">
          <p>
            <ons-search-input
              name="search_word"
              placeholder="Search"
              id="search_input"
            ></ons-search-input>
          </p>
          <p style="margin-top: 30px;">
            <button class="button button--outline">search</button>
          </p>
          <p style="margin-top: 40px;">
            タイトル・出演者・主題歌・スタッフ<br />季節・年・曜日・ジャンルなど
          </p>
        </div>
      </form>
    </ons-page>
  </ons-dialog>
</template>
