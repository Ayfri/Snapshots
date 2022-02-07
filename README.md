![GitHub Workflow Status](https://img.shields.io/github/workflow/status/Ayfri/minecraft-wiki-scrapper/Update%20versions)

# VersionCraft

TLDR : **VersionCraft is a website listing all versions of minecraft.**

This repo is the source of the WebSite VersionCraft, not hosted for now. It is an exhaustive list of all versions of
minecraft, starting from pre-classic to the latest snapshots. Updated automatically every week, it is always up-to-date.

If a version is missing or some wrong information are present, please make an issue.

This project was made for a school project but will be maintained by me and maybe later published.

## How information is gathered

The information is gathered from the official website of minecraft using
a [separate repo](https://github.com/Ayfri/minecraft-wiki-scrapper) that scrapes the fandom wiki. Then the information
is parsed and cleaned, and finally the information is stored in a JSON file on GitHub. The script is re-run every week
to update the JSON file.

When you load the website, some JavaScript scripts will fetch the JSON file and display the information.

## How to use the website

Go to the website on your browser, via localhost or internet (if published, not available for now). Then select the
version you want to see via the dropdown list.

There you can select the snapshot or pre-release or other version of minecraft you want to see. You can download the
version also, a tutorial is available on the main page.

## Compatibility

The website has been tested on Chrome, Firefox, Brave and Opera. For mobile, it has been tested on Android Brave.<br>
We do not have exact version numbers but the browsers versions where from Q3 2021. The responsive design is imperfect,
but it is a good compromise, font may be a bit tiny.

## Contributors

- [Ayfri](https://github.com/Ayfri)
- [Samydj0](https://github.com/Samydj0)