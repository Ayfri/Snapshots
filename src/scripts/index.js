export * from './snapshot.js';
export * from './version.js';
import {INPUT_URL} from './constants.js';
import {fixDropDowns} from './CSSFixes.js';
import {createMenu} from './menu.js';
import {assignButtons} from './buttons.js';

/**
 * @type {Version[]}
 */
export let json;

document.addEventListener('DOMContentLoaded', onLoad, {once: true});
dayjs.extend(window.dayjs_plugin_localizedFormat);

export async function onLoad() {
	json = await (await fetch(INPUT_URL)).json();
	createMenu();
	fixDropDowns();
	assignButtons();
}