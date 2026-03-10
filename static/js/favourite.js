function toggleFavourite(btn) {
    const id = btn.dataset.id;
    const icon = btn.querySelector('.heart-icon');
    
    let favourites = JSON.parse(localStorage.getItem('favourite_list') || '{}');
    
    if (favourites[id]) {
        delete favourites[id];
        icon.setAttribute('fill', 'none');
        icon.setAttribute('stroke', 'white');
    } else {
        favourites[id] = Date.now();
        icon.setAttribute('fill', 'red');
        icon.setAttribute('stroke', 'red');
    }
    
    localStorage.setItem('favourite_list', JSON.stringify(favourites));
    console.log("favourite_list", favourites);
}