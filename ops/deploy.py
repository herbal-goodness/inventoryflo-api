from config import const, env
from sys import argv
import zipfile as zf
import boto3
import six
import os


lambda_client = None
env_vars = None


def create_lambda():
    print('Creating function...')

    val = lambda_client.create_function(
        Publish=True,
        Runtime=const['runtime'],
        Handler=const['handler'],
        Description=const['description'],
        Timeout=const['timeout'],
        MemorySize=const['memory'],
        FunctionName=env_vars['name'],
        VpcConfig=env_vars['vpc'],
        Environment=env_vars['env-vars'],
        Role=env_vars['role'],
        Code={'ZipFile': open(const['zipfile'], 'rb').read()},
    )

    lambda_client.create_alias(
        FunctionVersion=val['Version'],
        Name=const['alias'],
        FunctionName=env_vars['name'],
    )

    return val


def update_lambda():
    print('Updating function...')

    lambda_client.update_function_configuration(
        Runtime=const['runtime'],
        Handler=const['handler'],
        Timeout=const['timeout'],
        MemorySize=const['memory'],
        FunctionName=env_vars['name'],
        Role=env_vars['role'],
        VpcConfig=env_vars['vpc'],
        Environment=env_vars['env-vars'],
    )
    val = lambda_client.update_function_code(
        Publish=True,
        FunctionName=env_vars['name'],
        ZipFile=open(const['zipfile'], 'rb').read(),
    )

    print('Updating alias to version ' + val['Version'] + '...')

    lambda_client.update_alias(
        FunctionVersion=val['Version'],
        Name=const['alias'],
        FunctionName=env_vars['name'],
    )

    return val


def lambda_exists():
    try:
        lambda_client.get_function(FunctionName=env_vars['name'])
        return True
    except:
        return False


if __name__ == "__main__":
    os.environ['AWS_DEFAULT_REGION'] = const['region']
    lambda_client = boto3.client('lambda')
    env_vars = env[argv[1]]

    if lambda_exists():
        update_lambda()
    else:
        create_lambda()

    print('Done.')
