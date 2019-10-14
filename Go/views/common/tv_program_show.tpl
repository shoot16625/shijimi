
<ons-list>
  <ons-list-header style="background-color:ghostwhite;">
    <div style="text-align: left; float:left;">
      {{.TvProgram.Year}}年 {{.TvProgram.Season.Name}}（{{.TvProgram.Week.Name}}<span id="tv_program_hour"></span>）
    </div>
    <div style="text-align: right;margin-right:5px;">
      閲覧数：{{.TvProgram.CountClicked}}
    </div>
  </ons-list-header>
  <ons-list-item id="expandable-list-item" expandable>
    {{.TvProgram.Title}}
    <div class="expandable-content">
      <ons-row>
        <ons-col>
          <div class="content">
            <ons-row style="margin-bottom:5px;">
              <ons-col width="20%">出演：</ons-col>
              <ons-col>{{.TvProgram.Cast}}</ons-col>
            </ons-row>
            <ons-row style="margin-bottom:5px;">
              <ons-col width="20%">歌：</ons-col>
              <ons-col>{{.TvProgram.Themesong}}</ons-col>
            </ons-row>
            <ons-row style="margin-bottom:5px;">
              <ons-col width="20%">監督：</ons-col>
              <ons-col>{{.TvProgram.Supervisor}}</ons-col>
            </ons-row>
            <ons-row style="margin-bottom:5px;">
              <ons-col width="20%">脚本：</ons-col>
              <ons-col>{{.TvProgram.Dramatist}}</ons-col>
            </ons-row>
            <ons-row style="margin-bottom:5px;">
              <ons-col width="20%">演出：</ons-col>
              <ons-col>{{.TvProgram.Director}}</ons-col>
            </ons-row>
          </div>
        </ons-col>
      </ons-row>
      <div class="image" align="center">
        <img src="{{.TvProgram.ImageUrl}}" alt="{{.Title}}" width="80%">
      </div>
      <ons-list-item expandable>
        あらすじ・見どころ
        <div class="right">
        </div>
        <div class="expandable-content">{{.TvProgram.Content}}</div>
      </ons-list-item>
      <ons-list-item>
        <div class="left">
          <i class="fa-laugh-beam" id="check_watched" onclick="ClickWatchStatus(this)"></i>
          <div id="check_watched_text" style="float:right; margin-left: 5px;margin-right: 8px;">見た：{{.TvProgram.CountWatched}}</div>
          <i class="fa-bookmark" id="check_wtw" onclick="ClickWatchStatus(this)"></i>
          <div id="check_wtw_text" style="float:right; margin-left: 5px;">また今度：{{.TvProgram.CountWantToWatch}}</div>
        </div>
        <div class="right">
          <div id="edit_tv_program">
          <ons-button modifier="quiet" onclick="GoOtherPage({{.User}},'/tv/tv_program/edit/{{.TvProgram.Id}}')">編集</ons-button>
          </div>
        </div>
      </ons-list-item>
    </div>
  </ons-list-item>
</ons-list>