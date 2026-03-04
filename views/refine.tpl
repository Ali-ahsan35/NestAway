<!DOCTYPE html>
<html>
<head>
    <title>Vacation Rentals</title>
    <link href="https://fonts.googleapis.com/css2?family=Playfair+Display:wght@600&family=DM+Sans:wght@300;400;500&display=swap" rel="stylesheet">
    
    <link rel="stylesheet" href="static/css/style.css">

</head>
<body>

<!-- HEADER -->
<div class="header">
    <div class="breadcrumb" id="breadcrumb">
        <span>Loading...</span>
    </div>
    <h1 id="page-title">Find a Place to Stay</h1>
    <p id="page-subtitle">Loading properties...</p>
</div>

<!-- FILTERS — exact class names from real site -->
<div class="refine-filters d-flex align-item-center sp-flex-wrap justify-between pt-btn-wrap">
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
            <button class="cursor-pointer relative pt-filter-btn fl-btn more-fl-btn" id="filter-btn">
                More
            </button>
        </div>

    </div>

    <!-- SORT -->
    <div class="wrapper d-flex align-item-center justify-content-end pt-sort-wrap">
    <span class="title color-primary ellipsis" title="Sort by">Sort by </span>
    <div id="js-filter-sort" class="select-wrap js-dropdown">
        <input type="hidden" class="js-selected-value" id="sort-properties">
        <ul class="default-option pt-sort-default">
            <li data-value="1">
                <div class="option">
                    <p class="ellipsis pt-sort-item">Most Popular</p>
                </div>
            </li>
        </ul>
        <ul class="select-ul">
            <li id="js-order-1" data-value="1">
                <div class="option"><p class="ellipsis pt-sort-item">Most Popular</p></div>
            </li>
            <li id="js-order-3" data-value="3">
                <div class="option"><p class="ellipsis pt-sort-item">Highest Price</p></div>
            </li>
            <li id="js-order-2" data-value="2">
                <div class="option"><p class="ellipsis pt-sort-item">Lowest Price</p></div>
            </li>
            <li id="js-order-5" data-value="5">
                <div class="option"><p class="ellipsis pt-sort-item">Highest Rating</p></div>
            </li>
            <li id="js-order-4" data-value="4">
                <div class="option"><p class="ellipsis pt-sort-item">Lowest Rating</p></div>
            </li>
        </ul>
    </div>
</div>
</div>

<!-- PROPERTY GRID -->
<div class="container">
    <div class="result-count" id="result-count"></div>
    <div class="grid" id="grid">
        <div class="loading">
            <span class="loading-dot">●</span>
            <span class="loading-dot">●</span>
            <span class="loading-dot">●</span>
            &nbsp; Loading properties...
        </div>
    </div>
</div>

<script src="/static/js/refine.js"></script>
<script>
    const sortWrap = document.getElementById('js-filter-sort');
    const defaultOpt = sortWrap.querySelector('.default-option');
    const hiddenInput = document.getElementById('sort-properties');

    defaultOpt.addEventListener('click', () => {
        sortWrap.classList.toggle('active');
    });

    sortWrap.querySelectorAll('.select-ul li').forEach(li => {
        li.addEventListener('click', () => {
            const value = li.getAttribute('data-value');
            const text  = li.querySelector('p').textContent;
            defaultOpt.querySelector('p').textContent = text;
            hiddenInput.value = value;
            sortWrap.classList.remove('active');

            // Call loadProperties from refine.js
            if (window.loadProperties && window.currentCategory) {
                window.loadProperties(window.currentCategory, value);
            }
        });
    });

    document.addEventListener('click', (e) => {
        if (!sortWrap.contains(e.target)) sortWrap.classList.remove('active');
    });
</script>

</body>
</html>