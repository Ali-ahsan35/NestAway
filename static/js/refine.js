document.addEventListener("DOMContentLoaded", function () {
    const keyword = "Barcelona, Spain";
    let locationName = "";

    // Step 1: breadcrumb
    fetch("/api/breadcrumb?keyword=" + encodeURIComponent(keyword), {
        headers: { "Content-Type": "application/json", "X-Requested-With": "XMLHttpRequest" }
    })
    .then(res => res.json())
    .then(breadcrumbData => {
        console.log("Step 1 done:", breadcrumbData?.GeoInfo?.ShortName);

        // Render breadcrumb
        const breadcrumbs = breadcrumbData?.GeoInfo?.Breadcrumbs || [];
        const bc = document.getElementById("breadcrumb");
        bc.innerHTML = "Vacation Rentals in " + (breadcrumbs[breadcrumbs.length-1]?.Name || "") + " &nbsp;|&nbsp; ";
        breadcrumbs.forEach((item, i) => {
            bc.innerHTML += `<span>${item.Name}</span>`;
            if (i < breadcrumbs.length - 1) bc.innerHTML += `<span class="sep"> › </span>`;
        });

        locationName = breadcrumbData?.GeoInfo?.ShortName || "Barcelona";
        document.getElementById("page-title").textContent = "Find a Place to Stay in " + locationName;
        document.getElementById("page-subtitle").textContent = "Find Your Perfect Stay in " + locationName;

        // Step 2: get properties
        const category = breadcrumbData?.GeoInfo?.LocationSlug;
        console.log("Step 2 category:", category);

        return fetch("/api/properties?category=" + encodeURIComponent(category), {
            headers: { "Content-Type": "application/json", "X-Requested-With": "XMLHttpRequest" }
        });
    })
    .then(res => res.json())
    .then(propertiesData => {
        console.log("Step 2 done, keys:", Object.keys(propertiesData));

        const ids = propertiesData?.Result?.ItemIDs || [];
        console.log("Items count:", ids.length);

        const idString = ids.join(",");
        console.log("IDs:", idString.substring(0, 100));

        // Step 3: get property details
        return fetch("/api/propertydetails?ids=" + encodeURIComponent(idString), {
            headers: { "Content-Type": "application/json", "X-Requested-With": "XMLHttpRequest" }
        });
    })
    .then(res => res.json())
    .then(detailsData => {
        console.log("Step 3 done, keys:", Object.keys(detailsData));
        const items = detailsData?.Items || [];
        console.log("Details items count:", items.length);
        renderCards(items);
    })
    .catch(err => {
        console.error("Error:", err);
        document.getElementById("grid").innerHTML = "<div class='loading'>Failed to load properties.</div>";
    });

    function renderCards(items) {
        const grid = document.getElementById("grid");
        const count = document.getElementById("result-count");

        count.textContent = items.length + " properties found";
        grid.innerHTML = "";

        if (items.length === 0) {
            grid.innerHTML = "<div class='loading'>No properties found.</div>";
            return;
        }

        items.forEach((item, index) => {
            const p = item.Property;
            const geo = item.GeoInfo;
            const partner = item.Partner;

            const name = p?.PropertyName || "Unnamed Property";
            const type = p?.PropertyType || "";
            const price = p?.Price ? "From BD ৳ " + Math.round(p.Price * 120).toLocaleString() : "Price on request";
            const rating = p?.ReviewScore ? p.ReviewScore.toFixed(1) : null;
            const reviews = p?.Counts?.Reviews || 0;
            const location = geo?.Display || "";
            const imgName = p?.FeatureImage || "";
            const imgUrl = imgName
                ? "https://imgservice.smartours.com/600x600/" + imgName
                : null;
            const source = partner?.URL?.includes("booking.com") ? "Booking.com" : "Expedia";

            const card = document.createElement("div");
            card.className = "card";
            card.style.animationDelay = (index * 0.05) + "s";

            card.innerHTML = `
                ${imgUrl
                    ? `<img class="card-img" src="${imgUrl}" alt="${name}" onerror="this.style.display='none';this.nextElementSibling.style.display='flex'">`
                    : ""}
                <div class="card-img-placeholder" style="display:${imgUrl ? 'none' : 'flex'}">🏠</div>
                <div class="card-body">
                    <div class="card-top">
                        <div class="card-rating">
                            ${rating ? `<span class="rating-badge">${rating}</span><span>(${reviews} Reviews)</span>` : '<span style="color:#ccc">No reviews</span>'}
                        </div>
                        <span class="card-type">${type}</span>
                    </div>
                    <div class="card-name">${name}</div>
                    <div class="card-location">${location}</div>
                    <div class="card-price">${price} <span>/ night</span></div>
                    <div class="card-footer">
                        <span class="card-source">${source}</span>
                        <button class="card-btn" onclick="window.open('${partner?.URL || '#'}', '_blank')">View Availability</button>
                    </div>
                </div>
            `;

            grid.appendChild(card);
        });
    }
});
