export * from './snapshot.js';
export * from './version.js';
import {fixDropDowns} from './CSSFixes.js';

document.addEventListener('DOMContentLoaded', fixDropDowns);
dayjs.extend(window.dayjs_plugin_localizedFormat);
