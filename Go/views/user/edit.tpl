<!DOCTYPE html>
<html lang="ja">
  <head>
    {{ template "/common/header.tpl" . }}
  </head>

  <body>
    <ons-page>
      {{ template "/common/toolbar.tpl" . }}
      {{ template "/common/alert.tpl" . }}

      <form id="create-user" action="/tv/user/{{.User.Id}}" method="post">
        <div class="input-table">
          <p>
            <ons-input
              name="username"
              modifier="underbar"
              placeholder="ユーザー名"
              value="{{.User.Username}}"
              minlength="2"
              maxlength="20"
              float
              required
            ></ons-input>
          </p>
          <p>
            <label for="age">＜ 年齢 ＞</label>
            <ons-input
              type="number"
              name="age"
              id="age"
              value="{{.User.Age}}"
              modifier="underbar"
              placeholder="年齢"
              min="0"
              max="120"
              float
              required
            ></ons-input>
          </p>
          <p>
            <ons-select
              class="select"
              name="gender"
              id="gender"
              modifier="underbar"
              float
              required
            >
              <select class="select-input">
                <option>男性</option>
                <option>女性</option>
                <option>LGBT</option>
              </select>
            </ons-select>
          </p>
          <p>
            <ons-select
              class="select"
              name="marital"
              id="marital"
              modifier="underbar"
              float
              required
            >
              <select class="select-input">
                <option>未婚</option>
                <option>既婚</option>
              </select>
            </ons-select>
          </p>
          <p>
            <ons-select
              class="select"
              name="bloodType"
              id="bloodType"
              modifier="underbar"
              float
              required
            >
              <select class="select-input">
                <option>A型</option>
                <option>B型</option>
                <option>O型</option>
                <option>AB型</option>
              </select>
            </ons-select>
          </p>
          <p>
            <ons-select
              class="select"
              name="address"
              id="address"
              modifier="underbar"
              float
              required
            >
              <label for="prefecture">＜居住地＞</label>
              <select class="select-input" id="prefecture">
                <option value="北海道">北海道</option>
                <option value="青森県">青森県</option>
                <option value="岩手県">岩手県</option>
                <option value="宮城県">宮城県</option>
                <option value="秋田県">秋田県</option>
                <option value="山形県">山形県</option>
                <option value="福島県">福島県</option>
                <option value="茨城県">茨城県</option>
                <option value="栃木県">栃木県</option>
                <option value="群馬県">群馬県</option>
                <option value="埼玉県">埼玉県</option>
                <option value="千葉県">千葉県</option>
                <option value="東京都">東京都</option>
                <option value="神奈川県">神奈川県</option>
                <option value="新潟県">新潟県</option>
                <option value="富山県">富山県</option>
                <option value="石川県">石川県</option>
                <option value="福井県">福井県</option>
                <option value="山梨県">山梨県</option>
                <option value="長野県">長野県</option>
                <option value="岐阜県">岐阜県</option>
                <option value="静岡県">静岡県</option>
                <option value="愛知県">愛知県</option>
                <option value="三重県">三重県</option>
                <option value="滋賀県">滋賀県</option>
                <option value="京都府">京都府</option>
                <option value="大阪府">大阪府</option>
                <option value="兵庫県">兵庫県</option>
                <option value="奈良県">奈良県</option>
                <option value="和歌山県">和歌山県</option>
                <option value="鳥取県">鳥取県</option>
                <option value="島根県">島根県</option>
                <option value="岡山県">岡山県</option>
                <option value="広島県">広島県</option>
                <option value="山口県">山口県</option>
                <option value="徳島県">徳島県</option>
                <option value="香川県">香川県</option>
                <option value="愛媛県">愛媛県</option>
                <option value="高知県">高知県</option>
                <option value="福岡県">福岡県</option>
                <option value="佐賀県">佐賀県</option>
                <option value="長崎県">長崎県</option>
                <option value="熊本県">熊本県</option>
                <option value="大分県">大分県</option>
                <option value="宮崎県">宮崎県</option>
                <option value="鹿児島県">鹿児島県</option>
                <option value="沖縄県">沖縄県</option>
                <option value="海外">海外</option>
              </select>
            </ons-select>
          </p>
          <p>
            <ons-select
              class="select"
              name="job"
              id="job"
              modifier="underbar"
              float
              required
            >
              <label for="jobs">＜職業＞</label>
              <select class="select-input" id="jobs">
                <option>学生</option>
                <option>エンジニア</option>
                <option>会社員</option>
                <option>公務員</option>
                <option>自営業</option>
                <option>会社役員</option>
                <option>御隠居</option>
                <option>専業主婦（夫）</option>
                <option>パート・アルバイト</option>
                <option>フリーター</option>
                <option>その他</option>
              </select>
            </ons-select>
          </p>

          <p>
            <label for="IconURL">＜プロフィール画像のURL＞</label>
            <ons-input
              name="IconURL"
              id="IconURL"
              value="{{.User.IconURL}}"
              modifier="underbar"
              placeholder="必須ではない"
              maxlength="200"
              float
            ></ons-input>
          </p>
          <p class="create-top-margin">
            <input type="hidden" name="_method" value="PUT" />
            <button class="button button--outline">更新</button>
          </p>
          <p class="area-right">
            <a href="forget_password_page">パスワードの変更はこちら</a>
          </p>
          <p class="area-right">
            <ons-button
              modifier="quiet"
              onclick="dialogBox('unsubscribe-dialog',{{.User.Id}})"
              >退会はこちら</ons-button
            >
          </p>
        </div>
      </form>
    </ons-page>

    <script type="text/javascript" src="/static/js/common.js"></script>
    <script type="text/javascript">
      const nameFlag = {{.NameFlag}};
      if (nameFlag === true) {
       dialogBoxEveryone("alert-username-duplicate");
      };
    </script>
    <script>
      if ({{.User}} === null){;
      } else {
        document.getElementById('gender').value = {{.User.Gender}};
        document.getElementById('marital').value = {{.User.Marital}};
        document.getElementById('job').value = {{.User.Job}};
        document.getElementById('address').value = {{.User.Address}};
        document.getElementById('bloodType').value = {{.User.BloodType}};
      }
    </script>
  </body>
</html>
