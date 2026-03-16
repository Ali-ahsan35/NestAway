# NestAway — Vacation Rental Search Engine

> A blazing-fast vacation rental discovery platform built with Go (Beego), inspired by Smartours. Search, filter, and explore thousands of vacation rentals worldwide — powered by real-time data from top booking partners like Expedia, VRBO, and Booking.com.

---

## Features

- **Location-based search** — Browse rentals by country, state, city, and neighborhood
- **Smart filtering** — Filter by price, guests, amenities, eco-friendly, and dates
- **Date picker** — Interactive dual-month calendar with availability filtering
- **Favourites** — Save properties locally with persistent localStorage
- **Sort & order** — Sort by popularity, price, and rating
- **Affiliate tracking** — Partner redirects with full tracking parameters
- **Breadcrumb navigation** — SEO-friendly hierarchical location URLs
- **Section-based layout** — Curated property sections per destination
- **Currency support** — BDT/USD conversion with live filter sync
- **Responsive grid** — 4-column layout adapting to all screen sizes

---

## Tech Stack

| Layer | Technology |
|-------|-----------|
| Backend | Go 1.21+ with [Beego v2](https://beego.vip/) |
| Templating | Beego HTML Templates |
| Frontend | Vanilla JS, CSS (CDN-powered) |
| Datepicker | [Hotel Datepicker](https://hoteldatepicker.org/) |
| Styling | 123presto CDN CSS + custom variables |
| Data | 123presto API (proxy architecture) |

---

## Project Structure

```
nestaway/
├── main.go                          # App entry point + template functions
├── routers/
│   └── router.go                    # All route definitions
├── controllers/
│   ├── all.go                       # /all/* — category page
│   ├── categorydetails.go           # /api/v1/category/details/*
│   ├── properties.go                # /api/properties
│   ├── propertydetails.go           # /api/propertydetails
│   ├── refine.go                    # /refine — search results page
│   └── redirectpartner.go           # /redirect-partner — affiliate redirect
├── views/
│   ├── all.tpl                      # Category page template
│   ├── refine.tpl                   # Search/filter page template
│   └── card.tpl                     # Reusable property card component
└── static/
    ├── css/
    │   └── variables.css            # CSS custom properties
    ├── js/
    │   ├── refine.js                # Property loading + rendering
    │   ├── sort.js                  # Sort dropdown logic
    │   ├── modal.js                 # Filter modal (price/guests/amenities)
    │   ├── datepicker.js            # Dual datepicker instances
    │   ├── favourite.js             # Heart toggle + localStorage
    │   └── redirect.js              # Affiliate URL builder
    └── img/
        └── partners-logo/           # Booking, Expedia, VRBO logos
```

---

## Getting Started

### Prerequisites

Make sure you have the following installed:

- [Go 1.21+](https://go.dev/dl/)
- [Bee CLI tool](https://github.com/beego/bee)

```bash
# Install Bee CLI
go install github.com/beego/bee/v2@latest
```

---

### Installation

**1. Clone the repository**

```bash
git clone https://github.com/yourusername/nestaway.git
cd nestaway
```

**2. Install dependencies**

```bash
go mod tidy
```

**3. Run the development server**

```bash
bee run
```

The app will be available at **http://localhost:8080**

---

## URL Structure

| URL | Description |
|-----|-------------|
| `/refine?search=Barcelona, Spain` | Search results page |
| `/all/spain` | All rentals in Spain |
| `/all/spain/catalonia` | All rentals in Catalonia |
| `/all/spain/catalonia/barcelona` | All rentals in Barcelona |
| `/api/v1/category/details/spain/catalonia/barcelona` | Category details API |
| `/api/properties?category=spain&order=1` | Properties list API |
| `/redirect-partner?...` | Affiliate tracking redirect |

---

## Filter URL Parameters

Filters are stored in the URL for shareability and browser history support:

```
/refine?search=France&order=1&amenities=1-2-7&amount=43344-122096&selectedCurrency=BDT&pax=2&dateStart=2026-04-14&dateEnd=2026-04-16
```

| Parameter | Description | Example |
|-----------|-------------|---------|
| `search` | Location keyword | `Barcelona, Spain` |
| `order` | Sort order (1-5) | `1` = Most Popular |
| `amenities` | Amenity IDs (dash-separated) | `1-2-7` |
| `amount` | Price range in BDT | `43344-122096` |
| `selectedCurrency` | Currency code | `BDT` |
| `pax` | Number of guests | `2` |
| `dateStart` | Check-in date | `2026-04-14` |
| `dateEnd` | Check-out date | `2026-04-16` |
| `ecoFriendly` | Eco-friendly filter | `true` |

---

## API Architecture

NestAway acts as a **proxy server** — all external API calls are made server-side from Go controllers, keeping credentials secure and enabling server-side rendering for the category pages.

```
Browser → NestAway (Go) → 123presto API
                       → Ownerdirect API
```

**External API Base URLs:**
- `https://ownerdirect.beta.123presto.com` — category details, property data
- `https://travel-bangladesh.beta.123presto.com` — property search & filtering

---

## Sort Options

| Value | Label |
|-------|-------|
| `1` | Most Popular |
| `2` | Lowest Price |
| `3` | Highest Price |
| `4` | Lowest Rating |
| `5` | Highest Rating |

---

## Key Implementation Notes

- **Currency conversion** — Prices displayed in BDT (×120 from USD). URL stores BDT, API receives USD.
- **Slug format** — External API uses `:` separator (`spain:catalonia:barcelona`), URLs use `/` — conversion handled transparently in controllers.
- **Template caching** — Beego caches templates in production. Run `bee run` to auto-reload during development.
- **ID limit** — Property detail API is limited to 72 IDs per request to avoid URL length limits.
- **Affiliate tracking** — Each property click generates a unique `menu_id` (timestamp) and `referral_id` (index + property ID).

---

## Partner Integrations

| Partner | Feed ID | Logo |
|---------|---------|------|
| Booking.com | 11 | `booking.svg` |
| VRBO | 12 | `vrbo.svg` |
| Expedia | 24 | `expedia_v2.svg` |

---

## Roadmap

- [ ] Property detail page (`/property/:slug/:id`)
- [ ] Pagination
- [ ] Homepage with search bar
- [ ] Mobile responsive design
- [ ] Map view integration
- [ ] Multi-currency support

---



<p align="center">Built with ❤️ using Go + Beego</p>
