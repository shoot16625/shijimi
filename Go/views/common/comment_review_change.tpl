<div style="text-align: center;">
  <ons-row>
    <!-- <i class="fas fa-arrow-left"></i>左にスワイプ -->
    <ons-col style="text-align: right;">
      <button class="button--cta" style="line-height: 15px; background-color: cornflowerblue;" onclick="location.href='/tv/tv_program/comment/{{.TvProgram.Id}}'"><i class="fas fa-comments"></i>Timeline</button>
    </ons-col>
    <ons-col style="text-align: left;">
      <button class="button--cta" style="line-height: 15px; background-color: darkorange;" onclick="location.href='/tv/tv_program/review/{{.TvProgram.Id}}'"><i class="fas fa-user"></i>Review</button>
    </ons-col>
    <!-- 右にスワイプ<i class="fas fa-arrow-right"></i> -->
  </ons-row>
</div>

<div style="text-align: center;">
  <i class="fas fa-arrow-left"></i>左にスワイプ
  右にスワイプ<i class="fas fa-arrow-right"></i>
</div>
