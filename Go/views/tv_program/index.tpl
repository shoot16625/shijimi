<!DOCTYPE html>
<html lang="ja">
  <head>
    {{ template "/common/header.tpl" . }}
  </head>

  <style type="text/css">
    select {
      width: 80%;
      max-width: 500px;
      height: 100px;
    }
  </style>

  <body>
    <ons-page>
      {{ template "/common/toolbar.tpl" . }}
      {{ template "/common/alert.tpl" . }}
      <ons-speed-dial position="bottom right" direction="up" ripple>
        <ons-fab>
          <ons-icon icon="md-share"></ons-icon>
        </ons-fab>
        <ons-speed-dial-item>
          <ons-icon
            icon="fa-search"
            onclick="DialogBoxEveryone('search_dialog')"
          ></ons-icon>
        </ons-speed-dial-item>
        <ons-speed-dial-item>
          <ons-icon
            icon="fa-chart-bar"
            onclick="GoAnotherCarousel(1)"
          ></ons-icon>
        </ons-speed-dial-item>
        <ons-speed-dial-item>
          <ons-icon
            icon="ion-arrow-up-a"
            onclick="GoTop()"
            style="vertical-align: 0px;"
          ></ons-icon>
        </ons-speed-dial-item>
      </ons-speed-dial>
      <ons-carousel
        swipeable
        overscrollable
        auto-scroll
        auto-refresh
        var="carousel"
      >
        <ons-carousel-item>
          <ons-list>
            <ons-lazy-repeat id="tv_programs"></ons-lazy-repeat>
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
    <template id="search_dialog.html">
      <ons-dialog id="search_dialog" modifier="large" cancelable fullscreen>
        <ons-page>
          <ons-toolbar>
            <div class="left">
              <ons-button
                id="cancel_button"
                onclick="hideAlertDialog('search_dialog')"
                style="background:left;color: grey;"
                ><i class="fas fa-window-close"></i
              ></ons-button>
            </div>
            <div class="center">
              <i class="fas fa-search" style="color: brown;"></i> 詳細検索
            </div>
            <div class="right">
              <ons-button
                id="reset_button"
                onclick="ResetSelect()"
                style="color:chocolate;background:left;"
                ><i class="far fa-trash-alt"></i
              ></ons-button>
            </div>
          </ons-toolbar>
          <div class="scroller">
            <form
              name="search_tv_program"
              id="search_tv_program"
              action="/tv/tv_program/search_tv_program"
              method="post"
            >
              <div style="text-align: center; margin-top: 30px;">
                <p>
                  <ons-input
                    type="text"
                    name="title"
                    id="title"
                    value="{{.SearchWords.Title}}"
                    modifier="underbar"
                    placeholder="タイトル"
                    float
                  ></ons-input>
                </p>
                <p>
                  <ons-input
                    type="text"
                    name="staff"
                    id="staff"
                    value="{{.SearchWords.Staff}}"
                    modifier="underbar"
                    placeholder="出演者・スタッフ・局"
                    float
                  ></ons-input>
                </p>
                <p>
                  <ons-input
                    type="text"
                    name="themesong"
                    id="themesong"
                    value="{{.SearchWords.Themesong}}"
                    modifier="underbar"
                    placeholder="アーティスト・曲名"
                    float
                  ></ons-input>
                </p>
                <p>
                  <label for="year" style="margin-right:8px;margin-left:8px;"
                    >＜放送年＞</label
                  >
                  <select
                    name="year"
                    id="year"
                    class="select-input select-input--underbar"
                    multiple
                  >
                  </select>
                </p>
                <p>
                  <label for="season">＜放送曜日＞</label>
                  <select
                    name="week"
                    id="week"
                    class="select-input select-input--underbar"
                    multiple
                  >
                    <option>月</option>
                    <option>火</option>
                    <option>水</option>
                    <option>木</option>
                    <option>金</option>
                    <option>土</option>
                    <option>日</option>
                  </select>
                </p>
                <p>
                  <label for="season" style="margin-right:8px;margin-left:8px;"
                    >＜時間帯＞</label
                  >
                  <select
                    name="hour"
                    id="hour"
                    class="select-input select-input--underbar"
                    multiple
                  >
                  </select>
                </p>
                <p>
                  <label for="season">＜シーズン＞</label>
                  <select
                    name="season"
                    id="season"
                    class="select-input select-input--underbar"
                    multiple
                  >
                    <option>春(4~6)</option>
                    <option>夏(7~9)</option>
                    <option>秋(10~12)</option>
                    <option>冬(1~3)</option>
                    <option>スペシャル</option>
                    <option>映画</option>
                  </select>
                </p>
                <p>
                  <label for="season">＜ジャンル＞</label>
                  <select
                    name="category"
                    id="category"
                    class="select-input select-input--underbar"
                    multiple
                  >
                    <option>アクション</option>
                    <option>アニメ映画</option>
                    <option>SF</option>
                    <option>学園・青春</option>
                    <option>グルメ</option>
                    <option>企業・オフィス</option>
                    <option>刑事・検事</option>
                    <option>コメディ</option>
                    <option>時代劇</option>
                    <option>スポーツ</option>
                    <option>政治</option>
                    <option>探偵・推理</option>
                    <option>ドキュメンタリー</option>
                    <option>犯罪・復讐</option>
                    <option>パロディ</option>
                    <option>不倫</option>
                    <option>弁護士</option>
                    <option>ホーム・ヒューマン</option>
                    <option>ホラー</option>
                    <option>ミステリー・サスペンス</option>
                    <option>恋愛</option>
                  </select>
                </p>
                <p>
                  <select
                    name="sortby"
                    id="sortby"
                    class="select-input select-input--underbar"
                  >
                    <option>新しい順</option>
                    <option>古い順</option>
                    <option>タイトル順</option>
                    <option>アーティスト順</option>
                    <option>閲覧数が多い順</option>
                    <option>見た人が多い順</option>
                    <option>見たい人が多い順</option>
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
                    float
                  ></ons-input>
                </p>
                <p style="margin-top: 30px;">
                  <button class="button button--outline">search</button>
                </p>
              </div>
            </form>
          </div>
        </ons-page>
        <script>
          var target = document.getElementById("hour");
          let text = "";
          let t;
          for (let i = 0; i < 48; i++) {
            if (i % 2 == 0) {
              t = String(i / 2) + ":00";
              text += "<option>" + t + "</option>";
            } else {
              t = String((i - 1) / 2) + ":30";
              text += "<option>" + t + "</option>";
            }
          }
          target.innerHTML = text;
        </script>
        <script>
          const today = new Date();
          const year = today.getFullYear() + 2;
          var target = document.getElementById("year");
          let text = "";
          for (let i = year; i >= 1970; i--) {
            text += "<option>" + i + "</option>";
          }
          target.innerHTML = text;
        </script>
        <script type="text/javascript">
          if ({{.SearchWords}} != null){
            console.log({{.SearchWords}});
            SetMultipleSelection("year", {{.SearchWords.Year}});
            SetMultipleSelection("week", {{.SearchWords.Week}});
            SetMultipleSelection("hour", {{.SearchWords.Hour}});
            SetMultipleSelection("season", {{.SearchWords.Season}});
            SetMultipleSelection("category", {{.SearchWords.Category}});
          }
          if ({{.SearchWords.Sortby}} != null){
            document.getElementById('sortby').value = {{.SearchWords.Sortby}};
          }
        </script>
        <script type="text/javascript">
          function ResetSelect() {
            document.search_tv_program.reset();
            document.getElementById("title").value = "";
            document.getElementById("staff").value = "";
            document.getElementById("themesong").value = "";
            document.getElementById("limit").value = "";
          }
        </script>
      </ons-dialog>
    </template>
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

    <script type="text/javascript">
      document
        .querySelector("ons-carousel")
        .addEventListener("postchange", function() {
          if (carousel.getActiveIndex() == 1) {
            GoTop();
          }
        });
    </script>
  </body>
</html>
