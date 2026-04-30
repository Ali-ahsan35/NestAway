document.addEventListener("DOMContentLoaded", function () {
  const keyword = window.searchKeyword || "Barcelona, Spain";
  let currentCategory = "";
  const reservedKeys = new Set([
    "search",
    "order",
    "amenities",
    "ecoFriendly",
    "amount",
    "selectedCurrency",
    "pax",
    "dateStart",
    "dateEnd"
  ]);

  function getExtraParamsFromUrl() {
    const extras = {};
    const urlParams = new URLSearchParams(window.location.search);
    urlParams.forEach((value, key) => {
      if (!reservedKeys.has(key)) {
        extras[key] = value;
      }
    });
    return extras;
  }

  const initialExtraParams = getExtraParamsFromUrl();

  function currentCurrency() {
    if (window.CurrencyManager && window.CurrencyManager.getCurrentCurrency) {
      return window.CurrencyManager.getCurrentCurrency();
    }
    return { code: "USD", symbol: "US $", rate: 1 };
  }

  function displayBounds() {
    const rate = Number(currentCurrency().rate) || 1;
    return {
      min: Math.round(2 * rate),
      max: Math.round(1000 * rate),
    };
  }

  function parseAmountToUSD(amountText, codeFromUrl) {
    if (!amountText) return "";
    const parts = amountText.split("-");
    if (parts.length !== 2) return "";
    const minDisplay = parseInt(parts[0], 10);
    const maxDisplay = parseInt(parts[1], 10);
    if (!Number.isFinite(minDisplay) || !Number.isFinite(maxDisplay)) return "";

    let rate = Number(currentCurrency().rate) || 1;
    if (window.CurrencyManager && codeFromUrl) {
      const byCode = window.CurrencyManager.getCurrencyByCode(codeFromUrl);
      if (byCode) rate = Number(byCode.rate) || 1;
    }

    const minUSD = Math.round(minDisplay / rate);
    const maxUSD = Math.round(maxDisplay / rate);
    return minUSD + "-" + maxUSD;
  }

  // Step 1: breadcrumb
  fetch("/api/breadcrumb?keyword=" + encodeURIComponent(keyword), {
    headers: {
      "Content-Type": "application/json",
      "X-Requested-With": "XMLHttpRequest",
    },
  })
    .then((res) => res.json())
    .then((breadcrumbData) => {
      const breadcrumbs = breadcrumbData?.GeoInfo?.Breadcrumbs || [];
      const bc = document.getElementById("breadcrumb");
      bc.innerHTML =
        "Vacation Rentals in " +
        (breadcrumbs[breadcrumbs.length - 1]?.Name || "") +
        " &nbsp;|&nbsp; ";
      breadcrumbs.forEach((item, i) => {
        const slug = item.Display.join("/");
        bc.innerHTML += `<a href="/all/${slug}" style="color:inherit; text-decoration:none;">${item.Name}</a>`;
        if (i < breadcrumbs.length - 1)
          bc.innerHTML += `<span class="sep"> › </span>`;
      });

      const locationName = breadcrumbData?.GeoInfo?.ShortName || "Barcelona";
      document.getElementById("page-title").textContent =
        "Find a Place to Stay in " + locationName;
      document.getElementById("page-subtitle").textContent =
        "Find Your Perfect Stay in " + locationName;

      currentCategory = breadcrumbData?.GeoInfo?.LocationSlug;
      window.currentCategory = currentCategory;

      const urlParams = new URLSearchParams(window.location.search);
      const defaultOrder = parseInt(urlParams.get("order")) || 1;

      const savedFilters = {};
      window.extraParams = getExtraParamsFromUrl();

      if (urlParams.get("amenities")) {
        savedFilters.amenities = urlParams.get("amenities").split("-").map(Number);
      }
      if (urlParams.get("ecoFriendly")) {
        savedFilters.ecoFriendly = true;
      }
      if (urlParams.get("amount")) {
        savedFilters.amountDisplay = urlParams.get("amount");
        savedFilters.selectedCurrency =
          urlParams.get("selectedCurrency") || currentCurrency().code;
        savedFilters.amount = parseAmountToUSD(
          savedFilters.amountDisplay,
          savedFilters.selectedCurrency
        );
      }
      if (urlParams.get("pax")) {
        savedFilters.guests = parseInt(urlParams.get("pax"));
      }
      if (urlParams.get("dateStart")) {
        savedFilters.checkin = urlParams.get("dateStart");
        window.checkin = savedFilters.checkin;
      }
      if (urlParams.get("dateEnd")) {
        savedFilters.checkout = urlParams.get("dateEnd");
        window.checkout = savedFilters.checkout;
      }

      // Restore modal state from URL
      if (savedFilters.amenities) {
        savedFilters.amenities.forEach((id) => {
          const cb = document.getElementById("amenity-" + id);
          if (cb) cb.checked = true;
        });
      }
      if (savedFilters.ecoFriendly) {
        const eco = document.getElementById("js-eco-friendly");
        if (eco) eco.checked = true;
      }
      if (urlParams.get("pax")) {
        const guestCount = document.getElementById("js-guest-count");
        if (guestCount) guestCount.textContent = urlParams.get("pax");
      }
      if (urlParams.get("amount")) {
        const parts = urlParams.get("amount").split("-");
        const minDisplay = parseInt(parts[0], 10);
        const maxDisplay = parseInt(parts[1], 10);
        const minSlider = document.getElementById("js-min-price-slider");
        const maxSlider = document.getElementById("js-max-price-slider");
        const minInput = document.getElementById("js-min-price");
        const maxInput = document.getElementById("js-max-price");
        if (minSlider) minSlider.value = minDisplay;
        if (maxSlider) maxSlider.value = maxDisplay;
        if (minInput) minInput.value = minDisplay;
        if (maxInput) maxInput.value = maxDisplay;
        const sliderRange = document.getElementById("js-slider-range");
        if (sliderRange) {
          const bounds = displayBounds();
          const leftPct =
            ((minDisplay - bounds.min) / (bounds.max - bounds.min)) * 100;
          const widthPct =
            ((maxDisplay - minDisplay) / (bounds.max - bounds.min)) * 100;
          sliderRange.style.left = leftPct + "%";
          sliderRange.style.width = widthPct + "%";
        }
      }

      // Restore date button text
      if (window.checkin && window.checkout) {
        setTimeout(function () {
          const btn = document.getElementById("standalone-dp");
          if (btn) btn.textContent = window.checkin + " → " + window.checkout;
          const modalDp = document.getElementById("modal-datepicker");
          if (modalDp) modalDp.value = window.checkin + " - " + window.checkout;
        }, 500);
      }

      window.loadProperties(currentCategory, defaultOrder, savedFilters);
    })
    .catch((err) => {
      console.error("Breadcrumb error:", err);
    });

  function updateURL(order, filters = {}) {
    const params = new URLSearchParams(window.location.search);

    params.set("search", window.searchKeyword);
    params.set("order", order);

    if (filters.amenities && filters.amenities.length > 0) {
      params.set("amenities", filters.amenities.join("-"));
    }
    if (filters.ecoFriendly) {
      params.set("ecoFriendly", "true");
    }
    if (filters.amount) {
      params.set("amount", filters.amountDisplay || filters.amount);
      params.set(
        "selectedCurrency",
        filters.selectedCurrency || currentCurrency().code
      );
    }
    if (filters.guests && filters.guests > 0) {
      params.set("pax", filters.guests);
    }
    if (filters.checkin) {
      params.set("dateStart", filters.checkin);
    }
    if (filters.checkout) {
      params.set("dateEnd", filters.checkout);
    }

    // Preserve extra params like pt, isWinter etc
    const extras =
      Object.keys(window.extraParams || {}).length > 0
        ? window.extraParams
        : initialExtraParams;
    Object.keys(extras).forEach((key) => {
      params.set(key, extras[key]);
    });

    window.history.pushState({}, "", "/refine?" + params.toString());
  }

  window.loadProperties = function (category, order, filters = {}) {
    updateURL(order, filters);
    showLoading();

    let url =
      "/api/properties?category=" +
      encodeURIComponent(category) +
      "&order=" +
      order;

    if (filters.amenities && filters.amenities.length > 0) {
      url += "&amenities=" + filters.amenities.join("-");
    }
    if (filters.ecoFriendly) {
      url += "&ecoFriendly=true";
    }
    if (filters.amount) {
      url +=
        "&amount=" +
        filters.amount +
        "&selectedCurrency=" +
        encodeURIComponent(
          filters.selectedCurrency || currentCurrency().code
        );
    }
    if (filters.guests && filters.guests > 0) {
      url += "&pax=" + filters.guests;
    }
    if (filters.checkin && filters.checkout) {
      url += "&dateStart=" + filters.checkin + "&dateEnd=" + filters.checkout;
    }

    // Add extra params like pt, isWinter etc
    const extras =
      Object.keys(window.extraParams || {}).length > 0
        ? window.extraParams
        : initialExtraParams;
    Object.keys(extras).forEach((key) => {
      url +=
        "&" +
        encodeURIComponent(key) +
        "=" +
        encodeURIComponent(extras[key]);
    });

    console.log("Fetching URL:", url);

    fetch(url, {
      headers: {
        "Content-Type": "application/json",
        "X-Requested-With": "XMLHttpRequest",
      },
    })
      .then((res) => res.json())
      .then((propertiesData) => {
        const ids = propertiesData?.Result?.ItemIDs || [];
        console.log("Items count:", ids.length, "order:", order, "filters:", filters);

        const limitedIds = ids.slice(0, 72);
        const idString = limitedIds.join(",");

        return fetch(
          "/api/propertydetails?ids=" + encodeURIComponent(idString),
          {
            headers: {
              "Content-Type": "application/json",
              "X-Requested-With": "XMLHttpRequest",
            },
          }
        );
      })
      .then((res) => res.json())
      .then((detailsData) => {
        const items = detailsData?.Items || [];
        renderCards(items);
      })
      .catch((err) => {
        console.error("Load error:", err);
        document.getElementById("grid").innerHTML =
          "<div class='loading'>Failed to load properties.</div>";
      });
  };

  function showLoading() {
    document.getElementById("grid").innerHTML = `
      <div class="loading">
        <span class="loading-dot">●</span>
        <span class="loading-dot">●</span>
        <span class="loading-dot">●</span>
        &nbsp; Loading properties...
      </div>`;
    document.getElementById("result-count").textContent = "";
  }

  function renderCards(items) {
    const grid = document.getElementById("grid");

    grid.innerHTML = "";

    if (items.length === 0) {
      grid.innerHTML = "<div class='loading'>No properties found.</div>";
      return;
    }

    const fallbackImages = [
      "demo_img_01.jpg",
      "demo_img_02.jpg",
      "demo_img_03.jpg",
      "demo_img_04.jpg",
    ];

    items.forEach((item, index) => {
      const p = item.Property;
      const geo = item.GeoInfo;
      const partner = item.Partner;
      const id = item.ID || "";
      const name = p?.PropertyName || "Unnamed Property";
      const type = p?.PropertyType || "";
      const priceUsd = Number(p?.Price);
      const hasPrice = Number.isFinite(priceUsd) && priceUsd > 0;
      const rating = p?.ReviewScore ? p.ReviewScore.toFixed(1) : null;
      const reviews = p?.Counts?.Reviews || 0;
      const imgName = p?.FeatureImage || "";
      const imgUrl = imgName
        ? "https://imgservice.ownerdirect.com/600x600/" + imgName
        : null;
      const fallbackUrl =
        "/static/img/reserve_image/" +
        fallbackImages[Math.floor(Math.random() * fallbackImages.length)];
      const partnerUrl = partner?.URL || "#";
      const isExpedia = partnerUrl.includes("expedia");
      const isVrbo = partnerUrl.includes("vrbo");
      const amenities = p?.Amenities
        ? Object.values(p.Amenities).slice(0, 3)
        : [];

      const categories = geo?.Categories || [];
      const locationItems =
        categories.length > 0
          ? categories
              .map(
                (cat) =>
                  `<li><a href="/all/${cat.Slug}" class="pt-tile-bdc" style="color:inherit; text-decoration:none;">${cat.Name}</a></li>`
              )
              .join("")
          : `<li><span class="pt-tile-bdc">${geo?.Display || ""}</span></li>`;

      const amenityItems = amenities
        .map((a) => `<li class="pt-am-item">${a}</li>`)
        .join("");

      const stars = p?.StarRating || 0;
      const starClass = stars > 0 ? `ratings star-icons-${stars}` : "ratings";

      const card = document.createElement("div");
      card.setAttribute("data-property_id", id);
      card.className = "sp-property-card";
      card.style.animationDelay = index * 0.04 + "s";
      card.setAttribute("data-feed", item.Feed);
      card.setAttribute("data-published", item.Published);
      card.setAttribute("data-upat", p?.UpdatedAt || "");
      card.setAttribute("data-lat", geo?.Lat || "");
      card.setAttribute("data-lng", geo?.Lng || "");
      card.setAttribute("data-type", p?.PropertyType || "");
      card.setAttribute("data-dest_id", partner?.ID || "");
      card.setAttribute("data-owner_id", partner?.OwnerID || "");
      card.setAttribute("data-direct_url", partner?.URL || "");
      card.setAttribute("data-display", geo?.Display || "");
      card.setAttribute("data-city", geo?.City || "");
      card.setAttribute("data-country", geo?.Country || "");
      card.setAttribute("data-country_code", geo?.CountryCode || "");
      card.setAttribute("data-epc", partner?.EpCluster || "");
      card.setAttribute("data-eplid", geo?.LocationID || "");
      card.setAttribute("data-index", index);
      if (!imgName) {
        card.setAttribute("data-missing-feature-image", "true");
      }

      // Partner logo
      let partnerLogo = `<img src="/static/img/partners-logo/booking.svg" height="14" width="80" alt="Booking.com" class="pt-partner-logo">`;
      if (isExpedia) {
        partnerLogo = `<img src="/static/img/partners-logo/expedia_v2.svg" height="14" width="80" alt="Expedia" class="pt-partner-logo">`;
      } else if (isVrbo) {
        partnerLogo = `<img src="/static/img/partners-logo/vrbo.svg" height="14" width="80" alt="VRBO" class="pt-partner-logo">`;
      }

      card.innerHTML = `
        <div class="image-section relative" id="js-${id}-image-section">
          <div class="tiles-icons absolute">
            <div class="tiles-btn fav-icon heart-btn" data-id="${id}" title="Bookmark" onclick="toggleFavourite(this)">
              <svg class="heart-icon" xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="white" stroke-width="2">
                <path d="M20.84 4.61a5.5 5.5 0 0 0-7.78 0L12 5.67l-1.06-1.06a5.5 5.5 0 0 0-7.78 7.78l1.06 1.06L12 21.23l7.78-7.78 1.06-1.06a5.5 5.5 0 0 0 0-7.78z"></path>
              </svg>
            </div>
          </div>
          <a href="#" target="_blank" class="sp-property-image js-tiles-redirect"
            onmouseenter="buildRedirectUrl(this)"
            onclick="redirectToPartner(this); return false;">
            <img
              class="featured-image pt-featured-image"
              src="${imgUrl || fallbackUrl}"
              alt="${name}"
              onerror="this.onerror=null; const card=this.closest('.sp-property-card'); if(card){card.dataset.missingFeatureImage='true';} this.src='${fallbackUrl}';"
            >
            ${hasPrice
              ? `<span class="property-price js-price-value"
                  data-base-usd="${priceUsd}"
                  data-price-prefix="From "
                  data-price-suffix="">
                  From US $ ${Math.round(priceUsd).toLocaleString()}
                </span>`
              : ""}
          </a>
        </div>

        <div class="sp-property-details js-tiles-redirect">
          <div class="pt-content-wrap">
            <div class="property-review pt-property-review">
              <div class="rating-review pt-rating-review">
                <div class="${starClass}"></div>
                ${rating ? `<span class="divider"> | </span>` : ""}
                ${rating
                  ? `<div class="reviews pt-reviews">
                      <span class="text-bold review-general">${rating}</span>
                      <span class="number-of-review">(${reviews} Reviews)</span>
                    </div>`
                  : ""}
              </div>
              <span class="property-type">${type}</span>
            </div>

            <div class="property-title">
              <a title="${name}"
                href="https://ownerdirect.beta.123presto.com/property/${p?.PropertySlug}/${id}"
                target="_blank"
                class="pt-property-title refine-page-redirect">
                ${name}
              </a>
            </div>

            <div class="property-info-wrap">
              <div class="property-info">
                <ul class="ellipsis pt-amenities">${amenityItems}</ul>
              </div>
              <div class="property-location">
                <ul class="ellipsis pt-breadcrumbs">${locationItems}</ul>
              </div>
            </div>
          </div>

          <div class="property-bottom">
            <div class="property-brand">
              <a rel="nofollow" class="pt-logo-wrap" href="#" target="_blank"
                onmouseenter="buildRedirectUrl(this)"
                onclick="redirectToPartner(this); return false;">
                ${partnerLogo}
              </a>
            </div>
            <a href="#" rel="nofollow" target="_blank"
              class="availability-button pt-availability"
              onmouseenter="buildRedirectUrl(this)"
              onclick="redirectToPartner(this); return false;">
              View Availability
            </a>
            ${hasPrice
              ? `<span class="list-tile-price property-price js-price-value"
                  data-base-usd="${priceUsd}"
                  data-price-prefix="From "
                  data-price-suffix=" / night">
                  From US $ ${Math.round(priceUsd).toLocaleString()} / night
                </span>`
              : ""}
          </div>
        </div>
      `;

      grid.appendChild(card);
    });

    if (typeof initImageSlider === "function") {
      initImageSlider();
    }
    if (window.CurrencyManager && window.CurrencyManager.applyCurrencyToDom) {
      window.CurrencyManager.applyCurrencyToDom(grid);
    }

    // Restore favourites
    let favourites = JSON.parse(localStorage.getItem("favourite_list") || "{}");
    Object.keys(favourites).forEach((favId) => {
      const btn = document.querySelector(`.heart-btn[data-id="${favId}"]`);
      if (btn) {
        const icon = btn.querySelector(".heart-icon");
        if (icon) {
          icon.setAttribute("fill", "red");
          icon.setAttribute("stroke", "red");
        }
      }
    });
  }

  window.addEventListener("currency:changed", function () {
    const grid = document.getElementById("grid");
    if (
      window.CurrencyManager &&
      window.CurrencyManager.applyCurrencyToDom &&
      grid
    ) {
      window.CurrencyManager.applyCurrencyToDom(grid);
    }
  });
});