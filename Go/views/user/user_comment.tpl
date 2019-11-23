<!DOCTYPE html>
<html lang="ja">
  <head>
    {{ template "/common/header.tpl" . }}
  </head>
  <body>
    <ons-page>
      {{ template "/common/toolbar.tpl" . }}
      {{ template "/common/alert.tpl" . }}
      {{ template "/user/profile_everyone.tpl" . }}
      {{ template "/common/comment_review_change_everyone.tpl" . }}

      <ons-list class="list-margin">
        <ons-lazy-repeat id="comments"></ons-lazy-repeat>
      </ons-list>
    </ons-page>
    {{ template "/common/js.tpl" . }}
    <script>
      let comments = {{.Comment}};
      if (comments.length === 0) {
        comments = null;
      }
      const user = {{.User}};
      const tvPrograms = {{.TvProgram}};
      let commentLikes = {{.CommentLike}};
      if ({{.CommentLike}} === null && comments != null){
        // ログインしていない場合
        commentLikes = [comments.length];
        for (let i = comments.length - 1; i >= 0; i--) {
          commentLikes[i] = {Like:false};
        }
      }
      let myUserID = {{.MyUserId}};
      if (myUserID === null) {
        myUserID = "";
      }
      ons.ready(function() {
        var infiniteList = document.getElementById('comments');
        if (comments != null) {
          infiniteList.delegate = {
            createItemContent: function(i) {

              return ons.createElement('<div id="' + comments[i].Id + '"><ons-list-header style="background-color:aliceblue;text-transform:none;"><div class="area-left profile-comment-list-header-font"><a href="/tv/tv_program/comment/'+tvPrograms[i].Id+'" style="color:black;">' + tvPrograms[i].Title + '</a></div><div class="area-right list-margin">' + moment(comments[i].Created, "YYYY-MM-DDHH:mm:ss").format("YYYY/MM/DD HH:mm") + '</div></ons-list-header><ons-list-item><div class="left"><a href="/tv/user/show/' + user.Id + '" title="user_comment"><img class="list-item__thumbnail" src="' + user.IconUrl + '" onerror="this.src=\'/static/img/user_img/s256_f_01.png\'"></a></div><div class="center"><span class="list-item__subtitle"id="comment-content-' + comments[i].Id + '" class="comment-list-content-font">' + comments[i].Content.replace(/(\r\n|\n|\r)/gm, "<br>") + '</span><span class="list-item__subtitle" class="area-right"><div style="float:right;" id="count-like-' + i + '">：' + comments[i].CountLike + '</div><div class="area-right"><i class="' + setLikeBold(commentLikes[i].Like) + ' fa-thumbs-up" id="' + i + '" onclick="clickLike(this,myUserID,comments,\'comment\')" style="color:' + setLikeStatus(commentLikes[i].Like, 'orchid') + ';"></i></div></span></div></ons-list-item></div>');
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

    <script>
      globalCommentLikeStatus = {{.CommentLike}};
    </script>

    <script>
      reshapeBadges({{.User.Badge}});
    </script>
  </body>
</html>
