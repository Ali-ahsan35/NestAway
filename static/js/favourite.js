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

function syncFavouriteIcons() {
    const favourites = JSON.parse(localStorage.getItem('favourite_list') || '{}');
    Object.keys(favourites).forEach(id => {
        const btn = document.querySelector('.heart-btn[data-id="' + id + '"]');
        if (!btn) return;
        const icon = btn.querySelector('.heart-icon');
        if (!icon) return;
        icon.setAttribute('fill', 'red');
        icon.setAttribute('stroke', 'red');
    });
}

document.addEventListener('DOMContentLoaded', function () {
    syncFavouriteIcons();
    const grid = document.getElementById('grid');
    if (!grid || typeof MutationObserver !== 'function') {
        return;
    }

    const observer = new MutationObserver(function () {
        syncFavouriteIcons();
    });
    observer.observe(grid, { childList: true, subtree: true });
});