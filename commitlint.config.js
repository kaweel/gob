module.exports = {
    extends: ['@commitlint/config-conventional'],
    rules: {
        'type-enum': [
            2,
            'always',
            ['build','feat', 'fix', 'docs', 'chore', 'style', 'refactor', 'test', 'perf']
        ],
        'scope-enum': [
            2,
            'always',
            ['root', 'list','info','copy','other']
        ],
        'subject-empty': [2, 'never'],
        'type-case': [2, 'always', 'lower-case'],
        'subject-case': [0],
        'header-max-length': [2, 'always', 80]
    },
};