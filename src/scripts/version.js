import {MINECRAFT_WIKI_LINK} from './constants.js';
import {fixGridCount} from './CSSFixes.js';
import images from '../resources/images.json' assert {type: 'json'}
import {generateSnapshotCard, Snapshot} from './snapshot.js';

/**
 * @param {Version} version
 */
export async function generateVersionPage(version) {
	const content = document.getElementsByClassName('content')[0];
	const versionPage = document.createElement('div');
	versionPage.className = 'version-page';

	const versionHeader = document.createElement('div');
	versionHeader.className = 'version-header';
	const title = document.createElement('h1');
	title.textContent = version.name;
	versionHeader.appendChild(title);
	versionPage.appendChild(versionHeader);

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
	versionImage.src = await fetch(`../resources/images/${version.imageUrl}`).then(r => r.url);
	versionImage.className = 'version-image';
	versionImage.alt = `${version.name} image`;
	versionImageDiv.appendChild(versionImage);


	const versionDescription = document.createElement('p');
	versionDescription.className = 'version-description';
	versionDescription.innerHTML = version.description.replace(/\. (<sup>.*?<\/sup>)?/g, ".$1<br>");
	[...versionDescription.querySelectorAll('a')]
		.filter(a => a.href.startsWith('/wiki/'))
        .forEach(a => a.href = MINECRAFT_WIKI_LINK + a.getAttribute('href'));
	versionImageAndDescription.appendChild(versionDescription);

	const versionImportantDescription = document.createElement('p');
	versionImportantDescription.className = 'version-important-description';
	versionImportantDescription.textContent = version.importantDescription;

	versionImageDiv.appendChild(versionImportantDescription);

	const versionSnapshots = document.createElement('div');
	versionSnapshots.className = 'version-snapshots';
	versionBody.appendChild(versionSnapshots);

	version.snapshots.forEach(snapshot => generateSnapshotCard(snapshot, versionSnapshots));

	content.appendChild(versionPage);
	fixGridCount();
	window.addEventListener('resize', () => fixGridCount());
}

export class Version extends Snapshot {
	/**
	 * @type {string}
	 */
	imageUrl;
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
		super(name, date, description, url);
		this.imageUrl = imageUrl;
		this.importantDescription = importantDescription;
		this.snapshots = snapshots;
	}

	/**
	 * @param {Version} json
	 * @returns {Version}
	 */
	static getFromJSON(json) {
		const description = json.description;

		let descriptionElement = document.createElement('p');
		descriptionElement.innerHTML = description;
		let importantDescription = descriptionElement.innerText.split(/\. [A-Z]/).slice(0, 2).join('\n') + '.';
		if (importantDescription.length > 100) importantDescription = descriptionElement.innerText.split(/\. [A-Z]/)[0] + '.';

		const imageUrl = Object.entries(images).find(([key, value], _) => json.name.toLowerCase().includes(key[0]))[1];

		const snapshots = json.snapshots.map(s => Snapshot.getFromJSON(s));
		return new Version(json.name, new Date(json.releaseTime), description, json.url, imageUrl, importantDescription, snapshots);
	}
}
