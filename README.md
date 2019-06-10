# Gimme Quote

Get your dose of famous quotes - [gimme-quote.robinmitra.now.sh](https://gimme-quote.robinmitra.now.sh/)

## Development

This project has been organised as a mono-repo, and includes the source code for the static site
and the API. The project makes use of Zeit Now, a serverless platform.

### Local Development

Using Now CLI, the following command starts a local development environment, which replicates the
production environment as closely as possible: 

```bash
now dev
```

### Deployment

#### Automatic Deployments with Git

Any changes pushed to master are automatically deployed by Now.

Changes to other branches (include pull requests) are also deployed automatically by Now to 
unique staging aliases.

#### Deploying from Terminal

Using Now CLI, the following command deploys to a staging environment.

```bash
now
```

For deployment to production, use the following command:

```bash
now --target production
```
