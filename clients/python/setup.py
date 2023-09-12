# coding: utf-8

"""
    nextcloud

    Nextcloud APIs

    The version of the OpenAPI document: 0.0.1
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


from setuptools import setup, find_packages  # noqa: H301

# To install the library, run the following
#
# python setup.py install
#
# prerequisite: setuptools
# http://pypi.python.org/pypi/setuptools
NAME = "nextcloud-client-sdk"
VERSION = "1.0.0"
PYTHON_REQUIRES = ">=3.7"
REQUIRES = [
    "urllib3 >= 1.25.3, < 2.1.0",
    "python-dateutil",
    "pydantic >= 1.10.5, < 2",
    "aenum"
]

setup(
    name=NAME,
    version=VERSION,
    description="nextcloud",
    author="OpenAPI Generator community",
    author_email="team@openapitools.org",
    url="https://github.com/nextcloud/client-sdk",
    keywords=["OpenAPI", "OpenAPI-Generator", "nextcloud"],
    install_requires=REQUIRES,
    packages=find_packages(exclude=["test", "tests"]),
    include_package_data=True,
    license="agpl",
    long_description_content_type='text/markdown',
    long_description="""\
    Nextcloud APIs
    """,  # noqa: E501
    package_data={"nextcloud_client": ["py.typed"]},
)
