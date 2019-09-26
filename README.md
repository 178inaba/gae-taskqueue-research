# gae-taskqueue-research

## Deploy

```console
# kicker
$ gcloud app deploy --version v1-0-0

# task
$ gcloud app deploy --version v1-0-0 --no-stop-previous-version

# queue
$ gcloud app deploy queue.yaml

# appcfg
$ appcfg.py --application $PROJECT_ID --version v1-0-0 update app.yaml
```
