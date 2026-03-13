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
                console.log("Submit clicked, checkin:", window.checkin, "checkout:", window.checkout);
                if (window.loadProperties && window.currentCategory) {
                    const sortEl = document.getElementById('sort-properties');
                    const currentOrder = (sortEl && sortEl.value) ? sortEl.value : '1';
                    window.loadProperties(window.currentCategory, currentOrder, {
                        checkin:  window.checkin  || '',
                        checkout: window.checkout || ''
                    });
                }
            }
        }
    });

    // --- INSTANCE 2: Modal datepicker (inside filter modal) ---
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
            submitButton: false,  // ← no submit button
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