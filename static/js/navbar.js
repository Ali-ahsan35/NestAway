(function () {
  function renderCurrency(currency) {
    var codeEls = document.querySelectorAll(".js-navbar-currency-code");
    var symbolEls = document.querySelectorAll(".js-navbar-currency-symbol");

    codeEls.forEach(function (el) {
      el.textContent = currency.code;
    });
    symbolEls.forEach(function (el) {
      el.textContent = currency.symbol;
    });

    var items = document.querySelectorAll(".st-currency__item");
    items.forEach(function (item) {
      var key = item.getAttribute("data-currency-key");
      item.classList.toggle("st-currency__item--active", key === currency.key);
    });
  }

  function closeAllMenus() {
    var wrappers = document.querySelectorAll("[data-currency-select]");
    wrappers.forEach(function (wrap) {
      wrap.classList.remove("st-currency--open");
      var toggle = wrap.querySelector("[data-currency-toggle]");
      if (toggle) {
        toggle.setAttribute("aria-expanded", "false");
      }
    });
  }

  function buildCurrencyMenu(menuEl, currencies) {
    if (!menuEl) {
      return;
    }

    menuEl.innerHTML = "";
    currencies.forEach(function (currency) {
      var li = document.createElement("li");
      li.setAttribute("role", "none");

      var btn = document.createElement("button");
      btn.type = "button";
      btn.setAttribute("role", "menuitem");
      btn.className = "st-currency__item";
      btn.setAttribute("data-currency-key", currency.key);
      btn.textContent = currency.country + " - " + currency.code + " (" + currency.symbol + ")";

      btn.addEventListener("click", function () {
        if (window.CurrencyManager) {
          window.CurrencyManager.setCurrencyByKey(currency.key);
        }
        closeAllMenus();
      });

      li.appendChild(btn);
      menuEl.appendChild(li);
    });
  }

  function initCurrencyMenu(wrapper) {
    var toggle = wrapper.querySelector("[data-currency-toggle]");
    if (!toggle) {
      return;
    }

    toggle.addEventListener("click", function (event) {
      event.stopPropagation();
      var isOpen = wrapper.classList.contains("st-currency--open");
      closeAllMenus();
      if (!isOpen) {
        wrapper.classList.add("st-currency--open");
        toggle.setAttribute("aria-expanded", "true");
      }
    });

  }

  document.addEventListener("click", function () {
    closeAllMenus();
  });

  document.addEventListener("DOMContentLoaded", function () {
    var wrappers = document.querySelectorAll("[data-currency-select]");
    wrappers.forEach(initCurrencyMenu);

    if (!window.CurrencyManager) {
      return;
    }

    window.CurrencyManager.ready().then(function () {
      var currencies = window.CurrencyManager.getCurrencies();
      wrappers.forEach(function (wrapper) {
        var menu = wrapper.querySelector("[data-currency-menu]");
        buildCurrencyMenu(menu, currencies);
      });
      renderCurrency(window.CurrencyManager.getCurrentCurrency());
    });

    window.addEventListener("currency:changed", function (event) {
      renderCurrency(event.detail);
    });
  });
})();
