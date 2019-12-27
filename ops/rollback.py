from config import const, env
from sys import argv
import zipfile as zf
import boto3
import six
import os


lambda_client = None
env_vars = None


def delete_version(version):
    print('Deleting version ' + version)
    return lambda_client.delete_function(
        FunctionName=env_vars['name'],
        Qualifier=version,
    )


def update_alias(version):
    print('Updating alias to point to version ' + version)
    return lambda_client.update_alias(
        Name=const['alias'],
        FunctionName=env_vars['name'],
        FunctionVersion=version,
    )


def get_latest_versions():
    versions = lambda_client.list_versions_by_function(
        FunctionName=env_vars['name'],
    )['Versions']
    versions = list(filter(lambda x: x['Version'] != '$LATEST', versions))
    versions.sort(key=lambda x: int(x['Version']), reverse=True)
    return (versions[0]['Version'], versions[1]['Version'])


if __name__ == "__main__":
    os.environ['AWS_DEFAULT_REGION'] = const['region']
    lambda_client = boto3.client('lambda')
    env_vars = env[argv[1]]

    cur, prev = get_latest_versions()
    update_alias(prev)
    delete_version(cur)

    print('Done.')
