export function assignButtons() {
	const links = document.querySelector('.footer-links').children;
	links[1].addEventListener('click', randomVersion);
	links[2].addEventListener('click', latestVersion);
}

function randomVersion() {
	const menu = [...document.querySelectorAll('.dropdown-content')].flatMap( e => [...e.children]);
	const random = Math.floor(Math.random() * menu.length);
	menu[random].click();
}

function latestVersion() {
	const menu = [...document.querySelectorAll('.dropdown-content')]
		.flatMap(e => [...e.children])
		.filter(e => e.parentElement.parentElement.querySelector('button').innerText !== 'Others');
	menu[menu.length - 1].click();
}