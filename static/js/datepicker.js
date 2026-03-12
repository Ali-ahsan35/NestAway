document.addEventListener('DOMContentLoaded', function() {
    const input = document.getElementById('js-dp-input');
    if (!input) return;

    const modal    = document.getElementById('js-datepicker-modal');
    const overlay  = document.getElementById('dp-overlay');
    const closeBtn = document.getElementById('js-dp-close');

    // Move input inside container before initializing
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
        inline: true,  // ← render inline, not as popup
        onSelectRange: function() {
            const val     = picker.getDatepickerValue();
            const parts   = val.split(' - ');
            window.checkin  = parts[0] || '';
            window.checkout = parts[1] || '';

            const btn = document.getElementById('standalone-dp');
            if (btn && window.checkin && window.checkout) {
                btn.textContent = window.checkin + ' → ' + window.checkout;
            }
        }
    });

    console.log("input id:", input.id);
    console.log("looking for:", 'datepicker-' + input.id);
    console.log("found:", document.getElementById('datepicker-' + input.id));

    function openDatepicker() {
        modal.classList.remove('hidden');
    }

    function closeDatepicker() {
        modal.classList.add('hidden');
    }

    document.getElementById('standalone-dp').addEventListener('click', openDatepicker);
    closeBtn.addEventListener('click', closeDatepicker);
    overlay.addEventListener('click', closeDatepicker);

    document.addEventListener('click', function(e) {
        if (e.target && e.target.classList.contains('datepicker__submit-button')) {
            closeDatepicker();
        }
    });
});