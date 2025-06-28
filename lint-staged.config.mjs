/**
 * @type {import('lint-staged').Configuration}
 */
const lintStagedConfig = {
    "*.{js,ts,jsx,tsx}": 'prettier --write'
}

export default lintStagedConfig;