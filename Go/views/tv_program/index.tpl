<!DOCTYPE html>
<html lang="ja">
  <head>
    {{ template "/common/header.tpl" . }}
  </head>

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
            icon="md-search"
            onclick="dialogBoxEveryone('search-dialog')"
          ></ons-icon>
        </ons-speed-dial-item>
        <ons-speed-dial-item>
          <ons-icon icon="md-chart" onclick="goAnotherCarousel(1)"></ons-icon>
        </ons-speed-dial-item>
        <ons-speed-dial-item>
          <ons-icon icon="md-home" onclick="goTop()"></ons-icon>
        </ons-speed-dial-item>
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
          <div class="area-center">
            <p>詳細情報や分析結果を表示</p>
            <p>工事中<i class="fas fa-truck-pickup"></i></p>
          </div>
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
              <div class="area-center create-top-margin">
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
                    class="select-input select-input--underbar select-search-table"
                    multiple
                  >
                  </select>
                </p>
                <p>
                  <label for="season">＜放送曜日＞</label>
                  <select
                    name="week"
                    id="week"
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
                    <option>?</option>
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
                  <label for="season">＜ジャンル＞</label>
                  <select
                    name="category"
                    id="category"
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
                <p class="create-top-margin">
                  <button class="button button--outline">search</button>
                </p>
              </div>
            </form>
          </div>
        </ons-page>
        <script>
          var target = document.getElementById('hour');
          let text = '';
          let t;
          for (let i = 0; i <= 48; i++) {
            if (i % 2 == 0) {
              t = String(i / 2) + ':00';
              text += '<option>' + t + '</option>';
            } else {
              t = String((i - 1) / 2) + ':30';
              text += '<option>' + t + '</option>';
            }
          }
          target.innerHTML = text;
        </script>
        <script>
          const today = new Date();
          const year = today.getFullYear() + 2;
          var target = document.getElementById('year');
          text = '';
          for (let i = year; i >= 1970; i--) {
            text += '<option>' + i + '</option>';
          }
          target.innerHTML = text;
        </script>
        <script type="text/javascript">
          if ({{.SearchWords}} != null){
            // console.log({{.SearchWords}});
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
    <script type="text/javascript" src="/static/js/common.js"></script>

    <script>
      let tvPrograms = {{.TvProgram}};
      if (tvPrograms.length === 0) {
        tvPrograms = null;
      }
      console.log("表示数：", tvPrograms.length);
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
              let moviePosition;
              if (tvPrograms[i].MovieURL===""){
                moviePosition = '<img class="image" id="image-' + tvPrograms[i].Id + '" src="'+tvPrograms[i].ImageURL+'" alt="' + tvPrograms[i].Title + '" width="80%" onerror="this.src=\'http://hankodeasobu.com/wp-content/uploads/animals_02.png\'">';
              } else {
                moviePosition = '<iframe id="movie-' + tvPrograms[i].Id + '" class="movie" src="'+tvPrograms[i].MovieURL+'?modestbranding=1&rel=0&playsinline=1" frameborder="0" alt="' + tvPrograms[i].Title + '" allow="accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>';
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
              let casts = tvPrograms[i].Cast;
              casts = casts.split("、").slice(0, 5).join("、");
              let referenceSite = tvPrograms[i].ImageURL;
              if (referenceSite.includes("walkerplus")){
                referenceSite = "@MovieWalker"
              } else {
                referenceSite = ""
              }
              return ons.createElement('<div id="' + tvPrograms[i].Id + '"><ons-list-header style="background-color:'+ headerColor +';"><div class="area-left">' + tvPrograms[i].Year + '年 ' + tvPrograms[i].Season.Name + '（' + tvPrograms[i].Week.Name + time + '）</div><div class="area-right list-margin">閲覧数：' + tvPrograms[i].CountClicked + '</div></ons-list-header><ons-list-item><div class="tv-program-list-title-font">' + tvPrograms[i].Title + '</div></ons-list-item><ons-list-item><ons-row><ons-col><ons-row class="list-margin-bottom"><ons-col width="20%">出演：</ons-col><ons-col>' + casts + '</ons-col></ons-row><ons-row class="list-margin-bottom"><ons-col width="20%">歌：</ons-col><ons-col>' + tvPrograms[i].Themesong+ '</ons-col></ons-row><ons-row class="list-margin-bottom"><ons-col width="20%">監督：</ons-col><ons-col>' + tvPrograms[i].Supervisor+ '</ons-col></ons-row><ons-row class="list-margin-bottom"><ons-col width="20%">脚本：</ons-col><ons-col>' + tvPrograms[i].Dramatist+ '</ons-col></ons-row><ons-row class="list-margin-bottom"><ons-col width="20%">演出：</ons-col><ons-col>' + tvPrograms[i].Director+ '</ons-col></ons-row></ons-col><ons-row></ons-list-item><div class="area-center" style="margin:5px;">' + moviePosition + '<div class="reference">'+referenceSite+'</div></div><ons-list-item expandable>あらすじ・見どころ<div class="expandable-content">' + tvPrograms[i].Content + '</div></ons-list-item><ons-list-item modifier="nodivider"><i class="'+ setLikeBold(watchStatus[i].Watched) +' fa-laugh-beam" id="check-watched-' + i + '" onclick="clickWatchStatus(this)" style="color:' + setLikeStatus(watchStatus[i].Watched, 'deeppink') + ';"></i><div id="check-watched-' + i + '-text" class="tv-program-watch" style="margin-right: 8px;">見た：' + tvPrograms[i].CountWatched + '</div><i class="'+ setLikeBold(watchStatus[i].WantToWatch) +' fa-bookmark" id="check-wan2wat-' + i + '" onclick="clickWatchStatus(this)" style="color:' + setLikeStatus(watchStatus[i].WantToWatch, 'lightseagreen') + ';"></i><div id="check-wan2wat-' + i + '-text" class="tv-program-watch">また今度：' + tvPrograms[i].CountWantToWatch + '</div></ons-list-item><ons-list-item><div class="right list-item__right"><a href="/tv/tv_program/comment/' + tvPrograms[i].Id + '" style="text-decoration: none;">コメントを見る</a></div></ons-list-item></div>');
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
      $('img').on('error', function() {
        $(this).attr(
          'src',
          'http://hankodeasobu.com/wp-content/uploads/animals_02.png'
        );
      });
    </script>
  </body>
</html>
