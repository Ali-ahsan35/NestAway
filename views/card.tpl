{{range .Items}}
{{$item := .}}
{{$prop := index $item "Property"}}
{{$geo := index $item "GeoInfo"}}
{{$partner := index $item "Partner"}}
{{$id := index $item "ID"}}

{{if $prop}}
<div class="sp-property-card" data-property_id="{{$id}}">

    <!-- Image -->
    <div class="image-section relative">
        <div class="tiles-icons absolute">
            <div class="tiles-btn fav-icon heart-btn" data-id="{{$id}}" title="Bookmark">♡</div>
        </div>
        <a rel="nofollow" target="_blank" class="sp-property-image" href="{{index $partner "URL"}}">
            {{$imgName := index $prop "FeatureImage"}}
            {{if $imgName}}
            <img class="featured-image pt-featured-image"
                src="https://imgservice.smartours.com/600x600/{{$imgName}}"
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
                href="{{index $partner "URL"}}"
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
                <a rel="nofollow" class="pt-logo-wrap" href="{{$url}}" target="_blank">
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
                class="availability-button pt-availability"
                href="{{index $partner "URL"}}">
                View Availability
            </a>
        </div>

    </div>
</div>
{{end}}
{{end}}