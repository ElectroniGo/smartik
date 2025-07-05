/**@type {import("lint-staged").Configuration} */
module.exports = {
    "*.{js,ts,json,jsonc,yaml,yml,md}": "biome check --write --no-errors-on-unmatched"
}