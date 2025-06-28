# Project Structure

This document outlines the structure of the Smartik monorepo.

```
smartik/
├───.gitignore
├───lint-staged.config.mjs
├───LICENSE
├───package.json
├───pnpm-lock.yaml
├───pnpm-workspace.yaml
├───prettier.config.mjs
├───README.md
├───turbo.json
├───.husky/
│   ├───.gitignore
│   └───pre-commit
├───.idea/
│   ├───.gitignore
│   └───workspace.xml
├───src/
│   ├───client/
│   │   └───desktop/
│   ├───server/
│   │   ├───api/
│   │   └───core/
│   └───shared/
│       └───.gitkeep
└───turbo/
```

## Top-Level Directories

- **`.husky`**: Contains Git hooks for the project. (Checks for formatting before committing to avoid useless conflicts.)
- **`docs`**: Documentation for the project. (All the documentation or anything team members should read for themselves can come in here.)
- **`src`**: Source code for the applications and packages in the monorepo. (All the source code for the applications and packages in the repo can be found here.)

## `src` Directory

### Client (Thembi's workspace)
- Contains the source code for the client-side applications. 
- A desktop application.

### Server (Lebo & Tlhalefo's workspaces)
- Contains the source code for the server-side applications. 

- **`api(Lebo)`**: The API for the server. (Client-server communication, authentication, etc.)
- **`core(Tlhalefo)`**: The core logic for the server. (OCR, PDF & Image processing, etc.)

### Shared (All of us as umdeni)
- **`shared(All of us as umdeni)`**: Contains code that is shared between the client and server applications.
