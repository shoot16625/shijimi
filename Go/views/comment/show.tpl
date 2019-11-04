<!DOCTYPE html>
<html lang="ja">
  <head>
    {{ template "/common/header.tpl" . }}
  </head>

  <body>
    <ons-page id="tv-comments">
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
        <ons-fab> </ons-fab>
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
          <div style="height: 200px; padding: 1px 0 0 0;">
            <div class="card" style="background-color: rgb(249, 252, 255);">
              <h2 class="card__title">工事中<i class="fas fa-wrench"></i></h2>
              <div class="card__content">
                ログの解析結果を表示.
              </div>
            </div>
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
            <div class="center">
              ポスト
            </div>
            <div class="right">
              <ons-button
                id="post_button"
                onclick="postComment()"
                style="color:chocolate;background:left;"
                ><i class="fas fa-book"></i
              ></ons-button>
            </div>
          </ons-toolbar>
          <div class="area-center">
            <textarea
              class="textarea-tweet"
              rows="10"
              id="tweet-dialog-content"
              name="content"
              type="text"
              minlength="5"
              maxlength="250"
              required
            ></textarea>
          </div>
          <div class="area-right">あと<span class="count"></span>文字</div>
        </ons-page>
        <script type="text/javascript">
          $(function() {
            const text_max = 250;
            let text_length;
            let countdown;
            $('.count').text(
              text_max - $('#tweet-dialog-content').val().length
            );

            $('#tweet-dialog-content').on(
              'keydown keyup keypress change',
              function() {
                text_length = $(this).val().length;
                countdown = text_max - text_length;
                $('.count').text(countdown);
              }
            );
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
              action="/tv/tv_program/comment/search_comment/{{.TvProgram.Id}}"
              method="post"
            >
              <div style="text-align: center; margin-top: 30px;">
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
                  <select
                    name="sortby"
                    id="sortby"
                    class="select-input select-input--underbar select-search-table"
                  >
                    <option>新しい順</option>
                    <option>古い順</option>
                    <option>いいねが多い順</option>
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
        <script type="text/javascript">
          console.log({{.SearchWords}});
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
      if ({{.CommentLike}} ===null && comments != null){
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

            return ons.createElement('<div class="user-' + users[i].Id + '"><ons-list-header style="background-color:aliceblue;text-transform:none;"><div class="area-left comment-list-header-font">@' + users[i].Username + '</div><div class="area-right list-margin">' + moment(comments[i].Created, "YYYY-MM-DDHH:mm:ss").format("YYYY/MM/DD HH:mm") + '</div></ons-list-header><ons-list-item><div class="left"><a href="/tv/user/show/' + users[i].Id + '" title="user_comment"><img class="list-item__thumbnail" src="' + users[i].IconURL + '" alt="@' + users[i].Username + '"></a></div><div class="center"><span class="list-item__subtitle comment-list-content-font" id="comment-content-' + String(i) + '">' + comments[i].Content.replace(/(\r\n|\n|\r)/gm, "<br>") + '</span><span class="list-item__subtitle area-right"><div style="float:right;" id="count-like-' + i + '">：' + comments[i].CountLike + '</div><div style="float:right;"><i class="' + setLikeBold(commentLikes[i].Like) + ' fa-thumbs-up" id="' + i + '" onclick="clickLike(this)" style="color:' + setLikeStatus(commentLikes[i].Like, 'orchid') + ';"></i></div></span></div></ons-list-item></div>');
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
        let url = URL+"/tv/comment_like/";
        var data = globalCommentLikeStatus[elem.id];
        let method;
        if (data.Id === 0){
          method = 'POST';
          data.UserId = {{.User.Id}};
          globalCommentLikeStatus[elem.id].UserId = data.UserId;
          data.CommentId = {{.Comment}}[elem.id].Id;
          globalCommentLikeStatus[elem.id].CommentId = data.CommentId;
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
        } else{
          method = 'PUT';
          url = url+data.Id;
        }
        const str ="check-watched"
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
          return dialogBox('alert-min-length', {{.User.Id}});
        }
        let url = URL+"/tv/comment/";
        let data = {};
        data.Id  = 0;
        data.UserId = {{.User.Id}};
        data.TvProgramId  = {{.TvProgram.Id}};
        data.Content = document.getElementById("tweet-dialog-content").value;
        data.CountLike  = 0;
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
        "<ons-fab><ons-icon icon='md-share'></ons-icon></ons-fab><ons-speed-dial-item><ons-icon icon='md-comment-dots' onclick='dialogBox(\"tweet-dialog\", {{.User.Id}})'></ons-icon></ons-speed-dial-item><ons-speed-dial-item><ons-icon icon='md-search' onclick='dialogBoxEveryone(\"search-dialog\")'></ons-icon></ons-speed-dial-item><ons-speed-dial-item><ons-icon icon='md-chart' onclick='goAnotherCarousel(1)'></ons-icon></ons-speed-dial-item><ons-speed-dial-item><i class='fas fa-arrow-up' onclick='goTop()'></i></ons-speed-dial-item>";
    </script>
    <script>
      let categories = {{.TvProgram.Category}}.split('、');
      if ({{.TvProgram.Category}} === ""){
        categories = [];
      }
      let category = "";
      for (let j = categories.length - 1; j >= 0; j--) {
        category += "<span style='padding:3px;'>#"+categories[j]+"</span>";
      }
      document.getElementById("category-area").innerHTML = category;
    </script>
  </body>
</html>
