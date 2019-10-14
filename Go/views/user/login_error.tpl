<!DOCTYPE html>
<html lang="ja">
<head>
  {{ template "/common/header.tpl" . }}
  {{ template "/common/alert.tpl" . }}

</head>

<body>
  <ons-page>
    {{ template "/common/toolbar.tpl" . }}
    <div class="toast toast--material">
      <div class="toast__message toast--material__message">ログインに失敗しました</div>
    </div>
  </ons-page>

  <script type="text/javascript" src="/static/js/common.js"></script>
  <script type="text/javascript">
    setTimeout(function(){
     window.location.href = URL;
   }, 3*1000);
 </script>
</body>
</html>