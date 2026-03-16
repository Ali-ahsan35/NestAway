document.addEventListener("DOMContentLoaded", function () {
  const keyword = window.searchKeyword || "Barcelona, Spain";
  let currentCategory = "";

  // breadcrumb
  fetch("/api/breadcrumb?keyword=" + encodeURIComponent(keyword), {
    headers: {
      "Content-Type": "application/json",
      "X-Requested-With": "XMLHttpRequest",
    },
  })
    .then((res) => res.json())
    .then((breadcrumbData) => {
      // Render breadcrumb
      const breadcrumbs = breadcrumbData?.GeoInfo?.Breadcrumbs || [];
      const bc = document.getElementById("breadcrumb");
      bc.innerHTML =
        "Vacation Rentals in " +
        (breadcrumbs[breadcrumbs.length - 1]?.Name || "") +
        " &nbsp;|&nbsp; ";
      breadcrumbs.forEach((item, i) => {
          // Build slug from Display array joined with "/"
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

      // Save category globally
      currentCategory = breadcrumbData?.GeoInfo?.LocationSlug;
      window.currentCategory = currentCategory;

      // Read filters from URL on page load
      const urlParams = new URLSearchParams(window.location.search);
      const defaultOrder = parseInt(urlParams.get('order')) || 1;

      const savedFilters = {};
      if (urlParams.get('amenities')) {
          savedFilters.amenities = urlParams.get('amenities').split('-').map(Number);
      }
      if (urlParams.get('ecoFriendly')) {
          savedFilters.ecoFriendly = true;
      }
      if (urlParams.get('amount')) {
          savedFilters.amount = urlParams.get('amount');
      }
      if (urlParams.get('pax')) {
          savedFilters.guests = parseInt(urlParams.get('pax'));
      }
      if (urlParams.get('dateStart')) {
          savedFilters.checkin = urlParams.get('dateStart');
          window.checkin = savedFilters.checkin;
      }
      if (urlParams.get('dateEnd')) {
          savedFilters.checkout = urlParams.get('dateEnd');
          window.checkout = savedFilters.checkout;
      }

      if (savedFilters.amenities) {
          savedFilters.amenities.forEach(id => {
              const cb = document.getElementById('amenity-' + id);
              if (cb) cb.checked = true;
          });
      }
      if (savedFilters.ecoFriendly) {
          const eco = document.getElementById('js-eco-friendly');
          if (eco) eco.checked = true;
      }
      if (urlParams.get('pax')) {
          const guestCount = document.getElementById('js-guest-count');
          if (guestCount) guestCount.textContent = urlParams.get('pax');
      }
      if (urlParams.get('amount')) {
          const parts = urlParams.get('amount').split('-');
          const minBDT = parseInt(parts[0]);
          const maxBDT = parseInt(parts[1]);
          const minSlider = document.getElementById('js-min-price-slider');
          const maxSlider = document.getElementById('js-max-price-slider');
          const minInput  = document.getElementById('js-min-price');
          const maxInput  = document.getElementById('js-max-price');
          if (minSlider) minSlider.value = minBDT;
          if (maxSlider) maxSlider.value = maxBDT;
          if (minInput)  minInput.value  = minBDT;
          if (maxInput)  maxInput.value  = maxBDT;
          const sliderRange = document.getElementById('js-slider-range');
          if (sliderRange) {
              const MIN = 244, MAX = 122096;
              const leftPct  = ((minBDT - MIN) / (MAX - MIN)) * 100;
              const widthPct = ((maxBDT - minBDT) / (MAX - MIN)) * 100;
              sliderRange.style.left  = leftPct + '%';
              sliderRange.style.width = widthPct + '%';
          }
      }
      console.log("amount from URL:", urlParams.get('amount'));
      console.log("minSlider found:", document.getElementById('js-min-price-slider'));
      console.log("guestCount found:", document.getElementById('js-guest-count'));
      console.log("amenity-1 found:", document.getElementById('amenity-1'));

      window.loadProperties(currentCategory, defaultOrder, savedFilters);
    })
    .catch((err) => {
      console.error("Breadcrumb error:", err);
    });


  function updateURL(order, filters = {}) {
    const params = new URLSearchParams();
    
    // Always include search and order
    params.set('search', window.searchKeyword);
    params.set('order', order);
    
    // Add filters if set
    if (filters.amenities && filters.amenities.length > 0) {
        params.set('amenities', filters.amenities.join("-"));
    }
    if (filters.ecoFriendly) {
        params.set('ecoFriendly', 'true');
    }
    if (filters.amount) {
    // Use BDT amount in URL if available, otherwise use amount as-is
    params.set('amount', filters.amountBDT || filters.amount);
    params.set('selectedCurrency', 'BDT');
    }
    if (filters.guests && filters.guests > 0) {
        params.set('pax', filters.guests);
    }
    if (filters.checkin) {
        params.set('dateStart', filters.checkin);
    }
    if (filters.checkout) {
        params.set('dateEnd', filters.checkout);
    }

    // Update URL without reloading page
    window.history.pushState({}, '', '/refine?' + params.toString());
}
  // Expose to window so the sort script in refine.tpl can call it
  window.loadProperties = function (category, order, filters = {}) {
    updateURL(order, filters);
    showLoading();

    let url =
      "/api/properties?category=" +
      encodeURIComponent(category) +
      "&order=" +
      order;

    // Add filters to URL if provided
    if (filters.amenities && filters.amenities.length > 0) {
      url += "&amenities=" + filters.amenities.join("-");
    }
    if (filters.ecoFriendly) {
      url += "&ecoFriendly=true";
    }
    if (filters.amount) {
        url += "&amount=" + filters.amount + "&selectedCurrency=BDT";
    }
    if (filters.guests && filters.guests > 0) {
      url += "&pax=" + filters.guests;
    }
    if (filters.checkin && filters.checkout) {
        url += "&dateStart=" + filters.checkin + "&dateEnd=" + filters.checkout;
        console.log("Added dates to URL:", filters.checkin, filters.checkout); 
    }

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
        console.log(
          "Items count:",
          ids.length,
          "order:",
          order,
          "filters:",
          filters,
        );

        // Limit to first 72 IDs to avoid API limit
        const limitedIds = ids.slice(0, 72);
        const idString = limitedIds.join(",");
        return fetch(
          "/api/propertydetails?ids=" + encodeURIComponent(idString),
          {
            headers: {
              "Content-Type": "application/json",
              "X-Requested-With": "XMLHttpRequest",
            },
          },
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
    const count = document.getElementById("result-count");

    grid.innerHTML = "";

    if (items.length === 0) {
      grid.innerHTML = "<div class='loading'>No properties found.</div>";
      return;
    }

    items.forEach((item, index) => {
      const p = item.Property;
      const geo = item.GeoInfo;
      const partner = item.Partner;
      const id = item.ID || "";
      const name = p?.PropertyName || "Unnamed Property";
      const type = p?.PropertyType || "";
      const price = p?.Price
        ? Math.round(p.Price * 120).toLocaleString()
        : null;
      const rating = p?.ReviewScore ? p.ReviewScore.toFixed(1) : null;
      const reviews = p?.Counts?.Reviews || 0;
      const imgName = p?.FeatureImage || "";
      const imgUrl = imgName
        ? "https://imgservice.smartours.com/600x600/" + imgName
        : null;
      const partnerUrl = partner?.URL || "#";
      const isExpedia = partnerUrl.includes("expedia");
      const amenities = p?.Amenities
        ? Object.values(p.Amenities).slice(0, 3)
        : [];

      // Build breadcrumb location list from Categories
      const categories = geo?.Categories || [];
      const locationItems = categories.length > 0
          ? categories.map((cat) => {
              return `<li><a href="/all/${cat.Slug}" class="pt-tile-bdc" style="color:inherit; text-decoration:none;">${cat.Name}</a></li>`;
          }).join("")
          : `<li><span class="pt-tile-bdc">${geo?.Display || ""}</span></li>`;

      // Build amenities list
      const amenityItems = amenities
        .map((a) => `<li class="pt-am-item">${a}</li>`)
        .join("");

      // Star rating class (1-5)
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

      card.innerHTML = `
                <!-- Image Section -->
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
                        ${
                          imgUrl
                            ? `<img class="featured-image pt-featured-image" src="${imgUrl}" alt="${name}" onerror="this.src=''">`
                            : `<div class="featured-image pt-featured-image" style="background:#e8eef5;display:flex;align-items:center;justify-content:center;font-size:36px;">🏠</div>`
                        }
                        ${
                          price
                            ? `<span class="property-price js-price-value">From BD ৳ ${price}</span>`
                            : ""
                        }
                    </a>
                </div>

                <!-- Details Section -->
                <div class="sp-property-details js-tiles-redirect">
                    <div class="pt-content-wrap">

                        <!-- Rating row -->
                        <div class="property-review pt-property-review">
                            <div class="rating-review pt-rating-review">
                                <div class="${starClass}"></div>
                                ${rating ? `<span class="divider"> | </span>` : ""}
                                ${
                                  rating
                                    ? `
                                <div class="reviews pt-reviews">
                                    <span class="text-bold review-general">${rating}</span>
                                    <span class="number-of-review">(${reviews} Reviews)</span>
                                </div>`
                                    : ""
                                }
                            </div>
                            <span class="property-type">${type}</span>
                        </div>

                        <!-- Property name -->
                        <div class="property-title">
                            <a title="${name}" href="https://presto:TRAV3LA1@ownerdirect.beta.123presto.com/property/${p?.PropertySlug}/${id}" target="_blank" class="pt-property-title refine-page-redirect">
                                ${name}
                            </a>
                        </div>

                        <!-- Amenities + Location -->
                        <div class="property-info-wrap">
                            <div class="property-info">
                                <ul class="ellipsis pt-amenities">
                                    ${amenityItems}
                                </ul>
                            </div>
                            <div class="property-location">
                                <ul class="ellipsis pt-breadcrumbs">
                                    ${locationItems}
                                </ul>
                            </div>
                        </div>

                    </div>

                    <!-- Bottom: logo + price + button -->
                    <div class="property-bottom">
                        <div class="property-brand">
                            <a rel="nofollow" class="pt-logo-wrap" href="#" target="_blank"
                              onmouseenter="buildRedirectUrl(this)"
                              onclick="redirectToPartner(this); return false;">
                                ${
                                  isExpedia
                                    ? `<img src="/static/img/partners-logo/expedia_v2.svg" height="14" width="80" alt="Expedia" class="pt-partner-logo">`
                                    : `<img src="/static/img/partners-logo/booking.svg" height="14" width="80" alt="Booking.com" class="pt-partner-logo">`
                                }
                            </a>
                        </div>
                        <a href="#" rel="nofollow" target="_blank" 
                            class="availability-button pt-availability"
                            onmouseenter="buildRedirectUrl(this)"
                            onclick="redirectToPartner(this); return false;">
                            View Availability
                        </a>
                        ${
                          price
                            ? `
                        <span class="list-tile-price property-price js-price-value">
                            <span class="pt-from">From</span>
                            BD ৳ ${price}
                            <span class="pt-per-night">/ night</span>
                        </span>`
                            : ""
                        }
                    </div>
                </div>
            `;

      grid.appendChild(card);
    });
  }


});
