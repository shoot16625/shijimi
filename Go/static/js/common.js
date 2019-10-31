// 定数
const URL = 'http://192.168.2.174:8080';
//const URL = "http://192.168.2.174:8081";
// const URL = "http://www.cmplx.cse.nagoya-u.ac.jp";
//const URL = "localhost:8080";

// 一度のみでいい
// ons.bootstrap();

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
  const topThreshold = 30;
  const scrollSpeedThreshold = 300;
  $('.page__content').on('scroll', function() {
    diff = pos - $(this).scrollTop();
    if (diff < scrollSpeedThreshold) {
      if ($(this).scrollTop() < topThreshold) {
        $('ons-toolbar').show();
      } else {
        if ($(this).scrollTop() < pos) {
          $('ons-toolbar').fadeIn();
        } else {
          $('ons-toolbar').fadeOut();
        }
        pos = $(this).scrollTop();
      }
    }
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
    if (userID === null) {
      return dialogBoxEveryone('alert-only-user-dialog');
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
    $('#' + elem.id).css({ color: newColor });
    checkFlag = true;
    newCount = parseInt(count.textContent.slice(1), 10) + 1;
  } else {
    elem.classList.remove('fas');
    elem.classList.add('far');
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
  const d = data.split('、');
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
function goOtherPage(userID, path) {
  if (userID === null) {
    return dialogBoxEveryone('alert-only-user-dialog');
  } else {
    window.location.href = path;
  }
}
