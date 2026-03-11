document.addEventListener("DOMContentLoaded", function () {
  const keyword = window.searchKeyword || "Barcelona, Spain";
  let currentCategory = "";

  // Step 1: breadcrumb — runs once on page load
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
        bc.innerHTML += `<span>${item.Name}</span>`;
        if (i < breadcrumbs.length - 1)
          bc.innerHTML += `<span class="sep"> › </span>`;
      });

      const locationName = breadcrumbData?.GeoInfo?.ShortName || "Barcelona";
      document.getElementById("page-title").textContent =
        "Find a Place to Stay in " + locationName;
      document.getElementById("page-subtitle").textContent =
        "Find Your Perfect Stay in " + locationName;

      // Save category globally so sort script can access it
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
      if (urlParams.get('guests')) {
          savedFilters.guests = parseInt(urlParams.get('pax'));
      }

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

    // count.textContent = items.length + " properties found";
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

      // Build breadcrumb location list
      const display = geo?.Display || [];
      const locationItems = Array.isArray(display)
        ? display
            .map((d) => `<li><span class="pt-tile-bdc">${d}</span></li>`)
            .join("")
        : `<li><span class="pt-tile-bdc">${display}</span></li>`;

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

      card.innerHTML = `
                <!-- Image Section -->
                <div class="image-section relative" id="js-${id}-image-section">
                    <div class="tiles-icons absolute">
                        <div class="tiles-btn fav-icon heart-btn" data-id="${id}" title="Bookmark">♡</div>
                    </div>

                    <a href="${partnerUrl}" target="_blank" class="sp-property-image js-tiles-redirect">
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
                            <a title="${name}" href="${partnerUrl}" target="_blank" class="pt-property-title refine-page-redirect">
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
                            <a rel="nofollow" class="pt-logo-wrap" href="${partnerUrl}" target="_blank">
                                ${
                                  isExpedia
                                    ? `<img src="/static/img/partners-logo/expedia_v2.svg" height="14" width="80" alt="Expedia" class="pt-partner-logo">`
                                    : `<img src="/static/img/partners-logo/booking.svg" height="14" width="80" alt="Booking.com" class="pt-partner-logo">`
                                }
                            </a>
                        </div>
                        <a href="${partnerUrl}" rel="nofollow" target="_blank" class="availability-button pt-availability">
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
