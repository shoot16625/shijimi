<ons-row>
  <ons-col>
    <button
      class="button--cta"
      style="line-height: 15px; background-color: cornflowerblue; width: 100%"
      onclick="location.href='/tv/user/show/{{.User.Id}}'"
    >
      <i class="fas fa-comments"></i>TM
    </button>
  </ons-col>
  <ons-col>
    <button
      class="button--cta"
      style="line-height: 15px; background-color: rgb(240, 141, 20); width: 100%"
      onclick="location.href='/tv/user/show_review/{{.User.Id}}'"
    >
      <i class="fas fa-user"></i>Re
    </button>
  </ons-col>
</ons-row>
