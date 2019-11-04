<ons-pull-hook id="pull-hook"></ons-pull-hook>
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
          <p>年齢　：{{.User.Age|Birthday2Age}}</p>
          <p>居住地：{{.User.Address}}</p>
        </div>
      </ons-col>
      <ons-col width="50%">
        <div class="profile-image">
          <img src="{{.User.IconURL}}" alt="{{.Username}}" width="100%" />
        </div>
      </ons-col>
    </ons-row>
    <ons-row>
      <span id="badges"></span>
    </ons-row>
  </ons-list-item>
</ons-list>
