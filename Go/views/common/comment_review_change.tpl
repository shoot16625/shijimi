<div>
  <ons-row>
    <ons-col>
      <button
        class="button--cta"
        style="width: 100%;line-height: 15px;background-color: cornflowerblue;"
        onclick="location.href='/tv/tv_program/comment/{{.TvProgram.Id}}'"
      >
        <i class="fas fa-comments"></i>Timeline
      </button>
    </ons-col>
    <ons-col>
      <button
        class="button--cta"
        style="width: 100%;line-height: 15px;background-color: rgb(240, 141, 20);"
        onclick="location.href='/tv/tv_program/review/{{.TvProgram.Id}}'"
      >
        <i class="fas fa-user"></i>Review
      </button>
    </ons-col>
  </ons-row>
</div>
