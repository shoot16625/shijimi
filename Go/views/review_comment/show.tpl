<!DOCTYPE html>
<html lang="ja">
<head>
  {{ template "/common/header.tpl" . }}
</head>

<style type="text/css">
.textarea {
  width: 100%;
  background-color: white;
}
select {
  width:80%;
  max-width: 500px;
  height: 100px;
}
</style>

<body>
  <ons-page>

    {{ template "/common/toolbar.tpl" . }}
    {{ template "/common/alert.tpl" . }}

    <ons-pull-hook id="pull-hook">
      Pull to refresh
    </ons-pull-hook>

    <ons-speed-dial position="bottom right" direction="up" ripple>
        <ons-fab>
          <ons-icon icon="md-share"></ons-icon>
        </ons-fab>
        <ons-speed-dial-item>
          <ons-icon
            icon="md-comment-dots"
            onclick="DialogBox('tweet_dialog', {{.User.Id}})"
          ></ons-icon>
        </ons-speed-dial-item>
        <ons-speed-dial-item>
          <ons-icon
            icon="md-search"
            onclick="DialogBoxEveryone('search_dialog')"
          ></ons-icon>
        </ons-speed-dial-item>
        <ons-speed-dial-item>
          <ons-icon icon="md-chart" onclick="GoAnotherCarousel(1)"></ons-icon>
        </ons-speed-dial-item>
        <ons-speed-dial-item>
          <ons-icon icon="md-home" onclick="GoTop()"></ons-icon>
        </ons-speed-dial-item>
      </ons-speed-dial>

    <ons-carousel swipeable overscrollable auto-scroll auto-refresh id="carousel">
      <ons-carousel-item>
       {{ template "/common/comment_review_change.tpl" . }}
       {{ template "/common/tv_program_show.tpl" . }}

       <ons-list style="margin-left: 3px;margin-right: 5px;">
        <ons-lazy-repeat id="comments"></ons-lazy-repeat>
      </ons-list>
    </ons-carousel-item>
    <ons-carousel-item>
        <p style="text-align:center;">詳細情報や分析結果を表示</p>
        <p style="text-align:center;">
          工事中<i class="fas fa-truck-pickup"></i>
        </p>
    </ons-carousel-item>
  </ons-carousel>
</ons-page>
<template id="tweet_dialog.html">
  <ons-dialog id="tweet_dialog" modifier="large" cancelable fullscreen>
    <ons-page>
      <ons-toolbar>
        <div class="left">
          <ons-button id="cancel_button" onclick="hideAlertDialog('tweet_dialog')" style="background:left;color: grey;"><i class="fas fa-window-close"></i></ons-button>
        </div>

        <div class="right">
          <ons-button id="post_button" onclick="PostComment()" style="color:chocolate;background:left;"><i class="fas fa-book"></i></ons-button>
        </div>
      </ons-toolbar>
      <div class="scroller">
      <div  style="text-align:center;">
        <textarea class="textarea" rows="10" id="tweet_dialog_content" name="content" type="text" minlength="5" maxlength="450" required></textarea>
      </div>
      <div style="text-align:left;float: left;">あと<span class="count"></span>文字</div>
      <div style="text-align:right;">
        ネタバレ
        <ons-switch id="check_spoiler">
        </ons-switch>
      </div>
      <div style="text-align: center;">
        <ons-row style="margin-top: 20px;">
          <ons-col width="40px" style="text-align: center; line-height: 31px;">
            <ons-icon icon="md-thumb-down" style="color:cadetblue;"></ons-icon>
          </ons-col>
          <ons-col>
           <label for="StarPoint">＜おすすめ度＞</label>
           <ons-range id="StarPoint" name="StarPoint" style="width: 100%;" value="5" min="0" max="10" step="1" required></ons-range>
           <div id="star_display" style="text-align:center;">score： 5
           </div>
         </ons-col>
         <ons-col width="40px" style="text-align: center; line-height: 31px;">
          <ons-icon icon="md-thumb-up" style="color:violet;"></ons-icon>
        </ons-col>
      </ons-row>
      <p>
       <label for="FavoritePoint">＜おすすめポイント＞※3つまで</label>
       <select name="FavoritePoint" id="FavoritePoint" class="select-input select-input--underbar restrict" style="height: 100px;" required multiple>
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
  $(function(){
    var text_max = 450;
    $(".count").text(text_max - $("#tweet_dialog_content").val().length);

    $("#tweet_dialog_content").on("keydown keyup keypress change",function(){
      var text_length = $(this).val().length;
      var countdown = text_max - text_length;
      $(".count").text(countdown);
    });
  });
</script>
<script type="text/javascript">
  $('.restrict').change(function() {
    var count = $(".restrict option:selected").length;
    var not = $('.restrict option').not(':selected');
    if(count >= 3) {
      not.attr("disabled",true);
    }else{
      not.attr("disabled",false);
    }
  });
</script>
<script type="text/javascript">
  $("#StarPoint").change(function() {
    var star = $("#StarPoint").val();
    document.getElementById("star_display").innerHTML = "score："+star;
  });
</script>
</ons-dialog>
</template>

<template id="search_dialog.html">
  <ons-dialog id="search_dialog" modifier="large" cancelable fullscreen>
    <ons-page>
      <ons-toolbar>
        <div class="left">
          <ons-button id="cancel_button" onclick="hideAlertDialog('search_dialog')" style="background:left;color: grey;"><i class="fas fa-window-close"></i></ons-button>
        </div>
        <div class="center">
          <i class="fas fa-search" style="color: brown;"></i> 詳細検索
        </div>
        <div class="right">
          <ons-button id="reset_button" onclick="resetSelect()" style="color:chocolate;background:left;"><i class="far fa-trash-alt"></i></ons-button>
        </div>
      </ons-toolbar>
      <div class="scroller">
        <form name="search_comment" id="search_comment" action="/tv/tv_program/review/search_comment/{{.TvProgram.Id}}" method="post">
          <div style="text-align: center; margin-top: 30px;">
            <p>
              <ons-input type="text" name="word" id="word" value="{{.SearchWords.Word}}" modifier="underbar" placeholder="フリーワード" float></ons-input>
            </p>
            <p>
              <label for="star" style="margin-right:8px;margin-left:8px;">＜おすすめ度＞</label>
              <select name="star" id="star" class="select-input select-input--underbar" multiple>
              </select>
            </p>
            <p>
             <label for="FavoritePoint">＜おすすめポイント＞</label>
             <select name="FavoritePoint" id="FavoritePoint" class="select-input select-input--underbar restrict" style="height: 100px;" multiple>
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
            <label for="FavoritePoint">＜ネタバレ＞</label>
            <select name="spoiler" id="spoiler" class="select-input select-input--underbar" multiple>
              <option>ネタバレなし</option>
              <option>ネタバレあり</option>
            </select>
          </p>
          <p>
            <select name="sortby" id="sortby" class="select-input select-input--underbar">
              <option>新しい順</option>
              <option>古い順</option>
              <option>いいねが多い順</option>
              <option>評価が高い順</option>
              <option>評価が低い順</option>
            </select>
          </p>
          <p>
            <ons-input type="number" name="limit" id="limit" modifier="underbar" value="{{.SearchWords.Limit}}" placeholder="表示数(デフォルト:100)" float></ons-input>
          </p>
          <p style="margin-top: 30px;">
            <button class="button button--outline">search</button>
          </p>
        </div>
      </form>
    </div>
  </ons-page>
  <script>
    var target = document.getElementById('star');
    let text = "";
    for (let i = 10; i >= 0; i--) {
      text += "<option>"+i+"</option>"
    }
    target.innerHTML = text;
  </script>
  <script type="text/javascript">
    console.log({{.SearchWords}});
    if ({{.SearchWords}} != null){
      console.log({{.SearchWords}});
      SetMultipleSelection("FavoritePoint", {{.SearchWords.Category}});
      SetMultipleSelection("spoiler", {{.SearchWords.Spoiler}});
      SetMultipleSelection("star", {{.SearchWords.Star}});
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
    PullHook();
  });
</script>


<script>

  let comments = {{.Comment}};
  // var ratings = {{.RatingTvProgram}};
  if (comments == "") {
    comments = null;
    // ratings = null;
  }
  // console.log(ratings);
  const users = {{.Users}};
  let comment_likes;
  if ({{.CommentLike}} == null && comments != null){
    comment_likes = [comments.length];
    for (let i = comments.length - 1; i >= 0; i--) {
      comment_likes[i] = {Like:false};
    }
  } else {
    comment_likes = {{.CommentLike}};
  }
  ons.ready(function() {
    var infiniteList = document.getElementById('comments');
    if (comments != null) {
      console.log(comments);

      infiniteList.delegate = {
        createItemContent: function(i) {
        // console.log(ratings);
        const fps = comments[i].FavoritePoint.split('、');
        let fp_text = "";
        for (let j = fps.length - 1; j >= 0; j--) {
          fp_text += "<span style='padding:3px;color:blue;'>#"+fps[j]+"</span>";
        }
        if(comments[i].Spoiler){
          fp_text += "<i class='fas fa-hand-paper' style='color:palevioletred;'></i>";
        }
        return ons.createElement('<div class="comment"><ons-list-header style="background-color:antiquewhite;text-transform:none;"><div style="text-align:left; float:left;font-size:16px;">@' + users[i].Username + '</div><div style="text-align: right;margin-right:5px;">' + moment(comments[i].Created, "YYYY-MM-DDHH:mm:ss").format("YYYY/MM/DD HH:mm:ss") + '</div></ons-list-header><ons-list-item><ons-row><ons-col width="15%"><i class="fas fa-star" style="color:gold;"></i>：' + comments[i].Star +'</ons-col><ons-col style="font-size:12px;">'+ fp_text + '</ons-col></ons-row></ons-list-item><ons-list-item><div class="left"><a href="/tv/user/show/' + users[i].Id + '" title="user_page"><img class="list-item__thumbnail" src="' + users[i].IconUrl + '" alt="@' + users[i].Username + '"></a></div><div class="center"><span class="list-item__subtitle"id="comment_content_' + String(i) + '" style="font-size:14px;">' + comments[i].Content.replace(/(\r\n|\n|\r)/gm, "<br>") + '</span><span class="list-item__subtitle" style="text-align: right;"><div style="float:right;" id="count_like_' + i + '">：' + comments[i].CountLike + '</div><div style="float:right;"><i class="' + SetLikeBold(comment_likes[i].Like) + ' fa-thumbs-up" id="' + i + '" onclick="ClickLike(this)" style="color:' + SetLikeStatus(comment_likes[i].Like, 'orchid') + ';"></i></div></span></div></ons-list-item></div>');
      },
      countItems: function() {
        return comments.length;
      }
    };
    infiniteList.refresh();
  } else {
    infiniteList.innerHTML = "<div style='text-align:center;margin-top:40px;'><i class='far fa-surprise'>Not Found !!</i></div>"
  }
});
</script>

<script type="text/javascript">

  SetWatchBold("check_watched", {{.WatchStatus.Watched}});
  SetWatchBold("check_wtw", {{.WatchStatus.WantToWatch}});
  SetWatchStatus("check_watched", "lightcoral", {{.WatchStatus.Watched}});
  SetWatchStatus("check_wtw", "lightseagreen", {{.WatchStatus.WantToWatch}});

</script>


<script>
  global_comment_like_status = {{.CommentLike}};
</script>

<script type="text/javascript">
  function CommentLikeStatus(elem, check_flag) {
    let url = URL+"/tv/review_comment_like/";
    var data = global_comment_like_status[elem.id];
    let method;
    if (data.Id === 0){
      method = 'POST';
      data.UserId = {{.User.Id}};
      global_comment_like_status[elem.id].UserId = data.UserId;
      data.ReviewCommentId = {{.Comment}}[elem.id].Id;
      global_comment_like_status[elem.id].ReviewCommentId = data.ReviewCommentId;
    } else{
      method = 'PUT';
      url = url+data.Id;
    }
    data.Like = check_flag;
      // console.log("flag",global_comment_like_status[elem.id], check_flag);
      global_comment_like_status[elem.id].Like = data.Like;

      // console.log("last", global_comment_like_status[elem.id]);
      var json = JSON.stringify(data);
      var request = new XMLHttpRequest();
      request.open(method, url, true);
      request.setRequestHeader('Content-type','application/json; charset=utf-8');
      request.onload = function () {
        var x = JSON.parse(request.responseText);
        if (request.readyState == 4 && request.status == "200") {
          console.table(x);
        } else {
          global_comment_like_status[elem.id].Id = x.Id;
        }
      }
      request.send(json);
    };
  </script>

  <script type="text/javascript">
    global_watch_status = {{.WatchStatus}};
  </script>


  <script type="text/javascript">
    function WatchStatus(elem, check_flag) {
      let url = URL+"/tv/watching_status/";
      var data = global_watch_status;
      let method;
      if (data.Id === 0){
        method = 'POST';
        data.UserId = {{.User.Id}};
        global_watch_status.UserId = data.UserId;
        data.TvProgramId = {{.TvProgram.Id}};
        global_watch_status.TvProgramId = data.TvProgramId;
      } else{
        method = 'PUT';
        url = url+data.Id;
      }
      const str ="check_watched"
      if (elem.id.indexOf(str)===0) {
        data.Watched = check_flag;
        global_watch_status.Watched = data.Watched;
      } else {
        data.WantToWatch = check_flag;
        global_watch_status.WantToWatch = data.WantToWatch;

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
          global_watch_status.Id = x.Id;
        }
      }
      request.send(json);
    };
  </script>

  <!-- リロードしたらexpandable-listを閉じた状態にする -->
 <!--  <script>
    if (window.performance) {
      if (performance.navigation.type != 1) {
        document.querySelector('#expandable-list-item').showExpansion();
      }
    }
  </script> -->

  <!-- ツイートを保存する -->
  <script type="text/javascript">
    function PostComment() {
      const text_length = document.getElementById("tweet_dialog_content").value.length;
      if (text_length < 5){
        return DialogBox('alert_minlength');
      }
      let url = URL+"/tv/review_comment/";
      let data = {};
      data.Id  = 0;
      data.UserId = {{.User.Id}};
      data.TvProgramId  = {{.TvProgram.Id}};
      data.Content = document.getElementById("tweet_dialog_content").value;
      data.CountLike  = 0;
      data.Spoiler  = $("#check_spoiler").prop("checked");
      var fp = document.getElementById("FavoritePoint");
      let fps = [];
      for (let i = fp.length - 1; i >= 0; i--) {
        if (fp[i].selected){
          fps.push(fp[i].value);
        }
      }
      data.FavoritePoint = fps.join("、");
      data.Star = Number(document.getElementById("StarPoint").value);
      console.log(data);
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
      hideAlertDialog('tweet_dialog')
      setTimeout(window.location.reload(false), 500);
    };
  </script>

  <script>
    let time = String({{.TvProgram.Hour}});
    str = ".5";
    if (time == "100"){
      time = "";
    } else {
      if (time.indexOf(str) > -1){
        time = time.replace(str, ":30")
      } else {
        time += ":00";
      }
    }
    document.getElementById('tv_program_hour').innerHTML = time;
  </script>
  <script type="text/javascript">
    document.querySelector('ons-carousel').addEventListener('postchange', function() {
      if (carousel.getActiveIndex() == 1){
        GoTop();
      }
    });
  </script>
  <script type="text/javascript">
    function resetSelect() {
      document.search_comment.reset();
      document.getElementById("word").value = "";
      document.getElementById("limit").value = "";
    };
  </script>
</body>
</html>