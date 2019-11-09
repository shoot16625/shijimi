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
    {{ template "/common/js.tpl" . }}
    <script>
      const tvPrograms = {{.TvProgram}};
      const watchStatus= {{.WatchStatus}};
      ons.ready(function() {
        var infiniteList = document.getElementById('tv-programs');
        if (tvPrograms != null) {
          infiniteList.delegate = {
            createItemContent: function(i) {
              let moviePosition = reshapeMovieCode(tvPrograms[i]);
              let time = reshapeHour(String(tvPrograms[i].Hour));
              let seasonName = "";
              if (tvPrograms[i].Season != null) {
                seasonName = tvPrograms[i].Season.Name;
              }
              let weekName = "";
              if (tvPrograms[i].Week != null) {
                weekName = tvPrograms[i].Week.Name;
              }
              let headerColor = seasonHeaderColor(seasonName);
              let referenceSite = reshapeReferenceSite(tvPrograms[i]);
              let casts = tvPrograms[i].Cast;
              casts = casts.split(",").slice(0, 5);
              // let supervisors = tvPrograms[i].Supervisor;
              // supervisors = supervisors.split(",").slice(0, 3);
              let themesongs = tvPrograms[i].Themesong;
              themesongs = themesongs.split(",");
              let categories = tvPrograms[i].Category.split(',');
              if (tvPrograms[i].Category === ""){
                categories = [];
              }
              let category = "";
              for (let j = categories.length - 1; j >= 0; j--) {
                category += "<span style='padding:3px;'>#"+categories[j]+"</span>";
              }
              return ons.createElement('<div id="' + tvPrograms[i].Id + '"><ons-list-header style="background-color:'+ headerColor +';"><div class="area-left">' + tvPrograms[i].Year + '年 ' + seasonName + '（' + weekName + time + '）</div><div class="area-right list-margin">閲覧数：' + tvPrograms[i].CountClicked + '</div></ons-list-header><ons-list-item expandable>' + tvPrograms[i].Title + '<div class="expandable-content"><ons-list-item><ons-row><ons-col><ons-row class="list-margin-bottom"><ons-col width="20%">出演：</ons-col><ons-col>' + reshapeContents(casts) + '</ons-col></ons-row><ons-row class="list-margin-bottom"><ons-col width="20%">歌：</ons-col><ons-col>' + reshapeContents(themesongs) + '</ons-col></ons-row><ons-col class="category-area">' + category + '</ons-col></ons-row><ons-row></ons-list-item><div class="area-center" style="margin:5px;">' + moviePosition + '<div class="reference"><a href="'+tvPrograms[i].ImageUrl+'"target="_blank">'+referenceSite+'</a></div></div><ons-list-item expandable>あらすじ・見どころ<div class="expandable-content">' + tvPrograms[i].Content + '</div></ons-list-item><ons-list-item modifier="nodivider"><i class="'+ setLikeBold(watchStatus[i].Watched) +' fa-laugh-beam" id="check-watched-' + i + '" onclick="clickWatchStatus(this)" style="color:' + setLikeStatus(watchStatus[i].Watched, 'deeppink') + ';"></i><div id="check-watched-' + i + '-text" class="tv-program-watch" style="margin-right: 8px;">見た：' + tvPrograms[i].CountWatched + '</div><i class="'+ setLikeBold(watchStatus[i].WantToWatch) +' fa-bookmark" id="check-wan2wat-' + i + '" onclick="clickWatchStatus(this)" style="color:' + setLikeStatus(watchStatus[i].WantToWatch, 'lightseagreen') + ';"></i><div id="check-wan2wat-' + i + '-text" class="tv-program-watch">また今度：' + tvPrograms[i].CountWantToWatch + '</div></ons-list-item><ons-list-item><div class="right list-item__right"><a href="/tv/tv_program/comment/' + tvPrograms[i].Id + '" style="text-decoration: none;">コメントを見る</a></div></ons-list-item></div></ons-list-item></div>');
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
    <script>
      reshapeBadges({{.User.Badge}});
    </script>
  </body>
</html>
