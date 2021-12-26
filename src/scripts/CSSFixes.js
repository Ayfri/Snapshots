export const fixDropDowns = () => {
	const dropdown = document.getElementsByClassName('dropdown');
	[...dropdown].forEach(element => {
		if (element.getElementsByClassName('dropdown-button')) {
			const button = element.getElementsByClassName('dropdown-button')[0];
			button.addEventListener('mouseover', () => {
				const list = element.getElementsByClassName('dropdown-content')[0];
				list.style.width = `${button.offsetWidth}px`;
			});
		}
	});
};
