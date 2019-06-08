# Gimme Quote

Get your dose of famous quotes.

## Development

This project makes use of Zeit Now, as serverless platform.

### Local Development

Using Now CLI, the following command starts a local development environment, which replicates the production environment as closely as possible: 

```bash
now dev
```

### Deployment

#### Automatic Deployments with Git

Any changes pushed to master are automatically deployed by Now.

#### Deploying from Terminal

Using Now CLI, the following command deploys to a staging environment.

```bash
now
```

For deployment to production, use the following command:

```bash
now --target production
```
