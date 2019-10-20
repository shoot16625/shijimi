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
        <ons-lazy-repeat id="tv_programs"></ons-lazy-repeat>
      </ons-list>
    </ons-page>
    <script type="text/javascript" src="/static/js/common.js"></script>
    <script>
      let tv_programs = {{.TvProgram}};
      if (tv_programs == "") {
        tv_programs = null;
      }
      let watch_status;
      if ({{.WatchStatus}} == null && tv_programs != null){
        watch_status = [tv_programs.length];
        for (let i = tv_programs.length - 1; i >= 0; i--) {
          watch_status[i] = {Watched:false, WantToWatch:false};
        }
      } else {
        watch_status = {{.WatchStatus}};
      }

      ons.ready(function() {
        var infiniteList = document.getElementById('tv_programs');
        if (tv_programs != null) {
          infiniteList.delegate = {
            createItemContent: function(i) {
              let movie_position;
              if (tv_programs[i].MovieUrl==""){
                movie_position = '<img id="image_' + i + '" src="'+tv_programs[i].ImageUrl+'" alt="' + tv_programs[i].Title + '" width="80%">';
              } else {
                movie_position = '<iframe id="movie_' + i + '" src="'+tv_programs[i].MovieUrl+'?modestbranding=1&rel=0&playsinline=1" frameborder="0" alt="' + tv_programs[i].Title + '" allow="accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>';
              }
              let time = String(tv_programs[i].Hour);
              str = ".5";
              if (time == "100"){
                time = "";
              } else {
                if (time.indexOf(str) > -1){
                  time = time.replace(str, ":30");
                } else {
                  time += ":00";
                }
              }
              let header_color;
              if (tv_programs[i].Season.Name == "春") {
                header_color = "lavenderblush";
              } else if (tv_programs[i].Season.Name == "夏") {
                header_color = "aliceblue";
              } else if (tv_programs[i].Season.Name == "秋") {
                header_color = "khaki";
              } else if (tv_programs[i].Season.Name == "冬") {
                header_color = "thistle";
              } else {
                header_color = "ghostwhite";
              }
              return ons.createElement('<div class="tv_program"><ons-list-header style="background-color:'+ header_color +';"><div style="text-align: left; float:left;">' + tv_programs[i].Year + '年' + tv_programs[i].Season.Name + '（' + tv_programs[i].Week.Name + time + '）</div><div style="text-align: right; margin-right:5px;">閲覧数：' + tv_programs[i].CountClicked + '</div></ons-list-header><ons-list-item><div class="title" style="font-size: 20px;">' + tv_programs[i].Title + '</div></ons-list-item><ons-list-item><ons-row><ons-col><div class="content"><ons-row style="margin-bottom:5px;"><ons-col width="20%">出演：</ons-col><ons-col>' + tv_programs[i].Cast+ '</ons-col></ons-row><ons-row style="margin-bottom:5px;"><ons-col width="20%">歌：</ons-col><ons-col>' + tv_programs[i].Themesong+ '</ons-col></ons-row><ons-row style="margin-bottom:5px;"><ons-col width="20%">監督：</ons-col><ons-col>' + tv_programs[i].Supervisor+ '</ons-col></ons-row><ons-row style="margin-bottom:5px;"><ons-col width="20%">脚本：</ons-col><ons-col>' + tv_programs[i].Dramatist+ '</ons-col></ons-row><ons-row style="margin-bottom:5px;"><ons-col width="20%">演出：</ons-col><ons-col>' + tv_programs[i].Director+ '</ons-col></ons-row></div></ons-col><ons-row></ons-list-item><div style="text-align:center; margin:5px;">' + movie_position + '</div><ons-list-item expandable>あらすじ・見どころ<div class="expandable-content">' + tv_programs[i].Content + '</div></ons-list-item><ons-list-item modifier="nodivider"><i class="'+ SetLikeBold(watch_status[i].Watched) +' fa-laugh-beam" id="check_watched_' + i + '" onclick="ClickWatchStatus(this)" style="color:' + SetLikeStatus(watch_status[i].Watched, 'lightcoral') + ';"></i><div id="check_watched_' + i + '_text" style="float:right; margin-left: 5px;margin-right: 8px;">見た：' + tv_programs[i].CountWatched + '</div><i class="'+ SetLikeBold(watch_status[i].WantToWatch) +' fa-bookmark" id="check_wan2wat_' + i + '" onclick="ClickWatchStatus(this)" style="color:' + SetLikeStatus(watch_status[i].WantToWatch, 'lightseagreen') + ';"></i><div id="check_wan2wat_' + i + '_text" style="float:right; margin-left: 5px;">また今度：' + tv_programs[i].CountWantToWatch + '</div></ons-list-item><ons-list-item><div class="right list-item__right"><a href="/tv/tv_program/comment/' + tv_programs[i].Id + '" style="text-decoration: none;">コメントを見る</a></div></ons-list-item></div>');
            },
            countItems: function() {
              return tv_programs.length;
            }
          };
          infiniteList.refresh();
          } else {
            infiniteList.innerHTML = "<div style='text-align:center;margin-top:40px;'><i class='far fa-surprise'>Not Found !!</i></div><div style='text-align:center;'>トップページから番組を登録してね<i class='fas fa-male'></i></div>"
          }
      });
    </script>

    <script type="text/javascript">
      global_watch_status = {{.WatchStatus}};
    </script>

    <script type="text/javascript">
      function WatchStatus(elem, check_flag) {
        let url = URL+"/tv/watching_status/";
        const index = elem.id.slice(14);
        let data = global_watch_status[index];
        let method;
        if (data.Id === 0){
          method = 'POST';
          data.UserId = {{.User.Id}};
          global_watch_status[index].UserId = data.UserId;
          data.TvProgramId = {{.TvProgram}}[index].Id;
          global_watch_status[index].TvProgramId = data.TvProgramId;
        } else {
          method = 'PUT';
          url = url+data.Id;
        }
        const str ="check_watched"
        if (elem.id.indexOf(str)===0) {
          data.Watched = check_flag;
          global_watch_status[index].Watched = data.Watched;
        } else {
          data.WantToWatch = check_flag;
          global_watch_status[index].WantToWatch = data.WantToWatch;

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
            global_watch_status[index].Id = x.Id;
          }
        }
        request.send(json);
      };
    </script>
  </body>
</html>
