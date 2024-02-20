const scoreSelector = document.querySelector("#scoreSelector");
const increaseScorePlayerOne = document.querySelector(
  "#increaseScorePlayerOne",
);
const increaseScorePlayerTwo = document.querySelector(
  "#increaseScorePlayerTwo",
);
const reset = document.querySelector("#reset");

const scorePlayerOne = document.querySelector("#scorePlayerOne");
const scorePlayerTwo = document.querySelector("#scorePlayerTwo");

scoreSelector.addEventListener("change", function (eve) {
  console.log(eve);
});

function disableButtons() {
  increaseScorePlayerOne.disabled = true;
  increaseScorePlayerTwo.disabled = true;
}

increaseScorePlayerOne.addEventListener("click", function () {
  score = scorePlayerOne.innerText;
  if (score === scoreSelector.value) {
    disableButtons();
  } else {
    scorePlayerOne.innerText = Number(score) + 1;
    scorePlayerOne.classList.add("winner");
    scorePlayerTwo.classList.add("loser");
  }
});

increaseScorePlayerTwo.addEventListener("click", function () {
  score = scorePlayerTwo.innerText;
  if (score === scoreSelector.value) {
    disableButtons();
  } else {
    scorePlayerTwo.innerText = Number(score) + 1;
    scorePlayerOne.classList.add("winner");
    scorePlayerTwo.classList.add("loser");
  }
});

reset.addEventListener("click", function () {
  scorePlayerOne.innerText = 0;
  scorePlayerTwo.innerText = 0;
  disableButtons();
});
