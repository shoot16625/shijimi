<!DOCTYPE html>
<html lang="ja">
  <head>
    {{ template "/common/header.tpl" . }}
    {{ template "/common/alert.tpl" . }}
  </head>

  <body>
    <ons-page>
      {{ template "/common/toolbar.tpl" . }}
      <div style="margin:50px 0 30px 0;">
        <ons-row>
        <div
          class="tenor-gif-embed"
          data-postid="10252927"
          data-share-method="host"
          data-width="100%"
          data-aspect-ratio="3.090909090909091"
        >
          <a href="https://tenor.com/view/thank-you-animated-gif-10252927"
            >Thank You Animated GIF</a
          >
          from
          <a href="https://tenor.com/search/thankyou-gifs">Thankyou GIFs</a>
        </div>
      </div>
    </ons-row>
      <ons-row>
        <ons-col class="area-right">登録ユーザ数：</ons-col>
        <ons-col>{{.Info.CntUsers}}</ons-col>
      </ons-row>
      <ons-row>
        <ons-col class="area-right">登録テレビ数：</ons-col>
        <ons-col>{{.Info.CntTvPrograms}}</ons-col>
      </ons-row>
      <div class="toast">
        <div class="toast__message">
          {{.Status}}
        </div>
      </div>
    </ons-page>

    {{ template "/common/js.tpl" . }}
    <script
    type="text/javascript"
    async
    src="https://tenor.com/embed.js"
  ></script>
    <script type="text/javascript">
        setTimeout(function() {
          window.location.href = URL;
        }, 4 * 1000);
    </script>
  </body>
</html>
