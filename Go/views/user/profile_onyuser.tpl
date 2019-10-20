<div style="text-align: center;">
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
        <div  class="title">
          {{.User.Username}}
        </div>
        <div class="content">
          <p>年齢：{{.User.Age}}</p>
          <p>居住地：{{.User.Address}}</p>
        </div>
      </ons-col>
      <ons-col width="33%" align="center">
        <div class="image" align="center"><img src="{{.User.IconUrl}}" alt="{{.Username}}" width="100%"></div>
      </ons-col>
    </ons-row>
    <form id="delete_user" action="/tv/user/{{.User.Id}}" method="post">
      <input type="hidden" name="_method" value="DELETE">
      <button type="submit">ユーザーを削除する</button>
    </form>
    <a href="/tv/user/edit">edit</a>
  </ons-list-item>
</ons-list>