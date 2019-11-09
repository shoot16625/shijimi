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
        commentLikes = [comments.length];
        for (let i = comments.length - 1; i >= 0; i--) {
          commentLikes[i] = {Like:false};
        }
      }
        ons.ready(function() {
          var infiniteList = document.getElementById('comments');
          if (comments != null) {
            infiniteList.delegate = {
              createItemContent: function(i) {
                const fps = comments[i].FavoritePoint.split(' ');
              let fpText = "";
              for (let j = fps.length - 1; j >= 0; j--) {
                fpText += "<span style='padding:3px;color:blue;'>#"+fps[j]+"</span>";
              }
              if(comments[i].Spoiler){
                fpText += "<i class='fas fa-hand-paper' style='color:palevioletred;'></i>";
              }
              return ons.createElement('<div class="comment-' + comments[i].Id + '"><ons-list-header style="background-color:antiquewhite;text-transform:none;"><div class="area-left profile-comment-list-header-font"><a href="/tv/tv_program/review/'+tvPrograms[i].Id+'" style="color:black;">' + tvPrograms[i].Title + '</a></div><div class="area-right list-margin">' + moment(comments[i].Created, "YYYY-MM-DDHH:mm:ss").format("YYYY/MM/DD HH:mm") + '</div></ons-list-header><ons-list-item><ons-row><ons-col width="15%"><i class="fas fa-star" style="color:gold;"></i>：' + comments[i].Star +'</ons-col><ons-col style="font-size:12px;">'+ fpText + '</ons-col></ons-row></ons-list-item><ons-list-item><div class="left"><a href="/tv/user/show/' + user.Id + '" title="user-page"><img class="list-item__thumbnail" src="' + user.IconUrl + '" alt="@' + user.Username + '"></a></div><div class="center"><span class="list-item__subtitle"id="comment-content-' + comments[i].Id + '" class="comment-list-content-font">' + comments[i].Content.replace(/(\r\n|\n|\r)/gm, "<br>") + '</span><span class="list-item__subtitle" class="area-right"><div style="float:right;" id="count-like-' + i + '">：' + comments[i].CountLike + '</div><div class="area-right"><i class="' + setLikeBold(commentLikes[i].Like) + ' fa-thumbs-up" id="' + i + '" onclick="clickLike(this)" style="color:' + setLikeStatus(commentLikes[i].Like, 'orchid') + ';"></i></div></span></div></ons-list-item></div>');
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

    <script type="text/javascript">
      function commentLikeStatus(elem, checkFlag) {
        let url = URL+"/tv/review_comment_like/";
        let data = globalCommentLikeStatus[elem.id];
        let method;
        if (data.Id === 0){
          method = 'POST';
          data.UserId = {{.MyUserId}};
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
            console.table(x);
          } else {
            globalCommentLikeStatus[elem.id].Id = x.Id;
          }
        }
        request.send(json);
      }
    </script>
    <script>
      reshapeBadges({{.User.Badge}});
    </script>
  </body>
</html>
