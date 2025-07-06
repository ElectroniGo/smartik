/**@type {import("lint-staged").Configuration} */
module.exports = {
    "*.{js,ts,json,jsonc,yaml,yml,md}": "biome check --write --files-ignore-unknown=true --no-errors-on-unmatched"
}