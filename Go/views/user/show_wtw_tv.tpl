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
      {{ template "/common/comment_review_change_only_user.tpl" . }}
      <ons-list class="list-margin">
        <ons-lazy-repeat id="tv-programs"></ons-lazy-repeat>
      </ons-list>
    </ons-page>
    <script type="text/javascript" src="/static/js/common.js"></script>
    <script>
      const tvPrograms = {{.TvProgram}};
      const watchStatus= {{.WatchStatus}};
      ons.ready(function() {
        var infiniteList = document.getElementById('tv-programs');
        if (tvPrograms != null) {
          infiniteList.delegate = {
            createItemContent: function(i) {
              let moviePosition;
              if (tvPrograms[i].MovieURL==""){
                moviePosition = '<img id="image-' + i + '" src="'+tvPrograms[i].ImageURL+'" alt="' + tvPrograms[i].Title + '" class="image">';
              } else {
                moviePosition = '<iframe id="movie-' + i + '" src="'+tvPrograms[i].MovieURL+'?modestbranding=1&rel=0&playsinline=1" frameborder="0" alt="' + tvPrograms[i].Title + '" allow="accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>';
              }
              let time = String(tvPrograms[i].Hour);
              str = ".5";
              if (time === "100"){
                time = "";
              } else {
                if (time.indexOf(str) > -1){
                  time = time.replace(str, ":30");
                } else {
                  time += ":00";
                }
              }
              let headerColor;
              if (tvPrograms[i].Season.Name === "春") {
                headerColor = "lavenderblush";
              } else if (tvPrograms[i].Season.Name === "夏") {
                headerColor = "aliceblue";
              } else if (tvPrograms[i].Season.Name === "秋") {
                headerColor = "khaki";
              } else if (tvPrograms[i].Season.Name === "冬") {
                headerColor = "thistle";
              } else {
                headerColor = "ghostwhite";
              }
              let categories = tvPrograms[i].Category.split('、');
              let category = "";
              for (let j = categories.length - 1; j >= 0; j--) {
                category += "<span style='padding:3px;'>#"+categories[j]+"</span>";
              }
              return ons.createElement('<div id="' + tvPrograms[i].Id + '"><ons-list-header style="background-color:'+ headerColor +';"><div class="area-left">' + tvPrograms[i].Year + '年' + tvPrograms[i].Season.Name + '（' + tvPrograms[i].Week.Name + time + '）</div><div class="area-right list-margin">閲覧数：' + tvPrograms[i].CountClicked + '</div></ons-list-header><ons-list-item><div class="tv-program-list-title-font">' + tvPrograms[i].Title + '</div></ons-list-item><ons-list-item><ons-row><ons-col><ons-row class="list-margin-bottom"><ons-col width="20%">出演：</ons-col><ons-col>' + tvPrograms[i].Cast+ '</ons-col></ons-row><ons-row class="list-margin-bottom"><ons-col width="20%">歌：</ons-col><ons-col>' + tvPrograms[i].Themesong+ '</ons-col></ons-row><ons-row class="list-margin-bottom"><ons-col width="20%">監督：</ons-col><ons-col>' + tvPrograms[i].Supervisor+ '</ons-col></ons-row><ons-row class="list-margin-bottom"><ons-col class="category-area">' + category + '</ons-col></ons-row><ons-row></ons-list-item><div class="area-center" style="margin:5px;">' + moviePosition + '</div><ons-list-item expandable>あらすじ・見どころ<div class="expandable-content">' + tvPrograms[i].Content + '</div></ons-list-item><ons-list-item modifier="nodivider"><i class="'+ setLikeBold(watchStatus[i].Watched) +' fa-laugh-beam" id="check-watched-' + i + '" onclick="clickWatchStatus(this)" style="color:' + setLikeStatus(watchStatus[i].Watched, 'deeppink') + ';"></i><div id="check-watched-' + i + '-text" class="tv-program-watch" style="margin-right: 8px;">見た：' + tvPrograms[i].CountWatched + '</div><i class="'+ setLikeBold(watchStatus[i].WantToWatch) +' fa-bookmark" id="check-wan2wat-' + i + '" onclick="clickWatchStatus(this)" style="color:' + setLikeStatus(watchStatus[i].WantToWatch, 'lightseagreen') + ';"></i><div id="check-wan2wat-' + i + '-text" class="tv-program-watch">また今度：' + tvPrograms[i].CountWantToWatch + '</div></ons-list-item><ons-list-item><div class="right list-item__right"><a href="/tv/tv_program/comment/' + tvPrograms[i].Id + '" style="text-decoration: none;">コメントを見る</a></div></ons-list-item></div>');
            },
            countItems: function() {
              return tvPrograms.length;
            }
          };
          infiniteList.refresh();
          } else {
            infiniteList.innerHTML = "<div style='text-align:center;margin-top:40px;'><i class='far fa-surprise' style='color:chocolate;'></i> Not Found !!</div>"
        }
      });
    </script>

    <script type="text/javascript">
      globalWatchStatus = {{.WatchStatus}};
    </script>

    <script type="text/javascript">
      function WatchStatus(elem, checkFlag) {
        let url = URL+"/tv/watching_status/";
        const index = elem.id.slice(14);
        let data = globalWatchStatus[index];
        let method;
        if (data.Id === 0){
          method = 'POST';
          data.UserId = {{.User.Id}};
          globalWatchStatus[index].UserId = data.UserId;
          data.TvProgramId = {{.TvProgram}}[index].Id;
          globalWatchStatus[index].TvProgramId = data.TvProgramId;
        } else {
          method = 'PUT';
          url = url+data.Id;
        }
        const str ="check-watched"
        if (elem.id.indexOf(str)===0) {
          data.Watched = checkFlag;
          globalWatchStatus[index].Watched = data.Watched;
        } else {
          data.WantToWatch = checkFlag;
          globalWatchStatus[index].WantToWatch = data.WantToWatch;

        }
        var json = JSON.stringify(data);
        var request = new XMLHttpRequest();
        request.open(method, url, true);
        request.setRequestHeader('Content-type','application/json; charset=utf-8');
        request.onload = function () {
          var x = JSON.parse(request.responseText);
          if (request.readyState == 4 && request.status == "200") {
            // console.table(x);
          } else {
            globalWatchStatus[index].Id = x.Id;
          }
        }
        request.send(json);
      };
    </script>
  </body>
</html>
