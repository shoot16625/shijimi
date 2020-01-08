<div class="background" style="background-color: white;"></div>

<ons-toolbar class="toolbar">
  <div class="left" id="mypage-toolbar">
    <img
      class="toolbar-image"
      src="{{.User.IconUrl}}"
      alt="mypage"
      title="mypage-icon"
      height="30px;"
      style="margin: 7px;"
      onerror="this.src='/static/img/user_img/mypage-icon.png'"
      onclick="location.href='/tv/user/show'"
    />
  </div>
  <div class="center" id="image-toolbar">
    <!-- androidだとtext-aline:leftのため -->
    <div class="area-center">
      <img
        class="toolbar-image"
        src="/static/img/shijimi-transparence.png"
        alt="shijimi"
        height="42px;"
        onclick="location.href='/'"
      />
    </div>
  </div>
  <div class="right">
    <ons-toolbar-button
      class="toolbar-image"
      icon="fa-search"
      onclick="dialogBoxEveryone('search-toolbar')"
    ></ons-toolbar-button>
  </div>
</ons-toolbar>
