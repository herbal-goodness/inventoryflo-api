env = {
    'qa': {
        'name': 'if-api-test',
        'role': 'arn:aws:iam::093604411390:role/if_lambda_dev',
        'vpc': {
            'SubnetIds': ['subnet-0b76b1f8b9b3b733c', 'subnet-0bbc1cc681ff5bfcb'],
            'SecurityGroupIds': ['sg-0d86c9d8f56c01491', 'sg-0ecc89d05b25dca70'],
        },
        'env-vars': {
            'Variables': {
                'env': 'qa',
                'data_key': 'AQIDAHh0DdonDMJJ3Zr8dpKyoPcWm9dlLHlh5Wy+tfrFm4vJ1wG2duQQ8ROeKvrtAGDumbR3AAAAfjB8BgkqhkiG9w0BBwagbzBtAgEAMGgGCSqGSIb3DQEHATAeBglghkgBZQMEAS4wEQQM1uzn4qmjQp8QUB1fAgEQgDvFE6DpF9gufiZ0iFBc6T4JrDidd75aYuuda7vWE6eenq8+3wt9Y40kzs9Q2PLftimO0+ooJCr9aMjPCg==',
            }
        }
    },
    'prod': {
        'name': 'if-api-prod',
        'role': 'todo',
        'vpc': {
            'SubnetIds': ['todo'],
            'SecurityGroupIds': ['todo'],
        },
        'env-vars': {
            'Variables': {
                'env': 'prod',
                'data_key': 'AQIDAHidvvXkO7IeRBLTpKSW5mDR5NMz/5LJM3jM/ssNl34EBgHx4e+GMrXDz6u1hBlmlHx6AAAAfjB8BgkqhkiG9w0BBwagbzBtAgEAMGgGCSqGSIb3DQEHATAeBglghkgBZQMEAS4wEQQMbpnR0Xc+WGECxL/LAgEQgDvn33+YUMDcaO0hnmKLZL0BiyNph8QruaVjqXxAC69O86WRWqUUD2lE6Zq1IeqFa0Fo0DzfIpiPrYBP0w==',
            }
        }
    },
}

const = {
    'runtime': 'go1.x',
    'handler': 'main',
    'description': 'RESTful API layer of InventoryFlo',
    'timeout': 60,
    'memory': 1024,
    'region': 'us-east-2',
    'alias': 'active',
    'zipfile': './main.zip'
}
