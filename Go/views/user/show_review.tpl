<!DOCTYPE html>
<html lang="ja">
  <head>
    {{ template "/common/header.tpl" . }}
  </head>

  <body>
    <ons-page>
      {{ template "/common/toolbar.tpl" . }}
      {{ template "/common/alert.tpl" . }}
      {{ template "/user/profile_onlyuser.tpl" . }}

      <ons-list style="margin-left: 3px;margin-right: 5px;">
        <ons-lazy-repeat id="comments"></ons-lazy-repeat>
      </ons-list>
    </ons-page>
    <script type="text/javascript" src="/static/js/common.js"></script>
    <script>

      var comments = {{.Comment}};
      var user = {{.User}};
      if ({{.CommentLike}} == null){
        var comment_likes = [comments.length];
        for (var i = comments.length - 1; i >= 0; i--) {
          comment_likes[i] = {Like:false};
        }
      } else {
        var comment_likes = {{.CommentLike}};
      }
      ons.ready(function() {
        var infiniteList = document.getElementById('comments');

        infiniteList.delegate = {
          createItemContent: function(i) {
            return ons.createElement('<div class="comment"><ons-list-header style="background-color:antiquewhite;text-transform:none;"><div style="text-align:left; float:left;font-size:16px;">@' + user.Username + '</div><div style="text-align: right;margin-right:5px;">' + moment(comments[i].Created, "YYYY-MM-DDHH:mm:ss").format("YYYY/MM/DD HH:mm:ss") + '</div></ons-list-header><ons-list-item><div class="left"><a href="/tv/user/show/' + user.Id + '" title="user_page"><img class="list-item__thumbnail" src="' + user.IconUrl + '" alt="@' + user.Username + '"></a></div><div class="center"><span class="list-item__subtitle"id="comment_content_' + String(i) + '" style="font-size:14px;">' + comments[i].Content.replace(/(\r\n|\n|\r)/gm, "<br>") + '</span><span class="list-item__subtitle" style="text-align: right;"><div style="float:right;" id="count_like_' + i + '">：' + comments[i].CountLike + '</div><div style="float:right;"><i class="' + SetLikeBold(comment_likes[i].Like) + ' fa-thumbs-up" id="' + i + '" onclick="ClickLike(this)" style="color:' + SetLikeStatus(comment_likes[i].Like, 'orchid') + ';"></i></div></span><span class="list-item__subtitle" style="text-align: right;"><form id="delete_comment" action="/tv/comment/' + comments[i].Id + '" method="post"><input type="hidden" name="_method" value="DELETE"><input type="hidden"><button type="submit">削除</button></form></span></div></ons-list-item></div>');
          },
          countItems: function() {
            return comments.length;
          }
        };
        infiniteList.refresh();
      });
    </script>

    <script>
      global_comment_like_status = {{.CommentLike}};
    </script>

    <script type="text/javascript">
      function CommentLikeStatus(elem, check_flag) {
        var url = URL+"/tv/review_comment_like/";
        var data = global_comment_like_status[elem.id];
        // console.log("here1;",data);
        if (data.Id === 0){
          var method = 'POST';
          data.UserId = {{.User.Id}};
          global_comment_like_status[elem.id].UserId = data.UserId;
          data.ReviewCommentId = {{.Comment}}[elem.id].Id;
          global_comment_like_status[elem.id].ReviewCommentId = data.ReviewCommentId;
        } else{
          var method = 'PUT';
          url = url+data.Id;
        }
        data.Like = check_flag;
        // console.log("flag",global_comment_like_status[elem.id], check_flag);
        // console.log("data",data);
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
  </body>
</html>
