<!DOCTYPE html>
<html lang="ja">
  <head>
    {{ template "/common/header.tpl" . }}
  </head>

  <body>
    <ons-page>
      {{ template "/common/toolbar.tpl" . }}
      {{ template "/common/alert.tpl" . }}

      <ons-pull-hook id="pull-hook">
        Pull to refresh
      </ons-pull-hook>

      <ons-speed-dial
        id="speed-dial"
        position="bottom right"
        direction="up"
        ripple
      >
        <ons-fab></ons-fab>
      </ons-speed-dial>

      <ons-carousel
        swipeable
        overscrollable
        auto-scroll
        auto-refresh
        id="carousel"
      >
        <ons-carousel-item>
          {{ template "/common/tv_program_show.tpl" . }}
          {{ template "/common/comment_review_change.tpl" . }}

          <ons-list class="list-margin">
            <ons-lazy-repeat id="comments"></ons-lazy-repeat>
          </ons-list>
        </ons-carousel-item>
        <ons-carousel-item>
          <div class="area-center">
            <p>詳細情報や分析結果を表示</p>
            <p>工事中<i class="fas fa-truck-pickup"></i></p>
          </div>
        </ons-carousel-item>
      </ons-carousel>
    </ons-page>
    <template id="tweet-dialog.html">
      <ons-dialog id="tweet-dialog" modifier="large" cancelable fullscreen>
        <ons-page>
          <ons-toolbar>
            <div class="left">
              <ons-button
                id="cancel-button"
                onclick="hideAlertDialog('tweet-dialog')"
                style="background:left;color: grey;"
                ><i class="fas fa-window-close"></i
              ></ons-button>
            </div>

            <div class="right">
              <ons-button
                id="post-button"
                onclick="postComment()"
                style="color:chocolate;background:left;"
                ><i class="fas fa-book"></i
              ></ons-button>
            </div>
          </ons-toolbar>
          <div class="scroller">
            <div class="area-center">
              <textarea
                class="textarea-tweet"
                rows="10"
                id="tweet-dialog-content"
                name="content"
                type="text"
                minlength="5"
                maxlength="450"
                required
              ></textarea>
            </div>
            <div class="area-left">あと<span class="count"></span>文字</div>
            <div class="area-right">
              ネタバレ
              <ons-switch id="check-spoiler"> </ons-switch>
            </div>
            <div style="text-align: center;">
              <ons-row style="margin-top: 20px;">
                <ons-col
                  width="40px"
                  class="area-center"
                  style="line-height: 31px;"
                >
                  <ons-icon
                    icon="md-thumb-down"
                    style="color:rgb(71, 119, 121);"
                  ></ons-icon>
                </ons-col>
                <ons-col>
                  <label for="star-point">＜おすすめ度＞</label>
                  <ons-range
                    id="star-point"
                    name="star-point"
                    style="width: 100%;"
                    value="5"
                    min="0"
                    max="10"
                    step="1"
                    required
                  ></ons-range>
                  <div id="star-display" class="area-center">score： 5</div>
                </ons-col>
                <ons-col
                  width="40px"
                  class="area-center"
                  style="line-height: 31px;"
                >
                  <ons-icon
                    icon="md-thumb-up"
                    style="color:rgb(241, 161, 86);"
                  ></ons-icon>
                </ons-col>
              </ons-row>
              <p>
                <label for="favorite-point">＜おすすめポイント＞※3つまで</label>
                <select
                  name="favorite-point"
                  id="favorite-point"
                  class="select-input select-input--underbar select-search-table restrict"
                  style="height: 130px;"
                  required
                  multiple
                >
                  <option>演技すごい</option>
                  <option>配役ばっちり</option>
                  <option>見ごたえ鬼</option>
                  <option>伏線に次ぐ伏線</option>
                  <option>神曲</option>
                  <option>泣きっぱなし</option>
                  <option>超ギャグセンス</option>
                  <option>ゆる～い</option>
                  <option>キュン死</option>
                  <option>ホラーナイト</option>
                </select>
              </p>
            </div>
          </div>
        </ons-page>
        <script type="text/javascript">
          $(function() {
            var text_max = 450;
            $('.count').text(
              text_max - $('#tweet-dialog-content').val().length
            );

            $('#tweet-dialog-content').on(
              'keydown keyup keypress change',
              function() {
                var text_length = $(this).val().length;
                var countdown = text_max - text_length;
                $('.count').text(countdown);
              }
            );
          });
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
          $('#star-point').change(function() {
            var star = $('#star-point').val();
            document.getElementById('star-display').innerHTML =
              'score：' + star;
          });
        </script>
      </ons-dialog>
    </template>

    <template id="search-dialog.html">
      <ons-dialog id="search-dialog" modifier="large" cancelable fullscreen>
        <ons-page>
          <ons-toolbar>
            <div class="left">
              <ons-button
                id="cancel-button"
                onclick="hideAlertDialog('search-dialog')"
                style="background:left;color: grey;"
                ><i class="fas fa-window-close"></i
              ></ons-button>
            </div>
            <div class="center">
              <i class="fas fa-search" style="color: brown;"></i> 詳細検索
            </div>
            <div class="right">
              <ons-button
                id="reset-button"
                onclick="resetSelect()"
                style="color:chocolate;background:left;"
                ><i class="far fa-trash-alt"></i
              ></ons-button>
            </div>
          </ons-toolbar>
          <div class="scroller">
            <form
              name="search_comment"
              id="search_comment"
              action="/tv/tv_program/review/search_comment/{{.TvProgram.Id}}"
              method="post"
            >
              <div class="area-center create-top-margin">
                <p>
                  <ons-input
                    type="text"
                    name="word"
                    id="word"
                    value="{{.SearchWords.Word}}"
                    modifier="underbar"
                    placeholder="フリーワード"
                    float
                  ></ons-input>
                </p>
                <p>
                  <label for="star" style="margin-right:8px;margin-left:8px;"
                    >＜おすすめ度＞</label
                  >
                  <select
                    name="star"
                    id="star"
                    class="select-input select-input--underbar select-search-table"
                    style="height: 130px;"
                    multiple
                  >
                  </select>
                </p>
                <p>
                  <label for="favorite-point">＜おすすめポイント＞</label>
                  <select
                    name="favorite-point"
                    id="favorite-point"
                    class="select-input select-input--underbar select-search-table restrict"
                    style="height: 130px;"
                    multiple
                  >
                    <option>演技すごい</option>
                    <option>配役ばっちり</option>
                    <option>見ごたえ鬼</option>
                    <option>伏線に次ぐ伏線</option>
                    <option>神曲</option>
                    <option>泣きっぱなし</option>
                    <option>超ギャグセンス</option>
                    <option>ゆる～い</option>
                    <option>キュン死</option>
                    <option>ホラーナイト</option>
                  </select>
                </p>
                <p>
                  <label for="spoiler">＜ネタバレ＞</label>
                  <select
                    name="spoiler"
                    id="spoiler"
                    class="select-input select-input--underbar select-search-table"
                    multiple
                  >
                    <option>ネタバレなし</option>
                    <option>ネタバレあり</option>
                  </select>
                </p>
                <p>
                  <select
                    name="sortby"
                    id="sortby"
                    class="select-input select-input--underbar select-search-table"
                  >
                    <option>新しい順</option>
                    <option>古い順</option>
                    <option>いいねが多い順</option>
                    <option>評価が高い順</option>
                    <option>評価が低い順</option>
                  </select>
                </p>
                <p>
                  <ons-input
                    type="number"
                    name="limit"
                    id="limit"
                    modifier="underbar"
                    value="{{.SearchWords.Limit}}"
                    placeholder="表示数(デフォルト:100)"
                    float
                  ></ons-input>
                </p>
                <p class="create-top-margin">
                  <button class="button button--outline">search</button>
                </p>
              </div>
            </form>
          </div>
        </ons-page>
        <script>
          var target = document.getElementById('star');
          let text = '';
          for (let i = 10; i >= 0; i--) {
            text += '<option>' + i + '</option>';
          }
          target.innerHTML = text;
        </script>
        <script type="text/javascript">
          console.log({{.SearchWords}});
          if ({{.SearchWords}} != null){
            // console.log({{.SearchWords}});
            setMultipleSelection("favorite-point", {{.SearchWords.Category}});
            setMultipleSelection("spoiler", {{.SearchWords.Spoiler}});
            setMultipleSelection("star", {{.SearchWords.Star}});
          }
          if ({{.SearchWords.Sortby}} != null){
            document.getElementById('sortby').value = {{.SearchWords.Sortby}};
          }
        </script>
      </ons-dialog>
    </template>

    <script type="text/javascript" src="/static/js/common.js"></script>

    <script>
      ons.ready(function() {
        pullHook();
      });
    </script>

    <script>
        let comments = {{.Comment}};
        if (comments.length === 0) {
          comments = null;
        }
        const users = {{.Users}};
        let commentLikes;
        if ({{.CommentLike}} === null && comments != null){
          commentLikes = [comments.length];
          for (let i = comments.length - 1; i >= 0; i--) {
            commentLikes[i] = {Like:false};
          }
        } else {
          commentLikes = {{.CommentLike}};
        }
        ons.ready(function() {
          var infiniteList = document.getElementById('comments');
          if (comments != null) {
            infiniteList.delegate = {
              createItemContent: function(i) {
              const fps = comments[i].FavoritePoint.split('、');
              let fpText = "";
              for (let j = fps.length - 1; j >= 0; j--) {
                fpText += "<span style='padding:3px;color:blue;'>#"+fps[j]+"</span>";
              }
              if(comments[i].Spoiler){
                fpText += "<i class='fas fa-hand-paper' style='color:palevioletred;'></i>";
              }
              return ons.createElement('<div class="user-' + users[i].Id + '"><ons-list-header style="background-color:antiquewhite;text-transform:none;"><div class="area-left comment-list-header-font">@' + users[i].Username + '</div><div class="area-right list-margin">' + moment(comments[i].Created, "YYYY-MM-DDHH:mm:ss").format("YYYY/MM/DD HH:mm") + '</div></ons-list-header><ons-list-item><ons-row><ons-col width="15%"><i class="fas fa-star" style="color:gold;"></i>：' + comments[i].Star +'</ons-col><ons-col style="font-size:12px;">'+ fpText + '</ons-col></ons-row></ons-list-item><ons-list-item><div class="left"><a href="/tv/user/show/' + users[i].Id + '" title="' + users[i].Username + '"><img class="list-item__thumbnail" src="' + users[i].IconURL + '" alt="@' + users[i].Username + '"></a></div><div class="center"><span class="list-item__subtitle"id="comment-content-' + users[i].Id + '" style="font-size:14px;">' + comments[i].Content.replace(/(\r\n|\n|\r)/gm, "<br>") + '</span><span class="list-item__subtitle" class="area-right"><div style="float:right;" id="count-like-' + i + '">：' + comments[i].CountLike + '</div><div style="float:right;"><i class="' + setLikeBold(commentLikes[i].Like) + ' fa-thumbs-up" id="' + i + '" onclick="clickLike(this)" style="color:' + setLikeStatus(commentLikes[i].Like, 'orchid') + ';"></i></div></span></div></ons-list-item></div>');
            },
            countItems: function() {
              return comments.length;
            }
          };
          infiniteList.refresh();
          } else {
                  infiniteList.innerHTML = "<div style='text-align:center;margin-top:40px;'><i class='far fa-surprise' style='color:chocolate;'></i> Not Found !!</div>"
              }
      });
    </script>

    <script type="text/javascript">
      setWatchBold("check-watched", {{.WatchStatus.Watched}});
      setWatchBold("check-wtw", {{.WatchStatus.WantToWatch}});
      setWatchStatus("check-watched", "deeppink", {{.WatchStatus.Watched}});
      setWatchStatus("check-wtw", "lightseagreen", {{.WatchStatus.WantToWatch}});
    </script>

    <script>
      globalCommentLikeStatus = {{.CommentLike}};
    </script>

    <script type="text/javascript">
      function commentLikeStatus(elem, checkFlag) {
        let url = URL+"/tv/review_comment_like/";
        var data = globalCommentLikeStatus[elem.id];
        let method;
        if (data.Id === 0){
          method = 'POST';
          data.UserId = {{.User.Id}};
          globalCommentLikeStatus[elem.id].UserId = data.UserId;
          data.ReviewCommentId = {{.Comment}}[elem.id].Id;
          globalCommentLikeStatus[elem.id].ReviewCommentId = data.ReviewCommentId;
        } else {
          method = 'PUT';
          url = url+data.Id;
        }
        data.Like = checkFlag;
          globalCommentLikeStatus[elem.id].Like = data.Like;
          var json = JSON.stringify(data);
          var request = new XMLHttpRequest();
          request.open(method, url, true);
          request.setRequestHeader('Content-type','application/json; charset=utf-8');
          request.onload = function () {
            var x = JSON.parse(request.responseText);
            if (request.readyState == 4 && request.status == "200") {
              // console.table(x);
            } else {
              globalCommentLikeStatus[elem.id].Id = x.Id;
            }
          }
          request.send(json);
        };
    </script>

    <script type="text/javascript">
      globalWatchStatus = {{.WatchStatus}};
    </script>

    <script type="text/javascript">
      function WatchStatus(elem, checkFlag) {
        let url = URL+"/tv/watching_status/";
        var data = globalWatchStatus;
        let method;
        if (data.Id === 0){
          method = 'POST';
          data.UserId = {{.User.Id}};
          globalWatchStatus.UserId = data.UserId;
          data.TvProgramId = {{.TvProgram.Id}};
          globalWatchStatus.TvProgramId = data.TvProgramId;
        } else {
          method = 'PUT';
          url = url+data.Id;
        }
        const str ="check-watched";
        if (elem.id.indexOf(str)===0) {
          data.Watched = checkFlag;
          globalWatchStatus.Watched = data.Watched;
        } else {
          data.WantToWatch = checkFlag;
          globalWatchStatus.WantToWatch = data.WantToWatch;

        }
        var json = JSON.stringify(data);
        var request = new XMLHttpRequest();
        request.open(method, url, true);
        request.setRequestHeader('Content-type','application/json; charset=utf-8');
        request.onload = function () {
          var x = JSON.parse(request.responseText);
          if (request.readyState == 4 && request.status == "200") {
            console.table(x);
          } else {
            globalWatchStatus.Id = x.Id;
          }
        }
        request.send(json);
      };
    </script>

    <!-- ツイートを保存する -->
    <script type="text/javascript">
      function postComment() {
        const text_length = document.getElementById("tweet-dialog-content").value.length;
        if (text_length < 5){
          return dialogBox('alert-min-length');
        }
        let url = URL+"/tv/review_comment/";
        let data = {};
        data.Id  = 0;
        data.UserId = {{.User.Id}};
        data.TvProgramId  = {{.TvProgram.Id}};
        data.Content = document.getElementById("tweet-dialog-content").value;
        data.CountLike  = 0;
        data.Spoiler  = $("#check-spoiler").prop("checked");
        var fp = document.getElementById("favorite-point");
        let fps = [];
        for (let i = fp.length - 1; i >= 0; i--) {
          if (fp[i].selected){
            fps.push(fp[i].value);
          }
        }
        data.FavoritePoint = fps.join("、");
        data.Star = Number(document.getElementById("star-point").value);
        // console.log(data);
        var json = JSON.stringify(data);
        var request = new XMLHttpRequest();
        request.open('POST', url, true);
        request.setRequestHeader('Content-type','application/json; charset=utf-8');
        request.onload = function () {
          var x = JSON.parse(request.responseText);
          if (request.readyState == 4 && request.status == "200") {
            console.table(x);
          } else {
            console.error(x);
          }
        }
        request.send(json);
        // PostRatingTvProgram();
        hideAlertDialog('tweet-dialog')
        setTimeout(window.location.reload(false), 500);
      };
    </script>
    <script>
      let time = String({{.TvProgram.Hour}});
      str = ".5";
      if (time === "100"){
        time = "";
      } else {
        if (time.indexOf(str) > -1){
          time = time.replace(str, ":30")
        } else {
          time += ":00";
        }
      }
      document.getElementById('tv-program-hour').innerHTML = time;
    </script>
    <script type="text/javascript">
      document
        .querySelector('ons-carousel')
        .addEventListener('postchange', function() {
          if (carousel.getActiveIndex() == 1) {
            goTop();
          }
        });
    </script>
    <script type="text/javascript">
      function resetSelect() {
        document.search_comment.reset();
        document.getElementById('word').value = '';
        document.getElementById('limit').value = '';
      }
    </script>
    <script>
      var dial = document.getElementById('speed-dial');
      dial.innerHTML =
        "<ons-fab><ons-icon icon='md-share'></ons-icon></ons-fab><ons-speed-dial-item><ons-icon icon='md-comment-dots' onclick='dialogBox(\"tweet-dialog\", {{.User.Id}})'></ons-icon></ons-speed-dial-item><ons-speed-dial-item><ons-icon icon='md-search' onclick='dialogBoxEveryone(\"search-dialog\")'></ons-icon></ons-speed-dial-item><ons-speed-dial-item><ons-icon icon='md-chart' onclick='goAnotherCarousel(1)'></ons-icon></ons-speed-dial-item><ons-speed-dial-item><ons-icon icon='md-home' onclick='goTop()'></ons-icon></ons-speed-dial-item>";
    </script>
  </body>
</html>
