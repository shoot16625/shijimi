<!DOCTYPE html>
<html lang="ja">
  <head>
    {{ template "/common/header.tpl" . }}
  </head>

  <body>
    <!-- <div class="floating-top"></div> -->
    <ons-page>
      {{ template "/common/toolbar.tpl" . }}
      {{ template "/common/alert.tpl" . }}

      <ons-pull-hook id="pull-hook"></ons-pull-hook>

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
            <ons-list-item>総レビュー数：{{ .CommentNum}}</ons-list-item>
            <ons-list-item>
              <div id="average-star"></div id="average-star">
              <!-- <i class="fas fa-star"></i>：{{ .TvProgram.Star}} -->
            </ons-list-item>
            <ons-list-item id="favorite-point-ranking-expandable" expandable>
              おすすめポイントランキング
              <div class="expandable-content">
                <ons-list modifier="inset" id="favorite-point-ranking">
                </ons-list>
              </div>
            </ons-list-item>
          </ons-list>
        </ons-carousel-item>
      </ons-carousel>
    </ons-page>

    {{ template "/common/js.tpl" . }}
    <script src="/static/js/star-raty/jquery.raty.js"></script>
    <script>
      $.fn.raty.defaults.path = '/static/js/star-raty/images';

      $(function raty() {
        $('#average-star').raty({
          number: 10,
          readOnly: true,
          precision: true,
          score: {{ .TvProgram.Star}},
        });
      });
    </script>

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

            <div class="right">
              <ons-button
                id="post-button"
                onclick="postComment()"
                style="color:chocolate;background:left;"
                ><i class="fas fa-paper-plane"></i
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
                minlength="20"
                maxlength="400"
                required
              ></textarea>
            </div>
            <div class="area-left">あと<span class="count"></span>文字</div>
            <div class="area-right">
              ネタバレ
              <ons-switch id="check-spoiler"> </ons-switch>
            </div>
            <div style="text-align: center;">
              <ons-row style="margin: 20px 0 20px 0;">
                <ons-col width="40px" class="area-center">
                  <ons-icon
                    icon="md-thumb-down"
                    style="color:rgb(71, 119, 121);"
                  ></ons-icon>
                </ons-col>
                <ons-col>
                  <div id="star-display" class="area-center">
                    <!-- <i class="fas fa-star"></i>： 5 -->
                  </div>
                  <!-- <ons-range
                    id="star-point"
                    name="star-point"
                    style="width: 100%;"
                    value="5"
                    min="0"
                    max="10"
                    step="1"
                    required
                  ></ons-range> -->
                </ons-col>
                <ons-col width="40px" class="area-center">
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
            let text_max = 400;
            $('.count').text(
              text_max - $('#tweet-dialog-content').val().length
            );

            $('#tweet-dialog-content').on(
              'keydown keyup keypress change',
              function() {
                let text_length = $(this).val().length;
                let countdown = text_max - text_length;
                $('.count').text(countdown);
              }
            );
          });

          $('.restrict').change(function() {
            let count = $('.restrict option:selected').length;
            let not = $('.restrict option').not(':selected');
            if (count >= 3) {
              not.attr('disabled', true);
            } else {
              not.attr('disabled', false);
            }
          });

          $(function raty() {
            $('#star-display').raty({
              number: 10,
              score: 5,
              hints: ['1', '2', '3', '4', '5', '6', '7', '8', '9', '10'],
              click: function(score, evt) {
                starScore = score;
              }
            });
          });

          // $('#star-point').change(function() {
          //   let star = $('#star-point').val();
          //   document.getElementById('star-display').innerHTML =
          //     '<i class="fas fa-star"></i>： ' + star;
          // });

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
              action="/tv/tv_program/review/search_comment/{{.TvProgram.Id}}"
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
                    min="1"
                    max="100"
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
        <script>
          var target = document.getElementById('star');
          let text = '';
          for (let i = 10; i >= 0; i--) {
            text += '<option>' + i + '</option>';
          }
          target.innerHTML = text;

          if ({{.SearchWords}} != null){
            setMultipleSelection("favorite-point", {{.SearchWords.Category}});
            setMultipleSelection("spoiler", {{.SearchWords.Spoiler}});
            setMultipleSelection("star", {{.SearchWords.Star}});
          }
          if ({{.SearchWords.Sortby}} != null){
            document.getElementById('sortby').value = {{.SearchWords.Sortby}};
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
                  複数選択方法(PC)：ctrlを押しながらクリック
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
        globalCommentLikeStatus = {{.CommentLike}};
        globalWatchStatus = {{.WatchStatus}};
        let comments = {{.Comment}};
        if (comments === null || comments.length === 0) {
          comments = null;
        }
        const users = {{.Users}};
        let commentLikes;
        if ({{.CommentLike}} === null && comments != null){
          // ログインしていない場合
          commentLikes = [comments.length];
          for (let i = comments.length - 1; i >= 0; i--) {
            commentLikes[i] = {Like:false};
          }
        } else {
          commentLikes = {{.CommentLike}};
        }
        let myUserID = {{.User.Id}};
              if (myUserID === null) {
                myUserID = "";
              }
        ons.ready(function() {
          var infiniteList = document.getElementById('comments');
          if (comments != null) {
            infiniteList.delegate = {
              createItemContent: function(i) {
              let fpText = reshapeFavoritePoint(comments[i]);
              if(comments[i].Spoiler){
                fpText += "<i class='fas fa-hand-paper' style='color:palevioletred;'></i>";
              }
              return ons.createElement('<div class="user-' + users[i].Id + '"><ons-list-header style="background-color:antiquewhite;text-transform:none;"><div class="area-left comment-list-header-font">@' + users[i].Username + '</div><div class="area-right list-margin">' + moment(comments[i].Created, "YYYY-MM-DDHH:mm:ss").format("YYYY/MM/DD HH:mm") + '</div></ons-list-header><ons-list-item><ons-row><ons-col width="15%"><i class="fas fa-star"></i>：' + comments[i].Star +'</ons-col><ons-col style="font-size:12px;">'+ fpText + '</ons-col></ons-row></ons-list-item><ons-list-item><div class="left"><a href="/tv/user/show/' + users[i].Id + '" title="' + users[i].Username + '"><img class="list-item__thumbnail" src="' + users[i].IconUrl + '" onerror="this.src=\'/static/img/user_img/s256_f_01.png\'"></a></div><div class="center"><span class="list-item__subtitle"id="comment-content-' + users[i].Id + '" style="font-size:14px;">' + comments[i].Content.replace(/(\r\n|\n|\r)/gm, "<br>") + '</span><span class="list-item__subtitle" class="area-right"><div style="float:right;" id="count-like-' + i + '">：' + comments[i].CountLike + '</div><div style="float:right;"><i class="' + setLikeBold(commentLikes[i].Like) + ' fa-thumbs-up" id="' + i + '" onclick="clickLike(this,myUserID,comments,\'review\')" style="color:' + setLikeStatus(commentLikes[i].Like, 'orchid') + ';"></i></div></span></div></ons-list-item></div>');
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

    <script type="text/javascript">
      setWatchBold("check-watched", {{.WatchStatus.Watched}});
      setWatchBold("check-wtw", {{.WatchStatus.WantToWatch}});
      setWatchStatus("check-watched", "deeppink", {{.WatchStatus.Watched}});
      setWatchStatus("check-wtw", "lightseagreen", {{.WatchStatus.WantToWatch}});
    </script>

    <!-- ツイートを保存する -->
    <script type="text/javascript">
      function postComment() {
        const text_length = document.getElementById("tweet-dialog-content").value.length;
        if (text_length < 5){
          return dialogBox('alert-min-length', {{.User.Id}}, 0);
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
        data.FavoritePoint = fps.join(",");
        if (typeof starScore === "undefined") {
          starScore = 5;
        }
        data.Star = starScore;
        var json = JSON.stringify(data);
        var request = new XMLHttpRequest();
        request.open('POST', url, true);
        request.setRequestHeader('Content-type','application/json; charset=utf-8');
        request.send(json);

        $('#tweet-dialog-content').val("");
        hideAlertDialog('tweet-dialog');
        document.querySelector('ons-speed-dial').hideItems();

        // $(".floating-top").html('<div class="toast"><div class="toast__message">リロードして反映してね <i class="fas fa-thumbs-up"></i></div></div>');
        // $(".floating-top").fadeIn();
        // setTimeout(function() {
        //   $(".floating-top").fadeOut();
        // }, 3000);

        setTimeout(function() {
          window.location.reload(false);
        }, 500);
      };
    </script>
    <script type="text/javascript">
      document
        .querySelector('ons-carousel')
        .addEventListener('postchange', function() {
          if (carousel.getActiveIndex() == 1) {
            goTop();
          }
        });

      // 選択の削除機能
      function resetSelect() {
        document.search_comment.reset();
        document.getElementById('word').value = '';
        document.getElementById('limit').value = '';
      }

        // スピードダイアルの表示
        var dial = document.getElementById('speed-dial');
        let userID = {{.User.Id}};
        if (comments != null){
        for (let index = 0; index < comments.length; index++) {
          if (comments[index].UserId === userID) {
            userID = -2;
          }
        }
      }
        dial.innerHTML =
          "<ons-fab><ons-icon icon='md-share'></ons-icon></ons-fab><ons-speed-dial-item><ons-icon icon='md-comment-dots' onclick='dialogBox(\"tweet-dialog\", "+userID+", {{.TvProgram.Id}})'></ons-icon></ons-speed-dial-item><ons-speed-dial-item><ons-icon icon='md-search' onclick='dialogBoxEveryone(\"search-dialog\")'></ons-icon></ons-speed-dial-item><ons-speed-dial-item><ons-icon icon='md-chart' onclick='goAnotherCarousel(1)'></ons-icon></ons-speed-dial-item><ons-speed-dial-item><i class='fas fa-arrow-up' onclick='goTop()'></i></ons-speed-dial-item>";
    </script>
    <script>
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
      let categories = {{.TvProgram.Category}}.split(',');
      if ({{.TvProgram.Category}} === ""){
        categories = [];
      }
      let category = reshapeCategory(categories);
      document.getElementById("tv-category").innerHTML = category;
      document.getElementById("tv-image").innerHTML = reshapeMovieCode({{ .TvProgram }});
      document.getElementById("tv-reference").innerHTML = reshapeReferenceSite({{ .TvProgram }});
    </script>
    <script>
        let fprank = {{ .FavoritePointRanking }};
        if (fprank === null) {
          fprank = "";
        }
        text = "";
        headerColor = ["mistyrose","seashell","lavenderblush","antiquewhite","azure"];
        for (let i = 0; i < fprank.length; i++) {
          text += '<ons-list-item style="background:'+headerColor[i]+'"><div class="left">'+fprank[i].Name+'</div><div class="center tv-program-list-content-font">' + fprank[i].Value + '</div></ons-list-item>';
        }
        document.querySelector('#favorite-point-ranking').innerHTML = text;
      document
        .querySelector('#favorite-point-ranking-expandable')
        .showExpansion();
    </script>
  </body>
</html>
