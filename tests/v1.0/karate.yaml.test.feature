
Feature: Basic

# The secrects can be used in the payload with the following syntax #(mysecretname)
Background:
* def sshSecret = karate.properties['sshSecret']


Scenario: get request

	Given url karate.properties['testURL']

	And path '/'
	And header Direktiv-ActionID = 'development'
	And header Direktiv-TempDir = '/tmp'
	And request
	"""
	{	
		"host": {
			"name": "10.100.6.16",
			"port": 2022
		},
		"auth": {
			"username": "direktiv",
			"certificate": "/cert",
			"password": "direktiv"
		},
		"commands": [
		{
			"command": "date",
			"silent": false,
			"print": true,
		}
		]
	}
	"""
	When method POST
	Then status 200
	And match $ ==
	"""
	{
	"ssh": [
	{
		"result": "#notnull",
		"success": true
	}
	]
	}
	"""
	