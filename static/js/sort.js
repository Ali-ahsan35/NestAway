const sortWrap = document.getElementById("js-filter-sort");
const defaultOpt = sortWrap.querySelector(".default-option");
const hiddenInput = document.getElementById("sort-properties");

defaultOpt.addEventListener("click", () => {
  sortWrap.classList.toggle("active");
});

sortWrap.querySelectorAll(".select-ul li").forEach((li) => {
  li.addEventListener("click", () => {
    const value = li.getAttribute("data-value");
    const text = li.querySelector("p").textContent;
    defaultOpt.querySelector("p").textContent = text;
    hiddenInput.value = value;
    sortWrap.classList.remove("active");

    // Call loadProperties from refine.js
    if (window.loadProperties && window.currentCategory) {
        // Read current filters from URL so they are preserved
        const urlParams = new URLSearchParams(window.location.search);
        const savedFilters = {};
        if (urlParams.get('amenities')) {
            savedFilters.amenities = urlParams.get('amenities').split('-');
        }
        if (urlParams.get('ecoFriendly')) {
            savedFilters.ecoFriendly = true;
        }
        if (urlParams.get('amount')) {
          const selectedCode = urlParams.get('selectedCurrency') || (window.CurrencyManager && window.CurrencyManager.getCurrentCurrency ? window.CurrencyManager.getCurrentCurrency().code : 'USD');
          let rate = 1;
          if (window.CurrencyManager && window.CurrencyManager.getCurrencyByCode) {
            const byCode = window.CurrencyManager.getCurrencyByCode(selectedCode);
            if (byCode) {
              rate = Number(byCode.rate) || 1;
            }
          }

          // amount in URL is display currency — convert to USD for API.
            const parts = urlParams.get('amount').split('-');
          const minUSD = Math.round(parseInt(parts[0], 10) / rate);
          const maxUSD = Math.round(parseInt(parts[1], 10) / rate);
            savedFilters.amount = minUSD + '-' + maxUSD;   
          savedFilters.amountDisplay = urlParams.get('amount');
          savedFilters.selectedCurrency = selectedCode;
        }
        if (urlParams.get('pax')) {
            savedFilters.guests = parseInt(urlParams.get('pax'));
        }
        if (urlParams.get('dateStart')) {
            savedFilters.checkin = urlParams.get('dateStart');
        }
        if (urlParams.get('dateEnd')) {
            savedFilters.checkout = urlParams.get('dateEnd');
        }
        console.log("sort savedFilters:", savedFilters);
        window.loadProperties(window.currentCategory, value, savedFilters);
    }
  });
});

document.addEventListener("click", (e) => {
  if (!sortWrap.contains(e.target)) sortWrap.classList.remove("active");
});
