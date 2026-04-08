document.addEventListener("DOMContentLoaded", function () {
  function getCurrency() {
    if (window.CurrencyManager && window.CurrencyManager.getCurrentCurrency) {
      return window.CurrencyManager.getCurrentCurrency();
    }
    return { code: "USD", symbol: "US $", rate: 1 };
  }

  function getBounds() {
    const rate = Number(getCurrency().rate) || 1;
    return {
      min: Math.round(2 * rate),
      max: Math.round(1000 * rate),
      step: Math.max(1, Math.round(rate)),
    };
  }

  function updateCurrencyIcons() {
    const symbol = getCurrency().symbol;
    document.querySelectorAll(".currency-icon").forEach((el) => {
      el.textContent = symbol;
    });
  }

  // Open modal
  function openFilterModal() {
    document.getElementById("js-filter").classList.remove("hidden");
    document.body.classList.add("noscroll");
  }

  // Close modal
  function closeFilterModal() {
    document.getElementById("js-filter").classList.add("hidden");
    document.body.classList.remove("noscroll");
  }

  // Buttons that open the modal
  document
    .getElementById("js-price-range")
    .addEventListener("click", openFilterModal);
  document
    .getElementById("js-guest-picker")
    .addEventListener("click", openFilterModal);
  document
    .getElementById("filter-btn")
    .addEventListener("click", openFilterModal);

  // Close on X button
  document
    .querySelector(".js-trigger-hide")
    .addEventListener("click", closeFilterModal);

  // Close on background layer click
  document
    .querySelector("#js-filter .popup-layer")
    .addEventListener("click", closeFilterModal);

  // Guest counter
  const guestCount = document.getElementById("js-guest-count");
  document.getElementById("js-guest-increase").addEventListener("click", () => {
    guestCount.textContent = parseInt(guestCount.textContent) + 1;
  });
  document.getElementById("js-guest-decrease").addEventListener("click", () => {
    const current = parseInt(guestCount.textContent);
    if (current > 0) guestCount.textContent = current - 1;
  });

  // Price slider logic
  const minSlider = document.getElementById("js-min-price-slider");
  const maxSlider = document.getElementById("js-max-price-slider");
  const minInput = document.getElementById("js-min-price");
  const maxInput = document.getElementById("js-max-price");
  const sliderRange = document.getElementById("js-slider-range");

  function syncPriceBounds(resetToLimits) {
    const bounds = getBounds();
    minSlider.min = bounds.min;
    minSlider.max = bounds.max;
    minSlider.step = bounds.step;
    maxSlider.min = bounds.min;
    maxSlider.max = bounds.max;
    maxSlider.step = bounds.step;
    minInput.min = bounds.min;
    minInput.max = bounds.max;
    maxInput.min = bounds.min;
    maxInput.max = bounds.max;

    if (resetToLimits) {
      minSlider.value = bounds.min;
      maxSlider.value = bounds.max;
      minInput.value = bounds.min;
      maxInput.value = bounds.max;
    }
  }

  function updateSliderRange() {
    const bounds = getBounds();
    const minVal = parseInt(minSlider.value);
    const maxVal = parseInt(maxSlider.value);
    const leftPct = ((minVal - bounds.min) / (bounds.max - bounds.min)) * 100;
    const widthPct = ((maxVal - minVal) / (bounds.max - bounds.min)) * 100;
    sliderRange.style.left = leftPct + "%";
    sliderRange.style.width = widthPct + "%";
  }

  minSlider.addEventListener("input", () => {
    let minVal = parseInt(minSlider.value);
    let maxVal = parseInt(maxSlider.value);
    const bounds = getBounds();
    if (minVal >= maxVal) {
      minVal = maxVal - bounds.step;
      minSlider.value = minVal;
    }
    minInput.value = minVal;
    updateSliderRange();
  });

  maxSlider.addEventListener("input", () => {
    let minVal = parseInt(minSlider.value);
    let maxVal = parseInt(maxSlider.value);
    const bounds = getBounds();
    if (maxVal <= minVal) {
      maxVal = minVal + bounds.step;
      maxSlider.value = maxVal;
    }
    maxInput.value = maxVal;
    updateSliderRange();
  });

  minInput.addEventListener("input", () => {
    const bounds = getBounds();
    let minVal = parseInt(minInput.value) || bounds.min;
    let maxVal = parseInt(maxInput.value);
    if (minVal < bounds.min) minVal = bounds.min;
    if (minVal >= maxVal) minVal = maxVal - bounds.step;
    minInput.value = minVal;
    minSlider.value = minVal;
    updateSliderRange();
  });

  maxInput.addEventListener("input", () => {
    const bounds = getBounds();
    let maxVal = parseInt(maxInput.value) || bounds.max;
    let minVal = parseInt(minInput.value);
    if (maxVal > bounds.max) maxVal = bounds.max;
    if (maxVal <= minVal) maxVal = minVal + bounds.step;
    maxInput.value = maxVal;
    maxSlider.value = maxVal;
    updateSliderRange();
  });

  // Initialize slider on page load
  updateCurrencyIcons();
  syncPriceBounds(true);
  updateSliderRange();

  // Clear button — resets everything
  document.getElementById("js-clear-filter").addEventListener("click", () => {
    document
      .querySelectorAll("#js-filter input[type=checkbox]")
      .forEach((cb) => (cb.checked = false));
    guestCount.textContent = "0";
    syncPriceBounds(true);
    updateSliderRange();

    window.checkin  = '';
    window.checkout = '';
    const modalDp = document.getElementById('modal-datepicker');
    if (modalDp) modalDp.value = '';
    const btn = document.getElementById('standalone-dp');
    if (btn) btn.textContent = 'Dates';
  });

  // Apply filter button
  document.getElementById("js-apply-filter").addEventListener("click", () => {
    const checkedAmenities = [];
    document
      .querySelectorAll(
        "#js-dynamic-amenities-filter input[type=checkbox]:checked",
      )
      .forEach((cb) => {
        checkedAmenities.push(cb.value);
      });
    if (document.getElementById("amenity-11").checked) {
      checkedAmenities.push("11");
    }

    const ecoFriendly = document.getElementById("js-eco-friendly").checked;
    const bounds = getBounds();
    const currency = getCurrency();
    const minPriceDisplay = parseInt(minInput.value) || bounds.min;
    const maxPriceDisplay = parseInt(maxInput.value) || bounds.max;
    const guests = guestCount.textContent;

    // Convert selected display currency to USD for API.
    const minPriceUSD = Math.round(minPriceDisplay / currency.rate);
    const maxPriceUSD = Math.round(maxPriceDisplay / currency.rate);

    const filters = {
      amenities: [...new Set(checkedAmenities)],
      ecoFriendly: ecoFriendly,
      amount:
        minPriceDisplay > bounds.min || maxPriceDisplay < bounds.max
          ? `${minPriceUSD}-${maxPriceUSD}`
          : "",
      amountDisplay:
        minPriceDisplay > bounds.min || maxPriceDisplay < bounds.max
          ? `${minPriceDisplay}-${maxPriceDisplay}`
          : "",
      selectedCurrency: currency.code,
      guests: guests !== "0" ? guests : "",
      checkin:     window.checkin  || '',
      checkout:    window.checkout || ''
    };

    console.log("Applying filters:", filters);
    closeFilterModal();

    if (window.loadProperties && window.currentCategory) {
      const sortEl = document.getElementById('sort-properties');
      const currentOrder = (sortEl && sortEl.value) ? sortEl.value : '1';
      window.loadProperties(window.currentCategory, currentOrder, filters);
    }
  });

  window.addEventListener("currency:changed", function () {
    updateCurrencyIcons();
    syncPriceBounds(true);
    updateSliderRange();
  });
});
