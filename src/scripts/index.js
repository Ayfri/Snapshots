import {INPUT_URL} from './constants.js';
import {fixDropDowns} from './CSSFixes.js';
import {createMenu} from './menu.js';
import {assignButtons} from './buttons.js';

document.addEventListener('DOMContentLoaded', onLoad, {once: true});
dayjs.extend(window.dayjs_plugin_localizedFormat);

async function onLoad() {
	const json = await (await fetch(INPUT_URL)).json();
	sessionStorage.setItem('json', JSON.stringify(json));
	createMenu();
	fixDropDowns();
	assignButtons();
}
