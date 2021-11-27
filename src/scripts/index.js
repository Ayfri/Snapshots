export * from './snapshot.js';


/**
 * @param {Snapshot} snapshot
 */
export function generateSnapshotPage(snapshot) {
	const snapshotPage = document.createElement('div');
	snapshotPage.className = 'snapshot-page';
	snapshotPage.appendChild(generateSnapshotTitle(snapshot));
	snapshotPage.appendChild(generateSnapshotBody(snapshot));

	document.getElementsByClassName("content")[0].append(snapshotPage);
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
	snapshotDate.textContent = snapshot.date;
	snapshotDate.className = 'snapshot-date';
	snapshotBodyDiv.appendChild(snapshotDate);
	return snapshotBodyDiv;
}
