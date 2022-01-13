export * from './snapshot.js';
export * from './version.js';
import {fixDropDowns} from './CSSFixes.js';

document.addEventListener('DOMContentLoaded', onLoad);
dayjs.extend(window.dayjs_plugin_localizedFormat);

function onLoad() {
	fixDropDowns();
	/**
	 * @type {HTMLButtonElement}
	 */
	const button = document.querySelector('#generate');
	button.click();
	button.onclick = null;
}
