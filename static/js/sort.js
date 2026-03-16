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
            // amount in URL is BDT — convert to USD for API
            const parts = urlParams.get('amount').split('-');
            const minUSD = Math.round(parseInt(parts[0]) / 120);
            const maxUSD = Math.round(parseInt(parts[1]) / 120);
            savedFilters.amount = minUSD + '-' + maxUSD;   
            savedFilters.amountBDT = urlParams.get('amount');
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
