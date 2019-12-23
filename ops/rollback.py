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


def get_latest_version():
    alias = lambda_client.get_alias(
        Name=const['alias'],
        FunctionName=env_vars['name'],
    )
    version = alias['FunctionVersion']
    print('Latest version is ' + version)
    return version


if __name__ == "__main__":
    os.environ['AWS_DEFAULT_REGION'] = const['region']
    lambda_client = boto3.client('lambda')
    env_vars = env[argv[1]]

    version = get_latest_version()
    update_alias(str(int(version) - 1))
    delete_version(version)

    print('Done.')
