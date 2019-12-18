<!DOCTYPE html>
<html lang="ja">
  <head>
    {{ template "/common/header.tpl" . }}
  </head>

  <body>
    <ons-page>
      {{ template "/common/toolbar.tpl" . }}
      {{ template "/common/alert.tpl" . }}
      <ons-pull-hook id="pull-hook"></ons-pull-hook>
      <div class="area-right" style="height: 0px;">
        <ons-button modifier="quiet" onclick="dialogBoxEveryone('hint-dialog')"
          ><i class="far fa-question-circle hint-icon-right"></i
        ></ons-button>
      </div>
      <form id="create-user" action="/tv/user/" method="post">
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
            <ons-input
              name="password"
              modifier="underbar"
              type="password"
              value="{{.User.Password}}"
              placeholder="パスワード（8字以上）"
              id="password"
              minlength="8"
              maxlength="50"
              float
              required
            ></ons-input>
          </p>
          <p>
            <label class="left">
              <ons-checkbox input-id="password-check"></ons-checkbox>
            </label>
            <label for="password-check" class="center">
              パスワードを表示
            </label>
          </p>
          <p>
            <select
              name="gender"
              id="gender"
              class="select-input select-input--underbar select-search-table"
              required
            >
              <option>男性</option>
              <option>女性</option>
              <option>LGBT</option>
            </select>
          </p>
          <p>
            <select
              class="select-input select-input--underbar select-search-table"
              name="marital"
              id="marital"
              required
            >
              <option>未婚</option>
              <option>既婚</option>
            </select>
          </p>
          <p>
            <select
              class="select-input select-input--underbar select-search-table"
              id="bloodType"
              name="bloodType"
            >
              <option>A型</option>
              <option>B型</option>
              <option>O型</option>
              <option>AB型</option>
            </select>
          </p>
          <p>
            <label for="address" class="label-margin">＜居住地＞</label>
            <select
              class="select-input select-input--underbar select-search-table"
              id="address"
              name="address"
              required
            >
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
          </p>
          <p>
            <label for="job" class="label-margin">＜職業＞</label>
            <select
              class="select-input select-input--underbar select-search-table"
              id="job"
              name="job"
              required
            >
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
          </p>
          <p>
            <label for="age" class="label-margin">＜生年月日＞</label>
            <ons-input
              type="date"
              name="age"
              id="age"
              value="{{.User.Age}}"
              modifier="underbar"
              min="1920-01-01"
              max="2020-01-01"
              float
              required
            ></ons-input>
          </p>
          <p>
            <label for="IconURL" class="label-margin"
              >＜プロフィール画像のURL＞</label
            >
            <ons-input
              name="IconURL"
              id="IconURL"
              value="{{.User.IconUrl}}"
              modifier="underbar"
              placeholder="必須ではない"
              maxlength="400"
              float
            ></ons-input>
          </p>
          <p>
            <label for="SecondPassword" class="label-margin"
              >＜第2パスワードの設定＞</label
            >
            <ons-input
              id="SecondPassword"
              name="SecondPassword"
              value="{{.User.SecondPassword}}"
              modifier="underbar"
              placeholder="あなたの小学校の名前は?"
              maxlength="100"
              required
            ></ons-input>
          </p>
          <p class="create-top-bottom-margin">
            <ons-button modifier="quiet" onclick="previewUser('preview-dialog')"
              >プレビュー</ons-button
            >
          </p>
          <p>
            <ons-checkbox
              value="consent"
              id="consent-checkbox"
              style="vertical-align: middle;"
              required
            ></ons-checkbox>
            <ons-button
              modifier="quiet"
              id="consent-button"
              onclick="dialogBoxEveryone('terms-of-service')"
              >利用規約に同意</ons-button
            >
          </p>
          <p>
            <ons-checkbox
              value="privacy"
              id="privacy-checkbox"
              style="vertical-align: middle;"
              required
            ></ons-checkbox>
            <ons-button
              modifier="quiet"
              id="privacy-button"
              onclick="dialogBoxEveryone('privacy-policy')"
              >プライバシーポリシーに同意</ons-button
            >
          </p>
          <p class="create-top-bottom-margin">
            <button class="button button--outline">作成する</button>
          </p>
        </div>
      </form>
    </ons-page>
    <template id="terms-of-service.html">
      <ons-dialog id="terms-of-service" modifier="large" cancelable fullscreen>
        <ons-page>
          <ons-toolbar>
            <div class="left"></div>
            <div class="center">利用規約</div>
          </ons-toolbar>
          <div class="scroller list-margin" style="height:85%;">
            {{ template "/common/terms_of_service.tpl" . }}
          </div>
          <p class="area-center create-top-margin-5">
            <ons-button
              id="ok-button"
              onclick="hideAlertDialog('terms-of-service')"
            >
              OK
            </ons-button>
          </p>
        </ons-page>
      </ons-dialog>
    </template>
    <template id="privacy-policy.html">
      <ons-dialog id="privacy-policy" modifier="large" cancelable fullscreen>
        <ons-page>
          <ons-toolbar>
            <div class="left"></div>
            <div class="center">プライバシーポリシー</div>
          </ons-toolbar>
          <div class="scroller list-margin" style="height:85%;">
            {{ template "/common/privacy_policy.tpl" . }}
          </div>
          <p class="area-center create-top-margin-5">
            <ons-button
              id="ok-button"
              onclick="hideAlertDialog('privacy-policy')"
            >
              OK
            </ons-button>
          </p>
        </ons-page>
      </ons-dialog>
    </template>
    <template id="preview-dialog.html">
      <ons-dialog id="preview-dialog" modifier="large" cancelable fullscreen>
        <ons-page>
          <ons-toolbar id="preview-hide-swipe">
            <div class="left">
              <ons-button
                id="cancel-button"
                onclick="hideAlertDialog('preview-dialog')"
                style="background:left;color: grey;"
                ><i class="fas fa-window-close"></i
              ></ons-button>
            </div>
            <div class="center">
              プレビュー
            </div>
          </ons-toolbar>
          <div class="scroller">
            <ons-list>
              <ons-list-header>
                <div class="area-left">作成日：○○</div>
              </ons-list-header>
              <ons-list-item>
                <ons-row>
                  <ons-col>
                    <div class="title" id="preview-username"></div>
                    <div class="content" id="preview-contents"></div>
                    <div>※年齢・居住地・アイコンのみ公開されます。</div>
                  </ons-col>
                  <ons-col width="50%">
                    <div class="profile-image" id="preview-image"></div>
                  </ons-col>
                </ons-row>
              </ons-list-item>
            </ons-list>
          </div>
        </ons-page>
        <script>
          hideSwipeToolbar('preview-hide-swipe', 'preview-dialog');
        </script>
      </ons-dialog>
    </template>
    <!-- ヒントページ -->
    <template id="hint-dialog.html">
      <ons-dialog id="hint-dialog" modifier="large" cancelable fullscreen>
        <ons-page>
          <ons-toolbar id="hint-hide-swipe">
            <div class="left">
              <ons-button
                id="cancel-button"
                onclick="hideAlertDialog('hint-dialog')"
              >
                <i class="fas fa-window-close"></i>
              </ons-button>
            </div>
            <div class="center">
              ヒント <i class="far fa-question-circle hint-font-size"></i>
            </div>
          </ons-toolbar>
          <div class="scroller list-margin">
            <ul class="list">
              <li class="list-header">
                ユーザ登録のすすめ
              </li>
              <li class="list-item hint-list-dialog">
                <div class="list-item__center">
                  ユーザ名：文字種制限はありません
                </div>
              </li>
              <li class="list-item hint-list-dialog">
                <div class="list-item__center">
                  パスワード：ハッシュ化して保存しています
                </div>
              </li>
              <li class="list-item">
                <div class="list-item__center hint-list-dialog">
                  生年月日(android)：カレンダーが出現します。上部の年をクリックすると変更が可能です。
                </div>
              </li>
              <li class="list-item">
                <div class="list-item__center hint-list-dialog">
                  プロフィール画像：画像アドレスで指定（http://の場合、表示されない場合があります）。正方形が良！。なお、指定しない場合はランダムに設定されます。
                </div>
              </li>
            </ul>
          </div>
        </ons-page>
        <script>
          hideSwipeToolbar('hint-hide-swipe', 'hint-dialog');
        </script>
      </ons-dialog>
    </template>
    {{ template "/common/js.tpl" . }}
    <script type="text/javascript">
      var previewUser = function(elemID) {
        ons.ready(function() {
          var dialog = document.getElementById(elemID);
          // 同じページで2度目以降のプレビュークリック時発動
          if (dialog) {
            inputPreviewUserData();
            dialog.show();
          } else {
            ons
              .createElement(elemID + '.html', { append: true })
              .then(function(dialog) {
                inputPreviewUserData();
                dialog.show();
              });
          }
        });
      };
    </script>
    <script type="text/javascript">
      let consentFlag = false;
      $(function() {
        $('#consent-button').click(function() {
          consentFlag = true;
        });
      });
      $(function() {
        $('#consent-checkbox').change(function() {
          if (consentFlag) {
          } else {
            $('#consent-checkbox').prop('checked', false);
          }
        });
      });
    </script>
    <script type="text/javascript">
      let privacyFlag = false;
      $(function() {
        $('#privacy-button').click(function() {
          privacyFlag = true;
        });
      });
      $(function() {
        $('#privacy-checkbox').change(function() {
          if (privacyFlag) {
          } else {
            $('#privacy-checkbox').prop('checked', false);
          }
        });
      });
    </script>
    <script type="text/javascript">
      const name = {{.User.Username}};
      if (name != null) {
        dialogBoxEveryone("alert-username-duplicate");
      };
    </script>
    <script>
      if ({{.User}} === null) {
        ;
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
