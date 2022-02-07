import {MINECRAFT_WIKI_LINK} from './constants.js';

const snapshotBodyLeaveListener = () => {
	document.querySelectorAll('.snapshot-card-hover').forEach(card => card.classList.remove('snapshot-card-hover'));
	document.querySelector('#snapshot-page').style.display = 'none';
};

/**
 * @param {Snapshot} snapshot
 * @param {HTMLDivElement} div
 */
export function generateSnapshotCard(snapshot, div) {
	const snapshotCard = document.createElement('div');
	snapshotCard.className = 'snapshot-card';
	switch (snapshot.type) {
		case 'snapshot':
			snapshotCard.classList.add('version-card-snapshot');
			break;
		case 'pre-release':
			snapshotCard.classList.add('version-card-pre-release');
			break;
		case 'candidate':
			snapshotCard.classList.add('version-card-candidate');
			break;
		default:
			snapshotCard.classList.add('version-card-release');
			break;
	}


	const snapshotCardTitle = document.createElement('h3');
	snapshotCardTitle.innerText = snapshot.name;
	snapshotCard.appendChild(snapshotCardTitle);
	div.appendChild(snapshotCard);

	snapshotCard.addEventListener('mouseenter', async () => {
		let removeHover = true;
		snapshotCard.classList.add('snapshot-card-hover');
		const timer = setTimeout(() => {
			removeHover = false;
			generateSnapshotPage(snapshot, document.body);
		}, 750);

		snapshotCard.onmouseleave = async () => {
			clearTimeout(timer);
			if (removeHover) snapshotCard.classList.remove('snapshot-card-hover');
		};
	});

	snapshotCard.addEventListener('click', () => {
		snapshotCard.classList.add('snapshot-card-hover');
		generateSnapshotPage(snapshot, document.body);
	});
}

/**
 * @param {Snapshot} snapshot
 * @param {HTMLDivElement} div
 */
export function generateSnapshotPage(snapshot, div) {
	let snapshotPage;
	if (document.querySelector('#snapshot-page')) {
		snapshotPage = document.querySelector('#snapshot-page');
		snapshotPage.innerHTML = '';
	} else {
		snapshotPage = document.createElement('div');
		snapshotPage.id = 'snapshot-page';
		div.appendChild(snapshotPage);
	}

	snapshotPage.appendChild(generateSnapshotBody(snapshot));
	document.querySelector('.snapshot-body')?.addEventListener('mouseleave', snapshotBodyLeaveListener);
	snapshotPage.style.display = 'block';
}

/**
 * @param {Snapshot} snapshot
 * @returns {HTMLDivElement}
 */
function generateSnapshotBody(snapshot) {
	const snapshotBodyDiv = document.createElement('div');
	snapshotBodyDiv.className = 'snapshot-body';
	generateExitButton(snapshot, snapshotBodyDiv);

	const snapshotTitle = document.createElement('h1');
	snapshotTitle.textContent = snapshot.name;
	snapshotBodyDiv.appendChild(snapshotTitle);

	const snapshotBody = document.createElement('p');
	snapshotBody.className = 'snapshot-description';
	snapshotBody.innerHTML = snapshot.description;
	snapshotBodyDiv.appendChild(snapshotBody);
	snapshotBody.querySelectorAll('a').forEach(a => (a.href = MINECRAFT_WIKI_LINK + a.getAttribute('href')));

	const snapshotDate = document.createElement('p');
	const parsedDate = dayjs(snapshot.releaseTime);
	snapshotDate.textContent = parsedDate.isValid() ? parsedDate.format('L') : 'Invalid Date or not found.';
	snapshotDate.className = 'snapshot-date';
	snapshotBodyDiv.appendChild(snapshotDate);

	const downloadClientButton = document.createElement('button');
	downloadClientButton.classList.add('snapshot-download', 'download-client', 'btn');
	downloadClientButton.type = 'submit';
	downloadClientButton.setAttribute('onclick', `window.open('${snapshot.downloadClient ?? '#'}')`);
	downloadClientButton.innerText = snapshot.downloadClient ? 'Download Client' : 'Unavailable';
	if (!snapshot.downloadClient) downloadClientButton.disabled = true;
	snapshotBodyDiv.appendChild(downloadClientButton);

	const downloadJSONButton = document.createElement('button');
	downloadJSONButton.classList.add('snapshot-download', 'download-json', 'btn');
	downloadJSONButton.type = 'submit';
	downloadJSONButton.setAttribute('onclick', `window.open('${snapshot.downloadJSON ?? '#'}')`);
	downloadJSONButton.innerText = snapshot.downloadJSON ? 'Download JSON' : 'Unavailable';
	if (!snapshot.downloadJSON) downloadJSONButton.disabled = true;
	snapshotBodyDiv.appendChild(downloadJSONButton);

	return snapshotBodyDiv;
}

/**
 * @param {Snapshot} snapshot snapshot
 * @param {HTMLDivElement} body
 */
function generateExitButton(snapshot, body) {
	const exit = document.createElement('a');
	exit.className = 'exit-button';
	exit.innerHTML = `<svg xmlns="http://www.w3.org/2000/svg" x="0px" y="0px"
width="24" height="24"
viewBox="0 0 24 24"
style=" fill:#000000;"><path d="M 4.7070312 3.2929688 L 3.2929688 4.7070312 L 10.585938 12 L 3.2929688 19.292969 L 4.7070312 20.707031 L 12 13.414062 L 19.292969 20.707031 L 20.707031 19.292969 L 13.414062 12 L 20.707031 4.7070312 L 19.292969 3.2929688 L 12 10.585938 L 4.7070312 3.2929688 z"></path></svg>`;
	body.appendChild(exit);
	exit.title = 'Close';
	exit.addEventListener('click', snapshotBodyLeaveListener);
}

export class Snapshot {
	/**
	 * @type {string}
	 */
	name;
	/**
	 * @type {Date}
	 */
	releaseTime;
	/**
	 * @type {string}
	 */
	description;
	/**
	 * @type {string}
	 */
	downloadClient;
	/**
	 * @type {string}
	 */
	downloadJSON;
	/**
	 * @type {'snapshot'|'pre-release'|'candidate'|'release'}
	 */
	type;

	/**
	 * @param {string} name
	 * @param {Date} date
	 * @param {string} description
	 * @param {string} downloadClient
	 * @param {string} downloadJSON
	 */
	constructor(name, date, description, downloadClient, downloadJSON) {
		this.name = name;
		this.releaseTime = date;
		this.description = description;
		this.downloadClient = downloadClient;
		this.downloadJSON = downloadJSON;

		if (/(\d{1,2}w\d{1,2}[a-z])|(preview)|(test)|(experimental snapshot)/i.test(name)) {
			this.type = 'snapshot';
		} else if (/(1\.\d{1,2}.\d{1,2}-pre)|(1\.\d{1,2}(\.\d{1,2})? pre-?release)/i.test(name)) {
			this.type = 'pre-release';
		} else if (/(1\.\d{1,2}.\d{1,2}-rc)|(RC\d)|(1\.\d{1,2}(\.\d{1,2})? release[- ]candidate)/i.test(name)) {
			this.type = 'candidate';
		} else {
			this.type = 'release';
		}
	}

	/**
	 * @param {Snapshot} json
	 * @returns {Snapshot}
	 */
	static getFromJSON(json) {
		return new Snapshot(json.name, new Date(json.releaseTime * 1000), json.description, json.downloadClient, json.downloadJSON);
	}
}
