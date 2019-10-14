<!--     <div style="text-align: center;">
      <ons-row>
        <ons-col>
          <i class="fas fa-comments"></i>
          <button class="button--cta" style="line-height: 20px; background-color: cornflowerblue;" onclick="location.href='/tv/user/show/{{.User.Id}}'">リアルタイム</button>
        </ons-col>
        <ons-col>
          <i class="fas fa-user"></i>
          <button class="button--cta" style="line-height: 20px; background-color: cornflowerblue;" onclick="location.href='/tv/user/show_review/{{.User.Id}}'">レビュー</button>
        </ons-col>
      </ons-row>
    </div> -->
<div style="text-align: center;">
  <ons-row>
    <!-- <i class="fas fa-arrow-left"></i>左にスワイプ -->
    <ons-col style="text-align: right;">
      <button class="button--cta" style="line-height: 15px; background-color: cornflowerblue;" onclick="location.href='/tv/user/show/{{.User.Id}}'"><i class="fas fa-comments"></i>Timeline</button>
    </ons-col>
    <ons-col style="text-align: left;">
      <button class="button--cta" style="line-height: 15px; background-color: darkorange;" onclick="location.href='/tv/user/show_review/{{.User.Id}}'"><i class="fas fa-user"></i>Review</button>
    </ons-col>
    <!-- 右にスワイプ<i class="fas fa-arrow-right"></i> -->
  </ons-row>
</div>

<div style="text-align: center;">
  <i class="fas fa-arrow-left"></i>左にスワイプ
  右にスワイプ<i class="fas fa-arrow-right"></i>
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
<!--         <form id="delete_user" action="/tv/user/{{.User.Id}}" method="post">
          <input type="hidden" name="_method" value="DELETE">
          <button type="submit">ユーザーを削除する</button>
        </form>
        <a href="/tv/user/edit">edit</a> -->
  </ons-list-item>
</ons-list>