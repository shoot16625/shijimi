<!DOCTYPE html>
<html lang="ja">
  <head>
    {{ template "/common/header.tpl" . }}
  </head>

  <body>
    <div class="floating-top"></div>
    <ons-page id="tv-comments">
      {{ template "/common/toolbar.tpl" . }}
      {{ template "/common/alert.tpl" . }}

      <ons-pull-hook id="pull-hook">
        <!-- Pull to refresh -->
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
          <ons-list>
            <ons-list-item>総コメント数：{{ .CommentNum}}</ons-list-item>
          </ons-list>
        </ons-carousel-item>
      </ons-carousel>
    </ons-page>

    {{ template "/common/js.tpl" . }}

    <template id="tweet-dialog.html">
      <ons-dialog id="tweet-dialog" modifier="large" cancelable fullscreen>
        <ons-page>
          <ons-toolbar id="tweet-hide-swipe">
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
                ><i class="fas fa-paper-plane"></i
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
              maxlength="180"
              required
            ></textarea>
          </div>
          <div class="area-right">あと<span class="count"></span>文字</div>
        </ons-page>
        <script type="text/javascript">
          $(function() {
            const textMax = 180;
            let textLength;
            let countdown;
            $('.count').text(textMax - $('#tweet-dialog-content').val().length);

            $('#tweet-dialog-content').on(
              'keydown keyup keypress change',
              function() {
                textLength = $(this).val().length;
                countdown = textMax - textLength;
                $('.count').text(countdown);
              }
            );
          });

          hideSwipeToolbar('tweet-hide-swipe', 'tweet-dialog');
        </script>
      </ons-dialog>
    </template>

    <template id="search-dialog.html">
      <ons-dialog id="search-dialog" modifier="large" cancelable fullscreen>
        <ons-page>
          <ons-toolbar id="search-hide-swipe">
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
              <div class="area-center create-top-bottom-margin">
                <div class="area-right" style="height: 0px;">
                  <ons-button
                    modifier="quiet"
                    onclick="dialogBoxEveryone('hint-dialog')"
                    ><i class="far fa-question-circle hint-icon-right"></i
                  ></ons-button>
                </div>
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
                  <ons-input
                    type="text"
                    name="username"
                    id="username"
                    value="{{.SearchWords.Username}}"
                    modifier="underbar"
                    placeholder="ユーザーを指定"
                    float
                  ></ons-input>
                </p>
                <p>
                  時間帯
                  <ons-row>
                    <ons-col width="47%">
                      <ons-input
                        type="date"
                        name="before-date"
                        id="before-date"
                        value="{{.SearchWords.BeforeDate}}"
                        modifier="underbar"
                        float
                      ></ons-input>
                      <ons-input
                        type="time"
                        name="before-time"
                        id="before-time"
                        value="{{.SearchWords.BeforeTime}}"
                        modifier="underbar"
                      ></ons-input>
                    </ons-col>
                    <ons-col width="6%">～</ons-col>
                    <ons-col>
                      <ons-input
                        type="date"
                        name="after-date"
                        id="after-date"
                        value="{{.SearchWords.AfterDate}}"
                        modifier="underbar"
                      ></ons-input>
                      <ons-input
                        type="time"
                        name="after-time"
                        id="after-time"
                        value="{{.SearchWords.AfterTime}}"
                        modifier="underbar"
                      ></ons-input>
                    </ons-col>
                  </ons-row>
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
                    placeholder="表示数(デフォルト:200)"
                    min="1"
                    max="200"
                    float
                  ></ons-input>
                </p>
                <p class="create-top-bottom-margin">
                  <button class="button button--outline">search</button>
                </p>
              </div>
            </form>
          </div>
        </ons-page>
        <script type="text/javascript">
          if ({{.SearchWords.Sortby}} != null){
            document.getElementById('sortby').value = {{.SearchWords.Sortby}};
          }

          let today = new Date();
          today.setDate(today.getDate());
          let yyyy = today.getFullYear();
          let mm = ('0' + (today.getMonth() + 1)).slice(-2);
          let dd = ('0' + today.getDate()).slice(-2);
          let h = ('0' + today.getHours()).slice(-2);
          let min = ('0' + today.getMinutes()).slice(-2);
          let date = yyyy + '-' + mm + '-' + dd;
          // 1分後を指定
          let minutes = parseInt(min, 10) + 1
          let time = h + ':' + String(minutes);
          if({{.SearchWords.BeforeDate}}===null||{{.SearchWords.BeforeDate}}===""){
            document.getElementById('before-date').value = date;
          }
          // if({{.SearchWords.BeforeTime}}===null||{{.SearchWords.BeforeTime}}===""){
          //   document.getElementById('before-time').value = time;
          // }
          if({{.SearchWords.AfterDate}}===null||{{.SearchWords.AfterDate}}===""){
            document.getElementById('after-date').value = date;
          }
          if({{.SearchWords.AfterTime}}===null||{{.SearchWords.AfterTime}}===""){
            document.getElementById('after-time').value = time;
          }

          hideSwipeToolbar('search-hide-swipe', 'search-dialog');
        </script>
      </ons-dialog>
    </template>
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
                検索
              </li>
              <li class="list-item hint-list-dialog">
                <div class="list-item__center">
                  ユーザーを指定：OR検索です（スペース区切り）
                </div>
              </li>
              <li class="list-item">
                <div class="list-item__center hint-list-dialog">
                  時間帯指定と古い順を組み合わせれば、いつでもリアルタイムにタイムラインを眺めることができます。
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
    <script>
      ons.ready(function() {
        pullHook();
      });
    </script>
    <script>
      comments = {{.Comment}};
      if (comments === null || comments.length === 0) {
        comments = null;
      }
      users = {{.Users}};
      if ({{.CommentLike}} === null && comments != null){
        // commentLike：グローバル
        // ログインしていない場合
        commentLikes = [comments.length];
        for (let i = comments.length - 1; i >= 0; i--) {
          commentLikes[i] = {Like:false};
        }
      } else {
        commentLikes = {{.CommentLike}};
      }
      globalCommentLikeStatus = commentLikes;
      globalWatchStatus = {{.WatchStatus}};
      let myUserID = {{.User.Id}};
      if (myUserID === null) {
        myUserID = "";
      }
      ons.ready(function() {
        var infiniteList = document.getElementById('comments');
        if (comments != null) {

        infiniteList.delegate = {
          createItemContent: function(i) {

            return ons.createElement('<div id="commentID-' + comments[i].Id + '"><ons-list-header style="background-color:aliceblue;text-transform:none;"><div class="area-left comment-list-header-font">@' + users[i].Username + '</div><div class="area-right list-margin">' + moment(comments[i].Created, "YYYY-MM-DDHH:mm:ss").format("YYYY/MM/DD HH:mm") + '</div></ons-list-header><ons-list-item><div class="left"><a href="/tv/user/show/' + users[i].Id + '" title="user_comment"><img class="list-item__thumbnail" src="' + users[i].IconUrl + '" onerror="this.src=\'/static/img/user_img/s256_f_01.png\'"></a></div><div class="center"><span class="list-item__subtitle comment-list-content-font" id="comment-content-' + String(i) + '">' + comments[i].Content.replace(/(\r\n|\n|\r)/gm, "<br>") + '</span><span class="list-item__subtitle area-right"><div style="float:right;" id="count-like-' + i + '">：' + comments[i].CountLike + '</div><div style="float:right;"><i class="' + setLikeBold(commentLikes[i].Like) + ' fa-thumbs-up" id="' + i + '" onclick="clickLike(this,myUserID,comments,\'comment\')" style="color:' + setLikeStatus(commentLikes[i].Like, 'orchid') + ';"></i></div></span></div></ons-list-item></div>');
          },
          countItems: function() {
            return comments.length;
          }
        };
        infiniteList.refresh();
        } else {
          infiniteList.innerHTML = "<div style='text-align:center;margin-top:40px;'><i class='far fa-surprise' style='color:chocolate;'></i> Not Found !!</div>";
        }
      });
    </script>

    <script>
      // 非同期通信でのコメント取得・更新
      function getNewComment(){
        if ({{.SearchWords}}===null && comments != null){
        let url = URL+"/tv/comment/update/"+{{.TvProgram.Id}}+"/"+comments[0].Id;
        let method = "GET"
        var request = new XMLHttpRequest();
        let pos;
        request.open(method, url, true);
        request.setRequestHeader('Content-type','application/json; charset=utf-8');
        request.send();
        request.onreadystatechange = function() {
          if(request.readyState === 4 && request.status === 200) {
            pos = $('.page__content').scrollTop();
            let commentsAndUsers = JSON.parse(request.responseText);
            if(commentsAndUsers.Comments != null){
              let newComments = commentsAndUsers.Comments;
              let newUsers = commentsAndUsers.Users;
              let newCommentLikes = commentsAndUsers.CommentLikes;
              comments = newComments.concat(comments);
              if (comments === null || comments.length === 0) {
                comments = null;
              }
              commentLikes = newCommentLikes.concat(globalCommentLikeStatus);
              globalCommentLikeStatus = commentLikes;
              let myUserID = {{.User.Id}};
              if (myUserID === null) {
                myUserID = "";
              }
              users = newUsers.concat(users);
              var infiniteList = document.getElementById('comments');
              if (comments != null) {

              infiniteList.delegate = {
                createItemContent: function(i) {

                  return ons.createElement('<div id="commentID-' + comments[i].Id + '"><ons-list-header style="background-color:aliceblue;text-transform:none;"><div class="area-left comment-list-header-font">@' + users[i].Username + '</div><div class="area-right list-margin">' + moment(comments[i].Created, "YYYY-MM-DDHH:mm:ss").format("YYYY/MM/DD HH:mm") + '</div></ons-list-header><ons-list-item><div class="left"><a href="/tv/user/show/' + users[i].Id + '" title="user_comment"><img class="list-item__thumbnail" src="' + users[i].IconUrl + '" onerror="this.src=\'/static/img/user_img/s256_f_01.png\'"></a></div><div class="center"><span class="list-item__subtitle comment-list-content-font" id="comment-content-' + String(i) + '">' + comments[i].Content.replace(/(\r\n|\n|\r)/gm, "<br>") + '</span><span class="list-item__subtitle area-right"><div style="float:right;" id="count-like-' + i + '">：' + comments[i].CountLike + '</div><div style="float:right;"><i class="' + setLikeBold(commentLikes[i].Like) + ' fa-thumbs-up" id="' + i + '" onclick="clickLike(this,myUserID,comments,\'comment\')" style="color:' + setLikeStatus(commentLikes[i].Like, 'orchid') + ';"></i></div></span></div></ons-list-item></div>');
                },
                countItems: function() {
                  return comments.length;
                }
              };
              infiniteList.refresh();
              if (newComments && carousel.getActiveIndex() === 0) {
                $(".floating-top").html('<div class="toast"><div class="toast__message">' + newComments.length + ' New Comments !!<button class="toast-hide-button" onclick="hideToast(\'.floating-top\')">ok</button></div></div>');
                $(".floating-top").fadeIn();
                setTimeout(function() {
                  $(".floating-top").fadeOut();
                }, 3000);
              }
              }
              let firstOldCommentId = comments[newComments.length].Id;
              let firstNewCommentId = newComments[0].Id;
              let firstOldCommentTopPos = $('#commentID-'+firstOldCommentId).offset().top;
              let firstNewCommentTopPos = $('#commentID-'+firstNewCommentId).offset().top;
              scrollToTarget(pos + firstOldCommentTopPos - firstNewCommentTopPos);
            }
          }
        }
      }
      };
      // setInterval(getNewComment, 30000);
      setInterval(getNewComment, 10000);
    </script>

    <script type="text/javascript">
      setWatchBold("check-watched", {{.WatchStatus.Watched}});
      setWatchBold("check-wtw", {{.WatchStatus.WantToWatch}});
      setWatchStatus("check-watched", "deeppink", {{.WatchStatus.Watched}});
      setWatchStatus("check-wtw", "lightseagreen", {{.WatchStatus.WantToWatch}});
    </script>

    <script type="text/javascript">
      // ツイートを保存する
      function postComment() {
        const textLength = document.getElementById("tweet-dialog-content").value.length;
        if (textLength < 5){
          return dialogBox('alert-min-length', {{.User.Id}}, 0);
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
        request.send(json);
        // request.onload = function () {
        //   console.log(request.responseText);
        //   var x = JSON.parse(request.responseText);
        //   if (request.readyState == 4 && request.status == "200") {
        //     console.table(x);
        //   } else {
        //     console.error(x);
        //   }
        // }
        $('#tweet-dialog-content').val("");
        hideAlertDialog('tweet-dialog');
        document.querySelector('ons-speed-dial').hideItems();
        setTimeout(window.location.reload(false), 1000);
      };

      document
        .querySelector('ons-carousel')
        .addEventListener('postchange', function() {
          if (carousel.getActiveIndex() === 1) {
            goTop();
          }
        });

      function resetSelect() {
        document.search_comment.reset();
        document.getElementById('word').value = '';
        document.getElementById('username').value = '';
        document.getElementById('before-date').value = '';
        document.getElementById('before-time').value = '';
        document.getElementById('after-date').value = '';
        document.getElementById('after-time').value = '';
        document.getElementById('limit').value = '';
      }

      var dial = document.getElementById('speed-dial');
      let userID = {{.User.Id}};
      dial.innerHTML =
        "<ons-fab><ons-icon icon='md-share'></ons-icon></ons-fab><ons-speed-dial-item><ons-icon icon='md-comment-dots' onclick='dialogBox(\"tweet-dialog\", " + userID + ", {{.TvProgram.Id}})'></ons-icon></ons-speed-dial-item><ons-speed-dial-item><ons-icon icon='md-search' onclick='dialogBoxEveryone(\"search-dialog\")'></ons-icon></ons-speed-dial-item><ons-speed-dial-item><ons-icon icon='md-chart' onclick='goAnotherCarousel(1)'></ons-icon></ons-speed-dial-item><ons-speed-dial-item><i class='fas fa-arrow-up' onclick='goTop()'></i></ons-speed-dial-item>";

      document.getElementById('tv-program-hour').innerHTML = reshapeHour(String({{.TvProgram.Hour}}))+"）";
      let seasonName = avoidStructNameError({{.TvProgram.Season}});
      let weekName = avoidStructNameError({{.TvProgram.Week}});
      document.getElementById('tv-program-week').innerHTML = "{{.TvProgram.Year}}年 "+seasonName+"（"+weekName;
      document.getElementById("tv-cast").innerHTML = reshapeContent({{.TvProgram.Cast}});
      document.getElementById("tv-themesong").innerHTML = reshapeContent({{.TvProgram.Themesong}});
      document.getElementById("tv-supervisor").innerHTML = reshapeContent({{.TvProgram.Supervisor}});
      document.getElementById("tv-dramatist").innerHTML = reshapeContent({{.TvProgram.Dramatist}});
      document.getElementById("tv-director").innerHTML = reshapeContent({{.TvProgram.Director}});
      document.getElementById("tv-production").innerHTML = reshapeContent({{.TvProgram.Production}});
      document.getElementById("tv-image").innerHTML = reshapeMovieCode({{ .TvProgram }});
      document.getElementById("tv-reference").innerHTML = reshapeReferenceSite({{ .TvProgram }});
      let categories = {{.TvProgram.Category}}.split(',');
      if ({{.TvProgram.Category}} === ""){
        categories = [];
      }
      let category = reshapeCategory(categories);
      document.getElementById("tv-category").innerHTML = category;
    </script>
  </body>
</html>
