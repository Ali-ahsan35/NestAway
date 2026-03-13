document.addEventListener("DOMContentLoaded", function () {
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
  const MIN = 244;
  const MAX = 122096;
  const minSlider = document.getElementById("js-min-price-slider");
  const maxSlider = document.getElementById("js-max-price-slider");
  const minInput = document.getElementById("js-min-price");
  const maxInput = document.getElementById("js-max-price");
  const sliderRange = document.getElementById("js-slider-range");

  function updateSliderRange() {
    const minVal = parseInt(minSlider.value);
    const maxVal = parseInt(maxSlider.value);
    const leftPct = ((minVal - MIN) / (MAX - MIN)) * 100;
    const widthPct = ((maxVal - minVal) / (MAX - MIN)) * 100;
    sliderRange.style.left = leftPct + "%";
    sliderRange.style.width = widthPct + "%";
  }

  minSlider.addEventListener("input", () => {
    let minVal = parseInt(minSlider.value);
    let maxVal = parseInt(maxSlider.value);
    if (minVal >= maxVal) {
      minVal = maxVal - 100;
      minSlider.value = minVal;
    }
    minInput.value = minVal;
    updateSliderRange();
  });

  maxSlider.addEventListener("input", () => {
    let minVal = parseInt(minSlider.value);
    let maxVal = parseInt(maxSlider.value);
    if (maxVal <= minVal) {
      maxVal = minVal + 100;
      maxSlider.value = maxVal;
    }
    maxInput.value = maxVal;
    updateSliderRange();
  });

  minInput.addEventListener("input", () => {
    let minVal = parseInt(minInput.value) || MIN;
    let maxVal = parseInt(maxInput.value);
    if (minVal < MIN) minVal = MIN;
    if (minVal >= maxVal) minVal = maxVal - 100;
    minInput.value = minVal;
    minSlider.value = minVal;
    updateSliderRange();
  });

  maxInput.addEventListener("input", () => {
    let maxVal = parseInt(maxInput.value) || MAX;
    let minVal = parseInt(minInput.value);
    if (maxVal > MAX) maxVal = MAX;
    if (maxVal <= minVal) maxVal = minVal + 100;
    maxInput.value = maxVal;
    maxSlider.value = maxVal;
    updateSliderRange();
  });

  // Initialize slider on page load
  updateSliderRange();

  // Clear button — resets everything
  document.getElementById("js-clear-filter").addEventListener("click", () => {
    document
      .querySelectorAll("#js-filter input[type=checkbox]")
      .forEach((cb) => (cb.checked = false));
    guestCount.textContent = "0";
    minSlider.value = MIN;
    maxSlider.value = MAX;
    minInput.value = MIN;
    maxInput.value = MAX;
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
    const minPriceBDT = parseInt(minInput.value) || MIN;
    const maxPriceBDT = parseInt(maxInput.value) || MAX;
    const guests = guestCount.textContent;

    // Convert BDT to USD for API
    const minPriceUSD = Math.round(minPriceBDT / 120);
    const maxPriceUSD = Math.round(maxPriceBDT / 120);

    const filters = {
      amenities: [...new Set(checkedAmenities)],
      ecoFriendly: ecoFriendly,
      amount:
        minPriceBDT > MIN || maxPriceBDT < MAX
          ? `${minPriceUSD}-${maxPriceUSD}`
          : "",
      amountBDT:
        minPriceBDT > MIN || maxPriceBDT < MAX
          ? `${minPriceBDT}-${maxPriceBDT}`
          : "",
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
});
