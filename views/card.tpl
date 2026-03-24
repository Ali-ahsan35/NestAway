{{range .Items}}
{{$item := .}}
{{$prop := index $item "Property"}}
{{$geo := index $item "GeoInfo"}}
{{$partner := index $item "Partner"}}
{{$id := index $item "ID"}}

{{if $prop}}
<div class="sp-property-card"
    data-property_id="{{$id}}"
    data-feed="{{index $item "Feed"}}"
    data-published="{{index $item "Published"}}"
    data-upat="{{index $prop "UpdatedAt"}}"
    data-lat="{{index $geo "Lat"}}"
    data-lng="{{index $geo "Lng"}}"
    data-type="{{index $prop "PropertyType"}}"
    data-dest_id="{{index $partner "ID"}}"
    data-owner_id="{{index $partner "OwnerID"}}"
    data-direct_url="{{index $partner "URL"}}"
    data-display="{{index $geo "Display"}}"
    data-city="{{index $geo "City"}}"
    data-country="{{index $geo "Country"}}"
    data-country_code="{{index $geo "CountryCode"}}"
    data-epc="{{index $partner "EpCluster"}}"
    data-eplid="{{index $geo "LocationID"}}"
    data-index="{{index $item "Index"}}"
    data-slug="{{index $prop "PropertySlug"}}">

    <!-- Image -->
    <div class="image-section relative" style="position:relative;">
        <div class="tiles-icons" style="position:absolute; top:8px; right:8px; z-index:10;">
            <div class="tiles-btn fav-icon heart-btn" data-id="{{$id}}" title="Bookmark" onclick="toggleFavourite(this)">
                <svg class="heart-icon" xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="white" stroke-width="2">
                    <path d="M20.84 4.61a5.5 5.5 0 0 0-7.78 0L12 5.67l-1.06-1.06a5.5 5.5 0 0 0-7.78 7.78l1.06 1.06L12 21.23l7.78-7.78 1.06-1.06a5.5 5.5 0 0 0 0-7.78z"></path>
                </svg>
            </div>
        </div>
        <a rel="nofollow" target="_blank" class="sp-property-image" href="#" onmouseenter="buildRedirectUrl(this)" onclick="redirectToPartner(this); return false;">
            {{$imgName := index $prop "FeatureImage"}}
            {{if $imgName}}
            <img class="featured-image pt-featured-image"
                src="https://imgservice.ownerdirect.com/600x600/{{$imgName}}"
                alt="{{index $prop "PropertyName"}}">
            {{end}}
            {{$price := index $prop "Price"}}
            {{if $price}}
            <span class="property-price js-price-value">
                From BD ৳ {{printf "%.0f" (mul $price 120.0)}}
            </span>
            {{end}}
        </a>
    </div>

    <!-- Details -->
    <div class="sp-property-details">

        <div class="property-review pt-property-review">
            <div class="rating-review pt-rating-review">
                <div class="reviews pt-reviews">
                    {{$score := index $prop "ReviewScore"}}
                    {{if $score}}
                    <span class="review-score">{{printf "%.1f" $score}}</span>
                    {{$counts := index $prop "Counts"}}
                    {{if $counts}}
                    <span class="review-count">({{index $counts "Reviews"}} Reviews)</span>
                    {{end}}
                    {{end}}
                </div>
            </div>
            <span class="property-type">{{index $prop "PropertyType"}}</span>
        </div>

        <div class="property-title">
            <a title="{{index $prop "PropertyName"}}"
                target="_blank"
                href="https://presto:TRAV3LA1@ownerdirect.beta.123presto.com/property/{{index $prop "PropertySlug"}}/{{$id}}"
                class="pt-property-title">
                {{index $prop "PropertyName"}}
            </a>
        </div>

        <div class="property-info">
            <ul class="pt-amenities ellipsis">
                {{$amenities := index $prop "TopAmenities"}}
                {{range $amenities}}
                <li class="pt-am-item" title="{{.}}">{{.}}</li>
                {{end}}
            </ul>
        </div>

        <div class="property-location">
            <ul class="pt-breadcrumbs">
                {{$categories := index $geo "Categories"}}
                {{range $categories}}
                <li><a href="/all/{{index . "Slug"}}" target="_blank">{{index . "Name"}}</a></li>
                {{end}}
            </ul>
        </div>

        <div class="property-bottom">
            <div class="property-brand">
                {{$url := index $partner "URL"}}
                <a rel="nofollow" class="pt-logo-wrap" href="#" target="_blank"
                    onmouseenter="buildRedirectUrl(this)"
                    onclick="redirectToPartner(this); return false;">
                    {{if $url}}
                        {{if contains $url "booking.com"}}
                        <img loading="lazy" src="/static/img/partners-logo/booking.svg" height="14" width="80" class="pt-partner-logo" alt="Booking.com">
                        {{else if contains $url "expedia.com"}}
                        <img loading="lazy" src="/static/img/partners-logo/expedia_v2.svg" height="14" width="80" class="pt-partner-logo" alt="Expedia">
                        {{else}}
                        <img loading="lazy" src="/static/img/partners-logo/vrbo.svg" height="14" width="80" class="pt-partner-logo" alt="VRBO">
                        {{end}}
                    {{end}}
                </a>
            </div>
            <a rel="nofollow" target="_blank"
                class="availability-button pt-availability js-view-availability"
                href="#"
                onmouseenter="buildRedirectUrl(this)"
                onclick="redirectToPartner(this); return false;">
                View Availability
            </a>
        </div>

    </div>
</div>
{{end}}
{{end}}