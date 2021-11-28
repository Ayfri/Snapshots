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
	 * @param {Object} json
	 * @returns {Version}
	 */
	static getFromJSON(json) {
		const importantDescription = json.description.split('\n').slice(0, 4).join('\n');
		const snapshots = json.snapshots.map(s => Snapshot.getFromJSON(s));
		return new Version(json.name, new Date(json.date), json.description, json.url, json.imageUrl, importantDescription, snapshots);
	}
}
