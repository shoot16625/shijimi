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
