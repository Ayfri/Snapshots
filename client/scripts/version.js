import {MINECRAFT_WIKI_LINK} from './constants.js';
import {fixGridCount} from './CSSFixes.js';
import images from '../resources/images.json' assert {type: 'json'};
import {generateSnapshotCard, Snapshot} from './snapshot.js';

function generateVersionBody(versionPage, version) {
	const versionBody = document.createElement('div');
	versionBody.className = 'version-body';
	versionPage.appendChild(versionBody);

	const versionImageAndDescription = document.createElement('div');
	versionImageAndDescription.className = 'version-image-and-description';
	versionBody.appendChild(versionImageAndDescription);

	const versionImageDiv = document.createElement('div');
	versionImageDiv.className = 'version-image-div';
	versionImageAndDescription.appendChild(versionImageDiv);

	const versionImage = document.createElement('img');
	versionImage.src = version.imageUrl;
	versionImage.className = 'version-image';
	versionImage.alt = `${version.name} image`;
	versionImageDiv.appendChild(versionImage);

	const versionDescription = document.createElement('p');
	versionDescription.className = 'version-description';
	versionDescription.innerHTML = version.description.replace(/\. (<sup>.*?<\/sup>)?/g, '.$1<br>');
	[...versionDescription.querySelectorAll('a')]
		.filter(a => a.getAttribute('href').startsWith('/wiki/'))
		.forEach(a => (a.href = MINECRAFT_WIKI_LINK + a.getAttribute('href')));
	versionImageAndDescription.appendChild(versionDescription);

	const versionImportantDescription = document.createElement('p');
	versionImportantDescription.className = 'version-important-description';
	versionImportantDescription.textContent = version.importantDescription;

	versionImageDiv.appendChild(versionImportantDescription);

	const versionSnapshots = document.createElement('div');
	versionSnapshots.className = 'version-snapshots';
	versionBody.appendChild(versionSnapshots);

	generateVersionListButtons(version, versionSnapshots);

	const versionSnapshotsList = document.createElement('div');
	versionSnapshotsList.className = 'version-snapshots-list';
	versionSnapshots.appendChild(versionSnapshotsList);
	versionBody.appendChild(versionSnapshots);
}

/**
 * @param {HTMLElement} div
 * @param {string} name
 * @param {(event: Event) => void} onChange
 * @param {boolean} checked
 */
export function createCheckBox(div, name, onChange = () => {}, checked = true) {
	const checkBox = document.createElement('input');
	checkBox.type = 'checkbox';
	checkBox.className = 'version-checkbox';
	checkBox.checked = checked;
	checkBox.onchange = onChange;

	const checkBoxLabel = document.createElement('label');
	checkBoxLabel.className = 'version-checkbox-label';
	checkBoxLabel.textContent = name;
	checkBoxLabel.appendChild(checkBox);

	div.appendChild(checkBoxLabel);
}

/**
 * @param {Version} version
 * @param {HTMLDivElement} div
 */
export function generateVersionListButtons(version, div) {
	const filters = document.createElement('form');
	filters.className = 'version-filters';
	div.appendChild(filters);

	createCheckBox(filters, `Snapshots (${version.getTypeSize('snapshot')})`, () => {
		const elementNodeListOf = document.querySelectorAll('.version-card-snapshot');
		elementNodeListOf.forEach(card => card.classList.toggle('hidden'));
	});

	createCheckBox(filters, `Pre-Releases (${version.getTypeSize('pre-release')})`, () => {
		const elementNodeListOf = document.querySelectorAll('.version-card-pre-release');
		elementNodeListOf.forEach(card => card.classList.toggle('hidden'));
	});

	createCheckBox(filters, `Release Candidates (${version.getTypeSize('candidate')})`, () => {
		const elementNodeListOf = document.querySelectorAll('.version-card-candidate');
		elementNodeListOf.forEach(card => card.classList.toggle('hidden'));
	});

	createCheckBox(filters, `Releases (${version.getTypeSize('release')})`, () => {
		const elementNodeListOf = document.querySelectorAll('.version-card-release');
		elementNodeListOf.forEach(card => card.classList.toggle('hidden'));
	});
}

/**
 * @param {Version} version
 */
export async function generateVersionPage(version) {
	const content = document.querySelector('.content');
	const versionPage = document.createElement('div');
	versionPage.className = 'version-page';

	const versionHeader = document.createElement('div');
	versionHeader.className = 'version-header';

	const title = document.createElement('h1');
	title.textContent = version.name;
	versionHeader.appendChild(title);
	versionPage.appendChild(versionHeader);

	generateVersionBody(versionPage, version);

	version.snapshots.forEach(snapshot => generateSnapshotCard(snapshot, versionPage.querySelector('.version-snapshots-list')));

	content.appendChild(versionPage);
	document.querySelector('.home').style.display = 'none';
	document.title = version.name;
	fixGridCount();
	window.addEventListener('resize', fixGridCount);
}

export class Version extends Snapshot {
	/**
	 * @type {string}
	 */
	imageUrl;
	/**
	 * @type {string}
	 */
	url;
	/**
	 * @type {string}
	 */
	importantDescription;
	/**
	 * @type {Snapshot[]}
	 */
	snapshots;

	/**
	 * @param {string} name
	 * @param {Date} date
	 * @param {string} description
	 * @param {string} url
	 * @param {string} imageUrl
	 * @param {string} importantDescription
	 * @param {Snapshot[]} snapshots
	 */
	constructor(name, date, description, url, imageUrl, importantDescription, snapshots = []) {
		super(name, date, description, '', '');
		this.imageUrl = imageUrl;
		this.importantDescription = importantDescription;
		this.url = url;
		this.snapshots = snapshots;
	}

	/**
	 * @param {Version} json
	 * @returns {Version}
	 */
	static getFromJSON(json) {
		const description = json.description ?? '';

		const descriptionElement = document.createElement('p');
		descriptionElement.innerHTML = description;
		const text = descriptionElement.innerText.replace(/\.?\[\d]\.? /g, '. ');
		let importantDescription = `${text
			.split(/\. [A-Z]/)
			.slice(0, 2)
			.join('\n')}.`;
		if (importantDescription.length > 100) importantDescription = `${text.replace(/<sup><a>.*?<\/a><\/sup>/g, '').split(/\. [A-Z]/)[0]}.`;

		const imageUrl = Object.entries(images).find(([key, value], _) => json.name.toLowerCase().endsWith(key))[1];
		const snapshots = (json.snapshots ?? []).map(Snapshot.getFromJSON).filter(snapshot => !snapshot.name.toLowerCase().includes('server'));
		return new Version(json.name, new Date(json.releaseTime), description, json.url, imageUrl, importantDescription, snapshots);
	}

	/**
	 * @param {'snapshot'|'pre-release'|'candidate'|'release'}type
	 * @returns {number}
	 */
	getTypeSize(type) {
		return this.snapshots.filter(s => s.snapshotType === type).length;
	}
}
