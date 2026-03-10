<!DOCTYPE html>
<html>
<head>
    <title>{{.Country}} - Vacation Rentals</title>
    <link href="https://fonts.googleapis.com/css2?family=Playfair+Display:wght@600&family=DM+Sans:wght@300;400;500&display=swap" rel="stylesheet">
    


    <link rel="stylesheet" href="/static/css/variables.css">
    <link rel="stylesheet" type="text/css" href="https://cdn.123presto.com/prod/static/css/global-1.1.80.css"/>
    <link rel="stylesheet" type="text/css" href="https://cdn.123presto.com/prod/static/css/refine-1.1.80.css"/>
    <link rel="stylesheet" type="text/css" href="https://cdn.123presto.com/prod/static/css/tile-1.1.80.css"/>

    <style>
        .grid {
            display: grid !important;
            grid-template-columns: repeat(4, 1fr) !important;
            gap: 18px !important;
        }
        body {
            padding-left: 40px;
            padding-right: 30px;
        }

        .sp-property-card {
            width: 100% !important;
            margin: 0 !important;
            float: none !important;
        }
        .heart-btn { cursor: pointer; font-size: 18px; color: white; transition: color 0.2s; }
        .heart-btn.active { color: red; }
        .image-section { position: relative !important; }
        .tiles-icons { position: absolute !important; top: 8px !important; right: 8px !important; z-index: 10 !important; }
        .sp-property-card .image-section { position: relative !important; }
        .sp-property-card .image-section .tiles-icons { position: absolute !important; top: 8px !important; right: 8px !important; z-index: 10 !important; }

        @media (max-width: 1200px) { .grid { grid-template-columns: repeat(3, 1fr) !important; } }
        @media (max-width: 860px)  { .grid { grid-template-columns: repeat(2, 1fr) !important; } }
        @media (max-width: 520px)  { .grid { grid-template-columns: 1fr !important; } }
    </style>
</head>
<body>

    <div class="breadcrumb-share">
        <div class="refine-breadcrumb" id="js-breadcrumb">
            <span class="js-place-count">{{.PropertyCount}}+</span> Vacation Rentals Near {{.LocationName}} |
            <ol itemscope="itemscope" itemtype="http://schema.org/BreadcrumbList" style="display:inline; padding:0; margin:0; list-style:none;">
                {{range $i, $bc := .Breadcrumbs}}
                <li itemprop="itemListElement" itemscope="itemscope" itemtype="http://schema.org/ListItem" style="display:inline;">
                    <a itemprop="item" href="/all/{{index $bc "Slug"}}">
                        <span itemprop="name">{{index $bc "Name"}}</span>
                    </a>
                    <meta itemprop="Position" content="{{$i}}">
                </li>
                {{end}}
            </ol>
        </div>
        <div class="refine-share">
            <div class="share-button relative text-right">
                <button id="pt-share-expand" class="share-btn cursor-pointer js-close-share-btn" aria-label="Share">
                    Share
                </button>
            </div>
        </div>
    </div>

    <!-- Title + Subtitle -->
    <h1 class="title" id="js-location-name-h1">Trips to {{.LocationName}}</h1>
    <h2 class="category-sub-title" id="js-location-name-h2" style="display:inline; font-weight:400;">
        Plan Your Next Trip to {{.LocationName}} and explore unforgettable sights, local flavors, and exciting experiences.
    </h2>

    <!-- Filter Bar — same as refine page -->
    <div class="d-flex refine-filters align-item-center sp-flex-wrap justify-between pt-btn-wrap">
        <div id="pt-filter-wrap" class="refine-buttons">

            <div class="relative pt-datepicker" id="js-filter-dp-div">
                <div class="dp-inline" id="filter-dp">
                    <div class="datepicker-input sp-datepicker fl-btn">
                        <button id="standalone-dp">Dates</button>
                    </div>
                </div>
            </div>

            <div class="relative">
                <div class="filter-currency fl-btn" id="js-filter-currency-div">
                    <button class="cursor-pointer relative pt-price-btn" id="js-price-range">Price</button>
                </div>
            </div>

            <div class="relative">
                <div class="filter-guest-div fl-btn" id="js-filter-guest-div">
                    <button class="cursor-pointer relative pt-guest-btn" id="js-guest-picker">Guests</button>
                </div>
            </div>

            <div class="relative">
                <button class="cursor-pointer relative pt-filter-btn fl-btn more-fl-btn" id="filter-btn">More</button>
            </div>

        </div>
    </div>

    <!-- Property Grid -->
    <div class="category-content">
        <div class="grid" style="display:grid; grid-template-columns:repeat(4,1fr); gap:18px; padding:20px 0;">
            {{template "card.tpl" .}}
        </div>
    </div>
    <!-- Sections -->
    {{range .Sections}}
    {{$section := .}}
    <div class="category-section" style="padding: 40px 0 20px;">

        <!-- Section Title -->
        <div class="section-header" style="text-align:center; margin-bottom:20px;">
            <h2 class="section-title" style="font-size:22px; font-weight:800; color:#0b1833;">
                {{index $section "Title"}}
            </h2>
            <p class="section-subtitle" style="font-size:14px; color:#6b7280; margin-top:4px;">
                Over {{index $section "Count"}}+ {{index $section "SubTitle"}}
            </p>
        </div>

        <!-- Section Cards -->
        <div class="grid" style="display:grid; grid-template-columns:repeat(4,1fr); gap:18px;">
            {{$items := index $section "ProcessedItems"}}
            {{template "card.tpl" (dict "Items" $items)}}
        </div>

    </div>
    {{end}}

    <script>
    function redirectToPartner(btn) {
        // Find the parent card
        const card = btn.closest('.sp-property-card');

        console.log("all data attributes:", card.dataset);
        console.log("index value:", card.dataset.index);

        // Read all data attributes
        const propertyId  = card.dataset.property_id;
        const feed        = card.dataset.feed;
        const published   = card.dataset.published;
        const upat        = card.dataset.upat;
        const lat         = card.dataset.lat;
        const lng         = card.dataset.lng;
        const type        = card.dataset.type;
        const destId      = card.dataset.dest_id;
        const ownerId     = card.dataset.owner_id;
        const directUrl   = card.dataset.direct_url;
        const display     = card.dataset.display;
        const city        = card.dataset.city;
        const country     = card.dataset.country;
        const countryCode = card.dataset.country_code;
        const epc         = card.dataset.epc;
        const eplId       = card.dataset.eplid;
        const index       = card.dataset.index; 

        // Generate menu_id — timestamp of this moment
        const menuId = Date.now();

        // Build referral_id — index not available so use 0
        const referralId = index + "-" + propertyId;

        // Build params
        const params = new URLSearchParams({
            menu_id:            menuId,
            lang:               "en-US",
            feed:               feed,
            property_id:        propertyId,
            published:          published,
            upat:               upat,
            latitude:           lat,
            longitude:          lng,
            type:               type.toLowerCase(),
            referrer_page:      "category",
            guests:             "2",
            dest_id:            destId,
            owner_id:           ownerId,
            direct_redirect_url: directUrl,
            search_string:      display,
            epc:                epc,
            referral_id:        referralId,
            pl:                 city,
            pc:                 country,
            pcc:                countryCode,
            eplId:              eplId,
            user_type:          "dd",
            currency:           "BDT"
        });

        const redirectUrl = "http://localhost:8080/redirect-partner?" + params.toString();
        console.log("Redirecting to:", redirectUrl);

        window.open(redirectUrl, "_blank");
    }
    </script>
    
    <script src="/static/js/favourite.js"></script>
    <script>
        document.addEventListener('DOMContentLoaded', function() {
        let favourites = JSON.parse(localStorage.getItem('favourite_list') || '{}');
        Object.keys(favourites).forEach(id => {
                const btn = document.querySelector(`.heart-btn[data-id="${id}"]`);
                if (btn) {
                    const icon = btn.querySelector('.heart-icon');
                    icon.setAttribute('fill', 'red');
                    icon.setAttribute('stroke', 'red');
                }
            });
        });
    </script>

</body>
</html>