(function () {
  var STORAGE_KEY = "selectedCurrency";
  var DATA_URL = "/static/data/currencies.json";
  var FALLBACK = {
    key: "US",
    country: "US",
    code: "USD",
    symbol: "US $",
    rate: 1,
  };

  var state = {
    currencies: [],
    byKey: {},
    current: FALLBACK,
    initialized: false,
    initPromise: null,
  };

  function normalize(raw) {
    var list = [];
    var byKey = {};

    Object.keys(raw || {}).forEach(function (key) {
      var item = raw[key] || {};
      var parsed = {
        key: key,
        country: item.Country || key,
        code: item.Code || "USD",
        symbol: item.Symbol || "US $",
        rate: Number(item.Rate) || 1,
      };
      byKey[key] = parsed;
      list.push(parsed);
    });

    list.sort(function (a, b) {
      return a.country.localeCompare(b.country);
    });

    return { list: list, byKey: byKey };
  }

  function readSaved() {
    try {
      var raw = localStorage.getItem(STORAGE_KEY);
      if (!raw) {
        return null;
      }
      var parsed = JSON.parse(raw);
      if (!parsed || !parsed.code) {
        return null;
      }
      return parsed;
    } catch (error) {
      return null;
    }
  }

  (function hydrateFromStorage() {
    var saved = readSaved();
    if (saved && saved.code && saved.symbol && Number(saved.rate)) {
      state.current = {
        key: saved.key || "US",
        country: saved.country || (saved.key || "US"),
        code: saved.code,
        symbol: saved.symbol,
        rate: Number(saved.rate) || 1,
      };
    }
  })();

  function saveCurrent() {
    localStorage.setItem(STORAGE_KEY, JSON.stringify(state.current));
  }

  function findByCode(code) {
    for (var i = 0; i < state.currencies.length; i += 1) {
      if (state.currencies[i].code === code) {
        return state.currencies[i];
      }
    }
    return null;
  }

  function findFromSaved(saved) {
    if (!saved) {
      return null;
    }
    if (saved.key && state.byKey[saved.key]) {
      return state.byKey[saved.key];
    }
    if (saved.code) {
      return findByCode(saved.code);
    }
    return null;
  }

  function formatNumber(value) {
    return new Intl.NumberFormat("en-US", {
      maximumFractionDigits: 0,
    }).format(Math.round(value));
  }

  function updatePriceNode(node, currency) {
    var base = Number(node.getAttribute("data-base-usd"));
    if (!Number.isFinite(base)) {
      node.textContent = "Price unavailable";
      return;
    }

    var prefix = node.getAttribute("data-price-prefix") || "";
    var suffix = node.getAttribute("data-price-suffix") || "";
    var converted = base * currency.rate;
    node.textContent = prefix + currency.symbol + " " + formatNumber(converted) + suffix;
  }

  function applyCurrencyToDom(root) {
    var target = root || document;
    var currency = state.current;

    target.querySelectorAll(".js-price-value[data-base-usd]").forEach(function (node) {
      updatePriceNode(node, currency);
    });

    target.querySelectorAll(".js-currency-symbol-dynamic").forEach(function (node) {
      node.textContent = currency.symbol;
    });
  }

  function notifyChange() {
    window.dispatchEvent(
      new CustomEvent("currency:changed", {
        detail: state.current,
      })
    );
  }

  function setCurrent(currency) {
    if (!currency) {
      return;
    }
    state.current = {
      key: currency.key,
      country: currency.country,
      code: currency.code,
      symbol: currency.symbol,
      rate: Number(currency.rate) || 1,
    };
    saveCurrent();
    applyCurrencyToDom(document);
    notifyChange();
  }

  function init() {
    if (state.initPromise) {
      return state.initPromise;
    }

    state.initPromise = fetch(DATA_URL)
      .then(function (res) {
        return res.json();
      })
      .then(function (json) {
        var normalized = normalize(json);
        state.currencies = normalized.list;
        state.byKey = normalized.byKey;

        var saved = readSaved();
        var selected = findFromSaved(saved) || state.byKey.US || findByCode("USD") || FALLBACK;
        state.current = selected;
        saveCurrent();
        state.initialized = true;
        applyCurrencyToDom(document);
        notifyChange();
      })
      .catch(function () {
        state.currencies = [FALLBACK];
        state.byKey = { US: FALLBACK };
        state.current = FALLBACK;
        saveCurrent();
        state.initialized = true;
        applyCurrencyToDom(document);
        notifyChange();
      });

    return state.initPromise;
  }

  window.CurrencyManager = {
    init: init,
    ready: init,
    isReady: function () {
      return state.initialized;
    },
    getCurrencies: function () {
      return state.currencies.slice();
    },
    getCurrentCurrency: function () {
      return state.current;
    },
    getCurrencyByCode: function (code) {
      return findByCode(code);
    },
    setCurrencyByKey: function (key) {
      if (state.byKey[key]) {
        setCurrent(state.byKey[key]);
      }
    },
    setCurrencyByCode: function (code) {
      var found = findByCode(code);
      if (found) {
        setCurrent(found);
      }
    },
    convertFromUSD: function (amountUsd) {
      var value = Number(amountUsd);
      if (!Number.isFinite(value)) {
        return null;
      }
      return value * state.current.rate;
    },
    applyCurrencyToDom: applyCurrencyToDom,
    formatNumber: formatNumber,
  };

  // Start loading currencies as soon as script is loaded.
  init();
})();
