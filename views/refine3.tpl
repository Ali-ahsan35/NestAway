<!DOCTYPE html>
<html>
<head>
    <title>Vacation Rentals</title>
    <link href="https://fonts.googleapis.com/css2?family=Mulish:wght@300;400;500;600;700;800&display=swap" rel="stylesheet">
    <style>
        /* ── CSS Variables ── */
        :root {
            --primary-color: #013573;
            --secondary-color: #00cd92;
            --font-family: 'Mulish', sans-serif;
            --text-dark: #0b1833;
            --text-gray: #6b7280;
            --border: #e5e7eb;
            --bg: #f5f6fa;
        }

        * { margin: 0; padding: 0; box-sizing: border-box; }

        body {
            font-family: var(--font-family);
            background: #fff;
            color: var(--text-dark);
            font-size: 14px;
        }

        /* ── Utility classes (matching real site) ── */
        .d-flex, .sp-flex { display: flex; }
        .sp-flex-wrap { flex-wrap: wrap; }
        .align-item-center { align-items: center; }
        .justify-between { justify-content: space-between; }
        .justify-content-end { justify-content: flex-end; }
        .hidden { display: none !important; }
        .relative { position: relative; }
        .ellipsis { overflow: hidden; text-overflow: ellipsis; white-space: nowrap; }

        /* ── Page Header ── */
        .page-header {
            padding: 16px 32px 20px;
            background: #fff;
            border-bottom: 1px solid var(--border);
        }

        .breadcrumb-nav {
            font-size: 12px;
            color: var(--text-gray);
            margin-bottom: 10px;
            display: flex;
            align-items: center;
            gap: 4px;
        }

        .breadcrumb-nav a {
            color: var(--text-gray);
            text-decoration: none;
        }

        .breadcrumb-nav a:hover { color: var(--primary-color); }

        .breadcrumb-nav .sep {
            color: #d1d5db;
            font-size: 11px;
        }

        .page-header h1 {
            font-size: 28px;
            font-weight: 800;
            color: var(--text-dark);
            margin-bottom: 4px;
            line-height: 1.2;
        }

        .page-header p {
            font-size: 14px;
            color: var(--text-gray);
            font-weight: 400;
        }

        /* ── Filter Bar ── */
        .refine-filters {
            padding: 12px 32px;
            background: #fff;
            border-bottom: 1px solid var(--border);
            display: flex;
            align-items: center;
            justify-content: space-between;
            gap: 12px;
        }

        /* Filter buttons group */
        .refine-buttons {
            display: flex;
            align-items: center;
            gap: 8px;
            flex-wrap: wrap;
        }

        /* Each filter button pill */
        .fl-btn {
            display: inline-flex;
            align-items: center;
            gap: 6px;
            padding: 7px 16px;
            border: 1.5px solid #d1d5db;
            border-radius: 999px;
            background: #fff;
            cursor: pointer;
            font-family: var(--font-family);
            font-size: 13px;
            font-weight: 600;
            color: var(--text-dark);
            transition: all 0.15s;
            white-space: nowrap;
        }

        .fl-btn:hover {
            border-color: var(--primary-color);
            color: var(--primary-color);
        }

        .fl-btn button {
            background: none;
            border: none;
            font-family: var(--font-family);
            font-size: 13px;
            font-weight: 600;
            color: inherit;
            cursor: pointer;
            padding: 0;
        }

        /* dp-inline wrapper — must NOT be poup-container */
        .dp-inline { position: relative; }

        /* ── Sort Dropdown (select-wrap) ── */
        .pt-sort-wrap {
            display: flex;
            align-items: center;
            gap: 8px;
        }

        .pt-sort-wrap .title {
            font-size: 14px;
            font-weight: 700;
            color: var(--text-dark);
            white-space: nowrap;
        }

        .select-wrap {
            width: 175px;
            position: relative;
            user-select: none;
            display: inline-block;
            z-index: 2;
        }

        /* The visible "button" showing current selection */
        .select-wrap .default-option {
            position: relative;
            cursor: pointer;
            list-style: none;
            margin: 0;
            padding: 0;
            border: 1.5px solid var(--primary-color);
            border-radius: 999px;
            background: #fff;
        }

        .select-wrap .default-option li {
            padding: 7px 36px 7px 14px;
            font-size: 13px;
            font-weight: 600;
            color: var(--primary-color);
            line-height: normal;
        }

        /* Arrow indicator */
        .select-wrap .default-option::before {
            content: "";
            position: absolute;
            top: 50%;
            right: 14px;
            width: 8px;
            height: 8px;
            margin-top: -6px;
            border: 2px solid var(--primary-color);
            border-top: none;
            border-left: none;
            transform: rotate(45deg);
            transition: transform 0.2s;
        }

        .select-wrap.active .default-option::before {
            margin-top: -2px;
            transform: rotate(-135deg);
        }

        /* Dropdown list */
        .select-wrap .select-ul {
            list-style: none;
            margin: 4px 0 0;
            padding: 4px 0;
            position: absolute;
            top: 100%;
            left: 0;
            width: 100%;
            background: #fff;
            border-radius: 10px;
            display: none;
            z-index: 99;
            box-shadow: 0 4px 20px rgba(0,0,0,0.12);
        }

        .select-wrap.active .select-ul { display: block; }

        .select-wrap .select-ul li {
            cursor: pointer;
            padding: 8px 14px;
        }

        .select-wrap .select-ul li:hover { background: #f0f5ff; }

        .select-wrap .select-ul li .option p {
            margin: 0;
            font-size: 13px;
            color: var(--text-dark);
            font-weight: 500;
        }

        /* ── Properties Container ── */
        .properties-container {
            padding: 20px 32px 40px;
            background: var(--bg);
            min-height: 80vh;
        }

        .result-count {
            font-size: 13px;
            color: var(--text-gray);
            margin-bottom: 16px;
            font-weight: 500;
        }

        /* ── Property Grid ── */
        .properties-grid {
            display: grid;
            grid-template-columns: repeat(4, 1fr);
            gap: 18px;
        }

        /* ── Property Card ── */
        .property-card {
            background: #fff;
            border-radius: 10px;
            overflow: hidden;
            border: 1px solid var(--border);
            transition: transform 0.2s, box-shadow 0.2s;
            animation: fadeUp 0.35s ease both;
        }

        .property-card:hover {
            transform: translateY(-3px);
            box-shadow: 0 8px 24px rgba(1,53,115,0.1);
        }

        @keyframes fadeUp {
            from { opacity: 0; transform: translateY(14px); }
            to   { opacity: 1; transform: translateY(0); }
        }

        /* Card image */
        .card-img {
            width: 100%;
            height: 185px;
            object-fit: cover;
            display: block;
        }

        .card-img-placeholder {
            width: 100%;
            height: 185px;
            background: linear-gradient(135deg, #e8eef5, #dfe6f1);
            display: flex;
            align-items: center;
            justify-content: center;
            font-size: 36px;
        }

        /* Card body */
        .card-body { padding: 11px 12px 12px; }

        .card-top {
            display: flex;
            justify-content: space-between;
            align-items: center;
            margin-bottom: 5px;
        }

        .card-rating {
            display: flex;
            align-items: center;
            gap: 5px;
            font-size: 12px;
            color: var(--text-gray);
        }

        .rating-badge {
            background: var(--secondary-color);
            color: #fff;
            padding: 2px 7px;
            border-radius: 4px;
            font-size: 11px;
            font-weight: 700;
        }

        .card-type {
            font-size: 10px;
            color: var(--text-gray);
            text-transform: uppercase;
            letter-spacing: 0.6px;
            font-weight: 600;
        }

        .card-name {
            font-size: 13px;
            font-weight: 700;
            color: var(--text-dark);
            margin-bottom: 3px;
            line-height: 1.4;
            display: -webkit-box;
            -webkit-line-clamp: 2;
            -webkit-box-orient: vertical;
            overflow: hidden;
        }

        .card-location {
            font-size: 11px;
            color: var(--text-gray);
            margin-bottom: 6px;
        }

        .card-price {
            font-size: 13px;
            font-weight: 700;
            color: var(--text-dark);
            margin-bottom: 10px;
        }

        .card-price span {
            font-size: 11px;
            color: var(--text-gray);
            font-weight: 400;
        }

        .card-footer {
            display: flex;
            justify-content: space-between;
            align-items: center;
            border-top: 1px solid var(--border);
            padding-top: 9px;
        }

        .card-source {
            font-size: 11px;
            color: var(--primary-color);
            font-weight: 600;
        }

        .card-btn {
            background: var(--secondary-color);
            color: #fff;
            border: none;
            padding: 6px 13px;
            border-radius: 6px;
            font-size: 11px;
            font-family: var(--font-family);
            font-weight: 700;
            cursor: pointer;
            transition: background 0.2s;
        }

        .card-btn:hover { background: #00b57e; }

        /* ── Loading spinner ── */
        .loading-wrap {
            grid-column: span 4;
            text-align: center;
            padding: 80px 20px;
            color: var(--text-gray);
        }

        .spinner {
            display: inline-flex;
            gap: 6px;
            margin-bottom: 12px;
        }

        .spinner > div {
            width: 10px;
            height: 10px;
            background: var(--primary-color);
            border-radius: 50%;
            animation: bounce 1.2s infinite ease-in-out both;
        }

        .spinner .bounce1 { animation-delay: -0.32s; }
        .spinner .bounce2 { animation-delay: -0.16s; }

        @keyframes bounce {
            0%, 80%, 100% { transform: scale(0); opacity: 0.3; }
            40%            { transform: scale(1); opacity: 1; }
        }

        /* ── Responsive ── */
        @media (max-width: 1200px) { .properties-grid { grid-template-columns: repeat(3, 1fr); } }
        @media (max-width: 860px)  { .properties-grid { grid-template-columns: repeat(2, 1fr); } }
        @media (max-width: 520px)  {
            .properties-grid { grid-template-columns: 1fr; }
            .refine-filters { flex-wrap: wrap; }
            .page-header { padding: 14px 16px; }
            .properties-container { padding: 16px; }
        }
    </style>
</head>

<body>

<!-- ── PAGE HEADER ── -->
<div class="page-header">
    <nav class="breadcrumb-nav" id="breadcrumb">
        <span>Loading...</span>
    </nav>
    <h1 id="page-title">Find a Place to Stay</h1>
    <p id="page-subtitle">Loading properties...</p>
</div>

<!-- ── FILTER BAR ── exact class names from real site -->
<div class="refine-filters d-flex align-item-center sp-flex-wrap justify-between pt-btn-wrap">

    <!-- Left: filter buttons -->
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
                <button class="pt-price-btn" id="js-price-range">Price</button>
            </div>
        </div>

        <div class="relative">
            <div class="filter-guest-div fl-btn" id="js-filter-guest-div">
                <button class="pt-guest-btn" id="js-guest-picker">Guests</button>
            </div>
        </div>

        <div class="relative">
            <button class="pt-filter-btn fl-btn more-fl-btn" id="filter-btn">More</button>
        </div>

    </div>

    <!-- Right: sort dropdown -->
    <div class="wrapper d-flex align-item-center justify-content-end pt-sort-wrap">
        <span class="title ellipsis">Sort by</span>
        <div id="js-filter-sort" class="select-wrap js-dropdown">
            <input type="hidden" class="js-selected-value" id="sort-properties" value="1">
            <ul class="default-option pt-sort-default">
                <li data-value="1">
                    <div class="option"><p class="ellipsis pt-sort-item">Most Popular</p></div>
                </li>
            </ul>
            <ul class="select-ul">
                <li id="js-order-1" data-value="1"><div class="option"><p>Most Popular</p></div></li>
                <li id="js-order-3" data-value="3"><div class="option"><p>Highest Price</p></div></li>
                <li id="js-order-2" data-value="2"><div class="option"><p>Lowest Price</p></div></li>
                <li id="js-order-5" data-value="5"><div class="option"><p>Highest Rating</p></div></li>
                <li id="js-order-4" data-value="4"><div class="option"><p>Lowest Rating</p></div></li>
            </ul>
        </div>
    </div>

</div>

<!-- ── PROPERTIES ── -->
<div class="properties-container">
    <div class="result-count" id="result-count"></div>
    <div class="properties-grid" id="grid">
        <div class="loading-wrap">
            <div class="spinner">
                <div class="bounce1"></div>
                <div class="bounce2"></div>
                <div class="bounce3"></div>
            </div>
            <div>Loading properties...</div>
        </div>
    </div>
</div>

<script src="/static/js/refine.js"></script>

<script>
    // Sort dropdown toggle
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
        });
    });

    document.addEventListener('click', (e) => {
        if (!sortWrap.contains(e.target)) sortWrap.classList.remove('active');
    });
</script>

</body>
</html>
