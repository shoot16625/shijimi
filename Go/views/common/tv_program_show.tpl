<ons-list>
  <ons-list-header style="background-color:ghostwhite;">
    <div class="area-left">
      <span id="tv-program-week"></span><span id="tv-program-hour"></span>
    </div>
    <div class="area-right list-margin">
      <i class="fas fa-eye"></i>：{{.TvProgram.CountClicked}}
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
              <ons-col id="tv-cast"></ons-col>
            </ons-row>
            <ons-row class="list-margin-bottom">
              <ons-col width="20%"
                ><i class="fas fa-music" style="color: cornflowerblue;"></i
                >：</ons-col
              >
              <ons-col id="tv-themesong"></ons-col>
            </ons-row>
            <ons-row class="list-margin-bottom">
              <ons-col width="20%">監督：</ons-col>
              <ons-col id="tv-supervisor"></ons-col>
            </ons-row>
            <ons-row class="list-margin-bottom">
              <ons-col width="20%">脚本：</ons-col>
              <ons-col id="tv-dramatist"></ons-col>
            </ons-row>
            <ons-row class="list-margin-bottom">
              <ons-col width="20%">演出：</ons-col>
              <ons-col id="tv-director"></ons-col>
            </ons-row>
            <ons-row class="list-margin-bottom">
              <ons-col width="20%">制作：</ons-col>
              <ons-col id="tv-production"></ons-col>
            </ons-row>
            <ons-row class="list-margin-bottom">
              <ons-col class="category-area" id="tv-category"></ons-col>
            </ons-row>
          </div>
        </ons-col>
      </ons-row>
      <div id="tv-image" class="area-center"></div>
      <div id="tv-reference" class="reference"></div>
      <ons-list-item expandable>
        あらすじ・見どころ
        <div class="right"></div>
        <div class="expandable-content">{{.TvProgram.Content}}</div>
      </ons-list-item>
      <ons-list-item>
        <div class="left">
          {{ if .User }}
          <i
            class="fa-laugh-beam"
            id="check-watched"
            onclick="clickWatchStatus(this, {{.User.Id}}, {{.TvProgram.Id}})"
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
            onclick="clickWatchStatus(this, {{.User.Id}}, {{.TvProgram.Id}})"
          ></i>
          <div id="check-wtw-text" style="float:right; margin-left: 5px;">
            また今度：{{.TvProgram.CountWantToWatch}}
          </div>
          {{ else }}
          <i
            class="fa-laugh-beam"
            id="check-watched"
            onclick="clickWatchStatus(this, '', {{.TvProgram.Id}})"
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
            onclick="clickWatchStatus(this, '', {{.TvProgram.Id}})"
          ></i>
          <div id="check-wtw-text" style="float:right; margin-left: 5px;">
            また今度：{{.TvProgram.CountWantToWatch}}
          </div>
          {{ end }}
        </div>
        <div class="right">
          <div id="edit-tv-program">
            <ons-button
              modifier="quiet"
              onclick="goOtherPage({{.User.Id}},{{.TvProgram.Id}},'/tv/tv_program/edit/{{.TvProgram.Id}}')"
              >編集</ons-button
            >
          </div>
        </div>
      </ons-list-item>
    </div>
  </ons-list-item>
</ons-list>
