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
              // let moviePosition = reshapeMovieCode(tvPrograms[i]);
              let time = reshapeHour(String(tvPrograms[i].Hour));
              let seasonName = avoidStructNameError(tvPrograms[i].Season);
              let weekName = avoidStructNameError(tvPrograms[i].Week);
              let headerColor = seasonHeaderColor(seasonName);
              // let referenceSite = reshapeReferenceSite(tvPrograms[i]);
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
              let category = reshapeCategory(categories);
              return ons.createElement('<div id="' + tvPrograms[i].Id + '"><ons-list-header style="background-color:'+ headerColor +';height:10px"><div class="area-left">' + tvPrograms[i].Year + '年 ' + seasonName + '（' + weekName + time + '）</div><div class="area-right list-margin"><i class="fas fa-eye"></i>：' + tvPrograms[i].CountClicked + '</div></ons-list-header><ons-list-item expandable>' + tvPrograms[i].Title + '<div class="expandable-content"><ons-list-item><ons-row><ons-col><ons-row class="list-margin-bottom"><ons-col width="20%">出演：</ons-col><ons-col>' + reshapeContents(casts) + '</ons-col></ons-row><ons-row class="list-margin-bottom"><ons-col width="20%"><i class="fas fa-music" style="color: cornflowerblue;"></i>：</ons-col><ons-col>' + reshapeContents(themesongs) + '</ons-col></ons-row><ons-col class="category-area">' + category + '</ons-col></ons-row><ons-row></ons-list-item><ons-list-item expandable>あらすじ・見どころ<div class="expandable-content">' + tvPrograms[i].Content + '</div></ons-list-item><ons-list-item modifier="nodivider"><i class="'+ setLikeBold(watchStatus[i].Watched) +' fa-laugh-beam" id="check-watched-' + i + '" onclick="clickWatchStatus(this, {{.User.Id}}, tvPrograms)" style="color:' + setLikeStatus(watchStatus[i].Watched, 'deeppink') + ';"></i><div id="check-watched-' + i + '-text" class="tv-program-watch" style="margin-right: 8px;">見た：' + tvPrograms[i].CountWatched + '</div><i class="'+ setLikeBold(watchStatus[i].WantToWatch) +' fa-bookmark" id="check-wan2wat-' + i + '" onclick="clickWatchStatus(this, {{.User.Id}}, tvPrograms)" style="color:' + setLikeStatus(watchStatus[i].WantToWatch, 'lightseagreen') + ';"></i><div id="check-wan2wat-' + i + '-text" class="tv-program-watch">また今度：' + tvPrograms[i].CountWantToWatch + '</div></ons-list-item><ons-list-item><div class="right list-item__right"><a href="/tv/tv_program/comment/' + tvPrograms[i].Id + '" style="text-decoration: none;">コメントを見る</a></div></ons-list-item></div></ons-list-item></div>');
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
    <script>
      reshapeBadges({{.User.Badge}});
    </script>
  </body>
</html>
