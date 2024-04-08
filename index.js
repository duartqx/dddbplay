const playButton = document.querySelector("#play");

playButton?.addEventListener("click", function () {
  fetch("http://127.0.0.1:8088/play", {
    method: "GET",
  }).then((res) => {
    res.text().then((t) => {
      playButton.innerHTML = t;
    });
  });
});
