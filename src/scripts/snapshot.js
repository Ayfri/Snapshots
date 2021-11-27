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

	constructor(name, date, description, url, image, importantDescription) {
		super(name, date, description, url);
		this.imageUrl = image;
		this.importantDescription = importantDescription;
	}
}
