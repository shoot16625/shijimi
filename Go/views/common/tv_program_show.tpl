<ons-list>
  <ons-list-header style="background-color:ghostwhite;">
    <div class="area-left">
      <span id="tv-program-week">
        {{.TvProgram.Year}}年 {{.TvProgram.Season.Name}}（{{.TvProgram.Week.Name}}
      </span>
      <span id="tv-program-hour"></span>）
    </div>
    <div class="area-right list-margin">
      閲覧数：{{.TvProgram.CountClicked}}
    </div>
  </ons-list-header>
  <ons-list-item id="expandable-list-item" expandable>
    {{.TvProgram.Title}}
    <div class="expandable-content">
      <ons-row>
        <ons-col>
          <div class="content">
            <ons-row class="list-margin-bottom">
              <ons-col width="20%">出演：</ons-col>
              <ons-col>{{.TvProgram.Cast}}</ons-col>
            </ons-row>
            <ons-row class="list-margin-bottom">
              <ons-col width="20%">歌：</ons-col>
              <ons-col>{{.TvProgram.Themesong}}</ons-col>
            </ons-row>
            <ons-row class="list-margin-bottom">
              <ons-col width="20%">監督：</ons-col>
              <ons-col>{{.TvProgram.Supervisor}}</ons-col>
            </ons-row>
            <ons-row class="list-margin-bottom">
              <ons-col width="20%">脚本：</ons-col>
              <ons-col>{{.TvProgram.Dramatist}}</ons-col>
            </ons-row>
            <ons-row class="list-margin-bottom">
              <ons-col width="20%">演出：</ons-col>
              <ons-col>{{.TvProgram.Director}}</ons-col>
            </ons-row>
            <ons-row class="list-margin-bottom">
              <ons-col width="20%">制作：</ons-col>
              <ons-col>{{.TvProgram.Production}}</ons-col>
            </ons-row>
            <ons-row class="list-margin-bottom">
              <ons-col
                class="category-area"
                id="category-area"
              >
              </ons-col>
            </ons-row>
          </div>
        </ons-col>
      </ons-row>
      <div class="area-center">
        <img
          src="{{.TvProgram.ImageURL}}"
          alt="{{.Title}}"
          class="image"
          onerror="this.src='http:\/\/hankodeasobu.com/wp-content/uploads/animals_02.png'"
        />
      </div>
      <ons-list-item expandable>
        あらすじ・見どころ
        <div class="right"></div>
        <div class="expandable-content">{{.TvProgram.Content}}</div>
      </ons-list-item>
      <ons-list-item>
        <div class="left">
          <i
            class="fa-laugh-beam"
            id="check-watched"
            onclick="clickWatchStatus(this)"
          ></i>
          <div
            id="check-watched-text"
            style="float:right; margin-left: 5px;margin-right: 8px;"
          >
            見た：{{.TvProgram.CountWatched}}
          </div>
          <i
            class="fa-bookmark"
            id="check-wtw"
            onclick="clickWatchStatus(this)"
          ></i>
          <div id="check-wtw-text" style="float:right; margin-left: 5px;">
            また今度：{{.TvProgram.CountWantToWatch}}
          </div>
        </div>
        <div class="right">
          <div id="edit-tv-program">
            <ons-button
              modifier="quiet"
              onclick="goOtherPage({{.User}},'/tv/tv_program/edit/{{.TvProgram.Id}}')"
              >編集</ons-button
            >
          </div>
        </div>
      </ons-list-item>
    </div>
  </ons-list-item>
</ons-list>