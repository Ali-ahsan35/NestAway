<!DOCTYPE html>
<html>
<head>
    <title>{{.Country}} - Vacation Rentals</title>
    <link href="https://fonts.googleapis.com/css2?family=Playfair+Display:wght@600&family=DM+Sans:wght@300;400;500&display=swap" rel="stylesheet">
    
    <link rel="stylesheet" type="text/css" href="https://cdn.123presto.com/prod/static/css/hotel-datepicker-1.1.83.css"/>
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
        /* Datepicker Modal */
        #js-datepicker-modal .datepicker__month {
            width: auto !important;
        }
        #js-datepicker-modal .datepicker {
            box-shadow: none !important;
            border: none !important;
            width: 100% !important;
        }
        #js-datepicker-modal .datepicker__inner {
            padding: 0 !important;
        }
        #js-datepicker-modal .datepicker__months {
            display: flex !important;
            gap: 40px !important;
            justify-content: center !important;
        }
        #js-datepicker-modal .datepicker__month-name {
            font-size: 14px !important;
            font-weight: 700 !important;
            color: #0b1833 !important;
            text-transform: uppercase !important;
            letter-spacing: 0.5px !important;
        }
        #js-datepicker-modal .datepicker__week-name {
            font-size: 12px !important;
            color: #6b7280 !important;
            font-weight: 600 !important;
        }
        #js-datepicker-modal .datepicker__month-day {
            font-size: 13px !important;
            padding: 8px !important;
            border-radius: 50% !important;
            cursor: pointer !important;
        }
        #js-datepicker-modal .datepicker__month-day--invalid {
            color: #ccc !important;
            cursor: default !important;
        }
        #js-datepicker-modal .datepicker__month-day--valid:hover {
            background: #013573 !important;
            color: white !important;
        }
        #js-datepicker-modal .datepicker__month-day--first-day-selected,
        #js-datepicker-modal .datepicker__month-day--last-day-selected {
            background: #013573 !important;
            color: white !important;
            border-radius: 50% !important;
        }
        #js-datepicker-modal .datepicker__month-day--selected {
            background: #e8f0fe !important;
            border-radius: 0 !important;
        }
        #js-datepicker-modal .datepicker__month-day--today {
            font-weight: 700 !important;
            color: #013573 !important;
        }
        #js-datepicker-modal .datepicker__topbar {
            background: white !important;
            border-top: 1px solid #eee !important;
            padding: 16px 0 0 0 !important;
            display: flex !important;
            align-items: center !important;
            justify-content: space-between !important;
        }
        #js-datepicker-modal .datepicker__info--selected {
            font-size: 13px !important;
            color: #333 !important;
            font-weight: 600 !important;
        }
        #js-datepicker-modal .datepicker__info-text--start-day,
        #js-datepicker-modal .datepicker__info-text--end-day {
            text-transform: uppercase !important;
            font-weight: 700 !important;
        }
        #js-datepicker-modal .datepicker__submit-button {
            background: #ff6b35 !important;
            color: white !important;
            border: none !important;
            border-radius: 8px !important;
            padding: 12px 28px !important;
            font-weight: 700 !important;
            font-size: 14px !important;
            cursor: pointer !important;
            display: flex !important;
            flex-direction: column !important;
            align-items: center !important;
            line-height: 1.3 !important;
        }
        #js-datepicker-modal .btn-skip {
            background: none !important;
            border: none !important;
            color: #ff6b35 !important;
            font-weight: 600 !important;
            font-size: 14px !important;
            cursor: pointer !important;
        }
        #js-datepicker-modal .datepicker__month-button {
            font-size: 18px !important;
            color: #013573 !important;
            cursor: pointer !important;
        }
        #js-datepicker-modal .datepicker__info--feedback {
            display: none !important;
        }

        #datepicker-modal-datepicker {
        background: white !important;
        border: 1px solid #ddd !important;
        box-shadow: 0 8px 30px rgba(0,0,0,0.15) !important;
        z-index: 99999 !important;
        }
        #datepicker-modal-datepicker .datepicker__inner {
            background: white !important;
        }
        #datepicker-modal-datepicker .datepicker__month {
            background: white !important;
            width: auto !important;
        }
        #datepicker-modal-datepicker .datepicker__month-day--valid:hover {
            background: #013573 !important;
            color: white !important;
        }
        #datepicker-modal-datepicker .datepicker__month-day--first-day-selected,
        #datepicker-modal-datepicker .datepicker__month-day--last-day-selected {
            background: #013573 !important;
            color: white !important;
        }
        #datepicker-modal-datepicker .datepicker__month-day--selected {
            background: #e8f0fe !important;
        }
        #datepicker-modal-datepicker .datepicker__topbar {
            background: white !important;
            border-top: 1px solid #eee !important;
        }

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
                    <button class="fl-btn cursor-pointer" id="standalone-dp">Dates</button>
                    <input type="text" id="js-dp-input" style="display:none;">
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

    <!-- FILTERS MODAL -->
    <div class="poup-container filter-popup js-modal hidden" id="js-filter">
    <div class="popup-layer"></div>
    <div class="modal-container">
        <div class="modal-header">
        <div class="popup-title font-20 font-semibold sp-text-center">
            Filters
            <div class="filter-cross-btn close-btn js-trigger-hide">✕</div>
        </div>
        </div>
        <div class="modal-body">
        <div class="top">
            <div class="filter-pop-up-container">

            <!-- Pet friendly -->
            <label class="relative cursor-pointer font-22 filter-level filter-pet-eco">
                <div class="d-flex">
                <div class="icon-text">
                    <div class="icon-title">
                    <div class="popup-check-box-label-title filter-pet-eco-title font-16">Pet-friendly only</div>
                    </div>
                    <input class="pet-eco-checkmark pet-checkmark js-pet-friendly amenities-cb" value="11" type="checkbox" id="amenity-11">
                </div>
                </div>
            </label>

            <!-- Eco friendly -->
            <label class="relative cursor-pointer font-22 filter-level filter-pet-eco">
                <div class="d-flex">
                <div class="icon-text">
                    <div class="icon-title">
                    <div class="popup-check-box-label-title filter-pet-eco-title font-16">Eco-friendly only</div>
                    </div>
                    <input class="pet-eco-checkmark eco-checkmark" type="checkbox" id="js-eco-friendly">
                </div>
                </div>
            </label>

            <!-- Dates -->
            <div class="filter-check-in-out">
                <label class="filter-input-label" for="modal-datepicker"></label>
                <div class="choose-date">
                <div class="dp-inline" id="modal-dp">
                    <input class="datepicker-input sp-datepicker" readonly placeholder="Select a date" id="modal-datepicker">
                    <div class="calendar-icon">📅</div>
                </div>
                </div>
            </div>

            <!-- Guests -->
            <div class="filter-guest" style="border:none;">
                <div class="d-flex align-item-center justify-between w-100">
                <div class="popup-check-box-label-title font-16">Guests</div>
                <div class="guestPopup">
                    <div class="top">
                    <div class="wrapper">
                        <div class="js-guest-filter guest-buttons d-flex align-item-center" style="gap:10px;">
                        <button type="button" class="guest decrease" id="js-guest-decrease" style="width:28px;height:28px;border-radius:4px;background:var(--secondary-color);color:#fff;font-size:18px;display:flex;align-items:center;justify-content:center;">−</button>
                        <span class="guest guest-count" id="js-guest-count" style="font-size:18px;min-width:30px;text-align:center;">0</span>
                        <button type="button" class="guest increase" id="js-guest-increase" style="width:28px;height:28px;border-radius:4px;background:var(--secondary-color);color:#fff;font-size:18px;display:flex;align-items:center;justify-content:center;">+</button>
                        </div>
                    </div>
                    </div>
                </div>
                </div>
            </div>

            <!-- Price Range -->
            <div class="popup-price-range row" style="border:2px solid var(--secondary-color);">
                <div class="filter-price-range">
                    <div class="popup-check-box-label-title font-16 sp-text-center">Price range</div>
                    <div class="price-slider js-price-slider">
                        <div class="slider-range" id="js-slider-range" style="left:0%;width:100%;"></div>
                        <input value="244" min="244" max="122096" step="100" id="js-min-price-slider" type="range">
                        <input value="122096" min="244" max="122096" step="100" id="js-max-price-slider" type="range">
                    </div>
                </div>
                <div class="popup-min-max-price align-item-center js-price-input">
                    <div class="popup-min-price-main">
                        <label class="popup-min-max-label">Min price</label>
                        <div class="popup-min-price cursor-pointer">
                            <span class="currency-text-box">
                                <span class="currency-icon">BD ৳</span>
                                <input class="popup-min-price-bottom" type="number" id="js-min-price" value="244" min="244" max="122096">
                            </span>
                        </div>
                    </div>
                    <div class="hyphen"></div>
                    <div class="popup-max-price-main">
                        <div class="popup-max-price cursor-pointer">
                            <span class="currency-text-box">
                                <span class="currency-icon">BD ৳</span>
                                <input class="popup-min-price-bottom" type="number" id="js-max-price" value="122096" min="244" max="122096">
                            </span>
                        </div>
                        <label class="popup-min-max-label">Max price</label>
                    </div>
                </div>
            </div>

            <!-- Amenities -->
            <div class="popup-amenities row">
                <div class="popup-amenities-container">
                <div class="popup-check-box-label-title popup-amenities-title font-16">Amenities</div>
                <div class="popup-amenities-checkbox" id="js-dynamic-amenities-filter">
                        <label class="popup-single-check-box relative cursor-pointer filter-amenities-level">
                            <span class="inline-block ellipsis filter-amenities-span">Air Conditioner</span>
                            <input class="absolute cursor-pointer amenities-cb" type="checkbox" id="amenity-1" value="1">
                            <span class="checkmark absolute inline-block">&nbsp;</span>
                        </label>
                        <label class="popup-single-check-box relative cursor-pointer filter-amenities-level">
                            <span class="inline-block ellipsis filter-amenities-span">Balcony/Terrace</span>
                            <input class="absolute cursor-pointer amenities-cb" type="checkbox" id="amenity-2" value="2">
                            <span class="checkmark absolute inline-block">&nbsp;</span>
                        </label>
                        <label class="popup-single-check-box relative cursor-pointer filter-amenities-level">
                            <span class="inline-block ellipsis filter-amenities-span">Bedding/Linens</span>
                            <input class="absolute cursor-pointer amenities-cb" type="checkbox" id="amenity-3" value="3">
                            <span class="checkmark absolute inline-block">&nbsp;</span>
                        </label>
                        <label class="popup-single-check-box relative cursor-pointer filter-amenities-level">
                            <span class="inline-block ellipsis filter-amenities-span">Breakfast</span>
                            <input class="absolute cursor-pointer amenities-cb" type="checkbox" id="amenity-4" value="4">
                            <span class="checkmark absolute inline-block">&nbsp;</span>
                        </label>
                        <label class="popup-single-check-box relative cursor-pointer filter-amenities-level">
                            <span class="inline-block ellipsis filter-amenities-span">Child Friendly</span>
                            <input class="absolute cursor-pointer amenities-cb" type="checkbox" id="amenity-5" value="5">
                            <span class="checkmark absolute inline-block">&nbsp;</span>
                        </label>
                        <label class="popup-single-check-box relative cursor-pointer filter-amenities-level">
                            <span class="inline-block ellipsis filter-amenities-span">Hot Tub</span>
                            <input class="absolute cursor-pointer amenities-cb" type="checkbox" id="amenity-6" value="6">
                            <span class="checkmark absolute inline-block">&nbsp;</span>
                        </label>
                        <label class="popup-single-check-box relative cursor-pointer filter-amenities-level">
                            <span class="inline-block ellipsis filter-amenities-span">Internet/Wifi</span>
                            <input class="absolute cursor-pointer amenities-cb" type="checkbox" id="amenity-7" value="7">
                            <span class="checkmark absolute inline-block">&nbsp;</span>
                        </label>
                        <label class="popup-single-check-box relative cursor-pointer filter-amenities-level">
                            <span class="inline-block ellipsis filter-amenities-span">Kitchen</span>
                            <input class="absolute cursor-pointer amenities-cb" type="checkbox" id="amenity-8" value="8">
                            <span class="checkmark absolute inline-block">&nbsp;</span>
                        </label>
                        <label class="popup-single-check-box relative cursor-pointer filter-amenities-level">
                            <span class="inline-block ellipsis filter-amenities-span">Laundry</span>
                            <input class="absolute cursor-pointer amenities-cb" type="checkbox" id="amenity-9" value="9">
                            <span class="checkmark absolute inline-block">&nbsp;</span>
                        </label>
                        <label class="popup-single-check-box relative cursor-pointer filter-amenities-level">
                            <span class="inline-block ellipsis filter-amenities-span">Parking</span>
                            <input class="absolute cursor-pointer amenities-cb" type="checkbox" id="amenity-10" value="10">
                            <span class="checkmark absolute inline-block">&nbsp;</span>
                        </label>
                        <label class="popup-single-check-box relative cursor-pointer filter-amenities-level">
                            <span class="inline-block ellipsis filter-amenities-span">Pool</span>
                            <input class="absolute cursor-pointer amenities-cb" type="checkbox" id="amenity-12" value="12">
                            <span class="checkmark absolute inline-block">&nbsp;</span>
                        </label>
                        <label class="popup-single-check-box relative cursor-pointer filter-amenities-level">
                            <span class="inline-block ellipsis filter-amenities-span">TV</span>
                            <input class="absolute cursor-pointer amenities-cb" type="checkbox" id="amenity-14" value="14">
                            <span class="checkmark absolute inline-block">&nbsp;</span>
                        </label>
                        <label class="popup-single-check-box relative cursor-pointer filter-amenities-level">
                            <span class="inline-block ellipsis filter-amenities-span">Wheelchair Accessible</span>
                            <input class="absolute cursor-pointer amenities-cb" type="checkbox" id="amenity-16" value="16">
                            <span class="checkmark absolute inline-block">&nbsp;</span>
                        </label>
                    </div>
                </div>
            </div>

            </div>
        </div>
        </div>

        <!-- Footer buttons -->
        <div class="modal-footer">
        <div class="filter-pop-up-container d-flex justify-content-between">
            <div class="popup-buttons d-flex align-item-center cursor-pointer">
            <button class="popup-clear-btn cursor-pointer" id="js-clear-filter">Clear</button>
            <div class="search-btn-area">
                <button class="popup-submit-btn cursor-pointer" id="js-apply-filter">
                <span id="js-filter-select-text">Search</span>
                </button>
            </div>
            </div>
        </div>
    </div>
    </div>
    </div>
    <div class="datepicker-modal hidden" id="js-datepicker-modal" style="position:fixed; top:0; left:0; width:100%; height:100%; z-index:9999;">
        <div class="popup-layer" id="dp-overlay" style="position:absolute; top:0; left:0; width:100%; height:100%; background:rgba(0,0,0,0.5);"></div>
        <div style="position:absolute; top:50%; left:50%; transform:translate(-50%,-50%); background:white; border-radius:12px; padding:24px; width:720px; max-width:95vw;">
            <div style="display:flex; justify-content:space-between; align-items:center; margin-bottom:16px;">
                <h3 style="font-size:18px; font-weight:700; color:#0b1833;">When do you want to travel?</h3>
                <div style="display:flex; gap:12px; align-items:center;">
                    <button id="js-dp-clear" style="background:none; border:none; color:#ff6b35; font-weight:600; font-size:14px; cursor:pointer;">Clear</button>
                    <span id="js-dp-close" style="cursor:pointer; font-size:20px;">✕</span>
                </div>
            </div>
            <div id="js-dp-container">
                <input type="text" id="js-dp-input" style="display:none;">
            </div>
        </div>
    </div>

    <script src="/static/js/redirect.js"></script>
    <script src="/static/js/modal.js"></script>
    <script>
        window.locationName = "{{.LocationName}}";
        window.currentCategory = "{{.Country}}";
        
        window.loadProperties = function(category, order, filters = {}) {
            console.log("all.tpl loadProperties called:", category, order, filters);
            const params = new URLSearchParams();
            params.set('search', window.locationName);
            params.set('order', order || '1');

            if (filters.amenities && filters.amenities.length > 0) {
                params.set('amenities', filters.amenities.join('-'));
            }
            if (filters.ecoFriendly) {
                params.set('ecoFriendly', 'true');
            }
            if (filters.amountBDT) {
                params.set('amount', filters.amountBDT);
                params.set('selectedCurrency', 'BDT');
            }
            if (filters.guests && filters.guests !== '0') {
                params.set('pax', filters.guests);
            }
            if (filters.checkin) {
                params.set('dateStart', filters.checkin);
            }
            if (filters.checkout) {
                params.set('dateEnd', filters.checkout);
            }

            window.location.href = '/refine?' + params.toString();
        };
    </script>
    <script src="https://cdn.jsdelivr.net/npm/fecha@4.2.3/lib/fecha.umd.js"></script>
    <script src="https://cdn.jsdelivr.net/npm/hotel-datepicker@4.5.0/dist/js/hotel-datepicker.min.js"></script>
    <script src="/static/js/datepicker.js"></script>

</body>
</html>