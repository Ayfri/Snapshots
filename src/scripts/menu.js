import {menuList, names} from './constants.js';
import {generateVersionPage, json, Version} from './index.js';

export function createMenu() {
	const navbar = document.querySelector('.navbar');

	menuList.forEach((dropdown, index) => {
		const dropdownMenu = document.createElement('div');
		dropdownMenu.classList.add('dropdown');

		const button = document.createElement('button');
		button.classList.add('dropdown-button');
		button.innerText = names[index];

		const dropdownItem = document.createElement('div');
		dropdownItem.classList.add('dropdown-content');
		json.forEach(version => {
			dropdown.forEach(item => {
				if (version.name.toLowerCase().endsWith(item.replace('1.0', '1.0.0'))) {
					const item = document.createElement('a');
					item.href = '#';
					item.innerText = version.name.replace('1.0.0', '1.0');
					if ([...dropdownItem.children].find(e => e.innerText.endsWith(item.innerText))) return;
					dropdownItem.appendChild(item);

					item.onclick = async () => {
						const ver = json.find(v => v.name === version.name);
						const content = document.querySelector('.content');
						content.innerHTML = '';
						await generateVersionPage(Version.getFromJSON(ver));
					};
				}
			});
		});
		dropdownMenu.appendChild(button);
		dropdownMenu.appendChild(dropdownItem);
		navbar.appendChild(dropdownMenu);
	});
}
