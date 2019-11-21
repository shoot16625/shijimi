<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
<meta http-equiv="Content-Style-Type" content="text/css" />
<meta http-equiv="Content-Script-Type" content="text/javascript" />
<meta
  name="viewport"
  content="width=device-width, initial-scale=1, shrink-to-fit=no"
/>
<meta
  name="description"
  content="ドラマや映画のコメントを投稿。みんなで盛り上がろう！！AIによるおすすめ機能もあるよ。"
/>
<meta name="keywords" content="レビュー,映画,ドラマ,ツイート,おすすめ" />

<link rel="shijimi icon" href="/static/img/shijimi-48x48.ico" />

<title>ShiJimi</title>

<!-- <link rel="stylesheet" href="https://unpkg.com/onsenui/css/onsenui.min.css" /> -->
<link
  rel="stylesheet"
  href="https://cdnjs.cloudflare.com/ajax/libs/onsen/2.10.10/css/onsenui.min.css"
/>
<!-- <link
  rel="stylesheet"
  href="https://unpkg.com/onsenui/css/onsen-css-components.min.css"
/> -->
<link
  rel="stylesheet"
  href="https://cdnjs.cloudflare.com/ajax/libs/onsen/2.10.10/css/onsen-css-components.min.css"
/>

<link rel="stylesheet" type="text/css" href="/static/css/common.css" />
<link rel="manifest" href="/manifest.json" />
<script>
  window.addEventListener('load', function() {
    if ('serviceWorker' in navigator) {
      navigator.serviceWorker
        .register('/serviceWorker.js')
        .then(function(registration) {
          console.log('serviceWorker registed.');
        })
        .catch(function(error) {
          console.warn('serviceWorker error.', error);
        });
    }
  });
</script>
