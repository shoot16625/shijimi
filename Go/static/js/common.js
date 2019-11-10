// 定数
const URL = 'http://192.168.2.174:8080';
//const URL = "http://192.168.2.174:8081";
// const URL = 'http://www.cmplx.cse.nagoya-u.ac.jp';
//const URL = "localhost:8080";
// const URL = "https://shijimi.herokuapp.com";

// 自動スクロール
function autoScroll(varName, len) {
  let indexState = -1;
  setInterval(function() {
    let activeIndex = varName.getActiveIndex();
    if (indexState === activeIndex) {
      varName.first();
      indexState = -1;
    } else {
      varName.next();
      indexState = activeIndex;
    }
  }, 2000);
}

// toolbarを隠す
$(function() {
  let pos = 0;
  let diff = 0;
  const topThreshold = 100;
  const scrollSpeedTop = 100;
  const scrollSpeedBottom = 12;
  $('.page__content').on('scroll', function() {
    diff = pos - $(this).scrollTop();
    if (-scrollSpeedTop < diff && diff < scrollSpeedTop) {
      if ($(this).scrollTop() < topThreshold) {
        $('ons-toolbar').show();
      } else {
        if (scrollSpeedBottom < diff) {
          $('ons-toolbar').fadeIn();
        } else if (-scrollSpeedBottom > diff) {
          $('ons-toolbar').fadeOut();
        }
      }
    }
    pos = $(this).scrollTop();
  });
});

// パスワードを表示するチェックボックス
$(function() {
  $('#password-check').change(function() {
    if ($(this).prop('checked')) {
      $('#password').attr('type', 'text');
    } else {
      $('#password').attr('type', 'password');
    }
  });
});

// アラートを閉じる
var hideAlertDialog = function(elem) {
  document.getElementById(elem).hide();
};

// ツイートボックス
var dialogBox = function(elemID, userID) {
  ons.ready(function() {
    var dialog = document.getElementById(elemID);
    // == でないとダメ
    if (userID == null) {
      return dialogBoxEveryone('alert-only-user-dialog');
    } else if (userID === -1) {
      return dialogBoxEveryone('alert-review-twice');
    }
    if (dialog) {
      dialog.show();
    } else {
      ons
        .createElement(elemID + '.html', { append: true })
        .then(function(dialog) {
          dialog.show();
        });
    }
  });
};

// not-userも許可するダイアログボックス
var dialogBoxEveryone = function(elemID) {
  ons.ready(function() {
    var dialog = document.getElementById(elemID);
    if (dialog) {
      dialog.show();
    } else {
      ons
        .createElement(elemID + '.html', { append: true })
        .then(function(dialog) {
          dialog.show();
        });
    }
  });
};

// プルフック
var pullHook = function() {
  var pullHook = document.getElementById('pull-hook');
  if (pullHook != null) {
    pullHook.addEventListener('changestate', function(event) {
      let message = '';
      switch (event.state) {
        case 'initial':
          message = 'Pull to refresh';
          break;
        case 'preaction':
          message = 'Release';
          break;
        case 'action':
          message = 'Loading...';
          break;
      }
      pullHook.innerHTML = message;
    });
    pullHook.onAction = function(done) {
      setTimeout(window.location.reload(false), 2000);
    };
  }
};

// いいねボタンの色切り替え
function setLikeStatus(likeStatus, newColor) {
  if (likeStatus) {
    return newColor;
  } else {
    return 'black';
  }
}

// いいねボタンの色切り替え
function setLikeBold(likeStatus) {
  if (likeStatus) {
    return 'fas';
  } else {
    return 'far';
  }
}

// 見たボタンの色の切り替え
function setWatchStatus(elemID, newColor, status) {
  var target = document.getElementById(elemID);
  if (status) {
    target.style.color = newColor;
  } else {
    target.style.color = 'black';
  }
}

// 見たボタンの色の切り替え
function setWatchBold(elemID, status) {
  var target = document.getElementById(elemID);
  if (status) {
    target.classList.add('fas');
  } else {
    target.classList.add('far');
  }
}

// いいねボタンがクリックされたら色を変える
function clickLike(elem) {
  const newColor = 'orchid';
  if (globalCommentLikeStatus === null) {
    return dialogBoxEveryone('alert-only-user-dialog');
  }
  var count = document.getElementById('count-like-' + elem.id);
  let checkFlag;
  let newCount;
  if (elem.style['color'] != newColor) {
    elem.classList.remove('far');
    elem.classList.add('fas');
    elem.classList.add('bounce-animation');
    $('#' + elem.id).css({ color: newColor });
    checkFlag = true;
    newCount = parseInt(count.textContent.slice(1), 10) + 1;
  } else {
    elem.classList.remove('fas');
    elem.classList.add('far');
    elem.classList.remove('bounce-animation');
    $('#' + elem.id).css({ color: 'black' });
    checkFlag = false;
    newCount = parseInt(count.textContent.slice(1), 10) - 1;
  }
  count.textContent = '：' + newCount;
  commentLikeStatus(elem, checkFlag);
}

// 見たボタンのクリック処理
function clickWatchStatus(elem) {
  if (globalWatchStatus === null) {
    return dialogBoxEveryone('alert-only-user-dialog');
  }
  var count = document.getElementById(elem.id + '-text');
  const str = 'check-watched';
  let newColor = 'lightseagreen';
  if (elem.id.indexOf(str) === 0) {
    newColor = 'deeppink';
  }
  let checkFlag;
  let rawText = count.textContent.trim();
  if (elem.style['color'] != newColor) {
    $('#' + elem.id).css({ color: newColor });
    elem.classList.remove('far');
    elem.classList.add('fas');
    elem.classList.add('bounce-animation');
    checkFlag = true;
    if (elem.id.indexOf(str) === 0) {
      newCount = parseInt(rawText.slice(3), 10) + 1;
    } else {
      newCount = parseInt(rawText.slice(5), 10) + 1;
    }
  } else {
    $('#' + elem.id).css({ color: 'black' });
    elem.classList.remove('fas');
    elem.classList.add('far');
    elem.classList.remove('bounce-animation');
    checkFlag = false;
    if (elem.id.indexOf(str) === 0) {
      newCount = parseInt(rawText.slice(3), 10) - 1;
    } else {
      newCount = parseInt(rawText.slice(5), 10) - 1;
    }
  }
  if (elem.id.indexOf(str) === 0) {
    count.textContent = '見た：' + newCount;
  } else {
    count.textContent = 'また今度：' + newCount;
  }
  WatchStatus(elem, checkFlag);
}

// セレクタが複数設定されていた時の再描画処理
function setMultipleSelection(elem, data) {
  const d = data.split(',');
  for (let i = d.length - 1; i >= 0; i--) {
    var target = document.getElementById(elem);
    for (let j = target.length - 1; j >= 0; j--) {
      if (target.options[j].value == d[i]) {
        target.options[j].selected = true;
      }
    }
  }
}

// ページの上部へ移動
function goTop() {
  $('.page__content').animate({ scrollTop: 0 }, 500, 'swing');
}

// カルーセルを移動してページのトップへ移動
function goAnotherCarousel(index) {
  ons.ready(function() {
    carousel.setActiveIndex(index);
    goTop();
  });
}

// pathのページへ移動
function goOtherPage(userID, tvProgramID, path) {
  if (userID == null) {
    return dialogBoxEveryone('alert-only-user-dialog');
  } else if (userID != 1 && tvProgramID === 1) {
    return dialogBoxEveryone('alert-only-user-dialog');
  } else {
    window.location.href = path;
  }
}

// ローディングマークの表示
function showLoading() {
  $('.page__content').html(
    '<ons-progress-bar indeterminate></ons-progress-bar>'
  );
}

// ポイントで買えるバッジの整形
function reshapeBadges(badge) {
  let badgeText = '';
  if (badge != '') {
    let badges = badge.split(',');
    for (let index = 0; index < badges.length; index++) {
      if (badges[index] == 'thanks') {
        badgeText +=
          "<span class='badge-icon' style='margin-right:7px;'><i class='fab fa-angellist' style ='color:cornflowerblue; font-size:30px;'></i><span style='font-size:10px;'>寄付</span></span>";
      }
    }
    document.getElementById('badges').innerHTML = badgeText;
  }
}

// 放送時間の整形
function reshapeHour(time) {
  let str = '.5';
  if (time === '100') {
    time = '';
  } else {
    if (time.indexOf(str) > -1) {
      time = time.replace(str, ':30');
    } else {
      time += ':00';
    }
  }
  return time;
}

// 季節ごとのヘッダーカラー
function seasonHeaderColor(seasonName) {
  let headerColor;
  if (seasonName === '春') {
    headerColor = 'lavenderblush';
  } else if (seasonName === '夏') {
    headerColor = 'aliceblue';
  } else if (seasonName === '秋') {
    headerColor = 'khaki';
  } else if (seasonName === '冬') {
    headerColor = 'thistle';
  } else {
    headerColor = 'ghostwhite';
  }
  return headerColor;
}

// 時間帯選択用
function getSelectHour(textTop) {
  let textHour = textTop;
  let t;
  for (let i = 40; i <= 48; i++) {
    if (i % 2 == 0) {
      t = String(i / 2) + ':00';
      textHour += '<option>' + t + '</option>';
    } else {
      t = String((i - 1) / 2) + ':30';
      textHour += '<option>' + t + '</option>';
    }
  }
  for (let i = 1; i <= 39; i++) {
    if (i % 2 == 0) {
      t = String(i / 2) + ':00';
      textHour += '<option>' + t + '</option>';
    } else {
      t = String((i - 1) / 2) + ':30';
      textHour += '<option>' + t + '</option>';
    }
  }
  return textHour;
}

// コンテンツの整形用
function reshapeContents(elements) {
  let text = '';
  for (let index = 0; index < elements.length; index++) {
    text += "<span class='new-line'>" + elements[index] + '&nbsp;</span>';
  }
  return text;
}
// コンテンツの整形用
function reshapeContent(element) {
  let text = '';
  let elements = element.split(',');
  for (let index = 0; index < elements.length; index++) {
    text += "<span class='new-line'>" + elements[index] + '&nbsp;</span>';
  }
  return text;
}

// 動画・画像の埋め込み変換
function reshapeMovieCode(tvPrograms) {
  let moviePosition;
  if (tvPrograms.MovieUrl === '') {
    moviePosition =
      '<img class="image" id="image-' +
      tvPrograms.Id +
      '" src="' +
      tvPrograms.ImageUrl +
      '" alt="' +
      tvPrograms.Title +
      '" width="80%" onerror="this.src=\'/static/img/tv_img/hanko_02.png\'">';
  } else {
    moviePosition =
      '<iframe id="movie-' +
      tvPrograms.Id +
      '" class="movie" src="' +
      tvPrograms.MovieUrl +
      '?modestbranding=1&rel=0&playsinline=1" frameborder="0" alt="' +
      tvPrograms.Title +
      '" allow="accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>';
  }
  return moviePosition;
}
// 出典の所
function reshapeReferenceSite(tvPrograms) {
  let referenceSite = '';
  if (tvPrograms.ImageUrlReference != '' && tvPrograms.MovieUrl == '') {
    referenceSite =
      '<a href=' +
      tvPrograms.ImageUrl +
      " target='_blank'>出典:" +
      tvPrograms.ImageUrlReference +
      '</a>';
  } else if (tvPrograms.MovieUrl != '') {
    referenceSite =
      '<a href=' + tvPrograms.MovieUrl + " target='_blank'>出典:Youtube</a>";
  }
  return referenceSite;
}

function inputPreviewData() {
  let hour;
  if (document.getElementsByName('hour')[0].value === '指定なし') {
    hour = '';
  } else {
    hour = document.getElementsByName('hour')[0].value;
  }
  document.getElementById('preview-on-air-info').innerHTML =
    document.getElementsByName('year')[0].value +
    '年 ' +
    document.getElementsByName('season')[0].value.replace(/\(.+\)/, '') +
    '（' +
    document.getElementsByName('week')[0].value +
    hour +
    '）';
  document.getElementById(
    'preview-title'
  ).innerHTML = document.getElementsByName('title')[0].value;
  document.getElementById(
    'preview-content'
  ).innerHTML = document.getElementsByName('content')[0].value;
  document.getElementById(
    'preview-cast'
  ).innerHTML = document.getElementsByName('cast')[0].value.replace(/\,/g, ' ');
  document.getElementById(
    'preview-themesong'
  ).innerHTML = document
    .getElementsByName('themesong')[0]
    .value.replace(/\,/g, ' ');
  let categories = document.getElementById('category');
  let category = '';
  let tag;
  for (let index = 0; index < categories.length; index++) {
    tag = categories[index];
    if (tag.selected) {
      category += tag.value + ' ';
    }
  }
  document.getElementById('preview-category').innerHTML = category;
  document.getElementById(
    'preview-production'
  ).innerHTML = document
    .getElementsByName('production')[0]
    .value.replace(/\,/g, ' ');
  document.getElementById(
    'preview-dramatist'
  ).innerHTML = document
    .getElementsByName('dramatist')[0]
    .value.replace(/\,/g, ' ');
  document.getElementById(
    'preview-supervisor'
  ).innerHTML = document
    .getElementsByName('supervisor')[0]
    .value.replace(/\,/g, ' ');
  document.getElementById(
    'preview-director'
  ).innerHTML = document
    .getElementsByName('director')[0]
    .value.replace(/\,/g, ' ');
  let imageURL = document.getElementsByName('ImageURL')[0].value;
  if (imageURL === '') {
    imageURL = '/static/img/tv_img/hanko_02.png';
  }
  document.getElementById('preview-img').innerHTML =
    '<img src="' +
    imageURL +
    '" alt="イメージ" width="80%" onerror="this.src=\'/static/img/tv_img/hanko_02.png\'">';
  var movieURL = document.getElementsByName('MovieURL')[0].value;
  if (movieURL != '') {
    movieURL = movieURL.replace('watch?v=', 'embed/');
    movieURL = movieURL.replace('m.youtube.com', 'www.youtube.com');
    movieURL = movieURL.replace('youtu.be/', 'www.youtube.com/embed/');
    document.getElementById('preview-movie').innerHTML =
      '<iframe src="' +
      movieURL +
      '?modestbranding=1&rel=0&playsinline=1" frameborder="0" alt="ムービー" width="200" height="112.5" allow="accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>';
  }
}
