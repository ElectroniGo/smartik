# `@repo/biome-config`

This repo holds the shared configs for biome.

We're using this instead of having to install `eslint` & `prettier`. Biome does what both do **and** faster

## Usage

1. Add it to all JavaScript/TypeScript apps & packages in this project like this:

```json
// package.json
{
    // ...other configs
    "devDependencies": {
        "@repo/biome-config": "workspace:*"
    }
    // ...rest of configs
}
```

Then install it in the package

> Run this in the root of the app/package you are working on
> 
> _in the same folder where your package's package.json file is_

```bash
pnpm install
```

2. Extend the base config

```json
// biome.json
{
    "extends": "@repo/biome-config/base.json",
    // ...rest of your config overwrites and extensions if any
}
```