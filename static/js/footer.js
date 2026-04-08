document.addEventListener("DOMContentLoaded", function () {
  var yearEl = document.getElementById("js-footer-year");
  if (yearEl) {
    yearEl.textContent = String(new Date().getFullYear());
  }
});
