<!DOCTYPE html>
<html>
<head>
    <title>{{.PropertyType}} Rentals in {{.LocationName}}</title>
    <link href="https://fonts.googleapis.com/css2?family=Playfair+Display:wght@600&family=DM+Sans:wght@300;400;500&display=swap" rel="stylesheet">
    <link rel="stylesheet" href="/static/css/variables.css">
    <link rel="stylesheet" href="/static/css/navbar.css">
    <link rel="stylesheet" href="/static/css/footer.css">
    <link rel="stylesheet" type="text/css" href="https://cdn.123presto.com/prod/static/css/global-1.1.80.css"/>
    <link rel="stylesheet" type="text/css" href="https://cdn.123presto.com/prod/static/css/refine-1.1.80.css"/>
    <link rel="stylesheet" type="text/css" href="https://cdn.123presto.com/prod/static/css/tile-1.1.80.css"/>
    <style>
        .grid {
            display: grid !important;
            grid-template-columns: repeat(4, 1fr) !important;
            gap: 18px !important;
        }
        body { padding-left: 40px; padding-right: 30px; }
        .sp-property-card { width: 100% !important; margin: 0 !important; float: none !important; }
        @media (max-width: 1200px) { .grid { grid-template-columns: repeat(3, 1fr) !important; } }
        @media (max-width: 860px)  { .grid { grid-template-columns: repeat(2, 1fr) !important; } }
        @media (max-width: 520px)  { .grid { grid-template-columns: 1fr !important; } }
    </style>
</head>
<body>
    {{template "navbar.tpl" .}}

    {{if .Error}}
    <div style="margin:16px 0; padding:12px; border:1px solid #ef4444; color:#991b1b; background:#fef2f2; border-radius:8px;">
        {{.Error}}
    </div>
    {{end}}

    <div class="refine-breadcrumb" style="margin: 10px 0 18px;">
        <span class="js-place-count">{{.PropertyCount}}</span> {{.PropertyType}} Rentals Near {{.LocationName}} |
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

    <h1 class="title">Best {{.PropertyType}} Rentals in {{.LocationName}}</h1>
    <h2 class="category-sub-title" style="display:inline; font-weight:400;">
        Explore top {{.PropertyType}} stays in {{.LocationName}}.
    </h2>

    <div class="category-content">
        {{if gt (len .Items) 0}}
        <div class="grid" style="display:grid; grid-template-columns:repeat(4,1fr); gap:18px; padding:20px 0;">
            {{template "card.tpl" .}}
        </div>
        {{else}}
        <div style="padding:20px 0; color:#6b7280;">No properties found for this sub-category.</div>
        {{end}}
    </div>

    <!-- View More Properties -->
    <div style="text-align:center; padding:40px 0;">
        <a href="{{.RefineURL}}"
        style="background:var(--primary-color); color:white; padding:14px 32px; border-radius:8px; text-decoration:none; font-weight:700; font-size:16px;">
            View More Properties
        </a>
    </div>

    {{template "footer.tpl" .}}

    <script src="/static/js/redirect.js"></script>
    <script src="/static/js/modal.js"></script>
    <script src="/static/js/currency.js"></script>
    <script src="/static/js/navbar.js"></script>
    <script src="/static/js/footer.js"></script>
    <script src="/static/js/favourite.js"></script>
    <script src="/static/js/imageslider.js"></script>
    <script>
        document.addEventListener("DOMContentLoaded", function () {
            if (typeof initImageSlider === "function") initImageSlider();
        });
    </script>
</body>
</html>