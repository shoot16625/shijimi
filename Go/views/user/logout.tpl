<!DOCTYPE html>
<html lang="ja">
  <head>
    {{ template "/common/header.tpl" . }}
    {{ template "/common/alert.tpl" . }}
  </head>

  <body>
    <ons-page>
      {{ template "/common/toolbar.tpl" . }}
      <ons-row>
        <ons-col class="area-center">
          <img
            style="width: 80%;height: auto; max-width: 600px;"
            src="http://gahag.net/img/201604/05s/gahag-0072969531-1.jpg"
            alt="3秒後にTopへ移動．"
          />
        </ons-col>
      </ons-row>
      <ons-row>
        <ons-col class="area-right">登録ユーザ数：</ons-col>
        <ons-col>{{.Info.CntUsers}}</ons-col>
      </ons-row>
      <ons-row>
        <ons-col class="area-right">登録テレビ数：</ons-col>
        <ons-col>{{.Info.CntTvPrograms}}</ons-col>
      </ons-row>
      <div class="toast toast--material">
        <div class="toast__message toast--material__message">
          {{.Status}}
        </div>
      </div>
    </ons-page>

    <script type="text/javascript" src="/static/js/common.js"></script>
    <script type="text/javascript">
      setTimeout(function() {
        window.location.href = URL;
      }, 3 * 1000);
    </script>
  </body>
</html>
