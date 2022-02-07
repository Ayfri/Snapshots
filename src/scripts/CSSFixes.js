export const fixDropDowns = () => {
	const dropdown = document.querySelectorAll('.dropdown');
	[...dropdown].forEach(element => {
		if (element.querySelectorAll('.dropdown-button')) {
			const button = element.querySelector('.dropdown-button');
			button.addEventListener('mouseover', () => {
				const list = element.querySelector('.dropdown-content');
				list.style.width = `${button.offsetWidth}px`;
			});
		}
	});
};

export const fixGridCount = () => {
	/**
	 * @type {HTMLDivElement}
	 */
	const grid = document.querySelector('.version-snapshots-list');
	grid.style.gridTemplateColumns = `repeat(${Math.floor(grid.offsetWidth / 250)}, minmax(250px, 1fr))`;
};
