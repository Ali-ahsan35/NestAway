document.addEventListener('DOMContentLoaded', function() {
    // --- INSTANCE 1: Standalone datepicker (Dates button) ---
    const input = document.getElementById('js-dp-input');
    if (!input) return;

    const modal    = document.getElementById('js-datepicker-modal');
    const overlay  = document.getElementById('dp-overlay');
    const closeBtn = document.getElementById('js-dp-close');
    const clearBtn = document.getElementById('js-dp-clear');

    const container = document.getElementById('js-dp-container');
    container.appendChild(input);

    const picker = new HotelDatepicker(input, {
        format: 'YYYY-MM-DD',
        infoFormat: 'MMM D, YYYY',
        separator: ' - ',
        startDate: new Date(),
        minNights: 1,
        topbarPosition: 'bottom',
        showTopbar: true,
        submitButton: true,
        submitButtonName: 'Continue',
        inline: true,
        onSelectRange: function() {
            const val    = input.value;
            const parts  = val.split(' - ');
            window.checkin  = parts[0] || '';
            window.checkout = parts[1] || '';

            const btn = document.getElementById('standalone-dp');
            if (btn && window.checkin && window.checkout) {
                btn.textContent = window.checkin + ' → ' + window.checkout;
            }
            // Sync to modal datepicker input
            const modalDpInput = document.getElementById('modal-datepicker');
            if (modalDpInput && window.checkin && window.checkout) {
                modalDpInput.value = window.checkin + ' - ' + window.checkout;
            }
        }
    });

    setTimeout(function() {
        const urlParams = new URLSearchParams(window.location.search);
        const dateStart = urlParams.get('dateStart');
        const dateEnd   = urlParams.get('dateEnd');
        if (dateStart && dateEnd) {
            window.checkin  = dateStart;
            window.checkout = dateEnd;
            const btn = document.getElementById('standalone-dp');
            if (btn) btn.textContent = dateStart + ' → ' + dateEnd;
            const modalDpInput = document.getElementById('modal-datepicker');
            if (modalDpInput) modalDpInput.value = dateStart + ' - ' + dateEnd;
        }
    }, 300);


    function openDatepicker() { modal.classList.remove('hidden'); }
    function closeDatepicker() { modal.classList.add('hidden'); }
    function clearDates() {
        window.checkin  = '';
        window.checkout = '';
        input.value = '';
        const modalDpInput = document.getElementById('modal-datepicker');
        if (modalDpInput) modalDpInput.value = '';
        const btn = document.getElementById('standalone-dp');
        if (btn) btn.textContent = 'Dates';
    }

    document.getElementById('standalone-dp').addEventListener('click', openDatepicker);
    closeBtn.addEventListener('click', closeDatepicker);
    overlay.addEventListener('click', closeDatepicker);
    if (clearBtn) clearBtn.addEventListener('click', clearDates);

    // Submit → redirect
    document.addEventListener('click', function(e) {
        if (e.target && (e.target.classList.contains('datepicker__submit-button') ||
            e.target.closest('.datepicker__submit-button'))) {
            if (modal && !modal.classList.contains('hidden')) {
                closeDatepicker();
                if (window.loadProperties && window.currentCategory) {
                    const sortEl = document.getElementById('sort-properties');
                    const currentOrder = (sortEl && sortEl.value) ? sortEl.value : '1';

                    // Read existing filters from URL and merge with new dates
                    const urlParams = new URLSearchParams(window.location.search);
                    const mergedFilters = {};
                    if (urlParams.get('amenities')) {
                        mergedFilters.amenities = urlParams.get('amenities').split('-');
                    }
                    if (urlParams.get('ecoFriendly')) {
                        mergedFilters.ecoFriendly = true;
                    }
                    if (urlParams.get('amount')) {
                        const parts = urlParams.get('amount').split('-');
                        const minUSD = Math.round(parseInt(parts[0]) / 120);
                        const maxUSD = Math.round(parseInt(parts[1]) / 120);
                        mergedFilters.amount    = minUSD + '-' + maxUSD;
                        mergedFilters.amountBDT = urlParams.get('amount');
                    }
                    if (urlParams.get('pax')) {
                        mergedFilters.guests = parseInt(urlParams.get('pax'));
                    }
                    // Override dates with new selection
                    mergedFilters.checkin  = window.checkin  || '';
                    mergedFilters.checkout = window.checkout || '';

                    window.loadProperties(window.currentCategory, currentOrder, mergedFilters);
                }
            }
        }
    });

    // INSTANCE 2: Modal datepicker 
    const modalDpInput = document.getElementById('modal-datepicker');
    const calendarIcon = document.querySelector('#modal-dp .calendar-icon');

    if (modalDpInput) {
        const modalPicker = new HotelDatepicker(modalDpInput, {
            format: 'YYYY-MM-DD',
            infoFormat: 'MMM D, YYYY',
            separator: ' - ',
            startDate: new Date(),
            minNights: 1,
            topbarPosition: 'bottom',
            showTopbar: true,
            submitButton: false,
            onSelectRange: function() {
                const val   = modalDpInput.value;
                const parts = val.split(' - ');
                window.checkin  = parts[0] || '';
                window.checkout = parts[1] || '';

                // Sync to standalone button
                const btn = document.getElementById('standalone-dp');
                if (btn && window.checkin && window.checkout) {
                    btn.textContent = window.checkin + ' → ' + window.checkout;
                }
                // Sync to standalone input
                input.value = val;
            }
        });

        if (calendarIcon) {
            calendarIcon.addEventListener('click', function() {
                modalDpInput.focus();
                modalDpInput.click();
            });
        }
    }
});