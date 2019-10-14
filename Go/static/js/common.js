// 定数
// const URL = "http://192.168.2.174:8080";
// URL = "http://192.168.2.174:8081";
// URL = "http://www.cmplx.cse.nagoya-u.ac.jp";
const URL = "localhost:8080";

// 一度のみでいい
ons.bootstrap();

// 自動スクロール
function AutoScroll(var_name, len) {
  let index_state = -1;
  setInterval(function() {
    let index = var_name.getActiveIndex();
    if (index >= len - 2) {
      var_name.first();
    } else {
      var_name.next();
    }
    if (index == index_state) {
      var_name.first();
    }
    index_state = index;
  }, 2000);
}

// toolbarを隠す
// var scroll_position = 0;
// ons.ready(function() {
// $('.page__content').on('scroll', function(){
//   var scrollTop = $(this).scrollTop();
//   if (scrollTop - scroll_position > 0){
//     document.querySelector('ons-toolbar').hide();
//   } else {
//     document.querySelector('ons-toolbar').show();
//   }
//   scroll_position = scrollTop;
// });
// });

// パスワードを表示するチェックボックス
$(function() {
  $("#password-check").change(function() {
    if ($(this).prop("checked")) {
      $("#password").attr("type", "text");
    } else {
      $("#password").attr("type", "password");
    }
  });
});

// アラートを閉じる
var hideAlertDialog = function(elem) {
  document.getElementById(elem).hide();
};

// ツイートボックス
var DialogBox = function(elem_id, user_id) {
  ons.ready(function() {
    var dialog = document.getElementById(elem_id);
    if (user_id === null) {
      return DialogBoxEveryone("alert_onlyuser_dialog");
    }
    if (dialog) {
      dialog.show();
    } else {
      ons
        .createElement(elem_id + ".html", { append: true })
        .then(function(dialog) {
          dialog.show();
        });
    }
  });
};

// not-userも許可するダイアログボックス
var DialogBoxEveryone = function(elem_id) {
  ons.ready(function() {
    var dialog = document.getElementById(elem_id);
    if (dialog) {
      dialog.show();
    } else {
      ons
        .createElement(elem_id + ".html", { append: true })
        .then(function(dialog) {
          dialog.show();
        });
    }
  });
};

// プルフック
var PullHook = function() {
  var pullHook = document.getElementById("pull-hook");
  if (pullHook != null) {
    pullHook.addEventListener("changestate", function(event) {
      let message = "";
      switch (event.state) {
        case "initial":
          message = "Pull to refresh";
          break;
        case "preaction":
          message = "Release";
          break;
        case "action":
          message = "Loading...";
          break;
      }
      pullHook.innerHTML = message;
    });
    pullHook.onAction = function(done) {
      setTimeout(window.location.reload(false), 1000);
    };
  }
};

// いいねボタンの色切り替え
function SetLikeStatus(like_status, new_color) {
  if (like_status) {
    return new_color;
  } else {
    return "black";
  }
}

// いいねボタンの色切り替え
function SetLikeBold(like_status) {
  if (like_status) {
    return "fas";
  } else {
    return "far";
  }
}

// 見たボタンの色の切り替え
function SetWatchStatus(elem_id, new_color, status) {
  var target = document.getElementById(elem_id);
  if (status) {
    target.style.color = new_color;
  } else {
    target.style.color = "black";
  }
}

// 見たボタンの色の切り替え
function SetWatchBold(elem_id, status) {
  var target = document.getElementById(elem_id);
  if (status) {
    target.classList.add("fas");
  } else {
    target.classList.add("far");
  }
}

// いいねボタンがクリックされたら色を変える
function ClickLike(elem) {
  const new_color = "orchid";
  if (global_comment_like_status === null) {
    return DialogBoxEveryone("alert_onlyuser_dialog");
  }
  var count = document.getElementById("count_like_" + elem.id);
  let check_flag;
  let new_count;
  if (elem.style["color"] != new_color) {
    elem.classList.remove("far");
    elem.classList.add("fas");
    $("#" + elem.id).css({ color: new_color });
    check_flag = true;
    new_count = parseInt(count.textContent.slice(1), 10) + 1;
  } else {
    elem.classList.remove("fas");
    elem.classList.add("far");
    $("#" + elem.id).css({ color: "black" });
    check_flag = false;
    new_count = parseInt(count.textContent.slice(1), 10) - 1;
  }
  count.textContent = "：" + new_count;
  CommentLikeStatus(elem, check_flag);
}

// 見たボタンのクリック処理
function ClickWatchStatus(elem) {
  if (global_watch_status === null) {
    return DialogBoxEveryone("alert_onlyuser_dialog");
  }
  var count = document.getElementById(elem.id + "_text");
  const str = "check_watched";
  let new_color;
  if (elem.id.indexOf(str) === 0) {
    new_color = "lightcoral";
  } else {
    new_color = "lightseagreen";
  }
  let check_flag;
  let new_count;
  if (elem.style["color"] != new_color) {
    $("#" + elem.id).css({ color: new_color });
    elem.classList.remove("far");
    elem.classList.add("fas");
    check_flag = true;
    if (elem.id.indexOf(str) === 0) {
      new_count = parseInt(count.textContent.slice(3), 10) + 1;
    } else {
      new_count = parseInt(count.textContent.slice(5), 10) + 1;
    }
  } else {
    $("#" + elem.id).css({ color: "black" });
    elem.classList.remove("fas");
    elem.classList.add("far");
    check_flag = false;
    if (elem.id.indexOf(str) === 0) {
      new_count = parseInt(count.textContent.slice(3), 10) - 1;
    } else {
      new_count = parseInt(count.textContent.slice(5), 10) - 1;
    }
  }
  if (elem.id.indexOf(str) === 0) {
    count.textContent = "見た：" + new_count;
  } else {
    count.textContent = "また今度：" + new_count;
  }
  WatchStatus(elem, check_flag);
}

// セレクタが複数設定されていた時の再描画処理
function SetMultipleSelection(elem, data) {
  const d = data.split("、");
  for (var i = d.length - 1; i >= 0; i--) {
    var target = document.getElementById(elem);
    for (var j = target.length - 1; j >= 0; j--) {
      if (target.options[j].value == d[i]) {
        target.options[j].selected = true;
      }
    }
  }
}

// ページの上部へ移動
function GoTop() {
  $(".page__content").animate({ scrollTop: 0 }, 500, "swing");
}

// カルーセルを移動してページのトップへ移動
function GoAnotherCarousel(index) {
  ons.ready(function() {
    carousel.setActiveIndex(index);
    GoTop();
  });
}

// pathのページへ移動
function GoOtherPage(user_id, path) {
  if (user_id == null) {
    return DialogBoxEveryone("alert_onlyuser_dialog");
  } else {
    window.location.href = path;
  }
}
