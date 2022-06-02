Feature: greeting end-point

Background:
* url demoBaseUrl

Scenario: say my name

    Given path '/'
    Given header Direktiv-ActionID = 'development'
    Given header Direktiv-Tempdir = '/tmp'
    And request 
    """
    { "commands": [
        {
            "command": "ssh -V"
        }
    ]}
    """"
    When method post
    Then status 200