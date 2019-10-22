<!DOCTYPE html>
<html lang="ja">
  <head>
    {{ template "/common/header.tpl" . }}
  </head>

  <body>
    <ons-page>
      {{ template "/common/toolbar.tpl" . }}
      {{ template "/common/alert.tpl" . }}
      {{ template "/user/profile_only_user.tpl" . }}
      {{ template "/common/comment_review_change_profile.tpl" . }}
      <ons-list class="list-margin">
        <ons-lazy-repeat id="comments"></ons-lazy-repeat>
      </ons-list>
    </ons-page>
    <script type="text/javascript" src="/static/js/common.js"></script>
    <script>

      let comments = {{.Comment}};
      let user = {{.User}};
      let tvPrograms = {{.TvProgram}};
      let commentLikes;
      if ({{.CommentLike}} == null){
        commentLikes = [comments.length];
        for (let i = comments.length - 1; i >= 0; i--) {
          commentLikes[i] = {Like:false};
        }
      } else {
        commentLikes = {{.CommentLike}};
      }
      ons.ready(function() {
        var infiniteList = document.getElementById('comments');

        infiniteList.delegate = {
          createItemContent: function(i) {

            return ons.createElement('<div id="' + comments[i].Id + '"><ons-list-header style="background-color:aliceblue;text-transform:none;"><div class="area-left comment-list-header-font">' + tvPrograms[i].Title + '</div><div class="area-right list-margin">' + moment(comments[i].Created, "YYYY-MM-DDHH:mm:ss").format("YYYY/MM/DD HH:mm:ss") + '</div></ons-list-header><ons-list-item><div class="left"><a href="/tv/user/show/' + user.Id + '" title="user_page"><img class="list-item__thumbnail" src="' + user.IconURL + '" alt="@' + user.Username + '"></a></div><div class="center"><span class="list-item__subtitle"id="comment-content-' + comments[i].Id + '" class="comment-list-content-font">' + comments[i].Content.replace(/(\r\n|\n|\r)/gm, "<br>") + '</span><span class="list-item__subtitle" class="area-right"><div style="float:right;" id="count-like-' + i + '">：' + comments[i].CountLike + '</div><div class="area-right"><i class="' + setLikeBold(commentLikes[i].Like) + ' fa-thumbs-up" id="' + i + '" onclick="clickLike(this)" style="color:' + setLikeStatus(commentLikes[i].Like, 'orchid') + ';"></i></div></span><span class="list-item__subtitle area-right"><div style="float:right;"><form id="delete-comment-' + comments[i].Id + '" action="/tv/comment/' + comments[i].Id + '" method="post"><input type="hidden" name="_method" value="DELETE"><input type="hidden"><button class="button button--light" style="line-height: 8px; font-size:12px;" type="submit">del</button></form></div></span></div></ons-list-item></div>');
          },

          countItems: function() {
            return comments.length;
          }
        };
        infiniteList.refresh();
      });
    </script>

    <script>
      globalCommentLikeStatus = {{.CommentLike}};
    </script>

    <script type="text/javascript">
      function commentLikeStatus(elem, checkFlag) {
        let url = URL+"/tv/comment_like/";
        let data = globalCommentLikeStatus[elem.id];
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
        // console.log("flag",globalCommentLikeStatus[elem.id], checkFlag);
        globalCommentLikeStatus[elem.id].Like = data.Like;

        // console.log("last", globalCommentLikeStatus[elem.id]);
        var json = JSON.stringify(data);
        var request = new XMLHttpRequest();
        request.open(method, url, true);
        request.setRequestHeader('Content-type','application/json; charset=utf-8');
        request.onload = function () {
          var x = JSON.parse(request.responseText);
          if (request.readyState == 4 && request.status == "200") {
            console.table(x);
          } else {
            globalCommentLikeStatus[elem.id].Id = x.Id;
          }
        }
        request.send(json);
      }
    </script>
  </body>
</html>
