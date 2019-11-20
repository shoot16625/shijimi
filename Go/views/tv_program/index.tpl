<!DOCTYPE html>
<html lang="ja">
  <head>
    {{ template "/common/header.tpl" . }}
  </head>

  <body>
    <ons-page>
      {{ template "/common/toolbar.tpl" . }}
      {{ template "/common/alert.tpl" . }}
      <ons-pull-hook id="pull-hook"> </ons-pull-hook>
      <ons-speed-dial
        id="speed-dial"
        position="bottom right"
        direction="up"
        ripple
      >
        <ons-fab>
          <ons-icon icon="md-share"></ons-icon>
        </ons-fab>
      </ons-speed-dial>
      <ons-carousel
        swipeable
        overscrollable
        auto-scroll
        auto-refresh
        id="carousel"
      >
        <ons-carousel-item>
          <ons-list>
            <ons-lazy-repeat id="tv-programs"></ons-lazy-repeat>
          </ons-list>
        </ons-carousel-item>
        <ons-carousel-item>
          <ons-list>
            <ons-list-item id="browsing-ranking-24-expandable" expandable>
              閲覧ランキング（24時間以内）
              <div class="expandable-content">
                <ons-list modifier="inset" id="browsing-ranking-24"></ons-list>
              </div>
            </ons-list-item>
            <ons-list-item id="star-ranking-on-air-expandable" expandable>
              レビューランキング（放送中）
              <div class="expandable-content">
                <ons-list modifier="inset" id="star-ranking-on-air"></ons-list>
              </div>
            </ons-list-item>
          </ons-list>
        </ons-carousel-item>
      </ons-carousel>
    </ons-page>
    <template id="search-dialog.html">
      <ons-dialog id="search-dialog" modifier="large" cancelable fullscreen>
        <ons-page>
          <ons-toolbar>
            <div class="left">
              <ons-button
                id="cancel-button"
                onclick="hideAlertDialog('search-dialog')"
                style="background:left;color: grey;"
                ><i class="fas fa-window-close"></i
              ></ons-button>
            </div>
            <div class="center">
              <i class="fas fa-search" style="color: brown;"></i> 詳細検索
            </div>
            <div class="right">
              <ons-button
                id="reset-button"
                onclick="resetSelect()"
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
              <div class="area-center create-top-bottom-margin">
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
                    style="height: 130px;"
                    class="select-input select-input--underbar select-search-table"
                    multiple
                  >
                  </select>
                </p>
                <p>
                  <label for="week">＜放送曜日＞</label>
                  <select
                    name="week"
                    id="week"
                    style="height: 130px;"
                    class="select-input select-input--underbar select-search-table"
                    multiple
                  >
                    <option>月</option>
                    <option>火</option>
                    <option>水</option>
                    <option>木</option>
                    <option>金</option>
                    <option>土</option>
                    <option>日</option>
                    <option>平日</option>
                    <option>スペシャル</option>
                    <option>映画</option>
                    <option>映画以外</option>
                    <option>?</option>
                  </select>
                </p>
                <p>
                  <label for="hour" style="margin-right:8px;margin-left:8px;"
                    >＜時間帯＞</label
                  >
                  <select
                    name="hour"
                    id="hour"
                    style="height: 130px;"
                    class="select-input select-input--underbar select-search-table"
                    multiple
                  >
                  </select>
                </p>
                <p>
                  <label for="season">＜シーズン＞</label>
                  <select
                    name="season"
                    id="season"
                    style="height: 130px;"
                    class="select-input select-input--underbar select-search-table"
                    multiple
                  >
                    <option>春(4~6)</option>
                    <option>夏(7~9)</option>
                    <option>秋(10~12)</option>
                    <option>冬(1~3)</option>
                  </select>
                </p>
                <p>
                  <label for="category">＜ジャンル＞</label>
                  <select
                    name="category"
                    id="category"
                    style="height: 130px;"
                    class="select-input select-input--underbar select-search-table"
                    multiple
                  >
                    <option>コメディ・パロディ</option>
                    <option>恋愛</option>
                    <option>学園・青春</option>
                    <option>グルメ</option>
                    <option>ホーム・ヒューマン</option>
                    <option>企業・オフィス</option>
                    <option>刑事・検事</option>
                    <option>弁護士</option>
                    <option>医療</option>
                    <option>時代劇</option>
                    <option>スポーツ</option>
                    <option>政治</option>
                    <option>不倫</option>
                    <option>ミステリー・サスペンス</option>
                    <option>探偵・推理</option>
                    <option>犯罪・復讐</option>
                    <option>ホラー</option>
                    <option>ドキュメンタリー</option>
                    <option>アクション</option>
                    <option>アニメ映画</option>
                    <option>SF</option>
                  </select>
                </p>
                <p>
                  <select
                    name="sortby"
                    id="sortby"
                    class="select-input select-input--underbar select-search-table"
                  >
                    <option>新しい順</option>
                    <option>古い順</option>
                    <option>タイトル順</option>
                    <option>アーティスト順</option>
                    <option>閲覧数が多い順</option>
                    <option>評価が高い順</option>
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
                    min="1"
                    max="100"
                    float
                  ></ons-input>
                </p>
                <p class="create-top-bottom-margin">
                  <button class="button button--outline">search</button>
                </p>
              </div>
            </form>
          </div>
        </ons-page>
        <script>
          let textTop = '';
          document.getElementById('hour').innerHTML = getSelectHour(textTop);
        </script>
        <script>
          const today = new Date();
          const year = today.getFullYear() + 2;
          text = '';
          for (let i = year; i >= 1970; i--) {
            text += '<option>' + i + '</option>';
          }
          document.getElementById('year').innerHTML = text;
        </script>
        <script type="text/javascript">
          if ({{.SearchWords}} != null){
            setMultipleSelection("year", {{.SearchWords.Year}});
            setMultipleSelection("week", {{.SearchWords.Week}});
            setMultipleSelection("hour", {{.SearchWords.Hour}});
            setMultipleSelection("season", {{.SearchWords.Season}});
            setMultipleSelection("category", {{.SearchWords.Category}});
          }
          if ({{.SearchWords.Sortby}} != null){
            document.getElementById('sortby').value = {{.SearchWords.Sortby}};
          }
        </script>
        <script type="text/javascript">
          function resetSelect() {
            document.search_tv_program.reset();
            document.getElementById('title').value = '';
            document.getElementById('staff').value = '';
            document.getElementById('themesong').value = '';
            document.getElementById('limit').value = '';
          }
        </script>
      </ons-dialog>
    </template>
    {{ template "/common/js.tpl" . }}

    <script>
      let tvPrograms = {{.TvProgram}};
      if (tvPrograms != null && tvPrograms.length != 0) {
        console.log("表示数：", tvPrograms.length);
      } else {
        tvPrograms = null;
      }
      let watchStatus;
      if ({{.WatchStatus}} === null && tvPrograms != null){
        watchStatus = [tvPrograms.length];
        for (let i = tvPrograms.length - 1; i >= 0; i--) {
          watchStatus[i] = {Watched:false, WantToWatch:false};
        }
      } else {
        watchStatus = {{.WatchStatus}};
      }

      ons.ready(function() {
        var infiniteList = document.getElementById('tv-programs');
        if (tvPrograms != null) {
          infiniteList.delegate = {
            createItemContent: function(i) {
              let moviePosition = reshapeMovieCode(tvPrograms[i]);
              let time = reshapeHour(String(tvPrograms[i].Hour));

              let seasonName = avoidStructNameError(tvPrograms[i].Season);
              let weekName = avoidStructNameError(tvPrograms[i].Week);
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
              let category = reshapeCategory(categories);

              return ons.createElement('<div id="' + tvPrograms[i].Id + '"><ons-list-header style="background-color:'+ headerColor +';"><div class="area-left">' + tvPrograms[i].Year + '年 ' + seasonName + '（' + weekName + time + '）</div><div class="area-right list-margin">閲覧数：' + tvPrograms[i].CountClicked + '</div></ons-list-header><ons-list-item><div class="tv-program-list-title-font">' + tvPrograms[i].Title + '</div></ons-list-item><ons-list-item><ons-row><ons-col><ons-row class="list-margin-bottom"><ons-col width="15%">出演：</ons-col><ons-col>' + reshapeContents(casts) + '</ons-col></ons-row><ons-row class="list-margin-bottom"><ons-col width="15%">歌：</ons-col><ons-col>' + reshapeContents(themesongs)+ '</ons-col></ons-row><ons-row class="list-margin-bottom"><ons-col class="category-area">' + category + '</ons-col></ons-row><ons-row></ons-list-item><div class="area-center" style="margin:5px;">' + moviePosition + '<div class="reference">' + referenceSite + '</div></div><ons-list-item expandable>あらすじ・見どころ<div class="expandable-content">' + tvPrograms[i].Content + '</div></ons-list-item><ons-list-item modifier="nodivider"><i class="'+ setLikeBold(watchStatus[i].Watched) +' fa-laugh-beam" id="check-watched-' + i + '" onclick="clickWatchStatus(this)" style="color:' + setLikeStatus(watchStatus[i].Watched, 'deeppink') + ';"></i><div id="check-watched-' + i + '-text" class="tv-program-watch" style="margin-right: 8px;">見た：' + tvPrograms[i].CountWatched + '</div><i class="'+ setLikeBold(watchStatus[i].WantToWatch) +' fa-bookmark" id="check-wan2wat-' + i + '" onclick="clickWatchStatus(this)" style="color:' + setLikeStatus(watchStatus[i].WantToWatch, 'lightseagreen') + ';"></i><div id="check-wan2wat-' + i + '-text" class="tv-program-watch">また今度：' + tvPrograms[i].CountWantToWatch + '</div></ons-list-item><ons-list-item><div class="right list-item__right"><a href="/tv/tv_program/comment/' + tvPrograms[i].Id + '">コメントを見る</a></div></ons-list-item></div>');
            },
            countItems: function() {
              return tvPrograms.length;
            }
          };
          infiniteList.refresh();
          } else {
            infiniteList.innerHTML = "<div style='text-align:center;margin-top:40px;'><i class='far fa-surprise' style='color:chocolate;'></i> Not Found !!</div><div style='text-align:center;'>トップページから番組を登録しよう<i class='fas fa-male'></i></div>"
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
      let browsingCount = {{ .ViewTvProgramIn24 }};
      if (browsingCount === null) {
        browsingCount = ""
      }
      let text = "";
      let headerColor = ["lightpink","seashell","lavenderblush","antiquewhite","azure"];
      for (let i = 0; i < browsingCount.length; i++) {
        let time = reshapeHour(String(browsingCount[i].hour));
        text += '<ons-list-header style="background-color:'+ headerColor[i] +';"><div class="area-left">' + browsingCount[i].year + '年 ' + browsingCount[i].season_id + '（' + browsingCount[i].week_id + time + '）</div><div class="area-right list-margin">閲覧数：' + browsingCount[i].Num + '</div></ons-list-header><ons-list-item><div class="left">'+(i+1)+'</div><div class="center tv-program-list-content-font"><a href="/tv/tv_program/comment/' + browsingCount[i].id + '">' + browsingCount[i].title + '</a></div></ons-list-item>';
      }
      document.querySelector('#browsing-ranking-24').innerHTML = text;
      document.querySelector('#browsing-ranking-24-expandable').showExpansion();
    </script>
    <script>
      let starCount = {{ .goodStarTvProgramOnAir }};
      if (starCount === null) {
        starCount = "";
      }
      text = "";
      headerColor = ["lightpink","seashell","lavenderblush","antiquewhite","azure"];
      for (let i = 0; i < starCount.length; i++) {
        let seasonName = avoidStructNameError(starCount[i].Season);
        let weekName = avoidStructNameError(starCount[i].Week);
        let time = reshapeHour(String(starCount[i].Hour));
        text += '<ons-list-header style="background-color:'+ headerColor[i] +';"><div class="area-left">' + starCount[i].Year + '年 ' + seasonName + '（' + weekName + time + '）</div><div class="area-right list-margin">閲覧数：' + starCount[i].CountClicked + '</div></ons-list-header><ons-list-item><div class="left"><i class="fas fa-star"></i>'+starCount[i].Star+'</div><div class="center tv-program-list-content-font"><a href="/tv/tv_program/comment/' + starCount[i].Id + '">' + starCount[i].Title + '</a></div></ons-list-item>';
      }
      document.querySelector('#star-ranking-on-air').innerHTML = text;
      document.querySelector('#star-ranking-on-air-expandable').showExpansion();
    </script>

    <script type="text/javascript">
      document
        .querySelector('ons-carousel')
        .addEventListener('postchange', function() {
          if (carousel.getActiveIndex() == 1) {
            goTop();
          }
        });
    </script>
    <script>
      document.getElementById('speed-dial').innerHTML =
        "<ons-fab><ons-icon icon='md-share'></ons-icon></ons-fab><ons-speed-dial-item><ons-icon icon='md-search' onclick='dialogBoxEveryone(\"search-dialog\")'></ons-icon></ons-speed-dial-item><ons-speed-dial-item><ons-icon icon='md-chart' onclick='goAnotherCarousel(1)'></ons-icon></ons-speed-dial-item><ons-speed-dial-item><i class='fas fa-arrow-up' onclick='goTop()'></i></ons-speed-dial-item>";
    </script>
  </body>
</html>
