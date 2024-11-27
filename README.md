# Postbidship Scrapper

This is a scrapper for the postbidship website. It is a simple scrapper that scrapes the postbidship website and uploads the data to a google sheet.

To work it requires following environment variables:

Google Drive:

- `GOOGLE_DRIVE_CLIENT_ID`
- `GOOGLE_DRIVE_CLIENT_SECRET`
- `GOOGLE_DRIVE_API_KEY`
- `GOOGLE_DRIVE_PARENT_FOLDER_ID`

Postbidship:

- `EMAIL`
- `PASSWORD`

This project uses [task](https://taskfile.dev) as a task runner.

To build and run the scrapper, you can use the following command:

```bash
task run
```
