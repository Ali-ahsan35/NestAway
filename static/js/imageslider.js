function initImageSlider() {
    document.querySelectorAll('.sp-property-card').forEach(card => {
        const imageSection = card.querySelector('.image-section');
        if (!imageSection) return;

        const nextBtn = document.createElement('button');
        nextBtn.className = 'slider-next-btn';
        nextBtn.innerHTML = '❯';
        nextBtn.style.cssText = `
            display: none;
            position: absolute;
            right: 8px;
            top: 50%;
            transform: translateY(-50%);
            z-index: 20;
            background: rgba(255,255,255,0.9);
            border: none;
            border-radius: 50%;
            width: 32px;
            height: 32px;
            cursor: pointer;
            font-size: 14px;
            font-weight: bold;
        `;
        imageSection.appendChild(nextBtn);

        const prevBtn = document.createElement('button');
        prevBtn.className = 'slider-prev-btn';
        prevBtn.innerHTML = '❮';
        prevBtn.style.cssText = `
            display: none;
            position: absolute;
            left: 8px;
            top: 50%;
            transform: translateY(-50%);
            z-index: 20;
            background: rgba(255,255,255,0.9);
            border: none;
            border-radius: 50%;
            width: 32px;
            height: 32px;
            cursor: pointer;
            font-size: 14px;
            font-weight: bold;
        `;
        imageSection.appendChild(prevBtn);

        // Show arrows on hover
        imageSection.addEventListener('mouseenter', () => {
            if (card.dataset.missingFeatureImage === 'true') {
                nextBtn.style.display = 'none';
                prevBtn.style.display = 'none';
                return;
            }
            if (card.dataset.imagesLoaded === 'true') {
                nextBtn.style.display = 'block';
                const idx = parseInt(card.dataset.currentIndex || '0');
                if (idx > 0) prevBtn.style.display = 'block';
                return;
            }

            if (card.dataset.imagesChecked === 'true') {
                return;
            }

            card.dataset.imagesChecked = 'true';
            const propertyId = card.dataset.property_id;
            fetch('/api/property/images?propertyId=' + propertyId)
                .then(res => res.json())
                .then(data => {
                    const images = data.Images || [];
                    if (images.length === 0) {
                        nextBtn.style.display = 'none';
                        prevBtn.style.display = 'none';
                        return;
                    }

                    card.dataset.images = JSON.stringify(images);
                    card.dataset.currentIndex = '0';
                    card.dataset.imagesLoaded = 'true';
                    nextBtn.style.display = 'block';
                })
                .catch(err => console.error('Image fetch error:', err));
        });
        imageSection.addEventListener('mouseleave', () => {
            nextBtn.style.display = 'none';
            prevBtn.style.display = 'none';
        });

        nextBtn.addEventListener('click', function(e) {
            e.stopPropagation();
            const propertyId = card.dataset.property_id;

            if (card.dataset.fallbackImage === 'true') {
                return;
            }

            if (card.dataset.imagesLoaded === 'true') {
                navigateSlider(card, 1);
                return;
            }

            fetch('/api/property/images?propertyId=' + propertyId)
                .then(res => res.json())
                .then(data => {
                    const images = data.Images || [];
                    if (images.length === 0) {
                        card.dataset.fallbackImage = 'true';
                        nextBtn.style.display = 'none';
                        prevBtn.style.display = 'none';
                        return;
                    }

                    card.dataset.images = JSON.stringify(images);
                    card.dataset.currentIndex = '0';
                    card.dataset.imagesLoaded = 'true';

                    createDots(card, imageSection, images.length);

                    navigateSlider(card, 1);
                })
                .catch(err => console.error('Image fetch error:', err));
        });

        prevBtn.addEventListener('click', function(e) {
            e.stopPropagation();
            navigateSlider(card, -1);
        });
    });
}

function createDots(card, imageSection, count) {
    if (card.querySelector('.slider-dots')) return;

    const dots = document.createElement('div');
    dots.className = 'slider-dots';
    dots.style.cssText = `
        position: absolute;
        bottom: 8px;
        left: 50%;
        transform: translateX(-50%);
        display: flex;
        gap: 5px;
        align-items: center;
        z-index: 20;
        pointer-events: none;
    `;

    for (let i = 0; i < count; i++) {
        const dot = document.createElement('span');
        dot.style.cssText = `
            width: 6px;
            height: 6px;
            border-radius: 50%;
            background: white;
            opacity: 0.6;
            transition: width 0.2s ease, height 0.2s ease, opacity 0.2s ease;
            display: inline-block;
        `;
        dots.appendChild(dot);
    }

    imageSection.appendChild(dots);
}

function navigateSlider(card, direction) {
    const images = JSON.parse(card.dataset.images || '[]');
    if (images.length === 0) return;

    let currentIndex = parseInt(card.dataset.currentIndex || '0');
    currentIndex = (currentIndex + direction + images.length) % images.length;
    card.dataset.currentIndex = currentIndex;

    updateSliderImage(card);

    const prevBtn = card.querySelector('.slider-prev-btn');
    const nextBtn = card.querySelector('.slider-next-btn');
    if (prevBtn) prevBtn.style.display = currentIndex === 0 ? 'none' : 'block';
    if (nextBtn) nextBtn.style.display = 'block';
}

function updateSliderImage(card) {
    const images = JSON.parse(card.dataset.images || '[]');
    const currentIndex = parseInt(card.dataset.currentIndex || '0');
    const imgEl = card.querySelector('.pt-featured-image');

    if (imgEl && images[currentIndex]) {
        imgEl.style.transition = 'opacity 0.3s ease';
        imgEl.style.opacity = '0';
        setTimeout(() => {
            imgEl.src = 'https://imgservice.ownerdirect.com/600x600/' + images[currentIndex];
            imgEl.style.opacity = '1';
        }, 300);
    }

    // Update active dot
    const dots = card.querySelector('.slider-dots');
    if (!dots) return;

    dots.querySelectorAll('span').forEach((dot, i) => {
        if (i === currentIndex) {
            dot.style.width   = '9px';
            dot.style.height  = '9px';
            dot.style.opacity = '1';
        } else {
            dot.style.width   = '6px';
            dot.style.height  = '6px';
            dot.style.opacity = '0.6';
        }
    });
}