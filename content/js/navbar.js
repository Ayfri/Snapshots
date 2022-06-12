window.addEventListener('load', () => {
	const openNavbarBtn = document.querySelector('#show-navbar');
	openNavbarBtn.addEventListener('change', () => {
		const navbar = document.querySelector('.navbar-menu');
		if (openNavbarBtn.checked) {
			navbar.classList.remove('hidden');
		} else {
			navbar.classList.add('hidden');
		}
	});
});
