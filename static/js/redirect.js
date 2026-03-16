function buildRedirectUrl(btn) {
  const card = btn.closest(".sp-property-card");
  const propertyId = card.dataset.property_id;
  const feed = card.dataset.feed;
  const published = card.dataset.published;
  const upat = card.dataset.upat;
  const lat = card.dataset.lat;
  const lng = card.dataset.lng;
  const type = card.dataset.type;
  const destId = card.dataset.dest_id;
  const ownerId = card.dataset.owner_id;
  const directUrl = card.dataset.direct_url;
  const display = card.dataset.display;
  const city = card.dataset.city;
  const country = card.dataset.country;
  const countryCode = card.dataset.country_code;
  const epc = card.dataset.epc;
  const eplId = card.dataset.eplid;
  const index = card.dataset.index;
  const menuId = Date.now();
  const referralId = index + "-" + propertyId;

  const params = new URLSearchParams({
    menu_id: menuId,
    lang: "en-US",
    feed: feed,
    property_id: propertyId,
    published: published,
    upat: upat,
    latitude: lat,
    longitude: lng,
    type: type.toLowerCase(),
    referrer_page: "category",
    guests: "2",
    dest_id: destId,
    owner_id: ownerId,
    direct_redirect_url: directUrl,
    search_string: display,
    epc: epc,
    referral_id: referralId,
    pl: city,
    pc: country,
    pcc: countryCode,
    eplId: eplId,
    user_type: "dd",
    currency: "BDT",
  });

  btn.href = "http://localhost:8080/redirect-partner?" + params.toString();
}
