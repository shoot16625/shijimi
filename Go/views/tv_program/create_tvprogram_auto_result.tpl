<!DOCTYPE html>
<html lang="ja">
  <head>
    {{ template "/common/header.tpl" . }}
  </head>

  <body>
    <ons-page>
      {{ template "/common/toolbar.tpl" . }}
      {{ template "/common/alert.tpl" . }}
      <ons-pull-hook id="pull-hook"></ons-pull-hook>
      <p>番組登録が完了しました！！</p>
    </ons-page>
    {{ template "/common/js.tpl" . }}
  </body>
</html>
