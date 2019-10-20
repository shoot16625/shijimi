<!--     <div style="text-align: center;">
      <ons-row>
        <ons-col>
          <i class="fas fa-comments"></i>
          <button class="button--cta" style="line-height: 20px; background-color: cornflowerblue;" onclick="location.href='/tv/user/show'">リアルタイム</button>
        </ons-col>
        <ons-col>
          <i class="fas fa-user"></i>
          <button class="button--cta" style="line-height: 20px; background-color: cornflowerblue;" onclick="location.href='/tv/user/show_review'">レビュー</button>
        </ons-col>
      </ons-row>
    </div> -->
<div style="text-align: center;">
  <ons-row>
    <ons-col style="text-align: right;">
      <button
        class="button--cta"
        style="line-height: 15px; background-color: cornflowerblue;"
        onclick="location.href='/tv/user/show'"
      >
        <i class="fas fa-comments"></i>Timeline
      </button>
    </ons-col>
    <ons-col style="text-align: left;">
      <button
        class="button--cta"
        style="line-height: 15px; background-color: darkorange;"
        onclick="location.href='/tv/user/show_review'"
      >
        <i class="fas fa-user"></i>Review
      </button>
    </ons-col>
  </ons-row>
</div>

<div style="text-align: center;">
  <ons-row>
    <ons-col style="text-align: right;">
      <button
        class="button--cta"
        style="line-height: 15px; background-color: darkorange;"
        onclick="location.href='/tv/user/show_watched_tv'"
      >
        <i class="fas fa-comments"></i>見た
      </button>
    </ons-col>
    <ons-col style="text-align: left;">
      <button
        class="button--cta"
        style="line-height: 15px; background-color: cornflowerblue;"
        onclick="location.href='/tv/user/show_wtw_tv'"
      >
        <i class="fas fa-user"></i>また今度
      </button>
    </ons-col>
  </ons-row>
</div>

<ons-list>
  <ons-list-header>
    <div style="text-align: left; float:left;">
      作成日：{{.User.Created|dateformatJst}}
    </div>
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
      <ons-col width="33%" style="text-align: center;">
        <div class="image" style="text-align: center;">
          <img src="{{.User.IconUrl}}" alt="{{.Username}}" width="100%" />
        </div>
      </ons-col>
    </ons-row>
    <a href="/tv/user/edit"><button>プロフィール編集</button></a>
    <form id="delete_user" action="/tv/user/{{.User.Id}}" method="post">
      <input type="hidden" name="_method" value="DELETE" />
      <button type="submit">ユーザーを削除</button>
    </form>
    <!-- <a href="/tv/user/edit">edit</a> -->
  </ons-list-item>
</ons-list>
