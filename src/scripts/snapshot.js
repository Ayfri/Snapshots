export class Snapshot {
	/**
	 * @param {string} name
	 */
	name;
	/**
	 * @param {Date} date
	 */
	date;
	/**
	 * @param {string} description
	 */
	description;
	/**
	 * @param {string} url
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
}

export class Version extends Snapshot {
	/**
	 * @param {string} image
	 */
	imageUrl;
	/**
	 * @param {string} importantDescription
	 */
	importantDescription;

	/**
	 *
	 * @param {string} name
	 * @param {Date} date
	 * @param {string} description
	 * @param {string} url
	 * @param {string} imageUrl
	 * @param {string} importantDescription
	 */
	constructor(name, date, description, url, imageUrl, importantDescription) {
		super(name, date, description, url);
		this.imageUrl = imageUrl;
		this.importantDescription = importantDescription;
	}
}
