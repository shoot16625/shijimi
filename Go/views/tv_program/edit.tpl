<!DOCTYPE html>
<html lang="ja">
  <head>
    {{ template "/common/header.tpl" . }}
  </head>

  <style type="text/css">
    p {
      margin-top: 20px;
      margin-bottom: 20px;
    }
  </style>

  <body>
    <ons-page>
      {{ template "/common/toolbar.tpl" . }}
      {{ template "/common/alert.tpl" . }}
      <ons-pull-hook id="pull-hook"></ons-pull-hook>
      <form
        id="edit-tv-program"
        action="/tv/tv_program/{{.TvProgram.Id}}"
        method="post"
      >
        <div class="input-table">
          <p>
            <i class="fas fa-flag flag-color"></i>
            <ons-input
              name="title"
              modifier="underbar"
              value="{{.TvProgram.Title}}"
              placeholder="タイトル"
              maxlength="40"
              float
              required
            ></ons-input>
          </p>
          <p>
            <textarea
              class="textarea"
              style="width: 80%;"
              name="content"
              id="content"
              rows="7"
              placeholder="あらすじ・見どころ"
              maxlength="200"
            ></textarea>
          </p>
          <p>
            <i class="fas fa-flag flag-color"></i>
            <ons-input
              name="cast"
              modifier="underbar"
              value="{{.TvProgram.Cast}}"
              placeholder="キャスト(ex.小栗旬,石原さとみ)"
              maxlength="100"
              float
              required
            ></ons-input>
          </p>
          <p>
            <i class="fas fa-flag flag-color"></i>
            <ons-input
              type="number"
              name="year"
              value="{{.TvProgram.Year}}"
              modifier="underbar"
              placeholder="放送年(ex.2012)"
              min="1900"
              max="2100"
              float
              required
            ></ons-input>
          </p>
          <p>
            <ons-input
              name="themesong"
              value="{{.TvProgram.Themesong}}"
              modifier="underbar"
              placeholder="主題歌(ex.miwa「ヒカリへ」)"
              maxlength="100"
              float
            ></ons-input>
          </p>
          <p>
            <i class="fas fa-flag flag-color"></i>
            <label for="season">＜シーズン＞</label>
            <select
              name="season"
              id="season"
              class="select-input select-input--underbar select-search-table"
              required
            >
              <option>春(4~6)</option>
              <option>夏(7~9)</option>
              <option>秋(10~12)</option>
              <option>冬(1~3)</option>
            </select>
          </p>
          <p>
            <i class="fas fa-flag flag-color"></i>
            <label for="week">＜放送曜日＞</label>
            <select
              name="week"
              id="week"
              class="select-input select-input--underbar select-search-table"
              required
            >
              <option>月</option>
              <option>火</option>
              <option>水</option>
              <option>木</option>
              <option>金</option>
              <option>土</option>
              <option>日</option>
              <option>平日</option>
              <option>スペシャル</option>
              <option>映画</option>
              <option>?</option>
            </select>
          </p>
          <p>
            <label for="hour">＜時間帯＞</label>
            <select
              name="hour"
              id="hour"
              class="select-input select-input--underbar select-search-table"
            >
            </select>
          </p>
          <p>
            <i class="fas fa-flag flag-color"></i>
            <label for="category">＜ジャンル＞※複数可</label>
            <select
              name="category"
              id="category"
              style="height: 130px;"
              class="select-input select-input--underbar select-search-table restrict"
              required
              multiple
            >
              <option>コメディ・パロディ</option>
              <option>恋愛</option>
              <option>学園・青春</option>
              <option>グルメ</option>
              <option>ホーム・ヒューマン</option>
              <option>企業・オフィス</option>
              <option>刑事・検事</option>
              <option>弁護士</option>
              <option>医療</option>
              <option>時代劇</option>
              <option>スポーツ</option>
              <option>政治</option>
              <option>不倫</option>
              <option>ミステリー・サスペンス</option>
              <option>探偵・推理</option>
              <option>犯罪・復讐</option>
              <option>ホラー</option>
              <option>ドキュメンタリー</option>
              <option>アクション</option>
              <option>アニメ映画</option>
              <option>SF</option>
            </select>
          </p>
          <p>
            <ons-input
              name="ImageURL"
              modifier="underbar"
              value="{{.TvProgram.ImageUrl}}"
              placeholder="イメージ画像URL"
              maxlength="400"
              float
            ></ons-input>
          </p>
          <p>
            <ons-input
              name="MovieURL"
              modifier="underbar"
              value="{{.TvProgram.MovieUrl}}"
              placeholder="youtube動画URL 公式チャンネルonly"
              maxlength="400"
              float
            ></ons-input>
          </p>
          <p>
            <ons-input
              name="WikiReference"
              modifier="underbar"
              value="{{.TvProgram.WikiReference}}"
              placeholder="URL @Wikipedia"
              maxlength="400"
              float
            ></ons-input>
          </p>
          <p>
            <ons-input
              name="dramatist"
              modifier="underbar"
              value="{{.TvProgram.Dramatist}}"
              placeholder="脚本"
              maxlength="100"
              float
            ></ons-input>
          </p>
          <p>
            <ons-input
              name="supervisor"
              modifier="underbar"
              placeholder="監督"
              value="{{.TvProgram.Supervisor}}"
              maxlength="100"
              float
            ></ons-input>
          </p>
          <p>
            <ons-input
              name="director"
              value="{{.TvProgram.Director}}"
              modifier="underbar"
              placeholder="演出"
              maxlength="100"
              float
            ></ons-input>
          </p>
          <p>
            <ons-input
              name="production"
              modifier="underbar"
              value="{{.TvProgram.Production}}"
              placeholder="制作会社"
              maxlength="20"
              float
            ></ons-input>
          </p>
          <p class="create-top-bottom-margin">
            <ons-button
              modifier="quiet"
              onclick="previewTvProgram('preview-dialog')"
              >プレビュー</ons-button
            >
          </p>
          <p class="create-top-bottom-margin">
            <input type="hidden" name="_method" value="PUT" />
            <button class="button button--outline">更新する</button>
          </p>
          <p class="area-right">
            <ons-button
              modifier="quiet"
              onclick="goOtherPage({{.User.Id}},{{.TvProgram.Id}},'delete-tvprogram')"
              >テレビ削除</ons-button
            >
          </p>
        </div>
      </form>
    </ons-page>
    <template id="alert-tv-title.html">
      <ons-alert-dialog id="alert-tv-title" modifier="rowfooter">
        <div class="alert-dialog-title">Alert</div>
        <div class="alert-dialog-content">
          「{{.TvProgram.Title}}」はすでに存在しています。
        </div>
        <div class="alert-dialog-footer">
          <ons-alert-dialog-button onclick="hideAlertDialog('alert-tv-title')"
            >OK</ons-alert-dialog-button
          >
        </div>
      </ons-alert-dialog>
    </template>

    <template id="preview-dialog.html">
      <ons-dialog id="preview-dialog" modifier="large" cancelable fullscreen>
        <ons-page>
          <ons-toolbar>
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
              <ons-list-header style="background-color:ghostwhite;">
                <div class="area-left" id="preview-on-air-info"></div>
                <div class="area-right list-margin">
                  <i class="fas fa-eye"></i>：0
                </div>
              </ons-list-header>
              <ons-list-item id="expandable-list-item" expandable>
                <div id="preview-title"></div>
                <div class="expandable-content">
                  <ons-row>
                    <ons-col>
                      <div class="content">
                        <ons-row class="list-margin-bottom">
                          <ons-col width="20%">出演：</ons-col>
                          <ons-col id="preview-cast"></ons-col>
                        </ons-row>
                        <ons-row class="list-margin-bottom">
                          <ons-col width="20%"
                            ><i
                              class="fas fa-music"
                              style="color: cornflowerblue;"
                            ></i
                            >：</ons-col
                          >
                          <ons-col id="preview-themesong"></ons-col>
                        </ons-row>
                        <ons-row class="list-margin-bottom">
                          <ons-col width="20%">タグ：</ons-col>
                          <ons-col id="preview-category"></ons-col>
                        </ons-row>
                        <ons-row class="list-margin-bottom">
                          <ons-col width="20%">制作：</ons-col>
                          <ons-col id="preview-production"></ons-col>
                        </ons-row>
                        <ons-row class="list-margin-bottom">
                          <ons-col width="20%">監督：</ons-col>
                          <ons-col id="preview-supervisor"></ons-col>
                        </ons-row>
                        <ons-row class="list-margin-bottom">
                          <ons-col width="20%">脚本：</ons-col>
                          <ons-col id="preview-dramatist"></ons-col>
                        </ons-row>
                        <ons-row class="list-margin-bottom">
                          <ons-col width="20%">演出：</ons-col>
                          <ons-col id="preview-director"></ons-col>
                        </ons-row>
                      </div>
                    </ons-col>
                  </ons-row>
                  <ons-row>
                    <ons-col width="30%" align="center">
                      <div class="image" id="preview-img"></div>
                    </ons-col>
                    <ons-col width="70%" align="center">
                      <div class="div-iframe">
                        <div id="preview-movie"></div>
                      </div>
                    </ons-col>
                  </ons-row>
                  <ons-list-item expandable>
                    あらすじ・見どころ
                    <div class="right"></div>
                    <div class="expandable-content" id="preview-content"></div>
                  </ons-list-item>
                </div>
              </ons-list-item>
            </ons-list>
          </div>
        </ons-page>
      </ons-dialog>
    </template>
    <template id="tvprogram-delete-dialog.html">
      <ons-alert-dialog id="tvprogram-delete-dialog" modifier="rowfooter">
        <div class="alert-dialog-title">Alert</div>
        <div class="alert-dialog-content">
          本当に削除しますか？<br />番組・ブックマーク・閲覧ログ・アップデートログの情報が削除されます。
        </div>
        <div class="alert-dialog-footer">
          <ons-alert-dialog-button
            onclick="hideAlertDialog('tvprogram-delete-dialog')"
            >Cancel</ons-alert-dialog-button
          >
          <ons-alert-dialog-button>
            <form
              id="delete-tvprogram"
              action="/tv/tv_program/{{.TvProgram.Id}}"
              method="post"
              onSubmit="showLoading();"
            >
              <input type="hidden" name="_method" value="DELETE" />
              <button
                id="delete-tvprogram-button"
                class="button--quiet"
                type="submit"
              >
                OK
              </button>
            </form>
          </ons-alert-dialog-button>
        </div>
      </ons-alert-dialog>
    </template>

    {{ template "/common/js.tpl" . }}

    <script>
      let textTop = '<option>指定なし</option>';
      document.getElementById('hour').innerHTML = getSelectHour(textTop);
    </script>
    <script type="text/javascript">
      $('.restrict').change(function() {
        var count = $('.restrict option:selected').length;
        var not = $('.restrict option').not(':selected');
        if (count >= 3) {
          not.attr('disabled', true);
        } else {
          not.attr('disabled', false);
        }
      });
    </script>
    <script type="text/javascript">
      var previewTvProgram = function(elemID) {
        ons.ready(function() {
          var dialog = document.getElementById(elemID);
          if (dialog) {
            inputPreviewData();
            document.querySelector('#expandable-list-item').showExpansion();
            dialog.show();
          } else {
            ons
              .createElement(elemID + '.html', { append: true })
              .then(function(dialog) {
                inputPreviewData();
                document.querySelector('#expandable-list-item').showExpansion();
                dialog.show();
              });
          }
        });
      };
    </script>

    <script type="text/javascript">
      const tvProgram = {{.TvProgram}};
      if (tvProgram.Season != null){
          const seasonName = tvProgram.Season.Name;
          if (seasonName === "春"){
            document.getElementById('season').value = seasonName+"(4~6)";
          }
          else if (seasonName === "夏"){
            document.getElementById('season').value = seasonName+"(7~9)";
          }
          else if (seasonName === "秋"){
            document.getElementById('season').value = seasonName+"(10~12)";
          }
          else if (seasonName === "冬"){
            document.getElementById('season').value = seasonName+"(1~3)";
          } else {
            document.getElementById('season').value = seasonName;
          }
        }
      let time = String(tvProgram.Hour);
      str = ".5";
      if (time === "100"){
        time = "指定なし";
      } else {
        if (time.indexOf(str) > -1){
            time = time.replace(str, ":30")
        } else {
          time += ":00";
        }
      }
      document.getElementById('hour').value = time;
      setMultipleSelection("category", tvProgram.Category);
      document.getElementById('content').value = tvProgram.Content;
      if (tvProgram.Week != null){
          document.getElementById('week').value = tvProgram.Week.Name;
          }
      if ({{ .TitleFlag}}) {
        dialogBoxEveryone("alert-tv-title");
      };
    </script>
    <script>
      // うまくいかない．．．
      console.log({{ .TitleFlag}});
    </script>
  </body>
</html>
