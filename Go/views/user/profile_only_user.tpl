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
          <p>年齢　：{{.User.Age}}</p>
          <p>居住地：{{.User.Address}}</p>
        </div>
      </ons-col>
      <ons-col width="50%">
        <div class="image" style="max-height: 170px;">
          <img src="{{.User.IconURL}}" alt="{{.Username}}" width="100%" />
        </div>
      </ons-col>
    </ons-row>
    <a href="/tv/user/edit">
      <button class="button button--light" style="line-height: 12px;">
        編集
      </button>
    </a>
  </ons-list-item>
</ons-list>
