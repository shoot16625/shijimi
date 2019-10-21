<ons-list>
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
          <p>年齢：{{.User.Age}}</p>
          <p>居住地：{{.User.Address}}</p>
          <p>職業：{{.User.Job}}</p>
        </div>
      </ons-col>
      <ons-col width="33%">
        <div class="image area-center">
          <img src="{{.User.IconURL}}" alt="{{.Username}}" width="100%" />
        </div>
      </ons-col>
    </ons-row>
    <a href="/tv/user/edit"><button>プロフィール編集</button></a>
    <form id="delete-user" action="/tv/user/{{.User.Id}}" method="post">
      <input type="hidden" name="_method" value="DELETE" />
      <button type="submit">ユーザーを削除</button>
    </form>
  </ons-list-item>
</ons-list>
<ons-row>
  <ons-col class="area-right">
    <button
      class="button--cta"
      style="line-height: 15px; background-color: cornflowerblue; width: 100%"
      onclick="location.href='/tv/user/show'"
    >
      <i class="fas fa-comments"></i>Timeline
    </button>
  </ons-col>
  <ons-col style="text-align: left;">
    <button
      class="button--cta"
      style="line-height: 15px; background-color: rgb(240, 141, 20); width: 100%"
      onclick="location.href='/tv/user/show_review'"
    >
      <i class="fas fa-user"></i>Review
    </button>
  </ons-col>
  <ons-col style="text-align: left;">
    <button
      class="button--cta"
      style="line-height: 15px; background-color: cornflowerblue; width: 100%;"
      onclick="location.href='/tv/user/show_watched_tv'"
    >
      <i class="fas fa-laugh-beam"></i>見た
    </button>
  </ons-col>
  <ons-col style="text-align: left;">
    <button
      class="button--cta"
      style="line-height: 15px; background-color: rgb(240, 141, 20); width: 100%"
      onclick="location.href='/tv/user/show_wtw_tv'"
    >
      <i class="fas fa-bookmark"></i>見たい
    </button>
  </ons-col>
</ons-row>
