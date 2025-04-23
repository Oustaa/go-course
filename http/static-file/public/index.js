let counter = 0;

const counterElem = document.querySelector(".counterElem");

const incBtn = document.querySelector("button#inc");
const decBtn = document.querySelector("button#dec");

console.log(incBtn);

counterElem.innerHTML = counter;
incBtn.addEventListener("click", () => {
  counter++;
  counterElem.innerHTML = counter;
});

decBtn.addEventListener("click", () => {
  counter--;
  counterElem.innerHTML = counter;
});

