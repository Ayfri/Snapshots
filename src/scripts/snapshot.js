/**
 * @param {Snapshot} snapshot
 * @param {HTMLDivElement} div
 */
export function generateSnapshotCard(snapshot, div) {
	const snapshotCard = document.createElement('div');
	snapshotCard.className = 'snapshot-card';

	const snapshotCardTitle = document.createElement('h3');
	snapshotCardTitle.innerText = snapshot.name;
	snapshotCard.appendChild(snapshotCardTitle);

	const snapshotCardDownloadButton = document.createElement('button');
	snapshotCardDownloadButton.innerText = 'Download';
	snapshotCardDownloadButton.addEventListener('click', () => snapshot.download());
	snapshotCard.appendChild(snapshotCardDownloadButton);

	div.appendChild(snapshotCard);
}

/**
 * @param {Snapshot} snapshot
 */
export function generateSnapshotPage(snapshot) {
	const snapshotPage = document.createElement('div');
	snapshotPage.className = 'snapshot-page';
	snapshotPage.appendChild(generateSnapshotTitle(snapshot));
	snapshotPage.appendChild(generateSnapshotBody(snapshot));

	document.getElementsByClassName('content')[0].append(snapshotPage);
}

/**
 * @param {Snapshot} snapshot
 * @returns {HTMLDivElement}
 */
function generateSnapshotTitle(snapshot) {
	const snapshotTitleDiv = document.createElement('div');
	snapshotTitleDiv.className = 'snapshot-title';
	const snapshotTitle = document.createElement('h1');
	snapshotTitle.textContent = snapshot.name;
	snapshotTitleDiv.appendChild(snapshotTitle);

	return snapshotTitleDiv;
}

/**
 * @param {Snapshot} snapshot
 * @returns {HTMLDivElement}
 */
function generateSnapshotBody(snapshot) {
	const snapshotBodyDiv = document.createElement('div');
	snapshotBodyDiv.className = 'snapshot-body';
	const snapshotBody = document.createElement('p');
	snapshotBody.textContent = snapshot.description;
	snapshotBodyDiv.appendChild(snapshotBody);
	const snapshotDate = document.createElement('p');
	snapshotDate.textContent = dayjs(snapshot.date).format('L');
	snapshotDate.className = 'snapshot-date';
	snapshotBodyDiv.appendChild(snapshotDate);

	return snapshotBodyDiv;
}

export class Snapshot {
	/**
	 * @type {string}
	 */
	name;
	/**
	 * @type {Date}
	 */
	date;
	/**
	 * @type {string}
	 */
	description;
	/**
	 * @type {string}
	 */
	url;

	/**
	 * @param {string} name
	 * @param {Date} date
	 * @param {string} description
	 * @param {string} url
	 */
	constructor(name, date, description, url) {
		this.name = name;
		this.date = date;
		this.description = description;
		this.url = url;
	}

	/**
	 * @param {Object} json
	 * @returns {Snapshot}
	 */
	static getFromJSON(json) {
		return new Snapshot(json.name, new Date(json.date), json.description, json.url);
	}

	download() {}
}
